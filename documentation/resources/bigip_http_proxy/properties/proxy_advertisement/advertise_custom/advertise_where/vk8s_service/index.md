---
page_title: "proxy_advertisement.advertise_custom.advertise_where.vk8s_service"
subcategory: ""
description: "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network."
xcsh_docs: {"aliases": ["proxy advertisement advertise custom advertise where vk8s service"], "body_bytes": 3133, "body_sha256": "sha256:97343a377603fc7304b80fa46638f9cd77aa64c4e2a88af7a733a836517b2401", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:site", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "documentation/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3201012210311222-0013322302303000-0033333010112020-1110200303031321-2310032110213210-2223222313100032-1222203010020321-1021221212131032", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--name", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:site", "type": "requires"}], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service", "site"], "syntax": "block", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site", "type": "requires"}], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.vk8s_service

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/)
- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
