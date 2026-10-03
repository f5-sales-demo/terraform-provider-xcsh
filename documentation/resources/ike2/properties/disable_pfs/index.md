---
page_title: "disable_pfs"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable pfs"], "body_bytes": 1050, "body_sha256": "sha256:99efea411d17c18ea5733926b4c62738d1be7c7cbc8f1ea5a75e2441c64d8985", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike2:properties:disable_pfs", "parent_id": "xcsh-docs:resources:ike2:reference", "path": "documentation/resources/ike2/properties/disable_pfs/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3123101102221122-2012103322303011-1203112322120320-3320111010111123-2232303211033321-3302202122311112-0322000001010023-2203320122033312", "registry_path": "docs/guides/resources--ike2--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_pfs"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/properties/disable_pfs/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_pfs

Breadcrumbs:

- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/)
- disable_pfs

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable pfs.

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
disable_pfs = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/)
- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
