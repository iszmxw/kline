package mobile

import "strings"

func HandleMobileParams(mobile string, nation string) string {
	mobile = strings.TrimSpace(mobile)
	if isNull(mobile) {
		return mobile
	}

	switch nation {
	case "Indonesia":
		if strings.HasPrefix(mobile, "0") {
			mobile = "62" + mobile[1:]
		}
		if strings.HasPrefix(mobile, "8") {
			mobile = "62" + mobile
		}
		if strings.HasPrefix(mobile, "62") {
			mobile = mobile
		}
	case "India":
		if !strings.HasPrefix(mobile, "91") {
			mobile = "91" + mobile
		}
	default:
		// 不处理
	}

	return mobile
}

func isNull(s string) bool {
	return len(s) == 0
}
