package aws

import (
	"fmt"
	"strconv"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glue"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsGlueClassifierGrok() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsGlueClassifierGrokCreate,
		Read:   resourceAwsGlueClassifierGrokRead,
		Update: resourceAwsGlueClassifierGrokUpdate,
		Delete: resourceAwsGlueClassifierGrokDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"classification": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"grok_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_patterns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"creation_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAwsGlueClassifierGrokCreate(d *schema.ResourceData, meta interface{}) error {
	glueconn := meta.(*AWSClient).glueconn
	name := d.Get("name").(string)

	input := &glue.CreateClassifierInput{
		GrokClassifier: &glue.CreateGrokClassifierRequest{
			Name:           aws.String(name),
			Classification: aws.String(d.Get("classification").(string)),
			GrokPattern:    aws.String(d.Get("grok_pattern").(string)),
		},
	}

	_, err := glueconn.CreateClassifier(input)
	if err != nil {
		return fmt.Errorf("Error creating Grok Classifier: %s", err)
	}

	d.SetId(name)

	return resourceAwsGlueClassifierGrokUpdate(d, meta)
}

func resourceAwsGlueClassifierGrokUpdate(d *schema.ResourceData, meta interface{}) error {
	glueconn := meta.(*AWSClient).glueconn
	doUpdate := false
	input := &glue.UpdateClassifierInput{
		GrokClassifier: &glue.UpdateGrokClassifierRequest{
			Name: aws.String(d.Id()),
		},
	}

	if ok := d.HasChange("classification"); ok {
		doUpdate = true
		input.GrokClassifier.Classification = aws.String(
			d.Get("classification").(string),
		)
	}

	if ok := d.HasChange("grok_pattern"); ok {
		doUpdate = true
		input.GrokClassifier.GrokPattern = aws.String(
			d.Get("grok_pattern").(string),
		)
	}

	if ok := d.HasChange("custom_patterns"); ok {
		doUpdate = true
		input.GrokClassifier.CustomPatterns = aws.String(
			d.Get("custom_patterns").(string),
		)
	}

	if doUpdate {
		if _, err := glueconn.UpdateClassifier(input); err != nil {
			return err
		}
	}

	return resourceAwsGlueClassifierGrokRead(d, meta)
}

func resourceAwsGlueClassifierGrokRead(d *schema.ResourceData, meta interface{}) error {
	glueconn := meta.(*AWSClient).glueconn

	input := &glue.GetClassifierInput{
		Name: aws.String(d.Id()),
	}

	out, err := glueconn.GetClassifier(input)
	if err != nil {
		return fmt.Errorf("Error reading Glue Grok Classifier: %s", err.Error())
	}

	d.Set("name", d.Id())
	gClassifier := out.Classifier.GrokClassifier
	d.Set("classification", gClassifier.Classification)
	d.Set("grok_pattern", gClassifier.GrokPattern)
	d.Set("custom_patterns", gClassifier.CustomPatterns)

	creation := gClassifier.CreationTime.UTC().Format("2006-01-02T15:04:05-0700")
	updated := gClassifier.LastUpdated.UTC().Format("2006-01-02T15:04:05-0700")

	d.Set("creation_time", creation)
	d.Set("last_updated", updated)
	d.Set("version", strconv.FormatInt(*gClassifier.Version, 10))

	return nil
}

func resourceAwsGlueClassifierGrokDelete(d *schema.ResourceData, meta interface{}) error {
	glueconn := meta.(*AWSClient).glueconn

	log.Printf("[DEBUG] Glue Grok Classifier: %s", d.Id())
	_, err := glueconn.DeleteClassifier(&glue.DeleteClassifierInput{
		Name: aws.String(d.Id()),
	})
	if err != nil {
		return fmt.Errorf("Error deleting Glue Grok Classifier: %s", err.Error())
	}
	return nil
}
