---
page_title: "routes.route_destination.spdy_config"
subcategory: ""
description: "Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1' Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'SPDY/3.1'"
xcsh_docs: {"aliases": ["routes route destination spdy config"], "body_bytes": 1729, "body_sha256": "sha256:7d39cf54a50580eb3403184635e45fa8d2861cfb76d783dd7c0cab74e08c36fc", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:spdy_config", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/spdy_config/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1211112100111122-3213121213133231-1231320323033110-2010301220033103-1021202000330100-0023023012000230-0022012210003120-3103323301010233", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "spdy_config"], "schema_version": 1, "sections": [{"aliases": ["routes route destination spdy config use spdy"], "anchor": "schema-routes--route_destination--spdy_config--use_spdy", "description": "Specifies that the HTTP client connection to this route is allowed to upgrade to a SPDY connection.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:spdy_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "spdy_config", "use_spdy"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/spdy_config/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1' Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'SPDY/3.1'", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["routeCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.spdy_config

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- routes.route_destination.spdy_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
spdy_config {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--route_destination--spdy_config--use_spdy"></a>

### use_spdy property

Type: `"bool"`. Optional.

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
