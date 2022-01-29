package main

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	cobraprompt "github.com/stromland/cobra-prompt"
	"github.com/stromland/cobra-prompt/_example/cmd"
)

func toSuggest(list [][]string) []prompt.Suggest {
	var suggestions []prompt.Suggest
	for _, s := range list {
		if len(s) == 0 {
			continue
		}
		if len(s) == 1 {
			suggestions = append(suggestions, prompt.Suggest{Text: s[0]})
		}
		if len(s) == 2 {
			suggestions = append(suggestions, prompt.Suggest{Text: s[0], Description: s[1]})
		}
	}
	return suggestions
}

var advancedOption = cobraprompt.CobraPromptOptions{
	PersistFlagValues:        true,
	DisableCompletionCommand: true,
	AddDefaultExitCommand:    true,
	GoPromptOptions: []prompt.Option{
		prompt.OptionTitle("cobra-prompt"),
		prompt.OptionPrefix(">(^!^)> "),
		prompt.OptionMaxSuggestion(10),
	},
	FindSuggestionsOptions: cobraprompt.FindSuggestionsOptions{
		ShowHelpCommandAndFlags: true,
		CustomSuggestionsFunc: func(c *cobra.Command, document *prompt.Document) []prompt.Suggest {
			if suggestions := cmd.GetFoodDynamic(c); suggestions != nil {
				return toSuggest(suggestions)
			}
			return []prompt.Suggest{}
		},
	},
	OnErrorFunc: func(err error) {
		if strings.Contains(err.Error(), "unknown command") {
			cmd.RootCmd.PrintErrln(err)
			return
		}

		cmd.RootCmd.PrintErr(err)
		os.Exit(1)
	},
}

var simpleOption = cobraprompt.CobraPromptOptions{
	AddDefaultExitCommand:    true,
	DisableCompletionCommand: true,
}

func main() {
	// Change to simpleOptions to see the difference
	cobraprompt.New(*cmd.RootCmd, advancedOption).Run()
}
