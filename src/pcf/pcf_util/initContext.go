package pcf_util

import (
	"free5gc/lib/openapi/models"
	"free5gc/src/pcf/factory"
	"free5gc/src/pcf/logger"
	"free5gc/src/pcf/pcf_context"

	"github.com/google/uuid"
)

func InitpcfContext(context *pcf_context.PCFContext) {
	config := factory.PcfConfig
	logger.UtilLog.Infof("pcfconfig Info: Version[%s] Description[%s]", config.Info.Version, config.Info.Description)
	configuration := config.Configuration
	context.NfId = uuid.New().String()
	if configuration.PcfName != "" {
		context.Name = configuration.PcfName
	}
	sbi := configuration.Sbi
	context.UriScheme = models.UriScheme(sbi.Scheme)
	context.HttpIPv4Address = "127.0.0.1" // default localhost
	context.HttpIpv4Port = 29507          // default port
	if sbi != nil {
		if sbi.IPv4Addr != "" {
			context.HttpIPv4Address = sbi.IPv4Addr
		}
		if sbi.Port != 0 {
			context.HttpIpv4Port = sbi.Port
		}
	}

	context.TimeFormat = configuration.TimeFormat
	context.DefaultBdtRefId = configuration.DefaultBdtRefId
}
