package cmd

import (
	"fmt"

	listSimple "github.com/Marrquito/GoToDoAPP/cmd/ui/list_simple"
	"github.com/Marrquito/GoToDoAPP/cmd/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// testeCmd represents the teste command
var testeCmd = &cobra.Command{
	Use:   "teste",
	Short: "test",
	Long: `Test some options of bubbletea`,
	Run: func(cmd *cobra.Command, args []string) {
		var output utils.Output
		// tprogram := tea.NewProgram(textInput.InitModel(&output, "Qual seu é nome?"), tea.WithAltScreen())

		// if _, err := tprogram.Run(); err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println("output = ", output.Output)
		
		
		items := []list.Item{
			listSimple.Item("OPA"),
			listSimple.Item("OPA2"),
			listSimple.Item("OPA3"),
		}

		tprogram := tea.NewProgram(listSimple.InitModel(&output, items, "Meu titulo"))
		if _, err := tprogram.Run(); err != nil {
			fmt.Println(err)
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
