---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3184, "body_sha256": "sha256:b2159cee57e46d12035e52bb715aefa3cc27c7dc2ee7eb9503615ab57ff09647", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:coalescing_options", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:coalescing_options:default_coalescing", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:coalescing_options", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--coalescing_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

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

- [default_coalescing](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--coalescing_options--default_coalescing.md): complete subsection reference.

- [strict_coalescing](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--coalescing_options--strict_coalescing.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--coalescing_options--default_coalescing.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--coalescing_options--strict_coalescing.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [xcsh_workload](../data-sources/workload.md)
