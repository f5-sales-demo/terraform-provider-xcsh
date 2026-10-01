---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.multi_ports.ports for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3296, "body_sha256": "sha256:21cd8658667dd88a36dd20409ddb7e412b66178116535a9b8c4a9571a093ea94", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:port", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:tcp_loadbalancer"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.multi_ports.ports for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.multi_ports.ports

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports.md)
- service.advertise_options.advertise_on_public.multi_ports.ports

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer.md): complete subsection reference.

- [port](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--port.md): complete subsection reference.

- [tcp_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--tcp_loadbalancer.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--port.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--tcp_loadbalancer.md)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports.md)
- [xcsh_workload](../resources/workload.md)
