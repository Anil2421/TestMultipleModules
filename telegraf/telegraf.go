package telegraf

import (
	"fmt"
	"main/service"
)

func MethodTelegraf() {
	fmt.Println("This is the method from the telegraf file")
	service.MethodService()
}
