package repository

import (
	"testing"

	"github.com/containous/lobicornis/v2/pkg/conf"
	"github.com/stretchr/testify/assert"
)

func TestRepository_getMinReview(t *testing.T) {
	testCases := []struct {
		name              string
		config            conf.RepoConfig
		markers           conf.Markers
		labels            []string
		expectedMinReview int
	}{
		{
			name: "with light review label",
			config: conf.RepoConfig{
				MinReview:      conf.Int(3),
				MinLightReview: conf.Int(1),
			},
			markers: conf.Markers{
				LightReview: "bot/light-review",
			},
			labels:            []string{"bot/light-review"},
			expectedMinReview: 1,
		},
		{
			name: "without light review label",
			config: conf.RepoConfig{
				MinReview:      conf.Int(3),
				MinLightReview: conf.Int(1),
			},
			markers: conf.Markers{
				LightReview: "bot/light-review",
			},
			expectedMinReview: 3,
		},
	}

	for i, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repository := Repository{
				markers: test.markers,
				config:  test.config,
			}

			pr := makePullRequestWithLabels(test.labels, i)

			minReview := repository.getMinReview(pr)

			assert.Equal(t, test.expectedMinReview, minReview)
		})
	}
}
