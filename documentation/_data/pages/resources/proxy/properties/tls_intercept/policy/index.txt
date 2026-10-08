---
page_title: "tls_intercept.policy"
subcategory: ""
description: "Policy to enable or disable TLS interception."
xcsh_docs: {"aliases": ["tls intercept policy"], "body_bytes": 1252, "body_sha256": "sha256:ea3476b5dd1d44321205571cf1345b1431124c662aed0b8854cc77c69f6c240c", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "path": "documentation/resources/proxy/properties/tls_intercept/policy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.policy:RequiredObjectAttributes:interception_rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "policy"], "schema_version": 1, "sections": [{"aliases": ["tls intercept policy interception rules"], "anchor": "section", "description": "List of ordered rules to enable or disable for TLS interception.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules:ConflictingListObjectAttributes:disable_interception,enable_interception", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:disable_interception", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.policy.interception_rules:ConflictingListObjectAttributes:disable_interception,enable_interception", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules:enable_interception", "type": "conflicts"}], "schema_path": ["tls_intercept", "policy", "interception_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/policy/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Policy to enable or disable TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
EnumExtractionComplete: false
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
