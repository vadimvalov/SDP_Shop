package cart

import (
	"shop/product"
	"sync"
)

type Cart struct {
	contents   []product.Product
	totalPrice float64
	observers  []Observer
}

var (
	cartInstance *Cart
	once         sync.Once
)

func GetCartInstance() *Cart {
	once.Do(func() {
		cartInstance = &Cart{}
	})
	return cartInstance
}

func (c *Cart) AddToCart(p product.Product) {
	cartInstance.contents = append(cartInstance.contents, p)
	cartInstance.totalPrice += p.Price
	c.NotifyObservers()
}

func (c *Cart) RemoveFromCart(p product.Product) {
	for i, item := range c.contents {
		if item == p {
			c.contents = append(c.contents[:i], c.contents[i+1:]...)
			c.totalPrice -= p.Price
			c.NotifyObservers()
			break
		}
	}
}

func (c *Cart) GetTotalPrice() float64 {
	return cartInstance.totalPrice
}

func (c *Cart) Attach(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *Cart) Detach(observer Observer) {
	for i, o := range c.observers {
		if o == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *Cart) NotifyObservers() {
	for _, observer := range c.observers {
		observer.Update(c)
	}
}
