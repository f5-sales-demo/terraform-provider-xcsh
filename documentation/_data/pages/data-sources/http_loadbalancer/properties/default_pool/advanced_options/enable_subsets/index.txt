---
page_title: "default_pool.advanced_options.enable_subsets"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.enable_subsets for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3168, "body_sha256": "sha256:6e978ac6d1b4e124baa557f94d09f1a53b5c488fcf93ff5c0a76e1b33b3180fc", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:any_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.enable_subsets for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/)
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

- [any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/any_endpoint/): complete subsection reference.

- [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/endpoint_subsets/): complete subsection reference.

- [fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/fail_request/): complete subsection reference.

## Next pages

- [default_pool.advanced_options.enable_subsets.any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/any_endpoint/)
- [default_pool.advanced_options.enable_subsets.default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/)
- [default_pool.advanced_options.enable_subsets.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/endpoint_subsets/)
- [default_pool.advanced_options.enable_subsets.fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/fail_request/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
