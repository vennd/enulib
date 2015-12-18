Client library for Go Programming Language
https://golang.org/

Usage
Get an API key from support@vennd.io and place the file enu_key.json in the current working directory of your application.

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
