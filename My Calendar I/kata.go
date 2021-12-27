package kata

type BookingInfo struct {
	start int
	end   int
}

type MyCalendar struct {
	booked []BookingInfo
}

func Constructor() MyCalendar {
	return MyCalendar{
		booked: make([]BookingInfo, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.booked) == 0 {
		this.booked = append(this.booked, BookingInfo{start: start, end: end})
		return true
	}

	for _, booking := range this.booked {
		if booking.start < end && start < booking.end {
			return false
		}
	}
	this.booked = append(this.booked, BookingInfo{start: start, end: end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
