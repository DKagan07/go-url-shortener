package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	userId = "randomUserId"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink1 := GenerateShortLink(initialLink1, userId)

	assert.Equal(t, "A2AZKReM", shortLink1)
}
