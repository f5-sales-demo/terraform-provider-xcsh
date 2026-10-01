---
page_title: "rate_limit.policies"
subcategory: "Load Balancing"
description: "rate_limit.policies for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1057, "body_sha256": "sha256:4272a1c62f2f0215f9e5fd8f428638bc6504bb112defd522c65fac4b6f2f5cc7", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:policies", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:policies:policies"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:policies", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit", "path": "docs/guides/data-sources--http_loadbalancer--properties--rate_limit--policies.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/rate_limit/policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.policies for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.policies

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [rate_limit](data-sources--http_loadbalancer--properties--rate_limit.md)
- rate_limit.policies

<a id="section"></a>

Type: `"single"`. Computed.

List of rate limiter policies to be applied.

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

- [policies](data-sources--http_loadbalancer--properties--rate_limit--policies--policies.md): complete subsection reference.

## Next pages

- [rate_limit.policies.policies](data-sources--http_loadbalancer--properties--rate_limit--policies--policies.md)
- [rate_limit](data-sources--http_loadbalancer--properties--rate_limit.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
