---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["service advertise options advertise on public port http loadbalancer https auto cert coalescing options default coalescing"], "body_bytes": 2495, "body_sha256": "sha256:ddc6a9dde3277646690c97ef65813d84ed3c393b202407a53216cb87034f5b6b", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:coalescing_options:default_coalescing", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:coalescing_options", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/coalescing_options/default_coalescing/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1002111202130012-2213212112200030-2231331233321001-3123133202012110-2111301012221231-0320131333132230-0113310203203222-2311213030231001", "registry_path": "docs/guides/resources--workload--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "coalescing_options", "default_coalescing"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/coalescing_options/default_coalescing/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/coalescing_options/)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.
