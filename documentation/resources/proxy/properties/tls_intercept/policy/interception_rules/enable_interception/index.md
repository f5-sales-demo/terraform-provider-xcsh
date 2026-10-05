---
page_title: "tls_intercept.policy.interception_rules.enable_interception"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls intercept policy interception rules enable interception"], "body_bytes": 1646, "body_sha256": "sha256:92c42a834168195b42cc5831191f2e9c14cea93ca109959b361f1b58b1e05863", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:enable_interception", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "path": "documentation/resources/proxy/properties/tls_intercept/policy/interception_rules/enable_interception/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1123312222230112-2211100331321323-3022231331031311-1131132221023111-2301031013132333-3201220202231311-2310211212210221-2330130331323033", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "policy", "interception_rules", "enable_interception"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/policy/interception_rules/enable_interception/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy.interception_rules.enable_interception

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/)
- [tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/)
- [tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/)
- tls_intercept.policy.interception_rules.enable_interception

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable interception.

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
enable_interception = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
