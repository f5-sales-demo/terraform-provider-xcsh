---
page_title: "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations"
subcategory: "Load Balancing"
description: "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3214, "body_sha256": "sha256:455dd80a3646b80c55810f6b75f9dbae7c51404c2f470274721821679650479a", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:code_base_integration", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "path": "docs/guides/data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [enable_api_discovery](data-sources--http_loadbalancer--properties--enable_api_discovery.md)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan.md)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="section"></a>

Type: `"list"`. Computed.

Configuration parameter for code base integrations.

Upstream description:

Configuration parameter for code base integrations

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [all_repos](data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--all_repos.md): complete subsection reference.

- [code_base_integration](data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration.md): complete subsection reference.

- [selected_repos](data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos.md): complete subsection reference.

## Next pages

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--all_repos.md)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration.md)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos.md)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
