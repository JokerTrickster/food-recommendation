package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"main/utils/aws"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleConfig = oauth2.Config{}

type tAuthMeta struct {
	GoogleIosID string   `json:"googleIosID" validate:"required"`
	GoogleAndID []string `json:"googleAndID" validate:"required"`
	// KakaoID     int      `json:"kakaoId" validate:"required"`
}

type OAuthData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
}
type AuthProvider uint8

const (
	AuthProviderGoogle = AuthProvider(0)
	AuthProviderKakao  = AuthProvider(1)
	AuthProviderNaver  = AuthProvider(2)
)

var authProviderName map[AuthProvider]string = map[AuthProvider]string{
	AuthProviderGoogle: "google",
	AuthProviderKakao:  "kakao",
	AuthProviderNaver:  "naver",
}
var authMeta tAuthMeta
var httpClient = http.DefaultClient

func InitGoogleOauth() error {
	clientID, err := getClientID()
	if err != nil {
		return err
	}
	clientSecret, err := getClientSecret()
	if err != nil {
		return err
	}
	GoogleConfig = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/v0.2/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	if !Env.IsLocal {
		GoogleConfig.RedirectURL = "https://food-recommendation.jokertrickster.com"
	}
	return nil
}

func getClientID() (string, error) {
	if Env.IsLocal {
		clientID, ok := os.LookupEnv("GOOGLE_CLIENT_ID")
		if !ok {
			return "", fmt.Errorf("GOOGLE_CLIENT_ID not found")
		}
		return clientID, nil
	} else {
		ClientID, err := aws.AwsSsmGetParam("food_google_client_id")
		if err != nil {
			return "", err
		}
		return ClientID, nil
	}
}

func getClientSecret() (string, error) {
	if Env.IsLocal {

		clientSecret, ok := os.LookupEnv("GOOGLE_CLIENT_SECRET")
		if !ok {
			return "", fmt.Errorf("GOOGLE_CLIENT_ID not found")
		}
		return clientSecret, nil

	} else {
		ClientID, err := aws.AwsSsmGetParam("food_google_client_secret")
		if err != nil {
			return "", err
		}
		return ClientID, nil
	}
}

func GoogleValidate(ctx context.Context, token string) (OAuthData, error) {
	claims, iErr := jwtVerifyWithKeySet(ctx, AuthProviderGoogle, token, "https://www.googleapis.com/oauth2/v3/certs")
	if iErr != nil {
		return OAuthData{}, iErr
	}
	aud, okAud := claims["aud"].(string)
	azp, _ := claims["azp"].(string)
	iss, okIss := claims["iss"].(string)
	sub, okSub := claims["sub"].(string)
	email, okEmail := claims["email"].(string)
	if !okAud || !okIss || !okSub || !okEmail ||
		(aud != authMeta.GoogleIosID && isGoogleIdAndNotExisted(aud) && azp != authMeta.GoogleIosID && isGoogleIdAndNotExisted(azp)) ||
		(iss != "accounts.google.com" && iss != "https://accounts.google.com") {
		return OAuthData{}, ErrorMsg(context.TODO(), ErrBadToken, Trace(), fmt.Sprintf("not valid token claims from provider Google - %+v", claims), ErrFromClient)
	}
	oauthData := OAuthData{
		ID:       sub,
		Email:    email,
		Provider: "google",
	}
	return oauthData, nil
}

func isGoogleIdAndNotExisted(key string) bool {
	for _, andKey := range authMeta.GoogleAndID {
		if key == andKey {
			return false
		}
	}
	return true
}

func InitAuth() error {
	authMetaString, iErr := aws.AwsSsmGetParam(fmt.Sprintf("%s_food_oauth", Env.Env))
	if iErr != nil {
		return iErr
	}
	if err := json.Unmarshal([]byte(authMetaString), &authMeta); err != nil {
		return ErrorMsg(context.TODO(), ErrInternalServer, Trace(), fmt.Sprintf("auth meta not available - %+v", authMetaString), ErrFromAwsSsm)
	}
	if err := ValidateStruct(authMeta); err != nil {
		return ErrorMsg(context.TODO(), ErrInternalServer, Trace(), fmt.Sprintf("auth meta not valid - %+v", authMeta), ErrFromAwsSsm)
	}
	return nil
}
