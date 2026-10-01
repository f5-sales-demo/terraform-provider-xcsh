---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2351, "body_sha256": "sha256:a0f99c0a4cf2f33e0ce1fe429da9ab48bd421d4e01a3f5ebe8e7f2b7b9898ed9", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:coalescing_options:strict_coalescing", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:coalescing_options:strict_coalescing", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:coalescing_options", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--coalescing_options--strict_coalescing.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "coalescing_options", "strict_coalescing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/coalescing_options/strict_coalescing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](resources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--coalescing_options.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--coalescing_options.md)
- [xcsh_workload](../resources/workload.md)
