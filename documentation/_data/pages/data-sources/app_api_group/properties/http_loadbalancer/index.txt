---
page_title: "http_loadbalancer"
subcategory: ""
description: "Set the scope of the API Group to a specific HTTP Loadbalancer."
xcsh_docs: {"aliases": ["http loadbalancer"], "body_bytes": 881, "body_sha256": "sha256:539f60cd2974110019d922c76d10c60d9c7489f602408bdefb64af9eafd3ee6d", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:app_api_group:properties:http_loadbalancer:http_loadbalancer"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_api_group:properties:http_loadbalancer", "parent_id": "xcsh-docs:data-sources:app_api_group:reference", "path": "documentation/data-sources/app_api_group/properties/http_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1012330221310120-1213213021310010-2120320000001322-3233313300332303-2221221001101022-1013130323212102-1323232121331022-2210223203322322", "registry_path": "docs/guides/data-sources--app_api_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_loadbalancer"], "schema_version": 1, "sections": [{"aliases": ["http loadbalancer http loadbalancer"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:app_api_group:properties:http_loadbalancer:http_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_loadbalancer", "http_loadbalancer"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/properties/http_loadbalancer/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Set the scope of the API Group to a specific HTTP Loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_loadbalancer

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/)
- http_loadbalancer

<a id="section"></a>

Type: `"single"`. Computed.

Set the scope of the API Group to a specific HTTP Loadbalancer.

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

- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/http_loadbalancer/http_loadbalancer/): complete subsection reference.
