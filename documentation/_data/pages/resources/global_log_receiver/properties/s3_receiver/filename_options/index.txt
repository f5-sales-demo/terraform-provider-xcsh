---
page_title: "s3_receiver.filename_options"
subcategory: ""
description: "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file."
xcsh_docs: {"aliases": ["s3 receiver filename options"], "body_bytes": 3074, "body_sha256": "sha256:651bad11bf9dd2f6b332b7ace353ec9a72385422b04b204aae8f902eb272a590", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options:log_type_folder", "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options:no_folder"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver", "path": "documentation/resources/global_log_receiver/properties/s3_receiver/filename_options/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0310300212301021-1030032312012230-0231233231132003-0310010220111031-0000031333323200-1103211313320310-2031132312223122-1323220303231200", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-005.md", "relationships": [{"anchor": "schema-s3_receiver--filename_options--custom_folder", "enforcement": "provider-schema", "group": "s3_receiver.filename_options:ConflictingObjectAttributes:custom_folder,log_type_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options", "type": "conflicts"}, {"anchor": "schema-s3_receiver--filename_options--custom_folder", "enforcement": "provider-schema", "group": "s3_receiver.filename_options:ConflictingObjectAttributes:custom_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "s3_receiver.filename_options:ConflictingObjectAttributes:custom_folder,log_type_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options:log_type_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "s3_receiver.filename_options:ConflictingObjectAttributes:log_type_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options:log_type_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "s3_receiver.filename_options:ConflictingObjectAttributes:custom_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options:no_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "s3_receiver.filename_options:ConflictingObjectAttributes:log_type_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options:no_folder", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["s3_receiver", "filename_options"], "schema_version": 1, "sections": [{"aliases": ["s3 receiver filename options custom folder"], "anchor": "schema-s3_receiver--filename_options--custom_folder", "description": "Exclusive with Use your own folder name as the name of the folder in the endpoint bucket or file The folder name must match `/^*$/i`", "document_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["s3_receiver", "filename_options", "custom_folder"], "syntax": "attribute", "type": "string"}, {"aliases": ["s3 receiver filename options log type folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options:log_type_folder", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["s3_receiver", "filename_options", "log_type_folder"], "syntax": "attribute", "type": "object"}, {"aliases": ["s3 receiver filename options no folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:filename_options:no_folder", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["s3_receiver", "filename_options", "no_folder"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/s3_receiver/filename_options/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# s3_receiver.filename_options

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [s3_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/)
- s3_receiver.filename_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_folder",
    "log_type_folder"),
  validators.ConflictingObjectAttributes("custom_folder",
    "no_folder"),
  validators.ConflictingObjectAttributes("log_type_folder",
    "no_folder")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-folder": "[\"custom_folder\",\"log_type_folder\",\"no_folder\"]"
}
```

Terraform syntax:

```terraform
filename_options {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-s3_receiver--filename_options--custom_folder"></a>

### custom_folder property

Type: `"string"`. Optional.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  }
}
```

- [log_type_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/filename_options/log_type_folder/): complete subsection reference.

- [no_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/filename_options/no_folder/): complete subsection reference.
