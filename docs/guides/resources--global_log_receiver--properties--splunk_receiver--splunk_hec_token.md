---
page_title: "splunk_receiver.splunk_hec_token"
subcategory: ""
description: "splunk_receiver.splunk_hec_token for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1941, "body_sha256": "sha256:4b876ee332a4dc3f6f5fdef69d388b306000a3545e57370d204be2a59b421820", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:splunk_hec_token", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:splunk_hec_token:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:splunk_hec_token:clear_secret_info"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:splunk_hec_token", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver", "path": "docs/guides/resources--global_log_receiver--properties--splunk_receiver--splunk_hec_token.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["splunk_receiver", "splunk_hec_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "splunk_receiver.splunk_hec_token for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# splunk_receiver.splunk_hec_token

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [splunk_receiver](resources--global_log_receiver--properties--splunk_receiver.md)
- splunk_receiver.splunk_hec_token

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
splunk_hec_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--clear_secret_info.md): complete subsection reference.

## Next pages

- [splunk_receiver.splunk_hec_token.blindfold_secret_info](resources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--blindfold_secret_info.md)
- [splunk_receiver.splunk_hec_token.clear_secret_info](resources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--clear_secret_info.md)
- [splunk_receiver](resources--global_log_receiver--properties--splunk_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
