package model

import (
	. "github.com/onsi/gomega"
	"sync"
	"testing"
	"time"
)

type testResource struct {
	value string
}

func (r *testResource) clone() cloneable {
	return &testResource{
		value: r.value,
	}
}

func TestModificationHandler(t *testing.T) {
	RegisterTestingT(t)

	bd := baseDomain{}
	resource := &testResource{"test"}
	updResource := &testResource{"updated"}

	amountHandlers := 5

	var wg sync.WaitGroup
	for i := 0; i < amountHandlers; i++ {
		wg.Add(3)
		bd.addHandler(&ModificationHandler{
			AddFunc: func(new interface{}) {
				defer wg.Done()
				Expect(new.(*testResource).value).To(Equal(resource.value))
			},
			UpdateFunc: func(old interface{}, new interface{}) {
				defer wg.Done()
				Expect(old.(*testResource).value).To(Equal(resource.value))
				Expect(new.(*testResource).value).To(Equal(updResource.value))
			},
			DeleteFunc: func(del interface{}) {
				defer wg.Done()
				Expect(del.(*testResource).value).To(Equal(resource.value))
			},
		})
	}

	bd.resourceAdded(resource)
	bd.resourceUpdated(resource, updResource)
	bd.resourceDeleted(resource)
	doneCh := make(chan struct{})

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <-time.After(5 * time.Second):
		t.Fatal("not all listeners have been emitted ")
	case <-doneCh:
		return
	}
}
