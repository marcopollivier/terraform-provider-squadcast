package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/squadcast/terraform-provider-squadcast/internal/api"
	"github.com/squadcast/terraform-provider-squadcast/internal/tf"
	"time"
)

func resourceOnCall() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceOnCallCreate,

		Schema: map[string]*schema.Schema{
			"schedule_id": {
				Description:  "Schedule id.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: tf.ValidateObjectID,
				ForceNew:     true,
			},
			"name": {
				Description:  "On-call event name.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"start_time": {
				Description:  "Date-time to start the rotation.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},
			"end_time": {
				Description:  "Date-time to finish the rotation.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},
			"timezone": {
				Description:  "Timezone target to schedule the rotation.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"is_forever": {
				Description: "Define if the rotation there's no end.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"rotate": {
				Description: "Define if the list of users/squads will rotate.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"rotation_frequency": {
				Description: "Frequency of the rotation.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"repeat": {
				Description: "Define if the rotation will repeat after the end.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"repeat_frequency": {
				Description: "",
				Type:        schema.TypeInt,
				Computed:    true,
			},

			"repetition_type": {
				Description: "",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_ids": {
				Description: "",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"squad_ids": {
				Description: "",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceToApi(d *schema.ResourceData) *api.OnCall {

	//rotationSets := &api.RotationSets{
	//	UserIds:  d.Get("user_ids").([]string),
	//	SquadIds: d.Get("squad_ids").([]string),
	//}

	timezone := &api.Timezone{}

	repetition := &api.Repetition{
		Frequency: 0,
		Type:      "",
	}

	config := &api.Config{
		IsForever:         d.Get("is_forever").(bool),
		Rotate:            d.Get("rotate").(bool),
		RotationFrequency: 0,
		Repeat:            false,
		Timezone:          *timezone,
		Repetition:        *repetition,
	}

	return &api.OnCall{
		Name:      d.Get("name").(string),
		StartTime: d.Get("start_time").(time.Time),
		EndTime:   d.Get("end_time").(time.Time),
		Config:    *config,
	}
}

func resourceOnCallCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*api.Client)

	tflog.Info(ctx, "Creating on-call events", tf.M{
		"name": d.Get("name").(string),
	})
	scheduleId := d.Get("schedule_id").(string)

	_, err := client.CreateOnCall(ctx, resourceToApi(d), scheduleId)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
