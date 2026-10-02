---
page_title: "jwt_validation.token_location"
subcategory: "Load Balancing"
description: "Location of JWT in HTTP request."
xcsh_docs: {"aliases": ["jwt validation token location"], "body_bytes": 1712, "body_sha256": "sha256:6e01be40f668720802d929be3fa67036827552741a5a98fb3fcf8b19a5b940e0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:token_location:bearer_token"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:token_location", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "path": "documentation/resources/http_loadbalancer/properties/jwt_validation/token_location/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "token_location"], "schema_version": 1, "sections": [{"aliases": ["bearer token"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:token_location:bearer_token", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "token_location", "bearer_token"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/token_location/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Location of JWT in HTTP request.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.token_location

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/)
- jwt_validation.token_location

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for token location.

Upstream description:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

Terraform syntax:

```terraform
token_location {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/token_location/bearer_token/): complete subsection reference.

## Next pages

- [jwt_validation.token_location.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/token_location/bearer_token/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
