---
page_title: "tls_intercept.policy"
subcategory: ""
description: "tls_intercept.policy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1258, "body_sha256": "sha256:b9c980b95613eec7baff097a17780fef50b09a1c85bf68b23d462781b6787420", "canonical_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "path": "docs/guides/resources--proxy--properties--tls_intercept--policy.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.policy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- tls_intercept.policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Policy to enable or disable TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interception_rules")}
```

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
policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interception_rules](resources--proxy--properties--tls_intercept--policy--interception_rules.md): complete subsection reference.

## Next pages

- [tls_intercept.policy.interception_rules](resources--proxy--properties--tls_intercept--policy--interception_rules.md)
- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- [xcsh_proxy](../resources/proxy.md)
