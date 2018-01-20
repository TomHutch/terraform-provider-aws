package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsGlueCrawler() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsGlueCrawlerCreate,
		Read:   resourceAwsGlueCrawlerRead,
		Update: resourceAwsGlueCrawlerUpdate,
		Delete: resourceAwsGlueCrawlerDelete,
		Exists: resourceAwsGlueCrawlerExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
				ValidateFunc: awsGlueSingleLineString,
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"s3_targets": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"jdbc_targets"},
				MinItems:      1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"exclusions": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
			"jdbc_targets": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"s3_targets"},
				MinItems:      1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"exclusions": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func resourceAwsGlueCrawlerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	glueconn := meta.(*AWSClient).glueconn
	name := d.Get("name").(string)
	role := d.Get("role").(string)
	database_name := d.Get("database_name").(string)

	input := &glue.CreateCrawlerInput{
		Name:         aws.String(name),
		Role:         aws.String(role),
		DatabaseName: aws.String(database_name),
		Targets:      &glue.CrawlerTargets{},
	}
	_, err = glueconn.CreateCrawler(input)

	return
}

func resourceAwsGlueCrawlerRead(d *schema.ResourceData, meta interface{}) (err error) {
	return
}

func resourceAwsGlueCrawlerUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	return
}

func resourceAwsGlueCrawlerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	return
}

func resourceAwsGlueCrawlerExists(d *schema.ResourceData, meta interface{}) (exists bool, err error) {
	return
}

func awsGlueSingleLineString(d interface{}, v string) (s []string, e []error) {
	return
}
