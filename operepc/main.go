package main

import "fmt"
import "flag"
import "os"
import "github.com/ma2ma/libepc"
func main() {
    var bar_code string
    flag.StringVar(&bar_code, "barcode", "", "barcode")
    var epc string
    flag.StringVar(&epc, "epc", "", "epc")
    flag.Parse()
    if bar_code != "" {
        res,_,e := libepc.Encode96bit(bar_code)
        if e != nil {
            os.Exit(1)
        }
        fmt.Println(res)
    } else if epc != ""{
        res,_,e := libepc.Decode96bit(epc)
        if e != nil {
            os.Exit(1)
        }
        fmt.Println(res)
    } else {
        fmt.Println("error")
        os.Exit(1)
    }
}
