package response

import "github.com/gin-gonic/gin"

type ParamsWithSuccess struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message,omitempty"`
	Result     any    `json:"result,omitempty"`
	Meta       any    `json:"meta,omitempty"`
	Success    bool   `json:"success"`
}
type ParamsWithError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Success    bool   `json:"success,omitempty"`
	Errors     any    `json:"errors,omitempty"`
}

func WithSuccess(ctx *gin.Context, params ParamsWithSuccess) {
	if params.StatusCode == 0 {
		params.StatusCode = 200
	}
	if !params.Success {
		params.Success = true
	}

	ctx.JSON(params.StatusCode, params)
}

func WithError(ctx *gin.Context, params ParamsWithError) {
	if params.StatusCode == 0 {
		params.StatusCode = 400
	}

	if params.Message == "" {
		params.Message = "Bad Request"
	}
	params.Success = false

	ctx.JSON(params.StatusCode, params)
}
