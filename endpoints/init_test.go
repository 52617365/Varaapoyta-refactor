package endpoints

import "testing"

func TestCityIsValid(t *testing.T) {
	invalidCity := "invalidCity"
	if cityIsValid(invalidCity) {
		t.Errorf("City %s should not be valid", invalidCity)
	}
}
func TestCityIsOnRaflaamoList(t *testing.T) {
	cityThatIsNotOnRaflaamoList := "invalidCity"
	if cityIsOnRaflaamoList(cityThatIsNotOnRaflaamoList) {
		t.Errorf("City: %s should not be in the list", cityThatIsNotOnRaflaamoList)
	}
}