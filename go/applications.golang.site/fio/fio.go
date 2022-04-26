package main

import (
	"io"
	"os"
)

func main() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close() // 지연 실행

	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024) // 채널 생성

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF { // 파일 끝을 에러로 처리
			panic(err)
		}

		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt]) // 서브 슬라이싱: 처음(0)부터 끝(cnt-1)까지
		if err != nil {
			panic(err)
		}
	}
}
