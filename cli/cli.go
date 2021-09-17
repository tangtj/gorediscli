package cli

import (
	"bufio"
	"errors"
	"github.com/tangtj/gorediscli/command"
	"io"
	"strconv"
	"strings"
)

var Nil = errors.New("nil")

func Command(bs []byte) []byte {
	c := command.FromInline(bs)
	return c.Bytes()
}

func Resp(reader io.Reader) (interface{}, error) {
	read := bufio.NewReader(reader)
	r, _ := read.ReadString(command.LF)
	result, _ := _resp(r, read)
	return result, nil
}

func _simpleString(head string, reader *bufio.Reader) (string, error) {
	strLen := len(head) - 2
	return head[1:strLen], nil
}

func _err(head string, reader *bufio.Reader) (string, error) {
	strLen := len(head) - 2
	return head[1:strLen], nil
}

func _string(head string, reader *bufio.Reader) (string, error) {
	s := strings.TrimSpace(head[1:])
	size, _ := strconv.Atoi(s)
	if size <= -1 {
		return "", Nil
	}
	b := make([]byte, size+2)
	reader.Read(b)
	return string(b[:size]), nil
}

func _number(head string, reader *bufio.Reader) (int, error) {
	s := strings.TrimSpace(head[1:])
	size, _ := strconv.Atoi(s)
	return size, nil
}

func _array(head string, reader *bufio.Reader) ([]interface{}, error) {
	s := strings.TrimSpace(head[1:])
	size, _ := strconv.Atoi(s)
	if size <= -1 {
		return nil, Nil
	}
	r := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		s, _ := reader.ReadString(command.LF)
		resp, _ := _resp(s, reader)
		r = append(r, resp)
	}
	return r, nil

}

func _resp(head string, reader *bufio.Reader) (interface{}, error) {
	c := head[0]
	switch c {
	case command.SimpleString:
		return _simpleString(head, reader)
	case command.Err:
		return _err(head, reader)
	case command.Number:
		return _number(head, reader)
	case command.Array:
		return _array(head, reader)
	case command.String:
		return _string(head, reader)
	}
	return nil, errors.New("unsupport operators")
}
