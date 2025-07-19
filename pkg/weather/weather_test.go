package weather

import (
	"fmt"
	"testing"
)

func TestGetWeather(t *testing.T) {
	tests := []struct {
		city string
		want string
	}{
		{"Beijing", "sunny"},
	}

	for i := 0; i < 100; i++ {
		got := GetWeather(fmt.Sprintf("city%d", i))
		if got == tests[i%len(tests)].want {
			break
		}
	}
}
