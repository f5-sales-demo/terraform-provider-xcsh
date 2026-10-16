---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["stateful service advertise options advertise on public port http loadbalancer specific routes routes custom route object caching disable"], "body_bytes": 2953, "body_sha256": "sha256:0ca45ca17dfda023837851ff44c67a1490cc6ff1dab10596bec756c135a85200", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/custom_route_object/caching_disable/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0321331020112113-2221312003120313-3000223222202031-3322213033010313-1221221212232312-1301221220313131-2021320313300001-1303011223112020", "registry_path": "docs/guides/resources--workload--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes", "routes", "custom_route_object", "caching_disable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/custom_route_object/caching_disable/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["workloadCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- [stateful_service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/custom_route_object/)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

Additional upstream details:

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
caching_disable = {}
```

This is an empty object or choice marker. It has no direct properties.
