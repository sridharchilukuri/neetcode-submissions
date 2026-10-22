func topKFrequent(nums []int, k int) []int {
    // 1) Count frequency
    freq := make(map[int]int)
    for _, n := range nums {
        freq[n]++
    }

    // 2) Bucket by frequency (max freq <= len(nums))
    buckets := make([][]int, len(nums)+1)
    for num, f := range freq {
        buckets[f] = append(buckets[f], num)
    }

    // 3) Collect from highest frequency down
    res := make([]int, 0, k)
    for f := len(buckets) - 1; f >= 1 && len(res) < k; f-- {
        for _, num := range buckets[f] {
            res = append(res, num)
            if len(res) == k {
                break
            }
        }
    }

    return res
}
