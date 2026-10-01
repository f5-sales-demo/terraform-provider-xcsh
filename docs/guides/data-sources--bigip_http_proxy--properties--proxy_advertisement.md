---
page_title: "proxy_advertisement"
subcategory: ""
description: "proxy_advertisement for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1407, "body_sha256": "sha256:c10b5880008f5977935bd22a8d207a40fe41c7893f733aa6f431722c6610f1d6", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:do_not_advertise"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "path": "docs/guides/data-sources--bigip_http_proxy--properties--proxy_advertisement.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_advertisement/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- proxy_advertisement

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"do_not_advertise\"]"
}
```

## Direct properties

- [advertise_custom](data-sources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom.md): complete subsection reference.

- [do_not_advertise](data-sources--bigip_http_proxy--properties--proxy_advertisement--do_not_advertise.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom.md)
- [proxy_advertisement.do_not_advertise](data-sources--bigip_http_proxy--properties--proxy_advertisement--do_not_advertise.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
