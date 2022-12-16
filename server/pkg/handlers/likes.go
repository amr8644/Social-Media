package handlers

import (
	"net/http"
	"server/pkg/models"

	"github.com/labstack/echo/v4"
)

func LikePost(c echo.Context) error {
	liked_by := new(models.Likes)

	err := c.Bind(&liked_by); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

  	like := models.Likes{
	PostID: liked_by.PostID,
	UserID: liked_by.UserID,
  	}

	// Add to database
	b,exists := like.LikePost()	
	if exists.Exists {
		return c.JSON(http.StatusNotAcceptable, "Already liked the post")
	}

	return c.JSON(http.StatusOK, b)
}
