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
	file, err := os.OpenFile("./myapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file, ", err.Error())
	}
	log.SetOutput(file)
	log.Println("Server started")

	return nil
}
