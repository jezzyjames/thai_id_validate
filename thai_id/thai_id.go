package thai_id

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ThaiID struct {
	ID string `json:"id"`
}

type ThaiIDHandler struct {
	db *sql.DB
}

func NewThaiIDHandler(db *sql.DB) ThaiIDHandler {
	return ThaiIDHandler{db: db}
}

func (handler ThaiIDHandler) ThaiIdValidateHandler(c *gin.Context) {
	var thaiID ThaiID

	if err := c.BindJSON(&thaiID); err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"error": err.Error(),
		})
	}

	if err := ValidateThaiID(thaiID.ID); err != nil {
		c.JSON(http.StatusOK, map[string]bool{
			"valid": false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]bool{
		"valid": true,
	})
}

func ValidateThaiID(id string) error {
	if len(id) != 13 {
		return errors.New("id digits incorrect")
	}

	splited := strings.Split(id, "")
	sum := 0
	for i, j := 0, 13; j > 1; i, j = i+1, j-1 {
		val, _ := strconv.Atoi(splited[i])
		sum += val * j
	}

	moded := sum % 11
	result := 11 - moded

	last := result % 10
	lastID, _ := strconv.Atoi(splited[len(splited)-1])

	if last != lastID {
		return errors.New("id incorrect")
	}
	return nil
}
