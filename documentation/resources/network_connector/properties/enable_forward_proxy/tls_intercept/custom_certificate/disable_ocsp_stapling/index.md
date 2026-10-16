---
page_title: "enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling"
subcategory: "Networking"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable forward proxy tls intercept custom certificate disable ocsp stapling"], "body_bytes": 1535, "body_sha256": "sha256:fbdd4627c9c5fd210ef967de8127b0203acc4462887c4244f97307aa000b119c", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/disable_ocsp_stapling/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3021203321310222-3123232223313100-2312011031021130-0113310311212011-1031322113213221-1011320213122300-0012121210212103-1210003130312211", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "disable_ocsp_stapling"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["network_connectorCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [enable_forward_proxy.tls_intercept.custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/)
- enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
