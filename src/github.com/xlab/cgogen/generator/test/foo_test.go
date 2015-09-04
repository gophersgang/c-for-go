package foo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassInt(t *testing.T) {
	expected := 5
	result := PassInt(2, 3)
	assert.Equal(t, expected, result)
}

func TestPassString(t *testing.T) {
	expected := "ab"
	result := PassString("a", "b")
	assert.Equal(t, expected, result)
}

func TestPassBytes(t *testing.T) {
	expected := []byte("abcd")
	result := PassBytes([]byte("ab"), 2, []byte("cd"), 2)
	assert.Equal(t, expected, result[:4])
}

func TestFindChar(t *testing.T) {
	tbl := []struct {
		Input  string
		Search byte
		Index  int64
	}{
		{"golang", 'c', -1},
		{"gopher", 'g', 0},
		{"fun", 'n', 2},
	}
	for _, test := range tbl {
		result := FindChar(test.Input, test.Search)
		assert.Equal(t, test.Index, result)
	}
}

func TestSendMessage(t *testing.T) {
	buf := make([]byte, 4096) // 4kB
	msg := &Message{
		FromID:  &[8]byte{0x1},
		ToID:    &[8]byte{0x2},
		Message: "Hey there! Check out these cool pictures attached. -xoxo",
		// AttachmentsLen: 2,
		// Attachments:    &Attachment{},
	}
	size := SendMessage(msg, buf)
	assert.EqualValues(t, 76, size)
	packed := append([]byte("msg:"), 0x01, 0x02)
	packed = append(packed, "Hey there! Check out these cool pictures attached. -xoxo"...)

	assert.Equal(t, packed, cleanBuf(buf))
}

func cleanBuf(buf []byte) []byte {
	tmp := new(bytes.Buffer)
	for i := range buf {
		if buf[i] != 0 {
			tmp.WriteByte(buf[i])
			continue
		}
	}
	return tmp.Bytes()
}

func TestA4Byte(t *testing.T) {
	a := [4]byte{'a', 'a', 'a', 'a'}
	b := [4]byte{'b', 'b', 'b', 'b'}
	A4Byte(&a)
	assert.Equal(t, b, a)
}

func TestA4String(t *testing.T) {
	a := [4]string{"g", "g", "g", "g"}
	b := [4]string{"go", "go", "go", "go"}
	A4String(&a)
	assert.Equal(t, b, a)
}

func TestA4SByte(t *testing.T) {
	a := [4][]byte{
		{'a', 'a'},
		{'a', 'a'},
		{'a', 'a'},
		{'a', 'a'},
	}
	b := [4][]byte{
		{'b', 'b'},
		{'b', 'b'},
		{'b', 'b'},
		{'b', 'b'},
	}
	A4SByte(&a, 2)
	assert.Equal(t, b, a)
}

func TestA4SString(t *testing.T) {
	a := [4][]string{
		{"g", "g"},
		{"g", "g"},
		{"g", "g"},
		{"g", "g"},
	}
	b := [4][]string{
		{"go", "go"},
		{"go", "go"},
		{"go", "go"},
		{"go", "go"},
	}
	A4SString(&a, 2)
	assert.Equal(t, b, a)
}

func TestA2A2Byte(t *testing.T) {
	a := [2][2]byte{
		{'a', 'a'},
		{'a', 'a'},
	}
	b := [2][2]byte{
		{'b', 'b'},
		{'b', 'b'},
	}
	A2A2Byte(&a)
	assert.Equal(t, b, a)
}

func TestA2A2String(t *testing.T) {
	a := [2][2]string{
		{"g", "g"},
		{"g", "g"},
	}
	b := [2][2]string{
		{"go", "go"},
		{"go", "go"},
	}
	A2A2String(&a)
	assert.Equal(t, b, a)
}

func TestA2A2SByte(t *testing.T) {
	a := [2][2][]byte{
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
	}
	b := [2][2][]byte{
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
	}
	A2A2SByte(&a, 3)
	assert.Equal(t, b, a)
}

func TestA2A2SString(t *testing.T) {
	a := [2][2][]string{
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
	}
	b := [2][2][]string{
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
	}
	A2A2SString(&a, 3)
	assert.Equal(t, b, a)
}

func TestSSByte(t *testing.T) {
	a := [][]byte{
		{'a', 'a'},
		{'a', 'a'},
		{'a', 'a'},
		{'a', 'a'},
	}
	b := [][]byte{
		{'b', 'b'},
		{'b', 'b'},
		{'b', 'b'},
		{'b', 'b'},
	}
	SSByte(a, 4, 2)
	assert.Equal(t, b, a)
}

func TestSSString(t *testing.T) {
	a := [][]string{
		{"g", "g"},
		{"g", "g"},
		{"g", "g"},
		{"g", "g"},
	}
	b := [][]string{
		{"go", "go"},
		{"go", "go"},
		{"go", "go"},
		{"go", "go"},
	}
	SSString(a, 4, 2)
	assert.Equal(t, b, a)
}

func TestA4SSByte(t *testing.T) {
	a := [4][][]byte{
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
	}
	b := [4][][]byte{
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
	}
	A4SSByte(&a, 2, 3)
	assert.Equal(t, b, a)
}

func TestA4SSString(t *testing.T) {
	a := [4][][]string{
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
	}
	b := [4][][]string{
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
	}
	A4SSString(&a, 2, 3)
	assert.Equal(t, b, a)
}

func TestA2A2SSByte(t *testing.T) {
	a := [2][2][][]byte{
		{
			{{'a', 'a'}, {'a', 'a'}, {'a', 'a'}},
			{{'a', 'a'}, {'a', 'a'}, {'a', 'a'}},
		},
		{
			{{'a', 'a'}, {'a', 'a'}, {'a', 'a'}},
			{{'a', 'a'}, {'a', 'a'}, {'a', 'a'}},
		},
	}
	b := [2][2][][]byte{
		{
			{{'b', 'b'}, {'b', 'b'}, {'b', 'b'}},
			{{'b', 'b'}, {'b', 'b'}, {'b', 'b'}},
		},
		{
			{{'b', 'b'}, {'b', 'b'}, {'b', 'b'}},
			{{'b', 'b'}, {'b', 'b'}, {'b', 'b'}},
		},
	}
	A2A2SSByte(&a, 3, 2)
	assert.Equal(t, b, a)
}

func TestA2A2SSString(t *testing.T) {
	a := [2][2][][]string{
		{
			{{"g", "g"}, {"g", "g"}, {"g", "g"}},
			{{"g", "g"}, {"g", "g"}, {"g", "g"}},
		},
		{
			{{"g", "g"}, {"g", "g"}, {"g", "g"}},
			{{"g", "g"}, {"g", "g"}, {"g", "g"}},
		},
	}
	b := [2][2][][]string{
		{
			{{"go", "go"}, {"go", "go"}, {"go", "go"}},
			{{"go", "go"}, {"go", "go"}, {"go", "go"}},
		},
		{
			{{"go", "go"}, {"go", "go"}, {"go", "go"}},
			{{"go", "go"}, {"go", "go"}, {"go", "go"}},
		},
	}
	A2A2SSString(&a, 3, 2)
	assert.Equal(t, b, a)
}
