package game

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	wordsURL = "https://raw.githubusercontent.com/kkrypt0nn/wordlists/main/wordlists/languages"
)

var supportedLangs = map[string]bool{
	"english": true,
	"danish":  true,
}

func generateWordString(n int) string {
	wordStr, err := getWords("english")
	if err != nil {
		fmt.Println("could not get words", err)
		os.Exit(1)
	}

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	words := strings.Fields(wordStr)
	l := len(words)

	var s string
	for i := 0; i < n; i++ {
		s += words[r.Intn(l)] + " "
	}

	return s
}

func getWords(lang string) (string, error) {
	if _, ok := supportedLangs[strings.ToLower(lang)]; !ok {
		return "", errors.New("language is not supported")
	}

	resp, err := http.Get(fmt.Sprintf("%s/%s.txt", wordsURL, lang))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	s, err := formatWhitespace(string(bs))
	if err != nil {
		return "", err
	}

	return s, nil
}
