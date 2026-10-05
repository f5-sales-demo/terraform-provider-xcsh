---
page_title: "gcp_bucket_receiver.filename_options"
subcategory: ""
description: "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file."
xcsh_docs: {"aliases": ["gcp bucket receiver filename options"], "body_bytes": 3366, "body_sha256": "sha256:e6e385cd73bf664b2e311bbd1f182838d03f2baf17d083d3610c0c2f0c9d0424", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:log_type_folder", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:no_folder"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver", "path": "documentation/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0111020321332333-3122322332001232-0020322202302102-1112032003200032-0000000030003212-3000313132001101-3003032022010100-1111212111123200", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_bucket_receiver", "filename_options"], "schema_version": 1, "sections": [{"aliases": ["gcp bucket receiver filename options custom folder"], "anchor": "schema-gcp_bucket_receiver--filename_options--custom_folder", "description": "Exclusive with Use your own folder name as the name of the folder in the endpoint bucket or file The folder name must match `/^*$/i`", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp_bucket_receiver", "filename_options", "custom_folder"], "syntax": "attribute", "type": "string"}, {"aliases": ["gcp bucket receiver filename options log type folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:log_type_folder", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp_bucket_receiver", "filename_options", "log_type_folder"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp bucket receiver filename options no folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:no_folder", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp_bucket_receiver", "filename_options", "no_folder"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_bucket_receiver.filename_options

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/)
- gcp_bucket_receiver.filename_options

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

<a id="schema-gcp_bucket_receiver--filename_options--custom_folder"></a>

### custom_folder property

Type: `"string"`. Computed.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match.

Upstream description:

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

- [log_type_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/log_type_folder/): complete subsection reference.

- [no_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/no_folder/): complete subsection reference.

## Next pages

- [gcp_bucket_receiver.filename_options.log_type_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/log_type_folder/)
- [gcp_bucket_receiver.filename_options.no_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/no_folder/)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
