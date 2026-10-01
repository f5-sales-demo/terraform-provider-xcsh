---
page_title: "enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1534, "body_sha256": "sha256:639baf3695f20c5afb7398b179a764798904b2fbd613b7f1c4de31e5f01d5c59", "canonical_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:disable_ocsp_stapling", "child_ids": [], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "path": "docs/guides/resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--disable_ocsp_stapling.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "disable_ocsp_stapling"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- [enable_forward_proxy](resources--network_connector--properties--enable_forward_proxy.md)
- [enable_forward_proxy.tls_intercept](resources--network_connector--properties--enable_forward_proxy--tls_intercept.md)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md)
- enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md)
- [xcsh_network_connector](../resources/network_connector.md)
