package proceduresgo

import (
	"fmt"
	"time"
)

// Funcion for write consele

func Writeconsole(sTipo, sProcesso, sMsg string) {

	// Funcion for write consele

	var sModo string

	switch sTipo {
	case "W":
		sModo = "WARNING!!! "
	case "E":
		sModo = "** ERROR **"
	case "T":
		sModo = "** TEXTO **"
	default:
		sModo = "           "
	}

	fmt.Printf("%s - %s - %s - %s \n", time.Now().Format("2006/01/02 - 03:04:05 PM"), sModo, sProcesso, sMsg)

}
