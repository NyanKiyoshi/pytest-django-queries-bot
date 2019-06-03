package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type gqlQuery struct {
	Query     string       `json:"query"`
	Variables GqlQueryData `json:"variables"`
}

type GqlQueryData struct {
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

func GetBaseRef(data GqlQueryData) (string, error) {
	query, err := json.Marshal(gqlQuery{
		Query:     getBaseCommitRefGQLQuery,
		Variables: data,
	})

	log.Printf("Requesting: %s", query)
	bytes.NewReader(query)

	if err != nil {
		return "", fmt.Errorf("failed to compile the request: %s", err)
	}

	resp, err := SendUploadRequest(
		githubAPIEndpoint, "application/json", bytes.NewReader(query), &map[string]string{
			"User-Agent":    userAgent,
			"Authorization": fmt.Sprintf("bearer %s", os.Getenv("GITHUB_GQL_TOKEN")),
		},
	)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("Received: %s", body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("expected HTTP 200 but go %d instead", resp.StatusCode)
	}

	jsonResp := gqlResponseData{}
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %s", err)
	}

	edges := jsonResp.Data.Repository.Ref.AssociatedPullRequests.Edges
	if len(edges) == 0 {
		return "", errors.New("did not find any pull request edges")
	}

	return edges[0].Node.BaseRefOid, nil
}
