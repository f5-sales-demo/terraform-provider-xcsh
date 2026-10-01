---
page_title: "rate_limit.custom_ip_allowed_list"
subcategory: "Load Balancing"
description: "rate_limit.custom_ip_allowed_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1222, "body_sha256": "sha256:297c269d61007f2c5c68352ef041680248c66f98793a9f35d87376058c74d46c", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit", "path": "docs/guides/data-sources--http_loadbalancer--properties--rate_limit--custom_ip_allowed_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "custom_ip_allowed_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/rate_limit/custom_ip_allowed_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.custom_ip_allowed_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.custom_ip_allowed_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [rate_limit](data-sources--http_loadbalancer--properties--rate_limit.md)
- rate_limit.custom_ip_allowed_list

<a id="section"></a>

Type: `"single"`. Computed.

IP Allowed list using existing ip\_prefix\_set objects.

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

- [rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--properties--rate_limit--custom_ip_allowed_list--rate_limiter_allowed_prefixes.md): complete subsection reference.

## Next pages

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--properties--rate_limit--custom_ip_allowed_list--rate_limiter_allowed_prefixes.md)
- [rate_limit](data-sources--http_loadbalancer--properties--rate_limit.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
