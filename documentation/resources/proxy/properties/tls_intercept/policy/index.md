---
page_title: "tls_intercept.policy"
subcategory: ""
description: "Policy to enable or disable TLS interception."
xcsh_docs: {"aliases": ["tls intercept policy"], "body_bytes": 1613, "body_sha256": "sha256:6d4f2c54f0d8be6eeed978e3ab5bb900ed69ffa6d71e8d2c56af48f17f884e37", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "path": "documentation/resources/proxy/properties/tls_intercept/policy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.policy:RequiredObjectAttributes:interception_rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "policy"], "schema_version": 1, "sections": [{"aliases": ["interception rules"], "anchor": "section", "description": "List of ordered rules to enable or disable for TLS interception.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules:ConflictingListObjectAttributes:disable_interception,enable_interception", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:disable_interception", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules:ConflictingListObjectAttributes:disable_interception,enable_interception", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:enable_interception", "type": "conflicts"}], "schema_path": ["tls_intercept", "policy", "interception_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Policy to enable or disable TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/)
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

- [interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/): complete subsection reference.

## Next pages

- [tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/interception_rules/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
