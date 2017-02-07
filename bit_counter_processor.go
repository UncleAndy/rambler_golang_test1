package main

import (
	"os"
	"log"
	"github.com/uncleandy/test_bits_count_golang/bit_counter"
	"io"
	"encoding/binary"
	"bytes"
	"sync"
	"fmt"
)

type InputValue struct {
	value uint64
	offset uint64
}

type OutputValue struct {
	value uint8
	offset uint64
}

var (
	result []uint8
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("File name required.")
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	chanel_in := make(chan InputValue, 8)
	chanel_out := make(chan OutputValue, 8)

	wg := &sync.WaitGroup{}
	// Параллельная обработка
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			for in := range chanel_in {
				out := OutputValue{
					value: bit_counter.BitsCount(uint64(in.value)),
					offset: in.offset,
				}
				chanel_out <- out
			}
			wg.Done()
		}()
	}

	// Постановка данных для процессов
	go func() {
		buf := make([]byte, 8)
		current_offset := uint64(0)
		for {
			_, err := io.ReadFull(file, buf)
			if err != nil {
				break;
			}

			var res uint64
			reader := bytes.NewReader(buf)
			err = binary.Read(reader, binary.BigEndian, &res)
			if err != nil {
				log.Print("Error wher read uint64 from bytes:", err)
			} else {
				in := InputValue{
					value: res,
					offset: current_offset,
				}
				chanel_in <- in
				current_offset++
			}
		}
		close(chanel_in)
	}()

	// Детектирование завершения всех процессов обработки и закрытие выходного канала
	go func() {
		wg.Wait()
		close(chanel_out)
	}()

	// Чтение результатов
	max_offset := uint64(0)
	for out := range chanel_out {
		if result == nil {
			result = make([]uint8, out.offset+16)
		} else if uint64(len(result)) <= out.offset {
			result = append(result, make([]uint8, (out.offset - uint64(len(result)) + 16))...)
		}

		result[out.offset] = out.value
		if out.offset > max_offset {
			max_offset = out.offset
		}
	}

	// Вывод результата
	fmt.Printf("Result: %+v\n", result[0:(max_offset+1)])
}
