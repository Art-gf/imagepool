package stuffs

//my little stuffs

import (
	"log"
)

func ErrorExit(err error) error {
	if err != nil {
		log.Fatalf("Attention: %s", err)
		return err
	}
	return nil
}

func MakeServerCfg() {

}
