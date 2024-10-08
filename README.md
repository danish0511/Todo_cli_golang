
# Todo CLI App

Command Line Interface Todo application built using golang.

## Setting up golang

Visit below given website to download golang
```
https://go.dev/doc/install
```

## Running the application

```
go mod init todo
```
Download table package by copying command given below and run in cli:
```
go get github.com/aquasecurity/table
```
After downloading, run the application
```
go run ./
```
To use a command write that command at the end.

Eg: To view all items in todo list:
```
go run ./ -list
```

Similarly use other commands by replacing ```-list``` 
```
go run ./ -add "<add title>"
go run ./ -edit "<item_number: new title>"
go run ./ -toggle <item_number>
go run ./ -delete <item_number>
```