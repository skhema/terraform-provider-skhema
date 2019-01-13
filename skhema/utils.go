package skhema

import (
	"github.com/skhema/terraform-provider-skhema/client"
)

func newTypeField(props map[string]interface{}) *skhema.TypeField {
	field := &skhema.TypeField{}

	if v, ok := props["name"]; ok {
		field.Name = v.(string)
	}

	if v, ok := props["type"]; ok {
		field.Type = v.(string)
	}

	return field
}

func newTypeFieldsList(seq []interface{}) []*skhema.TypeField {
	fields := make([]*skhema.TypeField, 0, len(seq))

	for _, field := range seq {
		props := field.(map[string]interface{})
		fields = append(fields, newTypeField(props))
	}

	return fields
}

func newApiOperationConsumable(props map[string]interface{}) *skhema.ApiOperationConsumable {
	r := &skhema.ApiOperationConsumable{}

	if v, ok := props["format"]; ok {
		r.Format = v.(string)
	}

	if v, ok := props["type"]; ok {
		r.Type = v.(string)
	}

	if schema, ok := props["schema"]; ok {
		s := schema.(map[string]interface{})
		r.Schema = &skhema.ApiOperationSchema{}

		if namespace, ok := s["namespace"]; ok {
			r.Schema.Namespace = namespace.(string)
		}

		if name, ok := s["name"]; ok {
			r.Schema.Name = name.(string)
		}

		if revision, ok := s["revision"]; ok {
			r.Schema.Revision = revision.(string)
		}
	}

	return r
}

func newApiOperationConsumablesList(seq []interface{}) []*skhema.ApiOperationConsumable {
	consumables := make([]*skhema.ApiOperationConsumable, 0, len(seq))

	for _, consumable := range seq {
		props := consumable.(map[string]interface{})
		consumables = append(consumables, newApiOperationConsumable(props))
	}

	return consumables
}

func newApiOperationProducible(props map[string]interface{}) *skhema.ApiOperationProducible {
	r := &skhema.ApiOperationProducible{}

	if v, ok := props["status"]; ok {
		r.Status = v.(string)
	}

	if v, ok := props["format"]; ok {
		r.Type = v.(string)
	}

	if v, ok := props["type"]; ok {
		r.Type = v.(string)
	}

	if schema, ok := props["schema"]; ok {
		s := schema.(map[string]interface{})
		r.Schema = &skhema.ApiOperationSchema{}

		if namespace, ok := s["namespace"]; ok {
			r.Schema.Namespace = namespace.(string)
		}

		if name, ok := s["name"]; ok {
			r.Schema.Name = name.(string)
		}

		if revision, ok := s["revision"]; ok {
			r.Schema.Revision = revision.(string)
		}
	}

	return r
}

func newApiOperationProduciblesList(seq []interface{}) []*skhema.ApiOperationProducible {
	producibles := make([]*skhema.ApiOperationProducible, 0, len(seq))

	for _, producible := range seq {
		props := producible.(map[string]interface{})
		producibles = append(producibles, newApiOperationProducible(props))
	}

	return producibles
}

func newApiOperationParam(props map[string]interface{}) *skhema.ApiOperationParam {
	p := &skhema.ApiOperationParam{}

	if v, ok := props["name"]; ok {
		p.Name = v.(string)
	}

	if v, ok := props["segment"]; ok {
		p.Segment = v.(string)
	}

	if v, ok := props["type"]; ok {
		p.Type = v.(string)
	}

	return p
}

func newApiOperationParamsList(seq []interface{}) []*skhema.ApiOperationParam {
	params := make([]*skhema.ApiOperationParam, 0, len(seq))

	for _, param := range seq {
		props := param.(map[string]interface{})
		params = append(params, newApiOperationParam(props))
	}

	return params
}

func newApiOperation(props map[string]interface{}) *skhema.ApiOperation {
	op := &skhema.ApiOperation{}

	if v, ok := props["name"]; ok {
		op.Name = v.(string)
	}

	if v, ok := props["path"]; ok {
		op.Path = v.(string)
	}

	if v, ok := props["method"]; ok {
		op.Method = v.(string)
	}

	if v, ok := props["param"]; ok {
		op.Params = newApiOperationParamsList(v.([]interface{}))
	}

	if v, ok := props["consume"]; ok {
		op.Consumables = newApiOperationConsumablesList(v.([]interface{}))
	}

	if v, ok := props["produce"]; ok {
		op.Producibles = newApiOperationProduciblesList(v.([]interface{}))
	}

	return op
}

func newApiOperationsList(seq []interface{}) []*skhema.ApiOperation {
	ops := make([]*skhema.ApiOperation, 0, len(seq))

	for _, op := range seq {
		props := op.(map[string]interface{})
		ops = append(ops, newApiOperation(props))
	}

	return ops
}

func flattenTypeFieldsList(in []*skhema.TypeField) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i, v := range in {
		m := make(map[string]interface{})
		m["name"] = v.Name
		m["type"] = v.Type

		out[i] = m
	}

	return out
}

func flattenApiOperationConsumables(in []*skhema.ApiOperationConsumable) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i, v := range in {
		s := make(map[string]interface{})
		s["namespace"] = v.Schema.Namespace
		s["name"] = v.Schema.Name
		s["revision"] = v.Schema.Revision

		m := make(map[string]interface{})
		m["format"] = v.Format
		m["type"] = v.Type
		m["schema"] = s

		out[i] = m
	}

	return out
}

func flattenApiOperationProducibles(in []*skhema.ApiOperationProducible) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i, v := range in {
		s := make(map[string]interface{})
		s["namespace"] = v.Schema.Namespace
		s["name"] = v.Schema.Name
		s["revision"] = v.Schema.Revision

		m := make(map[string]interface{})
		m["status"] = v.Status
		m["format"] = v.Format
		m["type"] = v.Type
		m["schema"] = s

		out[i] = m
	}

	return out
}

func flattenApiOperationParams(in []*skhema.ApiOperationParam) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i, v := range in {
		m := make(map[string]interface{})
		m["name"] = v.Name
		m["type"] = v.Type
		m["segment"] = v.Segment

		out[i] = m
	}

	return out
}

func flattenApiOperations(in []*skhema.ApiOperation) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i, v := range in {
		m := make(map[string]interface{})
		m["name"] = v.Name
		m["path"] = v.Path
		m["method"] = v.Method
		m["param"] = flattenApiOperationParams(v.Params)
		m["consumables"] = flattenApiOperationConsumables(v.Consumables)
		m["producibles"] = flattenApiOperationProducibles(v.Producibles)

		out[i] = m
	}

	return out
}
