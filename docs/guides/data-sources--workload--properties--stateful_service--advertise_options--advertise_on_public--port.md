---
page_title: "stateful_service.advertise_options.advertise_on_public.port"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.port for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2424, "body_sha256": "sha256:553f9544e68fec410301174ad9b281866feeb6f1470d540540ddb92b8e4fe169", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:port", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:tcp_loadbalancer"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.port for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.port

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- stateful_service.advertise_options.advertise_on_public.port

<a id="section"></a>

Type: `"single"`. Computed.

Advertise Port. Advertise single port.

Upstream description:

Advertise single port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"http_loadbalancer\",\"tcp_loadbalancer\"]"
}
```

## Direct properties

- [http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md): complete subsection reference.

- [port](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--port.md): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--tcp_loadbalancer.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.port](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--port.md)
- [stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--tcp_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [xcsh_workload](../data-sources/workload.md)
