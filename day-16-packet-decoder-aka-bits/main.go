package main

import (
	"fmt"
	"log"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	var input bits
	scanner := func(line string) error {
		input = bits{data: hexStringToBits(line)}
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	fmt.Println(input)

	result1 := puzzle1(input)
	result2 := puzzle2(input)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(in bits) string {
	verSum := 0
	callback := func(p packet) {
		logger(p)
		verSum += p.Header().ver
	}
	in.callback = callback
	in.read()

	return fmt.Sprintf("%d", verSum)
}

func puzzle2(in bits) string {
	in.callback = logger
	p := in.read()

	return fmt.Sprintf("%d", p.Value())
}

func logger(p packet) {
	fmt.Printf("  read: %+v\n", p)
}

type bits struct {
	data     []byte
	i        int
	callback func(packet)
}

func (b *bits) empty() bool {
	return b.i == len(b.data)
}

// read reads next 'n' bits and returns the  corresponding integer
func (b *bits) readInt(n int) int {
	if b.i+n > len(b.data) {
		log.Fatalf("cannot read %d bytes, only %d are left", n, len(b.data)-b.i)
	}

	result := 0
	for i := 0; i < n; i++ {
		result *= 2
		result += int(b.data[b.i])
		b.i++
	}

	return result
}

func (b *bits) readUntil(reader func(byte) bool) {
	if b.i >= len(b.data)-1 {
		log.Fatalf("cannot read next bit, data exhausted: i=%d, n=%d", b.i, len(b.data))
	}

	for ; reader(b.data[b.i]); b.i++ {
	}
	b.i++
}

func (b *bits) readHeader() header {
	ver := b.readInt(3)
	typ := b.readInt(3)
	return header{ver: ver, typ: Typ(typ)}
}

func (b *bits) readLiteral() int {
	lastChunk := true
	result := 0
	i := 0
	b.readUntil(func(d byte) bool {
		defer func() { fmt.Println("    --> readLiteral", i, ":", d, "->", lastChunk) }()
		i++
		if i%5 == 1 {
			lastChunk = d == 0
			return true
		}
		result *= 2
		result += int(d)
		if i%5 == 0 && lastChunk {
			return false
		}
		return true
	})
	return result
}

func (b *bits) readLen0Operator() []packet {
	n := b.readInt(15)
	fmt.Println("    --> len0 operator, n:", n)
	subPackets := []packet{}
	for {
		initialI := b.i
		sp := b.read()
		subPackets = append(subPackets, sp)
		n -= b.i - initialI
		fmt.Println("    --> len0 operator, remaining bits n:", n)
		if n <= 0 {
			break
		}
	}
	return subPackets
}

func (b *bits) readLen1Operator() []packet {
	n := b.readInt(11)
	fmt.Println("    --> len1 operator, n:", n)
	subPackets := []packet{}
	for ; n > 0; n-- {
		sp := b.read()
		subPackets = append(subPackets, sp)
		fmt.Println("    --> len1 operator, remaining  packets n:", n)
		if n <= 0 {
			break
		}
	}
	return subPackets
}

func (b *bits) read() packet {
	// if b.empty() {
	// 	return
	// }
	h := b.readHeader()
	fmt.Println("  --> header:", h)
	if isLiteral(h.typ) {
		i := b.readLiteral()
		l := literal{header: h, l: i}
		b.callCallback(l)
		return l
	}
	lenTyp := b.readInt(1)
	subPackets := []packet{}
	if lenTyp == 0 {
		subPackets = b.readLen0Operator()
	} else {
		subPackets = b.readLen1Operator()
	}
	o := operator{header: h, subPackets: subPackets}
	b.callCallback(o)
	return o
}

func (b *bits) callCallback(p packet) {
	if b.callback != nil {
		b.callback(p)
	}
}

func isLiteral(typ Typ) bool {
	return typ == Literal
}

type Typ int

const (
	Sum Typ = iota
	Prod
	Min
	Max
	Literal
	Greater
	Less
	Equal
)

type header struct {
	ver int
	typ Typ
}

type packet interface {
	Header() header
	Value() int
}

type literal struct {
	header
	l int
}

func (l literal) Header() header {
	return l.header
}

func (l literal) Value() int {
	return l.l
}

type operator struct {
	header
	subPackets []packet
}

func (o operator) Header() header {
	return o.header
}

func (o operator) Value() int {
	switch o.typ {
	case Sum:
		return o.sum()
	case Prod:
		return o.prod()
	case Min:
		return o.min()
	case Max:
		return o.max()
	case Greater:
		return o.greater()
	case Less:
		return o.less()
	case Equal:
		return o.equal()
	}
	log.Fatalf("Unknown type: %d", o.typ)
	return -1
}

func (o operator) sum() int {
	result := 0
	for _, p := range o.subPackets {
		result += p.Value()
	}
	return result
}

func (o operator) prod() int {
	result := 1
	for _, p := range o.subPackets {
		result *= p.Value()
	}
	return result
}

func (o operator) min() int {
	result := -1
	for _, p := range o.subPackets {
		v := p.Value()
		if result == -1 || v < result {
			result = v
		}
	}
	return result
}

func (o operator) max() int {
	result := -1
	for _, p := range o.subPackets {
		result = max(result, p.Value())
	}
	return result
}

func (o operator) greater() int {
	o.assertTwoSubpackets()
	if o.subPackets[0].Value() > o.subPackets[1].Value() {
		return 1
	}
	return 0
}

func (o operator) less() int {
	o.assertTwoSubpackets()
	if o.subPackets[0].Value() < o.subPackets[1].Value() {
		return 1
	}
	return 0
}

func (o operator) equal() int {
	o.assertTwoSubpackets()
	if o.subPackets[0].Value() == o.subPackets[1].Value() {
		return 1
	}
	return 0
}

func (o operator) assertTwoSubpackets() {
	if len(o.subPackets) != 2 {
		log.Fatalf("expected exactly 2 sub-packets, but got: %d", len(o.subPackets))
	}
}
