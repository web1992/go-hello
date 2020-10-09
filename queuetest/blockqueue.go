package queuetest

var maxSize = 10
var ch = make(chan interface{}, maxSize)

type BlockQueue struct {
}

func (q *BlockQueue) Offer(obj interface{}) {

	ch <- obj

}

func (q *BlockQueue) Take() interface{} {

	obj := <-ch

	return obj

}
