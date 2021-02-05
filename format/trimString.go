package main

import (
	"regexp"
	"strings"
)

// https://github.com/yaniswang/HTMLHint/wiki/attr-unsafe-chars
var reSpace = regexp.MustCompile("[\\s,\u0000-\u0009\u000b\u000c\u000e-\u001f\u007f-\u009f\u00ad\u0600-\u0604\u070f\u17b4\u17b5\u200c-\u200f\u2028-\u202f\u2060-\u206f\ufeff\ufff0-\uffff]")

func TrimSpaceOnlyInHeadAndEnd(in string) (out string) {
	out = strings.Trim(in, "　")
	return
}

func TrimSpace(in string) (out string) {
	in = strings.TrimSpace(in)
	in = strings.Trim(in, "　")
	// remove non-breaking space
	// https://en.wikipedia.org/wiki/Non-breaking_space
	in = strings.Replace(in, "\u00a0", "", -1)
	in = strings.Replace(in, "\u200b", "", -1)
	out = reSpace.ReplaceAllString(in, "")
	return
}

var reSpaceV2 = regexp.MustCompile("[\\s,　,\u00a0,\u200b,\u0000-\u0009\u000b\u000c\u000e-\u001f\u007f-\u009f\u00ad\u0600-\u0604\u070f\u17b4\u17b5\u200c-\u200f\u2028-\u202f\u2060-\u206f\ufeff\ufff0-\uffff]")

// TrimSpaceByRe 通过正则替换字符串中的空格和一下特殊字符
func TrimSpaceByRe(in string) (out string) {
	return reSpaceV2.ReplaceAllString(in, "")
}
