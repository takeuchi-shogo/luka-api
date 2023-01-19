package controllers

import (
	apierrors "github.com/takeuchi-shogo/luka-api/src/pkg/api-errors"
)

type ErrorResponse struct {
	Error apierrors.ApiError `json:"error"`
}

func NewErrorResponse(err apierrors.ApiError) *ErrorResponse {
	res := new(ErrorResponse)
	res.Error = err

	return res
}
