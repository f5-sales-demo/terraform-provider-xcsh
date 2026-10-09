---
page_title: "s3_receiver.filename_options"
subcategory: ""
description: "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file."
xcsh_docs: {"aliases": ["s3 receiver filename options"], "body_bytes": 2434, "body_sha256": "sha256:2ff16fcbf7181df30d3c81a7b22256d10ebe5f64e7114ecd32aae0be1cf6a1b6", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options:log_type_folder", "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options:no_folder"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver", "path": "documentation/data-sources/global_log_receiver/properties/s3_receiver/filename_options/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["s3_receiver", "filename_options"], "schema_version": 1, "sections": [{"aliases": ["s3 receiver filename options custom folder"], "anchor": "schema-s3_receiver--filename_options--custom_folder", "description": "Exclusive with Use your own folder name as the name of the folder in the endpoint bucket or file The folder name must match `/^*$/i`", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["s3_receiver", "filename_options", "custom_folder"], "syntax": "attribute", "type": "string"}, {"aliases": ["s3 receiver filename options log type folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options:log_type_folder", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["s3_receiver", "filename_options", "log_type_folder"], "syntax": "attribute", "type": "object"}, {"aliases": ["s3 receiver filename options no folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options:no_folder", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["s3_receiver", "filename_options", "no_folder"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/s3_receiver/filename_options/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# s3_receiver.filename_options

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [s3_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/)
- s3_receiver.filename_options

<a id="section"></a>

Type: `"single"`. Computed.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

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

## Direct properties

<a id="schema-s3_receiver--filename_options--custom_folder"></a>

### custom_folder property

Type: `"string"`. Computed.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [log_type_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/filename_options/log_type_folder/): complete subsection reference.

- [no_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/filename_options/no_folder/): complete subsection reference.
