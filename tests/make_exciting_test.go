package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMakeExcitingRequiresInput(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/make_exciting",
	})

	_, err := terraform.InitAndApplyE(t, terraformOptions)

	assert.Error(t, err)
}

func TestMakeExciting(t *testing.T) {

	var name string = "DontShaveTheYak"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/make_exciting",
		Vars: map[string]interface{}{
			"text": name,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "result")

	expectedOutput := fmt.Sprintf("%s!", name)

	assert.Equal(t, expectedOutput, output)
}
