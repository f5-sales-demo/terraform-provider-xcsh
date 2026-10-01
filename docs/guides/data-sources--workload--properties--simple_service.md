---
page_title: "simple_service"
subcategory: "Container"
description: "simple_service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2782, "body_sha256": "sha256:33c862196d917c4a691fbd074c6ee36f636f350cfc17fa2fcdd7292ba96dcc14", "canonical_id": "xcsh-docs:data-sources:workload:properties:simple_service", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:configuration", "xcsh-docs:data-sources:workload:properties:simple_service:container", "xcsh-docs:data-sources:workload:properties:simple_service:disabled", "xcsh-docs:data-sources:workload:properties:simple_service:do_not_advertise", "xcsh-docs:data-sources:workload:properties:simple_service:enabled", "xcsh-docs:data-sources:workload:properties:simple_service:simple_advertise"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service", "parent_id": "xcsh-docs:data-sources:workload:reference", "path": "docs/guides/data-sources--workload--properties--simple_service.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["simple_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- simple_service

<a id="section"></a>

Type: `"single"`. Computed.

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Upstream description:

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

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

## Direct properties

- [configuration](data-sources--workload--properties--simple_service--configuration.md): complete subsection reference.

- [container](data-sources--workload--properties--simple_service--container.md): complete subsection reference.

- [disabled](data-sources--workload--properties--simple_service--disabled.md): complete subsection reference.

- [do_not_advertise](data-sources--workload--properties--simple_service--do_not_advertise.md): complete subsection reference.

- [enabled](data-sources--workload--properties--simple_service--enabled.md): complete subsection reference.

<a id="schema-simple_service--scale_to_zero"></a>

### scale_to_zero property

Type: `"bool"`. Computed.

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

- [simple_advertise](data-sources--workload--properties--simple_service--simple_advertise.md): complete subsection reference.

## Next pages

- [simple_service.configuration](data-sources--workload--properties--simple_service--configuration.md)
- [simple_service.container](data-sources--workload--properties--simple_service--container.md)
- [simple_service.disabled](data-sources--workload--properties--simple_service--disabled.md)
- [simple_service.do_not_advertise](data-sources--workload--properties--simple_service--do_not_advertise.md)
- [simple_service.enabled](data-sources--workload--properties--simple_service--enabled.md)
- [simple_service.simple_advertise](data-sources--workload--properties--simple_service--simple_advertise.md)
- [Property reference](data-sources--workload--reference.md)
- [xcsh_workload](../data-sources/workload.md)
