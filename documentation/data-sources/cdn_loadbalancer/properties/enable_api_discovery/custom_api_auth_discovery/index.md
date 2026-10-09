---
page_title: "enable_api_discovery.custom_api_auth_discovery"
subcategory: "Load Balancing"
description: "API Discovery Advanced settings."
xcsh_docs: {"aliases": ["enable api discovery custom api auth discovery"], "body_bytes": 1127, "body_sha256": "sha256:2ea2b571e30c7b94f6bfec110a0da1cce93a544a5a2befbfb3e35d3219897609", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery", "path": "documentation/data-sources/cdn_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3120303222023021-2012311311330033-1012013330102033-1102212133223323-3013330220323313-2130103020333101-0223132211201203-3031221211130312", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "custom_api_auth_discovery"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery custom api auth discovery api discovery ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "API Discovery Advanced settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.custom_api_auth_discovery

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/enable_api_discovery/)
- enable_api_discovery.custom_api_auth_discovery

<a id="section"></a>

Type: `"single"`. Computed.

API Discovery Advanced Settings. API Discovery Advanced settings.

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

- [api_discovery_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/api_discovery_ref/): complete subsection reference.
