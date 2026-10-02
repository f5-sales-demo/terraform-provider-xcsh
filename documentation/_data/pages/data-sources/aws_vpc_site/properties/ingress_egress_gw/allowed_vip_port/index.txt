---
page_title: "ingress_egress_gw.allowed_vip_port"
subcategory: "Infrastructure"
description: "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site."
xcsh_docs: {"aliases": ["ingress egress gw allowed vip port"], "body_bytes": 3635, "body_sha256": "sha256:25af691ebd07a0cc2ad11597713ac50ba53e953c3cb56e05812ac89c16e42a2e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:custom_ports", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:disable_allowed_vip_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:use_http_https_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:use_http_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:use_https_port"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw", "path": "documentation/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "allowed_vip_port"], "schema_version": 1, "sections": [{"aliases": ["custom ports"], "anchor": "section", "description": "List of Custom port.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:custom_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port", "custom_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable allowed vip port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:disable_allowed_vip_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port", "disable_allowed_vip_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use http https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:use_http_https_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port", "use_http_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:use_http_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port", "use_http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:use_https_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port", "use_https_port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.allowed_vip_port

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- ingress_egress_gw.allowed_vip_port

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

- [custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/custom_ports/): complete subsection reference.

- [disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/disable_allowed_vip_port/): complete subsection reference.

- [use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_http_https_port/): complete subsection reference.

- [use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_http_port/): complete subsection reference.

- [use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_https_port/): complete subsection reference.

## Next pages

- [ingress_egress_gw.allowed_vip_port.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/custom_ports/)
- [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/disable_allowed_vip_port/)
- [ingress_egress_gw.allowed_vip_port.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_http_https_port/)
- [ingress_egress_gw.allowed_vip_port.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_http_port/)
- [ingress_egress_gw.allowed_vip_port.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_https_port/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
