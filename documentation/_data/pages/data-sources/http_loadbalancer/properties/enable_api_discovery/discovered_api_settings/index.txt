---
page_title: "enable_api_discovery.discovered_api_settings"
subcategory: "Load Balancing"
description: "Configure Discovered API Settings."
xcsh_docs: {"aliases": ["enable api discovery discovered api settings"], "body_bytes": 1862, "body_sha256": "sha256:f87ccc3d6b29c4dc0e879fd86c32ec16a470dfc7a3efee44480b4533b6b74ebd", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:discovered_api_settings", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery", "path": "documentation/data-sources/http_loadbalancer/properties/enable_api_discovery/discovered_api_settings/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2123111312022000-2333133110203330-0220223302200110-2202032222101322-1202332001000200-3311323320022011-1211213100322320-1233320302231323", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "discovered_api_settings"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery discovered api settings purge duration for inactive discovered apis"], "anchor": "schema-enable_api_discovery--discovered_api_settings--purge_duration_for_inactive_discovered_apis", "description": "Inactive discovered API will be deleted after configured duration.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:discovered_api_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "discovered_api_settings", "purge_duration_for_inactive_discovered_apis"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/enable_api_discovery/discovered_api_settings/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configure Discovered API Settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.discovered_api_settings

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_api_discovery/)
- enable_api_discovery.discovered_api_settings

<a id="section"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

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

<a id="schema-enable_api_discovery--discovered_api_settings--purge_duration_for_inactive_discovered_apis"></a>

### purge_duration_for_inactive_discovered_apis property

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```
