---
page_title: "ingress_gw.allowed_vip_port"
subcategory: "Infrastructure"
description: "ingress_gw.allowed_vip_port for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3488, "body_sha256": "sha256:1f0d07d002c632b1e1be824e7cec682b3f59cfa125114372fcc71a467625ed9c", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:allowed_vip_port", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw", "path": "documentation/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/index.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["ingress_gw", "allowed_vip_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw.allowed_vip_port for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.allowed_vip_port

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/)
- ingress_gw.allowed_vip_port

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

- [custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/custom_ports/): complete subsection reference.

- [disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/disable_allowed_vip_port/): complete subsection reference.

- [use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_https_port/): complete subsection reference.

- [use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_port/): complete subsection reference.

- [use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_https_port/): complete subsection reference.

## Next pages

- [ingress_gw.allowed_vip_port.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/custom_ports/)
- [ingress_gw.allowed_vip_port.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/disable_allowed_vip_port/)
- [ingress_gw.allowed_vip_port.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_https_port/)
- [ingress_gw.allowed_vip_port.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_port/)
- [ingress_gw.allowed_vip_port.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_https_port/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
