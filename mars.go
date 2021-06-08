// How long does it take to get to Mars?
package main

func main() {
	//最原始的形式,但是程序会报错
	var distance = 56000000
	var speed = 100800

	var (
		distance = 56000000
		speed    = 100800
	)

	//使用逗号分开
	var distance, speed = 56000000, 100800

	const hoursPerDay, minutesPerHour = 24, 60

}
