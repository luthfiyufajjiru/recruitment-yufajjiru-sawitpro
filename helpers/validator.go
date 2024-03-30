package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/SawitProRecruitment/UserService/generated"
)

// Rules:
// 1. Phone numbers must be at minimum 10 characters and maximum 13 characters.
// 2. Phone numbers must start with the Indonesia country code “+62”.
// 3. Full name must be at minimum 3 characters and maximum 60 characters.
// 4. Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 capital characters AND 1 number AND 1 special (non alpha-numeric) characters.
func RegistrationValidator(inp generated.UserRegistrationRequest) (err error) {
	rules := []string{
		"Phone numbers must be at minimum 10 characters and maximum 13 characters.",
		"Phone numbers must start with the Indonesia country code “+62”.",
		"Phone numbers must be a number.",
		"Full name must be at minimum 3 characters and maximum 60 characters.",
		"minimum 6 characters and maximum 64 characters",
		"containing at least 1 capital characters",
		"1 number",
		"1 special (non alpha-numeric) characters",
	}

	passValidation := make(chan string)

	go func(ch chan string) {
		var sbPass strings.Builder
		lnPass := len(inp.Password)
		ctr := 0

		if lnPass < 6 && lnPass > 64 {
			content := fmt.Sprintf(" %s", rules[4])
			sbPass.WriteString(content)
			ctr++
		}

		var (
			hasUpper     bool
			hasNumber    bool
			hasSpecial   bool
			specialChars = ",!#"
		)

		for _, r := range inp.Password {
			if unicode.IsUpper(r) && unicode.IsLetter(r) && !hasUpper {
				hasUpper = true
			}
			if strings.ContainsAny(string(r), specialChars) && !hasSpecial {
				hasSpecial = true
			}
			if _, err := strconv.Atoi(string(r)); err == nil && !hasNumber {
				hasNumber = true
			}
		}

		if !hasUpper {
			content := fmt.Sprintf(" %s", rules[5])
			if ctr > 0 {
				content = fmt.Sprintf(", %s", rules[5])
			}
			sbPass.WriteString(content)
			ctr++
		}

		if !hasNumber {
			content := fmt.Sprintf(" %s", rules[6])
			if ctr > 0 {
				content = fmt.Sprintf(" AND %s", rules[6])
			}
			sbPass.WriteString(content)
			ctr++
		}

		if !hasSpecial {
			content := fmt.Sprintf(" %s", rules[7])
			if ctr > 0 {
				content = fmt.Sprintf(" AND %s", rules[7])
			}
			sbPass.WriteString(content)
		}

		sbPass.WriteString(".")

		ch <- sbPass.String()
	}(passValidation)

	var sb strings.Builder

	lnPhone := len(inp.PhoneNumber)

	ctr := 0
	if inp.PhoneNumber == "" {
		sb.WriteString(rules[0])
		sb.WriteString(" " + rules[1])
		sb.WriteString(" " + rules[2])
	} else if lnPhone < 10 || lnPhone > 13 {
		sb.WriteString(rules[0])
		ctr++
	}

	if lnPhone > 2 {
		prefix := inp.PhoneNumber[:3]
		if prefix != "+62" {
			content := rules[1]
			if ctr > 0 {
				content = " " + content
			}
			sb.WriteString(content)
		}
	} else if lnPhone < 3 {
		content := rules[1]
		if ctr > 0 {
			content = " " + content
		}
		sb.WriteString(content)
	}

	if lnPhone > 1 {
		var nan bool
		for i := 1; i < lnPhone; i++ {
			_, err := strconv.Atoi(inp.PhoneNumber[i : i+1])
			if err != nil {
				nan = true
				break
			}
		}
		content := rules[2]
		if ctr > 0 {
			content = " " + content
		}
		if nan {
			sb.WriteString(content)
		}
	}

	lnName := len(inp.Name)
	if lnName < 3 || lnName > 60 {
		content := rules[3]
		if ctr > 0 {
			content = " " + content
		}
		sb.WriteString(content)
	}

	passValidationStr := <-passValidation
	if len(passValidationStr) > 0 {
		content := "Passwords must be"
		if ctr > 0 {
			content = " " + content
		}
		sb.WriteString(content)
		sb.WriteString(passValidationStr)
	}

	return
}
