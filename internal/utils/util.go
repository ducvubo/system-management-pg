package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// convert UUID to string, remove -
	uuidString := strings.ReplaceAll((newUUID).String(), "", "")
	// 10clitokenijkasdmfasikdjfpomgasdfgl,masdl;gmsdfpgk
	return strconv.Itoa(userId) + "clitoken" + uuidString
}

func CreateSlug(title string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	slug, _, _ := transform.String(t, strings.ToLower(title))
	
	var result strings.Builder
	for _, char := range slug {
		if unicode.IsLetter(char) || unicode.IsNumber(char) || char == ' ' {
			result.WriteRune(char)
		} else {
			result.WriteRune(' ')
		}
	}
	
	words := strings.Fields(result.String())
	slug = strings.Join(words, "-")
	
	return slug
}

