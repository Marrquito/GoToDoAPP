package cmd

import (
	"fmt"

	listSimple "github.com/Marrquito/GoToDoAPP/cmd/ui/list_simple"
	"github.com/Marrquito/GoToDoAPP/cmd/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts Application",
	Long:  `Starts the CLI Application`,
	Run: func(cmd *cobra.Command, args []string) {
		items := []list.Item{
			listSimple.Item("Nova tarefa"),
			listSimple.Item("Listar tarefas"),
			listSimple.Item("Atualizar tarefa"),
			listSimple.Item("Deletar tarefa"),
		}

		var output utils.Output
		tprogram := tea.NewProgram(listSimple.InitModel(&output, items, "O que deseja fazer?"))
		if _, err := tprogram.Run(); err != nil {
			fmt.Println(err)
		}

		switch output.Output {
		case "Nova tarefa":
			fmt.Println("Create")
		case "Listar tarefas":
			fmt.Println("List")
		case "Atualizar tarefa":
			fmt.Println("Update")
		case "Deletar tarefa":
			fmt.Println("Delete")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
