#!/bin/bash
git log --format="%H" -n 1 > build/version.txt
