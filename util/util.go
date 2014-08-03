package util

import (
	"log"
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
	log.Println("Byte32Uint32 data = ")
	log.Println(data)
	if 0 == endian {
		i = uint32(uint32(data[2]) + uint32(data[1])<<8 + uint32(data[0])<<16)
	}
	
	if 1 == endian {
		i = uint32(uint32(data[2]) + uint32(data[1])<<8 + uint32(data[0])<<16)
	}
	log.Println(i)

	return i
}