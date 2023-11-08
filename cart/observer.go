package cart

import (
	"fmt"
)

type Observer interface {
	Update(*Cart)
}

type CartObserver struct {
	Name string
}

func (co CartObserver) Update(cart *Cart) {
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Покупатель %s: Корзина была обновлена. Общая цена (без налога): $%.2f\n", co.Name, cart.totalPrice)
	fmt.Println("------------------------------------------------------")
}
