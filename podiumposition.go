package piscine

func PodiumPosition(podium [][]string) [][]string {
	depart := 0
	final := len(podium) - 1
	for depart < final {
		podium[final], podium[depart] = podium[depart], podium[final]
		depart++
		final--
	}
	return podium
}
