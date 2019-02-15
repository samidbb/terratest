package aws

import (
	"testing"
	amazon "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	tt "github.com/gruntwork-io/terratest/modules/aws"
)

func GetParameter(t *testing.T, awsRegion string, keyName string) string {
	keyValue, err := GetParameterE(t, awsRegion, keyName)
	if err != nil {
		t.Fatal(err)
	}
	return keyValue
}

func GetParameterE(t *testing.T, awsRegion string, keyName string) (string, error) {
	ssmClient, err := NewSsmClientE(t, awsRegion)
	if err != nil {
		return "", err
	}

	resp, err := ssmClient.GetParameter(&ssm.GetParameterInput{Name: amazon.String(keyName), WithDecryption: amazon.Bool(true)})
	if err != nil {
		return "", err
	}

	parameter := *resp.Parameter
	return *parameter.Value, nil
}

func PutParameter(t *testing.T, awsRegion string, keyName string, keyDescription string, keyValue string) int64 {
	version, err := PutParameterE(t, awsRegion, keyName, keyDescription, keyValue)
	if err != nil {
		t.Fatal(err)
	}
	return version
}

func PutParameterE(t *testing.T, awsRegion string, keyName string, keyDescription string, keyValue string, ) (int64, error) {
	ssmClient, err := NewSsmClientE(t, awsRegion)
	if err != nil {
		return 0, err
	}

	resp, err := ssmClient.PutParameter(&ssm.PutParameterInput{Name: amazon.String(keyName), Description: amazon.String(keyDescription), Value: amazon.String(keyValue), Type: amazon.String("SecureString")})
	if err != nil {
		return 0, err
	}

	return *resp.Version, nil
}

// NewSsmClient creates a ssm client.
func NewSsmClient(t *testing.T, region string) *ssm.SSM {
	client, err := NewSsmClientE(t, region)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

// NewSsmClientE creates an ssm client.
func NewSsmClientE(t *testing.T, region string) (*ssm.SSM, error) {
	sess, err := tt.NewAuthenticatedSession(region)
	if err != nil {
		return nil, err
	}

	return ssm.New(sess), nil
}