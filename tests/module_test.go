package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/hello_world",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "result")

	assert.Equal(t, "Hello, World!", output)
}

func TestHelloCustom(t *testing.T) {

	var name string = "DontShaveTheYak"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/hello_custom",
		Vars: map[string]interface{}{
			"name": name,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "result")

	expectedOutput := fmt.Sprintf("Hello, %s!", name)

	assert.Equal(t, expectedOutput, output)
}

func TestHelloMultiple(t *testing.T) {

	var names = []interface{}{"DSTY", "TFWeekly", "TerraTest"}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/hello_multiple",
		Vars: map[string]interface{}{
			"friends": names,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "result")

	expectedOutput := fmt.Sprintf("Hello, %s, %s and %s!", names...)

	assert.Equal(t, expectedOutput, output)
}
