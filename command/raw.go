package command

import "strconv"

type RawCommand struct {
	raw []byte
}

func (r *RawCommand) Size(size int) *RawCommand {
	r.raw = append(r.raw, []byte(strconv.Itoa(size))...)
	r.raw = append(r.raw, CR, LF)
	return r
}

func (r *RawCommand) String(str []byte) *RawCommand {
	r.raw = append(r.raw, String)
	r.raw = append(r.raw, []byte(strconv.Itoa(len(str)))...)
	r.raw = append(r.raw, CR, LF)
	r.raw = append(r.raw, str...)
	r.raw = append(r.raw, CR, LF)
	return r
}

func (r *RawCommand) Append(bytes []byte) *RawCommand {
	r.raw = append(r.raw, bytes...)
	return r
}

func (r RawCommand) Bytes() []byte {
	return r.raw
}

func NewRaw() *RawCommand {
	b := &RawCommand{raw: make([]byte, 0, 10)}
	b.raw = append(b.raw, '*')
	return b
}

func FromInline(command []byte) *RawCommand {
	r := NewRaw()

	args := make([][]byte, 0)

	escape := false
	argsIndex := 0

	args = append(args, []byte{})
	for i, l := 0, len(command); i < l; i++ {
		b := command[i]
		if b == CharSpace && !escape {
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

	r.Size(len(args))
	for _, b := range args {
		r.String(b)
	}
	return r
}
