package algs

import (
	"testing"

	"github.com/mjevans93308/platformscience/pkg/models"
)

type DummyFileContent struct {
	names     []string
	addresses []string
}

type CalculateSSCase struct {
	mockContent models.FileContent
	expected    []models.NameAddressSS
}

func TestCalculateSS(t *testing.T) {
	testCase := CalculateSSCase{
		mockContent: models.FileContent{
			Names: []string{
				"Daniel Davidson",
				"Isabella Lewis",
				"Joseph White",
			},
			Addresses: []string{
				"44 Fake Dr., San Diego, CA 92122",
				"765 Sycamore Avenue, Dallas, TX, 75201",
				"109 Pineapple Street, Philadelphia, PA, 19101",
			},
		},
		expected: []models.NameAddressSS{
			{
				Name:    "Isabella Lewis",
				Address: "44 Fake Dr., San Diego, CA 92122",
				SS:      13.5,
			},
			{
				Name:    "Daniel Davidson",
				Address: "109 Pineapple Street, Philadelphia, PA, 19101",
				SS:      12,
			},
			{
				Name:    "Joseph White",
				Address: "765 Sycamore Avenue, Dallas, TX, 75201",
				SS:      9,
			},
		},
	}
	result := CalculateSS(nil, &testCase.mockContent)
	for i := range result {
		checkName := result[i].Name == testCase.expected[i].Name
		checkAddr := result[i].Address == testCase.expected[i].Address
		checkSS := result[i].SS == testCase.expected[i].SS

		if !checkName || !checkAddr || !checkSS {
			t.Errorf("For CalculateSS, got name: %s, wanted %s, got address: %s, wanted %s, got SS: %f, wanted %f", result[i].Name, testCase.expected[i].Name, result[i].Address, testCase.expected[i].Address, result[i].SS, testCase.expected[i].SS)
		}
	}
}
