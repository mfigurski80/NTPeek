#!/bin/bash
# Get the database ID from the database name
export $(cat .env | xargs)
printf "$NOTION_DB" > build/db_id.txt
