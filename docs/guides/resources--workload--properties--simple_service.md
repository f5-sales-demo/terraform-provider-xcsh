---
page_title: "simple_service"
subcategory: "Container"
description: "simple_service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3108, "body_sha256": "sha256:a3b510b8aae4c81d0443ab7afbddd9f672ceb284d43b9f144dca4861187a869f", "canonical_id": "xcsh-docs:resources:workload:properties:simple_service", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:configuration", "xcsh-docs:resources:workload:properties:simple_service:container", "xcsh-docs:resources:workload:properties:simple_service:disabled", "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "xcsh-docs:resources:workload:properties:simple_service:enabled", "xcsh-docs:resources:workload:properties:simple_service:simple_advertise"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service", "parent_id": "xcsh-docs:resources:workload:reference", "path": "docs/guides/resources--workload--properties--simple_service.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["simple_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- simple_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Upstream description:

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disabled",
    "enabled"),
  validators.ConflictingObjectAttributes("do_not_advertise",
    "simple_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"do_not_advertise\",\"simple_advertise\"]",
  "x-ves-oneof-field-persistence_choice": "[\"disabled\",\"enabled\"]"
}
```

Terraform syntax:

```terraform
simple_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [configuration](resources--workload--properties--simple_service--configuration.md): complete subsection reference.

- [container](resources--workload--properties--simple_service--container.md): complete subsection reference.

- [disabled](resources--workload--properties--simple_service--disabled.md): complete subsection reference.

- [do_not_advertise](resources--workload--properties--simple_service--do_not_advertise.md): complete subsection reference.

- [enabled](resources--workload--properties--simple_service--enabled.md): complete subsection reference.

<a id="schema-simple_service--scale_to_zero"></a>

### scale_to_zero property

Type: `"bool"`. Optional.

Scale down replicas of the service to zero.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [simple_advertise](resources--workload--properties--simple_service--simple_advertise.md): complete subsection reference.

## Next pages

- [simple_service.configuration](resources--workload--properties--simple_service--configuration.md)
- [simple_service.container](resources--workload--properties--simple_service--container.md)
- [simple_service.disabled](resources--workload--properties--simple_service--disabled.md)
- [simple_service.do_not_advertise](resources--workload--properties--simple_service--do_not_advertise.md)
- [simple_service.enabled](resources--workload--properties--simple_service--enabled.md)
- [simple_service.simple_advertise](resources--workload--properties--simple_service--simple_advertise.md)
- [Property reference](resources--workload--reference.md)
- [xcsh_workload](../resources/workload.md)
