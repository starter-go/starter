package common

import "testing"

func TestUUIDGen(t *testing.T) {

	g := &UUIDGenService{}

	g.init()

	u1 := g.Generate("1")
	u2 := g.Generate("1")
	u3 := g.Generate("0")
	u4 := g.GenerateWithMap(map[string]string{"foo": "bar"})

	const f = "uuid[%d]: %v"
	t.Logf(f, 1, u1)
	t.Logf(f, 2, u2)
	t.Logf(f, 3, u3)
	t.Logf(f, 4, u4)
}
