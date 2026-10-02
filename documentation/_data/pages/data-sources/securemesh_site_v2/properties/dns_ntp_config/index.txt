---
page_title: "dns_ntp_config"
subcategory: ""
description: "Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site."
xcsh_docs: {"aliases": ["dns ntp config"], "body_bytes": 2504, "body_sha256": "sha256:cab671d4275a97076f3986cd86d909f6390499712f3e66398d2aa515d21aca2d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/dns_ntp_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3330210130131331-1111102213112232-3103010133222311-1323003033030132-2002312332303123-2122020221132202-3030100123303333-1322230230022212", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dns_ntp_config"], "schema_version": 1, "sections": [{"aliases": ["custom dns"], "anchor": "section", "description": "DNS Servers.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dns_ntp_config", "custom_dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom ntp"], "anchor": "section", "description": "NTP Servers.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dns_ntp_config", "custom_ntp"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 dns default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_ntp_config", "f5_dns_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 ntp default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_ntp_config", "f5_ntp_default"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/dns_ntp_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_ntp_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
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

- [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/custom_dns/): complete subsection reference.

- [custom_ntp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/custom_ntp/): complete subsection reference.

- [f5_dns_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/f5_dns_default/): complete subsection reference.

- [f5_ntp_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/f5_ntp_default/): complete subsection reference.

## Next pages

- [dns_ntp_config.custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/custom_dns/)
- [dns_ntp_config.custom_ntp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/custom_ntp/)
- [dns_ntp_config.f5_dns_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/f5_dns_default/)
- [dns_ntp_config.f5_ntp_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/f5_ntp_default/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
