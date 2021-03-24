package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "math/rand"
    "os"
    "strings"
    "time"
)

func makeReadme(filename string) error {

    // Read entire file content, giving us little control but
    // making it very simple. No need to close the file.
    content, err := ioutil.ReadFile("facts.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Convert []byte to string and print to screen
    // text := string(content)
    // fmt.Println(text)

    rand.Seed(time.Now().UnixNano())

    str := string(content)

    lines := strings.Split(str, "\n")
    fact := rand.Intn(len(lines))

    date := time.Now().Format("2 Jan 2006")

    // Whisk together static and dynamic content until stiff peaks form
    const social = `<a href="https://twitter.com/___apratham___" target="_blank" rel="nofollow"><img align="right" alt="Antariksh's Twitter" width="22px" src="https://cdn.jsdelivr.net/npm/simple-icons@v3/icons/twitter.svg" /></a><a href="https://www.linkedin.com/in/APratham" target="_blank" rel="nofollow"><img align="right" alt="Antariksh's LinkedIn" width="22px" src="https://cdn.jsdelivr.net/npm/simple-icons@v3/icons/linkedin.svg" /></a><a href="https://www.instagram.com/___apratham___" target="_blank" rel="nofollow"><img align="right" alt="Antariksh's Instagram" width="22px" src="https://cdn.jsdelivr.net/npm/simple-icons@v3/icons/instagram.svg" /></a>`
    const hello = `### Hi there 👋


                                 __
       ____     ____          __/ /_ __        __
      / _  \   / __ \________/_   _// /_  ____/ /.-..-.
     / __  /  / ____/ __/ _ / /  /_/ __ \/ _ / .-. /, /
    /_/ /_/../_/   /_/ /___/_/____/_/ /_/___/_/  // //

    `

  	quote := "⚡ Fun fact: " + lines[fact]
  	updated := "<sub>Last updated by magic on " + date + ".</sub>"
  	data := fmt.Sprintf("%s%s\n\n%s\n\n%s\n", social, hello, quote, updated)

	  // Prepare file with a light coating of os
	  file, err := os.Create(filename)
	  if err != nil {
		   return err
	   }
	   defer file.Close()

	   // Bake at n bytes per second until golden brown
	   _, err = io.WriteString(file, data)
	   if err != nil {
		     return err
	   }
	    return file.Sync()

}

func main() {

	makeReadme("./README.md")

}
