/*
Copyright © 2024 Marco Antonio Cunha Cossetin marcossetin@hotmail.com
*/
package main

import (
	"github.com/Marrquito/GoToDoAPP/cmd"
	"github.com/Marrquito/GoToDoAPP/common/logger"
	"github.com/Marrquito/GoToDoAPP/config"
)

func main() {
	cmd.Execute()
	
	logger := logger.NewLogger(config.InfoLevel, "Server")
	logger.Info("Starting ToDo")
}
