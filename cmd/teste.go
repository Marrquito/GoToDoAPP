package cmd

import (
	"fmt"

	"github.com/Marrquito/GoToDoAPP/cmd/ui/multiInput"
	"github.com/Marrquito/GoToDoAPP/cmd/ui/textInput"
	"github.com/Marrquito/GoToDoAPP/cmd/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// testeCmd represents the teste command
var testeCmd = &cobra.Command{
	Use:   "teste",
	Short: "test",
	Long: `Test some options of bubbletea`,
	Run: func(cmd *cobra.Command, args []string) {
		var output utils.Output_str
		
		tprogram2 := tea.NewProgram(textInput.InitModel(&output, "Qual seu é nome?"), tea.WithAltScreen())
		if _, err := tprogram2.Run(); err != nil {
			fmt.Println(err)
		}
		
		tprogram := tea.NewProgram(multiInput.InitModel(), tea.WithAltScreen())
		if _, err := tprogram.Run(); err != nil {
			fmt.Println(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(testeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
