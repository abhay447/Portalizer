package queryHandler

import (
	"strings"
)

func ActOnQuery(y string) (z string){
    i := strings.Index(y," ")
    if i > -1 {
        y = "https://google.com/search?q=" + strings.Replace(y, " ", "+", -1)
        return y
    } else{
        i := strings.Index(y,"http")
        if i==0{
            return y
        } else{
            i := strings.Index(y,".")
            if i>-1{
                return "https://"+y
            } else{
                return "https://google.com/search?q=" + y
            }
        }
    }
}

