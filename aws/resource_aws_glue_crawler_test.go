package aws

import (
	"testing"

	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSGlueCrawler_basic(t *testing.T) {
	rInt := acctest.RandInt()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGlueCrawlerDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccGlueCralwer_basic(rInt),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGlueCrawlerExists("aws_glue_crawler.test"),
					resource.TestCheckResourceAttr(
						"aws_glue_crawler.test",
						"name",
						fmt.Sprintf("my_test_crawler_%d", rInt),
					),
					resource.TestCheckResourceAttr(
						"aws_glue_crawler.test",
						"role",
						"arn:aws:iam::123456789012:role/GlueCrawlerAccess",
					),
					resource.TestCheckResourceAttr(
						"aws_glue_crawler.test",
						"database_name",
						"my_test_db",
					),
					resource.TestCheckResourceAttr(
						"aws_glue_crawler.test",
						"description",
						"A test crawler from terraform",
					),
					resource.TestCheckResourceAttr(
						"aws_glue_crawler.test",
						"s3_targets.0.path",
						"s3://my-test-bucet/path/to/data",
					),
					resource.TestCheckResourceAttr(
						"aws_glue_crawler.test",
						"s3_targets.0.exclusions.0",
						"my_glob_pattern",
					),
				),
			},
		},
	})
}

func testAccCheckGlueCrawlerDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).glueconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_glue_crawler" {
			continue
		}

		input := &glue.GetCrawlerInput{
			Name: aws.String(rs.Primary.ID),
		}
		if _, err := conn.GetCrawler(input); err != nil {
			//Verify the error is what we want
			if isAWSErr(err, glue.ErrCodeEntityNotFoundException, "") {
				continue
			}

			return err
		}
		return fmt.Errorf("still exists")
	}
	return nil
}

func testAccGlueCralwer_basic(rInt int) string {
	return fmt.Sprintf(`
resource "aws_glue_crawler" "test" {
  name          = "my_test_crawler_%s"
  role          = "arn:aws:iam::123456789012:role/GlueCrawlerAccess"
  database_name = "my_test_db"
}
`, rInt)
}

func testAccCheckGlueCrawlerExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		glueconn := testAccProvider.Meta().(*AWSClient).glueconn
		out, err := glueconn.GetCrawler(&glue.GetCrawlerInput{
			Name: aws.String(rs.Primary.ID),
		})

		if err != nil {
			return err
		}

		if out.Crawler == nil {
			return fmt.Errorf("No Glue Crawler Found")
		}

		if *out.Crawler.Name != rs.Primary.ID {
			return fmt.Errorf("Glue Crawler Mismatch - existing: %q, state: %q",
				*out.Crawler.Name, rs.Primary.ID)
		}

		return nil
	}
}
