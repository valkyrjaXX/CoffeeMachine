package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	m := NewCoffeeMachine(550, 400, 540, 120, []int{2, 4, 9})

	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scan(&action)

		switch action {
		case "buy":
			for {
				fmt.Println(fmt.Sprintf("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:"))
				var action string
				fmt.Scan(&action)

				if action == "back" {
					break
				}

				beverage, err := strconv.Atoi(action)
				if err != nil {
					fmt.Println(err)
					break
				}

				fmt.Println(fmt.Sprintf("Seleect volume: 1 -  small, 2  - medium, 3 - large"))
				var volume int
				fmt.Scan(&volume)

				err = m.Brew(beverage, volume)
				if err != nil {
					fmt.Println(err.Error())
					break
				}

				fmt.Println("I have enough resources, making you a coffee!")
				break
			}
		case "fill":
			fmt.Println(fmt.Sprintf("Write how many ml of water you want to add:"))
			var water int
			fmt.Scan(&water)

			fmt.Println(fmt.Sprintf("Write how many ml of milk you want to add:"))
			var milk int
			fmt.Scan(&milk)

			fmt.Println(fmt.Sprintf("Write how many grams of coffee beans you want to add:"))
			var coffee int
			fmt.Scan(&coffee)

			fmt.Println(fmt.Sprintf("Write how many disposable cups you want to add:"))
			cups := make([]int, 3)
			fmt.Println(fmt.Sprintf("Small:"))
			fmt.Scan(&cups[0])
			fmt.Println(fmt.Sprintf("Medium:"))
			fmt.Scan(&cups[1])
			fmt.Println(fmt.Sprintf("Large:"))
			fmt.Scan(&cups[2])

			m.Fill(water, milk, coffee, cups)
		case "take":
			fmt.Println(fmt.Sprintf("I gave you $%.0f", m.WithdrawMoney()))
		case "remaining":
			m.PrintStats()
		case "exit":
			return
		}
	}
}

type CoffeeRecipe struct {
	Water  int
	Milk   int
	Coffee int
	Price  float32
}

func (r CoffeeRecipe) CalculateForCups(numOfCups int) (int, int, int) {
	return r.Water * numOfCups, r.Milk * numOfCups, r.Coffee * numOfCups
}

type State struct {
	Money  float32
	Water  int
	Milk   int
	Coffee int
	Cups   []int
}

type CoffeeMachine struct {
	recipes []CoffeeRecipe
	state   *State
}

func (m *CoffeeMachine) Brew(beverage, volume int) error {
	recipe := m.recipes[beverage-1]
	cups := m.state.Cups[volume-1]

	if m.state.Water < recipe.Water {
		return errors.New("Sorry, not enough water!")
	}

	if m.state.Milk < recipe.Milk {
		return errors.New("Sorry, not enough milk!")
	}

	if m.state.Coffee < recipe.Coffee {
		return errors.New("Sorry, not enough coffee beeans!")
	}

	if cups < 1 {
		return errors.New("Sorry, not enough cups!")
	}

	m.state.Water -= recipe.Water
	m.state.Milk -= recipe.Milk
	m.state.Coffee -= recipe.Coffee
	cups--
	m.state.Money += recipe.Price
	return nil
}

func (m *CoffeeMachine) Fill(water, milk, coffee int, cups []int) {
	m.state.Water += water
	m.state.Milk += milk
	m.state.Coffee += coffee
	for idx, amount := range cups {
		m.state.Cups[idx] += amount
	}
}

func (m *CoffeeMachine) WithdrawMoney() float32 {
	amount := m.state.Money
	m.state.Money = 0
	return amount
}

func (m *CoffeeMachine) PrintStats() {
	fmt.Println("The coffee machine has:")
	fmt.Println(fmt.Sprintf("%d ml of water", m.state.Water))
	fmt.Println(fmt.Sprintf("%d ml of milk", m.state.Milk))
	fmt.Println(fmt.Sprintf("%d g of coffee beans", m.state.Coffee))
	fmt.Println(fmt.Sprintf("%d disposable cups", m.state.Cups))
	fmt.Println(fmt.Sprintf("%.0f of money", m.state.Money))
}

func NewCoffeeMachine(money float32, water, milk, coffeeBeans int, cups []int) *CoffeeMachine {
	return &CoffeeMachine{
		recipes: []CoffeeRecipe{
			{
				Water:  250,
				Coffee: 16,
				Price:  4,
			},
			{
				Water:  350,
				Milk:   75,
				Coffee: 20,
				Price:  7,
			},
			{
				Water:  200,
				Milk:   100,
				Coffee: 12,
				Price:  6,
			},
		},
		state: &State{
			Money:  money,
			Water:  water,
			Milk:   milk,
			Coffee: coffeeBeans,
			Cups:   cups,
		},
	}
}
