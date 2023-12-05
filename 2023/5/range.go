package main

type Range struct{ src, len int }

func (r Range) end() int { return r.src + r.len }

// intersect checks if a intersects with b.
// Returns the instersected range, as well as all parts that were not matched.
func (a Range) intersect(b Range) (intersected Range, unmatched []Range) {
	start := max(a.src, b.src)
	end := min(a.end(), b.end())
	length := end - start

	// No overlap.
	if length < 0 {
		return Range{}, append(unmatched, a)
	}

	intersected = Range{
		src: start - b.src + b.src,
		len: length,
	}

	// Overflow from bottom.
	if a.src < b.src {
		unmatched = append(unmatched, Range{src: a.src, len: b.src - a.src})
	}
	// Overflow from top.
	if a.end() > b.end() {
		unmatched = append(unmatched, Range{src: b.end(), len: a.end() - b.end()})
	}

	return intersected, unmatched
}
