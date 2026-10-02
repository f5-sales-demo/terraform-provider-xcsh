---
page_title: "disabled"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disabled"], "body_bytes": 1406, "body_sha256": "sha256:28d5022f3367eceda4de917f535e780699642af2e04a290d4c75e2c76ed7cfaf", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_k8s:properties:disabled", "parent_id": "xcsh-docs:data-sources:virtual_k8s:reference", "path": "documentation/data-sources/virtual_k8s/properties/disabled/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0103231212031130-1111030212223030-0201311232122103-3011222101321133-1002031202300101-1011013121010321-3130231201312333-0202121011021113", "registry_path": "docs/guides/data-sources--virtual_k8s--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_k8s/properties/disabled/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disabled

Breadcrumbs:

- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/)
- disabled

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disabled, isolated\] Enable this option

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

OneOf alternatives in this subsection:

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/disabled/#section)
- [isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/isolated/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/)
- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/)
