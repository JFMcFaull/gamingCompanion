package main

import (
	"github.com/JFMcFaull/gamingCompanion/initalizers"
	"github.com/JFMcFaull/gamingCompanion/models"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {
	initalizers.DB.AutoMigrate(&models.Game{})
}
