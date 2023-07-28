package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/kamalshkeir/kenv"
)

type Config struct {
	Host   string `kenv:"HOST|localhost"`
	Secret string `kenv:"SECRET|"`
	Id     string `kenv:"ID|"`
	Code   string `kenv:"CODE|"`
}

func Test_load_env_file(t *testing.T) {
	cfg := Config{}

	kenv.Load(".env")
	kenv.Fill(&cfg)

	if cfg.Host != "127.0.0.1" {
		t.Logf("config %+v \n", cfg)
		t.Fail()
		return
	}

	t.Logf("config %+v \n", cfg)
}

func Test_load_env_trim_space(t *testing.T) {
	cfg := Config{}

	kenv.Load(".env")
	kenv.Fill(&cfg)

	if cfg.Id != "100" {
		t.Logf("config %+v \n", cfg)
		t.Fail()
		return
	}

	t.Logf("config %+v \n", cfg)
}

func Test_load_env_base64_val(t *testing.T) {
	cfg := Config{}

	kenv.Load(".env")
	kenv.Fill(&cfg)

	if cfg.Secret != "MiU3czRkS3U8Zg==" {
		t.Logf("config %+v \n", cfg)
		t.Fail()
		return
	}

	t.Logf("config %+v \n", cfg)
}

func Test_env_file_not_found(t *testing.T) {
	cfg := Config{}

	t.Setenv("HOST", "TEST-ENV")

	kenv.Fill(&cfg)

	host := os.Getenv("HOST")

	if host != "TEST-ENV" {
		t.Logf("config %+v \n", cfg)
		t.Fail()
		return
	}

	t.Logf("config %+v \n", cfg)

}

func Test_load_default_tags_no_file(t *testing.T) {
	cfg := Config{}

	t.Setenv("HOST", "")

	kenv.Fill(&cfg)

	fmt.Printf("config %+v \n", cfg)

	if cfg.Host != "localhost" {
		t.Logf("config %+v \n", cfg)
		t.Fail()
		return
	}
}
