package persist_test

import (
	"os"
	"testing"
	"time"

	"github.com/cheekybits/is"
	"github.com/matryer/persist"
)

type obj struct {
	Name   string
	Number int
	When   time.Time
}

func TestNoFile(t *testing.T) {
	is := is.New(t)

	var v interface{}
	err := persist.Load("no-such-file.conf", &v)
	is.Equal(true, os.IsNotExist(err))

}

func TestPersist(t *testing.T) {
	is := is.New(t)
	defer os.Remove("./file.tmp")

	o := &obj{
		Name:   "Mat",
		Number: 47,
		When:   time.Now(),
	}

	// save it
	err := persist.Save("./file.tmp", o)
	is.NoErr(err)

	// load it
	var o2 obj
	err = persist.Load("./file.tmp", &o2)
	is.NoErr(err)

	is.Equal(o.Name, o2.Name)
	is.Equal(o.Number, o2.Number)
	is.Equal(o.When, o2.When)

}
