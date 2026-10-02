---
page_title: "tls_intercept.policy.interception_rules"
subcategory: ""
description: "List of ordered rules to enable or disable for TLS interception."
xcsh_docs: {"aliases": ["tls intercept policy interception rules"], "body_bytes": 3537, "body_sha256": "sha256:bf4250c747e5b71d0abf314ee986e913a1e1597b058f1ad90afc453f3577f77f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:disable_interception", "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:enable_interception"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "path": "documentation/resources/proxy/properties/tls_intercept/policy/interception_rules/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "policy", "interception_rules"], "schema_version": 1, "sections": [{"aliases": ["disable interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:disable_interception", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "policy", "interception_rules", "disable_interception"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain match"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tls_intercept--policy--interception_rules--domain_match--exact_value", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-tls_intercept--policy--interception_rules--domain_match--exact_value", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-tls_intercept--policy--interception_rules--domain_match--regex_value", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-tls_intercept--policy--interception_rules--domain_match--regex_value", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-tls_intercept--policy--interception_rules--domain_match--suffix_value", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-tls_intercept--policy--interception_rules--domain_match--suffix_value", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}], "schema_path": ["tls_intercept", "policy", "interception_rules", "domain_match"], "syntax": "block", "type": "object"}, {"aliases": ["enable interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:enable_interception", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "policy", "interception_rules", "enable_interception"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/policy/interception_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of ordered rules to enable or disable for TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy.interception_rules

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/)
- [tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/)
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

- [disable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/disable_interception/): complete subsection reference.

- [domain_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/domain_match/): complete subsection reference.

- [enable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/enable_interception/): complete subsection reference.

## Next pages

- [tls_intercept.policy.interception_rules.disable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/disable_interception/)
- [tls_intercept.policy.interception_rules.domain_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/domain_match/)
- [tls_intercept.policy.interception_rules.enable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/enable_interception/)
- [tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
