package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }

var replaceNumbers = map[rune]string{
	'1': "XONEX",
	'2': "XTWOX",
	'3': "XTHREEX",
	'4': "XFOURX",
	'5': "XFIVEX",
	'6': "XSIX",
	'7': "XSEVENX",
	'8': "XEIGHTX",
	'9': "XNINEX",
	'0': "XZEROX",
	' ': "XSPACEX",
}
var replaceNumbersBack = map[string]rune{
	"XONEX":   '1',
	"XTWOX":   '2',
	"XTHREEX": '3',
	"XFOURX":  '4',
	"XFIVEX":  '5',
	"XSIX":    '6',
	"XSEVENX": '7',
	"XEIGHTX": '8',
	"XNINEX":  '9',
	"XZEROX":  '0',
	"XSPACEX": ' ',
}

func normalizeText(s string) (strippedStr string) {
	normalizedStr := norm.NFD.String(strings.ToUpper(s))
	strippedStr = strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' || r == ' ' || unicode.IsDigit(r) {
			return r
		}
		return -1
	}, normalizedStr)
	return
}

func (app *App) Encrypt(plaintext, keyA, keyB string) string {
	fmt.Println(plaintext, keyA, keyB)
	a, err := strconv.Atoi(keyA)
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.Atoi(keyB)
	if err != nil {
		fmt.Println(err)
	}
	ciphertext := ""
	plaintext = normalizeText(plaintext)
	for _, r := range plaintext {
		if replace, ok := replaceNumbers[rune(r)]; ok {
			ciphertext += replace
			continue
		}
		// hledá číselnou hodnotu písmena z plaintextu,
		// pokud by r bylo 'c', tak ASCII hodnota 'c' je 99 a ASCII hodnota 'a' is 97. Takže výraz int(r - 'a') vrací int(99 - 97).
		x := int(r - 'A')
		c := (a*x + b) % 26
		ciphertext += string(rune(c + 'A'))
	}
	return fmt.Sprintf(ciphertext)
}

func (app *App) Decrypt(ciphertext, keyA, keyB string) string {
	a, err := strconv.Atoi(keyA)
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.Atoi(keyB)
	if err != nil {
		fmt.Println(err)
	}
	// v tomto for cyklu se hledá inverzní prvek k proměnné a, který ukládá ho do aInv, jako a inverse
	aInv := 0
	for i := 0; i < 26; i++ {
		if (a*i)%26 == 1 {
			aInv = i
			break
		}
	}

	for key, value := range replaceNumbersBack {
		ciphertext = strings.ReplaceAll(ciphertext, key, string(value))
	}

	plaintext := ""
	for _, r := range ciphertext {
		if unicode.IsDigit(r) || r == ' ' {
			plaintext += string(r)
			continue
		}
		x := int(r - 'A')
		m := (aInv * (x - b)) % 26
		if m < 0 {
			m = m + 26
		}
		plaintext += string(rune(m + 'A'))

	}
	return fmt.Sprintf("%s", plaintext)
}
