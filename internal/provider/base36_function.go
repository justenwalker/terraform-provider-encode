package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"strings"
)

func NewBase36EncoderFunction() function.Function {
	return &Base36EncoderFunction{}
}

var _ function.Function = &Base36EncoderFunction{}

type Base36EncoderFunction struct {
}

func (b Base36EncoderFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "base36"
}

func (b Base36EncoderFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:      false,
				AllowUnknownValues:  false,
				Description:         "the string that will be encoded into base36",
				MarkdownDescription: "the string that will be encoded into base36",
				Name:                "value",
			},
		},
		Return: function.StringReturn{},
	}
}

func (b Base36EncoderFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value basetypes.StringValue
	resp.Error = req.Arguments.Get(ctx, &value)
	if resp.Error != nil {
		return
	}
	encoded := strings.ToLower(encodeBase36(value.ValueString()))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, basetypes.NewStringValue(encoded)))
}
