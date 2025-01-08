package notification

import (
	"github.com/minisource/common_go/logging"
	"github.com/minisource/notifire/api/v1/dto"
	"github.com/minisource/notifire/config"
	"github.com/minisource/notifire/internal/platform/sms"
)

type SMSService struct {
	logger logging.Logger
	cfg    *config.Config
}

func NewSMSService(cfg *config.Config) *SMSService {
	logger := logging.NewLogger(&cfg.Logger)
	return &SMSService{
		logger: logger,
		cfg:    cfg,
	}
}

func (s *SMSService) SendNotification(req dto.SMSRequest) error {
	if(s.cfg.SMS.NotSendSms) {
		return nil
	}
	return sms.SendSMS(&s.cfg.SMS, s.logger, req.To, req.Body, req.Template)
}