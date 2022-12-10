package encode

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/martinlindhe/base36"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBase36Encoder() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBase36EncoderRead,
		Schema: map[string]*schema.Schema{
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value to base36 encode",
			},
			"lowercase": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Compute a lower-case result",
			},
			"result": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The result of the base36 encoding operation",
			},
		},
	}
}

func dataSourceBase36EncoderRead(ctx context.Context, rd *schema.ResourceData, v interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	value := rd.Get("value").(string)
	lowercase := rd.Get("lowercase").(bool)
	id := generateID(value)
	result := base36.EncodeBytes([]byte(value))
	if lowercase {
		result = strings.ToLower(result)
	}
	if err := rd.Set("result", result); err != nil {
		return diag.FromErr(err)
	}
	rd.SetId(id)
	return diags
}

func generateID(value string) string {
	h := sha256.Sum256([]byte(value))
	return hex.EncodeToString(h[:])
}
