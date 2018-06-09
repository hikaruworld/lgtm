package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

// Commands is supported CLI Command list
var Commands = []cli.Command{
	commandRand,
	commandIn,
}

var commandRand = cli.Command{
	Name:  "rand",
	Usage: "select random lgtm image",
	Description: `
	your's github repository search.
	select random lgtm image.
`,
	Action: doRand,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "list, l", Usage: "all list images."},
	},
}

var commandIn = cli.Command{
	Name:  "in",
	Usage: "select https://lgtm.in/g image",
	Description: `
	lgtm.in get image
`,
	Action: doIn,
}

func client() (context.Context, *http.Client) {
	ctx := context.Background()

	token := os.Getenv("LGTM_GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return ctx, tc
}

func doRand(c *cli.Context) error {
	ctx, tc := client()
	client := github.NewClient(tc)

	owner := os.Getenv("LGTM_GITHUB_OWNER")
	repo := os.Getenv("LGTM_GITHUB_REPO")
	path := os.Getenv("LGTM_GITHUB_ROOT_PATH")
	_, contents, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)

	if err != nil {
		return err
	}

	exact := func(c *github.RepositoryContent) string { return c.GetHTMLURL() }
	urls := mapString(contents, exact)

	if c.Bool("list") {
		for _, u := range urls {
			fmt.Println(u)
		}
	} else {
		url := choice(urls)
		fmt.Println(url)
	}

	return nil
}

type output struct {
	ImageURL string `json:"imageUrl"`
}

func doIn(c *cli.Context) error {
	url := "https://lgtm.in/g"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	output := output{}
	err = json.Unmarshal(body, &output)

	fmt.Println(output.ImageURL)

	return nil
}

func mapString(rc []*github.RepositoryContent, f func(*github.RepositoryContent) string) []string {
	rcm := make([]string, len(rc))

	for i, v := range rc {
		rcm[i] = f(v)
	}

	return rcm
}

func choice(urls []string) string {
	rand.Seed(time.Now().UnixNano())
	return urls[rand.Intn(len(urls))]
}
