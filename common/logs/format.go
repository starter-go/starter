package logs

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/starter-go/vlog"
)

// FormatData ...
type FormatData struct {
	LEVEL string
	TAG   string

	YEAR  string
	MONTH string
	DAY   string
	HH    string
	MM    string
	SS    string
	SSS   string
	ZONE  string
}

// Init ...
func (inst *FormatData) Init(msg *vlog.Message) {

	t := msg.Timestamp
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hh := t.Hour()
	mm := t.Minute()
	ss := t.Second()
	sss := int(t.UnixMilli() % 1000)

	inst.LEVEL = msg.Level.String()
	inst.TAG = msg.Tag

	inst.YEAR = inst.i2s(year, 4)
	inst.MONTH = inst.i2s(month, 2)
	inst.DAY = inst.i2s(day, 2)

	inst.HH = inst.i2s(hh, 2)
	inst.MM = inst.i2s(mm, 2)
	inst.SS = inst.i2s(ss, 2)

	inst.SSS = inst.i2s(sss, 3)
	inst.ZONE = inst.formatZone(t)
}

func (inst *FormatData) formatZone(t time.Time) string {
	_, sec := t.Zone()
	sign := "+"
	if sec < 0 {
		sec = -sec
		sign = "-"
	}
	h := sec / 3600
	m := (sec / 60) % 60
	hh := inst.i2s(h, 2)
	mm := inst.i2s(m, 2)
	return "UTC" + sign + hh + ":" + mm
}

func (inst *FormatData) i2s(i, width int) string {
	const max = 64
	if width > max {
		width = max
	}
	str := strconv.Itoa(i)
	if i < 0 {
		return str
	}
	for len(str) < width {
		str = "0" + str
	}
	return str
}

////////////////////////////////////////////////////////////////////////////////

// Formatter 是日志消息格式化工具
type Formatter struct {
	templ  *template.Template
	format string
	mu     sync.Mutex
}

// Format 格式化日志消息
func (inst *Formatter) Format(data *FormatData) string {

	inst.mu.Lock()
	defer inst.mu.Unlock()

	str, err := inst.formatWithError(data)
	if err != nil {
		return err.Error()
	}
	return str
}

func (inst *Formatter) formatWithError(data *FormatData) (string, error) {
	t := inst.templ
	if t == nil {
		t2, err := inst.load()
		if err != nil {
			return "", err
		}
		t = t2
		inst.templ = t2
	}
	if t == nil {
		return "", fmt.Errorf("no vlog format template")
	}
	builder := &strings.Builder{}
	err := t.Execute(builder, data)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func (inst *Formatter) load() (*template.Template, error) {
	const name = "logs.Formatter"
	return template.New(name).Parse(inst.format)
}
