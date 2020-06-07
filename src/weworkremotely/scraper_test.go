package weworkremotely

import "testing"

func TestFilterNonRelevantPositionListings(t *testing.T) {
	keywords := []string{"java"}

	validTitleModel := jobListingRSSModel{
		Title:       "Java",
		Description: "Groovy"}

	validDescriptionModel := jobListingRSSModel{
		Title:       "Groovy",
		Description: "Java"}

	invalidModel := jobListingRSSModel{
		Title:       "Kotlin",
		Description: "Groovy"}

	models := []jobListingRSSModel{
		validTitleModel, validDescriptionModel, invalidModel}

	expectedModels := []jobListingRSSModel{
		validTitleModel, validDescriptionModel}

	actualModels := filterNonRelevantPositionListings(keywords, models)

	if (expectedModels == nil) != (actualModels == nil) {
		t.Errorf("%s failed: Expected 2 identical slices, got %+v",
			t.Name(), actualModels)
	}

	if len(expectedModels) != len(actualModels) {
		t.Errorf("%s failed: Expected slice size %d, got %d",
			t.Name(), len(expectedModels), len(actualModels))
	}

	for index, expectedModel := range expectedModels {
		actualModel := actualModels[index]

		if expectedModel.Title != actualModel.Title {
			t.Errorf("%s failed: Expected title %s, got %s",
				t.Name(), expectedModel.Title, actualModel.Title)
		}

		if expectedModel.Description != actualModel.Description {
			t.Errorf("%s failed: Expected description %s, got %s",
				t.Name(), expectedModel.Description, actualModel.Description)
		}
	}
}
