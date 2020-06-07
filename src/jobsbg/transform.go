package jobsbg

import (
	"go-job-listing-aggregator/src/listing"
)

// SourceName defines the source as jobsbg
const SourceName = "jobsbg"

func transformToListingModel(queryID uint64, jobsBGListing *jobsBGListing) listing.Listing {
	return listing.Listing{
		ExternalID:  jobsBGListing.ID,
		Link:        jobsBGListing.Link,
		Name:        jobsBGListing.Title,
		Company:     jobsBGListing.Company,
		Location:    jobsBGListing.Location,
		PostingDate: jobsBGListing.PublishingDate,
		Description: jobsBGListing.Description,
		Keywords:    jobsBGListing.Keywords,
		QueryID:     queryID,
		SourceName:  SourceName}
}
