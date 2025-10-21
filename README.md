# Program for calculating the area and circumference of a circle using Go

## Function to Calculate area of circle

```go
    func area(r float32) float32{
        const phi = 3.14
        return phi * r * r
    }
```

## Function to Calculate circumference of circle

```go
    func circumference(r float32) float32{
        const phi = 3.14
        return 2*phi*r
    }
```
