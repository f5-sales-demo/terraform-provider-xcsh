---
page_title: "stateful_service.advertise_options.advertise_on_public"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1749, "body_sha256": "sha256:2e9bdcd5bca098f2cbdec4b71799fea891e19b62775017ece05dab46a8b7fea6", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- stateful_service.advertise_options.advertise_on_public

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

- [multi_ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports.md): complete subsection reference.

- [port](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports.md)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [xcsh_workload](../data-sources/workload.md)
