package main

import (
	"fmt"
	"shop/product"
)

func main() {
	product1 := product.CreateProduct(1, "laptop", 1000.01)
	fmt.Printf("%+v\n", product1)
}