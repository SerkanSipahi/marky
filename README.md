# marky

### Features of markdown module
* header, paragraph, emphasized, strong and link

### Requirements
* Golang >= 1.9

### Example usage
```golang
import (
	"io/ioutil"
	"github.com/serkansipahi/marky"
)

func main(){
    markdownTemplate, _ := ioutil.ReadFile("markdown.md")
    markdown := marky.NewMarkdown(string(markdownTemplate))
    code := markdown.Compile()
    fmt.Println(code)
}
```

### Build module
```golang
// type in terminal
git clone https://github.com/SerkanSipahi/marky.git
cd marky
go build *.go
./marky
```