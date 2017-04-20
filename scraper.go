package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"os/exec"
	_ "strconv"
	"time"
	"regexp"
)



func main() {
	//loop
	for {

		path := createDir()

		dateString := dateString()
		filename := dateString + ".md"

		// create markdown file
		createMarkDown(dateString, path, filename)

		// //TODO: use goroutinez
		scrape("elixir", path, filename)
		scrape("erlang", path, filename)
		scrape("ruby", path, filename)
		scrape("go", path, filename)
		scrape("python", path, filename)
		scrape("php", path, filename)
		scrape("javascript", path, filename)
		scrape("java", path, filename)

		gitPull()
		gitAddAll()
		gitCommit(dateString)
		gitPush()

		time.Sleep(time.Duration(24) * time.Hour)
	}
}

func dateString() string {
	y, m, d := time.Now().Date()
	mStr := fmt.Sprintf("%d", m)
	dStr := fmt.Sprintf("%d", d)
	if m < 10 {
		mStr = fmt.Sprintf("0%d", m)
	}
	if d < 10 {
		dStr = fmt.Sprintf("0%d", d)
	}
	return fmt.Sprintf("%d-%s-%s", y, mStr, dStr)

}

func createDir() string {
	t := time.Now()
	yStr := fmt.Sprintf("%d", t.Year())
	mStr := fmt.Sprintf("%d", t.Month())
	if t.Month() < 10 {
		mStr = fmt.Sprintf("0%d", t.Month())
	}
	return fmt.Sprintf("%v-%s", yStr, mStr)
}

func createMarkDown(date string, path string, filename string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
	    os.Mkdir(path, 0755)
	}

	// open output file
	fo, err := os.Create(path+"/"+filename)
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a write buffer
	w := bufio.NewWriter(fo)
	w.WriteString("### " + date + "\n")
	w.Flush()
}

func replaceSpace(str string) string {
	input := str
    re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
    re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
    final := re_leadclose_whtsp.ReplaceAllString(input, "")
    final = re_inside_whtsp.ReplaceAllString(final, " ")

    return fmt.Sprintf("%v", final)
}

func scrape(language string, path string, filename string) {
	var doc *goquery.Document
	var e error
	// var w *bufio.Writer

	f, err := os.OpenFile(path+"/"+filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// if _, err = f.WriteString(fmt.Sprintf("\n####%s\n", language)); err != nil {
	if _, err = f.WriteString(fmt.Sprintf("####%s", language)); err != nil {
		panic(err)
	}

	if doc, e = goquery.NewDocument(fmt.Sprintf("https://github.com/trending?l=%s", language)); e != nil {
		panic(e.Error())
	}

	doc.Find("ol.repo-list li").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3 a").Text()
		description := s.Find("p.col-9").Text()
		url, _ := s.Find("h3 a").Attr("href")
		url = "https://github.com" + url
		fmt.Println("URL: ", url)
		if _, err = f.WriteString("* [" + replaceSpace(title) + "](" + url + "): " + replaceSpace(description) + "\n"); err != nil {
			panic(err)
		}
	})
}
func gitPull() {
	app := "git"
	arg0 := "pull"
	arg1 := "origin"
	arg2 := "master"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitAddAll() {
	app := "git"
	arg0 := "add"
	arg1 := "."
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitCommit(date string) {
	app := "git"
	arg0 := "commit"
	arg1 := "-am"
	arg2 := date
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}
func gitPush() {
	app := "git"
	arg0 := "push"
	arg1 := "origin"
	arg2 := "master"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}
