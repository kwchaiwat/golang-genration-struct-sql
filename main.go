package main

import (
	"fmt"

	"github.com/gohouse/converter"
)

func main() {
	// Initialization
	t2t := converter.NewTable2Struct()

	t2t.Config(&converter.T2tConfig{
		// If the initial of the field is capitalized, tag will not be added, false will be added by default, and true will not be added
		RmTagIfUcFirsted: false,
		// Whether the field name of tag is converted to lowercase? If it has uppercase letters, the default value is false
		TagToLower: true,
		// If you want to convert other letters to lowercase while the initial of the field is uppercase, false will not be converted by default
		UcFirstOnly: false,
		////Put each struct into a separate file, false by default, and put it into the same file (not provided yet)
		//SeperatFile: false,
	})

	// Start migration transformation
	err := t2t.Table("information_receipts").
		TagKey("gorm").
		SavePath("/Users/m1schaiwat/Desktop/ETC/myWorkSpace/golang-genration-struct-sql/model/model.go").
		Dsn("root:@tcp(localhost:3306)/KC?charset=utf8").
		Run()

	fmt.Println(err)
}
