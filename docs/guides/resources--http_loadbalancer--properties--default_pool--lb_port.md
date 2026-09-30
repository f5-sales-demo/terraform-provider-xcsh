---
page_title: "default_pool.lb_port"
subcategory: "Load Balancing"
description: "default_pool.lb_port for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 898, "body_sha256": "sha256:7b210b8ec6ebfa5c72f9ac8a73ec4e3a5e675adfda4c902c6d99763b1a99db83", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:lb_port", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:lb_port", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--lb_port.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "lb_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/lb_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.lb_port for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool.lb_port

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- default_pool.lb_port

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
lb_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
