package app

import (
	amf_context "github.com/f0lkert/amf/internal/context"
	"github.com/f0lkert/amf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *amf_context.AMFContext
	Config() *factory.Config
}
