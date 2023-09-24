package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/v55/github"
	"io/ioutil"
	"strings"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	token64, err := ioutil.ReadFile("token")
	token_bytes, _ := base64.StdEncoding.DecodeString(string(token64))
	token := strings.TrimSpace(string(token_bytes))

	check(err)
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	repos, _, err := client.Repositories.List(ctx, "", nil)
	check(err)

	for _, repo := range repos {
		fmt.Printf("%s\n", github.Stringify(repo.Name))
	}
}
