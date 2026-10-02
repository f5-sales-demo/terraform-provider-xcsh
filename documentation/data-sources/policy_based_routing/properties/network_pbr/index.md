---
page_title: "network_pbr"
subcategory: ""
description: "Network(L3/L4) routing policy rule."
xcsh_docs: {"aliases": ["network pbr"], "body_bytes": 2417, "body_sha256": "sha256:2146d5b79f746cc918893a830d60f91ad33e4aea86c230f2b7f4df622310947d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:any", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:label_selector", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:prefix_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "parent_id": "xcsh-docs:data-sources:policy_based_routing:reference", "path": "documentation/data-sources/policy_based_routing/properties/network_pbr/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011", "registry_path": "docs/guides/data-sources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_pbr"], "schema_version": 1, "sections": [{"aliases": ["any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:any", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:label_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["network pbr rules"], "anchor": "section", "description": "Network(L3/L4) routing policy rule.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "prefix_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/network_pbr/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Network(L3/L4) routing policy rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- network_pbr

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for network pbr.

Upstream description:

Network(L3/L4) routing policy rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_choice": "[\"any\",\"label_selector\",\"prefix_list\"]"
}
```

## Direct properties

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/any/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/label_selector/): complete subsection reference.

- [network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/prefix_list/): complete subsection reference.

## Next pages

- [network_pbr.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/any/)
- [network_pbr.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/label_selector/)
- [network_pbr.network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/)
- [network_pbr.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/prefix_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
