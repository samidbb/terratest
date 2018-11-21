// Integration tests that validate S3-related code in AWS.
package aws

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
)

func TestZoneIsFound(t *testing.T) {
	t.Parallel()

	zoneId := "ZUZ8AC4JZBPK5"
	region := aws.GetRandomRegion(t, nil, nil)
	logger.Logf(t, "Selected random region: %s", region)
	logger.Logf(t, "Looking up Zone with id: %s", zoneId)
	zoneName := FindHostedZoneWithId(t, region, zoneId)
	logger.Logf(t, "Found zone with name %s", zoneName)
	expectedZoneName := "ded.onep.dk."
	assert.Equal(t, expectedZoneName, zoneName)

}