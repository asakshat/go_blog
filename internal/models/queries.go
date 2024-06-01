package models

import "gorm.io/gorm"

type PostDetails struct {
	Username     string
	Post         Post
	LikeCount    int64
	CommentCount int64
	Comments     []CommentDetails
}
type CommentDetails struct {
	CommentString   string
	CommentUsername string
}

func GetLikeCount(db *gorm.DB, postID uint) (int64, error) {
	var count int64
	err := db.Model(&Like{}).Where("post_id = ?", postID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func GetCommentCount(db *gorm.DB, postID uint) (int64, error) {
	var count int64
	err := db.Model(&Comment{}).Where("post_id = ?", postID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetAllPostDetails(db *gorm.DB) ([]PostDetails, error) {
	var posts []Post
	err := db.Preload("User").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	var postDetails []PostDetails
	for _, post := range posts {
		likeCount, err := GetLikeCount(db, post.PostID)
		if err != nil {
			return nil, err
		}

		commentCount, err := GetCommentCount(db, post.PostID)
		if err != nil {
			return nil, err
		}

		postDetails = append(postDetails, PostDetails{
			Username:     post.User.Username,
			Post:         post,
			LikeCount:    likeCount,
			CommentCount: commentCount,
		})
	}

	return postDetails, nil
}
func GetPostWithId(db *gorm.DB, postID uint) (PostDetails, error) {
	var post Post
	err := db.Preload("User").Preload("Comments.User").Where("post_id = ?", postID).First(&post).Error
	if err != nil {
		return PostDetails{}, err
	}

	likeCount, err := GetLikeCount(db, post.PostID)
	if err != nil {
		return PostDetails{}, err
	}

	commentCount, err := GetCommentCount(db, post.PostID)
	if err != nil {
		return PostDetails{}, err
	}

	var commentDetails []CommentDetails
	for _, comment := range post.Comments {
		commentDetails = append(commentDetails, CommentDetails{
			CommentString:   comment.Comment,
			CommentUsername: comment.User.Username,
		})
	}

	postDetails := PostDetails{
		Username:     post.User.Username,
		Post:         post,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		Comments:     commentDetails,
	}

	return postDetails, nil
}
