---
page_title: "default_pool.use_tls.default_session_key_caching"
subcategory: "Load Balancing"
description: "default_pool.use_tls.default_session_key_caching for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1194, "body_sha256": "sha256:3eb9e1522d9fb04c015c089f3721a8246a71980c2fef13adf18da6ffe54c2561", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:default_session_key_caching", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:default_session_key_caching", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--use_tls--default_session_key_caching.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "use_tls", "default_session_key_caching"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/use_tls/default_session_key_caching/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.use_tls.default_session_key_caching for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool.use_tls.default_session_key_caching

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.use_tls](resources--http_loadbalancer--properties--default_pool--use_tls.md)
- default_pool.use_tls.default_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
default_session_key_caching = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.use_tls](resources--http_loadbalancer--properties--default_pool--use_tls.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
