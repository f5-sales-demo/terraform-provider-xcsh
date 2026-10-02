---
page_title: "http_loadbalancer"
subcategory: ""
description: "Set the scope of the API Group to a specific HTTP Loadbalancer."
xcsh_docs: {"aliases": ["http loadbalancer"], "body_bytes": 1296, "body_sha256": "sha256:f98abac7b1393a2bfcfd65e6767da6e00051e95b1c345c4d6ac444003a188f92", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:app_api_group:properties:http_loadbalancer:http_loadbalancer"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_api_group:properties:http_loadbalancer", "parent_id": "xcsh-docs:data-sources:app_api_group:reference", "path": "documentation/data-sources/app_api_group/properties/http_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1012330221310120-1213213021310010-2120320000001322-3233313300332303-2221221001101022-1013130323212102-1323232121331022-2210223203322322", "registry_path": "docs/guides/data-sources--app_api_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_loadbalancer"], "schema_version": 1, "sections": [{"aliases": ["http loadbalancer"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:app_api_group:properties:http_loadbalancer:http_loadbalancer", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_loadbalancer", "http_loadbalancer"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/properties/http_loadbalancer/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Set the scope of the API Group to a specific HTTP Loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [http_loadbalancer.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/http_loadbalancer/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/)
- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
