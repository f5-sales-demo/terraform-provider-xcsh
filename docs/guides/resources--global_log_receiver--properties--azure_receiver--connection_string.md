---
page_title: "azure_receiver.connection_string"
subcategory: ""
description: "azure_receiver.connection_string for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1938, "body_sha256": "sha256:fdc5f52e22fa4d84eca55590888f9197719f209444cf62f13e124dfe22f0bf84", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:clear_secret_info"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "path": "docs/guides/resources--global_log_receiver--properties--azure_receiver--connection_string.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_receiver", "connection_string"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/connection_string/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_receiver.connection_string for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.connection_string

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [azure_receiver](resources--global_log_receiver--properties--azure_receiver.md)
- azure_receiver.connection_string

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
connection_string {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--global_log_receiver--properties--azure_receiver--connection_string--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--properties--azure_receiver--connection_string--clear_secret_info.md): complete subsection reference.

## Next pages

- [azure_receiver.connection_string.blindfold_secret_info](resources--global_log_receiver--properties--azure_receiver--connection_string--blindfold_secret_info.md)
- [azure_receiver.connection_string.clear_secret_info](resources--global_log_receiver--properties--azure_receiver--connection_string--clear_secret_info.md)
- [azure_receiver](resources--global_log_receiver--properties--azure_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
