package controllers

import (
	"github.com/asakshat/go_blog/db"
	"github.com/asakshat/go_blog/internal/models"
	"github.com/gofiber/fiber/v2"
)

func PostComment(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return handleError(c, fiber.StatusBadRequest, "Could not parse request body: "+err.Error())
	}
	userID, err := parseUserID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}
	// user eixts or not
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

	comment := models.Comment{
		PostID:  uint(postID),
		UserID:  uint(userID),
		Comment: data["comment"],
	}
	db.DB.Create(&comment)
	var postDetails models.PostDetails
	postDetails.CommentCount++
	db.DB.First(&comment, comment.ID)
	response := map[string]interface{}{
		"comment_id": comment.ID,
		"post_id":    comment.PostID,
		"user_id":    comment.UserID,
		"comment":    comment.Comment,
		"created_at": comment.CreatedAt,
	}
	return c.JSON(response)

}
func EditComment(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return handleError(c, fiber.StatusBadRequest, "Could not parse request body: "+err.Error())
	}
	userId, err := parseUserID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}

	commentID, err := parseCommentID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid comment ID: "+err.Error())
	}
	var comment models.Comment
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "Comment not found")
	}

	// user check
	if comment.UserID != uint(userId) {
		return handleError(c, fiber.StatusUnauthorized, "You are not authorized to edit this comment")
	}

	db.DB.Save(&comment)
	response := map[string]interface{}{
		"comment_id": comment.ID,
		"post_id":    comment.PostID,
		"user_id":    comment.UserID,
		"comment":    data["comment"],
	}
	return c.JSON(response)
}

func DeleteComment(c *fiber.Ctx) error {
	userID, err := parseUserID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}

	commentID, err := parseCommentID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid comment ID: "+err.Error())
	}
	var comment models.Comment
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "Comment not found")
	}

	// user check
	if comment.UserID != uint(userID) {
		return handleError(c, fiber.StatusUnauthorized, "You are not authorized to delete this comment")
	}
	db.DB.Delete(&comment)
	return c.JSON(fiber.Map{
		"message": "Comment deleted successfully",
	})
}
