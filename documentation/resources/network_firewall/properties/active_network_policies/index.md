---
page_title: "active_network_policies"
subcategory: "Security"
description: "List of firewall policy views."
xcsh_docs: {"aliases": ["active network policies"], "body_bytes": 1659, "body_sha256": "sha256:5e74f981920b157cf9afe103ec53b5a5b3bca2375863d5ffce6e1a34a18061ce", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_network_policies:network_policies"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_network_policies", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "documentation/resources/network_firewall/properties/active_network_policies/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2113302023113302-3020022101313123-0003310010231013-1230332313211203-3122211313002211-2212312210212203-2312100331130132-3113020002020033", "registry_path": "docs/guides/resources--network_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "active_network_policies:RequiredObjectAttributes:network_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_network_policies:network_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["active_network_policies"], "schema_version": 1, "sections": [{"aliases": ["network policies"], "anchor": "section", "description": "Ordered List of Firewall Policies active for this network firewall.", "document_id": "xcsh-docs:resources:network_firewall:properties:active_network_policies:network_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-active_network_policies--network_policies--name", "enforcement": "provider-schema", "group": "active_network_policies.network_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_network_policies:network_policies", "type": "requires"}], "schema_path": ["active_network_policies", "network_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_network_policies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of firewall policy views.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_network_policies

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- active_network_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_network_policies/network_policies/): complete subsection reference.

## Next pages

- [active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_network_policies/network_policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
