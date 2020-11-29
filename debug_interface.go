package main

import (
	. "aicup2020/model"
	"bufio"
)

type DebugInterface struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (debugInterface DebugInterface) Send(command DebugCommand) {
	ClientMessageDebugMessage{
		Command: command,
	}.Write(debugInterface.Writer)
	err := debugInterface.Writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (debugInterface DebugInterface) GetState() DebugState {
	ClientMessageRequestDebugState{}.Write(debugInterface.Writer)
	err := debugInterface.Writer.Flush()
	if err != nil {
		panic(err)
	}
	return ReadDebugState(debugInterface.Reader)
}
