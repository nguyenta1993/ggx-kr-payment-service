package validation

import (
	"encoding/json"
	httpcommon "github.com/gogovan/ggx-kr-payment-service/internal/api/http/common"
	utils "github.com/gogovan/ggx-kr-payment-service/pkg/string_utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gogovan/ggx-kr-service-utils/logger"
)

func CheckValidateHTTP(c *gin.Context, dest interface{}, logger logger.Logger) (err error) {
	if err = c.ShouldBindJSON(&dest); err != nil {
		checkErr(c, logger, err)
		return
	}
	return
}

func checkErr(c *gin.Context, logger logger.Logger, err error) {
	switch t := err.(type) {
	case *json.UnmarshalTypeError:
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(t.Field, httpcommon.RequestInvalid, t.Field))
		return
	case *json.SyntaxError:
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(err.Error(), httpcommon.RequestInvalid, ""))
		return
	case validator.ValidationErrors:
		errors := HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, httpcommon.NewPartialSuccess[any](false, nil, errors))
		return
	default:
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(err.Error(), httpcommon.RequestInvalid, ""))
		return
	}
}

func HandleValidationErrors(err error) (errors []httpcommon.ErrorResponse) {
	for _, fieldErr := range err.(validator.ValidationErrors) {
		fields := utils.LowerInitial(strings.Split(fieldErr.StructNamespace(), ".")[1:])
		field := strings.Join(fields, ".")
		errorCode, ok := TagMap[fieldErr.Tag()]
		if !ok {
			errorCode = httpcommon.InvalidFormat
		}
		errors = httpcommon.AddError(errors, fieldErr.Tag(), string(errorCode), field)
	}
	return
}
