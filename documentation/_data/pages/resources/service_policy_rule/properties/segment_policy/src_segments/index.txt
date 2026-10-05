---
page_title: "segment_policy.src_segments"
subcategory: ""
description: "List of references to Segments."
xcsh_docs: {"aliases": ["segment policy src segments"], "body_bytes": 1786, "body_sha256": "sha256:da9d7096abb589a3d06c3a8e12c9beed7bfae2d60a9a06f66ec78849f9e291da", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "path": "documentation/resources/service_policy_rule/properties/segment_policy/src_segments/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy.src_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_policy", "src_segments"], "schema_version": 1, "sections": [{"aliases": ["segment policy src segments segments"], "anchor": "section", "description": "Select list of segments.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-segment_policy--src_segments--segments--name", "enforcement": "provider-schema", "group": "segment_policy.src_segments.segments:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "type": "requires"}], "schema_path": ["segment_policy", "src_segments", "segments"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/src_segments/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of references to Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
