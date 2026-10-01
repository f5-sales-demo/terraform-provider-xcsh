---
page_title: "single_lb_app.disable_discovery"
subcategory: "Load Balancing"
description: "single_lb_app.disable_discovery for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1061, "body_sha256": "sha256:c542bd9b6868594fde4993b57e63899bef2b0abaea3d513b43e889f69e8f3d6d", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_discovery", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_discovery", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app", "path": "docs/guides/resources--http_loadbalancer--properties--single_lb_app--disable_discovery.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "disable_discovery"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/disable_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.disable_discovery for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.disable_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- single_lb_app.disable_discovery

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable discovery.

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
disable_discovery = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
