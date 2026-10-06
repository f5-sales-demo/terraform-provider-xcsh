---
page_title: "rules.spec.segment_policy"
subcategory: "Security"
description: "Configure source and destination segment for policy."
xcsh_docs: {"aliases": ["rules spec segment policy"], "body_bytes": 2591, "body_sha256": "sha256:6dc2ed5bdcd38f98767092f782232f0707fd575ff912bbc07e234b33a0ba3a5f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_any", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_segments"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "path": "documentation/resources/rate_limiter_policy/properties/rules/spec/segment_policy/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312", "registry_path": "docs/guides/resources--rate_limiter_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy:ConflictingObjectAttributes:dst_any,dst_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy:ConflictingObjectAttributes:dst_any,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy:ConflictingObjectAttributes:dst_any,dst_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy:ConflictingObjectAttributes:dst_segments,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy:ConflictingObjectAttributes:dst_any,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy:ConflictingObjectAttributes:dst_segments,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy:ConflictingObjectAttributes:src_any,src_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy:ConflictingObjectAttributes:src_any,src_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_segments", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "spec", "segment_policy"], "schema_version": 1, "sections": [{"aliases": ["rules spec segment policy dst any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_any", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "dst_any"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec segment policy dst segments"], "anchor": "section", "description": "List of references to Segments.", "document_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy.dst_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments:segments", "type": "requires"}], "schema_path": ["rules", "spec", "segment_policy", "dst_segments"], "syntax": "block", "type": "object"}, {"aliases": ["rules spec segment policy intra segment"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "intra_segment"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec segment policy src any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "src_any"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec segment policy src segments"], "anchor": "section", "description": "List of references to Segments.", "document_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_segments", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec.segment_policy.src_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_segments:segments", "type": "requires"}], "schema_path": ["rules", "spec", "segment_policy", "src_segments"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/segment_policy/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Configure source and destination segment for policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.segment_policy

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/)
- rules.spec.segment_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure source and destination segment for policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dst_any",
    "dst_segments"),
  validators.ConflictingObjectAttributes("dst_any",
    "intra_segment"),
  validators.ConflictingObjectAttributes("dst_segments",
    "intra_segment"),
  validators.ConflictingObjectAttributes("src_any",
    "src_segments")}
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
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

Terraform syntax:

```terraform
segment_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_any/): complete subsection reference.

- [dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/): complete subsection reference.

- [intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/): complete subsection reference.

- [src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_any/): complete subsection reference.

- [src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/): complete subsection reference.
