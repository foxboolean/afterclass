package main

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// message Websocket（JSON Frame)
type message struct {
	ver  int16	// 协议版本号
	op 	 int32	// 指令
	seq  int32  // 序列号（服务端返回和客户端发送一一对应）
	body []byte // 授权令牌，用于检验获取用户真实用户Id
}

const (
	packageLen   int = 4
	headerLen    int = 2
	verLen       int = 2
	operationLen int = 4
	seqLen       int = 4
)

func (m *message)Decode(msg []byte) error {
	if len(msg) < packageLen + headerLen {
		return errors.New("illegal msg: len is too short")
	}
	cur := 0
	recPackageLen := BytesToInt32(msg[cur:cur+packageLen])
	if  recPackageLen < 16 {
		return errors.New("illegal msg: package length is too short")
	}
	cur += packageLen

	recHeaderLen := BytesToInt16(msg[cur:cur+headerLen])
	if recHeaderLen != 16 {
		return errors.New("illegal msg: header length should be 16")
	}
	cur += headerLen

	m.ver = BytesToInt16(msg[cur:cur+verLen])
	cur += verLen

	m.op = BytesToInt32(msg[cur:cur+operationLen])
	cur += operationLen

	m.seq = BytesToInt32(msg[cur:cur+seqLen])
	cur += seqLen

	m.body = msg[recHeaderLen:recPackageLen]
	return nil
}

func (m *message)Encode() []byte {
	packLen := int32(16 + len(m.body))
	totalLenByte := Int32ToBytes(packLen)
	headerLenByte := Int16ToBytes(16)
	verByte := Int16ToBytes(m.ver)
	operationByte := Int32ToBytes(m.op)
	seqByte := Int32ToBytes(m.seq)
	msg := ToBytes(totalLenByte, headerLenByte, verByte, operationByte, seqByte)
	msg = append(msg, m.body...)
	return msg
}

func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToInt16(b []byte) int16 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int16
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func Int32ToBytes(n int32) []byte {
	tmp := n
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Int16ToBytes(n int16) []byte {
	tmp := n
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func ToBytes(bytes ...[]byte) []byte {
	if len(bytes) < 2 {
		return nil
	}
	res := append(bytes[0], bytes[1]...)
	for i := 2; i < len(bytes); i++ {
		res = append(res, bytes[i]...)
	}
	return res
}

