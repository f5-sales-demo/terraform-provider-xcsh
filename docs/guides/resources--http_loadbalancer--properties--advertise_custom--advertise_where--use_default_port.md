---
page_title: "advertise_custom.advertise_where.use_default_port"
subcategory: "Load Balancing"
description: "advertise_custom.advertise_where.use_default_port for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1230, "body_sha256": "sha256:9b93630646e6269465d7862d6ed4d3f296ff24d2a61342b51f63609100dfcecf", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where", "path": "docs/guides/resources--http_loadbalancer--properties--advertise_custom--advertise_where--use_default_port.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "use_default_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/advertise_custom/advertise_where/use_default_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.use_default_port for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.use_default_port

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [advertise_custom](resources--http_loadbalancer--properties--advertise_custom.md)
- [advertise_custom.advertise_where](resources--http_loadbalancer--properties--advertise_custom--advertise_where.md)
- advertise_custom.advertise_where.use_default_port

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
use_default_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advertise_custom.advertise_where](resources--http_loadbalancer--properties--advertise_custom--advertise_where.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
