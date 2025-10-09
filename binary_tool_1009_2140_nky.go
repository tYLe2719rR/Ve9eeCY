// 代码生成时间: 2025-10-09 21:40:41
package main

import (
    "bufio"
    "encoding/binary"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

// BinaryTool 是一个用于二进制文件读写的工具
type BinaryTool struct {
    // 这里可以包含任何需要的字段，例如文件指针等
}

// NewBinaryTool 创建一个新的 BinaryTool 实例
func NewBinaryTool() *BinaryTool {
    return &BinaryTool{}
}

// ReadFromFile 从指定的二进制文件中读取数据
// filename 是文件路径
// dataLength 是要读取的数据长度
func (bt *BinaryTool) ReadFromFile(filename string, dataLength int) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    buffer := make([]byte, dataLength)
    _, err = file.Read(buffer)
    if err != nil {
        return nil, err
    }
    return buffer, nil
}

// WriteToFile 将数据写入指定的二进制文件
// filename 是文件路径
// data 是要写入的数据
func (bt *BinaryTool) WriteToFile(filename string, data []byte) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = file.Write(data)
    if err != nil {
        return err
    }
    return nil
}

// ReadInt32FromReader 从bufio.Reader中读取一个int32
func ReadInt32FromReader(reader *bufio.Reader) (int32, error) {
    var number int32
    err := binary.Read(reader, binary.LittleEndian, &number)
    return number, err
}

// WriteInt32ToWriter 将一个int32写入bufio.Writer
func WriteInt32ToWriter(writer *bufio.Writer, number int32) error {
    return binary.Write(writer, binary.LittleEndian, number)
}

func main() {
    bt := NewBinaryTool()

    // 示例：读取二进制文件
    filename := "example.bin"
    dataLength := 1024
    data, err := bt.ReadFromFile(filename, dataLength)
    if err != nil {
        log.Fatalf("Failed to read from file: %v", err)
    }
    fmt.Printf("Read data: %x
", data)

    // 示例：写入二进制文件
    err = bt.WriteToFile(filename, data)
    if err != nil {
        log.Fatalf("Failed to write to file: %v", err)
    }
    fmt.Println("Data written successfully")
}
