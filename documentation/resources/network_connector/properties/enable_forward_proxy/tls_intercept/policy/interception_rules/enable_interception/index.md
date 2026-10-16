---
page_title: "enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception"
subcategory: "Networking"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable forward proxy tls intercept policy interception rules enable interception"], "body_bytes": 1740, "body_sha256": "sha256:ab553058eba5c9fbae821ef845582c1df5235c8dcb3dcf51609af6f4ea551949", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:enable_interception", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/enable_interception/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2233321120031103-2221330033023220-3303231331113011-2032300030102113-2320002001121322-1320031132001330-3122221121133001-2221002123022000", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "enable_interception"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/enable_interception/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["network_connectorCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [enable_forward_proxy.tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/)
- enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable interception.

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
enable_interception = {}
```

This is an empty object or choice marker. It has no direct properties.
