[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fmfigurski80%2FNTPeek&count_bg=%2379C83D&title_bg=%23555555&icon=github.svg&icon_color=%23FFFFFF&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Notion Database Peek

Library designed to read, and perform small updates to a Notion Database from your terminal, perfect for cloud-based todos. The latest version also includes the ability to specify query field names and pass in the database id and notion token as positional arguments.

## Basic Usage

Default usage requires just the token and database id (long string id in url of notion database). Due to the length of this command, we recommend setting up a bash alias in your bashrc, which will make it easy to type commands for a specific notion database, as shown here:

![Alias Peeking Usage](http://ntpeek-usage.surge.sh/alias_usage.gif)

Version and Help text can be viewed by calling the tool with `v` or `h` respectively

## Tailoring to Custom Database

This tool is somewhat optimized (read hard-coded), and currently expects a series of fields to display with: a title, timestamp, checkbox, select, and multiselect. The field name defaults are reasonable, but can be customized with the following flags:

- `--title` : name of the title field for actual task name

- `--done` : name of checkbox field to filter incomplete tasks

- `--date` : name of date field

- `--category` : name of category field to use with color

- `--tag` : name of multi-select field to define importance. Tasks with tags that are high-importance will get highlighted, while those without will be dimmed.

Importance of tags is also configurable:

- `--important` : comma-separated list of tags to highlight as important

- `--unimportant` : comma-separated list of tags to dim as un-important

## Installation

The first positional argument is reserved for the Notion Access Token: you will need to create one yourself in order to securely access your Notion content. You can do this by performing the following steps:

1) Navigating to [https://www.notion.so/my-integrations](https://www.notion.so/my-integrations).

2) Clicking on 'New Integration', and filling out the form as you wish (although we do recommend setting the name to NTPeek and using the [official NTPeek Integration logo](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fpublic.notion-static.com%2F9e0bc46d-c5eb-44d6-b1cb-c3542b4d08c0%2Ftenor.gif?id=170a6e36-bec1-44fa-906e-fe06c92f4e8e&table=bot&userId=d9f1afdc-e094-4675-bbb2-e8b8dd390e8e&cache=v2). Make sure to select the correct 'Associated Workspace'. Notably, the current version needs only read permission.

3) Saving the result 'Internal Integration Token' -- you can pass it as the first positional argument when calling the tool from the command line

You can install Notion Database Peek by downloading the pre-compiled binary from the releases tab in Github and selecting the architecture/os that corresponds to your machine. Remember to add the binary to your path.

## Building

This project is built in golang. You can install it yourself, if you have go development tools set up, by running `go install github.com/mfigurski80/NTPeek`.

## Screenshots

![Notion Database Peek](images/Demo.png)
