---
page_title: "service"
subcategory: "Container"
description: "service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3425, "body_sha256": "sha256:ae7aa0e305bf809ab7bcbe588e659265355970dec3da4af5a9327088533de183", "canonical_id": "xcsh-docs:data-sources:workload:properties:service", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options", "xcsh-docs:data-sources:workload:properties:service:configuration", "xcsh-docs:data-sources:workload:properties:service:containers", "xcsh-docs:data-sources:workload:properties:service:deploy_options", "xcsh-docs:data-sources:workload:properties:service:scale_to_zero", "xcsh-docs:data-sources:workload:properties:service:volumes"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service", "parent_id": "xcsh-docs:data-sources:workload:reference", "path": "docs/guides/data-sources--workload--properties--service.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- service

<a id="section"></a>

Type: `"single"`. Computed.

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers..

Upstream description:

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers,
traditional SQL databases, etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

## Direct properties

- [advertise_options](data-sources--workload--properties--service--advertise_options.md): complete subsection reference.

- [configuration](data-sources--workload--properties--service--configuration.md): complete subsection reference.

- [containers](data-sources--workload--properties--service--containers.md): complete subsection reference.

- [deploy_options](data-sources--workload--properties--service--deploy_options.md): complete subsection reference.

<a id="schema-service--num_replicas"></a>

### num_replicas property

Type: `"number"`. Computed.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [scale_to_zero](data-sources--workload--properties--service--scale_to_zero.md): complete subsection reference.

- [volumes](data-sources--workload--properties--service--volumes.md): complete subsection reference.

## Next pages

- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.configuration](data-sources--workload--properties--service--configuration.md)
- [service.containers](data-sources--workload--properties--service--containers.md)
- [service.deploy_options](data-sources--workload--properties--service--deploy_options.md)
- [service.scale_to_zero](data-sources--workload--properties--service--scale_to_zero.md)
- [service.volumes](data-sources--workload--properties--service--volumes.md)
- [Property reference](data-sources--workload--reference.md)
- [xcsh_workload](../data-sources/workload.md)
