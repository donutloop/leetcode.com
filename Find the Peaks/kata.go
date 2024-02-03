package kata

func findPeaks(mountain []int) []int {
    var peaks []int
    for i := 1; i < len(mountain)-1; i++ {
         if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
             peaks = append(peaks, i)
         }
    }
    return peaks
}