---
page_title: "routes.route_destination.spdy_config"
subcategory: ""
description: "routes.route_destination.spdy_config for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2011, "body_sha256": "sha256:116569481873d292e8e53fec3c0853b438a2e95aeae3768c0691f69948436ba5", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:spdy_config", "child_ids": [], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:spdy_config", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "docs/guides/data-sources--route--properties--routes--route_destination--spdy_config.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "spdy_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/spdy_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.spdy_config for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.spdy_config

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
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

- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
- [xcsh_route](../data-sources/route.md)
