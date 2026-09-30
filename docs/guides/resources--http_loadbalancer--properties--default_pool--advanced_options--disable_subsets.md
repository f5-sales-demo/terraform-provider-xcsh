---
page_title: "default_pool.advanced_options.disable_subsets"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.disable_subsets for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1200, "body_sha256": "sha256:a3834fab113668165a08aba032cc9894f78fb06d637f2889067540f4f23a334b", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_subsets", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_subsets", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--advanced_options--disable_subsets.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "disable_subsets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/disable_subsets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.disable_subsets for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool.advanced_options.disable_subsets

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- default_pool.advanced_options.disable_subsets

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable subsets. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_subsets = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
