package tools

import (
	"crypto/md5"
	"unicode"
	"unicode/utf8"
)

type Comment struct {
	CommentID int
	UserID    int
	Owner     string
	Comment   string
	CreatedAt string
	Parent    *int
}

type CommentForRender struct {
	CommentID   int
	Owner       string
	Content     string
	CreatedAt   string
	SonComments [](*CommentForRender)
}

// get a string's MD5 digest
func GetMD5Digest(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return string(m.Sum(nil))
}

// 1.   3 < length < 16
// 2.   only allow characters(a~z A~Z), _, and numbers(0~9)
func CheckUsernameValid(u string) bool {
	if len(u) <= 3 || len(u) >= 16 {
		return false
	}

	for _, v := range u {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) && v != '_' {
			return false
		}
	}

	return true
}

// 1.  5 < length < 21
// 2.  only allow characters(a~z A~Z), _, and numbers(0~9)
func CheckPasswordValid(p string) bool {
	if len(p) <= 5 || len(p) >= 21 {
		return false
	}

	for _, v := range p {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) && v != '_' {
			return false
		}
	}

	return true
}

// the length of md5 digest is 128 bits(16 bytes)
func CheckAuthcodeValid(a string) bool {
	return len(a) == 16
}

// 1 <= unicode length <= 20
func CheckSecurityAnswerValid(s string) bool {
	c := utf8.RuneCountInString(s)
	return c >= 1 && c <= 20
}

func CheckAppendCommentValid(s string) bool {
	c := utf8.RuneCountInString(s)
	return c >= 1 && c <= 50
}

func CheckNewCommentValid(s string) bool {
	c := utf8.RuneCountInString(s)
	return c >= 1 && c <= 256
}
