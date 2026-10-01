---
page_title: "advanced_options.http1_config.header_transformation"
subcategory: "Load Balancing"
description: "advanced_options.http1_config.header_transformation for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 3053, "body_sha256": "sha256:c39724d0e6bac5f40f2c10ec8453f155431f88d162de5c40979269b31a6403d9", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation:default_header_transformation", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config", "path": "documentation/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["advanced_options", "http1_config", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.http1_config.header_transformation for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.http1_config.header_transformation

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/)
- [advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/)
- advanced_options.http1_config.header_transformation

<a id="section"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

## Direct properties

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/proper_case_header_transformation/): complete subsection reference.

## Next pages

- [advanced_options.http1_config.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/default_header_transformation/)
- [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/)
- [advanced_options.http1_config.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/proper_case_header_transformation/)
- [advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
