package cmd

//
import (
	"fmt"
	"github.com/SyliusLabs/gh-kit/internal/generator"
	"github.com/SyliusLabs/gh-kit/internal/githubcli"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strconv"
)

var allowedMergeStrategies = map[string]bool{"merge": true, "squash": true, "rebase": true}

var mergeCmd = &cobra.Command{
	Use:   "merge <pull_request_number>",
	Short: "Merge a pull request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prNumber, err := strconv.Atoi(args[0])
		if nil != err {
			fmt.Println("🛑 The pull request number is not a valid integer")
			return
		}
		pr, prErr := ghClient.GetPullRequest(prNumber)
		if nil != prErr {
			fmt.Println(prErr.Error())
			return
		}

		if pr.GetMerged() {
			fmt.Printf(
				"ℹ️ The pull request #%d on %s/%s is already merged by %s\r\n",
				pr.GetNumber(),
				ghClient.Repository.Owner(),
				ghClient.Repository.Name(),
				pr.GetMergedBy().GetLogin(),
			)
			return
		}

		mergeStrategy, _ := cmd.Flags().GetString("strategy")

		if "" == mergeStrategy {
			strategyPrompt := promptui.Select{
				Label: "Select merge strategy",
				Items: []string{"merge", "squash", "rebase"},
			}

			_, mergeStrategy, _ = strategyPrompt.Run()
		}

		if !allowedMergeStrategies[mergeStrategy] || "" == mergeStrategy {
			fmt.Println("🛑 The merge strategy is not allowed or is empty. Allowed strategies are: merge, squash, rebase")
			return
		}

		mergeCategory, _ := cmd.Flags().GetString("category")

		if "" == mergeCategory {
			categoryPrompt := promptui.Select{
				Label: "Select merge category",
				Items: []string{"minor", "feature", "bugfix", "docs", "refactor", "test", "chore"},
			}

			_, mergeCategory, _ = categoryPrompt.Run()
		}

		if "" == mergeCategory {
			fmt.Println("⚠️ The Pull Request category must be provided.")
			return
		}

		mergeSubject := generator.GenerateCommitMessage(mergeCategory, pr.GetNumber(), pr.GetTitle(), pr.GetUser().GetLogin())

		commits, commitsErr := ghClient.GetCommitsInPullRequest(pr.GetNumber())
		if nil != commitsErr {
			fmt.Println("🛑 The pull request commits could not be fetched")
			return
		}

		mergeBody := generator.GenerateCommitBody(pr, commits)

		summaryMsg := `The pull request #%d on %s/%s will be merged with the following parameters:
- Subject: %s
- Strategy: %s
- Category: %s
- Body: %s`
		fmt.Printf(summaryMsg, pr.GetNumber(), ghClient.Repository.Owner(), ghClient.Repository.Name(), mergeSubject, mergeStrategy, mergeCategory, mergeBody)

		confirmPrompt := promptui.Select{
			Label: "Are you sure you want to merge it?",
			Items: []string{"yes", "no"},
		}

		_, confirmed, _ := confirmPrompt.Run()

		if "" == confirmed || "no" == confirmed {
			fmt.Println("🛑 The pull request has not been merged")
			return
		}

		mergeErr := githubcli.Merge(ghCli, pr.GetNumber(), mergeSubject, mergeBody, mergeStrategy)

		if nil != mergeErr {
			fmt.Printf("🛑 The pull request #%d could not be merged. Error message: %s\r\n", pr.GetNumber(), mergeErr.Error())
			return
		}

		fmt.Printf("✅ The pull request #%d on %s/%s has been merged\r\n", pr.GetNumber(), ghClient.Repository.Owner(), ghClient.Repository.Name())

		if noThankYou, _ := cmd.Flags().GetBool("skip-thankyou"); noThankYou {
			fmt.Println("⏩ The thank you comment has not been added")
			return
		}

		commentErr := ghClient.AddCommentToIssue(pr.GetNumber(), fmt.Sprintf("Thank you, @%s!", pr.GetUser().GetLogin()))
		if nil != commentErr {
			fmt.Printf("⚠️ The comment could not be added. Error message: %s\r\n", commentErr.Error())
			return
		}

		fmt.Printf("✅ The thank you comment has been added to the pull request #%d\r\n", pr.GetNumber())
	},
}

func init() {
	prCmd.AddCommand(mergeCmd)

	mergeCmd.PersistentFlags().StringP("category", "c", "", "Category of the pull request")
	mergeCmd.PersistentFlags().StringP("strategy", "s", "", "Merge strategy to be used")
	mergeCmd.PersistentFlags().Bool("skip-thankyou", false, "Do not add a thank you comment to the pull request")
}
