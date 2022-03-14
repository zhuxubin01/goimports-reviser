// Tool for Golang to sort goimports by 3-4 groups: std, general, local(which is optional) and project dependencies.
// It will help you to keep your code cleaner.
//
// Example:
//	bbimports -project-name github.com/incu6us/bbimports -file-path ./reviser/reviser.go -rm-unused
//
// Input:
// 	import (
//		"log"
//
//		"github.com/incu6us/bbimports/testdata/innderpkg"
//
//		"bytes"
//
//		"github.com/pkg/errors"
// 	)
//
// Output:
//
//	 import (
//		"bytes"
//		"log"
//
//		"github.com/pkg/errors"
//
//		"github.com/incu6us/bbimports/testdata/innderpkg"
//	 )
//
// If you need to set package names explicitly(in import declaration), you can use additional option `-set-alias`.
//
// More:
//
// 	bbimports -h
//
package main
