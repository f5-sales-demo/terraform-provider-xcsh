---
page_title: "every_day"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["every day"], "body_bytes": 1311, "body_sha256": "sha256:b81e3a4ae587ebfec774f7a6ed963063f2e0e14df3f556912e45289ca4989088", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:every_day", "parent_id": "xcsh-docs:data-sources:api_testing:reference", "path": "documentation/data-sources/api_testing/properties/every_day/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2332330000101223-2332033331213132-1301121202210033-3103023332020211-1332031203333033-1103110000210321-0012010013012201-0313032122013212", "registry_path": "docs/guides/data-sources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["every_day"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/every_day/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["api_testingCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# every_day

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- every_day

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: every\_day, every\_month, every\_week\] Enable this option

Additional upstream details:

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

- [every_day](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/every_day/#section)
- [every_month](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/every_month/#section)
- [every_week](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/every_week/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
