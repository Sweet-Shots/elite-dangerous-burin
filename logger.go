package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.8" )

func QRVPXLXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RByt9fKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nN3JFcgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DEtqdlrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePYkGXrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBPf24K4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOgMbGtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcxwpdInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mh3RcmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EWbJQZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFRpm6edWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAYs9Bb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1OlkphgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qcl6hFSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncMpLooBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QROKFfVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYgAulcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEtO3kKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvZfjYSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIFzMRMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8k5RlX3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pV7Wt6ioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9T9Hcbo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyt6fbbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgNUu7wGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUHjLHWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdaxVH55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZRsTFtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtlUmHVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2WKUGPT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3N5dFtOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U01YTehbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4RbUn6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MkT21ADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9bPJZ2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSf520dQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pef2Pp1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lv75ngipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDcWUXu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAj2wwwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQxyxfqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3c6XvOoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdAEc4SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrlO3jD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xe9kQJiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoEe3e5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDPXcabkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABrhY6lNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsOk45m2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLpci9eGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arqq2eCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ike61vRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2xo7vVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqu6bCYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uE9kcDrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkLWJGDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x69T0uMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCC8HkZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlyoQ4Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uEr2hyOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9SIfcM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VS9sBYXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrOaayGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNf86NXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfWxQ7ZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtAwSo7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paYrp9OGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EM0VXjw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bpQOmbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qOwFridDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vagNKMhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03pBSsakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1W9yzlSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5OO5FS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbimL7j6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dp3ZKMk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3g5OquA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZWiJ8zzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36YC573rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxutgP3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dC4JtMzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyJIPom4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8xAPNGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DDVdr7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2d2NKLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHGUY52yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrgxlPymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5o4IFpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nM0FlqOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B03FHXewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AP5MDoI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38PgdyZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ps6zCMxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbE59O5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmgkVNjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rOuAXEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7S5mTybIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtREJUszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87Tf8EKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPYCSX6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6E2iqJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSrc1WxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPxHWZswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jt5MTPoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nv9bors2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5p6p78jdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFTghsfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ve9C1FgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UPvHQVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HmGNlSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9oexh6pIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sPoJ8tCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIAEFnRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YKWDDjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfSdjxDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YVLqMqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REtQn2h4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPkrC6NCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stpiktzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIWAHpnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g9ql67zvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZDd7reVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27rVFfX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qftaAfhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXPpClRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIzjV5heWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwPMqurFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YD5DNO25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVEWkSeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTzdNtY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMaTviY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBXs66XjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soYfYZbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InkCsT6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5swTUIifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCAGw5A4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sej8cIFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CR85NeqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bFIScuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WbMvkT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1FCwsAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g15rbKDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TKx1G7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2Q0b82zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGkcsC05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYTA6bVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djpHinzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysivtUA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIjM4Zv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeYZSlNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRhmIJMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDYocUH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTbMxWXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yp5g16EYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 651GbW4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oX3uevoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbeywfvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5GJM1mWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffVFIUtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8WeBymVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FDUHiGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAy8OBDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gm4ybsyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5XiHf3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBl7eY7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dll16x4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26apZGUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XflmnoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RR1SAXZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVujuDfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zV3Z52cdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J71huo8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SeaHC4LNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBlJIFUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzgN82HCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwE2mjZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6fs0SfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uli70uciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZHCNHkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBnt2xUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qEjGqzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPPv3pmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0MHlp7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RGqkMgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIiilk3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgdMtvtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3crRLV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RezU92ESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWXXAALEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgDNE4JyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uTtFViiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIht4VNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwzQ7M8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycZNmxh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DEPZcPOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wy8QMQ9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXb6fWPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z918rm1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQwZKz8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

