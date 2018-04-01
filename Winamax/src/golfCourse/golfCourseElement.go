package golfCourse 

type GolfElement interface {
	X() int
	Y() int
}

// func FindAll(golfCourseMap []string) ([](*GolfElement)) {
// 	var golfElement [](*GolfElement) 
// 	H := len(golfCourseMap)
// 	for i := 0; i < H; i++ {
// 		for j := 0; j < len(golfCourseMap[i]); j++ {
// 			if isPointABall(string(golfCourseMap[i][j])) {
// 				golfElement = append(golfElement, golfElementSupplier(i, j))
// 			}
// 		}
// 	}
// 	return golfElement
// }