---
page_title: "segment_policy.src_segments"
subcategory: ""
description: "List of references to Segments."
xcsh_docs: {"aliases": ["segment policy src segments"], "body_bytes": 1786, "body_sha256": "sha256:da9d7096abb589a3d06c3a8e12c9beed7bfae2d60a9a06f66ec78849f9e291da", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "path": "documentation/resources/service_policy_rule/properties/segment_policy/src_segments/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy.src_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_policy", "src_segments"], "schema_version": 1, "sections": [{"aliases": ["segments"], "anchor": "section", "description": "Select list of segments.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-segment_policy--src_segments--segments--name", "enforcement": "provider-schema", "group": "segment_policy.src_segments.segments:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "type": "requires"}], "schema_path": ["segment_policy", "src_segments", "segments"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/src_segments/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of references to Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy.src_segments

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/)
- segment_policy.src_segments

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for src segments.

Upstream description:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
src_segments {
  # Configure direct properties listed below.
}
```

## Direct properties

- [segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/segments/): complete subsection reference.

## Next pages

- [segment_policy.src_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/segments/)
- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
