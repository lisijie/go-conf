package goconf_test

import (
	"github.com/lisijie/go-conf"
	"testing"
)

func TestConf(t *testing.T) {

	c, err := goconf.NewConfig("example.ini")

	if err != nil {
		t.Fatal(err)
	}

	if c.GetString("test_str") != "hello world" {
		t.Error("test_str not equal to hello world")
	}

	if c.GetInt("test_int") != 123456 {
		t.Error("test_int not equal to 123456")
	}

	c.SetString("foo", "bar")

	if c.GetString("foo") != "bar" {
		t.Error("foo not equal to bar")
	}

	c.SetInt("foo2", 2015)
	if c.GetInt("foo2") != 2015 {
		t.Error("foo2 not equal to 2015")
	}

}
