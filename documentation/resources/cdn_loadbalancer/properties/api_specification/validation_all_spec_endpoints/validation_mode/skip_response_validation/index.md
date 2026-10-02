---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints validation mode skip response validation"], "body_bytes": 1936, "body_sha256": "sha256:065919f8bd5fd4ebac77aa842a15ad2b9efd76ffa5fef8db0f303855675d89e6", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_response_validation/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1302321003220022-2132011332213003-1322331023203221-0200233011032321-3231101312030110-2301032211311103-3333123113230022-1231203132302101", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "skip_response_validation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_response_validation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

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
skip_response_validation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
