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
  //  hello := "## Hello! I’m Victoria Drake. 👋\n\nI’m a software developer at 💜 and Director of Engineering at work. I build my skill stack in public and share open source knowledge through the words I’ve written on [victoria.dev](https://victoria.dev). I hope to encourage people to learn openly and fearlessly, with wild child-like abandon."
  	quote = "⚡ Fun fact: " + lines[fact]
  	updated = "<sub>Last updated by magic on " + date + ".</sub>"
  	data := fmt.Sprintf("%s\n\n%s\n\n%s\n", social, quote, updated)

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
