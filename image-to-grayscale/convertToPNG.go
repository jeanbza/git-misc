package main

import (
    "bufio"
    "image"
    "image/png"
    "log"
    "os"

    _ "image/jpeg"
)

func main() {
    file := "testsmall.jpg"

    reader, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer reader.Close()

    m, _, err := image.Decode(reader)
    if err != nil {
        log.Fatal(err)
    }

    fw, err := os.Create("testsmall.png")
    defer fw.Close()
    if err != nil {
        log.Fatal(err)
    }
    writer := bufio.NewWriter(fw)

    png.Encode(writer, m)

    writer.Flush()
}