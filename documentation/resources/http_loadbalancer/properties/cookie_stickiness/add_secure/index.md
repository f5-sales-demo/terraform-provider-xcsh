---
page_title: "cookie_stickiness.add_secure"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cookie stickiness add secure"], "body_bytes": 1293, "body_sha256": "sha256:f9e5367bffefdea16fbff7f61e56c6be500a73cb16085f712361039eec601993", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:add_secure", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness", "path": "documentation/resources/http_loadbalancer/properties/cookie_stickiness/add_secure/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2010011221312311-2201313103330320-0310313033132213-3230012221222320-0122233030233133-2213133110022103-2201003011023032-1310213320202133", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cookie_stickiness", "add_secure"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/cookie_stickiness/add_secure/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_stickiness.add_secure

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [cookie_stickiness](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/)
- cookie_stickiness.add_secure

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
add_secure = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cookie_stickiness](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
