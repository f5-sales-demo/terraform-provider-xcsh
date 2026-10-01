---
page_title: "simple_service"
subcategory: "Container"
description: "simple_service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3916, "body_sha256": "sha256:b7b8679bbb1f0a548865ec1614427585ca13d2c0fdfe6ba8b65076a05bdba6b1", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:configuration", "xcsh-docs:resources:workload:properties:simple_service:container", "xcsh-docs:resources:workload:properties:simple_service:disabled", "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "xcsh-docs:resources:workload:properties:simple_service:enabled", "xcsh-docs:resources:workload:properties:simple_service:simple_advertise"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service", "parent_id": "xcsh-docs:resources:workload:reference", "path": "documentation/resources/workload/properties/simple_service/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["simple_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
