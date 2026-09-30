---
page_title: "enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1435, "body_sha256": "sha256:9df44c581ab3b747d15f8b6dc002aaa7f7bad9667d436d5a335ee987912d179b", "canonical_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:disable_ocsp_stapling", "child_ids": [], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "path": "docs/guides/resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--disable_ocsp_stapling.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "disable_ocsp_stapling"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
