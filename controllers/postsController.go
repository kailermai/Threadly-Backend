package controllers

import (
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	// Get data off request body
	var body map[string]string

	err := c.BodyParser(&body)

	if err != nil {
		return err
	}

	// Create a post
	post := models.Post{
		Title: body["title"],
		Body:  body["body"],
		User:  body["user"],
		Tag:   body["tag"],
	}

	result := database.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return nil
	}

	// Return it
	return c.JSON(post)

}

func FindPosts(c *fiber.Ctx) error {
	// Get the post
	var posts []models.Post

	database.DB.Find(&posts)

	// Respond with it
	return c.JSON(posts)
}

func FindSinglePost(c *fiber.Ctx) error {
	// Get id of url
	id := c.Params("id")

	// Get the post
	var post models.Post

	database.DB.First(&post, id)

	// Respond with it
	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	// Get id of url
	id := c.Params("id")

	// Get data from request body
	var body map[string]string

	err := c.BodyParser(&body)

	if err != nil {
		return err
	}

	// Find post we want to update
	var post models.Post

	database.DB.First(&post, id)

	// Update
	database.DB.Model(&post).Updates(models.Post{
		Title: body["title"],
		Body:  body["body"],
		User:  body["user"],
		Tag:   body["tag"],
	})

	// Respond with it
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	// Get id of url
	id := c.Params("id")

	// Delete post
	database.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
