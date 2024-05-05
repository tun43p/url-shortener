package shortener

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sqids/sqids-go"
	"gorm.io/gorm"

	"github.com/tun43p/tun43p.com/internal/api/failure"
)

func GetSingleOrAllShortenedUrls(ctx *gin.Context, db *gorm.DB) {
	u := ctx.Query("u")

	var urls []ShortenerResponse
	db.Find(&urls)

	if u != "" {
		for _, a := range urls {
			if a.Original == u {
				ctx.IndentedJSON(http.StatusOK, a)

				return
			}
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "URL not found",
		})

		return
	}

	ctx.IndentedJSON(http.StatusOK, urls)
}

func ShortenUrl(ctx *gin.Context, db *gorm.DB) {
	var url ShortenerResponse
	var urls []ShortenerResponse

	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest,
			&failure.FailureResponse{
				Status:  http.StatusBadRequest,
				Message: "Error binding JSON request body",
				Error:   err.Error(),
			})

		return
	}

	db.Find(&urls)

	for _, a := range urls {
		if a.Original == url.Original {
			ctx.IndentedJSON(http.StatusConflict, &failure.FailureResponse{
				Status:  http.StatusConflict,
				Message: "URL already exists",
				Error:   "URL already exists",
			})

			return
		}
	}

	s, err := sqids.New()

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, &failure.FailureResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error generating short URL",
			Error:   err.Error(),
		})
	}

	hash, err := s.Encode([]uint64{uint64(time.Now().Unix())})

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, &failure.FailureResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error generating short URL",
			Error:   err.Error(),
		})
	}

	db.Create(&ShortenerResponse{
		Original:  url.Original,
		Short:     "http://localhost:8080/s/" + hash,
		CreatedAt: time.Now().Unix(),
	})

	db.Find(&urls)

	ctx.IndentedJSON(http.StatusCreated, urls)
}

func RedirectShortenedUrlToOriginalUrl(ctx *gin.Context, db *gorm.DB) {
	u := ctx.Param("u")

	var data []ShortenerResponse
	db.Find(&data)

	for _, a := range data {
		fmt.Println(u, a.Short, a.Original)

		if a.Short == "http://localhost:8080/s/"+u {
			fmt.Println(u, a.Short, a.Original)
			ctx.Redirect(http.StatusMovedPermanently, a.Original)

			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, &failure.FailureResponse{
		Status:  http.StatusNotFound,
		Message: "Short URL not found",
		Error:   "Short URL not found",
	})
}
