---
page_title: "s3_receiver.filename_options"
subcategory: ""
description: "s3_receiver.filename_options for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2809, "body_sha256": "sha256:a525f6611c71f4df6d9180139d15b567084103cc96db7ab149336fffd1f81784", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options:log_type_folder", "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options:no_folder"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver", "path": "docs/guides/data-sources--global_log_receiver--properties--s3_receiver--filename_options.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["s3_receiver", "filename_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/s3_receiver/filename_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "s3_receiver.filename_options for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# s3_receiver.filename_options

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [s3_receiver](data-sources--global_log_receiver--properties--s3_receiver.md)
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

- [log_type_folder](data-sources--global_log_receiver--properties--s3_receiver--filename_options--log_type_folder.md): complete subsection reference.

- [no_folder](data-sources--global_log_receiver--properties--s3_receiver--filename_options--no_folder.md): complete subsection reference.

## Next pages

- [s3_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--properties--s3_receiver--filename_options--log_type_folder.md)
- [s3_receiver.filename_options.no_folder](data-sources--global_log_receiver--properties--s3_receiver--filename_options--no_folder.md)
- [s3_receiver](data-sources--global_log_receiver--properties--s3_receiver.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
