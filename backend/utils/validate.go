package utils

import (
	"regexp"
	"strings"
)

type Validate struct {};

func (v *Validate) Email(email string) bool {
    email = strings.ToLower(email);
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]{1,64}@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`);

    return re.MatchString(email)
}

func (v *Validate) Username(username string) bool {
    re := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9 -_]{0,20}[A-Za-z0-9_-]$`);
    return re.MatchString(username)
}

func (v *Validate) Fullname(fullname string) bool {
    re := regexp.MustCompile(`^[A-Za-zÀ-ÿ][a-zA-ZÀ-ÿ' -]{1,50}[a-zÀ-ÿ]$`)
    return re.MatchString(fullname)
}

