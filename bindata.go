package web

//go:generate go-bindata -pkg web -o bindata_gen.go -ignore \.map\$ build/...

var Prefix = "build"
