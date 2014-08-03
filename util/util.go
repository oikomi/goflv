//
// Copyright 2014 Hong Miao. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	//"log"
)

const (
	BigEndian = 0
)

func Byte42Uint32(data []byte, endian int) uint32 {
	var i uint32
	if 0 == endian {
		i = uint32(uint32(data[3]) + uint32(data[2])<<8 + uint32(data[1])<<16 + uint32(data[0])<<24)
	}
	
	if 1 == endian {
		i = uint32(uint32(data[3]) + uint32(data[2])<<8 + uint32(data[1])<<16 + uint32(data[0])<<24)
	}

	return i
}

func Byte32Uint32(data []byte, endian int) uint32 {
	var i uint32
	//log.Println("Byte32Uint32 data = ")
	//log.Println(data)
	if 0 == endian {
		i = uint32(uint32(data[2]) + uint32(data[1])<<8 + uint32(data[0])<<16)
	}
	
	if 1 == endian {
		i = uint32(uint32(data[2]) + uint32(data[1])<<8 + uint32(data[0])<<16)
	}
	//log.Println(i)

	return i
}