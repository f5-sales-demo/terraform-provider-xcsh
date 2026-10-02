---
page_title: "rules"
subcategory: "Security"
description: "Shape of Rule Choice."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 1550, "body_sha256": "sha256:a4d7a78dbc110bf3a0e0f4ba0adf614faec36764acb37e97f657adc11107c9d2", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_policy:properties:rules:egress_rules", "xcsh-docs:data-sources:network_policy:properties:rules:ingress_rules"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:rules", "parent_id": "xcsh-docs:data-sources:network_policy:reference", "path": "documentation/data-sources/network_policy/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212", "registry_path": "docs/guides/data-sources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["egress rules"], "anchor": "section", "description": "Ordered list of rules applied to connections from policy endpoints.", "document_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "egress_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules"], "anchor": "section", "description": "Ordered list of rules applied to connections to policy endpoints.", "document_id": "xcsh-docs:data-sources:network_policy:properties:rules:ingress_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "ingress_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Shape of Rule Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- rules

<a id="section"></a>

Type: `"single"`. Computed.

Rule Choice. Shape of Rule Choice.

Upstream description:

Shape of Rule Choice.

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

- [egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/): complete subsection reference.

- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/ingress_rules/): complete subsection reference.

## Next pages

- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/)
- [rules.ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/ingress_rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
