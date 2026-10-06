---
page_title: "simple_service"
subcategory: "Container"
description: "SimpleService is a service having one container and one replica that is deployed on all Regional Edges and advertised on Internet via HTTP loadbalancer on default VIP."
xcsh_docs: {"aliases": ["simple service"], "body_bytes": 2620, "body_sha256": "sha256:0bb769fb8d84adb3ac2b7f7fafdc1ce1ebfc8d4d6e276b16f4fb6b1c480d96ba", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:configuration", "xcsh-docs:resources:workload:properties:simple_service:container", "xcsh-docs:resources:workload:properties:simple_service:disabled", "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "xcsh-docs:resources:workload:properties:simple_service:enabled", "xcsh-docs:resources:workload:properties:simple_service:simple_advertise"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service", "parent_id": "xcsh-docs:resources:workload:reference", "path": "documentation/resources/workload/properties/simple_service/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "simple_service:ConflictingObjectAttributes:disabled,enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service:ConflictingObjectAttributes:do_not_advertise,simple_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service:ConflictingObjectAttributes:disabled,enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service:ConflictingObjectAttributes:do_not_advertise,simple_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service"], "schema_version": 1, "sections": [{"aliases": ["simple service configuration"], "anchor": "section", "description": "Configuration parameters of the workload.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "configuration"], "syntax": "block", "type": "object"}, {"aliases": ["simple service container"], "anchor": "section", "description": "ContainerType configures the container information.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--container--flavor", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container", "type": "conflicts"}, {"anchor": "schema-simple_service--container--flavor", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:default_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:default_flavor", "type": "conflicts"}, {"anchor": "schema-simple_service--container--name", "enforcement": "provider-schema", "group": "simple_service.container:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container", "type": "requires"}], "schema_path": ["simple_service", "container"], "syntax": "block", "type": "object"}, {"aliases": ["simple service disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["simple service do not advertise"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "do_not_advertise"], "syntax": "attribute", "type": "object"}, {"aliases": ["simple service enabled"], "anchor": "section", "description": "Persistent storage volume configuration for the workload.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:enabled", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--enabled--name", "enforcement": "provider-schema", "group": "simple_service.enabled:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:enabled", "type": "requires"}], "schema_path": ["simple_service", "enabled"], "syntax": "block", "type": "object"}, {"aliases": ["simple service scale to zero"], "anchor": "schema-simple_service--scale_to_zero", "description": "Scale down replicas of the service to zero.", "document_id": "xcsh-docs:resources:workload:properties:simple_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "scale_to_zero"], "syntax": "attribute", "type": "bool"}, {"aliases": ["simple service simple advertise"], "anchor": "section", "description": "Advertise OPTIONS for Simple Service.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--simple_advertise--domains", "enforcement": "provider-schema", "group": "simple_service.simple_advertise:RequiredObjectAttributes:domains,service_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "type": "requires"}, {"anchor": "schema-simple_service--simple_advertise--service_port", "enforcement": "provider-schema", "group": "simple_service.simple_advertise:RequiredObjectAttributes:domains,service_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "type": "requires"}], "schema_path": ["simple_service", "simple_advertise"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SimpleService is a service having one container and one replica that is deployed on all Regional Edges and advertised on Internet via HTTP loadbalancer on default VIP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- simple_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/): complete subsection reference.

- [container](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/): complete subsection reference.

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/disabled/): complete subsection reference.

- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/do_not_advertise/): complete subsection reference.

- [enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/enabled/): complete subsection reference.

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

- [simple_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/simple_advertise/): complete subsection reference.
