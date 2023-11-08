package main

import (
	"fmt"
	"os"
	"shop/cart"
	"shop/product"
)

func main() {

	fmt.Println("Добро пожаловать в магазин!")

	var clientName string
	fmt.Print("Введите Ваше имя: ")
	fmt.Fscan(os.Stdin, &clientName)
	observer1 := cart.CartObserver{Name: clientName}

	cartInstance := cart.GetCartInstance()
	taxDecorator := cart.NewTaxDecorator(cartInstance, 0.12)
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
				taxDecorator.AddToCart(selectedProduct)
			}

		case 2:
			fmt.Println("Вы открыли корзину.")

		case 3:
			fmt.Println("Вы закрыли программу.")
			fmt.Printf("Общая сумма к оплате: %f", taxDecorator.GetTotalPrice())
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите от 1 до 3.")
		}
	}
}
