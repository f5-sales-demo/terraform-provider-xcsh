---
page_title: "network_pbr.network_pbr_rules.ip_prefix_set"
subcategory: ""
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["network pbr network pbr rules ip prefix set"], "body_bytes": 1242, "body_sha256": "sha256:acb02ed42af317022fabc36695f6f6935f27aa39a416c410eb8d855c23157aac", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "parent_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules", "path": "documentation/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3132011220012130-0013120222210331-1321100223223310-0331212023012023-0321123303312223-2032302123310313-1131133233133112-1223301230221033", "registry_path": "docs/guides/data-sources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules", "ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["network pbr network pbr rules ip prefix set ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set:ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "ip_prefix_set", "ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.network_pbr_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/)
- [network_pbr.network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/)
- network_pbr.network_pbr_rules.ip_prefix_set

<a id="section"></a>

Type: `"single"`. Computed.

A list of references to ip\_prefix\_set objects.

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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/): complete subsection reference.
