package temperature

import "math"

func ExtractTemperature(v *[]int64) int {
	if len(*v) > 0 {
		var minDifference float64 = 5526
        for _, temperature := range *v {
			// t: a temperature expressed as an integer ranging from -273 to 5526
			tFloat := float64(temperature)
            if  math.Abs(tFloat) < math.Abs(minDifference) {
				minDifference = tFloat
			} else if math.Abs(tFloat) == math.Abs(minDifference) && int64(minDifference) + temperature == 0 { 
				minDifference = math.Abs(tFloat)
			} 
        }
        return int(minDifference)
	} else {
		return 0
	}
}