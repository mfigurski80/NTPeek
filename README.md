[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fmfigurski80%2FNTPeek&count_bg=%2379C83D&title_bg=%23555555&icon=github.svg&icon_color=%23FFFFFF&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Notion Database Peek

Library designed to read, and perform small updates to a Notion Database from your terminal, perfect for cloud-based todos.

This library is still under development. Note that it is currently somewhat 'optimized' (read hard-coded) -- installing this script implies you're:

- adding your own extension bearer token to the http headers -- ie defining `NOTION_TOKEN` in your `.env` file

- adjusting some of the query field names to match your own database -- in `query.go`

- building this script yourself with `go generate && go build && go install`.


## Usage

Default usage requires just the database id (long string id in url):

![Peeking Usage](images/usage.gif)

Version and Help text can be viewed by calling the tool with `v` or `h` respectively

## Screenshots

![Notion Database Peek](images/Demo.png)
