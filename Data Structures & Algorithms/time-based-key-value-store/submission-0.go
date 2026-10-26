type Entry struct {
    timestamp int
    value     string
}

type TimeMap struct {
    store map[string][]Entry
}

func Constructor() TimeMap {
    return TimeMap{
        store: make(map[string][]Entry),
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    this.store[key] = append(this.store[key], Entry{
        timestamp: timestamp,
        value:     value,
    })
}

func (this *TimeMap) Get(key string, timestamp int) string {
    entries := this.store[key]

    left, right := 0, len(entries)-1
    result := ""

    for left <= right {
        mid := left + (right-left)/2

        if entries[mid].timestamp <= timestamp {
            result = entries[mid].value
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return result
}