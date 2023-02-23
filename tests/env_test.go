package tests

import (
	"os"
	"testing"

	"github.com/kamalshkeir/kenv"
)

type Config struct {
	Host string `kenv:"HOST|localhost"`
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

	kenv.Fill(&cfg)

	if cfg.Host != "localhost" {
		t.Logf("config %+v \n", cfg)
		t.Fail()
		return
	}

	t.Logf("config %+v \n", cfg)
}
