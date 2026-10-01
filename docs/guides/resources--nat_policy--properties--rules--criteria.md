---
page_title: "rules.criteria"
subcategory: ""
description: "rules.criteria for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3982, "body_sha256": "sha256:e7c2dc859033bada6a2b6bfa9432600c805c04a5cbd8eb561f38dc6a2f798382", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_inside_network", "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_network", "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "docs/guides/resources--nat_policy--properties--rules--criteria.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "criteria"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.criteria for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- [rules](resources--nat_policy--properties--rules.md)
- rules.criteria

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match criteria of the packet to apply the NAT Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "icmp"),
  validators.ConflictingObjectAttributes("any",
    "tcp"),
  validators.ConflictingObjectAttributes("any",
    "udp"),
  validators.ConflictingObjectAttributes("icmp",
    "tcp"),
  validators.ConflictingObjectAttributes("icmp",
    "udp"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
```

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

Terraform syntax:

```terraform
criteria {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any](resources--nat_policy--properties--rules--criteria--any.md): complete subsection reference.

<a id="schema-rules--criteria--destination_cidr"></a>

### destination_cidr property

Type: `["list", "string"]`. Optional.

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

- [icmp](resources--nat_policy--properties--rules--criteria--icmp.md): complete subsection reference.

- [site_local_inside_network](resources--nat_policy--properties--rules--criteria--site_local_inside_network.md): complete subsection reference.

- [site_local_network](resources--nat_policy--properties--rules--criteria--site_local_network.md): complete subsection reference.

<a id="schema-rules--criteria--source_cidr"></a>

### source_cidr property

Type: `["list", "string"]`. Optional.

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

- [tcp](resources--nat_policy--properties--rules--criteria--tcp.md): complete subsection reference.

- [udp](resources--nat_policy--properties--rules--criteria--udp.md): complete subsection reference.

## Next pages

- [rules.criteria.any](resources--nat_policy--properties--rules--criteria--any.md)
- [rules.criteria.icmp](resources--nat_policy--properties--rules--criteria--icmp.md)
- [rules.criteria.site_local_inside_network](resources--nat_policy--properties--rules--criteria--site_local_inside_network.md)
- [rules.criteria.site_local_network](resources--nat_policy--properties--rules--criteria--site_local_network.md)
- [rules.criteria.tcp](resources--nat_policy--properties--rules--criteria--tcp.md)
- [rules.criteria.udp](resources--nat_policy--properties--rules--criteria--udp.md)
- [rules](resources--nat_policy--properties--rules.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
