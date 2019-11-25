package path

func (p AllShortest) Flush() {
	if p.dist != nil {
		p.dist.Reset()
		p.dist = nil
	}

	p.nodes = nil
	for k := range p.indexOf {
		delete(p.indexOf, k)
	}
	p.indexOf = nil
	p.next = nil
}
