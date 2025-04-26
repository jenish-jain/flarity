package transaction

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ValidateYearAndMonth validates and parses the year and month query parameters.
func ValidateYearAndMonth(ctx *gin.Context) (int, int, error) {
	yearParam := ctx.Query("year")
	monthParam := ctx.Query("month")

	// Parse and validate year
	year, err := strconv.Atoi(yearParam)
	if err != nil || year <= 0 {
		return 0, 0, errors.New("Invalid or missing 'year' query parameter")
	}

	// Parse and validate month
	month, err := strconv.Atoi(monthParam)
	if err != nil || month < 1 || month > 12 {
		return 0, 0, errors.New("Invalid or missing 'month' query parameter")
	}

	return year, month, nil
}
