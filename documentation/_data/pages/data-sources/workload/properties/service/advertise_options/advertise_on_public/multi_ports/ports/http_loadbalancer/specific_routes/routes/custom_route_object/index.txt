---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object"
subcategory: "Container"
description: "A custom route uses a route object created outside of this view."
xcsh_docs: {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer specific routes routes custom route object"], "body_bytes": 5172, "body_sha256": "sha256:35f03b6fa1412d3c69ac28dce87094480e75a6b8db64094454f3c9622ad5b47c", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_inherit", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:custom_route_object:route_ref"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/custom_route_object/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0023023121203023-2203013300322033-1103112202110303-2012230110223202-3101000000330131-2221232020131133-2231111102330200-0033313100132023", "registry_path": "docs/guides/data-sources--workload--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "custom_route_object"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer specific routes routes custom route object caching disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "custom_route_object", "caching_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer specific routes routes custom route object caching inherit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_inherit", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "custom_route_object", "caching_inherit"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer specific routes routes custom route object route ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:custom_route_object:route_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "custom_route_object", "route_ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/custom_route_object/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A custom route uses a route object created outside of this view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="section"></a>

Type: `"single"`. Computed.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

## Direct properties

- [caching_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/custom_route_object/caching_disable/): complete subsection reference.

- [caching_inherit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/custom_route_object/caching_inherit/): complete subsection reference.

- [route_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/custom_route_object/route_ref/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/custom_route_object/caching_disable/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/custom_route_object/caching_inherit/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/custom_route_object/route_ref/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
