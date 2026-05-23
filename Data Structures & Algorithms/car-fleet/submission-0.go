func carFleet(target int, position []int, speed []int) int {
	n := len(position)

	type car struct {
		post int
		speed int
	}
	cars := []car{}
	for i:= 0 ; i < n ; i++{
		cars = append(cars, car{
			post : position[i],
			speed : speed[i],
		})
	}
	sort.Slice(cars, func(i,j int) bool {
		return cars[i].post > cars[j].post
	})

	stack := []float64{}

	for _, car := range cars {

		time := float64(target- car.post)/float64(car.speed)

		if len(stack) == 0 || time > stack[len(stack)-1] {
			stack = append(stack, time)
		}
	}

	return len(stack)
}
