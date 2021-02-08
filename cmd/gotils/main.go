package main

import (
	"log"
	"os"
)

func init() {
	githubCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(githubCmd)

	commitCmd.PersistentFlags().StringVar(&githubUser, "user", "u", "user name")
	commitCmd.PersistentFlags().StringVar(&githubRepo, "repo", "r", "repo name")
	commitCmd.PersistentFlags().StringVar(&keywords, "keywords", "", "filtering keywords with command separated")
	commitCmd.PersistentFlags().IntVar(&maxCommits, "maxcommits", 10, "max read commits(default: 10)")
}

func main() {
	//os.Args = append(os.Args, "github", "commit", "--user", "zacscoding", "--repo", "go-workspace",
	//	"--keywords", "gorm:")
	//	os.Args = append(os.Args, "github")
	if err := rootCmd.Execute(); err != nil {
		log.Printf("failed to execute command. err: %v", err)
		os.Exit(1)
	}
}
