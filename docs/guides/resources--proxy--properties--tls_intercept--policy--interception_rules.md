---
page_title: "tls_intercept.policy.interception_rules"
subcategory: ""
description: "tls_intercept.policy.interception_rules for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2943, "body_sha256": "sha256:6084be92bb11363eb109f42d036d99f9439444e85a9d4d64b73cf17fd56eb4e4", "canonical_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:disable_interception", "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:enable_interception"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "path": "docs/guides/resources--proxy--properties--tls_intercept--policy--interception_rules.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "policy", "interception_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/policy/interception_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.policy.interception_rules for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy.interception_rules

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- [tls_intercept.policy](resources--proxy--properties--tls_intercept--policy.md)
- tls_intercept.policy.interception_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of ordered rules to enable or disable for TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("disable_interception",
    "enable_interception")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
interception_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_interception](resources--proxy--properties--tls_intercept--policy--interception_rules--disable_interception.md): complete subsection reference.

- [domain_match](resources--proxy--properties--tls_intercept--policy--interception_rules--domain_match.md): complete subsection reference.

- [enable_interception](resources--proxy--properties--tls_intercept--policy--interception_rules--enable_interception.md): complete subsection reference.

## Next pages

- [tls_intercept.policy.interception_rules.disable_interception](resources--proxy--properties--tls_intercept--policy--interception_rules--disable_interception.md)
- [tls_intercept.policy.interception_rules.domain_match](resources--proxy--properties--tls_intercept--policy--interception_rules--domain_match.md)
- [tls_intercept.policy.interception_rules.enable_interception](resources--proxy--properties--tls_intercept--policy--interception_rules--enable_interception.md)
- [tls_intercept.policy](resources--proxy--properties--tls_intercept--policy.md)
- [xcsh_proxy](../resources/proxy.md)
