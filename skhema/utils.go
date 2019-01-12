package skhema

import (
	"github.com/skhema/terraform-provider-skhema/client"
)

func newField(props map[string]interface{}) *skhema.Field {
	field := &skhema.Field{}

	if v, ok := props["name"]; ok {
		field.Name = v.(string)
	}

	if v, ok := props["type"]; ok {
		field.Type = v.(string)
	}

	return field
}

func newFieldList(seq []interface{}) []*skhema.Field {
	fields := make([]*skhema.Field, 0, len(seq))

	for _, field := range seq {
		props := field.(map[string]interface{})
		fields = append(fields, newField(props))
	}

	return fields
}

func flattenFieldList(in []*skhema.Field) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i, v := range in {
		m := make(map[string]interface{})
		m["name"] = v.Name
		m["type"] = v.Type

		out[i] = m
	}

	return out
}
