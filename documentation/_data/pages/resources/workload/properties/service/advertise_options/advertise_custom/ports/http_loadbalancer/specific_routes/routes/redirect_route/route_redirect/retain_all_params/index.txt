---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["service advertise options advertise custom ports http loadbalancer specific routes routes redirect route route redirect retain all params"], "body_bytes": 3111, "body_sha256": "sha256:a85e48f765ef910b20295ef72486437e491f2bf28621383ab52f12bce149894a", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect:retain_all_params", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/route_redirect/retain_all_params/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3322210232313033-0233122011002301-0322211331113203-1122233332212220-2200101030301312-1000312321303020-1331011230331221-0133023110033110", "registry_path": "docs/guides/resources--workload--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "redirect_route", "route_redirect", "retain_all_params"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/route_redirect/retain_all_params/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/route_redirect/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.
