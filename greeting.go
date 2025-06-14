package greeting

import "time"

func Do() string {
	switch h := time.Now().Hour(); {
	case h >= 4 && h < 10:
		return "おはよう"
	case h >= 10 && h < 17:
		return "こんにちは"
	default:
		return "こんばんは"
	}
}
