---
page_title: "isolated"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["isolated"], "body_bytes": 815, "body_sha256": "sha256:474e5f128ede3cb1d1241d9d2811cc3d38d2511aa4e762c1e9457304072ab023", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_k8s:properties:isolated", "parent_id": "xcsh-docs:resources:virtual_k8s:reference", "path": "documentation/resources/virtual_k8s/properties/isolated/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0231101212002130-3302021000230012-0022123212210330-1033002222133301-0201201030203202-3321311310222323-0323330100103303-0202330120221111", "registry_path": "docs/guides/resources--virtual_k8s--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["isolated"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_k8s/properties/isolated/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# isolated

Breadcrumbs:

- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/properties/)
- isolated

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
isolated = {}
```

This is an empty object or choice marker. It has no direct properties.
