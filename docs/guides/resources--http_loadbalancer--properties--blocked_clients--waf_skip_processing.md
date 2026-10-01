---
page_title: "blocked_clients.waf_skip_processing"
subcategory: "Load Balancing"
description: "blocked_clients.waf_skip_processing for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1051, "body_sha256": "sha256:92503fc6fa7095723495ed9173d24eb78dc1839940f1a85d8c952d619b9d3514", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:waf_skip_processing", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:waf_skip_processing", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients", "path": "docs/guides/resources--http_loadbalancer--properties--blocked_clients--waf_skip_processing.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_clients", "waf_skip_processing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/blocked_clients/waf_skip_processing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_clients.waf_skip_processing for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_clients.waf_skip_processing

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [blocked_clients](resources--http_loadbalancer--properties--blocked_clients.md)
- blocked_clients.waf_skip_processing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
waf_skip_processing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [blocked_clients](resources--http_loadbalancer--properties--blocked_clients.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
