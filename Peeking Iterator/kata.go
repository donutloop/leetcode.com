package kata

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *      // Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *      // Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: iter,
	}
}

func (pt *PeekingIterator) hasNext() bool {
	return pt.iter.hasNext()
}

func (pt *PeekingIterator) next() int {
	return pt.iter.next()
}

func (pt *PeekingIterator) peek() int {
	n := pt.iter.next()
	pt.iter.index--
	return n
}
