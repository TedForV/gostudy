package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient of the pizza decoreator it must return"+
			" the text %s the expected text, not '%s'", pizzaResult, expectedText)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	if err == nil {
		t.Errorf("%s,%s", onionResult, err.Error())
	}

	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(onionResult, "onion") {
		t.Errorf("%s,%s", onionResult, err.Error())
	}
}
