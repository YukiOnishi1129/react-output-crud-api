package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/interfaces/handler/constants"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase/auth"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	authUseCase auth.AuthUseCase
}

func NewAuthHandler(authUseCase auth.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// RegisterHandlers は、認証関連のハンドラーを登録します
func (h *AuthHandler) RegisterHandlers(r *mux.Router) {
	// /api/v1/auth のサブルーターを作成
	authRouter := r.PathPrefix(constants.AuthPath).Subrouter()

	authRouter.HandleFunc(constants.RegisterPath, h.Register).Methods("POST")
	authRouter.HandleFunc(constants.LoginPath, h.Login).Methods("POST")
	authRouter.HandleFunc(constants.LogoutPath, h.Logout).Methods("POST")
	authRouter.HandleFunc(constants.RefreshPath, h.RefreshToken).Methods("POST")
}

// Register は新規ユーザー登録を処理します
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input auth.RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.authUseCase.Register(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// Login はユーザーログインを処理します
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input auth.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.authUseCase.Login(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// JWTトークンをCookieにセット
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    output.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
} 