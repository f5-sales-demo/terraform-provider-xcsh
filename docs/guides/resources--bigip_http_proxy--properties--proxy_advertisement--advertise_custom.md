---
page_title: "proxy_advertisement.advertise_custom"
subcategory: ""
description: "proxy_advertisement.advertise_custom for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1444, "body_sha256": "sha256:c6057c738fe915430bdf92854628242be7b37d393809ab8e9b564ed6011c81f0", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.advertise_custom for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_advertisement.advertise_custom

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_advertisement](resources--bigip_http_proxy--properties--proxy_advertisement.md)
- proxy_advertisement.advertise_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
```

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
advertise_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_where](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- [proxy_advertisement](resources--bigip_http_proxy--properties--proxy_advertisement.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
