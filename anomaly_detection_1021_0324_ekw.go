// 代码生成时间: 2025-10-21 03:24:45
package main

import (
    "fmt"
    "math"
    "os"
    "strings"
    "time"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
)

// AnomalyDetection struct holds the configuration for anomaly detection
type AnomalyDetection struct {
    Data []float64
    Threshold float64
}

// NewAnomalyDetection creates a new AnomalyDetection instance
func NewAnomalyDetection(data []float64, threshold float64) *AnomalyDetection {
    return &AnomalyDetection{
        Data: data,
        Threshold: threshold,
    }
}

// DetectAnomalies checks for anomalies in the data
func (ad *AnomalyDetection) DetectAnomalies() ([]int, error) {
    if len(ad.Data) == 0 {
        return nil, fmt.Errorf("data slice is empty")
    }

    anomalies := make([]int, 0)
    mean := calculateMean(ad.Data)
    stdDev := calculateStdDev(ad.Data, mean)

    for i, value := range ad.Data {
        if math.Abs(value-mean) > ad.Threshold*stdDev {
            anomalies = append(anomalies, i)
        }
    }

    return anomalies, nil
}

// calculateMean calculates the mean of a slice of float64
func calculateMean(data []float64) float64 {
    sum := 0.0
    for _, value := range data {
        sum += value
    }
    return sum / float64(len(data))
}

// calculateStdDev calculates the standard deviation of a slice of float64
func calculateStdDev(data []float64, mean float64) float64 {
    variance := 0.0
    for _, value := range data {
        variance += math.Pow(value-mean, 2)
    }
    variance /= float64(len(data) - 1)
    return math.Sqrt(variance)
}

func main() {
    beego.BeeLogger.SetLogger("console", `{"level": "trace"}`)
    defer beego.BeeLogger.Close()

    data := []float64{10, 12, 23, 23, 19, 22, 33, 33, 50, 40}
    threshold := 2.5

    detector := NewAnomalyDetection(data, threshold)
    anomalies, err := detector.DetectAnomalies()
    if err != nil {
        logs.Error("Error detecting anomalies: ", err)
        os.Exit(1)
    }

    fmt.Printf("Detected anomalies at indices: %v
", anomalies)
}
