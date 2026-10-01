---
page_title: "api_rate_limit"
subcategory: "Load Balancing"
description: "api_rate_limit for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3340, "body_sha256": "sha256:dfdb5e2105e5d0792cc00b86fc8a32e7fc09dce032e4dff93d943d5a84b49321", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:ip_allowed_list", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_rate_limit.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- api_rate_limit

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\] Path-
or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and choose
inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Upstream description:

Path- or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and
choose inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md#section)
- [disable_rate_limit](data-sources--http_loadbalancer--properties--disable_rate_limit.md#section)
- [rate_limit](data-sources--http_loadbalancer--properties--rate_limit.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [api_endpoint_rules](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md): complete subsection reference.

- [bypass_rate_limiting_rules](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules.md): complete subsection reference.

- [custom_ip_allowed_list](data-sources--http_loadbalancer--properties--api_rate_limit--custom_ip_allowed_list.md): complete subsection reference.

- [ip_allowed_list](data-sources--http_loadbalancer--properties--api_rate_limit--ip_allowed_list.md): complete subsection reference.

- [no_ip_allowed_list](data-sources--http_loadbalancer--properties--api_rate_limit--no_ip_allowed_list.md): complete subsection reference.

- [server_url_rules](data-sources--http_loadbalancer--properties--api_rate_limit--server_url_rules.md): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules.md)
- [api_rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--properties--api_rate_limit--custom_ip_allowed_list.md)
- [api_rate_limit.ip_allowed_list](data-sources--http_loadbalancer--properties--api_rate_limit--ip_allowed_list.md)
- [api_rate_limit.no_ip_allowed_list](data-sources--http_loadbalancer--properties--api_rate_limit--no_ip_allowed_list.md)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--properties--api_rate_limit--server_url_rules.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
