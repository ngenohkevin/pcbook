package sample

import (
	"github.com/google/uuid"
	"github.com/ngenohkevin/pcbook/pb"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}
func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9 9980HK",
			"Core i7 9750H",
			"Core i5-9400F",
			"Core i3 1005G1",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}
func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}
func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "HP", "Lenovo")
}
func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "HP":
		return randomStringFromSet("Notebook", "Pro Desk", "Helios")
	case "Dell":
		return randomStringFromSet("Latitude", "XPS", "Vostro", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad p53")
	}
}
func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX VEGA-56",
	)
}
func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	Resolution := &pb.Screen_Resolution{
		Width:  uint32(height),
		Height: uint32(width),
	}
	return Resolution
}
func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
func randomID() string {
	return uuid.New().String()
}
