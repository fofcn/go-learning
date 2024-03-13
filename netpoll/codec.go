package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Codec 接口定义了 Encode 和 Decode 方法
type Codec interface {
	Encode(frame *Frame) ([]byte, error)
	Decode(data []byte) (*Frame, error)
}

// Frame 定义了数据传输的帧结构
type Frame struct {
	Version   uint8 // 协议版本
	Type      uint8 // 帧类型
	HeaderLen uint32
	Headers   map[string][]byte // 变长协议头（采用 Length-Value 编码）
	Payload   []byte            // 有效负载数据
}

// LengthBaseCodec 实现 Codec 接口
type LengthBaseCodec struct{}

func (c *LengthBaseCodec) Encode(frame *Frame) ([]byte, error) {
	buf := new(bytes.Buffer)

	// 编码定长字段，使用 Varint 编码
	buf.Write(EncodeVarint(uint64(frame.Version)))
	buf.Write(EncodeVarint(uint64(frame.Type)))

	// 编码变长协议头
	headerBuf := new(bytes.Buffer)
	for key, value := range frame.Headers {
		headerBuf.Write(EncodeVarint(uint64(len(key))))
		headerBuf.WriteString(key)
		headerBuf.Write(EncodeVarint(uint64(len(value))))
		headerBuf.Write(value)
	}
	// 变长协议头写入buf中
	buf.Write(EncodeVarint(uint64(headerBuf.Len())))
	buf.Write(headerBuf.Bytes())

	// 编码 Payload
	if _, err := buf.Write(frame.Payload); err != nil {
		return nil, err
	}

	// 实际编码的时候只需要调用Writer.Write()
	// 先写EncodeVarint(uint64(buf.Available()))这几个字节就行了
	lvBuf := new(bytes.Buffer)
	lvBuf.Write(EncodeVarint(uint64(buf.Available())))
	lvBuf.Write(buf.Bytes())

	return lvBuf.Bytes(), nil
}

func EncodeVarint(variable uint64) []byte {
	var cmdBuf [binary.MaxVarintLen64]byte
	encodeLen := binary.PutUvarint(cmdBuf[:], variable)
	return cmdBuf[:encodeLen]
}

// Decode 对数据进行解码，返回 Frame 指针
func (c *LengthBaseCodec) Decode(data []byte) (*Frame, error) {
	buf := bytes.NewReader(data)
	frame := &Frame{
		Headers: make(map[string][]byte),
	}

	// 实际编码中，调用Reader
	// 可以一个一个字节的读取
	// 判断这个字节最高位是否为0
	// 这个过程binary已经封装好了
	// len, err := binary.ReadUvarint(reader)
	len, _ := binary.ReadUvarint(buf)
	fmt.Printf("frame length: %d\n", len)

	// 解码定长字段，使用 Varint 解码
	version, _ := binary.ReadUvarint(buf)
	frame.Version = uint8(version)

	frameType, _ := binary.ReadUvarint(buf)
	frame.Type = uint8(frameType)

	// 解码变长协议头
	headerLen, _ := binary.ReadUvarint(buf)
	headLen := int(headerLen)
	for headLen > 0 {
		// 解析key长度
		keyLen64, _ := binary.ReadUvarint(buf)
		keyLen := int(keyLen64)
		// 这里我们已经知道了key的长度为1，所以就直接减1了，实际开发需要计算

		// 解析Key的值
		key := make([]byte, keyLen)
		buf.Read(key)

		headLen -= (keyLen + 1)

		// 解析Value的长度
		valLen64, _ := binary.ReadUvarint(buf)
		valLen := int(valLen64)
		// 解析Value的值
		value := make([]byte, valLen)
		buf.Read(value)
		frame.Headers[string(key)] = value
		headLen -= (valLen + 1)
	}

	// 其余的全部是 Payload
	frame.Payload, _ = io.ReadAll(buf)

	return frame, nil
}

func main() {
	// 示例使用
	codec := LengthBaseCodec{}
	frame := &Frame{
		Version: 1,
		Type:    10,
		Headers: map[string][]byte{
			"Key1": []byte("Value1"),
			"Key2": []byte("Value2"),
		},
		Payload: []byte("Hello, World!"),
	}

	encodedFrame, _ := codec.Encode(frame)
	fmt.Println("Encoded Frame:", encodedFrame)

	decodedFrame, _ := codec.Decode(encodedFrame)
	for key, value := range decodedFrame.Headers {
		fmt.Println("Header:", key, "Value:", string(value))
	}
	fmt.Println("Version:", decodedFrame.Version)
	fmt.Println("Type:", decodedFrame.Type)
	fmt.Println("Payload:", string(decodedFrame.Payload))
}
