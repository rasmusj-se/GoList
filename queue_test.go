package GoList

import "testing"

func TestQueue(t *testing.T) {
	queue := &Queue{}
	queue.Enqueue("Hello")
	if val := queue.Dequeue(); val != "Hello" {
		t.Error("Queue did not return expected value")
	}
	queue.Enqueue("Test")
	if val := queue.Dequeue(); val != "Test" {
		t.Error("Queue did not return expected value")
	}
	for i := 0; i < 512; i++ {
		queue.Enqueue(i)
	}
	for i := 0; i < 512; i++ {
		if queue.Dequeue() != i || queue.Count() != 511-i {
			t.Error("Queue did not handle large size correctly")
		}
	}
}
