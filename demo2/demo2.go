package main

import "bufio"

func main() {
	b := bufio.NewWriter(fd)
	b.Write(p0[a:b])
	b.Write(p1[c:d])
	b.Write(p2[e:f])
	// and so on
	if b.Flush() != nil {
		return b.Flush()
	}
}
