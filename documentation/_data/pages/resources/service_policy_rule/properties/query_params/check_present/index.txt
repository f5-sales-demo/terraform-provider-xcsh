---
page_title: "query_params.check_present"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["query params check present"], "body_bytes": 1310, "body_sha256": "sha256:6794bf32ea68288c0f274b61109e7925d0e89bdcfb4fc553c1d8b358a7f83338", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:query_params:check_present", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:query_params", "path": "documentation/resources/service_policy_rule/properties/query_params/check_present/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2213203220132122-1123311113310120-0200132121322222-1320110200011231-3213213300100303-2100122210033231-1120321110122000-1001021132121000", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["query_params", "check_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/query_params/check_present/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# query_params.check_present

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/)
- query_params.check_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
