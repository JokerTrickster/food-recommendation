package usecase

import (
	"context"
	"log"
	_interface "main/features/user/model/interface"
	"main/features/user/model/request"
	"main/utils"
	"time"

	"firebase.google.com/go/messaging"
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

	// 메시지 생성
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: req.Title,
			Body:  req.Message,
		},
	}

	// 메시지 전송
	response, err := utils.MessageClient.Send(ctx, message)
	if err != nil {
		log.Printf("error sending message: %v", err)
		return err
	}

	log.Printf("Successfully sent message: %s", response)
	// 4. 메시지를 저장한다.

	return nil
}
