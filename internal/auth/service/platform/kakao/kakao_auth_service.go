package platform

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/15011106/LetsWork/internal/auth/auth_dto"
)

type KakaoAuthService struct {
	RestAPIKey  string
	RedirectURI string
	Client      *http.Client
}

func NewKakaoAuthService(restAPIKey, redirectURI string) KakaoAuth {
	return &KakaoAuthService{
		RestAPIKey:  restAPIKey,
		RedirectURI: redirectURI,
		Client:      &http.Client{},
	}
}

func (s *KakaoAuthService) GetAccessToken(code string) (*auth_dto.KakaoTokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", s.RestAPIKey)
	data.Set("redirect_uri", s.RedirectURI)
	data.Set("code", code)

	req, err := http.NewRequest("POST", "https://kauth.kakao.com/oauth/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.New(string(bodyBytes))
	}

	var tokenResp auth_dto.KakaoTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func (s *KakaoAuthService) GetUserInfo(accessToken string) (*auth_dto.KakaoProfile, error) {
	req, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.New(string(bodyBytes))
	}

	var profile auth_dto.KakaoProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}

	return &profile, nil
}
