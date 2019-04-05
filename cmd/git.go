package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	//Response represents the detail received for a repository in the list
	Response struct {
		name      string
		error     error
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Star      int    `json:"stargazers_count"`
		Watch     int    `json:"watchers_count"`
		Fork      int    `json:"forks_count"`
		Issues    int    `json:"open_issues"`
	}
)

func readMD(url string) ([]string, error) {

	r := make([]string, 0)

	resp, err := http.Get(url)
	if err != nil {
		return r, err
	}

	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return r, err
	}
	return r, nil
}

func callGit(ch chan Response, fla *flags, l straredLine) {
	resp := &Response{}
	resp.name = l.Name
	url := apiURL + l.Repo
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		resp.error = err
		goto End
	}

	if fla.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("token %s", fla.token))
	}
	{
		client := &http.Client{}
		response, err := client.Do(req)
		if response.StatusCode != 200 {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				resp.error = err
				goto End
			}
			resp.error = fmt.Errorf("Error status code : %d:%s\n", response.StatusCode, string(contents))
			goto End
		} else {
			if err != nil {
				resp.error = err
				goto End
			} else {
				defer response.Body.Close()
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					resp.error = err
					goto End
				}

				if err := json.Unmarshal([]byte(contents), &resp); err != nil {
					resp.error = err
					goto End
				}
			}
		}
	}
End:
	ch <- *resp
}

func (r Response) sortingValue(sortingKey string) int {
	switch sortingKey {
	case keyStar:
		return r.Star
	case keyFork:
		return r.Fork
	case keyWatch:
		return r.Watch
	case keyIssues:
		return r.Issues
	default:
		return 0
	}
}
