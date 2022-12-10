package pkg

import "fmt"

func GetHello(name string) string {
	return fmt.Sprintf("Hello: %s", name)
}
