package integration

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getDiffCount(t *testing.T) {
	const (
		reportWithoutDiffs = `
# saleor.graphql.accountbenchmark permission group
test name                                                     left count      right count     duplicate count
-----------------------------------------------------------   -----------     -----------     ---------------
  degraded func                  15              15                   0
  improved func                  20              20                   0
unchanged func                  1               1                   0
`
		reportWithDiffs = `
# saleor.graphql.accountbenchmark permission group
test name                                                     left count      right count     duplicate count
-----------------------------------------------------------   -----------     -----------     ---------------
- degraded func                  15              15                   0
+ improved func                  20              20                   0
unchanged func                  1               1                   0
`
	)

	type args struct {
		diff []byte
	}
	tests := []struct {
		name string
		want uint
		args args
	}{
		{
			name: "No diffs",
			want: 0,
			args: args{diff: []byte(reportWithoutDiffs)},
		},
		{
			name: "Some diffs",
			want: 2,
			args: args{diff: []byte(reportWithDiffs)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiffCount(tt.args.diff); got != tt.want {
				t.Errorf("getDiffCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplate(t *testing.T) {
	expected := `
Here is the report for 6db0fab78656207b1040cc6eae4b41cc1ab4b736
Missing base report (f6d1e2b4b06b7ac7a40783b6879f66840cf9e75d). The results couldn't be compared.<details>
<summary><b>Found 3 differences!</b> (click me)</summary>
<p>

` + "```" + `diff
foo
bar
` + "```" + `

</p>
</details>
`
	ctx := Context{
		HeadSHA:           "6db0fab78656207b1040cc6eae4b41cc1ab4b736",
		BaseSHA:           "f6d1e2b4b06b7ac7a40783b6879f66840cf9e75d",
		BaseReportMissing: true,
		DiffCount:         3,
		RawDiff:           []byte("foo\nbar"),
	}
	tpl := commentTemplate
	w := bytes.NewBufferString("")
	err := tpl.Execute(w, ctx)
	assert.NoError(t, err)
	assert.Equal(t, expected, w.String())
}
