---
page_title: "single_lb_app.enable_discovery.discovered_api_settings"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.discovered_api_settings for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2096, "body_sha256": "sha256:c6e1c3022bcbd72393845eeb68f031a25943ed9beb960af8cc98ffa742df8a15", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:discovered_api_settings", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:discovered_api_settings", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery", "path": "docs/guides/data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--discovered_api_settings.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "discovered_api_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/discovered_api_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.discovered_api_settings for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.discovered_api_settings

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [single_lb_app](data-sources--http_loadbalancer--properties--single_lb_app.md)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- single_lb_app.enable_discovery.discovered_api_settings

<a id="section"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

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

<a id="schema-single_lb_app--enable_discovery--discovered_api_settings--purge_duration_for_inactive_discovered_apis"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
