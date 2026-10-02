---
page_title: "segment_policy"
subcategory: ""
description: "Configure source and destination segment for policy."
xcsh_docs: {"aliases": ["segment policy"], "body_bytes": 3276, "body_sha256": "sha256:84b0656a8b7905d2eda428d343dc85810893db741e8ae31db6fe752f1e1d7132", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_any", "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments", "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_any", "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/segment_policy/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_any,dst_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_any,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_any,dst_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_segments,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_any,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_segments,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:src_any,src_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:src_any,src_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_policy"], "schema_version": 1, "sections": [{"aliases": ["dst any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_any", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["segment_policy", "dst_any"], "syntax": "attribute", "type": "object"}, {"aliases": ["dst segments"], "anchor": "section", "description": "List of references to Segments.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy.dst_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments:segments", "type": "requires"}], "schema_path": ["segment_policy", "dst_segments"], "syntax": "block", "type": "object"}, {"aliases": ["intra segment"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["segment_policy", "intra_segment"], "syntax": "attribute", "type": "object"}, {"aliases": ["src any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_any", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["segment_policy", "src_any"], "syntax": "attribute", "type": "object"}, {"aliases": ["src segments"], "anchor": "section", "description": "List of references to Segments.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy.src_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "type": "requires"}], "schema_path": ["segment_policy", "src_segments"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure source and destination segment for policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- segment_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure source and destination segment for policy.

Provider validators and defaults (from schema source):

```go
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

- [dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_any/): complete subsection reference.

- [dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/): complete subsection reference.

- [intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/intra_segment/): complete subsection reference.

- [src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_any/): complete subsection reference.

- [src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/): complete subsection reference.

## Next pages

- [segment_policy.dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_any/)
- [segment_policy.dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/)
- [segment_policy.intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/intra_segment/)
- [segment_policy.src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_any/)
- [segment_policy.src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
