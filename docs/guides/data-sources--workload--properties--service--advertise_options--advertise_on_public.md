---
page_title: "service.advertise_options.advertise_on_public"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1623, "body_sha256": "sha256:3c0adc740cc4b06a3a11fbd767e22339814c39ef07972fda20372fd18c05afd3", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_on_public.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- service.advertise_options.advertise_on_public

<a id="section"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on Internet with default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

## Direct properties

- [multi_ports](data-sources--workload--properties--service--advertise_options--advertise_on_public--multi_ports.md): complete subsection reference.

- [port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--properties--service--advertise_options--advertise_on_public--multi_ports.md)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [xcsh_workload](../data-sources/workload.md)
