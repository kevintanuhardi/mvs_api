package webservice

import "gitlab.warungpintar.co/sales-platform/brook/pkg/router"

type WebService interface {
	Run() error
}
type WebRegistrator interface {
	Register(router router.Registrator)
}
