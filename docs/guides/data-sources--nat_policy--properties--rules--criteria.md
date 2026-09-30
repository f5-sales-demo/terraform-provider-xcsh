---
page_title: "rules.criteria"
subcategory: ""
description: "rules.criteria for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3261, "body_sha256": "sha256:779429f7e8c9af505efa246b37513dbbede53728ebfda17412fc308910a58d4b", "canonical_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:criteria:any", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:icmp", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:site_local_inside_network", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:site_local_network", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "path": "docs/guides/data-sources--nat_policy--properties--rules--criteria.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "criteria"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/criteria/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.criteria for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.criteria

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md)
- [Property reference](data-sources--nat_policy--reference.md)
- [rules](data-sources--nat_policy--properties--rules.md)
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

- [any](data-sources--nat_policy--properties--rules--criteria--any.md): complete subsection reference.

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

- [icmp](data-sources--nat_policy--properties--rules--criteria--icmp.md): complete subsection reference.

- [site_local_inside_network](data-sources--nat_policy--properties--rules--criteria--site_local_inside_network.md): complete subsection reference.

- [site_local_network](data-sources--nat_policy--properties--rules--criteria--site_local_network.md): complete subsection reference.

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

- [tcp](data-sources--nat_policy--properties--rules--criteria--tcp.md): complete subsection reference.

- [udp](data-sources--nat_policy--properties--rules--criteria--udp.md): complete subsection reference.

## Next pages

- [rules.criteria.any](data-sources--nat_policy--properties--rules--criteria--any.md)
- [rules.criteria.icmp](data-sources--nat_policy--properties--rules--criteria--icmp.md)
- [rules.criteria.site_local_inside_network](data-sources--nat_policy--properties--rules--criteria--site_local_inside_network.md)
- [rules.criteria.site_local_network](data-sources--nat_policy--properties--rules--criteria--site_local_network.md)
- [rules.criteria.tcp](data-sources--nat_policy--properties--rules--criteria--tcp.md)
- [rules.criteria.udp](data-sources--nat_policy--properties--rules--criteria--udp.md)
- [rules](data-sources--nat_policy--properties--rules.md)
- [xcsh_nat_policy](../data-sources/nat_policy.md)
