package main

import (
  "log"
  "os"
	"errors"

  "github.com/urfave/cli"

	"github.com/misrab/alfred/dev"
)

func main() {
  app := cli.NewApp()

	app.Commands = []cli.Command{
	 {
		 Name:    "ping",
		 Usage:   "basic test",
		 Action:  func(c *cli.Context) error {
       log.Println("pong")

			 return nil
		 },
	 },

	 {
		 Name:    "gopack",
		 Usage:   "make a new go package with initial files",
		 Action:  func(c *cli.Context) error {
			 package_name := c.Args().First()
			 if package_name == "" {
				 return errors.New("please provide a package name")
			 }

			 return dev.Gopack(package_name)
		 	},
	 },
 }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
