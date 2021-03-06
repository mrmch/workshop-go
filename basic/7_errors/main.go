package main

import (
    "fmt"
    "errors"
)

type Computer struct {
    Brand string
    Model string
    Price int
}

func (c *Computer) Describe() {
    fmt.Printf("%s %s $%d\n", c.Brand, c.Model, c.Price)
}

func (c *Computer) Purchase(dollars int) (error, int) {
    change := dollars - c.Price
    
    if change < 0 {
        return errors.New("Not enough money"), 0
    } else {
        return nil, change
    }
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
    computer.Describe()

    err, change := computer.Purchase(1500)
    
    if err != nil {
        fmt.Println("Failed to purchase")
    } else {
        fmt.Printf("Purchase returned $%d\n", change)
    }
}