package utils

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

// 프론트엔드 받을 에러 형식
type ResError struct {
	ErrType string `json:"errType,omitempty"`
	Msg     string `json:"msg,omitempty"`
}

// 에러 로깅을 위한 에러 형식
type Err struct {
	HttpCode int    `json:"httpCode,omitempty"`
	ErrType  string `json:"errType,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Trace    string `json:"trace,omitempty"`
	From     string `json:"from,omitempty"`
}

// 에러 타입을 구분
type ErrType string

// 에러가 어디서 발생했는지 확인용
type IErrFrom string

const (
	ErrFromClient   = IErrFrom("client")
	ErrFromInternal = IErrFrom("internal")
	ErrFromMongoDB  = IErrFrom("mongoDB")
	ErrFromMysqlDB  = IErrFrom("mysqlDB")
	ErrFromAws      = IErrFrom("aws")
	ErrFromAwsS3    = IErrFrom("aws_s3")
	ErrFromAwsSsm   = IErrFrom("aws_ssm")
	ErrFromNaver    = IErrFrom("naver")
)

// basic error
const (
	ErrBadParameter   = ErrType("PARAM_BAD")
	ErrNotFound       = ErrType("NOT_FOUND")
	ErrBadToken       = ErrType("TOKEN_BAD")
	ErrInternalServer = ErrType("INTERNAL_SERVER")
	ErrInternalDB     = ErrType("INTERNAL_DB")
	ErrPartner        = ErrType("PARTNER")
)

// game error
const (
	ErrNotAllUsersReady = ErrType("NOT_ALL_USERS_READY")
	ErrNotOwner         = ErrType("NOT_OWNER")
	ErrNotFirstPlayer   = ErrType("NOT_FIRST_PLAYER")
	ErrNotLoanCard      = ErrType("NOT_LOAN_CARD")
	ErrNotEnoughCard    = ErrType("NOT_ENOUGH_CARD")
	ErrNotEnoughCond    = ErrType("NOT_ENOUGH_CONDITION")
)

// room error
const (
	ErrUserNotFound       = ErrType("USER_NOT_FOUND")
	ErrInvalidAccessToken = ErrType("INVALID_ACCESS_TOKEN")
	ErrUserAlreadyExisted = ErrType("USER_ALREADY_EXISTED")
	ErrBadRequest         = ErrType("BAD_REQUEST")
	ErrRoomNotFound       = ErrType("ROOM_NOT_FOUND")
	ErrRoomFull           = ErrType("ROOM_FULL")
	ErrPlayerStateFailed  = ErrType("PLAYER_STATE_CHANGE_FAILED")
	ErrRoomUserNotFound   = ErrType("ROOM_USER_NOT_FOUND")
)

// auth error

// basic , game, room, auth error mapping
var ErrHttpCode = map[string]int{
	//400
	"PARAM_BAD":            http.StatusBadRequest,
	"USER_ALREADY_EXISTED": http.StatusBadRequest,
	"BAD_REQUEST":          http.StatusBadRequest,
	"NOT_ALL_USERS_READY":  http.StatusBadRequest,
	"NOT_OWNER":            http.StatusBadRequest,
	"NOT_FIRST_PLAYER":     http.StatusBadRequest,
	"ROOM_NOT_FOUND":       http.StatusBadRequest,
	"ROOM_USER_NOT_FOUND":  http.StatusBadRequest,
	"USER_NOT_FOUND":       http.StatusBadRequest,
	"ROOM_FULL":            http.StatusBadRequest,
	"NOT_LOAN_CARD":        http.StatusBadRequest,
	"NOT_ENOUGH_CARD":      http.StatusBadRequest,
	"NOT_ENOUGH_CONDITION": http.StatusBadRequest,

	//401
	"TOKEN_BAD":            http.StatusUnauthorized,
	"INVALID_ACCESS_TOKEN": http.StatusUnauthorized,
	//403
	"PARTNER": http.StatusForbidden,

	//404
	"NOT_FOUND": http.StatusNotFound,

	//500
	"INTERNAL_SERVER":            http.StatusInternalServerError,
	"INTERNAL_DB":                http.StatusInternalServerError,
	"PLAYER_STATE_CHANGE_FAILED": http.StatusInternalServerError,
}

func ErrorParsing(data string) Err {
	slice := strings.Split(data, "|")
	result := Err{
		HttpCode: ErrHttpCode[slice[0]],
		ErrType:  slice[0],
		Trace:    slice[1],
		Msg:      slice[2],
		From:     slice[3],
	}
	return result
}

func ErrorMsg(ctx context.Context, errType ErrType, trace string, msg string, from IErrFrom) error {

	return fmt.Errorf("%s|%s|%s|%s", errType, trace, msg, from)
}

func (e ErrType) New(errType string, msg string) *ResError {
	return &ResError{ErrType: errType, Msg: msg}
}

func Trace() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	_, line := runtime.FuncForPC(pc).FileLine(pc)
	return fmt.Sprintf("%s.L%d", funcName, line)
}
