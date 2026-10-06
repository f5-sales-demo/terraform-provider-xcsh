---
page_title: "dns_ntp_config"
subcategory: ""
description: "Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site."
xcsh_docs: {"aliases": ["dns ntp config"], "body_bytes": 1602, "body_sha256": "sha256:2171f4426bde87d11675e9a5a849dd37a9a9f41ab3541c498a1b5ea34f7997b3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/dns_ntp_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3330210130131331-1111102213112232-3103010133222311-1323003033030132-2002312332303123-2122020221132202-3030100123303333-1322230230022212", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dns_ntp_config"], "schema_version": 1, "sections": [{"aliases": ["dns ntp config custom dns"], "anchor": "section", "description": "DNS Servers.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dns_ntp_config", "custom_dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["dns ntp config custom ntp"], "anchor": "section", "description": "NTP Servers.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dns_ntp_config", "custom_ntp"], "syntax": "attribute", "type": "object"}, {"aliases": ["dns ntp config f5 dns default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_ntp_config", "f5_dns_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["dns ntp config f5 ntp default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_ntp_config", "f5_ntp_default"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/dns_ntp_config/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
