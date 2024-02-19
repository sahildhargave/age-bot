package main 

import (
	"fmt"
	"os"
	"log"
	"context"
	"github.com/shomali11/slacker"
)

func main(){
	fmt.Println("Welcome to age-bot");
	os.Setenv("SLACK_BOT_TOKEN","xoxb-6661129647267-6663744600948-EjAeTQMJF67UILDDeMpdMAds")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A06KTT31HKK-6658347548053-691d23abcbf4854f1713771a65c4e3011aaa8d2b71173fa90e513027cfa1e360");

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"))
}
