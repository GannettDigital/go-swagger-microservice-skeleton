package main

//go:generate mkdir -p $GOPATH/src/github.com/GannettDigital/$ALLIGATOR
//go:generate git init $GOPATH/src/github.com/GannettDigital/$ALLIGATOR
//go:generate bash -c "(cd $GOPATH/src/github.com/GannettDigital/$ALLIGATOR; git remote add origin https://github.com/GannettDigital/$ALLIGATOR)"
//go:generate go run main.go write.go

func main() {
	write()
}
