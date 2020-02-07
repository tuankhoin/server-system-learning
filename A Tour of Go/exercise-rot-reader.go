package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error){
	n, err := r.r.Read(b)
	for i:= range b {
		if (b[i]>='A' && b[i]<'N') || (b[i]>='a' && b[i]<'n') {
			b[i]+=13
		} else if (b[i]>='N' && b[i]<'Z') || (b[i]>='n' && b[i]<'z') {
			b[i]-=13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
