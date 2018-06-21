package matrix

// CompareConcurrently two matrix concurrently
// for each row this method create a goroutine
// TODO Check if this implementation is the best
func (m *Matrix) CompareConcurrently(o *Matrix) bool {
	for x := range m.matrix {
		go func(x int) bool {
			for y := range o.matrix[0] {
				if m.matrix[x][y] != o.matrix[x][y] {
					return false
				}
			}
			return true
		}(x)
	}
	return true
}
