---
page_title: "advertise_custom.advertise_where.vk8s_service"
subcategory: "Load Balancing"
description: "advertise_custom.advertise_where.vk8s_service for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2227, "body_sha256": "sha256:79bce8f6fadf1f725821a09000035887eef1a8218fef986087372664412519ab", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service:site", "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service:virtual_site"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where", "path": "docs/guides/resources--http_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "vk8s_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.vk8s_service for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.vk8s_service

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [advertise_custom](resources--http_loadbalancer--properties--advertise_custom.md)
- [advertise_custom.advertise_where](resources--http_loadbalancer--properties--advertise_custom--advertise_where.md)
- advertise_custom.advertise_where.vk8s_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--http_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.vk8s_service.site](resources--http_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md)
- [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--http_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md)
- [advertise_custom.advertise_where](resources--http_loadbalancer--properties--advertise_custom--advertise_where.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
