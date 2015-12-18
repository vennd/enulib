Client library for Go Programming Language
https://golang.org/

Usage
```
import "github.com/vennd/enulib/enulibgo"

func main() {
	enulibgo.Init()

		wallet, err := enulibgo.CreateWallet(blockchain)
		if err != nil {
			log.Fatal(err.Error())
		}
}
```
