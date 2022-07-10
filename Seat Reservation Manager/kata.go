package Seat_Reservation_Manager

type SeatManager struct {
	seats   []int
	current int
}

func Constructor(n int) SeatManager {
	return SeatManager{seats: make([]int, n)}
}

func (this *SeatManager) Reserve() int {
	if this.seats[this.current] == 1 {
		for this.seats[this.current] != 0 {
			this.current++
		}
	}

	this.seats[this.current] = 1
	seat := this.current + 1
	this.current++
	return seat
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.seats[seatNumber-1] = 0
	if this.current > (seatNumber - 1) {
		this.current = seatNumber - 1
	}
}
