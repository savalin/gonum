package path

func (p AllShortest) Flush() {
	if p.dist != nil {
		p.dist.Reset()
		p.dist = nil
	}

	p.nodes = nil
	p.indexOf = nil
	p.next = nil
}
