package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/nachoconques0/quizzer/internal/model"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quizzer",
	Short: "a lazy and easy quiz",
	Run: func(cmd *cobra.Command, args []string) {
		initQuiz(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initQuiz(cmd *cobra.Command, args []string) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F355 {{ .Label | red }}",
		Inactive: "{{ .Label | cyan }}",
		Selected: "\U0001F355 {{ .Label | red }}",
	}

	prompt := promptui.Select{
		Templates: templates,
		Label:     "Welcome to QUIZZZZZER",
		Items: []model.Item{
			{
				Label:  "Start",
				Answer: 0,
			},
			{
				Label:  "Exit",
				Answer: 1,
			},
		},
	}

	idx, _, err := prompt.Run()
	switch idx {
	case 0:
		go startApiCmd.Run(cmd, os.Args)
		startQuizCmd.Run(cmd, args)
	default:
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
}
