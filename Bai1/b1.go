package main
 
import ("fmt"
        //"bufio"
        //"os"
)
 
func main() {
    var n,p,q,r int
 
    n = input("n")
    p = input("p")
    q = input("q")
    r = input("r")
   
    ch :=make(chan int)
    ch2 :=make(chan int)
    ch3 :=make(chan int)
   
    go find(n, p, q, r, ch)
    go find(n, r, p, q, ch2)
    go find(n, q, r, p, ch3)
 
    for i := range ch {
    fmt.Printf(" %v chi het cho %v, %v, ko chia het cho %v\n",i, p, q, r)
    }
    for i := range ch2 {
    fmt.Printf(" %v chi het cho %v, %v, ko chia het cho %v\n",i, r, p, q)
    }
    for i := range ch3 {
    fmt.Printf(" %v chi het cho %v, %v, ko chia het cho %v\n",i, q, r, p)
    }
   
}
 
func input(name string) int {
    var in int
    fmt.Printf("Enter %v: ",name)
    fmt.Scanln(&in)
    fmt.Printf("Entered %v: %v\n",name, in)
    return in
}
 
func find(n, divisible1, divisible2, notDivisible int, ch chan int) {
    for i:=0; i<=n; i++ {
        if i % divisible1 == 0 && i % divisible2 == 0 && i % notDivisible != 0 {
            ch <- i
        }
    }
    close(ch)
}
