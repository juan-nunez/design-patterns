package main

import (
    "log"
)

type File struct{
    state string
}

func newFile() *File {
    return &File{"untouched"}
}

func (f *File) open() {
    f.state = "opened"
}

func (f *File) close() {
    f.state = "closed"
}

type Command interface {
    execute()
}

type OpenFileCommand struct {
    file *File
}

func (o OpenFileCommand) execute() {
    o.file.open()
}

type CloseFileCommand struct{
    file *File
}

func (c CloseFileCommand) execute() {
    c.file.close()
}

func main() {
    file := newFile()
    log.Println("File's state before being touched: "+ file.state) 
    openCommand := OpenFileCommand{file}
    closeCommand := CloseFileCommand{file}
    openCommand.execute()
    log.Println("File state after running open command: " + file.state)
    closeCommand.execute()
    log.Println("File state after running close command: " + file.state)


}
