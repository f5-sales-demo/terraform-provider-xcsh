---
page_title: "dns_ntp_config"
subcategory: ""
description: "dns_ntp_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1896, "body_sha256": "sha256:1c0547538a7f255c5662b2ea5d21ea968f10ec1fdd81428b11a0a877feac23a4", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "docs/guides/data-sources--securemesh_site_v2--properties--dns_ntp_config.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dns_ntp_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/dns_ntp_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dns_ntp_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_ntp_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- dns_ntp_config

<a id="section"></a>

Type: `"single"`. Computed.

Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_server_choice": "[\"custom_dns\",\"f5_dns_default\"]",
  "x-ves-oneof-field-ntp_server_choice": "[\"custom_ntp\",\"f5_ntp_default\"]"
}
```

## Direct properties

- [custom_dns](data-sources--securemesh_site_v2--properties--dns_ntp_config--custom_dns.md): complete subsection reference.

- [custom_ntp](data-sources--securemesh_site_v2--properties--dns_ntp_config--custom_ntp.md): complete subsection reference.

- [f5_dns_default](data-sources--securemesh_site_v2--properties--dns_ntp_config--f5_dns_default.md): complete subsection reference.

- [f5_ntp_default](data-sources--securemesh_site_v2--properties--dns_ntp_config--f5_ntp_default.md): complete subsection reference.

## Next pages

- [dns_ntp_config.custom_dns](data-sources--securemesh_site_v2--properties--dns_ntp_config--custom_dns.md)
- [dns_ntp_config.custom_ntp](data-sources--securemesh_site_v2--properties--dns_ntp_config--custom_ntp.md)
- [dns_ntp_config.f5_dns_default](data-sources--securemesh_site_v2--properties--dns_ntp_config--f5_dns_default.md)
- [dns_ntp_config.f5_ntp_default](data-sources--securemesh_site_v2--properties--dns_ntp_config--f5_ntp_default.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
