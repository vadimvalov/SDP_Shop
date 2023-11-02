package main

import (
	"fmt"
	"shop/product"
	"shop/cart"
)

func main() {
	product1 := product.CreateProduct(1, "laptop", 1000.01)
	fmt.Printf("%+v\n", product1)
	cartInstance := cart.GetCartInstance()
	cartInstance.AddToCart(product1)
	cartInstance.AddToCart(product1)
	totalPrice := cartInstance.GetTotalPrice()
	fmt.Printf("Total Price: $%.2f\n", totalPrice)
}