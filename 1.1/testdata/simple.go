package testdata

import "fmt"

func Simple() {
	s := "Hello, world!"
	fmt.Printf("%s", s)
}

func Complex(n int) {
	if n > 0 {
		for n > 1 {
			for n > 2 {
				for n > 3 {
					for n > 4 {
						for n > 5 {
							return
						}
					}
				}
			}
		}
	}
	if n > 0 {
		for n > 1 {
			for n > 2 {
				for n > 3 {
					for n > 4 {
						for n > 5 {
							return
						}
					}
				}
			}
		}
	}
}
