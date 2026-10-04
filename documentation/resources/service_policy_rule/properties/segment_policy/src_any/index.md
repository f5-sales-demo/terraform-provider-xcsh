---
page_title: "segment_policy.src_any"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["segment policy src any"], "body_bytes": 1280, "body_sha256": "sha256:6be338a69bc512d23d2d1af94ddab2e46fcbfce6aad4d831bc67690d76c3837a", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_any", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "path": "documentation/resources/service_policy_rule/properties/segment_policy/src_any/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1021033100031322-0210332232032110-0310012210003102-3310330313201232-1013122333100201-1010021111102100-0221022030003001-1011013331111111", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_policy", "src_any"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/src_any/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy.src_any

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/)
- segment_policy.src_any

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
src_any = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
