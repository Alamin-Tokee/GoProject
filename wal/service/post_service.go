package service

import (
	"net/http"
	"strconv"
	"wal/model"
	"wal/repository"
	"wal/util"

	"github.com/gin-gonic/gin"
)

// RoutesPost ...
func RoutesPost(rg *gin.RouterGroup) {
	post := rg.Group("/post")

	post.GET("/:id", util.TokenAuthMiddleware(), getPostByID)
	post.GET("/", getPosts)
}

// getPostByID godoc
// @Summary show Post by id
// @Description get string by ID
// @Tags Post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {object} model.MPost
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /post/{id} [get]
func getPostByID(c *gin.Context) {
	var post model.MPost
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	post, err = repository.GetPostByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	if (model.MPost{}) == post {
		c.JSON(http.StatusNotFound, post)
	} else {
		c.JSON(http.StatusOK, post)
	}
}

// getPosts godoc
// @Summary show list post
// @Description get posts
// @Tags Post
// @Accept  json
// @Produce  json
// @Success 200 {array} model.MPost
// @Failure 400 {string} string
// @Failure 404 {object} model.MPost
// @Failure 500 {string} string
// @Router /post/ [get]
func getPosts(c *gin.Context) {

	var posts []model.MPost
	posts, err := repository.GetPostAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
