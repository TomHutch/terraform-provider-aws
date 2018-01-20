package aws

import "github.com/hashicorp/terraform/helper/schema"

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
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
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
		},
	}
}

func resourceAwsGlueCrawlerCreate(d *schema.ResourceData, meta interface{}) (err error) {
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
