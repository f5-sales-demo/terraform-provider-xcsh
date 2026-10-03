---
page_title: "dualstack"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["dualstack"], "body_bytes": 1566, "body_sha256": "sha256:62bfa64fd1c05f7bec721833edd4fdab98e9abc01d1b6aa9b57022bc4d6bec4a", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:dualstack", "parent_id": "xcsh-docs:data-sources:advertise_policy:reference", "path": "documentation/data-sources/advertise_policy/properties/dualstack/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1300321132021101-2312230222232310-0310313232330230-3132121011011210-3012022120112200-0112311123223231-0322331101121122-2111101311031301", "registry_path": "docs/guides/data-sources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dualstack"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/dualstack/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dualstack

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
- dualstack

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: dualstack, ipv4, ipv6\] Enable this option

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

- [dualstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/dualstack/#section)
- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/ipv4/#section)
- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/ipv6/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
