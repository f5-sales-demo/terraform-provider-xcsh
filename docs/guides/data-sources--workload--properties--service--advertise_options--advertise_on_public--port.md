---
page_title: "service.advertise_options.advertise_on_public.port"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2154, "body_sha256": "sha256:dbe16193027460f3f08a62f1e9ddb9444bec83e4e51d35244b6f88b5bd6c6813", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:port", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:tcp_loadbalancer"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_on_public.port

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](data-sources--workload--properties--service--advertise_options--advertise_on_public.md)
- service.advertise_options.advertise_on_public.port

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

- [http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md): complete subsection reference.

- [port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--port.md): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--tcp_loadbalancer.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--port.md)
- [service.advertise_options.advertise_on_public.port.tcp_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--tcp_loadbalancer.md)
- [service.advertise_options.advertise_on_public](data-sources--workload--properties--service--advertise_options--advertise_on_public.md)
- [xcsh_workload](../data-sources/workload.md)
