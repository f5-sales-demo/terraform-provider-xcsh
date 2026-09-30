---
page_title: "proxy_advertisement.advertise_custom.advertise_where.vk8s_service"
subcategory: ""
description: "proxy_advertisement.advertise_custom.advertise_where.vk8s_service for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2492, "body_sha256": "sha256:65522bca225f5c6afd19298e4390d73e3169ecde9dde8a836ec8f55fee65da82", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:site", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.advertise_custom.advertise_where.vk8s_service for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_advertisement.advertise_custom.advertise_where.vk8s_service

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_advertisement](resources--bigip_http_proxy--properties--proxy_advertisement.md)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom.md)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

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

- [site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
