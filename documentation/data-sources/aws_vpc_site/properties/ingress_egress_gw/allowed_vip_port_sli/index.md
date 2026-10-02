---
page_title: "ingress_egress_gw.allowed_vip_port_sli"
subcategory: "Infrastructure"
description: "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site."
xcsh_docs: {"aliases": ["ingress egress gw allowed vip port sli"], "body_bytes": 3703, "body_sha256": "sha256:429b13ac0a9f02c6c5f8a58a024d574e19fba050c12c2dcaa0ef6b7d8b3d4513", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:custom_ports", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:disable_allowed_vip_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_http_https_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_http_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_https_port"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw", "path": "documentation/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli"], "schema_version": 1, "sections": [{"aliases": ["custom ports"], "anchor": "section", "description": "List of Custom port.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:custom_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "custom_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable allowed vip port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:disable_allowed_vip_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "disable_allowed_vip_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use http https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_http_https_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "use_http_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_http_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "use_http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_https_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "use_https_port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.allowed_vip_port_sli

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- ingress_egress_gw.allowed_vip_port_sli

<a id="section"></a>

Type: `"single"`. Computed.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

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

- [custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/custom_ports/): complete subsection reference.

- [disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/disable_allowed_vip_port/): complete subsection reference.

- [use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_http_https_port/): complete subsection reference.

- [use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_http_port/): complete subsection reference.

- [use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_https_port/): complete subsection reference.

## Next pages

- [ingress_egress_gw.allowed_vip_port_sli.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/custom_ports/)
- [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/disable_allowed_vip_port/)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_http_https_port/)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_http_port/)
- [ingress_egress_gw.allowed_vip_port_sli.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_https_port/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
