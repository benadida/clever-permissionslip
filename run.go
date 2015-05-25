package main

import (
	"github.com/go-martini/martini"
	"code.google.com/p/goauth2/oauth"
	clevergo "gopkg.in/Clever/clever-go.v1"
	"log"
)

func main(){
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: "DEMO_TOKEN"},
	}
	
	client := t.Client()
	clever := clevergo.New(client, "https://api.clever.com")
	
	paged := clever.QueryAll("/v1.1/districts", nil)
	
	var district clevergo.District
	for paged.Next() {
		
		if err := paged.Scan(&district); err != nil {
			log.Fatal(err)
		}
		
		log.Println(district)
	}
	
	if err := paged.Error(); err != nil {
		log.Fatal(err)
	}

	m := martini.Classic()
	m.Get("/", func() string {
		return district.Name
	})
	m.Run()
}
