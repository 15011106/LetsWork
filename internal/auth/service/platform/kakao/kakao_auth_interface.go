package platform

import "github.com/15011106/LetsWork/internal/auth/auth_dto"

type KakaoAuth interface {
	GetAccessToken(code string) (*auth_dto.KakaoTokenResponse, error)
	GetUserInfo(accessToken string) (*auth_dto.KakaoProfile, error)
}
