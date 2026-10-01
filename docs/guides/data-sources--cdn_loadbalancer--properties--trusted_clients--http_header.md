---
page_title: "trusted_clients.http_header"
subcategory: "Load Balancing"
description: "trusted_clients.http_header for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1161, "body_sha256": "sha256:f86dadade1c6417c14ebb90d7776bc355929aaaa64b79ec0d58359f269cb65cb", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:trusted_clients:http_header", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:trusted_clients:http_header:headers"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:trusted_clients:http_header", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:trusted_clients", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--trusted_clients--http_header.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["trusted_clients", "http_header"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/trusted_clients/http_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "trusted_clients.http_header for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# trusted_clients.http_header

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [trusted_clients](data-sources--cdn_loadbalancer--properties--trusted_clients.md)
- trusted_clients.http_header

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

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

- [headers](data-sources--cdn_loadbalancer--properties--trusted_clients--http_header--headers.md): complete subsection reference.

## Next pages

- [trusted_clients.http_header.headers](data-sources--cdn_loadbalancer--properties--trusted_clients--http_header--headers.md)
- [trusted_clients](data-sources--cdn_loadbalancer--properties--trusted_clients.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
