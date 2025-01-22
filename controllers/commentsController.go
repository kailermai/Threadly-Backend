package controllers

import (
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	// Get data off request body
	var body map[string]string

	err := c.BodyParser(&body)

	if err != nil {
		return err
	}

	// Create a post
	comment := models.Comment{
		Body:   body["body"],
		User:   body["user"],
		PostID: body["postid"],
	}

	result := database.DB.Create(&comment)

	if result.Error != nil {
		c.Status(400)
		return nil
	}

	// Return it
	return c.JSON(comment)

}

func FindComments(c *fiber.Ctx) error {
	// Get id of url
	postID := c.Params("id")

	// Get the post
	var comments []models.Comment

	database.DB.Where("post_id = ?", postID).Find(&comments)

	// Respond with it
	return c.JSON(comments)
}

func UpdateComment(c *fiber.Ctx) error {
	// Get id of url
	postID := c.Params("id")

	// Get data from request body
	var body map[string]string

	err := c.BodyParser(&body)

	if err != nil {
		return err
	}

	// Find post we want to update
	var comments models.Comment

	database.DB.First(&comments, postID)

	// Update
	database.DB.Model(&comments).Updates(models.Comment{
		Body:   body["commentBody"],
		User:   body["user"],
		PostID: body["postid"],
	})

	// Respond with it
	return c.JSON(comments)
}

func DeleteComment(c *fiber.Ctx) error {
	// Get id of url
	id := c.Params("id")

	// Delete post
	database.DB.Delete(&models.Comment{}, id)

	// Respond
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
