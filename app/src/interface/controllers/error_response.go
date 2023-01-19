package controllers

import (
	apierrors "github.com/takeuchi-shogo/luka-api/src/pkg/api-errors"
)

type ErrorResponse struct {
	Error apierrors.ApiError `json:"error"`
}

type Error struct {
	Next        error  `json:"error"`
	StatusCode  int    `json:"code"`
	Message     string `json:"message"`
	InfoMessage string `json:"infoMessage"`
	Type        string `json:"type"`
	// Code        string `json:"errorCode"`
}

func NewErrorResponse(err apierrors.ApiError) *ErrorResponse {
	res := new(ErrorResponse)
	res.Error = err

	return res
}
