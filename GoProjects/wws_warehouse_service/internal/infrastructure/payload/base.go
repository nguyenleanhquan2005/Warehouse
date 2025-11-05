package payload

import (
	appErrors "app/internal/error"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type FilterBase struct {
	Sort     string `form:"sort" json:"sort"`
	Keyword  string `form:"keyword" json:"keyword"`
	FromDate string `form:"from_date" json:"from_date"`
	ToDate   string `form:"to_date" json:"to_date"`
}

// Validate validates the payload and returns a standardized error if validation fails
func Validate(payl interface{}, message string) error {
	validate := validator.New()

	// Validate struct tags
	if err := validate.Struct(payl); err != nil {
		return appErrors.InvalidArgumentError{
			Msg:   fmt.Sprintf("Validation: %s: %s", message, err.Error()),
			Field: "Validation",
		}
	}

	// Check if payload has custom Validate method
	if _validator, ok := payl.(interface{ Validate() error }); ok {
		if err := _validator.Validate(); err != nil {
			return appErrors.InvalidArgumentError{
				Msg:   fmt.Sprintf("Custom Validation: %s: Detail: %s", message, err.Error()),
				Field: "Validation",
			}
		}
	}

	return nil
}

func ValidateVendorID(vendorID string) (int64, error) {
	if vendorID == "" {
		return 0, appErrors.InvalidArgumentError{
			Msg:   "vendor ID is required",
			Field: "vendor_id",
		}
	}

	vendorIDInt, err := strconv.ParseInt(vendorID, 10, 64)
	if err != nil {
		return 0, appErrors.InvalidArgumentError{
			Msg:   "invalid vendor ID format",
			Field: "vendor_id",
		}
	}

	if vendorIDInt <= 0 {
		return 0, appErrors.InvalidArgumentError{
			Msg:   "vendor ID must be a positive integer",
			Field: "vendor_id",
		}
	}

	return vendorIDInt, nil
}

func ValidateCategoryID(categoryID string) (int64, error) {
	if categoryID == "" {
		return 0, appErrors.InvalidArgumentError{
			Msg:   "category ID is required",
			Field: "category_id",
		}
	}

	categoryIDInt, err := strconv.ParseInt(categoryID, 10, 64)
	if err != nil {
		return 0, appErrors.InvalidArgumentError{
			Msg:   "invalid category ID format",
			Field: "category_id",
		}
	}

	if categoryIDInt <= 0 {
		return 0, appErrors.InvalidArgumentError{
			Msg:   "category ID must be a positive integer",
			Field: "category_id",
		}
	}

	return categoryIDInt, nil
}
