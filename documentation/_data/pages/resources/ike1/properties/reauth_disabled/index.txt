---
page_title: "reauth_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["reauth disabled"], "body_bytes": 1623, "body_sha256": "sha256:b941fc30f04325f71c22556c33ce7da357d091c5eb9f6370331b1ea0df23bbca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:properties:reauth_disabled", "parent_id": "xcsh-docs:resources:ike1:reference", "path": "documentation/resources/ike1/properties/reauth_disabled/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0130213003021102-3221102331303201-3120330230322233-0201010222030121-1020031323123322-0332031113231101-3031133203003122-3322322301033323", "registry_path": "docs/guides/resources--ike1--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["reauth_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/properties/reauth_disabled/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike1CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
