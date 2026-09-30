---
page_title: "blocked_clients.http_header"
subcategory: "Load Balancing"
description: "blocked_clients.http_header for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1316, "body_sha256": "sha256:051c727ce40a6343502b429e3f5dd5cffebfe5cea2c0dadb3ef67feea953911e", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients", "path": "docs/guides/resources--http_loadbalancer--properties--blocked_clients--http_header.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_clients", "http_header"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/blocked_clients/http_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_clients.http_header for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# blocked_clients.http_header

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [blocked_clients](resources--http_loadbalancer--properties--blocked_clients.md)
- blocked_clients.http_header

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
```

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
http_header {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](resources--http_loadbalancer--properties--blocked_clients--http_header--headers.md): complete subsection reference.

## Next pages

- [blocked_clients.http_header.headers](resources--http_loadbalancer--properties--blocked_clients--http_header--headers.md)
- [blocked_clients](resources--http_loadbalancer--properties--blocked_clients.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
