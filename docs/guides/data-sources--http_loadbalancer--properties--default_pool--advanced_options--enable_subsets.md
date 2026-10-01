---
page_title: "default_pool.advanced_options.enable_subsets"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.enable_subsets for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2478, "body_sha256": "sha256:ae2827d73ff42156bc85f1ba7e3142eb0a1fc4d5c1a7a55c20587d6af71771de", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:any_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.enable_subsets for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](data-sources--http_loadbalancer--properties--default_pool--advanced_options.md)
- default_pool.advanced_options.enable_subsets

<a id="section"></a>

Type: `"single"`. Computed.

Configure subset OPTIONS for origin pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fallback_policy_choice": "[\"any_endpoint\",\"default_subset\",\"fail_request\"]"
}
```

## Direct properties

- [any_endpoint](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--any_endpoint.md): complete subsection reference.

- [default_subset](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--default_subset.md): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--endpoint_subsets.md): complete subsection reference.

- [fail_request](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--fail_request.md): complete subsection reference.

## Next pages

- [default_pool.advanced_options.enable_subsets.any_endpoint](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--any_endpoint.md)
- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--default_subset.md)
- [default_pool.advanced_options.enable_subsets.endpoint_subsets](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--endpoint_subsets.md)
- [default_pool.advanced_options.enable_subsets.fail_request](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--fail_request.md)
- [default_pool.advanced_options](data-sources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
