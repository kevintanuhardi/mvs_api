package webservice

import "github.com/kevintanuhardi/mvs_api/pkg/router"

type WebService interface {
	Run() error
}
type WebRegistrator interface {
	Register(router router.Registrator)
}
