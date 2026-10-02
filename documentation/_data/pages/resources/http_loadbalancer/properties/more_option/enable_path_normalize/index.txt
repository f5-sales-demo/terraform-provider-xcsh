---
page_title: "more_option.enable_path_normalize"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["more option enable path normalize"], "body_bytes": 1290, "body_sha256": "sha256:0e607e3adb15c2a01fb76b20d432204d0ea79f9ad2590f8307ad19d7006fc007", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:enable_path_normalize", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option", "path": "documentation/resources/http_loadbalancer/properties/more_option/enable_path_normalize/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0200102323330033-2123023013233122-3220321300123212-3020110001131003-2313312012331231-1030020232012220-3302210111313313-3333332202102203", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-021.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["more_option", "enable_path_normalize"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/more_option/enable_path_normalize/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.enable_path_normalize

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/)
- more_option.enable_path_normalize

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
enable_path_normalize = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
