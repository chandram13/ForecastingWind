// Marvish Chandra

package CalculateWindLoad

import "fmt"

func windForce(l,w,v,cD):
projectedArea = l * w
windPressure = 0.613 * v * v
windLoad = projectedArea * windPressure * cD

