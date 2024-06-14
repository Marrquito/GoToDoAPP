/*
Copyright © 2024 Marco Antonio Cunha Cossetin marcossetin@hotmail.com
*/
package main

import (
	"github.com/Marrquito/GoToDoAPP/cmd"
	"github.com/Marrquito/GoToDoAPP/common/logger"
	"github.com/Marrquito/GoToDoAPP/config"
	"github.com/charmbracelet/lipgloss"

	tm "github.com/buger/goterm"
)

const logo = `
##   ##   #####   #####     #####   ##   ##   ######  ##########  #######
### ###  #######  #######  #######  ##   ##   ######  ##########  #######
#######  ##   ##       ##  ##   ##  ##   ##     ##        ##	  ##   ##
#######  ##   ##  ######   ##   ##  ##   ##     ##        ##	  ##   ##
## # ##  #######  ##   ##  ##   ##  ##   ##     ##        ##	  ##   ##
##   ##  ##   ##  ##   ##  #######  #######   ######      ##	  #######
##   ##  ##   ##  ##   ##   #### ##  #####    ######      ##	   #####
` 

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

func main() {
	logger := logger.NewLogger(config.InfoLevel, "Server")
	logger.Info("Starting ToDo")

	tm.Clear() // Clear current screen
	tm.MoveCursor(1, 1)
	tm.Println(logoStyle.Render(logo))
	tm.Flush()

	cmd.Execute()
}
