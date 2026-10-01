---
page_title: "splunk_receiver.use_tls.mtls_enable.key_url"
subcategory: ""
description: "splunk_receiver.use_tls.mtls_enable.key_url for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2297, "body_sha256": "sha256:4db025e676b3e317edb451effbd487c5d3b04095dbcda0c071e5dd5c8422bd69", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_enable:key_url", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_enable:key_url:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_enable:key_url:clear_secret_info"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_enable:key_url", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_enable", "path": "docs/guides/resources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["splunk_receiver", "use_tls", "mtls_enable", "key_url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "splunk_receiver.use_tls.mtls_enable.key_url for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# splunk_receiver.use_tls.mtls_enable.key_url

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [splunk_receiver](resources--global_log_receiver--properties--splunk_receiver.md)
- [splunk_receiver.use_tls](resources--global_log_receiver--properties--splunk_receiver--use_tls.md)
- [splunk_receiver.use_tls.mtls_enable](resources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable.md)
- splunk_receiver.use_tls.mtls_enable.key_url

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
key_url {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md): complete subsection reference.

## Next pages

- [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md)
- [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md)
- [splunk_receiver.use_tls.mtls_enable](resources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
