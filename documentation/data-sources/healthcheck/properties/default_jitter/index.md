---
page_title: "default_jitter"
subcategory: "Monitoring"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default jitter"], "body_bytes": 1254, "body_sha256": "sha256:74b795f824e09b8b77013ac2e2fde6c30453db71f3d87360a58afc8fa5e11a65", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:healthcheck:properties:default_jitter", "parent_id": "xcsh-docs:data-sources:healthcheck:reference", "path": "documentation/data-sources/healthcheck/properties/default_jitter/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0322032010000220-3333013213012021-1212301213233322-2302110220211022-2021313303330311-0102320321001210-0030103103101313-3320111112233323", "registry_path": "docs/guides/data-sources--healthcheck--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_jitter"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/properties/default_jitter/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["healthcheckCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_jitter

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/properties/)
- default_jitter

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_jitter, jitter\_percent; Default: default\_jitter\] Configuration parameter for
default jitter.

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

- [default_jitter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/properties/default_jitter/#section)
- [jitter_percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/properties/#schema-jitter_percent)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
