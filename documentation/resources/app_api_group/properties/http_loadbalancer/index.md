---
page_title: "http_loadbalancer"
subcategory: ""
description: "Set the scope of the API Group to a specific HTTP Loadbalancer."
xcsh_docs: {"aliases": ["http loadbalancer"], "body_bytes": 1404, "body_sha256": "sha256:7aef2996cdac8fbbafd3306b584a24a2af420bacec5e43d7331950e3277aedec", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:app_api_group:properties:http_loadbalancer:http_loadbalancer"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:properties:http_loadbalancer", "parent_id": "xcsh-docs:resources:app_api_group:reference", "path": "documentation/resources/app_api_group/properties/http_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0111131201032300-2301200102322231-2002333313023333-3002103130232211-0212123110011020-1030313101300031-3221101213030123-0032303202233032", "registry_path": "docs/guides/resources--app_api_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_loadbalancer"], "schema_version": 1, "sections": [{"aliases": ["http loadbalancer"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:app_api_group:properties:http_loadbalancer:http_loadbalancer", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_loadbalancer--http_loadbalancer--name", "enforcement": "provider-schema", "group": "http_loadbalancer.http_loadbalancer:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_api_group:properties:http_loadbalancer:http_loadbalancer", "type": "requires"}], "schema_path": ["http_loadbalancer", "http_loadbalancer"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/properties/http_loadbalancer/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Set the scope of the API Group to a specific HTTP Loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_loadbalancer

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- http_loadbalancer

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/http_loadbalancer/): complete subsection reference.

## Next pages

- [http_loadbalancer.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
