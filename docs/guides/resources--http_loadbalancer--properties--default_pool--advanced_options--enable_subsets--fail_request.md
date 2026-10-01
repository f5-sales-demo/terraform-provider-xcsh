---
page_title: "default_pool.advanced_options.enable_subsets.fail_request"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.enable_subsets.fail_request for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1418, "body_sha256": "sha256:c003bf2b9757eb38e1d2a02d44cc1efae2341a3c04fb1f0c1ececd8a9ebe860f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--fail_request.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets", "fail_request"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/fail_request/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.enable_subsets.fail_request for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets.fail_request

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets.md)
- default_pool.advanced_options.enable_subsets.fail_request

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fail request.

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

Terraform syntax:

```terraform
fail_request = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
