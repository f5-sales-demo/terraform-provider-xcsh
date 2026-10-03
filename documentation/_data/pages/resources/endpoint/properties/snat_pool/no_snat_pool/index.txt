---
page_title: "snat_pool.no_snat_pool"
subcategory: "Networking"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["snat pool no snat pool"], "body_bytes": 1211, "body_sha256": "sha256:8a21f69b7704a8f8bdceb839cbaf8b43ef44548eca15f02f53a6d8e3d6d665a8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:snat_pool:no_snat_pool", "parent_id": "xcsh-docs:resources:endpoint:properties:snat_pool", "path": "documentation/resources/endpoint/properties/snat_pool/no_snat_pool/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2220221221301300-0010220322312023-2102300102101000-3101021221201033-1120021010113302-1101312221123121-2022101121133301-0232103001330100", "registry_path": "docs/guides/resources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["snat_pool", "no_snat_pool"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/snat_pool/no_snat_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["endpointCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# snat_pool.no_snat_pool

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/)
- snat_pool.no_snat_pool

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
