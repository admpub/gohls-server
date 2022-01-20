//go:generate go install github.com/jteeuwen/go-bindata/go-bindata@latest
//go:generate go-bindata -pkg api -prefix ../../ui/build ../../ui/build/...
package api

import _ "github.com/jteeuwen/go-bindata"
