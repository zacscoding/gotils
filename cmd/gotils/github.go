package main

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/zacscoding/gotils/pkg/git"
	"os"
	"strings"
)

var (
	githubUser string
	githubRepo string
	keywords   string
	maxCommits int

	rootCmd = &cobra.Command{
		Use:   "gotils [sub]",
		Short: "My Commands :)",
	}
)

var (
	githubCmd = &cobra.Command{
		Use:  "github",
		Long: "Github Utility Commands",
	}
	commitCmd = &cobra.Command{
		Use: "commit",
		Run: func(cmd *cobra.Command, args []string) {
			opt := &git.SearchOpt{
				User:       githubUser,
				Repo:       githubRepo,
				Keywords:   strings.Split(keywords, ","),
				MaxCommits: maxCommits,
			}
			writer := table.NewWriter()
			writer.SetOutputMirror(os.Stdout)
			_ = git.SearchGithubCommitMessages(opt, writer)
		},
	}
)
