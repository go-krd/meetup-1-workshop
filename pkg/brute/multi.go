package brute

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Multi bruteforces password using multiple concurrent goroutines.
func Multi(hash string, num int) (password string, err error) {
	b, err := hex.DecodeString(hash)
	if err != nil {
		return "", fmt.Errorf("decode hex: %s", err)
	}

	var target [md5.Size]byte
	copy(target[:], b)

	// Канал, куда будет записано решение.
	solution := make(chan string)

	// Канал, куда воркер будет сигнализировать завершение работы.
	finish := make(chan struct{})

	in := generator()

	for i := 0; i < num; i++ {
		go worker(target, in, solution, finish)
	}

	finished := 0
	for {
		select {
		case password = <-solution:
			return password, nil
		case <-finish:
			finished++
			if finished == num {
				return "", fmt.Errorf("password not found")
			}
		}
	}
}

type job struct {
	start string
	end   byte
}

func generator() chan job {
	out := make(chan job)
	start := []byte("00000")
	go func() {
		for {
			b := NextByte(start[0])
			out <- job{string(start), b}
			if b == '0' {
				close(out)
				return
			}

			start[0] = b
		}
	}()
	return out
}

func worker(hash [md5.Size]byte, in <-chan job, out chan<- string, finish chan<- struct{}) {
	for p := range in {
		b := []byte(p.start)
		for b[0] != p.end {
			if md5.Sum(b) == hash {
				out <- string(b)
				return
			}
			NextWord(b)
		}
	}

	finish <- struct{}{}
}
