---
page_title: "tls_intercept.policy.interception_rules.disable_interception"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls intercept policy interception rules disable interception"], "body_bytes": 1313, "body_sha256": "sha256:bf0c0c25fd24d5bc5c5e83237bf33733d070411ca2a223c707df7dc153c796ee", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules:disable_interception", "parent_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules", "path": "documentation/data-sources/proxy/properties/tls_intercept/policy/interception_rules/disable_interception/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2120002331023310-2300002310011311-3023231120133323-3123331131120200-1110110202022323-3020312210233102-1302301033100013-2113011320103113", "registry_path": "docs/guides/data-sources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "policy", "interception_rules", "disable_interception"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/tls_intercept/policy/interception_rules/disable_interception/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["proxyCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy.interception_rules.disable_interception

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/)
- [tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/policy/)
- [tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/policy/interception_rules/)
- tls_intercept.policy.interception_rules.disable_interception

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable interception.

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

This is an empty object or choice marker. It has no direct properties.
