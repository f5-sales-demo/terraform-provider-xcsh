---
page_title: "rules.criteria"
subcategory: ""
description: "rules.criteria for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4106, "body_sha256": "sha256:53332e16870bd2eb2aed2b19c6ff8255583ca5d2d3e4e5a8e905533b99b6d4d6", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:criteria:any", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:icmp", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:site_local_inside_network", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:site_local_network", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "path": "documentation/data-sources/nat_policy/properties/rules/criteria/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["rules", "criteria"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/criteria/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.criteria for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
