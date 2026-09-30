---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3533, "body_sha256": "sha256:08390c43e3f0bd4c142630d3796d56070f8d72d34b134e8a26e5fca96ffd20ba", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:coalescing_options", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:coalescing_options:default_coalescing", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:coalescing_options", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--coalescing_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https.md)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_coalescing](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--coalescing_options--default_coalescing.md): complete subsection reference.

- [strict_coalescing](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--coalescing_options--strict_coalescing.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--coalescing_options--default_coalescing.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--coalescing_options--strict_coalescing.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https.md)
- [xcsh_workload](../resources/workload.md)
