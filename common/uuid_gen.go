package common

import (
	"bytes"

	"crypto/sha1"
	"io"
	"math/rand"
	"sort"
	"strconv"
	"sync"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
)

// UUIDGenService 是UUID生成服务
type UUIDGenService struct {
	//starter:component
	_as func(lang.UUIDGenerator) //starter:as("#")

	mutex sync.Mutex
	index int64
	t0    lang.Time
	nonce io.Reader
	prev  []byte
}

func (inst *UUIDGenService) _impl() (lang.UUIDGenerator, application.Lifecycle) {
	return inst, inst
}

// Life ...
func (inst *UUIDGenService) Life() *application.Life {
	return &application.Life{}
}

func (inst *UUIDGenService) init() error {

	now := lang.Now()
	src := rand.NewSource(now.Int())
	reader := rand.New(src)
	nonce := make([]byte, 32)

	reader.Read(nonce)

	inst.t0 = now
	inst.nonce = reader
	inst.prev = nonce
	return nil
}

// Generate ...
func (inst *UUIDGenService) Generate(params ...string) lang.UUID {
	return inst.gen(nil, params)
}

// GenerateWithMap ...
func (inst *UUIDGenService) GenerateWithMap(params map[string]string) lang.UUID {
	return inst.gen(params, nil)
}

func (inst *UUIDGenService) gen(p1 map[string]string, p2 []string) lang.UUID {

	if p1 == nil {
		p1 = make(map[string]string)
	}
	inst.prepare(p1)

	for k, v := range p1 {
		text := k + " = " + v + "\n"
		p2 = append(p2, text)
	}
	sort.Strings(p2)
	builder := &bytes.Buffer{}
	for _, row := range p2 {
		builder.WriteString(row)
	}
	builder.Write(inst.prev)
	sum := sha1.Sum(builder.Bytes())
	data := sum[:]
	inst.prev = data
	return lang.NewUUID(data)
}

func (inst *UUIDGenService) prepare(m map[string]string) {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	now := lang.Now()
	idx := inst.index
	t0 := inst.t0
	nonce := make([]byte, 32)
	nReader := inst.nonce

	inst.index++

	if nReader != nil {
		nReader.Read(nonce)
	}

	m["time0"] = t0.String()
	m["time1"] = now.String()
	m["index"] = strconv.FormatInt(idx, 10)
	m["nonce"] = lang.HexFromBytes(nonce).String()
}
