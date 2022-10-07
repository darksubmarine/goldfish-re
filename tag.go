package goldfish_re

import "strings"

const tag_ = "gre"

func parseTag(tag string) (object, attribute, value string, err error) {

	values := strings.Split(tag, ",")
	toRet := map[string]string{}
	for _, val := range values {
		part := strings.Split(val, "=")
		if len(part) != 2 {
			return emptyStr, emptyStr, emptyStr, ErrMalformedTag
		}
		toRet[part[0]] = part[1]
	}

	return toRet["object"], toRet["attribute"], toRet["value"], nil
}
