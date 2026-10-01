---
page_title: "enable_forward_proxy.tls_intercept.custom_certificate.private_key"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept.custom_certificate.private_key for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 2555, "body_sha256": "sha256:bb683301184652ba0f30b5322b582b0e1d2d3f6aaa1c3d6217db02ce4a964b22", "canonical_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key", "child_ids": ["xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:blindfold_secret_info", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "path": "docs/guides/resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept.custom_certificate.private_key for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.custom_certificate.private_key

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- [enable_forward_proxy](resources--network_connector--properties--enable_forward_proxy.md)
- [enable_forward_proxy.tls_intercept](resources--network_connector--properties--enable_forward_proxy--tls_intercept.md)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key

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
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info.md)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md)
- [xcsh_network_connector](../resources/network_connector.md)
