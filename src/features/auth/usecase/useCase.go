package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"main/features/auth/model/request"
	"main/utils"
	"main/utils/db/mysql"
	"net/http"
)

func CreateGoogleUserDTO(oauthData utils.OAuthData) *mysql.Users {
	return &mysql.Users{
		Email:    oauthData.Email,
		Password: "",
		Provider: "google",
		Birth:    "1990-01-01",
	}
}

// 영문 + 숫자 6글자 랜덤값 생성
func GeneratePasswordAuthCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func CreateTokenDTO(uID uint, accessToken string, accessTknExpiredAt int64, refreshToken string, refreshTknExpiredAt int64) mysql.Tokens {
	return mysql.Tokens{
		UserID:           uID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTknExpiredAt,
	}
}

func CreateSignupUser(req *request.ReqSignup) mysql.Users {
	return mysql.Users{
		Email:    req.Email,
		Password: req.Password,
		Provider: "email",
		Name:     req.Name,
		Birth:    req.Birth,
		Sex:      req.Sex,
	}
}

func VerifyAccessAndRefresh(req *request.ReqReissue) error {
	accessTokenUserID, accessTokenEmail, err := utils.ParseToken(req.AccessToken)
	if err != nil {
		return err
	}
	refresdhTokenUserID, refreshTokenEmail, err := utils.ParseToken(req.RefreshToken)
	if err != nil {
		return err
	}
	if accessTokenUserID != refresdhTokenUserID || accessTokenEmail != refreshTokenEmail {
		return utils.ErrorMsg(context.TODO(), utils.ErrBadParameter, utils.Trace(), "access token and refresh token are not matched", utils.ErrFromClient)
	}

	if err := utils.VerifyToken(req.RefreshToken); err != nil {
		return err
	}
	return nil
}

func GenerateStateOauthCookie(ctx context.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func getGoogleUserInfo(ctx context.Context, accessToken string) ([]byte, error) {
	token, err := utils.GoogleConfig.Exchange(ctx, accessToken)
	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), fmt.Sprintf("google exchange error %v", err), utils.ErrFromInternal)
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), fmt.Sprintf("bad request google access token %v", err), utils.ErrFromInternal)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
