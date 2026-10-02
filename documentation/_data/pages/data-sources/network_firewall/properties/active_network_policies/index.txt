---
page_title: "active_network_policies"
subcategory: "Security"
description: "List of firewall policy views."
xcsh_docs: {"aliases": ["active network policies"], "body_bytes": 1390, "body_sha256": "sha256:afabf0a75ceacb796e492b57fea49d478a58618c335522157664c8770e02f0f9", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_firewall:properties:active_network_policies:network_policies"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_firewall:properties:active_network_policies", "parent_id": "xcsh-docs:data-sources:network_firewall:reference", "path": "documentation/data-sources/network_firewall/properties/active_network_policies/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0300113221000300-3323031321112301-2303200000031221-3220101113022012-2132110033002231-3323003110211123-0002220301010033-3322101121111123", "registry_path": "docs/guides/data-sources--network_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["active_network_policies"], "schema_version": 1, "sections": [{"aliases": ["network policies"], "anchor": "section", "description": "Ordered List of Firewall Policies active for this network firewall.", "document_id": "xcsh-docs:data-sources:network_firewall:properties:active_network_policies:network_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["active_network_policies", "network_policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_firewall/properties/active_network_policies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of firewall policy views.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_network_policies

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/)
- active_network_policies

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

- [network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/active_network_policies/network_policies/): complete subsection reference.

## Next pages

- [active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/active_network_policies/network_policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/)
- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/)
