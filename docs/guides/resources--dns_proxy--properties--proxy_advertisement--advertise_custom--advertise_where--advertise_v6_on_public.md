---
page_title: "proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public"
subcategory: ""
description: "proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1969, "body_sha256": "sha256:43f991109edc2b9e82fc3732d0df0b7789c73f832046433587f53e54d0910308", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public:public_ip"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public", "parent_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "docs/guides/resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_v6_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [proxy_advertisement](resources--dns_proxy--properties--proxy_advertisement.md)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--properties--proxy_advertisement--advertise_custom.md)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public_ip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
