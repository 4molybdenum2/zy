package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"github.com/4molybdenum2/zy/pkg/sitemap"
)
func main() {
	// get the sitelink from command line flags
	var siteLink string
	var depth int
	// create the cli for the app
	app := &cli.App{
		Name: "zy",
		Usage: "Generate a sitemap for a given URL, zy runs a BFS algorithm to parse all the link tags in a HTML page and ignores external links",
		Flags: []cli.Flag {
			&cli.IntFlag{
			Name: "depth",
			Value: 0,
			Usage: "Define the depth for the sitemap",
			Destination: &depth,
			},
			&cli.StringFlag{
				Name: "sitelink",
				Value: "https://molybdenum.netlify.app/",
				Usage: "Define the link for the sitemap",
				Destination: &siteLink,
			},
		},
		Action: func(c *cli.Context) error {
			if depth > 0 {
				pages := sitemap.BuildSitemap(siteLink, depth)
				fmt.Println(pages)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}