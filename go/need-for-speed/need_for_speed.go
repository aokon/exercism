package speed

// TODO: define the 'Car' type struct
type Car struct {
	speed        int
	battery      int
	batteryDrain int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed,
		100,
		batteryDrain,
		0,
	}
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	battery := car.battery - car.batteryDrain
	distance := car.distance + car.speed

	if battery < 0 {
		battery = car.battery
		distance = car.distance
	}

	return Car{
		car.speed,
		battery,
		car.batteryDrain,
		distance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	lapsNumber := track.distance / car.speed
	batteryCost := car.batteryDrain * lapsNumber

	var result bool

	if (car.battery - batteryCost) >= 0 {
		result = true
	} else {
		result = false
	}

	return result
}
