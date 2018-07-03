package main

import (
	"bytes"
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
	// 明示的にbyte列を指定して実験
	writeFile("output_big.dat", []byte{0x12, 0x34, 0xAB, 0xCD})
	writeFile("output_lit.dat", []byte{0xcd, 0xab, 0x34, 0x12})

	processFile("output_lit.dat", binary.LittleEndian)
	processFile("output_big.dat", binary.BigEndian)

	// 同じ4byte整数がendianによって異なるbyte列としてファイルに出力され、
	// それぞれに応じたendianで読み込むことによって同じ数字を復元する
	buf1 := new(bytes.Buffer)
	// ここは環境依存のintでは無理で明示的に型を指定しないとbinary.Writeでエラーになる
	err := binary.Write(buf1, binary.LittleEndian, int32(0x1234ABCD))
	if err != nil {
		panic(err)
	}
	writeFile("output_lit_int.dat", buf1.Bytes())

	buf2 := new(bytes.Buffer)
	err2 := binary.Write(buf2, binary.BigEndian, int32(0x1234ABCD))
	if err2 != nil {
		panic(err2)
	}
	writeFile("output_big_int.dat", buf2.Bytes())

	processFile("output_lit_int.dat", binary.LittleEndian)
	processFile("output_big_int.dat", binary.BigEndian)
}
