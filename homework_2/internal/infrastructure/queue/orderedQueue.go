package queue

import (
	"code-cadets-2021/homework_2/internal/domain/models"
	"context"
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type OrderedQueue struct {
	queue  []models.Odd
	source chan models.Odd
}

func NewOrderedQueue() *OrderedQueue {
	return &OrderedQueue{
		source: make(chan models.Odd),
	}
}

func (o *OrderedQueue) String() string {
	return "ordered queue"
}

func (o *OrderedQueue) Start(ctx context.Context) error {
	// ignore ctx parameter, we will use it later :)

	// initially:
	// - load existing data from disk
	//
	// repeatedly:
	// - read source channel
	// - update queue slice
	// - when source channel is closed, exit
	//
	// finally:
	// - store queue slice to disk
	//fmt.Println("uso u queue")
	err := o.loadFromFile()
	if err != nil {
		//fmt.Println("izlazim iz queue", err)
		return err
	}

	//fmt.Println("rangeam po srcu")
	for odd := range o.source {
		//fmt.Println("uso u odd")
		o.queue = append(o.queue, odd)
	}

	err = o.storeToFile()
	if err != nil {
		//fmt.Println("storing file error", err)
		return err
	}

	return nil
}

func (o *OrderedQueue) GetSource() chan models.Odd {
	return o.source
}

func (o *OrderedQueue) loadFromFile() error {
	f, err := os.Open("queue.txt")
	if os.IsNotExist(err) {
		return nil

	} else if err != nil {
		return errors.Wrap(err, "load from file, open")
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&o.queue)
	if err != nil {
		return errors.Wrap(err, "load from file, decode")
	}

	return nil
}

func (o *OrderedQueue) storeToFile() error {
	f, err := os.Create("queue.txt")
	if err != nil {
		return errors.Wrap(err, "store to file, create")
	}
	defer f.Close()

	serializedQueue, err := json.MarshalIndent(o.queue, "", "    ")
	if err != nil {
		return errors.Wrap(err, "store to file, marshal")
	}

	n, err := f.Write(serializedQueue)
	if err != nil {
		return errors.Wrap(err, "store to file, write")

	} else if len(serializedQueue) != n {
		return errors.Wrapf(err, "store to file, write len; n: %d, serializedLen: %d", n, len(serializedQueue))
	}

	err = f.Sync()
	if err != nil {
		return errors.Wrap(err, "store to file, sync")
	}

	return nil
}
