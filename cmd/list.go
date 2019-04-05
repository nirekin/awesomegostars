package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type (
	title struct {
		name      string
		maxLength int
		content   []straredLine
	}

	straredLine struct {
		Origin string
		Repo   string
		Name   string
	}

	flags struct {
		sorting  string
		token    string
		filter   string
		category string
	}
)

var (
	rootCmd = &cobra.Command{
		Use:   "awesomegostars <sort-key>",
		Short: "Awesomegostars is a tool to get details on the Awesome Go content",
		Long: `Awesomegostars is a tool to get details on the Awesome Go content. 
		
 Available sorting keys are:
 - star: descending sort on the stargazers count
 - fork: descending sort on the forks count
 - watch: descending sort on the watchers count
 - issues: descending sort on the open issues count	
`,
		Example: `> awesomegostars star -f data
Give the choice to get the detail all categories containing "data"				

> awesomegostars star -c data-structures
Give the detail of the category "data-structures"				

 Available <sort-key> are:
 - star: descending sort on the stargazers count
 - fork: descending sort on the forks count
 - watch: descending sort on the watchers count
 - issues: descending sort on the open issues count	
`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(" Sorting on: %s\n", args[0])
			fla.sorting = args[0]
			if fla.token != "" {
				fmt.Printf(" Git Personal access token: %s\n", fla.token)
			}
			if fla.filter != "" {
				fmt.Printf(" Filtering categories with: %s\n", fla.filter)
			}
			if fla.category != "" {
				fmt.Printf(" Desired category: %s\n", fla.category)
			}
			run()
		},
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{keyStar, keyFork, keyWatch, keyIssues},
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := cobra.OnlyValidArgs(cmd, args); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if fla.category != "" && fla.filter != "" {
				fmt.Printf(" The filter \"%s\" will be ignore because the category \"%s\" has been provided", fla.filter, fla.category)
				fla.filter = ""
			}

			if fla.category != "" {
				fla.category = getCategory(fla.category)
			}

			if fla.filter != "" {
				fla.filter = getCategory(fla.filter)
			}
		},
	}
	fla *flags
)

// Execute runs the cobra command
func Execute() {
	fla = &flags{}

	rootCmd.Flags().StringVarP(&fla.token, "token", "t", "", "The Git personal access token.")
	rootCmd.Flags().StringVarP(&fla.filter, "filter", "f", "", "A filter on the listed content.")
	rootCmd.Flags().StringVarP(&fla.category, "category", "c", "", "The name of the desired category of content.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() {
	md, err := readMD(masterURL)
	if err != nil {
		panic(err)
	}

	stared := make(map[string]*title, 0)
	var titleStr string
	var inContent bool
	for _, mdl := range md {
		//Title
		if strings.HasPrefix(mdl, titleMarker) {
			titleStr = getTitle(mdl)

			if titleStr == "contents" {
				inContent = true
			}

			if !inContent {
				continue
			}

			// Apply the filter on titles if required
			if fla.filter != "" && !strings.Contains(titleStr, strings.ToLower(fla.filter)) {
				continue
			}

			stared[titleStr] = &title{
				name:      titleStr,
				maxLength: 0,
				content:   make([]straredLine, 0),
			}
		}

		//Lines
		if inContent && strings.Index(mdl, staredLineMarker) > -1 {

			if val, ok := stared[titleStr]; ok {
				n := getName(mdl)
				stared[titleStr].content = append(val.content, straredLine{
					Origin: mdl,
					Repo:   getRepo(mdl),
					Name:   n,
				})
				// We keep the max length of the title for formatting purpose... later
				tn := len(n)
				if tn > stared[titleStr].maxLength {
					stared[titleStr].maxLength = tn
				}
			}
		}
	}

	if fla.category != "" {
		if _, ok := stared[fla.category]; !ok {
			fmt.Printf("Desired category of content \"%s\" cannot be located\n", fla.category)
			os.Exit(1)
		}
	} else {
		if fla.filter != "" && len(stared) == 0 {
			fmt.Printf("There is no category corresponding to the filter \"%s\"\n", fla.filter)
			os.Exit(1)
		}

		// Prepare the list of available categories
		var keys []string
		for k, v := range stared {
			if len(v.content) > 0 {
				keys = append(keys, k)
			}
		}
		sort.Strings(keys)
		for i, k := range keys {
			fmt.Println(" ", i, " : ", k)
		}

		var key int
		fmt.Print(" Select the desired category: ")
		_, err = fmt.Scanf("%d", &key)

		if err != nil {
			panic(err)
		}

		if key < 0 || key+1 > len(keys) {
			fmt.Println(" Hey you need to choose something within the list... ")
			os.Exit(1)
		}

		fmt.Printf(" Desired category %d : %s\n", key, keys[key])
		fla.category = keys[key]
	}

	if val, ok := stared[fla.category]; ok {
		respCh := make(chan Response)

		for _, v := range val.content {
			go callGit(respCh, fla, v)
		}

		toSort := make([]Response, 0)
		for i := 0; i < len(val.content); i++ {
			toSort = append(toSort, <-respCh)
		}
		close(respCh)

		responses, err := sortResponses(fla.sorting, toSort)
		if err != nil {
			panic(err)
		}

		spad := strconv.Itoa(val.maxLength + 3)
		header := "|  Star  |  Fork  |  Watch  |  Issues  |  Last update"
		s := fmt.Sprintf(" %-"+spad+"s"+header, "NAME")
		fmt.Printf("%s\n", s)
		var br string
		for i := 0; i < val.maxLength+len(header)+20; i++ {
			br = br + "-"
		}
		fmt.Printf("%s\n", br)

		for _, resp := range responses {
			if resp.error != nil {
				s := fmt.Sprintf(" %-"+spad+"s| %s", resp.name, resp.error.Error())
				fmt.Printf("%s\n", s)
			} else {
				s := fmt.Sprintf(" %-"+spad+"s| %-7d| %-7d| %-8d| %-9d| %s", resp.name, resp.Star, resp.Fork, resp.Watch, resp.Issues, resp.UpdatedAt)
				fmt.Printf("%s\n", s)
			}

		}
	}
}
