---
page_title: "rules.action.dynamic_pool"
subcategory: ""
description: "Dynamic Pool. Dynamic Pool Configuration."
xcsh_docs: {"aliases": ["rules action dynamic pool"], "body_bytes": 1383, "body_sha256": "sha256:8e2655b08cdfb48a423188f3d42e3f625e9188b6ca94a99e63ab77380532db24", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips", "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:pools"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:action", "path": "documentation/resources/nat_policy/properties/rules/action/dynamic_pool/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0001302111202011-2132230122033202-3030010011233111-0223320203201202-3023020220121332-2233113020002132-1320211001201300-1230230011213313", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic_pool:ConflictingObjectAttributes:elastic_ips,pools", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic_pool:ConflictingObjectAttributes:elastic_ips,pools", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:pools", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "dynamic_pool"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic pool elastic ips"], "anchor": "section", "description": "List of references to Cloud Elastic IP Object.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic_pool.elastic_ips:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "type": "requires"}], "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips"], "syntax": "block", "type": "object"}, {"aliases": ["rules action dynamic pool pools"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:pools", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "pools"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/action/dynamic_pool/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Dynamic Pool. Dynamic Pool Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["nat_policyCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic_pool

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/)
- rules.action.dynamic_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Dynamic Pool. Dynamic Pool Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("elastic_ips",
    "pools")}
```

Terraform syntax:

```terraform
dynamic_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [elastic_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/): complete subsection reference.

- [pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic_pool/pools/): complete subsection reference.
