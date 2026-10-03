---
page_title: "enable_forward_proxy.tls_intercept.policy.interception_rules"
subcategory: "Networking"
description: "List of ordered rules to enable or disable for TLS interception."
xcsh_docs: {"aliases": ["enable forward proxy tls intercept policy interception rules"], "body_bytes": 4205, "body_sha256": "sha256:cd2885c1af29e778e905c1b37d6bb7d9d35c3a1deed1850459fa3e6439d82a7d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:disable_interception", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:enable_interception"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2212132222031223-2022113131220322-3333032311223130-1203021131321033-1301233321012022-0112103110300102-3133032102230313-3321030102212031", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.policy.interception_rules:ConflictingListObjectAttributes:disable_interception,enable_interception", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:disable_interception", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.policy.interception_rules:ConflictingListObjectAttributes:disable_interception,enable_interception", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:enable_interception", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules"], "schema_version": 1, "sections": [{"aliases": ["enable forward proxy tls intercept policy interception rules disable interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:disable_interception", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "disable_interception"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable forward proxy tls intercept policy interception rules domain match"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--exact_value", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--exact_value", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--regex_value", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--regex_value", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--suffix_value", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}, {"anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--suffix_value", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "type": "conflicts"}], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "domain_match"], "syntax": "block", "type": "object"}, {"aliases": ["enable forward proxy tls intercept policy interception rules enable interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:enable_interception", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "enable_interception"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of ordered rules to enable or disable for TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
