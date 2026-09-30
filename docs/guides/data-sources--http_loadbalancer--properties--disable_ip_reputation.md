---
page_title: "disable_ip_reputation"
subcategory: "Load Balancing"
description: "disable_ip_reputation for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1247, "body_sha256": "sha256:068dd2594c98d558475044c58932675314fa421f9dc178c774d46e0fb8cca91b", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:disable_ip_reputation", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:disable_ip_reputation", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--disable_ip_reputation.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_ip_reputation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/disable_ip_reputation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_ip_reputation for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_ip_reputation

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- disable_ip_reputation

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](data-sources--http_loadbalancer--properties--disable_ip_reputation.md#section)
- [enable_ip_reputation](data-sources--http_loadbalancer--properties--enable_ip_reputation.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
