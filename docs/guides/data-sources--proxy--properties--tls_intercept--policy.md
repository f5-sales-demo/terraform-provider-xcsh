---
page_title: "tls_intercept.policy"
subcategory: ""
description: "tls_intercept.policy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1007, "body_sha256": "sha256:4be369f0de3794de2506e609c160f033e4f226c3892ad2d7a30844374c0d15eb", "canonical_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy", "child_ids": ["xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy", "parent_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept", "path": "docs/guides/data-sources--proxy--properties--tls_intercept--policy.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/tls_intercept/policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.policy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [tls_intercept](data-sources--proxy--properties--tls_intercept.md)
- tls_intercept.policy

<a id="section"></a>

Type: `"single"`. Computed.

Policy to enable or disable TLS interception.

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

## Direct properties

- [interception_rules](data-sources--proxy--properties--tls_intercept--policy--interception_rules.md): complete subsection reference.

## Next pages

- [tls_intercept.policy.interception_rules](data-sources--proxy--properties--tls_intercept--policy--interception_rules.md)
- [tls_intercept](data-sources--proxy--properties--tls_intercept.md)
- [xcsh_proxy](../data-sources/proxy.md)
