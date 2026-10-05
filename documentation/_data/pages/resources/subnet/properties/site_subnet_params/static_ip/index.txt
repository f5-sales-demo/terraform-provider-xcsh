---
page_title: "site_subnet_params.static_ip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["site subnet params static ip"], "body_bytes": 1219, "body_sha256": "sha256:a6469d65d544cc1484ed4bc6779271c4c0d1377a3fe69ee7cb330bf407ea31d8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:site_subnet_params:static_ip", "parent_id": "xcsh-docs:resources:subnet:properties:site_subnet_params", "path": "documentation/resources/subnet/properties/site_subnet_params/static_ip/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0000120300121031-2031001210121031-2323022232221120-0013310013233000-3222123111101312-0122021331012020-3112332211122123-1012123321213200", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_subnet_params", "static_ip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/site_subnet_params/static_ip/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params.static_ip

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- [site_subnet_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/)
- site_subnet_params.static_ip

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
static_ip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [site_subnet_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/)
- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
