---
page_title: "do_not_advertise"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["do not advertise"], "body_bytes": 1500, "body_sha256": "sha256:7615998eef2feecbe7a98ce72f283260e83816228aa798567a3bb20ae0b7de84", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:do_not_advertise", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/do_not_advertise/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1300300333220210-2321100332123100-3100222323020132-2000023232233210-0122031023212311-0203230120332013-3020033130003013-3323122123123232", "registry_path": "docs/guides/resources--proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["do_not_advertise"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/do_not_advertise/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# do_not_advertise

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- do_not_advertise

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: do\_not\_advertise, site\_virtual\_sites\] Configuration parameter for do not advertise.

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

- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/do_not_advertise/#section)
- [site_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
do_not_advertise = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
