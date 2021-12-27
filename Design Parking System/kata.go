package kata

type ParkingSystem struct {
	big, medium, small int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

const (
	_ = iota
	CarTypeBig
	CarTypeMedium
	CarTypeSmall
)

func (ps *ParkingSystem) AddCar(carType int) bool {
	var found bool
	if carType == CarTypeSmall {
		if ps.small > 0 {
			found = true
			ps.small--
		}
	} else if carType == CarTypeMedium {
		if ps.medium > 0 {
			found = true
			ps.medium--
		}
	} else if carType == CarTypeBig {
		if ps.big > 0 {
			found = true
			ps.big--
		}
	}

	return found
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
