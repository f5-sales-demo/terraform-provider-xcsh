---
page_title: "origin_pools.pools.origin_servers.origin_servers.k8s_service"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers.k8s_service for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 6368, "body_sha256": "sha256:cbd56c028396f49a23a98c022910b403dc503d834b48fe722f5f96d33b4617db", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:inside_network", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:outside_network", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:site_locator", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:snat_pool", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:vk8s_networks"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers", "path": "documentation/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/index.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "k8s_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers.k8s_service for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pools.pools.origin_servers.origin_servers.k8s_service

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/)
- origin_pools.pools.origin_servers.origin_servers.k8s_service

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

## Direct properties

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/inside_network/): complete subsection reference.

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/outside_network/): complete subsection reference.

<a id="schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--protocol"></a>

### protocol property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--service_name"></a>

### service_name property

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/): complete subsection reference.

- [vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/vk8s_networks/): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/inside_network/)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/outside_network/)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/vk8s_networks/)
- [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
