package hw

import (
        "math"
        "testing"
)

func TestGeom_CalculateDistance(t *testing.T) {
        tests := []struct {
                name         string
                geom         Geom
                wantDistance float64
                wantErr      bool
        }{
                {
                        name:         "#1",
                        geom:         Geom{X1: 1, Y1: 1, X2: 4, Y2: 5},
                        wantDistance: 5,
                        wantErr:      false,
                },
                {
                        name:         "Negative Coordinates",
                        geom:         Geom{X1: -1, Y1: 1, X2: 4, Y2: 5},
                        wantDistance: -1,
                        wantErr:      true,
                },
                {
                        name:         "All Negative Coordinates",
                        geom:         Geom{X1: -1, Y1: -1, X2: -4, Y2: -5},
                        wantDistance: -1,
                        wantErr:      true,
                },
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        gotDistance, err := tt.geom.CalculateDistance()
                        if (err != nil) != tt.wantErr {
                                t.Errorf("Geom.CalculateDistance() error = %v, wantErr %v", err, tt.wantErr)
                                return
                        }
                        if math.Abs(gotDistance-tt.wantDistance) > 1e-9 { // Используем math.Abs для сравнения float64
                                t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
                        }
                })
        }
}