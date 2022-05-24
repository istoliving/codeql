// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for github.com/gobwas/ws, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: github.com/gobwas/ws (exports: ; functions: ReadFrame,WriteFrame,NewTextFrame,Dial)

// Package ws is a stub of github.com/gobwas/ws, generated by depstubber.
package ws

import (
	bufio "bufio"
	context "context"
	io "io"
	net "net"
)

func Dial(_ context.Context, _ string) (net.Conn, *bufio.Reader, Handshake, error) {
	return nil, nil, Handshake{}, nil
}

type Frame struct {
	Header  Header
	Payload []byte
}

type Handshake struct {
	Protocol   string
	Extensions []interface{}
}

type Header struct {
	Fin    bool
	Rsv    byte
	OpCode OpCode
	Masked bool
	Mask   [4]byte
	Length int64
}

func (_ Header) Rsv1() bool {
	return false
}

func (_ Header) Rsv2() bool {
	return false
}

func (_ Header) Rsv3() bool {
	return false
}

func NewTextFrame(_ []byte) Frame {
	return Frame{}
}

type OpCode uint8

func (_ OpCode) IsControl() bool {
	return false
}

func (_ OpCode) IsData() bool {
	return false
}

func (_ OpCode) IsReserved() bool {
	return false
}

func ReadFrame(_ io.Reader) (Frame, error) {
	return Frame{}, nil
}

func WriteFrame(_ io.Writer, _ Frame) error {
	return nil
}