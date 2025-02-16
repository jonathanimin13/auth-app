package customerror

type customError struct {
	Field    string
	Sentinel error
	Actual   error
}

func (e *customError) Error() string {
	return e.Actual.Error()
}

type BadRequestError struct {
	customError
}

func NewBadRequestError(field string, sentinel error, actual error) *BadRequestError {
	return &BadRequestError{
		customError: customError{
			Field:    field,
			Sentinel: sentinel,
			Actual:   actual,
		},
	}
}

type UnauthorizedError struct {
	customError
}

func NewUnauthorizedError(field string, sentinel error, actual error) *UnauthorizedError {
	return &UnauthorizedError{
		customError: customError{
			Field:    field,
			Sentinel: sentinel,
			Actual:   actual,
		},
	}
}

type NotFoundError struct {
	customError
}

func NewNotFoundError(field string, sentinel error, actual error) *NotFoundError {
	return &NotFoundError{
		customError: customError{
			Field:    field,
			Sentinel: sentinel,
			Actual:   actual,
		},
	}
}

type InternalServerError struct {
	customError
}

func NewInternalServerError(field string, sentinel error, actual error) *InternalServerError {
	return &InternalServerError{
		customError: customError{
			Field:    field,
			Sentinel: sentinel,
			Actual:   actual,
		},
	}
}