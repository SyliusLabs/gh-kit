package cmd

import (
	"fmt"
	"github.com/jakubtobiasz/gh-kit/internal/github/client"
	"github.com/jakubtobiasz/gh-kit/internal/github/client/issue"
	"github.com/jakubtobiasz/gh-kit/internal/github/current_repository"
	"github.com/jakubtobiasz/gh-kit/internal/github/pull_request"
	"github.com/spf13/cobra"
)

var allowedMergeStrategies = map[string]bool{"merge": true, "squash": true, "rebase": true}

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge a pull request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo, repoErr := current_repository.GetCurrentRepositoryData()
		if nil != repoErr {
			fmt.Println("The current directory is not a git repository")
			return
		}

		pr, prErr := client.GetPullRequestData(repo, args[0])
		if nil != prErr {
			fmt.Println("The pull request does not exist")
			return
		}

		if pr.GetMerged() {
			fmt.Println(fmt.Sprintf("The pull request #%d on %s/%s is already merged by %s", pr.GetNumber(), repo.Owner, repo.Name, pr.GetMergedBy().GetLogin()))
			return
		}

		mergeStrategy, _ := cmd.Flags().GetString("strategy")
		if !allowedMergeStrategies[mergeStrategy] {
			fmt.Println("The merge strategy is not allowed. Allowed strategies are: merge, squash, rebase")
			return
		}

		mergeCategory, _ := cmd.Flags().GetString("category")
		mergeSubject := pull_request.GenerateSubjectWithCategory(mergeCategory, pr.GetNumber(), pr.GetTitle(), pr.GetUser().GetLogin())

		mergeErr := pull_request.Merge(pr.GetNumber(), mergeSubject, mergeStrategy)

		if nil != mergeErr {
			fmt.Println(fmt.Sprintf("The pull request #%d could not be merged. Error message: %s", pr.GetNumber(), mergeErr.Error()))
			return
		}

		fmt.Println(fmt.Sprintf("✅ The pull request #%d on %s/%s has been merged", pr.GetNumber(), repo.Owner, repo.Name))

		commentErr := issue.AddComment(repo, pr.GetNumber(), fmt.Sprintf("Thank you, @%s!", pr.GetUser().GetLogin()))
		if nil != commentErr {
			fmt.Println(fmt.Sprintf("The comment could not be added. Error message: %s", commentErr.Error()))
			return
		}

		fmt.Println(fmt.Sprintf("✅ The thank you comment has been added to the pull request #%d", pr.GetNumber()))
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	mergeCmd.PersistentFlags().StringP("category", "c", "minor", "Category of the pull request")
	mergeCmd.PersistentFlags().StringP("strategy", "s", "merge", "Merge strategy to be used")
}
