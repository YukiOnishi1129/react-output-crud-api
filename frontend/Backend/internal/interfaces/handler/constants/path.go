package constants

import (
	"fmt"
)

// APIのベースパス
const (
	APIBasePath = "/api/v1"
)

// 各リソースのパス
const (
	TodosPath = APIBasePath + "/todos"
	AuthPath  = APIBasePath + "/auth"
	UsersPath = APIBasePath + "/users"
)

// 認証関連のエンドポイント
const (
	RegisterPath   = "/register"
	LoginPath     = "/login"
	LogoutPath    = "/logout"
	RefreshPath   = "/refresh"
)

// BuildResourcePath は、リソースIDを含むパスを生成します
func BuildResourcePath(basePath string, id string) string {
	return fmt.Sprintf("%s/%s", basePath, id)
}

// BuildAuthPath は、認証関連のパスを生成します
func BuildAuthPath(endpoint string) string {
	return AuthPath + endpoint
} 