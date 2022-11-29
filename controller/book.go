package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"books_memo_app/model"
	"books_memo_app/service"
	"strconv"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.bookService{}
	err = bookService.SetBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookList(c *gin.Context){
	bookService := service.bookService{}
}