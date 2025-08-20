package logs

import "github.com/starter-go/vlog"

func initLoggerAdapter(ada *vlog.LoggerAdapter) *vlog.LoggerAdapter {

	if ada == nil {
		return nil
	}

	dl := vlog.GetDefaultLogger()

	ada.SetLevelAccepted(vlog.INFO)
	ada.SetSender(ada)
	ada.SetTag("logger-adapter")
	ada.SetTargetLogger(dl)

	return ada
}
