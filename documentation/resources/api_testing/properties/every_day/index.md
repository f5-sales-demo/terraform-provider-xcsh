---
page_title: "every_day"
subcategory: ""
description: "every_day for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 1594, "body_sha256": "sha256:1a82f6e531dc110e1c8ab0aeddb5815587473a1d57dad3c19679f686bfb3502f", "child_ids": [], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:every_day", "parent_id": "xcsh-docs:resources:api_testing:reference", "path": "documentation/resources/api_testing/properties/every_day/index.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["every_day"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/every_day/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "every_day for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
