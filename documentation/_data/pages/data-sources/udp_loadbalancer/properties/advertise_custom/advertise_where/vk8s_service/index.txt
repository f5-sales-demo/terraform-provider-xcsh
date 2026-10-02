---
page_title: "advertise_custom.advertise_where.vk8s_service"
subcategory: ""
description: "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network."
xcsh_docs: {"aliases": ["advertise custom advertise where vk8s service"], "body_bytes": 2456, "body_sha256": "sha256:b0770f96792fdd2ac8254e92d7d92e251b4f83fe4cf2db681c62e196d48e9b91", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service:site", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service:virtual_site"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where", "path": "documentation/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1322123110212212-1033222023310221-1021233002213200-2121220132323000-1120320111033002-2012002311123220-0020131200103231-0323132202303320", "registry_path": "docs/guides/data-sources--udp_loadbalancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "vk8s_service"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "vk8s_service", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "vk8s_service", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.vk8s_service

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/)
- advertise_custom.advertise_where.vk8s_service

<a id="section"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/virtual_site/): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.vk8s_service.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/site/)
- [advertise_custom.advertise_where.vk8s_service.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/virtual_site/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/)
- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/)
