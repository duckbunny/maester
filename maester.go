// Copyright © 2016 Jason Smith <jasonrichardsmith@gmail.com>.
//
// Use of this source code is governed by the LGPL-3
// license that can be found in the LICENSE file.

package maester

import (
	"flag"
	"fmt"
	"os"

	"github.com/duckbunny/service"
)

// MaesterTypes are a collection of maester interface identified by strings
// associated with a maester type
var MaesterTypes map[string]Maester = make(map[string]Maester)

// Flags to be parsed for setting pool and declaration interfaces.
var maester string

var coreMaester Maester

func init() {
	flag.StringVar(&maester, "maester", os.Getenv("MAESTER"), "Maester to load config values.")
}

type Maester interface {
	Set(string) ([]byte, error)
	Get(string) ([]byte, error)
	Init(*service.Service) error
}

func Init() error {
	if !flag.Parsed() {
		flag.Parse()
	}
	if maester != "" {
		if _, ok := MaesterTypes[maester]; !ok {
			err := fmt.Errorf(
				"Attempt to utilize unrecognized Maester mechanism %v", maester)
			return err
		}
		coreMaester = MaesterTypes[maester]
	}
	s, err := service.This()
	if err != nil {
		return err
	}
	return coreMaester.Init(s)
}

// Get a foreign microservice definition.
func Get(s string) ([]byte, error) {
	return coreMaester.Get(s)
}

// AddMaester adds a Maester type.
func AddMaester(key string, m Maester) {
	MaesterTypes[key] = m
}

// AddMaesters adds Multiple Maesters at once.
func AddMaesters(ms map[string]Maester) {
	for key, maester := range ms {
		MaesterTypes[key] = maester
	}
}
