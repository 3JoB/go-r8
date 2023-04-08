package hash

import (
	"fmt"

	"github.com/3JoB/ulib/crypt"
)

func SHA512(cookieSecret, username, password string) string {
	return crypt.SHA256(fmt.Sprintf("%v%v%v", password, cookieSecret, username))
}
