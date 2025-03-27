package main

import (
	"fmt"
	"strings"
)

// ShadeCategory shade ke seasonal category ko define karta hai
type ShadeCategory struct {
	Name     string
	Seasonal map[string]bool // Kis season mein appropriate hai
}

// shadeDatabase ye lipstick shade database ko paribhasit karta hai jo skin tone aur age group par aadharit hai.
var shadeDatabase = map[string]map[string][]ShadeCategory{
	"Fair": {
		"Young": {
			{Name: "Soft Pink", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
			{Name: "Peach Coral", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
			{Name: "Glossy Nude", Seasonal: map[string]bool{"Winter": true, "Summer": true, "Spring": true, "Fall": true}},
			{Name: "Light Mauve", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": true, "Fall": true}},
		},
		"Middle-Aged": {
			{Name: "Warm Rose", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": true, "Fall": true}},
			{Name: "Mocha Brown", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Classic Red", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Dusty Pink", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
		},
		"Older": {
			{Name: "Deep Mauve", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Rich Plum", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Elegant Wine", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Soft Coral", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
		},
	},
	"Medium": {
		"Young": {
			{Name: "Berry Red", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": true, "Fall": true}},
			{Name: "Rosy Mauve", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
			{Name: "Glossy Nude", Seasonal: map[string]bool{"Winter": true, "Summer": true, "Spring": true, "Fall": true}},
			{Name: "Coral Pink", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
		},
		"Middle-Aged": {
			{Name: "Caramel Brown", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Spiced Coral", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
			{Name: "Deep Red", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Warm Nude", Seasonal: map[string]bool{"Winter": true, "Summer": true, "Spring": true, "Fall": true}},
		},
		"Older": {
			{Name: "Terracotta", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": false, "Fall": true}},
			{Name: "Burgundy", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Dark Mocha", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Brick Red", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
		},
	},
	"Dark": {
		"Young": {
			{Name: "Vibrant Plum", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": true, "Fall": true}},
			{Name: "Fiery Orange", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
			{Name: "Glossy Cocoa", Seasonal: map[string]bool{"Winter": true, "Summer": true, "Spring": false, "Fall": true}},
			{Name: "Hot Pink", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
		},
		"Middle-Aged": {
			{Name: "Chocolate Brown", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Bold Red", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Deep Burgundy", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Warm Coral", Seasonal: map[string]bool{"Winter": false, "Summer": true, "Spring": true, "Fall": false}},
		},
		"Older": {
			{Name: "Crimson Red", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Dark Violet", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Espresso Brown", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": false, "Fall": true}},
			{Name: "Deep Berry", Seasonal: map[string]bool{"Winter": true, "Summer": false, "Spring": true, "Fall": true}},
		},
	},
}

// recommendShades lipstick shades ka sujhav karta hai jo skin tone, age group, season, aur finish pasand par aadharit hota hai.
func recommendShades(skinTone, ageGroup, season, finishPreference string) []string {
	// Skin tone aur age group ki maanita ki jaanch karein
	if !isValidSkinTone(skinTone) || !isValidAgeGroup(ageGroup) {
		return []string{"No match found"}
	}

	// Shades ko prapt karein
	shadeCategories := shadeDatabase[skinTone][ageGroup]
	
	// Shades ke naam extract karein
	shadeNames := make([]string, len(shadeCategories))
	for i, category := range shadeCategories {
		shadeNames[i] = category.Name
	}
	
	// Mausam ke anusar shades ko adjust karein
	seasonalShades := applySeasonalAdjustment(shadeNames, shadeCategories, season)
	
	// Finish preference lagu karein
	finalShades := applyFinishPreference(seasonalShades, finishPreference)
	
	if len(finalShades) == 0 {
		return []string{"No suitable shades found for this combination"}
	}
	
	return finalShades
}

// isValidSkinTone dekhta hai ki prastavit skin tone maanaya hai ya nahi.
func isValidSkinTone(skinTone string) bool {
	_, exists := shadeDatabase[skinTone]
	return exists
}

// isValidAgeGroup dekhta hai ki prastavit age group maanaya hai ya nahi.
func isValidAgeGroup(ageGroup string) bool {
	validAgeGroups := map[string]bool{
		"Young":       true,
		"Middle-Aged": true,
		"Older":       true,
	}
	return validAgeGroups[ageGroup]
}

// applySeasonalAdjustment mausam ke aadharit sujhawit shades ko badalta hai.
// Improved version that uses the seasonal flags to filter shades
func applySeasonalAdjustment(shadeNames []string, shadeCategories []ShadeCategory, season string) []string {
	// Mausam ke anusar filter karein
	var seasonalShades []string
	
	for i, category := range shadeCategories {
		// Agar is shade ka current season mein true hai to add karein
		if category.Seasonal[season] {
			seasonalShades = append(seasonalShades, shadeNames[i])
		}
	}
	
	// Agar seasonal shades nahi mile to original shades lautayein
	if len(seasonalShades) == 0 {
		switch season {
		case "Winter":
			// Winter mein pehle 2 shades
			return shadeNames[:min(2, len(shadeNames))]
		case "Summer":
			// Summer mein madhya 2 shades
			return shadeNames[1:min(3, len(shadeNames))]
		default:
			// Anya mausam ke liye original shades
			return shadeNames
		}
	}
	
	return seasonalShades
}

// applyFinishPreference finish pasand ke aadharit sujhawit shades ko badal deta hai.
func applyFinishPreference(shades []string, finish string) []string {
	for i, shade := range shades {
		if finish == "Matte" {
			shades[i] = shade + " (Matte)"
		} else if finish == "Glossy" {
			shades[i] = shade + " (Glossy)"
		} else if finish == "Satin" {
			shades[i] = shade + " (Satin)"
		}
	}
	return shades
}

// min do poornankon mein se chhoti sankhya ko lautata hai.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Parikshan cases jo system ki kareksheelta dikhate hain
	testCases := [][]string{
		{"Fair", "Young", "Winter", "Matte"},
		{"Medium", "Middle-Aged", "Summer", "Glossy"},
		{"Dark", "Older", "Winter", "Matte"},
		{"Fair", "Middle-Aged", "Spring", "Satin"},
		{"Unknown", "Young", "Summer", "Matte"},
	}

	// Har parikshan case ko execute karein aur sujhaw dekhein
	for _, testCase := range testCases {
		skinTone, ageGroup, season, finishPreference := testCase[0], testCase[1], testCase[2], testCase[3]
		
		fmt.Printf("Skin Tone: %s, Age: %s, Season: %s, Finish: %s\n",
			skinTone, ageGroup, season, finishPreference)
		
		recommendations := recommendShades(skinTone, ageGroup, season, finishPreference)
		fmt.Println("Recommended Shades:", strings.Join(recommendations, ", "))
		fmt.Println("--------------------------")
	}
}