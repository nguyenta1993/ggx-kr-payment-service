package httpcommon

type SuccessResponse[T any] struct {
	Success bool            `json:"success"`
	Data    *T              `json:"data"`
	Errors  []ErrorResponse `json:"errors"`
}

type PagingSuccessResponse[T any] struct {
	Success bool            `json:"success"`
	Data    *T              `json:"data"`
	Errors  []ErrorResponse `json:"errors"`
	Paging
}

type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field"`
	Code    string `json:"code"`
}

func NewSuccessResponse[T any](data T) SuccessResponse[T] {
	return SuccessResponse[T]{
		Data:    &data,
		Success: true,
		Errors:  nil,
	}
}

func NewPartialSuccess[T any](success bool, data T, errors []ErrorResponse) SuccessResponse[T] {
	return SuccessResponse[T]{
		Data:    &data,
		Success: success,
		Errors:  errors,
	}
}

func NewPagingSuccessResponse[T any](success bool, data T, errors []ErrorResponse, paging Paging) PagingSuccessResponse[T] {
	return PagingSuccessResponse[T]{
		Data:    &data,
		Success: success,
		Errors:  errors,
		Paging:  paging,
	}
}

func NewErrorResponse(message string, code string, field string) SuccessResponse[any] {
	var errors []ErrorResponse
	error := ErrorResponse{
		Message: message,
		Field:   field,
		Code:    code,
	}
	errors = append(errors, error)
	return SuccessResponse[any]{
		Success: false,
		Data:    nil,
		Errors:  errors,
	}
}

func NewError(message string, code string, field string) []ErrorResponse {
	var errors []ErrorResponse
	error := ErrorResponse{
		Message: message,
		Field:   field,
		Code:    code,
	}
	errors = append(errors, error)
	return errors
}

func AddError(errOrigin []ErrorResponse, message string, code string, field string) []ErrorResponse {
	var errors []ErrorResponse
	err := ErrorResponse{
		Message: message,
		Field:   field,
		Code:    code,
	}
	errors = append(errOrigin, err)
	return errors
}
