---
page_title: "service.advertise_options.advertise_on_public.port"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2510, "body_sha256": "sha256:5dc27d84b82017e5d328c90c52dcb335ae4094e7cb18b0a900a42c27d36e47a6", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:port", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:tcp_loadbalancer"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--port.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- service.advertise_options.advertise_on_public.port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise Port. Advertise single port.

Upstream description:

Advertise single port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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
  "x-ves-oneof-field-advertise_choice": "[\"http_loadbalancer\",\"tcp_loadbalancer\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md): complete subsection reference.

- [port](resources--workload--properties--service--advertise_options--advertise_on_public--port--port.md): complete subsection reference.

- [tcp_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--port--tcp_loadbalancer.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--properties--service--advertise_options--advertise_on_public--port--port.md)
- [service.advertise_options.advertise_on_public.port.tcp_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--port--tcp_loadbalancer.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [xcsh_workload](../resources/workload.md)
