package main

import (
	"errors"
	"fmt"
	"strings"
)

type CarPart struct {
	engine   bool
	brakes   bool
	lights   bool
	oilLevel float64
}

func (c *CarPart) CheckEngine() error {
	if !c.engine {
		return errors.New("двигатель не запускается")
	}
	return nil
}

func (c *CarPart) CheckBrakes() error {
	if !c.brakes {
		return errors.New("неисправны тормоза")
	}
	return nil
}

func (c *CarPart) CheckLights() error {
	if !c.lights {
		return errors.New("не работает освещение")
	}
	return nil
}

func (c *CarPart) CheckOilLevel() error {
	if c.oilLevel < 1.0 {
		return fmt.Errorf("низкий уровень масла (%.2f)", c.oilLevel)
	}
	return nil
}

func (c *CarPart) CheckCar() error {
	var errs []error

	if err := c.CheckEngine(); err != nil {
		errs = append(errs, err)
	}
	if err := c.CheckBrakes(); err != nil {
		errs = append(errs, err)
	}
	if err := c.CheckLights(); err != nil {
		errs = append(errs, err)
	}
	if err := c.CheckOilLevel(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func main() {
	var n int
	fmt.Scanln(&n)

	cars := make([]CarPart, n)

	for i := 0; i < n; i++ {
		var part CarPart
		fmt.Scanln(&part.engine, &part.brakes, &part.lights, &part.oilLevel)
		cars[i] = part
	}

	for _, car := range cars {
		err := car.CheckCar()
		if err != nil {
			fmt.Println("Проблемы:")
			parts := strings.Split(err.Error(), "\n")
			for _, part := range parts {
				if part != "" {
					fmt.Println("-", part)
				}
			}
		} else {
			fmt.Println("Автомобиль OK")
		}
	}
}
