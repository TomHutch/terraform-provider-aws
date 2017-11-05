package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAWSGlueClassifierGrok_importBasic(t *testing.T) {
	resourceName := "aws_glue_classifier_grok.test"
	rInt := acctest.RandInt()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGlueClassifierGrokDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGlueClassifierGrok_basic(rInt),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
