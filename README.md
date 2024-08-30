# BibleGateway-to-Json

# WIP

Go script to convert BibleGateway.com Bible verses to JSON format
Inspired and based on the work of [BibleGateway-to-Obsidian](https://github.com/selfire1/BibleGateway-to-Obsidian)

## Pre-requisites

### First clone the repository

### Obsidian

You first need to download the Bible as markdown files from [BibleGateway-to-Obsidian](https://github.com/selfire1/BibleGateway-to-Obsidian)

### Go

You need to have Go installed on your machine. You can download it from [here](https://golang.org/dl/)

## Usage

Once you have the markdown files downloaded, copy the folder with the Markdown files (By default WEB) and paste it here.

Then run the following command:

````bash
go run main.go
```

This will generate a `bible.json` file with all the verses in JSON format.
````
