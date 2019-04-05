---
# Awesomegostars
---

Awesomegostars is a small tool to get details on the [**Awesome Go**](https://github.com/avelino/awesome-go) list content. For each listed project the detail returned will be :

- The name of the project 
- Its number of :
  - stars
  - forks
  - watchers
  - issues
- Its last update date


Awesomegostars allows to filter the list content by category name and then sort the category content.

#Usage:

```$ awesomegostars <sort-key> [flags]``` 


## Sorting content


Awesomegostars allows to sort a category content by the following `<sort-key>` :

- **star**: descending sort on the stargazers count
- **fork**: descending sort on the forks count
- **watch**: descending sort on the watchers count
- **issues**: descending sort on the open issues count




 

## Filtering the categories

You can filter the desired categories by:

- Its exact matching name
- A text match in the category name


### Filtering by matching name `-c`

Example: 

```$ awesomegostars star -c data-structures```

If the desired category exists the you will get the detail of its content based on the wanted sorting key, here **star**

```
 Sorting on: star
 Desired category: data-structures
 NAME                     |  Star  |  Fork  |  Watch  |  Issues  |  Last update
------------------------------------------------------------------------------------------
 gods                     | 5501   | 586    | 5501    | 16       | 2019-04-04T19:45:01Z
 go-datastructures        | 4851   | 539    | 4851    | 14       | 2019-04-04T06:25:23Z
 boomfilters              | 1091   | 68     | 1091    | 7        | 2019-03-28T13:36:03Z
 golang-set               | 1013   | 104    | 1013    | 6        | 2019-04-04T15:47:12Z
 gota                     | 775    | 86     | 775     | 32       | 2019-04-04T03:23:49Z
 hyperloglog              | 637    | 37     | 637     | 1        | 2019-04-03T01:28:44Z
 willf/bloom              | 601    | 94     | 601     | 5        | 2019-04-04T03:41:02Z
 roaring                  | 583    | 62     | 583     | 47       | 2019-04-03T04:25:00Z
 cuckoofilter             | 456    | 30     | 456     | 6        | 2019-03-29T21:44:32Z
 bitset                   | 451    | 85     | 451     | 2        | 2019-04-04T14:53:26Z
 trie                     | 382    | 60     | 382     | 9        | 2019-04-02T06:35:07Z
 go-geoindex              | 303    | 35     | 303     | 3        | 2019-04-02T04:38:13Z
 mafsa                    | 272    | 18     | 272     | 5        | 2019-04-04T18:38:02Z
 goskiplist               | 185    | 43     | 185     | 3        | 2019-04-03T08:30:31Z
...
```

> Note: The matching name is not case sensitive and the `-` are not required

> Those three invocations will produce the same result

> ```$ awesomegostars star -c data-structures```
>
> ```$ awesomegostars star -c "data structures"```
>
> ```$ awesomegostars star -c DATA-STRUCTURES```

### Filtering by text match `-f`

Example: 

```$ awesomegostars issues -f data```

If any category name contains *data* then you will have to select the desired category :

```
 Soring on: star
 Filtering categories with: data
  0  :  data-structures
  1  :  database
  2  :  database-drivers
  3  :  science-and-data-analysis
 Select the desired category:
```

> Note: The filter condition does not support regular expression nor words combination


## Git API limitation `-t`

Because the Git API has a rate limit for unauthenticated requests (up to 60 requests per hour) the following error message can show up into the detail report once you reach this limit.

```
Error status code : 403:{"message":"API rate limit exceeded for 194.250.98.243. (But here's the good news: Authenticated requests get a higher rate limit. Check out the documentation for more details.)","documentation_url":"https://developer.github.com/v3/#rate-limiting"}
```

If you want to get rid of this message and make up to 5000 requests per hour you can run the command with a Git Personal Acces Token

```
awesomegostars star -f data -t your-personal-access-token
```



