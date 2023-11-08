package main

import (
	"fmt"
	"os"
	"shop/cart"
	"shop/product"
)

func main() {
	//cartInstance := cart.GetCartInstance()
	//
	//observer1 := cart.CartObserver{Name: "Dias"}
	//observer2 := cart.CartObserver{Name: "Vadim"}
	//cartInstance.Attach(observer1)
	//cartInstance.Attach(observer2)
	//
	//product1 := product.CreateProduct(1, "laptop", 1000.01)
	//product2 := product.CreateProduct(2, "smartphone", 850.25)
	//product3 := product.CreateProduct(3, "Motorcycle", 7250.50)
	//fmt.Printf("%+v\n", product1)
	//
	//cartInstance.AddToCart(product1)
	//cartInstance.AddToCart(product2)
	//fmt.Printf("\n")
	//
	//cartInstance.RemoveFromCart(product2)
	//fmt.Printf("\n")
	//
	//cartInstance.AddToCart(product3)
	//fmt.Printf("\n")
	//
	//fmt.Println("--------------------------")
	//
	//totalPrice := cartInstance.GetTotalPrice()
	//fmt.Printf("Total Price: $%.2f\n\n\n", totalPrice)

	fmt.Println("Добро пожаловать в магазин!")

	var clientName string
	fmt.Print("Введите Ваше имя: ")
	fmt.Fscan(os.Stdin, &clientName)
	observer1 := cart.CartObserver{Name: clientName}

	cartInstance := cart.GetCartInstance()
	cartInstance.Attach(observer1)

	products := []product.Product{
		{1, "Лаптоп", 1000.01},
		{2, "Смартфон", 799.99},
		{3, "Наушники", 99.50},
	}

	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1. Открыть магазин")
		fmt.Println("2. Открыть корзину")
		fmt.Println("3. Закрыть")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите цифру от 1 до 3.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Вы открыли магазин.")

			for _, p := range products {
				fmt.Printf("%d. %s - $%.2f\n", p.ID, p.Name, p.Price)
			}
			for {
				fmt.Print("Введите номер продукта для добавления в корзину (или 0 для выхода): ")
				var productChoice int
				_, err := fmt.Scanf("%d", &productChoice)
				if err != nil {
					fmt.Println("Ошибка ввода. Пожалуйста, введите цифру.")
					continue
				}
				if productChoice == 0 {
					break
				}
				if productChoice < 1 || productChoice > len(products) {
					fmt.Println("Неверный выбор продукта.")
					continue
				}

				// Добавление выбранного продукта в корзину
				selectedProduct := products[productChoice-1]
				cartInstance.AddToCart(selectedProduct)
			}

		case 2:
			fmt.Println("Вы открыли корзину.")

		case 3:
			fmt.Println("Вы закрыли программу.")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите от 1 до 3.")
		}
	}
}
