#!/bin/bash
# Get AUTH_TOKEN from .env in root dir into token.txt
export $(cat .env | xargs)
printf "$NOTION_TOKEN" > build/auth_token.txt
