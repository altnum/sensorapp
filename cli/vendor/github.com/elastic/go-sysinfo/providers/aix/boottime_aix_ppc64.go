// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//go:build aix && ppc64
// +build aix,ppc64

package aix

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

// utmp can't be used by "encoding/binary" if generated by cgo,
// some pads will be missing.
type utmp struct {
	User        [256]uint8
	Id          [14]uint8
	Line        [64]uint8
	XPad1       int16
	Pid         int32
	Type        int16
	XPad2       int16
	Time        int64
	Termination int16
	Exit        int16
	Host        [256]uint8
	Xdblwordpad int32
	XreservedA  [2]int32
	XreservedV  [6]int32
}

const (
	typeBootTime = 2
)

// BootTime returns the time at which the machine was started, truncated to the nearest second
func BootTime() (time.Time, error) {
	return bootTime("/etc/utmp")
}

func bootTime(filename string) (time.Time, error) {
	// Get boot time from /etc/utmp
	file, err := os.Open(filename)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get host uptime: cannot open /etc/utmp: %w", err)
	}

	defer file.Close()

	for {
		var utmp utmp
		if err := binary.Read(file, binary.BigEndian, &utmp); err != nil {
			break
		}

		if utmp.Type == typeBootTime {
			return time.Unix(utmp.Time, 0), nil
		}
	}

	return time.Time{}, fmt.Errorf("failed to get host uptime: no utmp record: %w", err)
}
