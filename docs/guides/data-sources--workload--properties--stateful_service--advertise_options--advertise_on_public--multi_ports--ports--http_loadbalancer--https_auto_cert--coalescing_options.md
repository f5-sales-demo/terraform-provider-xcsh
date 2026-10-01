---
page_title: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3701, "body_sha256": "sha256:2449508d45b336460faac9822f99761cffdd095ec889827fde6ccbcb1d41dcc5", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https_auto_cert:coalescing_options", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https_auto_cert:coalescing_options:default_coalescing", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https_auto_cert:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https_auto_cert:coalescing_options", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https_auto_cert", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert--coalescing_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https_auto_cert", "coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https_auto_cert/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert.md)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="section"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

## Direct properties

- [default_coalescing](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert--coalescing_options--default_coalescing.md): complete subsection reference.

- [strict_coalescing](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert--coalescing_options--strict_coalescing.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert--coalescing_options--default_coalescing.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert--coalescing_options--strict_coalescing.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert.md)
- [xcsh_workload](../data-sources/workload.md)
