package controllers

import (
	"context"
	"io"
	"os"
	"strconv"

	"github.com/asakshat/go_blog/db"
	"github.com/asakshat/go_blog/internal/models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

func parseID(c *fiber.Ctx, param string) (int, error) {
	id, err := strconv.Atoi(c.Params(param))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func handleError(c *fiber.Ctx, status int, message string) error {
	c.Status(status)
	return c.JSON(fiber.Map{
		"message": message,
	})
}

func CreateBlog(c *fiber.Ctx) error {
	postTitle := c.FormValue("post_title")
	postDescription := c.FormValue("post_description")
	postContent := c.FormValue("post_content")

	if postTitle == "" || postDescription == "" || postContent == "" {
		return handleError(c, fiber.StatusBadRequest, "Title, description, and content cannot be empty")
	}

	userID, err := parseID(c, "user_id")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return handleError(c, fiber.StatusBadRequest, "User not found: "+err.Error())
	}
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Image upload failed: "+err.Error())
	}

	file, err := fileHeader.Open()
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Failed to open file: "+err.Error())
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("", "upload-*.png")
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Failed to create temporary file: "+err.Error())
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Failed to save uploaded file: "+err.Error())
	}

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	cld.Config.URL.Secure = true
	ctx := context.Background()
	uniqueFilename := false
	overwrite := true
	resp, err := cld.Upload.Upload(ctx, tempFile.Name(), uploader.UploadParams{
		PublicID:       postTitle,
		UniqueFilename: &uniqueFilename,
		Overwrite:      &overwrite,
	})
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Image upload failed: "+err.Error())
	}

	post := models.Post{
		PostTitle:       postTitle,
		PostDescription: postDescription,
		PostContent:     postContent,
		ImageURL:        resp.SecureURL, // Add this line
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
		"image_url":        post.ImageURL, // Add this line
		"created_at":       post.CreatedAt,
	}

	return c.JSON(response)
}
func EditPost(c *fiber.Ctx) error {
	postTitle := c.FormValue("post_title")
	postDescription := c.FormValue("post_description")
	postContent := c.FormValue("post_content")
	userID, err := parseID(c, "user_id")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "User not found")
	}
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Image upload failed: "+err.Error())
	}

	file, err := fileHeader.Open()
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Failed to open file: "+err.Error())
	}
	defer file.Close()

	// create a temp file
	tempFile, err := os.CreateTemp("", "upload-*.png")
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Failed to create temporary file: "+err.Error())
	}
	defer os.Remove(tempFile.Name())

	// copy gile
	_, err = io.Copy(tempFile, file)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Failed to save uploaded file: "+err.Error())
	}

	// upload the image to Cloudinary
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	cld.Config.URL.Secure = true
	ctx := context.Background()
	uniqueFilename := false
	overwrite := true
	resp, err := cld.Upload.Upload(ctx, tempFile.Name(), uploader.UploadParams{
		PublicID:       postTitle,
		UniqueFilename: &uniqueFilename,
		Overwrite:      &overwrite,
	})
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Image upload failed: "+err.Error())
	}

	var post models.Post
	db.DB.First(&post, userID)
	post.PostTitle = postTitle
	post.PostDescription = postDescription
	post.PostContent = postContent
	post.ImageURL = resp.SecureURL
	db.DB.Save(&post)

	response := map[string]interface{}{
		"post_id":          post.PostID,
		"user_id":          post.UserID,
		"post_title":       post.PostTitle,
		"post_description": post.PostDescription,
		"post_content":     post.PostContent,
		"image_url":        post.ImageURL,
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

func GetAllPostByUserId(c *fiber.Ctx) error {
	userID, err := parseID(c, "user_id")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}

	var posts []models.Post
	if err := db.DB.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Error getting posts: "+err.Error())
	}

	return c.JSON(posts)
}

func GetPostWithIdHandler(c *fiber.Ctx) error {
	postID, err := parseID(c, "post_id")
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
	userID, err := parseID(c, "user_id")
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user ID: "+err.Error())
	}
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return handleError(c, fiber.StatusNotFound, "User not found")
	}
	postID, err := parseID(c, "post_id")
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
