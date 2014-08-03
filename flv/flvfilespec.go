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

package flv

import (
	//"log"
)

type FlvFileHeader struct {
	Signature string
	Version uint8
	TypeFlags uint8
	DataOffset uint32
}

type FlvFiletags struct {
	TagType uint8
	DataSize uint32
	Timestamp uint32
	TimestampExtended uint8
	StreamID uint32
}

type FlvFileBody struct {
	PreviousTagSize uint32
	Tag FlvFiletags
}


type FlvFileSpec struct {
	FlvName string
	TotalSize int64
	Header FlvFileHeader
	Bodys []*FlvFileBody
}

func NewFlvFileSpec (name string) *FlvFileSpec {
	fs := &FlvFileSpec {
		FlvName : name,
		TotalSize : 0,
		
		Bodys : make([]*FlvFileBody, 0),
	}
	
	return fs
}





