---
page_title: "active_fast_acls"
subcategory: "Security"
description: "List of Fast ACL(s)."
xcsh_docs: {"aliases": ["active fast acls"], "body_bytes": 1671, "body_sha256": "sha256:5f32e6759c9e967284512a10bf824aaec2463059ed3ba4dbefb1f4af77aad10c", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "documentation/resources/network_firewall/properties/active_fast_acls/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3112210130022303-2030213233130033-2121231033222300-2033311300122320-0021030220003103-2012033313212000-3303200320201022-1230312021203230", "registry_path": "docs/guides/resources--network_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "active_fast_acls:RequiredObjectAttributes:fast_acls", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["active_fast_acls"], "schema_version": 1, "sections": [{"aliases": ["active fast acls fast acls"], "anchor": "section", "description": "Ordered List of Fast ACL(s) active for this network firewall.", "document_id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-active_fast_acls--fast_acls--name", "enforcement": "provider-schema", "group": "active_fast_acls.fast_acls:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls", "type": "requires"}], "schema_path": ["active_fast_acls", "fast_acls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_fast_acls/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of Fast ACL(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["network_firewallCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_fast_acls

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- active_fast_acls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Additional upstream details:

List of Fast ACL(s).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("fast_acls")}
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

OneOf alternatives in this subsection:

- [active_fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_fast_acls/#section)
- [disable_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/disable_fast_acl/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_fast_acls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_fast_acls/fast_acls/): complete subsection reference.
