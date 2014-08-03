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
	"os"
	"log"
	"bufio"
	"github.com/oikomi/goflv/util"
)

type FlvFileHandle struct {
	file *os.File
	r *bufio.Reader
}

func NewFlvFileHandle () *FlvFileHandle {
	fh := &FlvFileHandle {
		
	}
	
	return fh
}

func (self * FlvFileHandle) FlvOpen(fs *FlvFileSpec) error {
	var err error

	self.file, err = os.Open(fs.FlvName)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	self.r = bufio.NewReader(self.file)
	
	return nil
}

func (self * FlvFileHandle) FlvRead(size int64) ([]byte,  error) {
	buf := make([]byte, size)
	_, err := self.file.Read(buf)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return buf, nil
}

func (self *FlvFileHandle) FlvSeek(offset int64, whence int)  error {
	_, err := self.file.Seek(offset, whence)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	//log.Printf("seek to %d\n", ret)
	/*
	//log.Println(offset)
	buf := make([]byte, offset)
	n, err := self.r.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
		return  err
	}
	
	log.Println("seeking :")
	log.Println(n)
	*/
	
	return nil
}

func (self *FlvFileHandle) FlvFileStat(fs *FlvFileSpec) error {
	fi ,err := self.file.Stat() 
	
	if err != nil {
		log.Fatalln(err.Error())
		return  err
	}
	
	fs.FlvName = fi.Name()
	fs.TotalSize = fi.Size()
	
	return nil
}

func (self *FlvFileHandle) FlvReadHeader(fs *FlvFileSpec) error {
	buf := make([]byte, 9)
	_, err := self.file.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	//log.Println(buf)
	
	fs.Header.Signature = string(buf[0:3])
	fs.Header.Version = buf[3]
	fs.Header.TypeFlags = buf[4]
	fs.Header.DataOffset = util.Byte42Uint32(buf[5:], 0)

	return nil
}

func (self *FlvFileHandle) FlvReadBodyPreviousTagSize(fb *FlvFileBody) error {
	buf := make([]byte, 4)
	_, err := self.file.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fb.PreviousTagSize = util.Byte42Uint32(buf, 0)
	
	//log.Println(fb.PreviousTagSize)
	
	return nil
}

func (self *FlvFileHandle) FlvReadBodyTag(fb *FlvFileBody) error {
	buf := make([]byte, 11)
	_, err := self.file.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	//log.Println(buf)
	fb.Tag.TagType = uint8(buf[0])
	fb.Tag.DataSize = util.Byte32Uint32(buf[1:4], 0)
	fb.Tag.Timestamp = util.Byte32Uint32(buf[4:7], 0)
	fb.Tag.TimestampExtended = uint8(buf[7])
	fb.Tag.StreamID = util.Byte32Uint32(buf[8:11], 0)
	
	return nil
}

func (self *FlvFileHandle) FlvNextBody(pos int64) error {
	err := self.FlvSeek(pos, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	return nil
}

func (self *FlvFileHandle) FlvReadBody(fs *FlvFileSpec) error {
	var num int64
	var pos int64
	var err error
	
	pos = 9
	
	
	
	for fs.TotalSize > pos { 
		fb := new(FlvFileBody)
		err = self.FlvReadBodyPreviousTagSize(fb)
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		
		if pos == (fs.TotalSize - 4) {
			fs.Bodys = append(fs.Bodys, fb)
			return nil
		}
		
		err = self.FlvReadBodyTag(fb)
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		
		pos += (int64)(fb.Tag.DataSize) 
		pos += 15
		
		fs.Bodys = append(fs.Bodys, fb)
		
		err = self.FlvNextBody(pos)
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		num ++
	}

	return nil
}

