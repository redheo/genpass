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
	rand.Seed(time.Now().UnixNano())

	m := map[string]string{
		"letters":  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"numbers":  "0123456789",
		"specials": "!@#$%^&*()=+?",
	}

	app := &cli.App{
		Name:  "genpass",
		Usage: "generate random password",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "nospecial",
				Aliases: []string{"n"},
				Usage:   "don't include special characters",
			},
		},
		Action: func(c *cli.Context) error {
			i, err := strconv.Atoi(c.Args().Get(0))

			if err != nil {
				i = 24
			}

			b := make([]byte, i)

			for i := range b {
				s := randomKey(m)
				b[i] = randomCharacter(s)
			}

			fmt.Println(string(b))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func randomKey(m map[string]string) string {
	keys := []string{}

	for k := range m {
		keys = append(keys, k)
	}

	return m[keys[rand.Intn(len(keys))]]
}

func randomCharacter(s string) byte {
	return s[rand.Intn(len(s))]
}
