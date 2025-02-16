package middleware

import (
	"auth-app/internal/apperrors"
	"auth-app/internal/dto"
	"auth-app/pkg/customerror"
	"auth-app/pkg/jsonvalidator"
	"auth-app/pkg/typeconverter"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		err := ctx.Errors[0].Err

		switch e := err.(type) {
		case validator.ValidationErrors:
			var errors []dto.Error
			for _, fe := range e {
				errors = append(errors, dto.Error{
					Field:  fe.Field(),
					Detail: jsonvalidator.ExtractValidationError(fe),
				})
			}
			respondWithBadRequestError(ctx, errors)
		case *json.SyntaxError:
			respondWithBadRequestError(ctx, []dto.Error{{
				Field:  apperrors.FieldJSON,
				Detail: e.Error(),
			}})
		case *json.UnmarshalTypeError:
			respondWithBadRequestError(ctx, []dto.Error{{
				Field:  e.Field,
				Detail: fmt.Sprintf("%s should be %s", e.Field, typeconverter.ConvertToUserFrinedlyName(e.Type.Name())),
			}})
		case *customerror.BadRequestError:
			respondWithBadRequestError(ctx, []dto.Error{{
					Field: e.Field,
					Detail: e.Sentinel.Error(),
			}})
		case *customerror.UnauthorizedError:
			respondWithUnauthorizedError(ctx, []dto.Error{{
				Field: e.Field,
				Detail: e.Sentinel.Error(),
			}})
		case *customerror.NotFoundError:
			respondWithNotFoundError(ctx, []dto.Error{{
				Field: e.Field,
				Detail: e.Sentinel.Error(),
			}})
		case *customerror.InternalServerError:
			respondWithInternalServerError(ctx, []dto.Error{{
				Field: e.Field,
				Detail: e.Sentinel.Error(),
			}})
		}
	}
}

func respondWithBadRequestError(ctx *gin.Context, errors []dto.Error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
		Errors: errors,
	})
}

func respondWithUnauthorizedError(ctx *gin.Context, errors []dto.Error) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
		Errors: errors,
	})
}

func respondWithNotFoundError(ctx *gin.Context, errors []dto.Error) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, dto.Response{
		Errors: errors,
	})
}

func respondWithInternalServerError(ctx *gin.Context, errors []dto.Error) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
		Errors: errors,
	})
}