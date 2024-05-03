package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sqids/sqids-go"
	"github.com/tun43p/api/common"
	"gorm.io/gorm"
)

type URLRequest struct {
	Original string `json:"original"`
}

type URLResponse struct {
	Original  string `json:"original" gorm:"unique"`
	Short     string `json:"short"`
	CreatedAt int64  `json:"created_at"`
}

func GetURLs(ctx *gin.Context, db *gorm.DB) {
	u := ctx.Query("u")

	var urls []URLResponse
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

func ShrinkUrl(ctx *gin.Context, db *gorm.DB) {
	var url URLRequest
	var urls []URLResponse

	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest,
			&common.Error{
				Status:  http.StatusBadRequest,
				Message: "Error binding JSON request body",
				Error:   err.Error(),
			})

		return
	}

	db.Find(&urls)

	for _, a := range urls {
		if a.Original == url.Original {
			ctx.IndentedJSON(http.StatusConflict, &common.Error{
				Status:  http.StatusConflict,
				Message: "URL already exists",
				Error:   "URL already exists",
			})

			return
		}
	}

	s, err := sqids.New()

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, &common.Error{
			Status:  http.StatusInternalServerError,
			Message: "Error generating short URL",
			Error:   err.Error(),
		})
	}

	hash, err := s.Encode([]uint64{uint64(time.Now().Unix())})

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, &common.Error{
			Status:  http.StatusInternalServerError,
			Message: "Error generating short URL",
			Error:   err.Error(),
		})
	}

	db.Create(&URLResponse{
		Original:  url.Original,
		Short:     "http://localhost:8080/s/" + hash,
		CreatedAt: time.Now().Unix(),
	})

	db.Find(&urls)

	ctx.IndentedJSON(http.StatusCreated, urls)
}

func Redirect(ctx *gin.Context, db *gorm.DB) {
	s := ctx.Param("s")

	var urls []URLResponse
	db.Find(&urls)

	for _, a := range urls {
		if a.Short == "http://localhost:8080/s/"+s {
			ctx.Redirect(http.StatusMovedPermanently, a.Original)

			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, &common.Error{
		Status:  http.StatusNotFound,
		Message: "Short URL not found",
		Error:   "Short URL not found",
	})
}
