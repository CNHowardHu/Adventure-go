package IO

import (
	"fmt"
	"os"
)

func PressToExit() {
	fmt.Println("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
