package generator

import (
	"fmt"
)

func (g goFile) WriteReset(typeName string) error {
	_, err := fmt.Fprintf(g.file,
		"\nfunc (v *%[1]s) Reset() {\n"+
			"\t*v = %[1]s{} // dummy realisation.\n"+
			"}\n",
		typeName,
	)
	return err
}
