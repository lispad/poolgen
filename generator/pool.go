package generator

import (
	"fmt"
)

func (g goFile) WritePool(structName string, poolName string) error {
	if poolName == "" {
		poolName = genPoolName(structName)
	}

	_, err := fmt.Fprintf(g.file,
		"\ntype %[1]s struct {\n"+
			"\tsync.Pool\n"+
			"}\n\n"+
			"func New%[1]s(stdLen int) *%[1]s {"+
			"\treturn &JsonNumberSlicePool{\n"+
			"\t\tPool: sync.Pool{New: func() interface{} {\n"+
			"\t\t\treturn %[2]s{}\n"+
			"\t\t}},\n"+
			"\t}\n"+
			"}\n"+
			"func (v *%[1]s) Get() %[2]s {\n"+
			"\treturn %[1]s.Pool.Get().(%[2]s)\n"+
			"}\n"+
			"func (v *%[1]s) Put(el %[2]s) {\n"+
			"\tel.Reset()\n"+
			"\treturn v.Pool.Put(el)\n"+
			"}\n",
		poolName,
		structName,
	)

	return err
}

func genPoolName(structName string) string {
	return structName + "Pool"
}
