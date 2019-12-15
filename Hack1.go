package main
import "fmt"
import "os"
import "bufio"

func main() {
    fmt.Println("Hello, World.")
    scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 
}