package usecase

import (
	"context"
	"fmt"
	"log"
	_interface "main/features/user/model/interface"
	"main/features/user/model/request"
	"main/utils/aws"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"
)

type MessageUserUseCase struct {
	Repository     _interface.IMessageUserRepository
	ContextTimeout time.Duration
}

func NewMessageUserUseCase(repo _interface.IMessageUserRepository, timeout time.Duration) _interface.IMessageUserUseCase {
	return &MessageUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *MessageUserUseCase) Message(c context.Context, uID uint, req *request.ReqMessageUser) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// 1. 알람 여부를 체크한다.

	// 2. 푸시 토큰을 가져온다.
	token, err := d.Repository.FindOnePushToken(ctx, uID)
	if err != nil {
		return err
	}

	// 3. 푸시를 보낸다.

	// 4. 메시지를 저장한다.

	serviceKey, err := aws.AwsSsmGetParam("firebase_service_key")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 서비스 계정 JSON 키를 byte 배열로 변환합니다.
	credentials := []byte(serviceKey)

	opt := option.WithCredentialsJSON(credentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
		return err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v", err)
		return err
	}

	// 메시지 생성
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: req.Title,
			Body:  req.Message,
		},
	}

	// 메시지 전송
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Printf("error sending message: %v", err)
		return err
	}

	log.Printf("Successfully sent message: %s", response)

	return nil
}
