package main

import (
	"crud-go/initializers"
	"crud-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}
