package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	master_url         string = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
	api_url            string = "https://api.github.com/repos/"
	anchor_marker      string = "](#"
	title_marker       string = "#"
	stared_line_marker string = "](https://github.com/"
)

type (
	Response struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Star      int    `json:"stargazers_count"`
		Watch     int    `json:"watchers_count"`
		Fork      int    `json:"forks_count"`
		Issues    int    `json:"open_issues"`
	}
	StraredStuff map[string]*Title

	Title struct {
		name      string
		maxLength int
		content   []StraredLine
	}

	StraredLine struct {
		Origin string
		Repo   string
		Name   string
	}
)

func main() {
	f, err := readFile("./test.md")
	if err != nil {
		panic(err)
	}

	stuff := make(StraredStuff, 0)
	var title string
	var inContent bool
	for _, v := range f {

		//Title
		if strings.HasPrefix(v, title_marker) {
			title = getTitle(v)

			if title == "contents" {
				inContent = true
			}

			if !inContent {
				continue
			}

			stuff[title] = &Title{
				name:      title,
				maxLength: 0,
				content:   make([]StraredLine, 0),
			}

		}

		//Lines
		if inContent && strings.Index(v, stared_line_marker) > -1 {

			if val, ok := stuff[title]; ok {
				n := getName(v)
				stuff[title].content = append(val.content, StraredLine{
					Origin: v,
					Repo:   getRepo(v),
					Name:   n,
				})
				tn := len(n)
				if tn > stuff[title].maxLength {
					stuff[title].maxLength = tn
				}
			}
		}
	}

	var keys []string
	for k, v := range stuff {
		if len(v.content) > 0 {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	for i, k := range keys {
		fmt.Println(" ", i, " : ", k)
	}

	var key int
	fmt.Print(" Select the desired anchor\n")
	_, err = fmt.Scanf("%d", &key)

	if err != nil {
		panic(err)
	}

	log.Printf(" Desired entry %d : %s\n", key, keys[key])

	if val, ok := stuff[keys[key]]; ok {
		spad := strconv.Itoa(val.maxLength + 3)
		header := "|  Star  |  Fork  |  Watch  |  Issues  |  Udapte"
		s := fmt.Sprintf(" %-"+spad+"s"+header, "NAME")
		fmt.Printf("%s\n", s)
		var br string
		for i := 0; i < val.maxLength+len(header)+20; i++ {
			br = br + "-"
		}
		fmt.Printf("%s\n", br)
		for _, v := range val.content {
			url := api_url + v.Repo
			response, err := http.Get(url)
			if response.StatusCode != 200 {
				defer response.Body.Close()
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Error status code : %d:%s\n", response.StatusCode, string(contents))
			} else {
				if err != nil {
					panic(err)
				} else {
					defer response.Body.Close()
					contents, err := ioutil.ReadAll(response.Body)
					if err != nil {
						panic(err)
					}
					resp := &Response{}
					if err := json.Unmarshal([]byte(contents), &resp); err != nil {
						panic(err)
					}
					s := fmt.Sprintf(" %-"+spad+"s| %-7d| %-7d| %-8d| %-9d| %s", v.Name, resp.Star, resp.Fork, resp.Watch, resp.Issues, resp.UpdatedAt)
					fmt.Printf("%s\n", s)

				}
			}
			//s := fmt.Sprintf(" %-"+spad+"s| %-7d| %-7d| %-8d| %-9d| %s", v.Name, 11, 22, 33, 44, "2019-02-22T10:14:51Z")
		}
	}
}

func readFile(path string) ([]string, error) {
	r := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		return r, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return r, err
	}
	return r, nil
}

func readMaster(path string) ([]string, error) {
	r := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		return r, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return r, err
	}
	return r, nil
}

func getTitle(s string) string {
	return strings.ToLower(strings.Replace(strings.TrimSpace(strings.Replace(s, title_marker, "", -1)), " ", "-", -1))
}

func getRepo(s string) string {
	r := s[strings.Index(s, stared_line_marker)+len(stared_line_marker):]
	r = r[:strings.Index(r, ")")]
	r = strings.TrimSuffix(r, "/")
	return r
}

func getName(s string) string {
	return s[strings.Index(s, "[")+1 : strings.Index(s, "]")]
}
