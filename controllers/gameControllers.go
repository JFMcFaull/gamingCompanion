package controllers

import (
	"github.com/JFMcFaull/gamingCompanion/initalizers"
	"github.com/JFMcFaull/gamingCompanion/models"
	"github.com/gin-gonic/gin"
)

func GamesCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Title                 string
		Franchise             string
		Platform              string
		MainQuests            int
		SideQuests            int
		CompletedMainQuests   int
		CompletedSideQuests   int
		Collectibles          int
		CollectedCollectibles int
		CurrentlyPlaying      string
		GameGuide             string
		GameMap               string
	}

	c.Bind(&body)

	//create a Game

	game := models.Game{Title: body.Title, Franchise: body.Franchise, Platform: body.Platform, MainQuests: body.MainQuests, SideQuests: body.SideQuests, CompletedMainQuests: body.CompletedMainQuests, CompletedSideQuests: body.CompletedSideQuests, Collectibles: body.Collectibles, CollectedCollectibles: body.CollectedCollectibles, CurrentlyPlaying: body.CurrentlyPlaying, GameGuide: body.GameGuide, GameMap: body.GameMap}

	result := initalizers.DB.Create(&game)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it

	c.JSON(200, gin.H{
		"game": game,
	})
}

func GamesIndex(c *gin.Context) {
	// Get the posts
	var games []models.Game
	initalizers.DB.Find(&games)

	// Respond with them
	c.JSON(200, gin.H{
		"games": games,
	})

}

func GamesShow(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Get the game
	var game models.Game
	result := initalizers.DB.First(&game, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Game not found"})
		return
	}

	// Respond with the game details
	c.JSON(200, gin.H{
		"game": game,
	})
}

func GamesUpdate(c *gin.Context) {
	//get the ID from the URL
	id := c.Param("id")

	//get the data off request body
	var body struct {
		Title                 string
		Franchise             string
		Platform              string
		MainQuests            int
		SideQuests            int
		CompletedMainQuests   int
		CompletedSideQuests   int
		Collectibles          int
		CollectedCollectibles int
		CurrentlyPlaying      string
		GameGuide             string
		GameMap               string
	}

	c.Bind(&body)

	//find post being updated
	var game models.Game
	initalizers.DB.First(&game, id)

	//update it
	initalizers.DB.Model(&game).Updates(models.Game{
		Title:                 body.Title,
		Franchise:             body.Franchise,
		Platform:              body.Platform,
		MainQuests:            body.MainQuests,
		SideQuests:            body.SideQuests,
		CompletedMainQuests:   body.CompletedMainQuests,
		CompletedSideQuests:   body.CompletedSideQuests,
		Collectibles:          body.Collectibles,
		CollectedCollectibles: body.CollectedCollectibles,
		CurrentlyPlaying:      body.CurrentlyPlaying,
		GameGuide:             body.GameGuide,
		GameMap:               body.GameMap,
	})

	//respond with it
	c.JSON(200, gin.H{
		"game": game,
	})
}

func GamesDelete(c *gin.Context) {
	//get the id off the URL
	id := c.Param("id")

	//delete the posts
	initalizers.DB.Delete(&models.Game{}, id)

	//Respond
	c.Status(200)
}
