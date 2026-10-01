---
page_title: "trusted_clients.http_header"
subcategory: "Load Balancing"
description: "trusted_clients.http_header for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1415, "body_sha256": "sha256:3b511307f7b4aecd9d0d350f0ce04444366c3706677a18e3993347f90ee07e41", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients", "path": "docs/guides/resources--http_loadbalancer--properties--trusted_clients--http_header.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["trusted_clients", "http_header"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/trusted_clients/http_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "trusted_clients.http_header for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# trusted_clients.http_header

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [trusted_clients](resources--http_loadbalancer--properties--trusted_clients.md)
- trusted_clients.http_header

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

- [headers](resources--http_loadbalancer--properties--trusted_clients--http_header--headers.md): complete subsection reference.

## Next pages

- [trusted_clients.http_header.headers](resources--http_loadbalancer--properties--trusted_clients--http_header--headers.md)
- [trusted_clients](resources--http_loadbalancer--properties--trusted_clients.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
