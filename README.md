# Notion Database Peek

Script to pull tasks from a Notion database and display it in a formatted list. Note that this is 'optimized' (read hard-coded) -- installing this script implies you're:

a) adding your own extension bearer token to the http headers -- see `query.go`
b) adjusting some of the query field names to match your own database -- also `query.go`
c) building this script yourself with `go build && go install`.

![Notion Database Peek](images/Demo.png)
