---
page_title: "datadog_receiver.datadog_api_key"
subcategory: ""
description: "datadog_receiver.datadog_api_key for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1845, "body_sha256": "sha256:d90e1c6ac98834bf165509987f20913b1d230820abe7a49d616c5251f61b3206", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver", "path": "docs/guides/resources--global_log_receiver--properties--datadog_receiver--datadog_api_key.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["datadog_receiver", "datadog_api_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "datadog_receiver.datadog_api_key for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# datadog_receiver.datadog_api_key

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [datadog_receiver](resources--global_log_receiver--properties--datadog_receiver.md)
- datadog_receiver.datadog_api_key

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
datadog_api_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--global_log_receiver--properties--datadog_receiver--datadog_api_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--properties--datadog_receiver--datadog_api_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [datadog_receiver.datadog_api_key.blindfold_secret_info](resources--global_log_receiver--properties--datadog_receiver--datadog_api_key--blindfold_secret_info.md)
- [datadog_receiver.datadog_api_key.clear_secret_info](resources--global_log_receiver--properties--datadog_receiver--datadog_api_key--clear_secret_info.md)
- [datadog_receiver](resources--global_log_receiver--properties--datadog_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
