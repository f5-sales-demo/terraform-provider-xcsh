---
page_title: "api_rate_limit.server_url_rules.client_matcher.asn_matcher"
subcategory: "Load Balancing"
description: "api_rate_limit.server_url_rules.client_matcher.asn_matcher for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1610, "body_sha256": "sha256:1cf62ff1e8af72678ed70b6f993541b19cf9b7e6cb462c55be24aa0ddd792d77", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher--asn_matcher.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.server_url_rules.client_matcher.asn_matcher for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.client_matcher.asn_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_rate_limit](data-sources--cdn_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules.md)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher.md)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

## Direct properties

- [asn_sets](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher--asn_matcher--asn_sets.md): complete subsection reference.

## Next pages

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher--asn_matcher--asn_sets.md)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
