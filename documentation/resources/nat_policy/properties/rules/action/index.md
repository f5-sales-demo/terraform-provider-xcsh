---
page_title: "rules.action"
subcategory: ""
description: "Action to apply on the packet if the NAT rule is applied."
xcsh_docs: {"aliases": ["rules action"], "body_bytes": 3210, "body_sha256": "sha256:27250fbc59d4d480f391c4672c6f5a098c70d88f027b43996cf43f8525e8141d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:action:dynamic"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:action", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "documentation/resources/nat_policy/properties/rules/action/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "schema-rules--action--virtual_cidr", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:dynamic,virtual_cidr", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:dynamic,virtual_cidr", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action"], "schema_version": 1, "sections": [{"aliases": ["dynamic"], "anchor": "section", "description": "Dynamic Pool Configuration.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic:ConflictingObjectAttributes:elastic_ips,pools", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:elastic_ips", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic:ConflictingObjectAttributes:elastic_ips,pools", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:pools", "type": "conflicts"}], "schema_path": ["rules", "action", "dynamic"], "syntax": "block", "type": "object"}, {"aliases": ["virtual cidr"], "anchor": "schema-rules--action--virtual_cidr", "description": "Exclusive with Virtual Subnet NAT is static NAT that does a one-to-one translation between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR 192.0.2.0/24, the virtual CIDR has", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "virtual_cidr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action to apply on the packet if the NAT rule is applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nat_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- rules.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action to apply on the packet if the NAT rule is applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dynamic",
    "virtual_cidr")}
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
  "x-ves-oneof-field-source_nat_choice": "[\"dynamic\",\"virtual_cidr\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/): complete subsection reference.

<a id="schema-rules--action--virtual_cidr"></a>

### virtual_cidr property

Type: `"string"`. Optional.

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR..

Upstream description:

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR
192.0.2.0/24, the virtual CIDR has 100.100.100.0/24.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

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

- [rules.action.dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
