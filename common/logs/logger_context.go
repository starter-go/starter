package logs

import (
	"fmt"
	"time"

	"github.com/starter-go/vlog"
)

// LoggerContext ...
type LoggerContext struct {
	sender       any
	tag          string
	acceptLevels vlog.Level
	chain        vlog.MessageFilterChain
}

// Clone ...
func (inst *LoggerContext) Clone() *LoggerContext {
	lc := &LoggerContext{}
	*lc = *inst
	return lc
}

// IsLevelEnabled ...
func (inst *LoggerContext) IsLevelEnabled(l vlog.Level) bool {
	return l >= inst.acceptLevels
}

// LogWithLevel ...
func (inst *LoggerContext) LogWithLevel(level vlog.Level, format string, args ...any) {
	msg := &vlog.Message{}
	msg.Level = level
	msg.Format = format
	msg.Arguments = args
	msg.Timestamp = time.Now()
	msg.Tag = inst.tag
	msg.Sender = inst.sender
	inst.LogWithMessage(msg)
}

// LogWithMessage ...
func (inst *LoggerContext) LogWithMessage(msg *vlog.Message) {

	if msg == nil {
		return
	}

	if !inst.IsLevelEnabled(msg.Level) {
		return
	}

	chain := inst.chain
	if chain == nil {
		fmt.Println("[warn!] LoggerContext.chain is nil")
		return
	}
	chain.DoFilter(msg)
}
