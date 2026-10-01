---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2147, "body_sha256": "sha256:420b7aa8bf1f1948aff3183875da8eba15ac30828c200e0708c13da826e299f5", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:pass_through", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:pass_through", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--pass_through.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "pass_through"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/pass_through/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](resources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
pass_through = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- [xcsh_workload](../resources/workload.md)
