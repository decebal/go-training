package main
import (
    "fmt"
)

func main() {
    var n, m int
    fmt.Scan(&n)
    if n < 2 || n > 100 {
        return
    }
    fmt.Scan(&m)

    cnt := map[int]int{}
    nodes := make(map[int][]int, n)
    for i := 1; i <= n; i++ {
        nodes[i] = make([]int, 0)
        cnt[i] = 1
    }

    for i := 1; i <= m; i ++ {
        var edge1, edge2 int
        fmt.Scan(&edge1)
        fmt.Scan(&edge2)
        nodes[edge1] = append(nodes[edge1], edge2)
        nodes[edge2] = append(nodes[edge2], edge1)
    }

    parent := map[int]int{}
    levels := bfs(1, nodes, func (node int, n int) {
        parent[n] = node
    })

    for i := len(levels)-1; i > 0; i-- {
        for _,j := range levels[i] {
            cnt[parent[j]] += cnt[j]
        }
    }
    var answer int
    for i:= 1; i <= n; i++ {
        if (cnt[i] % 2 == 0) {
            answer++
        }
    }
    fmt.Println(answer-1)
}

func bfs(start int, nodes map[int][]int, fn func (int, int)) [][]int {
    frontier := []int{start}
    visited := map[int]bool{}
    levels := [][]int{}
    next := []int{}


    for 0 < len(frontier) {
        next = []int{}
        for _, node := range frontier {
            visited[node] = true
            for _, n := range bfs_frontier(node, nodes, visited) {
                next = append(next, n)
                fn(node, n)
            }
        }
        levels = append(levels, frontier)
        frontier = next
    }
    return levels
}

func bfs_frontier(node int, nodes map[int][]int, visited map[int]bool) []int {
    next := []int{}
    iter := func (n int) bool { _, ok := visited[n]; return !ok }
    for _, n := range nodes[node] {
        if iter(n) {
            next = append(next, n)
        }
    }
    return next
}
