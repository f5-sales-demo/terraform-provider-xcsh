---
page_title: "rules.action"
subcategory: ""
description: "rules.action for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2778, "body_sha256": "sha256:cb0ab683d80b17061897a0c9f09eb7c1ad0c89391c940b23d3ca0135a118e813", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:action", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "path": "documentation/data-sources/nat_policy/properties/rules/action/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["rules", "action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.action for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- rules.action

<a id="section"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_nat_choice": "[\"dynamic\",\"virtual_cidr\"]"
}
```

## Direct properties

- [dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic/): complete subsection reference.

<a id="schema-rules--action--virtual_cidr"></a>

### virtual_cidr property

Type: `"string"`. Computed.

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR..

Upstream description:

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR
192.0.2.0/24, the virtual CIDR has 100.100.100.0/24.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
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

## Next pages

- [rules.action.dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
