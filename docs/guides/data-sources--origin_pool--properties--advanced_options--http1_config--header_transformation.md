---
page_title: "advanced_options.http1_config.header_transformation"
subcategory: "Load Balancing"
description: "advanced_options.http1_config.header_transformation for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2459, "body_sha256": "sha256:2d6c50697692893e16c44188f7202034d02edc7da413c881ceca736fb92f8060", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation:default_header_transformation", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config", "path": "docs/guides/data-sources--origin_pool--properties--advanced_options--http1_config--header_transformation.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "http1_config", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.http1_config.header_transformation for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.http1_config.header_transformation

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- [advanced_options.http1_config](data-sources--origin_pool--properties--advanced_options--http1_config.md)
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

- [default_header_transformation](data-sources--origin_pool--properties--advanced_options--http1_config--header_transformation--default_header_transformation.md): complete subsection reference.

- [preserve_case_header_transformation](data-sources--origin_pool--properties--advanced_options--http1_config--header_transformation--preserve_case_header_transformation.md): complete subsection reference.

- [proper_case_header_transformation](data-sources--origin_pool--properties--advanced_options--http1_config--header_transformation--proper_case_header_transformation.md): complete subsection reference.

## Next pages

- [advanced_options.http1_config.header_transformation.default_header_transformation](data-sources--origin_pool--properties--advanced_options--http1_config--header_transformation--default_header_transformation.md)
- [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](data-sources--origin_pool--properties--advanced_options--http1_config--header_transformation--preserve_case_header_transformation.md)
- [advanced_options.http1_config.header_transformation.proper_case_header_transformation](data-sources--origin_pool--properties--advanced_options--http1_config--header_transformation--proper_case_header_transformation.md)
- [advanced_options.http1_config](data-sources--origin_pool--properties--advanced_options--http1_config.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
