package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSayHelloRequiresInput(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/say_hello",
		NoColor:      true,
	})

	_, err := terraform.InitAndApplyE(t, terraformOptions)

	assert.ErrorContains(t, err, `"name" is not set, and has no default value`)

}

func TestSayHello(t *testing.T) {

	var name string = "DontShaveTheYak"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/say_hello",
		Vars: map[string]interface{}{
			"name": name,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "result")

	expectedOutput := fmt.Sprintf("Hello, %s", name)

	assert.Equal(t, expectedOutput, output)
}
