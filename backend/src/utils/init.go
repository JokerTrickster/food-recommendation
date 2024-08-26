package utils

import (
	"fmt"
	"log"
	"main/utils/aws"
	_aws "main/utils/aws"
	"main/utils/db/mysql"
	_redis "main/utils/db/redis"
	"os"
)

var GeminiID string

func InitServer() error {
	if err := InitEnv(); err != nil {
		fmt.Sprintf("서버 에러 발생 : %s", err.Error())
		return err
	}

	if err := InitJwt(); err != nil {
		fmt.Sprintf("jwt 초기화 에러 : %s", err.Error())
		return err
	}
	if err := _aws.InitAws(); err != nil {
		fmt.Sprintf("aws 초기화 에러 : %s", err.Error())
		return err
	}

	if err := InitGoogleOauth(); err != nil {
		fmt.Sprintf("google oauth 초기화 에러 : %s", err.Error())
		return err
	}
	if err := mysql.InitMySQL(); err != nil {
		fmt.Sprintf("db 초기화 에러 : %s", err.Error())
		return err
	}
	if err := _redis.InitRedis(); err != nil {
		fmt.Sprintf("redis 초기화 에러 : %s", err.Error())
		return err
	}

	if Env.IsLocal {
		var ok bool
		GeminiID, ok = os.LookupEnv("GEMINI_ID")
		if !ok {
			fmt.Println("GEMINI_ID not found")
		}
	} else {
		var err error
		GeminiID, err = aws.AwsSsmGetParam("food_gemini_id")
		if err != nil {
			fmt.Println(err)
		}
	}

	// 로그 파일 열기 또는 생성
	logFile, err := os.OpenFile("/logs/myapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// 로그 출력 대상을 logFile로 설정
	log.SetOutput(logFile)

	// 로그 메시지 작성
	log.Println("This is a regular log message.")
	log.Println("Another log message.")

	return nil
}
