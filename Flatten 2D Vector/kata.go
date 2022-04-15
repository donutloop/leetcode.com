package Flatten_2D_Vector

type Vector2D struct {
	vec [][]int
	i   int
	j   int
	end bool
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{vec: vec}
}

func (this *Vector2D) Next() int {
	n := this.vec[this.j][this.i]
	if this.i+1 < len(this.vec[this.j]) {
		this.i = this.i + 1
	} else if this.j+1 < len(this.vec) {
		this.j = this.j + 1
		this.i = 0
	} else {
		this.end = true
	}
	return n
}

func (this *Vector2D) HasNext() bool {
	if this.end {
		return false
	}

	if len(this.vec) > 0 && this.j < len(this.vec) && len(this.vec[this.j]) == 0 {
		this.j++
		return this.HasNext()
	}
	if len(this.vec) > 0 && this.j < len(this.vec) && this.i < len(this.vec[this.j]) {
		return true
	}
	if this.j+1 < len(this.vec) {
		return true
	}
	return false
}
