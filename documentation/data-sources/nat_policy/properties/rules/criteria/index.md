---
page_title: "rules.criteria"
subcategory: ""
description: "Match criteria of the packet to apply the NAT Rule."
xcsh_docs: {"aliases": ["rules criteria"], "body_bytes": 4205, "body_sha256": "sha256:a38ab66f802607dbf3f0d68832f34c7ecb22e0f32c6546cfcefdd3002c4a6f66", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:criteria:any", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:icmp", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:site_local_inside_network", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:site_local_network", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "path": "documentation/data-sources/nat_policy/properties/rules/criteria/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria"], "schema_version": 1, "sections": [{"aliases": ["rules criteria any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:any", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria destination cidr"], "anchor": "schema-rules--criteria--destination_cidr", "description": "Destination IP of the packet to match.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "destination_cidr"], "syntax": "attribute", "type": "list"}, {"aliases": ["rules criteria icmp"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:icmp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "icmp"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:site_local_inside_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:site_local_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria source cidr"], "anchor": "schema-rules--criteria--source_cidr", "description": "Source IP of the packet to match.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "source_cidr"], "syntax": "attribute", "type": "list"}, {"aliases": ["rules criteria tcp"], "anchor": "section", "description": "Action to apply on the packet if the NAT rule is applied.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "criteria", "tcp"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria udp"], "anchor": "section", "description": "Action to apply on the packet if the NAT rule is applied.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "criteria", "udp"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/criteria/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Match criteria of the packet to apply the NAT Rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nat_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- rules.criteria

<a id="section"></a>

Type: `"single"`. Computed.

Match criteria of the packet to apply the NAT Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-protocol_choice": "[\"any\",\"icmp\",\"tcp\",\"udp\"]"
}
```

## Direct properties

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/any/): complete subsection reference.

<a id="schema-rules--criteria--destination_cidr"></a>

### destination_cidr property

Type: `["list", "string"]`. Computed.

Destination IP. Destination IP of the packet to match.

Upstream description:

Destination IP of the packet to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/icmp/): complete subsection reference.

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/site_local_network/): complete subsection reference.

<a id="schema-rules--criteria--source_cidr"></a>

### source_cidr property

Type: `["list", "string"]`. Computed.

Source IP. Source IP of the packet to match.

Upstream description:

Source IP of the packet to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/tcp/): complete subsection reference.

- [udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/udp/): complete subsection reference.

## Next pages

- [rules.criteria.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/any/)
- [rules.criteria.icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/icmp/)
- [rules.criteria.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/site_local_inside_network/)
- [rules.criteria.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/site_local_network/)
- [rules.criteria.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/tcp/)
- [rules.criteria.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/udp/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
