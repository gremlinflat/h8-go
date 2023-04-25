package controllers

import (
	. "C2-simple-rest-gin/model"
	repo "C2-simple-rest-gin/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllBook(ctx *gin.Context) {
	if len(repo.DataBooks) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No books, Empty DB",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"books": repo.DataBooks,
	})
}

func GetBook(ctx *gin.Context) {
	bookIDStr := ctx.Param("bookID")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Book ID should be number",
		})
		return
	}
	//bookID := ctx.Param("bookID")
	found := false
	var databook Book

	for i, book := range repo.DataBooks {
		if bookID == book.ID {
			found = true
			databook = repo.DataBooks[i]
			break
		}
	}

	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Books not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": databook,
	})
}

func AddBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var maxID int
	for _, book := range repo.DataBooks {
		if book.ID > maxID {
			maxID = book.ID
		}
	}

	newID := maxID + 1
	newBook.ID = newID
	repo.DataBooks = append(repo.DataBooks, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Book Succesfully Created",
		"book":    newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	for i := range repo.DataBooks {
		if strconv.Itoa(repo.DataBooks[i].ID) == bookID {
			updatedBook.ID = repo.DataBooks[i].ID
			repo.DataBooks[i] = updatedBook

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Book succesfully updated",
				"book":    updatedBook,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Invalid Book ID",
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	found := false

	bookIDInt, err := strconv.Atoi(bookID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid ID: books not found",
		})
	}

	for i, book := range repo.DataBooks {
		if bookIDInt == book.ID {
			found = true
			repo.DataBooks = append(repo.DataBooks[:i], repo.DataBooks[i+1:]...)
			break
		}
	}

	if found {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Book successfully deleted",
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid ID: books not found",
		})
	}
}
