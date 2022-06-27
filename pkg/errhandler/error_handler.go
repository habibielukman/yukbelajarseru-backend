package errhandler

import (
	"log"
	"os"
)

var (
    outfile, _ = os.Create("logs3.log")
    l      = log.New(outfile, "", 0)
)

func HandErr(err error) error {
	if err != nil {
		l.Println(err.Error())
		return err
	}
	return nil
}

