---
page_title: "disable_trust_client_ip_headers"
subcategory: "Load Balancing"
description: "disable_trust_client_ip_headers for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1508, "body_sha256": "sha256:badfa62a265744ee9ea95ee32d08d3a94e73e7533f3807d36ae9c281a8485882", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:disable_trust_client_ip_headers", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:disable_trust_client_ip_headers", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--disable_trust_client_ip_headers.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_trust_client_ip_headers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/disable_trust_client_ip_headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_trust_client_ip_headers for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_trust_client_ip_headers

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- disable_trust_client_ip_headers

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_trust\_client\_ip\_headers, enable\_trust\_client\_ip\_headers; Default:
disable\_trust\_client\_ip\_headers\] Enable this option. Defaults to \`map\[\]\`. Server applies
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

OneOf alternatives in this subsection:

- [disable_trust_client_ip_headers](resources--http_loadbalancer--properties--disable_trust_client_ip_headers.md#section)
- [enable_trust_client_ip_headers](resources--http_loadbalancer--properties--enable_trust_client_ip_headers.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_trust_client_ip_headers = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
