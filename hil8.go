/*
 * Copyright (c) 2022 Gilles Chehade <gilles@poolp.org>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var color string
	var useRegex bool

	flag.StringVar(&color, "color", "green", "color to use for hilighting")
	flag.BoolVar(&useRegex, "regexp", false, "use regular expression")
	flag.Parse()

	if flag.NArg() == 0 {
		os.Exit(1)
	}

	var colorCode string
	switch color {
	case "red":
		colorCode = "\u001b[31m"
	case "green":
		colorCode = "\u001b[32m"
	case "yellow":
		colorCode = "\u001b[33m"
	case "blue":
		colorCode = "\u001b[34m"
	default:
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		skip := false
		for _, pattern := range flag.Args() {
			matched := false
			if useRegex {
				matched, _ = regexp.MatchString(pattern, scanner.Text())
			} else {
				if strings.Contains(scanner.Text(), pattern) {
					matched = true
				}
			}
			if matched {
				fmt.Printf("%s%s%s\n", colorCode, scanner.Text(), "\033[0m")
				skip = true
				break
			}

		}
		if !skip {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
