package hostname

import (
	"fmt"
	"os"
)

func GetHostname()Result{
	return Get()
}
func  Get() Result {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("Host name request servers from : " + name)
	return Result{Hostname: name}
}