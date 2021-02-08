package git

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"net/http"
	"strings"
)

type SearchOpt struct {
	User       string
	Repo       string
	Keywords   []string
	MaxCommits int
}

func (opt *SearchOpt) includeCommits(msg string) bool {
	for _, word := range opt.Keywords {
		if strings.Contains(msg, word) {
			return true
		}
	}
	return false
}

// SearchGithubCommitMessages read commit messages from remote github repository
// and filters commit message with given keywords
func SearchGithubCommitMessages(opt *SearchOpt, writer table.Writer) error {
	var (
		ctx         = context.Background()
		client      = github.NewClient(nil)
		page        = 0
		perPage     = 100
		remain      = opt.MaxCommits
		readCommits = 0
		idx         = 1
		continueReq = true
	)
	if opt.MaxCommits > 0 {
		perPage = opt.MaxCommits
	}
	// write header
	writer.Style().Format.Header = text.FormatDefault
	writer.Style().Format.Footer = text.FormatDefault
	writer.SetTitle("Matched Commits from /%s/%s. Keywords:%s", opt.User, opt.Repo, strings.Join(opt.Keywords, ","))
	writer.AppendHeader(table.Row{"#", "Message", "URL"})

	for {
		if !continueReq {
			break
		}
		commits, response, err := client.Repositories.ListCommits(ctx, opt.User, opt.Repo,
			&github.CommitsListOptions{
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: perPage,
				},
			},
		)
		if err != nil {
			return err
		}
		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to fetch commits. status code:%d", response.StatusCode)
		}
		for _, commit := range commits {
			if opt.includeCommits(strings.ToLower(commit.GetCommit().GetMessage())) {
				writer.AppendRow(table.Row{
					idx, commit.GetCommit().GetMessage(), fmt.Sprintf("https://github.com/%s/%s/commit/%s", opt.User, opt.Repo, commit.GetSHA()),
				})
				writer.AppendSeparator()
				idx++
				remain--
			}
			if opt.MaxCommits > 0 && remain < 0 {
				continueReq = false
				break
			}
		}
		readCommits += len(commits)
		if response.NextPage == 0 {
			continueReq = false
		}
		page = response.NextPage
	}
	writer.AppendFooter(table.Row{"", "", fmt.Sprintf("Total Read Commits: #%d", readCommits)})
	writer.Render()
	return nil
}
