// Package main provides
package main

import (
	"bufio"
	"fmt"
	"github.com/ma2ma/libepc"
	"os"
	"strings"
)

func readLinestomap(path string) (map[string]string, error) {
	file, err := os.Open(path)
	var strs []string
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//lines = append(lines, scanner.Text())
		strs = strings.Split(scanner.Text(), ",")
		lines[strs[0]] = scanner.Text()
	}
	return lines, scanner.Err()

}
func encode_epc(m map[string]string) (ress []string, err error) {
	for k, v := range m {
		epc, _, _ := libepc.Encode96bit(k)
		ress = append(ress, epc+","+v)
	}
	return
}
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
func main() {
	var s map[string]string
	var r []string
	s, _ = readLinestomap(os.Args[1])
	r, _ = encode_epc(s)
	writeLines(r, os.Args[2])
	fmt.Println(len(r))
}
