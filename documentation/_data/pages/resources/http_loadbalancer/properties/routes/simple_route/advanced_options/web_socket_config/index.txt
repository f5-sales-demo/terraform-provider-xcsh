---
page_title: "routes.simple_route.advanced_options.web_socket_config"
subcategory: "Load Balancing"
description: "Configuration to allow Websocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'"
xcsh_docs: {"aliases": ["routes simple route advanced options web socket config"], "body_bytes": 2743, "body_sha256": "sha256:e3e7f79dd71a27144cad2bee98959bf834a412afa0d51f8e59e1943a21c0bdfa", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:web_socket_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/web_socket_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2221130310300232-3333311201210030-2220032332203101-0111312233003310-0330020033223201-2312212013003202-2330320321012310-3132220122331020", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "web_socket_config"], "schema_version": 1, "sections": [{"aliases": ["use websocket"], "anchor": "schema-routes--simple_route--advanced_options--web_socket_config--use_websocket", "description": "Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket connection.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:web_socket_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "web_socket_config", "use_websocket"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/web_socket_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration to allow Websocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.web_socket_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- routes.simple_route.advanced_options.web_socket_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
web_socket_config {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--simple_route--advanced_options--web_socket_config--use_websocket"></a>

### use_websocket property

Type: `"bool"`. Optional.

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

- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
