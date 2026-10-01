---
page_title: "proxy_advertisement.do_not_advertise"
subcategory: ""
description: "proxy_advertisement.do_not_advertise for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1037, "body_sha256": "sha256:308349f346b340a8e7f389dc10722d6d964c4e75340d8ab1121d141ca1ddbb39", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:do_not_advertise", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:do_not_advertise", "parent_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement", "path": "docs/guides/resources--dns_proxy--properties--proxy_advertisement--do_not_advertise.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "do_not_advertise"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/proxy_advertisement/do_not_advertise/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.do_not_advertise for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.do_not_advertise

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [proxy_advertisement](resources--dns_proxy--properties--proxy_advertisement.md)
- proxy_advertisement.do_not_advertise

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [proxy_advertisement](resources--dns_proxy--properties--proxy_advertisement.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
