---
page_title: "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos"
subcategory: "Load Balancing"
description: "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2141, "body_sha256": "sha256:1e9fd06376dadf1f9e97a6e83aef4a9aa042bcc802de7a71041ead083e983d5c", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations", "selected_repos"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [enable_api_discovery](data-sources--cdn_loadbalancer--properties--enable_api_discovery.md)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan.md)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations.md)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="section"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

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

<a id="schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo"></a>

### api_code_repo property

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

## Next pages

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
