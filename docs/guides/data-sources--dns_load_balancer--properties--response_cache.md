---
page_title: "response_cache"
subcategory: "DNS"
description: "response_cache for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1769, "body_sha256": "sha256:e36aad42bfbdd8532456be4a44b560eb0823e4a1248cc117f2b84ec0dad95b67", "canonical_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:disable_spec", "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:response_cache_parameters"], "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "path": "docs/guides/data-sources--dns_load_balancer--properties--response_cache.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["response_cache"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/response_cache/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cache for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cache

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md)
- [Property reference](data-sources--dns_load_balancer--reference.md)
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

- [default_response_cache_parameters](data-sources--dns_load_balancer--properties--response_cache--default_response_cache_parameters.md): complete subsection reference.

- [disable_spec](data-sources--dns_load_balancer--properties--response_cache--disable_spec.md): complete subsection reference.

- [response_cache_parameters](data-sources--dns_load_balancer--properties--response_cache--response_cache_parameters.md): complete subsection reference.

## Next pages

- [response_cache.default_response_cache_parameters](data-sources--dns_load_balancer--properties--response_cache--default_response_cache_parameters.md)
- [response_cache.disable_spec](data-sources--dns_load_balancer--properties--response_cache--disable_spec.md)
- [response_cache.response_cache_parameters](data-sources--dns_load_balancer--properties--response_cache--response_cache_parameters.md)
- [Property reference](data-sources--dns_load_balancer--reference.md)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md)
