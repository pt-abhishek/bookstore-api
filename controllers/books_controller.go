package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/bookstore-api/services"
	"github.com/pt-abhishek/bookstore-api/utils/errors"
)

//SearchBooks searchs for books using seartchtext from queryparam
func SearchBooks(c *gin.Context) {
	searchText := c.DefaultQuery("book_name", "")
	books, err := services.BookService.GetBySearchText(searchText)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, books)
}

//FindAllWithPagination All with pagination
func FindAllWithPagination(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		parseErr := errors.NewBadRequestError("Error parsing to integer")
		c.JSON(parseErr.Code, parseErr)
		return
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		parseErr := errors.NewBadRequestError("Error parsing to integer")
		c.JSON(parseErr.Code, parseErr)
		return
	}
	books, bookGetErr := services.BookService.GetAllWithPagination(pageInt, pageSizeInt)
	if bookGetErr != nil {
		c.JSON(bookGetErr.Code, bookGetErr)
		return
	}
	c.JSON(http.StatusOK, books)
}
