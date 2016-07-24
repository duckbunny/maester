// Copyright © 2016 Jason Smith <jasonrichardsmith@gmail.com>.
//
// Use of this source code is governed by the LGPL-3
// license that can be found in the LICENSE file.

package maester

import (
	"testing"

	"github.com/duckbunny/service"
)

type Maester interface {
	Set(string) ([]byte, error)
	Get(string) ([]byte, error)
	Init(*service.Service) error
}

func TestInit(t *testing.T) {
}

func TestGet(t *testing.T) {
}

func TestAddMaester(t *testing.T) {
}

func TestAddMaesters(t *testing.T) {
}
