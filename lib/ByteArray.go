package ByteArray

import (
	"bytes"
	"encoding/binary"
)

type ByteArray struct {
	Bytes *bytes.Buffer
}

// Constructors
func NewByteArrayEmpty() *ByteArray {
	return &ByteArray{Bytes: bytes.NewBuffer([]byte{})}
}

func NewByteArray(bytes *bytes.Buffer) *ByteArray {
	return &ByteArray{Bytes: bytes}
}

// Standard methods
func (bt *ByteArray) Clear() *ByteArray {
	bt.Bytes.Reset()
	return bt
}

func (bt *ByteArray) Get_Bytes() []byte {
	return bt.Bytes.Bytes()
}

func (bt *ByteArray) Len() int {
	return bt.Bytes.Len()
}

// Write methods
func (bt *ByteArray) Write_Byte(value byte) *ByteArray {
	bt.Bytes.WriteByte(value)
	return bt
}

func (bt *ByteArray) Write_Short(value uint16) *ByteArray {
	bt.Bytes.Write(binary.BigEndian.AppendUint16([]byte{}, value))
	return bt
}

func (bt *ByteArray) Write_Int(value int) *ByteArray {
	bt.Bytes.Write(binary.BigEndian.AppendUint32([]byte{}, uint32(value)))
	return bt
}

// Read methods
func (bt *ByteArray) Read_Byte() byte {
	value, err := bt.Bytes.ReadByte()

	if err != nil {
		panic(err)
	}

	return value
}

func (bt *ByteArray) Read_Short() uint16 {
	data := make([]byte, 2)
	bt.Bytes.Read(data)
	return binary.BigEndian.Uint16(data)
}

func (bt *ByteArray) Read_Int() int {
	data := make([]byte, 4)
	bt.Bytes.Read(data)
	return int(binary.BigEndian.Uint32(data))
}
