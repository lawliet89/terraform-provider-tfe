package tfe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTFEOrganizationRunTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTFEOrganizationRunTaskRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"organization": {
				Type:     schema.TypeString,
				Required: true,
			},

			"url": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceTFEOrganizationRunTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(ConfiguredClient)
	name := d.Get("name").(string)
	organization := d.Get("organization").(string)

	task, err := fetchOrganizationRunTask(name, organization, config.Client)
	if err != nil {
		return err
	}

	d.Set("url", task.URL)
	d.Set("category", task.Category)
	d.Set("enabled", task.Enabled)
	d.Set("description", task.Description)
	d.SetId(task.ID)

	return nil
}
