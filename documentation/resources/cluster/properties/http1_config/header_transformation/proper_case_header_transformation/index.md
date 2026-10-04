---
page_title: "http1_config.header_transformation.proper_case_header_transformation"
subcategory: ""
description: "Transform HTTP header names to proper case when explicit transformation is required."
xcsh_docs: {"aliases": ["http1 config header transformation proper case header transformation"], "body_bytes": 1495, "body_sha256": "sha256:6ed7a39b8e046bcaec5d89acfc284c619ddeafae4a7c91786433d26029e35d71", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:proper_case_header_transformation", "parent_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation", "path": "documentation/resources/cluster/properties/http1_config/header_transformation/proper_case_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2231100321330203-1210202101302300-1303003222032121-0312313011333021-3211020021011222-2122211033331200-2002101332232122-1202020021103120", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http1_config", "header_transformation", "proper_case_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http1_config/header_transformation/proper_case_header_transformation/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Transform HTTP header names to proper case when explicit transformation is required.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config.header_transformation.proper_case_header_transformation

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/)
- [http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/)
- http1_config.header_transformation.proper_case_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
