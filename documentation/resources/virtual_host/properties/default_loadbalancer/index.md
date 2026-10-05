---
page_title: "default_loadbalancer"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default loadbalancer"], "body_bytes": 1633, "body_sha256": "sha256:43e5b8bb52c91db3eb72ea9b43a8a18fe8a48dbe4882156ab1a8b9ad5d794985", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:default_loadbalancer", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "documentation/resources/virtual_host/properties/default_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3121022200033310-1220301130221312-2021030331023200-1312300232120133-1231101321130312-2222300001330221-1303311113210322-3100130111300312", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_loadbalancer"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/default_loadbalancer/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_loadbalancer

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- default_loadbalancer

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_loadbalancer, non\_default\_loadbalancer; Default: default\_loadbalancer\]
Configuration parameter for default loadbalancer.

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

- [default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/default_loadbalancer/#section)
- [non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/non_default_loadbalancer/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_loadbalancer = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
