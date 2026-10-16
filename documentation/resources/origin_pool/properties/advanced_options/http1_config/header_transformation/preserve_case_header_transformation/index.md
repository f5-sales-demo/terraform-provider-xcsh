---
page_title: "advanced_options.http1_config.header_transformation.preserve_case_header_transformation"
subcategory: "Load Balancing"
description: "Preserve HTTP header-name case when upstream case must remain unchanged."
xcsh_docs: {"aliases": ["advanced options http1 config header transformation preserve case header transformation"], "body_bytes": 1448, "body_sha256": "sha256:2e14e65f4aafcbf6d63ee9c0d312dd48676f26a570b129267911d78150f4446d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation", "path": "documentation/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0330021003111221-1332122300203310-0123000012003123-0223121210122230-3231211101203232-2312110300330112-2123001120301023-3030102300123220", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "http1_config", "header_transformation", "preserve_case_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Preserve HTTP header-name case when upstream case must remain unchanged.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["origin_poolCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.http1_config.header_transformation.preserve_case_header_transformation

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- [advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/)
- [advanced_options.http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/)
- advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.
