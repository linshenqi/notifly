package notify

import (
	"github.com/linshenqi/notifly/src/services/wechat"
	"github.com/linshenqi/sptty"
)

type Service struct {
	app         sptty.Sptty
	controllers *Controllers

	wechat *wechat.Service
}

func (s *Service) Init(service sptty.Sptty) error {
	s.app = service
	s.wechat = service.GetService("wechat").(*wechat.Service)

	s.controllers = &Controllers{
		service: s,
	}

	s.app.AddRoute("POST", "/customer-msgs", s.controllers.postCustomerMsg)
	s.app.AddRoute("POST", "/template-msgs", s.controllers.postTemplateMsg)

	return nil
}

func (s *Service) Release() {

}

func (s *Service) Enable() bool {
	return true
}
