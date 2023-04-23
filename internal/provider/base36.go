package provider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/martinlindhe/base36"
	"strings"
)

func NewBase36EncoderDataSource() datasource.DataSource {
	return &EncodeBase36DataSource{}
}

var _ datasource.DataSource = &EncodeBase36DataSource{}

type EncodeBase36DataSource struct {
}

func (d *EncodeBase36DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_base36"
}

func (d *EncodeBase36DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "data source identifier",
			},
			"value": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The value to base36 encode",
			},
			"lowercase": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Compute a lower-case result",
			},
			"result": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The result of the base36 encoding operation",
			},
		},
		MarkdownDescription: "Computes the base36 encoding of the given value",
	}
}

type EncodeBase36DataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	Result    types.String `tfsdk:"result"`
	Lowercase types.Bool   `tfsdk:"lowercase"`
	Value     types.String `tfsdk:"value"`
}

func (d *EncodeBase36DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var r EncodeBase36DataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &r)...)
	if resp.Diagnostics.HasError() {
		return
	}
	value := r.Value.ValueString()
	id := generateID(value)
	result := base36.EncodeBytes([]byte(value))
	if r.Lowercase.ValueBool() {
		result = strings.ToLower(result)
	}
	r.Result = types.StringValue(result)
	r.ID = types.StringValue(id)
	resp.Diagnostics.Append(resp.State.Set(ctx, r)...)
}

func generateID(value string) string {
	h := sha256.Sum256([]byte(value))
	return hex.EncodeToString(h[:])
}
