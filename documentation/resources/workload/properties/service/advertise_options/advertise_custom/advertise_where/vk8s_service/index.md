---
page_title: "service.advertise_options.advertise_custom.advertise_where.vk8s_service"
subcategory: "Container"
description: "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network."
xcsh_docs: {"aliases": ["service advertise options advertise custom advertise where vk8s service"], "body_bytes": 3233, "body_sha256": "sha256:325edb093ada671b5ec5624fe5f6f1178973163c58a821edf6d84f1712750c79", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:site", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:virtual_site"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/vk8s_service/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:virtual_site", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where", "vk8s_service"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise custom advertise where vk8s service site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--advertise_options--advertise_custom--advertise_where--vk8s_service--site--name", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where.vk8s_service.site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:site", "type": "requires"}], "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where", "vk8s_service", "site"], "syntax": "block", "type": "object"}, {"aliases": ["service advertise options advertise custom advertise where vk8s service virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--advertise_options--advertise_custom--advertise_where--vk8s_service--virtual_site--name", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:virtual_site", "type": "requires"}], "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where", "vk8s_service", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/vk8s_service/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.advertise_where.vk8s_service

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
```

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

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/vk8s_service/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/vk8s_service/virtual_site/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/vk8s_service/site/)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/vk8s_service/virtual_site/)
- [service.advertise_options.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
