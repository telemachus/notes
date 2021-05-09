# Introducing Go, Chapter 9: Testing

Go provides excellent support for testing. First, there is the `go test` command. Second, there is `testing` in the standard library. 

Here’s a bit of sample code.

```go
type testpair struct {
    ns       []float64
    expected float64
}

var testTable = []testpair{
    {[]float64{1, 2}, 1.5},
    {[]float64{1, 1, 1, 1, 1, 1}, 1},
    {[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
    for _, pair := range testTable {
        actual := Average(pair.ns)
        if pair.expected != actual {
            t.Errorf("expected %f, but got %f\n", pair.expected, actual)
        }
    }
}

func TestAverageEmptySlice(t *testing.T) {
    actual := Average([]float64{})
    if !math.IsNaN(actual) {
        t.Errorf("expected NaN, got %f\n", actual)
    }
}
```
