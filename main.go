package main

import (
	"fmt"
	"github.com/SyliusLabs/gh-kit/cmd"
	"github.com/SyliusLabs/gh-kit/internal/github"
	"go.uber.org/fx"
	"os"
)

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				cmd.NewRootCmd,
				fx.ParamTags(`group:"commands"`),
			),
			asCommand(cmd.NewPrCmd, "pr"),
			asSubCommand(cmd.NewPrMergeCmd, "pr"),
			asSubCommand(cmd.NewPrRerunCmd, "pr"),
			asCommand(cmd.NewUpmergeCmd, "upmerge"),
			asSubCommand(cmd.NewUpmergeCreateCmd, "upmerge"),
			asSubCommand(cmd.NewUpmergeMergeCmd, "upmerge"),
			github.NewCli,
			github.NewClient,
			github.NewRestClient,
			github.NewRepository,
		),
		fx.Invoke(func(rootCmd *cmd.RootCmd) {
			err := rootCmd.Execute()

			if err != nil {
				os.Exit(1)
			}

			os.Exit(0)
		}),
		fx.NopLogger,
	).Run()
}

func asCommand(f any, name string) any {
	return fx.Annotate(
		f,
		fx.As(new(cmd.Command)),
		fx.ResultTags(`group:"commands"`),
		fx.ParamTags(fmt.Sprintf(`group:"%s_subcommands"`, name)),
	)
}

func asSubCommand(f any, parentName string) any {
	return fx.Annotate(
		f,
		fx.As(new(cmd.Command)),
		fx.ResultTags(fmt.Sprintf(`group:"%s_subcommands"`, parentName)),
	)
}
