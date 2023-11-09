package main

import (
	"fmt"
	"os"
	"shop/cart"
	"shop/product"
)

func main() {

	fmt.Println("Добро пожаловать в магазин!\n")

	var clientName string
	fmt.Print("Введите Ваше имя: ")
	fmt.Scanln(&clientName)
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
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Открыть магазин")
		fmt.Println("2. Открыть корзину")
		fmt.Println("3. Закрыть")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите цифру от 1 до 3.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("\n###################")
			fmt.Println("Вы открыли магазин.")
			fmt.Println("###################\n")

			fmt.Println("Содержимое магазина:")
			for _, p := range products {
				fmt.Printf("%d. %s - $%.2f\n", p.ID, p.Name, p.Price)
			}
			for {
				fmt.Print("\nВведите номер продукта для добавления в корзину (или 0 для выхода): ")
				var productChoice int
				_, err := fmt.Scanln(&productChoice)
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

				selectedProduct := products[productChoice-1]
				taxDecorator.AddToCart(selectedProduct)
			}

		case 2:
			fmt.Println("\n###################")
			fmt.Println("Вы открыли корзину.")
			fmt.Println("###################")
			for {
				contents := cartInstance.GetContents()
				if len(contents) == 0 {
					fmt.Println("Корзина пуста.\n")
					break
				} else {
					fmt.Println("\nСодержимое корзины:")
					for i, item := range contents {
						fmt.Printf("%d. %s - $%.2f\n", i+1, item.Name, item.Price)
					}
					fmt.Print("\nВведите номер товара, который хотите удалить (или 0 для выхода): ")
					var choice int
					_, err := fmt.Scanln(&choice)
					if err != nil {
						fmt.Println("Ошибка ввода. Пожалуйста, введите цифру.")
						fmt.Scanln()
						continue
					}
					if choice == 0 {
						break
					}
					if choice < 1 || choice > len(contents) {
						fmt.Println("Неверный выбор товара.")
						fmt.Scanln()
						continue
					}
					removedProduct := contents[choice-1]
					cartInstance.RemoveFromCart(removedProduct)
				}
			}

		case 3:
			fmt.Println("\nВы закончили с выбором товаров.")
			fmt.Printf("Общая сумма к оплате: $%.2f \n", taxDecorator.GetTotalPrice())
			cartInstance.Detach(observer1)
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите от 1 до 3.")
		}
	}
}
