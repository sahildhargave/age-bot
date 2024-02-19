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
	os.Setenv("SLACK_APP_TOKEN","")
}
