package controllers

import (
	"github.com/asakshat/go_blog/db"
	"github.com/asakshat/go_blog/internal/models"
	"github.com/gofiber/fiber/v2"
)

func LikePost(c *fiber.Ctx) error {
	userID, err := parseID(c, "user_id")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Cannot parse user_id")
	}

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "User not found")
	}

	postID, err := parseID(c, "post_id")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Cannot parse post_id")
	}

	var post models.Post
	if err := db.DB.First(&post, postID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "Post not found")
	}

	var like models.Like
	db.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&like)
	if like.ID != 0 {
		return handleError(c, fiber.StatusBadRequest, "User has already liked this post")
	}

	like = models.Like{
		UserID: uint(userID),
		PostID: uint(postID),
	}

	db.DB.Create(&like)
	var postDetails models.PostDetails
	postDetails.LikeCount++
	db.DB.Save(&post)

	return c.JSON(fiber.Map{
		"message": "Post liked successfully",
	})
}

func UnlikePost(c *fiber.Ctx) error {
	userID, err := parseID(c, "user_id")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Cannot parse user_id")
	}
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "User not found")
	}

	postID, err := parseID(c, "post_id")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Cannot parse post_id")
	}

	var post models.Post
	if err := db.DB.First(&post, postID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "Post not found")
	}
	var like models.Like
	db.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&like)
	if like.ID == 0 {
		return handleError(c, fiber.StatusBadRequest, "User has not liked this post")
	}

	db.DB.Delete(&like)
	var postDetails models.PostDetails
	db.DB.First(&post, postID)
	postDetails.LikeCount--
	db.DB.Save(&post)

	return c.JSON(fiber.Map{
		"message": "Post unliked successfully",
	})
}
