---
page_title: "rules.action"
subcategory: ""
description: "Action to apply on the packet if the NAT rule is applied."
xcsh_docs: {"aliases": ["rules action"], "body_bytes": 2106, "body_sha256": "sha256:0084b9099e5e7da04858a2ec7df19c71d408a78c8fb8a86cae5413d8388c758e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:action", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "path": "documentation/data-sources/nat_policy/properties/rules/action/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "data-sources", "registry_anchor": "canonical-1113032111000113-3122312303102131-3301332200331203-1032122102301330-3023222303210123-3113201001000301-2112101101312111-3211320212303220", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic pool"], "anchor": "section", "description": "Dynamic Pool. Dynamic Pool Configuration.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "action", "dynamic_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules action virtual cidr"], "anchor": "schema-rules--action--virtual_cidr", "description": "Exclusive with Virtual Subnet NAT is static NAT that does a one-to-one translation between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR 192.0.2.0/24, the virtual CIDR has", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "virtual_cidr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/action/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Action to apply on the packet if the NAT rule is applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["nat_policyCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
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

- [dynamic_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic_pool/): complete subsection reference.

<a id="schema-rules--action--virtual_cidr"></a>

### virtual_cidr property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-13T13:24:52+00:00"
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
