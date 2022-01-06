package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	// default password length if no argument is pass from cli
	var defaultLength int = 24

	// random seed
	rand.Seed(time.Now().UnixNano())

	// define characters
	m := map[string]string{
		"letters":  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"numbers":  "0123456789",
		"specials": "!@#$%^&*()=+?",
	}

	app := &cli.App{
		// define our cli options here
		Name:            "genpass",
		Usage:           "generate random password",
		UsageText:       "genpass [global option] [password length]",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "nospecial",
				Aliases: []string{"n"},
				Usage:   "don't include special characters",
			},
		},
		// cli action
		Action: func(c *cli.Context) error {
			// convert password length argument to int
			i, err := strconv.Atoi(c.Args().Get(0))
			if err != nil {
				i = defaultLength
			}

			// remove special from map m if --nospecial option is passed
			if c.Bool("nospecial") {
				delete(m, "specials")
			}

			// create byte slice that will contain password
			b := make([]byte, i)

			// build our password
			for i := range b {
				// pick random key and grab its value
				s := randomKeyValue(m)
				// pick random character from the retrieved value from above
				b[i] = randomCharacter(s)
			}

			// print our password
			fmt.Println(string(b))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func randomKeyValue(m map[string]string) string {
	keys := []string{}

	for k := range m {
		keys = append(keys, k)
	}

	return m[keys[rand.Intn(len(keys))]]
}

func randomCharacter(s string) byte {
	return s[rand.Intn(len(s))]
}
