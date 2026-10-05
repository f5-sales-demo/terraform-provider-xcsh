---
page_title: "disable_gpu"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable gpu"], "body_bytes": 1645, "body_sha256": "sha256:b0bbe4b151025dd1d54aade2d11fff11221f86eee7d73bdbe342fc55b67a0aaa", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:disable_gpu", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "documentation/data-sources/voltstack_site/properties/disable_gpu/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2013133233301023-3002101131132333-0132130032032231-1201022321001011-2201310323103301-3212311322221321-3312032120220131-0232123003003001", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_gpu"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/disable_gpu/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_gpu

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- disable_gpu

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_gpu, enable\_gpu, enable\_vgpu; Default: disable\_gpu\] Configuration parameter
for disable gpu.

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

- [disable_gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/disable_gpu/#section)
- [enable_gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/enable_gpu/#section)
- [enable_vgpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/enable_vgpu/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
