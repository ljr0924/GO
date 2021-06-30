package algorithms

import (
    "strconv"
    "testing"
)


func QuickSort(list []int, low, high int) []int {

    if low >= high {
        return list
    }

    pivot := list[low]

    l := low
    h := high
    for low < high {
        for low < high && list[high] >= pivot {
            high--
        }
        list[low] = list[high]
        for low < high && list[low] <= pivot {
            low++
        }
        list[high] = list[low]
    }

    list[low] = pivot

    QuickSort(list, l, low-1)
    QuickSort(list, low+1, h)

    return list
}

func TestQuickSort(t *testing.T) {
    var l = []int{1,4,5,32,5,6,85,3,2,9,5,1}

    l = QuickSort(l, 0, len(l)-1)
    var ret string
    for _, i := range l {
        ret += " " + strconv.Itoa(i)
    }

    t.Logf(ret)
}
