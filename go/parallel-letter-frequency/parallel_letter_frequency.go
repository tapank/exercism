package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var mchan chan FreqMap = make(chan FreqMap)
	fmap := FreqMap{}
	for _, s := range l {
		go func(s string, mchan chan FreqMap) {
			mchan <- Frequency(s)
		}(s, mchan)
	}
	for i := 0; i < len(l); i++ {
		m := <-mchan
		for k, v := range m {
			fmap[k] += v
		}
	}
	return fmap
}
