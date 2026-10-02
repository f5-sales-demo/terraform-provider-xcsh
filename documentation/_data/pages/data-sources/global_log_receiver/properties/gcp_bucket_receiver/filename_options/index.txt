---
page_title: "gcp_bucket_receiver.filename_options"
subcategory: ""
description: "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file."
xcsh_docs: {"aliases": ["gcp bucket receiver filename options"], "body_bytes": 3366, "body_sha256": "sha256:49d86ee782bff5f02fcab716aadcf33b3bf43bc37d1aeaaf3cdb7e51820535b2", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:log_type_folder", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:no_folder"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver", "path": "documentation/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0111020321332333-3122322332001232-0020322202302102-1112032003200032-0000000030003212-3000313132001101-3003032022010100-1111212111123200", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_bucket_receiver", "filename_options"], "schema_version": 1, "sections": [{"aliases": ["custom folder"], "anchor": "schema-gcp_bucket_receiver--filename_options--custom_folder", "description": "Exclusive with Use your own folder name as the name of the folder in the endpoint bucket or file The folder name must match `/^*$/i`", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp_bucket_receiver", "filename_options", "custom_folder"], "syntax": "attribute", "type": "string"}, {"aliases": ["log type folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:log_type_folder", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp_bucket_receiver", "filename_options", "log_type_folder"], "syntax": "attribute", "type": "object"}, {"aliases": ["no folder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:no_folder", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp_bucket_receiver", "filename_options", "no_folder"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
