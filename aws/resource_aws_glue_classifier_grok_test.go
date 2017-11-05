package aws

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/glue"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSGlueClassifierGrok_basic(t *testing.T) {
	rInt := acctest.RandInt()
	name := "aws_glue_classifier_grok.test"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGlueClassifierGrokDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGlueClassifierGrok_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGlueClassifierGrokExists(name),
					resource.TestCheckResourceAttr(
						name,
						"classification",
						"my_classification",
					),
					resource.TestCheckResourceAttr(
						name,
						"grok_pattern",
						"my_grok_pattern",
					),
					resource.TestCheckResourceAttr(
						name,
						"custom_patterns",
						"my_custom_pattern",
					),
					resource.TestMatchResourceAttr(
						name,
						"creation_time",
						regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\+\\d{4}$"),
					),
					resource.TestMatchResourceAttr(
						name,
						"last_updated",
						regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\+\\d{4}$"),
					),
					resource.TestMatchResourceAttr(
						name,
						"version",
						regexp.MustCompile("^\\d+$"),
					),
				),
			},
		},
	})
}

func testAccCheckGlueClassifierGrokDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).glueconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_glue_classifier_grok" {
			continue
		}

		input := &glue.GetClassifierInput{
			Name: aws.String(rs.Primary.ID),
		}
		if _, err := conn.GetClassifier(input); err != nil {
			//Verify the error is what we want
			if ae, ok := err.(awserr.Error); ok && ae.Code() == "EntityNotFoundException" {
				continue
			}

			return err
		}
		return fmt.Errorf("still exists")
	}
	return nil
}

func testAccGlueClassifierGrok_basic(rInt int) string {
	return fmt.Sprintf(`
resource "aws_glue_classifier_grok" "test" {
  name            = "test_classifier_%d"
  classification  = "my_classification"
  grok_pattern    = "my_grok_pattern"
  custom_patterns = "my_custom_pattern"
}
`, rInt)
}

func testAccCheckGlueClassifierGrokExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		glueconn := testAccProvider.Meta().(*AWSClient).glueconn
		out, err := glueconn.GetClassifier(&glue.GetClassifierInput{
			Name: aws.String(rs.Primary.ID),
		})

		if err != nil {
			return err
		}

		if out.Classifier == nil {
			return fmt.Errorf("No Glue Grok Classifier Found")
		}

		if *out.Classifier.GrokClassifier.Name != rs.Primary.ID {
			return fmt.Errorf("Glue Grok Classifier Mismatch - existing: %q, state: %q",
				*out.Classifier.GrokClassifier.Name, rs.Primary.ID)
		}

		return nil
	}
}
