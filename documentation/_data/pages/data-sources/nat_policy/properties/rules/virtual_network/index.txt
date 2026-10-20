---
page_title: "rules.virtual_network"
subcategory: ""
description: "Carries the reference to virtual network."
xcsh_docs: {"aliases": ["rules virtual network"], "body_bytes": 942, "body_sha256": "sha256:ecb97319cde97d76a97368faf216668cd2f26b998783748898009992463980c7", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:virtual_network:refs"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:virtual_network", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "path": "documentation/data-sources/nat_policy/properties/rules/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1222133003113213-2211003133231313-3030010301222022-1032232221210133-2133032033131032-2133122120012312-3112332233032010-3032123201113221", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "virtual_network"], "schema_version": 1, "sections": [{"aliases": ["rules virtual network refs"], "anchor": "section", "description": "Reference to virtual network.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:virtual_network:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "virtual_network", "refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/virtual_network/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Carries the reference to virtual network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["nat_policyCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.virtual_network

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- rules.virtual_network

<a id="section"></a>

Type: `"single"`. Computed.

Carries the reference to virtual network.

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

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/virtual_network/refs/): complete subsection reference.
