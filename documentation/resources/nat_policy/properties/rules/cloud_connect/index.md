---
page_title: "rules.cloud_connect"
subcategory: ""
description: "Reference to Cloud connect Object."
xcsh_docs: {"aliases": ["rules cloud connect"], "body_bytes": 1286, "body_sha256": "sha256:b570a682070a1985ed5e40595f7b61cd62728872ebeb03e5bf98b405f144e75a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:cloud_connect:refs"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "documentation/resources/nat_policy/properties/rules/cloud_connect/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.cloud_connect:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect:refs", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "cloud_connect"], "schema_version": 1, "sections": [{"aliases": ["rules cloud connect refs"], "anchor": "section", "description": "Reference to Cloud Connect Object.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "cloud_connect", "refs"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/cloud_connect/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reference to Cloud connect Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["nat_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.cloud_connect

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- rules.cloud_connect

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cloud connect.

Additional upstream details:

Reference to Cloud connect Object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
```

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

Terraform syntax:

```terraform
cloud_connect {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/): complete subsection reference.
