package initial

import "errors"

func Check_for_letters(th string, arg string) error {
	switch th {
	case "standard", "shadow", "thinkertoy":
		for i := 0; i < len(arg); i++ {
			if !(arg[i] >= 32 && arg[i] <= 126) {
				return errors.New("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
			}
		}
	}
	return nil
}
