// cmd/chaincodeauditorsolutionsultra/main.go
package main

import (
"flag"
"log"
"os"

"chaincodeauditorsolutionsultra/internal/chaincodeauditorsolutionsultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := chaincodeauditorsolutionsultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
