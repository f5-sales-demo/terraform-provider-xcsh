---
page_title: "default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1702, "body_sha256": "sha256:d9afb28f8e1e51e9d4aa9b07cc3c0c5c598dcb5bf89117d467ec5efe225b8e83", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config--header_transformation--proper_case_header_transformation.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation", "proper_case_header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/proper_case_header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [default_pool.advanced_options.http1_config](resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config.md)
- [default_pool.advanced_options.http1_config.header_transformation](resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config--header_transformation.md)
- default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.advanced_options.http1_config.header_transformation](resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config--header_transformation.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
