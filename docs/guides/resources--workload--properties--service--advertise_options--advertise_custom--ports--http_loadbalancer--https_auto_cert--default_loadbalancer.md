---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2035, "body_sha256": "sha256:25f9dd9ccc614e5550eba44860ae565f4c4523df1d5cfdfe62c9c5958a8ee05d", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:default_loadbalancer", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:default_loadbalancer", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--default_loadbalancer.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "default_loadbalancer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/default_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](resources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- [xcsh_workload](../resources/workload.md)
