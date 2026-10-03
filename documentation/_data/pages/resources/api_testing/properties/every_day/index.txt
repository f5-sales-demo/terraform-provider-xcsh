---
page_title: "every_day"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["every day"], "body_bytes": 1594, "body_sha256": "sha256:1a82f6e531dc110e1c8ab0aeddb5815587473a1d57dad3c19679f686bfb3502f", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:every_day", "parent_id": "xcsh-docs:resources:api_testing:reference", "path": "documentation/resources/api_testing/properties/every_day/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0110133330203132-0300210311112100-1220023032121103-3000313102030220-3123012332002301-1030201222002031-3231121132130330-1120323230220100", "registry_path": "docs/guides/resources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["every_day"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/every_day/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# every_day

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- every_day

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: every\_day, every\_month, every\_week\] Enable this option

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

- [every_day](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/every_day/#section)
- [every_month](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/every_month/#section)
- [every_week](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/every_week/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
every_day = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
