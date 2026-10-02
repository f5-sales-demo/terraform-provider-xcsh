---
page_title: "routes.route_destination.spdy_config"
subcategory: ""
description: "Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1' Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'SPDY/3.1'"
xcsh_docs: {"aliases": ["routes route destination spdy config"], "body_bytes": 2317, "body_sha256": "sha256:fd73bc9c7fb49a2d2232456a84a89e93ec8b93d112b3bf493afee651c2ab36cb", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:spdy_config", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "documentation/data-sources/route/properties/routes/route_destination/spdy_config/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2033321310103033-0133103213233310-2321121102310112-0031320012321002-3003013021300212-0111211013213310-0120322130212131-2220332212000120", "registry_path": "docs/guides/data-sources--route--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "spdy_config"], "schema_version": 1, "sections": [{"aliases": ["use spdy"], "anchor": "schema-routes--route_destination--spdy_config--use_spdy", "description": "Specifies that the HTTP client connection to this route is allowed to upgrade to a SPDY connection.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:spdy_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "spdy_config", "use_spdy"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/spdy_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1' Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'SPDY/3.1'", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.spdy_config

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- routes.route_destination.spdy_config

<a id="section"></a>

Type: `"single"`. Computed.

Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1'
Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to
allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols
'Upgrade'..

Upstream description:

Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1'

Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to
allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade':
'SPDY/3.1' 'Connection': 'Upgrade'

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

<a id="schema-routes--route_destination--spdy_config--use_spdy"></a>

### use_spdy property

Type: `"bool"`. Computed.

Specifies that the HTTP client connection to this route is allowed to upgrade to a SPDY connection.

Upstream description:

Specifies that the HTTP client connection to this route is allowed to upgrade to a SPDY connection.

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
