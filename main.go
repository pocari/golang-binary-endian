package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func writeFile(filename string, bytes []byte) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(bytes)
}

func processFile(filename string, order binary.ByteOrder) {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var data int32
	for {
		err := binary.Read(f, order, &data)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", data)
	}
}
func main() {
	writeFile("output_big.dat", []byte{0x12, 0x34, 0xAB, 0xCD})
	writeFile("output_lit.dat", []byte{0xcd, 0xab, 0x34, 0x12})

	processFile("output_lit.dat", binary.LittleEndian)
	processFile("output_big.dat", binary.BigEndian)
}
