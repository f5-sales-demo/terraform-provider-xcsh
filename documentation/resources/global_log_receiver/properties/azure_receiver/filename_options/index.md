---
page_title: "azure_receiver.filename_options"
subcategory: ""
description: "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file."
xcsh_docs: {"aliases": ["azure receiver filename options"], "body_bytes": 3036, "body_sha256": "sha256:9373ff6574669c0440670511ccd4c4db2800b932755b9b9794bca4cc1876d797", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:log_type_folder", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "path": "documentation/resources/global_log_receiver/properties/azure_receiver/filename_options/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-001.md", "relationships": [{"anchor": "schema-azure_receiver--filename_options--custom_folder", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:custom_folder,log_type_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "type": "conflicts"}, {"anchor": "schema-azure_receiver--filename_options--custom_folder", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:custom_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:custom_folder,log_type_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:log_type_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:log_type_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:log_type_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:custom_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:log_type_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_receiver", "filename_options"], "schema_version": 1, "sections": [{"aliases": ["azure receiver filename options custom folder"], "anchor": "schema-azure_receiver--filename_options--custom_folder", "description": "Exclusive with Use your own folder name as the name of the folder in the endpoint bucket or file The folder name must match `/^*$/i`", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "filename_options", "custom_folder"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure receiver filename options log type folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:log_type_folder", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "filename_options", "log_type_folder"], "syntax": "attribute", "type": "object"}, {"aliases": ["azure receiver filename options no folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "filename_options", "no_folder"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/filename_options/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.filename_options

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/)
- azure_receiver.filename_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

Provider validators and defaults (from schema source):

```go
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

<a id="schema-azure_receiver--filename_options--custom_folder"></a>

### custom_folder property

Type: `"string"`. Optional.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [log_type_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/filename_options/log_type_folder/): complete subsection reference.

- [no_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/filename_options/no_folder/): complete subsection reference.
