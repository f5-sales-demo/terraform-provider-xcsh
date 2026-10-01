---
page_title: "default_pool.origin_servers.private_ip"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.private_ip for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3593, "body_sha256": "sha256:d3056bb31a6ad0a9406a4352110da6c9da98edede1a7a92521cfc8e88125b826", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:inside_network", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:outside_network", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:segment", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "private_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.private_ip for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.private_ip

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](data-sources--http_loadbalancer--properties--default_pool--origin_servers.md)
- default_pool.origin_servers.private_ip

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with private or public IP address and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

## Direct properties

- [inside_network](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--inside_network.md): complete subsection reference.

<a id="schema-default_pool--origin_servers--private_ip--ip"></a>

### ip property

Type: `"string"`. Computed.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--outside_network.md): complete subsection reference.

- [segment](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--segment.md): complete subsection reference.

- [site_locator](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--site_locator.md): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--snat_pool.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.private_ip.inside_network](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--inside_network.md)
- [default_pool.origin_servers.private_ip.outside_network](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--outside_network.md)
- [default_pool.origin_servers.private_ip.segment](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--segment.md)
- [default_pool.origin_servers.private_ip.site_locator](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--site_locator.md)
- [default_pool.origin_servers.private_ip.snat_pool](data-sources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--snat_pool.md)
- [default_pool.origin_servers](data-sources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
