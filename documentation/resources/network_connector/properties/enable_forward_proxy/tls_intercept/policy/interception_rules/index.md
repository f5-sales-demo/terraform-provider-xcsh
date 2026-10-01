---
page_title: "enable_forward_proxy.tls_intercept.policy.interception_rules"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept.policy.interception_rules for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 4205, "body_sha256": "sha256:cd2885c1af29e778e905c1b37d6bb7d9d35c3a1deed1850459fa3e6439d82a7d", "child_ids": ["xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:disable_interception", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:enable_interception"], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/index.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept.policy.interception_rules for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.policy.interception_rules

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [enable_forward_proxy.tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/)
- enable_forward_proxy.tls_intercept.policy.interception_rules

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

- [disable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/disable_interception/): complete subsection reference.

- [domain_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/): complete subsection reference.

- [enable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/enable_interception/): complete subsection reference.

## Next pages

- [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/disable_interception/)
- [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/)
- [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/enable_interception/)
- [enable_forward_proxy.tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
