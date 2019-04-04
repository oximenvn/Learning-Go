package main
 
import ("fmt"
        //"bufio"
        //"os"
		"time"
)
 
func main() {
    var n,p,q,r int
 
    n = input("n")
    p = input("p")
    q = input("q")
    r = input("r")
   
    // ch :=make(chan int)
    // ch2 :=make(chan int)
    // ch3 :=make(chan int)
   
    start := time.Now()
    count := 0

    // go find3(n, p, q, r, ch)
    // go find3(n, r, p, q, ch2)
    // go find3(n, q, r, p, ch3)
 
    // for i := range ch {
    // fmt.Printf(" %v chi het cho %v, %v, ko chia het cho %v\n",i, p, q, r)
    // }
    // for i := range ch2 {
    // fmt.Printf(" %v chi het cho %v, %v, ko chia het cho %v\n",i, r, p, q)
    // }
    // for i := range ch3 {
    // fmt.Printf(" %v chi het cho %v, %v, ko chia het cho %v\n",i, q, r, p)
    // }

    count = counter(n, p, q, r) + counter(n, r, p, q) + counter(n, q, r, p)

   	fmt.Println(time.Since(start))
   	fmt.Println(count)
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


func find2(n, divisible1, divisible2, notDivisible int, ch chan int) {

    BCNN := divisible1 / UCLN(divisible1, divisible2) * divisible2
    for i:=1;  i <= n /BCNN ; i++ {
        if BCNN * i % notDivisible != 0 {
            ch <- BCNN * i 
        }
    }
    close(ch)
}

func find3(n, divisible1, divisible2, notDivisible int, ch chan int) {

    BCNN := divisible1 / UCLN(divisible1, divisible2) * divisible2
    UCLN3 := UCLN(BCNN, notDivisible)
    p3 := notDivisible / UCLN3

    for i:=0;  i <= n /BCNN ; i+=p3 {
    	for x:=i+1; x < i+p3 && x <= n /BCNN; x++ {
    		ch <- BCNN * x
    	}
    }
    close(ch)
}

func counter(n, divisible1, divisible2, notDivisible int) int {

    BCNN := divisible1 / UCLN(divisible1, divisible2) * divisible2
    UCLN3 := UCLN(BCNN, notDivisible)
    p3 := notDivisible / UCLN3
    return n /BCNN - n /BCNN/p3
}

func UCLN(a, b int) int {
    if(a==0 || b==0){
        return a+b
    }
    for a!=b {
        if (a > b){
            a -= b
        }else{
            b -= a
        }
    }
    return a
}