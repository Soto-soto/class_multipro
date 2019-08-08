package main

import (
    "fmt"
	"bufio"
    "io"
    "os"
	"strconv"

)



func merge(left, right []int) (ret []int) {
    ret = []int{}
    for len(left) > 0 && len(right) > 0 {
        var x int
        // ソート済みのふたつのスライスからより小さいものを選んで追加していく (これがソート処理)
        if right[0] > left[0] {
            x, left = left[0], left[1:]
        } else {
            x, right = right[0], right[1:]
        }
        ret = append(ret, x)
    }
    // 片方のスライスから追加する要素がなくなったら残りは単純に連結できる (各スライスは既にソートされているため)
    ret = append(ret, left...)
    ret = append(ret, right...)
    return
}

func sort(left, right []int) (ret []int) {
    // ふたつのスライスをそれぞれ再帰的にソートする
    if len(left) > 1 {
        l, r := split(left)
        left = sort(l, r)
    }
    if len(right) > 1 {
        l, r := split(right)
        right = sort(l, r)
    }

    // ソート済みのふたつのスライスをひとつにマージする
    ret = merge(left, right)
    return
}

func split(values []int) (left, right []int) {
    // スライスを真ん中でふたつに分割する
    left = values[:len(values) / 2]
    right = values[len(values) / 2:]
    return
}

func Sort(values []int) (ret []int) {
    left, right := split(values)
    ret = sort(left, right)
    return
}

func main() {
	var fp *os.File
    var err error

    if len(os.Args) < 2 {
		fp = os.Stdin
    }else {
        fmt.Printf(">> read file: %s\n", os.Args[1])
        fp, err = os.Open(os.Args[1])
        if err != nil {
            panic(err)
        }
        defer fp.Close()
    }

    reader := bufio.NewReaderSize(fp, 4096)
    datas := []int{}
    for {
        line, _, err := reader.ReadLine()
		num := string(line)
		i, _ := strconv.Atoi(num)
		datas = append(datas,i)
        if err == io.EOF {
            break
		} else if err != nil {
            panic(err)
        }
    }

	N:=len(datas)

	q := N/4

    //fmt.Println(sortedValues1)
	ch1 := make(chan []int,q)
    go func() {
		ch1 <- Sort(datas[:q])
    }()

	ch2 := make(chan []int,q)
    go func() {
		ch2 <- Sort(datas[q:q*2])
    }()

	ch3 := make(chan []int,q)
    go func() {
		ch3 <- Sort(datas[q*2:q*3])
    }()

	ch4 := make(chan []int,q)
    go func() {
		ch4 <- Sort(datas[q*3:])
    }()

	aaa := <-ch1
	bbb := <-ch2
	ccc := <-ch3
	ddd := <-ch4

	eee := sort(aaa,bbb)
	fff := sort(ccc,ddd)
	ggg := sort(eee,fff)
	fmt.Println(ggg[:10])
}
