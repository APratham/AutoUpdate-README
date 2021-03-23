package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "math/rand"
    "os"
    "time"
)

func makeReadme(filename string) error {


    rand.Seed(time.Now().UnixNano())
    randsource := rand.NewSource(time.Now().UnixNano())

    fact := rand.Intn(len(content))
    date := time.Now().Format("2 Jan 2006")


                randgenerator := rand.New(randsource)
                firstLoc := randgenerator.Intn(10)
                fact := ""

      // Read entire file content, giving us little control but
      // making it very simple. No need to close the file.

        dat, err := ioutil.ReadFile("facts.txt")
        if err == nil {
                ascii := string(dat)
                splt := strings.Split(ascii, "\n")
    fact = splt[firstLoc]


    // Whisk together static and dynamic content until stiff peaks form
    hello := "### Hello! I’m Victoria Drake. 👋\n\nI’m a software developer at 💜 and Director of Engineering at work. I build my skill stack in public and share open source knowledge through the words I’ve written on [victoria.dev](https://victoria.dev). I hope to encourage people to learn openly and fearlessly, with wild child-like abandon."
  	quote := "⚡ Fun fact: " + fact
  	updated := "<sub>Last updated by magic on " + date + ".</sub>"
  	data := fmt.Sprintf("%s\n\n\n\n%s\n", hello, updated)

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

	makeReadme("../README.md")

}
