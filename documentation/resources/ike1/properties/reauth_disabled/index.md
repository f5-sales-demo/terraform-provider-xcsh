---
page_title: "reauth_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["reauth disabled"], "body_bytes": 1623, "body_sha256": "sha256:b941fc30f04325f71c22556c33ce7da357d091c5eb9f6370331b1ea0df23bbca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:properties:reauth_disabled", "parent_id": "xcsh-docs:resources:ike1:reference", "path": "documentation/resources/ike1/properties/reauth_disabled/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0130213003021102-3221102331303201-3120330230322233-0201010222030121-1020031323123322-0332031113231101-3031133203003122-3322322301033323", "registry_path": "docs/guides/resources--ike1--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["reauth_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/properties/reauth_disabled/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["ike1CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_disabled

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/properties/)
- reauth_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

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

- [reauth_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/properties/reauth_disabled/#section)
- [reauth_timeout_days](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/properties/reauth_timeout_days/#section)
- [reauth_timeout_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/properties/reauth_timeout_hours/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
reauth_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/properties/)
- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/)
