package story

import (
	"fmt"
	"strings"
)

type Interpolator struct{}

func (i Interpolator) Interpolate(template string, words []string) string {
	for i := 0; i < len(words); i++ {
		template = strings.Replace(template, fmt.Sprintf("&%d", i+1), words[i], 1)
	}

	return template
}
