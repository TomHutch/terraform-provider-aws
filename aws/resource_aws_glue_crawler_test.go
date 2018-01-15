package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"fmt"
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
