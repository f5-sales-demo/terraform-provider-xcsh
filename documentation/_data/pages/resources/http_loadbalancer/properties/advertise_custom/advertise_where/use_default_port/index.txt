---
page_title: "advertise_custom.advertise_where.use_default_port"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["advertise custom advertise where use default port"], "body_bytes": 1536, "body_sha256": "sha256:a51ca2c3c93872627e40a49da23297a9c2ccba4e9fbc6fef82ca8de4b7cfd1c9", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where", "path": "documentation/resources/http_loadbalancer/properties/advertise_custom/advertise_where/use_default_port/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2120112203220322-0333003120010022-3103323201200001-0120131312031010-0012133133121033-2222303231110323-0303321132023013-0320122113323221", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "use_default_port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/advertise_custom/advertise_where/use_default_port/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.use_default_port

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_custom/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_custom/advertise_where/)
- advertise_custom.advertise_where.use_default_port

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
use_default_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_custom/advertise_where/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
