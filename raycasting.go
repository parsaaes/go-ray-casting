package raycasting

type Polygon struct {
	Points [][]float64
}

func (p *Polygon) Contains(point [2]float64) bool {
	result := false
	x, y := point[0], point[1]

	for i, j := 0, len(p.Points)-1; i < len(p.Points); j, i = i, i+1 {
		firstX, firstY := p.Points[i][0], p.Points[i][1]
		secondX, secondY := p.Points[j][0], p.Points[j][1]

		if (firstY > y) != (secondY > y) &&
			x < (y-firstY)*(secondX-firstX)/(secondY-firstY)+firstX {
			result = !result
		}

	}

	return result
}
