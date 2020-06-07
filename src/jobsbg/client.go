package jobsbg

import (
	"fmt"
	"strings"
)

const searchPageURL = "https://www.jobs.bg/front_job_search.php?add_sh=1&location_sid=&distance=&categories[]=0&type_all=0&position_level_all=0&company_type[]=0&::keywords::keyword=&job_languages_all=0&salary_from=0&last=0"

func generateKeywordsQueryParameter(keywords []string) string {
	var builder strings.Builder

	for _, word := range keywords {
		parameterRepresentation := fmt.Sprintf("keywords[]=%s&", word)
		builder.WriteString(parameterRepresentation)
	}

	return builder.String()
}
