---
page_title: "jwt_validation.target.all_endpoint"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["jwt validation target all endpoint"], "body_bytes": 1454, "body_sha256": "sha256:eb2cf21eb81f5adbc6bac4928f8f0b344824af3e0472e4416fee8af52ce9ab44", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target:all_endpoint", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target", "path": "documentation/resources/http_loadbalancer/properties/jwt_validation/target/all_endpoint/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1020320031023321-0320202313120203-1020333130031033-0231120033031022-3001322011013132-1031131331321201-2333121233322230-3033322013201211", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "target", "all_endpoint"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/target/all_endpoint/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.target.all_endpoint

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/)
- [jwt_validation.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/target/)
- jwt_validation.target.all_endpoint

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
all_endpoint = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [jwt_validation.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/target/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
