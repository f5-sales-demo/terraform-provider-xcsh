---
page_title: "cdn_loadbalancer"
subcategory: ""
description: "Set the scope of the API Group to a specific CDN Loadbalancer."
xcsh_docs: {"aliases": ["cdn loadbalancer"], "body_bytes": 991, "body_sha256": "sha256:2ed1a86331726757d90b7c28e9ed772d27f46e0b6a19976cd1b738ced0b5b2f4", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer", "parent_id": "xcsh-docs:resources:app_api_group:reference", "path": "documentation/resources/app_api_group/properties/cdn_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2122310110220121-1311210010012320-2130100120201203-2012322123033101-1200100321200231-2131013122300131-2032300202032303-2103323121312333", "registry_path": "docs/guides/resources--app_api_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cdn_loadbalancer"], "schema_version": 1, "sections": [{"aliases": ["cdn loadbalancer cdn loadbalancer"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cdn_loadbalancer--cdn_loadbalancer--name", "enforcement": "provider-schema", "group": "cdn_loadbalancer.cdn_loadbalancer:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer", "type": "requires"}], "schema_path": ["cdn_loadbalancer", "cdn_loadbalancer"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/properties/cdn_loadbalancer/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Set the scope of the API Group to a specific CDN Loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cdn_loadbalancer

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- cdn_loadbalancer

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Set the scope of the API Group to a specific CDN Loadbalancer.

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
cdn_loadbalancer {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/cdn_loadbalancer/): complete subsection reference.
