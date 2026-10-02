---
page_title: "dynamic_reverse_proxy"
subcategory: ""
description: "In this mode of proxy, virtual host will resolve the destination endpoint dynamically. The dynamic resolution is done using a predefined field in the request. This predefined field depends on the ProxyType configured on the Virtual Host. For HTTP traffic, i.e. With ProxyType as HTTP_PROXY or HTTPS_PROXY, virtual host"
xcsh_docs: {"aliases": ["dynamic reverse proxy"], "body_bytes": 10537, "body_sha256": "sha256:3c75634aa8148d7ea0dbb5e111346007feb814d6f094ad5da27c518812c81275", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:dynamic_reverse_proxy:resolution_network"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:dynamic_reverse_proxy", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "documentation/data-sources/virtual_host/properties/dynamic_reverse_proxy/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2320322030133330-2313030201332123-3133200313313203-0300032102221333-1001120222131112-1330331333122030-2312013001023021-0010002323330332", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_reverse_proxy"], "schema_version": 1, "sections": [{"aliases": ["connection timeout", "duration", "operation timeout"], "anchor": "schema-dynamic_reverse_proxy--connection_timeout", "description": "The timeout for new network connections to upstream server. This is specified in milliseconds. The default value is 2000 (2 seconds)", "document_id": "xcsh-docs:data-sources:virtual_host:properties:dynamic_reverse_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_reverse_proxy", "connection_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["resolution network"], "anchor": "section", "description": "Reference to virtual network where the endpoint is resolved. Reference is valid only when the network type is VIRTUAL_NETWORK_PER_SITE or VIRTUAL_NETWORK_GLOBAL. It is ignored for all other network types.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:dynamic_reverse_proxy:resolution_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["dynamic_reverse_proxy", "resolution_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["resolution network type"], "anchor": "schema-dynamic_reverse_proxy--resolution_network_type", "description": "Different types of virtual networks understood by the system Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual", "document_id": "xcsh-docs:data-sources:virtual_host:properties:dynamic_reverse_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_reverse_proxy", "resolution_network_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["resolve endpoint dynamically"], "anchor": "schema-dynamic_reverse_proxy--resolve_endpoint_dynamically", "description": "X-example : true In this mode of proxy, virtual host will resolve the destination endpoint dynamically. The dynamic resolution is done using a predefined field in the request. This predefined field depends on the ProxyType configured on the Virtual Host. For HTTP traffic, i.e. With ProxyType as HTTP_PROXY or", "document_id": "xcsh-docs:data-sources:virtual_host:properties:dynamic_reverse_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_reverse_proxy", "resolve_endpoint_dynamically"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/dynamic_reverse_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "In this mode of proxy, virtual host will resolve the destination endpoint dynamically. The dynamic resolution is done using a predefined field in the request. This predefined field depends on the ProxyType configured on the Virtual Host. For HTTP traffic, i.e. With ProxyType as HTTP_PROXY or HTTPS_PROXY, virtual host", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_reverse_proxy

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- dynamic_reverse_proxy

<a id="section"></a>

Type: `"single"`. Computed.

In this mode of proxy, virtual host will resolve the destination endpoint dynamically. The dynamic
resolution is done using a predefined field in the request. This predefined field depends on the
ProxyType configured on the Virtual Host.

Upstream description:

In this mode of proxy, virtual host will resolve the destination endpoint dynamically.

The dynamic resolution is done using a predefined field in the request. This predefined field
depends on the ProxyType configured on the Virtual Host.

For HTTP traffic, i.e. With ProxyType as HTTP\_PROXY or HTTPS\_PROXY, virtual host will use the
"HOST" HTTP header from the request and perform DNS resolution to select destination endpoint.

For TCP traffic with SNI, (If the ProxyType is TCP\_PROXY\_WITH\_SNI), virtual host will perform DNS
resolution using the SNI.

The DNS resolution is performed in the virtual network specified in outside\_network\_type or
outside\_network

In both modes of operation(either using Host header or SNI), the DNS resolution could return
multiple addresses. First IPv4 address from such returned list is used as endpoint for the request.
The DNS response is cached for 60s by default.

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

<a id="schema-dynamic_reverse_proxy--connection_timeout"></a>

### connection_timeout property

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [resolution_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/dynamic_reverse_proxy/resolution_network/): complete subsection reference.

<a id="schema-dynamic_reverse_proxy--resolution_network_type"></a>

### resolution_network_type property

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-dynamic_reverse_proxy--resolve_endpoint_dynamically"></a>

### resolve_endpoint_dynamically property

Type: `"bool"`. Computed.

X-example : true In this mode of proxy, virtual host will resolve the destination endpoint
dynamically. The dynamic resolution is done using a predefined field in the request. This predefined
field depends on the ProxyType configured on the Virtual Host.

Upstream description:

X-example : true In this mode of proxy, virtual host will resolve the destination endpoint
dynamically.

The dynamic resolution is done using a predefined field in the request. This predefined field
depends on the ProxyType configured on the Virtual Host.

For HTTP traffic, i.e. With ProxyType as HTTP\_PROXY or HTTPS\_PROXY, virtual host will use the
"HOST" HTTP header from the request and perform DNS resolution to select destination endpoint.

For TCP traffic with SNI, (If the ProxyType is TCP\_PROXY\_WITH\_SNI), virtual host will perform DNS
resolution using the SNI.

The DNS resolution is performed in the virtual network specified in outside\_network\_type or
outside\_network

In both modes of operation(either using Host header or SNI), the DNS resolution could return
multiple addresses. First IPv4 address from such returned list is used as endpoint for the request.
The DNS response is cached for 60s by default.

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

## Next pages

- [dynamic_reverse_proxy.resolution_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/dynamic_reverse_proxy/resolution_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
