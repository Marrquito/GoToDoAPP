/*
Copyright © 2024 Marco Antonio Cunha Cossetin marcossetin@hotmail.com
*/
package main

import (
	"github.com/Marrquito/GoToDoAPP/cmd"
	"github.com/Marrquito/GoToDoAPP/common/logger"
)

func main() {
	cmd.Execute()
	
	logger := logger.NewLogger(2, "Server")
	logger.Info(	"DEU BOM!! INFO")
	logger.Warning(	"DEU BOM!! WARNING")
	logger.Error(	"DEU BOM!! ERROR")
	logger.Debug(	"DEU BOM!! DEBUG")
}
