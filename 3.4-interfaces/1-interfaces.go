package types

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DoInterfaces1() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	v1, ok1 := value1.(float64)
	v2, ok2 := value2.(float64)
	op, ok3 := operation.(string)

	if !ok1 {
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}
	if !ok2 {
		fmt.Printf("value=%v: %T", value2, value2)
		return
	}
	if !ok3 || (op != "+" && op != "-" && op != "*" && op != "/") {
		fmt.Printf("неизвестная операция")
		return
	}

	switch op {
	case "+":
		fmt.Printf("%.4f", v1+v2)
	case "-":
		fmt.Printf("%.4f", v1-v2)
	case "*":
		fmt.Printf("%.4f", v1*v2)
	case "/":
		if v2 == 0 {
			fmt.Printf("ошибка")
			return
		}
		fmt.Printf("%.4f", v1/v2)
	default:
		fmt.Printf("неизвестная операция")
	}
}

func readTask() (interface{}, interface{}, interface{}) {
	var value1, value2, value3 string

	value1, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	value1 = strings.TrimSpace(value1)
	v1, _ := strconv.ParseFloat(value1, 64)

	value2, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	v2, _ := strconv.ParseFloat(value2, 64)

	value3, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	value3 = strings.TrimSpace(value3)

	return v1, v2, value3
}

/*****************************************************************/

type calculator interface {
	add() float64
	subtract() float64
	multi() float64
	div() float64
}

type Expression struct {
	first     float64
	second    float64
	operation string
}

func newExpression(v1, v2, op interface{}) (*Expression, error) {
	value1, err := getFloat(v1)
	if err != nil {
		return nil, fmt.Errorf("value=%v: %T", v1, v1)
	}
	value2, err := getFloat(v2)
	if err != nil {
		return nil, fmt.Errorf("value=%v: %T", v2, v2)
	}
	operation, err := getString(op)
	if err != nil {
		return nil, fmt.Errorf("value=%v: %T", op, op)
	}

	return &Expression{value1, value2, operation}, nil
}

func getFloat(v interface{}) (float64, error) {
	if res, ok := v.(float64); !ok {
		return 0, fmt.Errorf("ошибка преобразования типа")
	} else {
		return res, nil
	}
}

func getString(s interface{}) (string, error) {
	if res, ok := s.(string); !ok {
		return "", fmt.Errorf("ошибка преобразования типа")
	} else {
		return res, nil
	}
}

func (e *Expression) doOperation() float64 {
	switch e.operation {
	case "+":
		return e.add()
	case "-":
		return e.subtract()
	case "*":
		return e.multi()
	case "/":
		res, err := e.div()
		if err != nil {
			panic(err.Error())
		} else {
			return res
		}
	default:
		panic("неизвестная операция")
	}
}

func (e *Expression) add() float64 {
	return e.first + e.second
}

func (e *Expression) subtract() float64 {
	return e.first - e.second
}

func (e *Expression) multi() float64 {
	return e.first * e.second
}

func (e *Expression) div() (float64, error) {
	if e.second == 0 {
		return 0, fmt.Errorf("zero division")

	}
	return e.first / e.second, nil
}

func DoInterface12() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	exp, err := newExpression(value1, value2, operation)

	if err != nil {
		panic(err.Error())
	}

	res := exp.doOperation()
	fmt.Printf("%0.4f\n", res)
}
