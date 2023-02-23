package kenv

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/kamalshkeir/kstrct"
)

var envFile = false

func Load(envFiles ...string) {
	var wg sync.WaitGroup
	if len(envFiles) == 0 {
		Load(".env")
	}
	wg.Add(len(envFiles))
	envFile = true
	for _, env := range envFiles {
		go func(env string) {
			defer wg.Done()

			f, err := os.OpenFile(env, os.O_RDONLY, os.ModePerm)
			if err != nil {
				envFile = false
				fmt.Println(err)
				return
			}
			defer f.Close()
			r := bufio.NewScanner(f)

			for r.Scan() {
				sp := strings.Split(r.Text(), "=")
				if len(sp) != 2 || r.Text()[0] == '#' {
					continue
				}
				sp[0] = strings.TrimSpace(sp[0])
				sp[1] = strings.TrimSpace(sp[1])
				err := os.Setenv(sp[0], sp[1])
				if err != nil {
					continue
				}
			}
		}(env)
	}
	wg.Wait()
}

// FillStructFromEnv fill the struct from env
func Fill(structure interface{}) error {
	if !envFile {
		fmt.Println("env: the file wasn't loaded. The environment value would be reading from os or set by default")
	}
	inputType := reflect.TypeOf(structure)
	if inputType != nil {
		if inputType.Kind() == reflect.Ptr {
			if inputType.Elem().Kind() == reflect.Struct {
				return fillStructFromEnv(reflect.ValueOf(structure).Elem())
			} else {
				return errors.New("env: element is not pointer to struct")
			}
		}
	}
	return errors.New("env: invalid structure")
}

// fillStructFromEnv sets a reflected struct fields with the equivalent OS environment variables.
func fillStructFromEnv(s reflect.Value) error {
	errored := []string{}
	for i := 0; i < s.NumField(); i++ {
		if t, exist := s.Type().Field(i).Tag.Lookup("kenv"); exist {
			// tag exist
			tag := t
			defau := ""
			required := false
			if strings.Contains(t, "|") {
				sp := strings.Split(t, "|")
				if len(sp) == 2 {
					tag = sp[0]
					defau = sp[1]

				}
			} else {
				required = true
			}
			if osv := os.Getenv(strings.TrimSpace(tag)); osv != "" {
				kstrct.SetReflectFieldValue(s.Field(i), osv)
			} else {
				if !required {
					fmt.Printf("env: %s don't set. %s will set by default \n", tag, tag)
					kstrct.SetReflectFieldValue(s.Field(i), defau)
				} else {
					errored = append(errored, t)
				}
			}
		} else if s.Type().Field(i).Type.Kind() == reflect.Struct {
			// tag not exist, check if struct
			if err := fillStructFromEnv(s.Field(i)); err != nil {
				return err
			}
		} else if s.Type().Field(i).Type.Kind() == reflect.Ptr {
			if !s.Field(i).IsZero() && s.Field(i).Elem().Type().Kind() == reflect.Struct {
				if err := fillStructFromEnv(s.Field(i).Elem()); err != nil {
					return err
				}
			}
		}
	}
	if len(errored) > 0 {
		return errors.New(strings.Join(errored, ",") + " required and has not been found")
	}
	return nil
}
