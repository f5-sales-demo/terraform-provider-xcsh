---
page_title: "default_pool.use_tls.use_host_header_as_sni"
subcategory: "Load Balancing"
description: "default_pool.use_tls.use_host_header_as_sni for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1184, "body_sha256": "sha256:10f4d7008779053eb3403682d4436b08432df24e3f41dda7366539b6d48da74b", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_host_header_as_sni", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_host_header_as_sni", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--use_tls--use_host_header_as_sni.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "use_tls", "use_host_header_as_sni"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_host_header_as_sni/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.use_tls.use_host_header_as_sni for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls.use_host_header_as_sni

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.use_tls](data-sources--http_loadbalancer--properties--default_pool--use_tls.md)
- default_pool.use_tls.use_host_header_as_sni

<a id="section"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.use_tls](data-sources--http_loadbalancer--properties--default_pool--use_tls.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
