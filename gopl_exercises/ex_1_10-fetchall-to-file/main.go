// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// With some changes by Leandro Motta Barros
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// The exercise asks about caching. I did some experiments in which the second
// and subsequent downloads took less time than the first one, suggesting there
// is some caching scheme (which surprised me quite a bit; I need to learn much
// more about Go and/or the Internet protocols). However, I couldn't reproduce
// this later, so I am clueless about caching in `http.Get`.

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// I guess I should make this simpler.
type channelData struct {
	downloadTimeInSecs float64
	status             string
	url                string
	resp               *http.Response
}

func main() {
	start := time.Now()
	ch := make(chan channelData)

	outFile, err := os.Create("whatever_i_was_asked_to_download.out") // Could read name from Args, but nah!

	if err != nil {
		fmt.Fprintf(os.Stderr, "Coulndn't create output file: %v", err)
		return
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		data := <-ch

		if data.resp == nil {
			fmt.Println(data.status)
			continue
		}

		nbytes, err := io.Copy(outFile, data.resp.Body)
		data.resp.Body.Close() // don't leak resources
		if err != nil {
			fmt.Fprintf(os.Stderr, "while reading %s: %v\n", data.url, err)
			return
		}
		secs := time.Since(start).Seconds()
		fmt.Printf("%.2fs  %7d  %s\n", secs, nbytes, data.url)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- channelData) {
	start := time.Now()
	resp, err := http.Get(url)
	var cd channelData

	if err != nil {
		cd.status = fmt.Sprint(err)
		ch <- cd
		return
	}

	cd.url = url
	cd.resp = resp
	cd.downloadTimeInSecs = time.Since(start).Seconds()
	ch <- cd
}

//!-
