package cli

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var Nil = errors.New("nil")

func _convertCommand(command []byte) [][]byte {
	args := make([][]byte, 0)

	escape := false
	argsIndex := 0

	args = append(args, []byte{})
	for i, l := 0, len(command); i < l; i++ {
		b := command[i]
		if b == CHAR_SPACE && !escape {
			argsIndex++
			args = append(args, []byte{})
			escape = false
		} else {
			if b == '"' || b == '\'' {
				escape = true
			}
			args[argsIndex] = append(args[argsIndex], b)
		}
	}
	return args
}

func Command(bs []byte) []byte {
	c := _convertCommand(bs)

	b := make([]byte, 0, len(bs))
	b = append(b, '*')

	size := len(c)
	b = append(b, []byte(strconv.Itoa(size))...)
	b = append(b, CR, LF)

	for _, i2 := range c {
		b = append(b, String)
		size := len(i2)
		b = append(b, []byte(strconv.Itoa(size))...)
		b = append(b, CR, LF)
		b = append(b, i2...)
		b = append(b, CR, LF)
	}
	return b
}

func Resp(reader io.Reader) (interface{}, error) {
	read := bufio.NewReader(reader)
	r, _ := read.ReadString(LF)
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
		s, _ := reader.ReadString(LF)
		resp, _ := _resp(s, reader)
		r = append(r, resp)
	}
	return r, nil

}

func _resp(head string, reader *bufio.Reader) (interface{}, error) {
	c := head[0]
	switch c {
	case SimpleString:
		return _simpleString(head, reader)
	case Err:
		return _err(head, reader)
	case Number:
		return _number(head, reader)
	case Array:
		return _array(head, reader)
	case String:
		return _string(head, reader)
	}
	return nil, errors.New("unsupport operators")
}
