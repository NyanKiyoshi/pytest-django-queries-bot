package integration

import (
	"bytes"
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/upstream"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/utils"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path"
	"text/template"
)

const mdCodeBlock = "```"

var commentTemplate = template.Must((&template.Template{}).Parse(`
Here is the report for {{ .HeadSHA }} ({{ .HeadLabel }})
{{ if .BaseReportMissing -}}
Missing base report ({{ .BaseSHA }}). The results couldn't be compared.
{{ else -}}
Base comparison is {{ .BaseSHA }}.
{{ end -}}

<details>
<summary>
{{- if eq .DiffCount 0 -}}
	No differences were found.
{{- else -}}
	<b>Found {{ .DiffCount }} differences!</b> (click me)
{{- end -}}
</summary>
<p>

` + mdCodeBlock + `diff
{{ printf "%s" .RawDiff }}
` + mdCodeBlock + `

</p>
</details>
`))

func writeBaseReport(targetPath string, body []byte) (err error) {
	if err = ioutil.WriteFile(targetPath, body, 0660); err != nil {
		err = fmt.Errorf("failed to write base report to '%s': %w", targetPath, err)
	}
	return
}

func readHeadReport(reportPath string) ([]byte, error) {
	return ioutil.ReadFile(reportPath)
}

type PullRequestHandler struct {
	event                     github.PullRequestEvent
	HeadReportLocation        string
	RetrieveQueriesPayloadUrl *url.URL
	PostDiffCommentUrl        *url.URL
	Logger                    *logrus.Logger

	isBaseMissing      bool
	baseReportLocation string
}

type Context struct {
	HeadSHA           string
	BaseSHA           string
	BaseReportMissing bool
	DiffCount         uint
	RawDiff           []byte
	HeadLabel string
}

func getDiffCount(diff []byte) uint {
	var diffCount uint

	for _, l := range bytes.Split(diff, []byte{'\n'}) {
		if len(l) < 3 {
			continue
		}
		if l[1] == ' ' && (l[0] == '-' || l[0] == '+') {
			diffCount++
		}
	}

	return diffCount
}

func (h *PullRequestHandler) GetEventPtr() interface{} {
	return &h.event
}

func (h *PullRequestHandler) retrieveBaseReport(baseSHA string) (err error) {
	reqUrl := h.RetrieveQueriesPayloadUrl
	reqUrl.Path = path.Join(reqUrl.Path, baseSHA)
	h.Logger.Infof("GET '%s'", reqUrl)
	resp, err := utils.HttpDo("GET", reqUrl.String(), nil, nil)

	if err != nil {
		err = fmt.Errorf("failed to retrieve base report: %w", err)
		return
	}

	if resp == nil {
		err = fmt.Errorf("received no response when getting the base report")
		return
	}

	if resp.StatusCode != 200 {
		h.Logger.Infof("Skipping base report, got status: %d", resp.StatusCode)
		head, err := readHeadReport(h.HeadReportLocation)
		if err != nil {
			err = fmt.Errorf("failed to read head report to '%s': %w", h.HeadReportLocation, err)
			return err
		}
		h.isBaseMissing = true
		return writeBaseReport(h.baseReportLocation, head)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read body of base report: %w", err)
		return
	}
	h.isBaseMissing = false
	return writeBaseReport(h.baseReportLocation, body)
}

func (h *PullRequestHandler) getDiff() (diff []byte, err error) {
	cmd := exec.Command(
		"python3",
		"-Im",
		"pytest_django_queries.cli",
		"diff", h.baseReportLocation, h.HeadReportLocation,
	)
	cmd.Stderr = os.Stderr
	diff, err = cmd.Output()
	return
}

func (h *PullRequestHandler) createCommentBody(ctx *Context) (comment []byte, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = commentTemplate.Execute(buf, ctx)
	comment = buf.Bytes()
	return
}

func (h *PullRequestHandler) Invoke() (err error) {
	event := h.event
	base := event.PullRequest.Base.SHA
	head := event.PullRequest.Head.SHA

	h.baseReportLocation = "./base-report.json"
	if err = h.retrieveBaseReport(*base); err != nil {
		err = fmt.Errorf("failed to retrieve base report: %w", err)
		return
	}

	diff, err := h.getDiff()
	if err != nil {
		err = fmt.Errorf(
			"failed to generate diff between '%s' and '%s': %w", h.baseReportLocation, h.HeadReportLocation, err)
	}

	body, err := h.createCommentBody(&Context{
		HeadSHA:           *head,
		HeadLabel: *event.PullRequest.Head.Label,
		BaseSHA:           *base,
		BaseReportMissing: h.isBaseMissing,
		DiffCount:         getDiffCount(diff),
		RawDiff:           diff,
	})
	if err != nil {
		err = fmt.Errorf("failed to generate comment: %w", err)
		return
	}
	h.Logger.Infof("Generated Comment: %s", body)

	input := upstream.DiffInput{
		PostCommentUrl: h.PostDiffCommentUrl.String(),
		SHA1Revision:   *head,
	}
	if _, err = input.PostFromReader(bytes.NewReader(body)); err != nil {
		err = fmt.Errorf("failed to write diff comment from stdin: %w", err)
	}
	return
}
