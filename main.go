package main
import (
	"github.com/api/app"
	"github.com/api/config"
)
func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
