package postman

import (
	"bufio"
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"os"
)

var app = models.App{Name: "felix", Remark: "default app for felix postman"}
var env = models.Env{}

func Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
