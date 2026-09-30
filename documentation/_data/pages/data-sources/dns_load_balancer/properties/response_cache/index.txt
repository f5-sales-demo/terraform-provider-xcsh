---
page_title: "response_cache"
subcategory: "DNS"
description: "response_cache for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 2178, "body_sha256": "sha256:d59ff9df8ec58e74910808ce06b5c4b36a089708817a908a17c10fd88b19ca65", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:disable_spec", "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:response_cache_parameters"], "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "path": "documentation/data-sources/dns_load_balancer/properties/response_cache/index.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["response_cache"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/response_cache/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cache for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# response_cache

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- response_cache

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for response cache.

Upstream description:

Response Cache x-required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_cache_parameters_choice": "[\"default_response_cache_parameters\",\"disable\",\"response_cache_parameters\"]"
}
```

## Direct properties

- [default_response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/disable_spec/): complete subsection reference.

- [response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/): complete subsection reference.

## Next pages

- [response_cache.default_response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/)
- [response_cache.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/disable_spec/)
- [response_cache.response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
