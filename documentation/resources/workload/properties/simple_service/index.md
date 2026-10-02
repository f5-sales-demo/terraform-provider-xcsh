---
page_title: "simple_service"
subcategory: "Container"
description: "SimpleService is a service having one container and one replica that is deployed on all Regional Edges and advertised on Internet via HTTP loadbalancer on default VIP."
xcsh_docs: {"aliases": ["simple service"], "body_bytes": 3916, "body_sha256": "sha256:b7b8679bbb1f0a548865ec1614427585ca13d2c0fdfe6ba8b65076a05bdba6b1", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:configuration", "xcsh-docs:resources:workload:properties:simple_service:container", "xcsh-docs:resources:workload:properties:simple_service:disabled", "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "xcsh-docs:resources:workload:properties:simple_service:enabled", "xcsh-docs:resources:workload:properties:simple_service:simple_advertise"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service", "parent_id": "xcsh-docs:resources:workload:reference", "path": "documentation/resources/workload/properties/simple_service/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "simple_service:ConflictingObjectAttributes:disabled,enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service:ConflictingObjectAttributes:do_not_advertise,simple_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service:ConflictingObjectAttributes:disabled,enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service:ConflictingObjectAttributes:do_not_advertise,simple_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service"], "schema_version": 1, "sections": [{"aliases": ["configuration"], "anchor": "section", "description": "Configuration parameters of the workload.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "configuration"], "syntax": "block", "type": "object"}, {"aliases": ["container"], "anchor": "section", "description": "ContainerType configures the container information.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--container--flavor", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container", "type": "conflicts"}, {"anchor": "schema-simple_service--container--flavor", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:default_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container:ConflictingObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:default_flavor", "type": "conflicts"}, {"anchor": "schema-simple_service--container--name", "enforcement": "provider-schema", "group": "simple_service.container:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container", "type": "requires"}], "schema_path": ["simple_service", "container"], "syntax": "block", "type": "object"}, {"aliases": ["disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["do not advertise"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "do_not_advertise"], "syntax": "attribute", "type": "object"}, {"aliases": ["enabled"], "anchor": "section", "description": "Persistent storage volume configuration for the workload.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:enabled", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--enabled--name", "enforcement": "provider-schema", "group": "simple_service.enabled:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:enabled", "type": "requires"}], "schema_path": ["simple_service", "enabled"], "syntax": "block", "type": "object"}, {"aliases": ["scale to zero"], "anchor": "schema-simple_service--scale_to_zero", "description": "Scale down replicas of the service to zero.", "document_id": "xcsh-docs:resources:workload:properties:simple_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "scale_to_zero"], "syntax": "attribute", "type": "bool"}, {"aliases": ["simple advertise"], "anchor": "section", "description": "Advertise OPTIONS for Simple Service.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--simple_advertise--domains", "enforcement": "provider-schema", "group": "simple_service.simple_advertise:RequiredObjectAttributes:domains,service_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "type": "requires"}, {"anchor": "schema-simple_service--simple_advertise--service_port", "enforcement": "provider-schema", "group": "simple_service.simple_advertise:RequiredObjectAttributes:domains,service_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "type": "requires"}], "schema_path": ["simple_service", "simple_advertise"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SimpleService is a service having one container and one replica that is deployed on all Regional Edges and advertised on Internet via HTTP loadbalancer on default VIP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [simple_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/)
- [simple_service.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/disabled/)
- [simple_service.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/do_not_advertise/)
- [simple_service.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/enabled/)
- [simple_service.simple_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/simple_advertise/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
