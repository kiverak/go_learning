package types

import (
	"fmt"
	"strconv"
)

func DoAnonym1() {
	fn := func() func(uint) uint {
		rsl := make([]rune, 0)
		return func(x uint) uint {
			runes := []rune(strconv.Itoa(int(x)))
			for i := range runes {
				n, _ := strconv.Atoi(string(runes[i]))
				if n%2 == 0 && n != 0 {
					rsl = append(rsl, runes[i])
				}
			}
			str := string(rsl)
			num, _ := strconv.Atoi(str)

			if len(str) == 0 {
				return 100
			}

			return uint(num)
		}
	}()

	fmt.Println(strconv.Itoa(int(fn(123))))
	fmt.Println(strconv.Itoa(int(fn(456))))
	fmt.Println(strconv.Itoa(int(fn(789))))
}
