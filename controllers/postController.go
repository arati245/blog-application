package controllers

import (
	"blog_api/config"
	"blog_api/models"
	"net/http"
    "github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

func GetPost(c *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found!"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found!"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := config.DB.Delete(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
