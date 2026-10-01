---
page_title: "azure_receiver.filename_options"
subcategory: ""
description: "azure_receiver.filename_options for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 3417, "body_sha256": "sha256:810faf7cf6d5a7d8c5aafaa2afc9fcad380ea65ef1cc1bed7662059016817c6b", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:log_type_folder", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "path": "docs/guides/resources--global_log_receiver--properties--azure_receiver--filename_options.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_receiver", "filename_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/filename_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_receiver.filename_options for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.filename_options

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [azure_receiver](resources--global_log_receiver--properties--azure_receiver.md)
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
in the endpoint bucket or file The folder name must match.

Upstream description:

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

- [log_type_folder](resources--global_log_receiver--properties--azure_receiver--filename_options--log_type_folder.md): complete subsection reference.

- [no_folder](resources--global_log_receiver--properties--azure_receiver--filename_options--no_folder.md): complete subsection reference.

## Next pages

- [azure_receiver.filename_options.log_type_folder](resources--global_log_receiver--properties--azure_receiver--filename_options--log_type_folder.md)
- [azure_receiver.filename_options.no_folder](resources--global_log_receiver--properties--azure_receiver--filename_options--no_folder.md)
- [azure_receiver](resources--global_log_receiver--properties--azure_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
