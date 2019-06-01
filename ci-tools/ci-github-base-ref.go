// +build base-ref

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"log"
	"os"
	"pytest-queries-bot/pytestqueries/ci-tools/utils"
)

type gqlQuery struct {
	Query     string       `json:"query"`
	Variables gqlQueryData `json:"variables"`
}

type gqlQueryData struct {
	User    string `short:"u" long:"user" description:"The repository owner's name'" required:"true" json:"owner"`
	Repo    string `short:"n" long:"name" description:"The name of the repository." required:"true" json:"repo"`
	RefName string `short:"r" long:"ref" description:"The reference name (branch name)." required:"true" json:"refName"`
}

type gqlResponseData struct {
	Data struct {
		Repository struct {
			Ref struct {
				AssociatedPullRequests struct {
					Edges []struct {
						Node struct {
							BaseRefOid string `json:"baseRefOid"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"associatedPullRequests"`
			} `json:"ref"`
		} `json:"repository"`
	} `json:"data"`
}

// The v4 API endpoint
const githubAPIEndpoint = "https://api.github.com/graphql"

// userAgent is mandatory to make requests in the GitHub APIs.
// In case of issues, GitHub must be able to contact the owner/maintainer.
const userAgent = "NyanKiyoshi/pytest-django-queries-bot"

const getBaseCommitRefGQLQuery = `
query($owner: String!, $repo: String!, $refName: String!) {
  repository(owner: $owner, name: $repo) {
    ref(qualifiedName: $refName) {
      associatedPullRequests(last: 1) {
        edges {
          node {
            baseRefOid
          }
        }
      }
    }
  }
}
`

func main() {
	data := gqlQueryData{}
	if _, err := flags.Parse(&data); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	query, err := json.Marshal(gqlQuery{
		Query:     getBaseCommitRefGQLQuery,
		Variables: data,
	})

	bytes.NewReader(query)

	if err != nil {
		log.Fatalf("Failed to compile the request: %s", err)
	}

	resp := utils.SendUploadRequest(
		githubAPIEndpoint, "application/json", bytes.NewReader(query), &map[string]string{
			"User-Agent":    userAgent,
			"Authorization": fmt.Sprintf("bearer %s", os.Getenv("GITHUB_GQL_TOKEN")),
		},
	)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Expected HTTP 200 but go %d instead", resp.StatusCode)
	}

	jsonResp := gqlResponseData{}
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		log.Fatalf("Failed to decode response: %s", err)
	}

	edges := jsonResp.Data.Repository.Ref.AssociatedPullRequests.Edges
	if len(edges) == 0 {
		log.Fatalf("Did not find any pull request edges")
	}

	if _, err := os.Stdout.WriteString(edges[0].Node.BaseRefOid); err != nil {
		panic(err)
	}
}
