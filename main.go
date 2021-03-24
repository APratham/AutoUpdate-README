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
    title := "### Hi there 👋"
    const hello = `


                                 __
       ____     ____          __/ /_ __        __
      / _  \   / __ \________/_   _// /_  ____/ /.-..-.
     / __  /  / ____/ __/ _ / /  /_/ __ \/ _ / .-. /, /
    /_/ /_/../_/   /_/ /___/_/____/_/ /_/___/_/  // //

    `

    points := "## About me\n- 🔭 Currently working on Cloud Certifications\n- 🌱 Learning AWS\n- 👯 I'm looking to collaborate on amazing projects as long as they are based on cloud computing\n- 💬 Reach out to me: https://github.com/APratham/APratham/discussions/1\n- 📫 antariksh.pratham@pccoepune.org\n- 😄 Pronouns: He/Him\n"

  	quote := "\n- ⚡ Fun fact: " + lines[fact] + "\n- 🛠️ Currently working on: https://github.com/vigneshshettyin/Flask-Generate-Certificate"
    spotify1 := "## I'm currently jamming to"
    spotify2 := `<a href="https://spotify-now-playing-iota-umber.vercel.app/now-playing?open"><img src="https://spotify-now-playing-iota-umber.vercel.app/now-playing" width="256" height="64" alt="Now Playing"></a>`
    lan1 := "## Languages and Tools"
    lan2 := `<p align="center">
    <div align="center">
        <code><img height="40" src=".images/c.svg"></code>
        <code><img height="40" src="./images/cplusplus.svg"></code>
        <code><img height="40" src="./images/html5.svg"></code>
        <code><img height="40" src="./images/css3.svg"></code>
        <code><img height="40" src="./images/javascript.svg"></code>
        <code><img height="40" src="./images/python.svg"></code>
        <code><img height="40" src="./images/jupyter.svg"></code>
        <code><img height="40" src="./images/docker.svg"></code>
        <code><img height="40" src="./images/kubernetes.svg"></code>
        <code><img height="40" src="./images/heroku.svg"></code>
        <code><img height="40" src="./images/firebase.svg"></code>
        <code><img height="40" src="./images/nodejs.svg"></code>
        <code><img height="40" src="./images/npm.svg"></code>
        <code><img height="40" src="./images/nginx.svg"></code>
        <code><img height="40" src="./images/yarn.svg"></code>
      </div>
    </p>`
    toplan1 := "## Some of my top languages"
    toplan2 := `[![Top Langs](https://github-readme-stats.vercel.app/api/top-langs/?username=APratham&&exclude_repo=now-playing-profile,natemoo-re&langs_count=3&bg_color=30,e96443,904e95&title_color=fff&text_color=fff)](https://github.com/anuraghazra/github-readme-stats)`
    cloud1 := "## Cloud Platforms"
    cloud2 := `![Google Cloud](https://raw.githubusercontent.com/devicons/devicon/master/icons/googlecloud/googlecloud-original.svg)`
    cloud3 := `<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/amazonwebservices/amazonwebservices-original-wordmark.svg" height="200">`
  	updated := "<sub>Last updated by magic on " + date + ".</sub>"
  	data1 := fmt.Sprintf("%s\n\n%s\n%s\n%s\n\n%s\n\n%s\n\n%s\n", social, title, hello, points, quote, spotify1, spotify2)
    data2 := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n", lan1, lan2, toplan1, toplan2, cloud1, cloud2, cloud3, updated)

    data := data1 + data2

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
