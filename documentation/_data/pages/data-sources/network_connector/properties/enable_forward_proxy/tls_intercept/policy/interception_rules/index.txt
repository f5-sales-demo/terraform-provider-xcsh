---
page_title: "enable_forward_proxy.tls_intercept.policy.interception_rules"
subcategory: "Networking"
description: "List of ordered rules to enable or disable for TLS interception."
xcsh_docs: {"aliases": ["enable forward proxy tls intercept policy interception rules"], "body_bytes": 3926, "body_sha256": "sha256:bd0a14c55a759fe682310dfd912024e4338872e310d06ccd1fe3a5640940c7ac", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:disable_interception", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:enable_interception"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules", "parent_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "path": "documentation/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3100030312323221-1220233133200221-2233203233210222-0120210102323100-1212122100312022-0112333012323333-3122220022013032-2132331300021210", "registry_path": "docs/guides/data-sources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules"], "schema_version": 1, "sections": [{"aliases": ["enable forward proxy tls intercept policy interception rules disable interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:disable_interception", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "disable_interception"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable forward proxy tls intercept policy interception rules domain match"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "domain_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable forward proxy tls intercept policy interception rules enable interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:enable_interception", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "enable_interception"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of ordered rules to enable or disable for TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.policy.interception_rules

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [enable_forward_proxy.tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/)
- enable_forward_proxy.tls_intercept.policy.interception_rules

<a id="section"></a>

Type: `"list"`. Computed.

List of ordered rules to enable or disable for TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [disable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/disable_interception/): complete subsection reference.

- [domain_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/): complete subsection reference.

- [enable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/enable_interception/): complete subsection reference.

## Next pages

- [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/disable_interception/)
- [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/)
- [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/enable_interception/)
- [enable_forward_proxy.tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
