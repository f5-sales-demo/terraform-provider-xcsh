---
page_title: "default_pool.origin_servers.consul_service"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.consul_service for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3419, "body_sha256": "sha256:b419f20e0483b55f6b6f103eb83d06eef73786df53f8057a5b9e721bbf1e8b90", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:inside_network", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:outside_network", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:site_locator", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:snat_pool"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "consul_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.consul_service for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.consul_service

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](data-sources--http_loadbalancer--properties--default_pool--origin_servers.md)
- default_pool.origin_servers.consul_service

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with HashiCorp Consul service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

## Direct properties

- [inside_network](data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service--inside_network.md): complete subsection reference.

- [outside_network](data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service--outside_network.md): complete subsection reference.

<a id="schema-default_pool--origin_servers--consul_service--service_name"></a>

### service_name property

Type: `"string"`. Computed.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Upstream description:

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [site_locator](data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service--site_locator.md): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service--snat_pool.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.consul_service.inside_network](data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service--inside_network.md)
- [default_pool.origin_servers.consul_service.outside_network](data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service--outside_network.md)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service--site_locator.md)
- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--properties--default_pool--origin_servers--consul_service--snat_pool.md)
- [default_pool.origin_servers](data-sources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
