package product

type Product struct {
	ID    int
	Name  string
	Price float64
}

func CreateProduct(id int, name string, price float64) Product {
	return Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}