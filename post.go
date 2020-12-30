package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	Post struct {
		ID      int    `json:"id"`
		Title   string `json:"title,omitempty"`
		Content string `json:"content,omitempty"`
	}

	PostService struct {
		Posts []*Post
	}
)

func prepareData() []*Post {
	return []*Post{
		{
			ID:      1,
			Title:   "First post",
			Content: "This is first post",
		},
		{
			ID:      2,
			Title:   "Second post",
			Content: "This is second post",
		},
	}
}

func RegisterPostAPI(v1 *echo.Group) {
	s := PostService{Posts: prepareData()}
	endpoint := v1.Group("/posts")
	endpoint.GET("", s.home)
	endpoint.GET("/:id", s.detail)
}

func (s *PostService) home(c echo.Context) error {
	return c.JSON(http.StatusOK, Wrapper{
		Code:    http.StatusOK,
		Content: s.Posts,
	})
}

func (s *PostService) detail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := s.find(id)
	if err != nil {
		if err, ok := err.(*Error); ok {
			return c.JSON(err.Code, Wrapper{Code: err.Code, ErrorMessage: err.Message})
		}
		return err
	}
	return c.JSON(http.StatusOK, Wrapper{
		Code:    http.StatusOK,
		Content: post,
	})
}

func (s *PostService) find(id int) (*Post, error) {
	for _, post := range s.Posts {
		if post.ID == id {
			return post, nil
		}
	}

	return nil, &Error{
		Code:    http.StatusNotFound,
		Message: "post does not exist",
	}
}
