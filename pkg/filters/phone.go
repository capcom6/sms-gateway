package filters

import (
	"errors"
	"regexp"
)

var phoneCleanRegex = regexp.MustCompile(`[^0-9]`)

func Phone(phone string, allowCity, allowForeign bool) (string, error) {
	phone = phoneCleanRegex.ReplaceAllString(phone, "")
	if len(phone) == 10 &&
		(allowCity || phone[0:1] == "9") {
		if allowForeign {
			return phone, errors.New("can't determine country code")
		}
		return "7" + phone, nil
	}
	if len(phone) == 11 {
		if phone[0:1] == "8" &&
			(allowCity || phone[1:2] == "9") {
			if allowForeign {
				return phone, errors.New("can't determine country code")
			}
			return "7" + phone[1:], nil
		}
		if (phone[0:1] == "7" || allowForeign) &&
			(allowCity || phone[1:2] == "9") {
			return phone, nil
		}
	}

	return phone, errors.New("not a phone number")
}
