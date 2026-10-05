---
page_title: "routes.route_destination.web_socket_config"
subcategory: ""
description: "Configuration to allow Websocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'"
xcsh_docs: {"aliases": ["routes route destination web socket config"], "body_bytes": 2313, "body_sha256": "sha256:645113e3537c0377e81f68bdd7d316487e50f2140001556b120ef2d4afad0106", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:web_socket_config", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "documentation/data-sources/route/properties/routes/route_destination/web_socket_config/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2220023021332222-0321110302230020-3013022312112230-1000032111332113-1232203003003021-1212100001333133-2322110231321100-1312103000111302", "registry_path": "docs/guides/data-sources--route--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "web_socket_config"], "schema_version": 1, "sections": [{"aliases": ["routes route destination web socket config use websocket"], "anchor": "schema-routes--route_destination--web_socket_config--use_websocket", "description": "Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket connection.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:web_socket_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "web_socket_config", "use_websocket"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/web_socket_config/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration to allow Websocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.web_socket_config

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- routes.route_destination.web_socket_config

<a id="section"></a>

Type: `"single"`. Computed.

Configuration to allow Websocket Request headers of such upgrade looks like below 'connection',
'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce
following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'.

Upstream description:

Configuration to allow Websocket

Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket'

With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101
Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'

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

<a id="schema-routes--route_destination--web_socket_config--use_websocket"></a>

### use_websocket property

Type: `"bool"`. Computed.

Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket
connection.

Upstream description:

Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket
connection.

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

- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
