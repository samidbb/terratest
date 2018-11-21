package aws

import (
	"testing"
	amazon "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	tt "github.com/gruntwork-io/terratest/modules/aws"
)

func FindHostedZoneWithId(t *testing.T, awsRegion string, zoneId string) string {
	zoneName, err := FindHostedZoneWithIdE(t, awsRegion, zoneId)
	if err != nil {
		t.Fatal(err)
	}
	return zoneName
}

func FindHostedZoneWithIdE(t *testing.T, awsRegion string, zoneId string) (string, error) {
	route53Client, err := NewRoute53ClientE(t, awsRegion)
	if err != nil {
		return "", err
	}

	resp, err := route53Client.GetHostedZone(&route53.GetHostedZoneInput{Id: amazon.String(zoneId)})
	if err != nil {
		return "", err
	}

	zone := *resp.HostedZone
	return *zone.Name, nil
}

// NewRoute53Client creates a Route53 client.
func NewRoute53Client(t *testing.T, region string) *route53.Route53 {
	client, err := NewRoute53ClientE(t, region)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

// NewRoute53ClientE creates an Route53 client.
func NewRoute53ClientE(t *testing.T, region string) (*route53.Route53, error) {
	sess, err := tt.NewAuthenticatedSession(region)
	if err != nil {
		return nil, err
	}

	return route53.New(sess), nil
}