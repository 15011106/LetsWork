package auth_dto

type KakaoTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

type KakaoUserResponse struct {
	ID           int64 `json:"id"`
	KakaoAccount struct {
		Email   string `json:"email"`
		Profile struct {
			Nickname     string `json:"nickname"`
			ProfileImage string `json:"profile_image_url"`
		} `json:"profile"`
	} `json:"kakao_account"`
}

type KakaoProfile struct {
	ID           int64
	Email        string
	Nickname     string
	ProfileImage string
}
