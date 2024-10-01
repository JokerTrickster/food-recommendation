package usecase

import (
	"context"
	_interface "main/features/system/model/interface"
	"main/features/system/model/request"
	"main/utils/aws"
	"strconv"
	"time"
)

type ReportSystemUseCase struct {
	Repository     _interface.IReportSystemRepository
	ContextTimeout time.Duration
}

func NewReportSystemUseCase(repo _interface.IReportSystemRepository, timeout time.Duration) _interface.IReportSystemUseCase {
	return &ReportSystemUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *ReportSystemUseCase) Report(c context.Context, uID uint, req *request.ReqReport) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//디비에 저장
	reportDTO := CreateReportDTO(uID, req)
	err := d.Repository.SaveReport(ctx, reportDTO)
	if err != nil {
		return err
	}
	//이메일 전송
	reqReport := &aws.ReqReportSES{
		UserID: strconv.Itoa(int(uID)),
		Reason: string(req.Reason),
	}
	go aws.EmailSendReport([]string{"pkjhj485@gmail.com", "dtw7225@naver.com"}, reqReport)

	return nil
}
