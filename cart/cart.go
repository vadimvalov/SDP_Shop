package cart

import (
    "sync"
    "shop/product"
)

type Cart struct {
    contents   []product.Product
    totalPrice float64
}

var (
    cartInstance *Cart
    once sync.Once
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
}

func (c *Cart) GetTotalPrice() float64 {
    return cartInstance.totalPrice
}
