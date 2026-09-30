---
page_title: "service.advertise_options.advertise_on_public"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1787, "body_sha256": "sha256:3da1f59ec7b599f018110f201e3c349422a048354be4f45b1e279dbc6057a32c", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_on_public

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- service.advertise_options.advertise_on_public

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on Internet with default VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("multi_ports",
    "port")}
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
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

Terraform syntax:

```terraform
advertise_on_public {
  # Configure direct properties listed below.
}
```

## Direct properties

- [multi_ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports.md): complete subsection reference.

- [port](resources--workload--properties--service--advertise_options--advertise_on_public--port.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports.md)
- [service.advertise_options.advertise_on_public.port](resources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [xcsh_workload](../resources/workload.md)
