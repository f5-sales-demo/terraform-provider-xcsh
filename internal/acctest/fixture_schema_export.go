// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"os"
)

func ExportAcceptanceSchema(path string) error {
	server, err := ProtoV6ProviderFactories["xcsh"]()
	if err != nil {
		return err
	}
	schema, err := server.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		return err
	}
	resources := map[string]interface{}{}
	sources := map[string]interface{}{}
	for name, s := range schema.ResourceSchemas {
		resources[name] = map[string]interface{}{"block": exportSchemaBlock(s.Block)}
	}
	for name, s := range schema.DataSourceSchemas {
		sources[name] = map[string]interface{}{"block": exportSchemaBlock(s.Block)}
	}
	data, err := json.Marshal(map[string]interface{}{"format_version": "1.0", "provider_schemas": map[string]interface{}{"xcsh": map[string]interface{}{"resource_schemas": resources, "data_source_schemas": sources}}})
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
func exportSchemaBlock(block *tfprotov6.SchemaBlock) map[string]interface{} {
	attributes := map[string]interface{}{}
	blocks := map[string]interface{}{}
	for _, attribute := range block.Attributes {
		attributes[attribute.Name] = map[string]interface{}{"required": attribute.Required, "optional": attribute.Optional, "computed": attribute.Computed, "type": attribute.ValueType()}
	}
	for _, nested := range block.BlockTypes {
		blocks[nested.TypeName] = map[string]interface{}{"block": exportSchemaBlock(nested.Block), "min_items": nested.MinItems}
	}
	return map[string]interface{}{"attributes": attributes, "block_types": blocks}
}
