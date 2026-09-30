---
page_title: "routes.route_destination.web_socket_config"
subcategory: ""
description: "routes.route_destination.web_socket_config for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1908, "body_sha256": "sha256:61ef03000db1df9094d3168c33792f84e2f1793a5ed0c9aeb63dc2984493800c", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:web_socket_config", "child_ids": [], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:web_socket_config", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "docs/guides/data-sources--route--properties--routes--route_destination--web_socket_config.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "web_socket_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/web_socket_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.web_socket_config for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.route_destination.web_socket_config

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
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

- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
- [xcsh_route](../data-sources/route.md)
