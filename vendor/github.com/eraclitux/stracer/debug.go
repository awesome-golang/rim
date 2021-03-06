// Copyright (c) 2015 Andrea Masi. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE.txt file.

// +build debug

package stracer

import (
	"encoding/json"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "[STRACER] ", log.Lmicroseconds)

func Traceln(args ...interface{}) {
	logger.Println(args...)
}

func Tracef(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func PrettyStruct(name string, s interface{}) {
	b, _ := json.MarshalIndent(s, "", "  ")
	logger.Println(name, "-", string(b))
}

func init() {
	Traceln("enabled!")
}
