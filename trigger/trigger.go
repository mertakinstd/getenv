package trigger

import getenv "github.com/mertakinstd/getenv"

func init() {
	getenv.Load().Default()
}
