---
page_title: "vn_config.allowed_vip_port"
subcategory: ""
description: "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site."
xcsh_docs: {"aliases": ["vn config allowed vip port"], "body_bytes": 2074, "body_sha256": "sha256:4fcd335dd07cb102125df2b3dcccd067d19a618b2ed8be074d4bb663acdf5325", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:disable_allowed_vip_port", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_https_port", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_https_port"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config", "path": "documentation/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "allowed_vip_port"], "schema_version": 1, "sections": [{"aliases": ["vn config allowed vip port custom ports"], "anchor": "section", "description": "List of Custom port.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "custom_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config allowed vip port disable allowed vip port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:disable_allowed_vip_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "disable_allowed_vip_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config allowed vip port use http https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_https_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "use_http_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config allowed vip port use http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "use_http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config allowed vip port use https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_https_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "use_https_port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.allowed_vip_port

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/)
- vn_config.allowed_vip_port

<a id="section"></a>

Type: `"single"`. Computed.

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

## Direct properties

- [custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/custom_ports/): complete subsection reference.

- [disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/disable_allowed_vip_port/): complete subsection reference.

- [use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_https_port/): complete subsection reference.

- [use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_port/): complete subsection reference.

- [use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_https_port/): complete subsection reference.
