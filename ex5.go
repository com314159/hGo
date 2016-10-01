package main

import (
  "fmt"
  "os"
  "bufio"
  "io"
  "log"
)

func main() {
  var in_file_name, out_file_name string
  if len(os.Args) > 1 {
    in_file_name = os.Args[1]
  }
  if len(os.Args) > 2 {
    out_file_name = os.Args[2]
  }
  fmt.Println(out_file_name)
  if in_file_name == out_file_name{
    fmt.Println("error")
    return
  }
  var err error
  in_file, out_file := os.Stdin, os.Stdout
  if in_file_name != "" {
    if in_file, err = os.Open(in_file_name); err != nil {
      log.Fatal(err)
    }
    defer in_file.Close()
  }
  if out_file_name != "" {
    if out_file, err = os.Create(out_file_name); err != nil {
      log.Fatal(err)
    }
    defer out_file.Close()
  }

  if err = process(in_file, out_file); err != nil {

  }
}

func process(infile io.Reader, outfile io.Writer)(err error){
  reader := bufio.NewReader(infile)
  writer := bufio.NewWriter(outfile)

  defer func() {
    if err == nil {
      err = writer.Flush()
    }
  }()

  eof := false
  for !eof {
    var line string
    line, err = reader.ReadString('\n')
    if err == io.EOF {
      err = nil
      eof = true
    } else if err != nil {
      return err
    }
    if _, err = writer.WriteString(line); err != nil {
      return err
    }
  }
  return nil
}
