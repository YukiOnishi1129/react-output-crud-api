package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/apperrors"
)

type BaseHandler struct{}

// エラーレスポンスの構造体
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// JSONレスポンスを返す共通メソッド
func (h *BaseHandler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		h.respondError(w, apperrors.NewInternalError("failed to marshal response", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// エラーレスポンスを返す共通メソッド
func (h *BaseHandler) respondError(w http.ResponseWriter, err error) {
	var status int
	var response ErrorResponse

	// アプリケーションのエラー型の場合
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		switch appErr.Type {
		case apperrors.NotFound:
			status = http.StatusNotFound
		case apperrors.ValidationError:
			status = http.StatusBadRequest
		case apperrors.PermissionDenied:
			status = http.StatusForbidden
		case apperrors.Unauthorized:
			status = http.StatusUnauthorized
		case apperrors.AlreadyExists:
			status = http.StatusConflict
		case apperrors.BusinessRuleError:
			status = http.StatusUnprocessableEntity
		default:
			status = http.StatusInternalServerError
		}

		response = ErrorResponse{
			Code:    string(appErr.Type),
			Message: appErr.Message,
		}
	} else {
		// 未知のエラーの場合
		status = http.StatusInternalServerError
		response = ErrorResponse{
			Code:    string(apperrors.InternalError),
			Message: "internal server error",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
} 