---
page_title: "segment_policy.dst_segments"
subcategory: ""
description: "List of references to Segments."
xcsh_docs: {"aliases": ["segment policy dst segments"], "body_bytes": 1786, "body_sha256": "sha256:b82ce44055cb960539ce7f44c30dab082fcc5224628c6d3ffe027131fa6eb5d1", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments:segments"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "path": "documentation/resources/service_policy_rule/properties/segment_policy/dst_segments/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy.dst_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments:segments", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_policy", "dst_segments"], "schema_version": 1, "sections": [{"aliases": ["segment policy dst segments segments"], "anchor": "section", "description": "Select list of segments.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments:segments", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-segment_policy--dst_segments--segments--name", "enforcement": "provider-schema", "group": "segment_policy.dst_segments.segments:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments:segments", "type": "requires"}], "schema_path": ["segment_policy", "dst_segments", "segments"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/dst_segments/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of references to Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy.dst_segments

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/)
- segment_policy.dst_segments

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dst segments.

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
dst_segments {
  # Configure direct properties listed below.
}
```

## Direct properties

- [segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/segments/): complete subsection reference.

## Next pages

- [segment_policy.dst_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/segments/)
- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
