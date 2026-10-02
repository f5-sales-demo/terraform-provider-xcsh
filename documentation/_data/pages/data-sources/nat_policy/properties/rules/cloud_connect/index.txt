---
page_title: "rules.cloud_connect"
subcategory: ""
description: "Reference to Cloud connect Object."
xcsh_docs: {"aliases": ["rules cloud connect"], "body_bytes": 1370, "body_sha256": "sha256:31fb4ef4e3bc5d134f91e44be1f249f003d1123a4c34a3420b4b42ee306c4654", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:cloud_connect:refs"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:cloud_connect", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "path": "documentation/data-sources/nat_policy/properties/rules/cloud_connect/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2002332211313233-0323320313100303-1101121333232222-0221010310231023-0110021131222030-1000023213330201-0023023212120113-3230132223300030", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "cloud_connect"], "schema_version": 1, "sections": [{"aliases": ["refs"], "anchor": "section", "description": "Reference to Cloud Connect Object.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:cloud_connect:refs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "cloud_connect", "refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/cloud_connect/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Reference to Cloud connect Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nat_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.cloud_connect

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- rules.cloud_connect

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for cloud connect.

Upstream description:

Reference to Cloud connect Object.

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

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/cloud_connect/refs/): complete subsection reference.

## Next pages

- [rules.cloud_connect.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/cloud_connect/refs/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
