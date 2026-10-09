---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["stateful service advertise options advertise custom ports http loadbalancer specific routes routes simple route disable host rewrite"], "body_bytes": 2884, "body_sha256": "sha256:54175800e68e53eb4e7e517eea91e109180ee702e3e5ea00d09afb1c117c1fa0", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route:disable_host_rewrite", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/simple_route/disable_host_rewrite/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2303303222023110-0003330013102013-2201332102011130-0031130223212130-0233113220330332-3111023220213132-3133211023321301-2033110210110120", "registry_path": "docs/guides/resources--workload--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "simple_route", "disable_host_rewrite"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/simple_route/disable_host_rewrite/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- [stateful_service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/simple_route/)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disable_host_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.
