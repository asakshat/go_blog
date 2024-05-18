package controllers

import (
	"strconv"

	"github.com/asakshat/go_blog/db"
	"github.com/asakshat/go_blog/internal/models"
	"github.com/gofiber/fiber/v2"
)

func parseUserID(c *fiber.Ctx) (int, error) {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return 0, err
	}
	return userID, nil
}
func parseCommentID(c *fiber.Ctx) (int, error) {
	commentID, err := strconv.Atoi(c.Params("comment_id"))
	if err != nil {
		return 0, err
	}
	return commentID, nil
}

func parsePostID(c *fiber.Ctx) (int, error) {
	postID, err := strconv.Atoi(c.Params("post_id"))
	if err != nil {
		return 0, err
	}
	return postID, nil
}

func handleError(c *fiber.Ctx, status int, message string) error {
	c.Status(status)
	return c.JSON(fiber.Map{
		"message": message,
	})
}

func CreateBlog(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return handleError(c, fiber.StatusBadRequest, "Could not parse request body: "+err.Error())
	}
	if data["post_title"] == "" || data["post_description"] == "" || data["post_content"] == "" {
		return handleError(c, fiber.StatusBadRequest, "Title, description, and content cannot be empty")
	}

	userID, err := parseUserID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return handleError(c, fiber.StatusBadRequest, "User not found: "+err.Error())
	}

	post := models.Post{
		PostTitle:       data["post_title"],
		PostDescription: data["post_description"],
		PostContent:     data["post_content"],
		UserID:          uint(userID),
	}
	db.DB.Create(&post)

	db.DB.First(&post, post.PostID)
	response := map[string]interface{}{
		"post_id":          post.PostID,
		"user_id":          post.UserID,
		"post_title":       post.PostTitle,
		"post_description": post.PostDescription,
		"post_content":     post.PostContent,
		"created_at":       post.CreatedAt,
	}

	return c.JSON(response)
}

func EditPost(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return handleError(c, fiber.StatusBadRequest, "Could not parse request body: "+err.Error())
	}
	userID, err := parseUserID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "User not found")
	}
	postID, err := parsePostID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid post ID: "+err.Error())
	}
	var post models.Post
	if err := db.DB.First(&post, postID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "Post not found")
	}
	if post.UserID != uint(userID) {
		return handleError(c, fiber.StatusUnauthorized, "You are not authorized to edit this post")
	}

	// Update the post fields with the new data
	post.PostTitle = data["post_title"]
	post.PostDescription = data["post_description"]
	post.PostContent = data["post_content"]

	db.DB.Save(&post)
	response := map[string]interface{}{
		"post_id":          post.PostID,
		"user_id":          post.UserID,
		"post_title":       post.PostTitle,
		"post_description": post.PostDescription,
		"post_content":     post.PostContent,
		"created_at":       post.CreatedAt,
	}
	return c.JSON(response)
}

func GetAllPosts(c *fiber.Ctx) error {
	postDetails, err := models.GetAllPostDetails(db.DB)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Error getting post details: "+err.Error())
	}

	return c.JSON(postDetails)
}

func GetPostWithIdHandler(c *fiber.Ctx) error {
	postID, err := parsePostID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Cannot parse post_id")
	}

	postDetails, err := models.GetPostWithId(db.DB, uint(postID))
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Error getting post details: "+err.Error())
	}

	return c.JSON(postDetails)
}

func DeletePost(c *fiber.Ctx) error {
	userID, err := parseUserID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "User not found")
	}
	postID, err := parsePostID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid post ID: "+err.Error())
	}
	var post models.Post
	if err := db.DB.First(&post, postID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "Post not found")
	}
	if post.UserID != uint(userID) {
		return handleError(c, fiber.StatusUnauthorized, "You are not authorized to delete this post")
	}
	db.DB.Delete(&post)
	return c.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}
