package request

// 닉네임, 비밀번호, 이메일 정도만 정보를 받는다.
type ReqSignup struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
