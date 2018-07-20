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
		t.Error("test_str not equal to 'hello world'")
	}

	if c.GetInt("test_int") != 123456 {
		t.Error("test_int not equal to 123456")
	}

	if c.GetString("noexists", "default value") != "default value" {
		t.Error("noexists not equal to 'default value'")
	}

	if c.GetSection("foo").GetString("bar") != "abc" {
		t.Error("foo.bar not equal to abc")
	}
}
