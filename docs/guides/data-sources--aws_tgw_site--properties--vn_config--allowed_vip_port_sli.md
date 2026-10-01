---
page_title: "vn_config.allowed_vip_port_sli"
subcategory: ""
description: "vn_config.allowed_vip_port_sli for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2788, "body_sha256": "sha256:26559320e0ea55c1309eecec9ae034cc32fc762f6a1778ea1e844a81f81f997a", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:disable_allowed_vip_port", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_https_port", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_port", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_https_port"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config", "path": "docs/guides/data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "allowed_vip_port_sli"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.allowed_vip_port_sli for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.allowed_vip_port_sli

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- vn_config.allowed_vip_port_sli

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

- [custom_ports](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--custom_ports.md): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--disable_allowed_vip_port.md): complete subsection reference.

- [use_http_https_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_http_https_port.md): complete subsection reference.

- [use_http_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_http_port.md): complete subsection reference.

- [use_https_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_https_port.md): complete subsection reference.

## Next pages

- [vn_config.allowed_vip_port_sli.custom_ports](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--custom_ports.md)
- [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--disable_allowed_vip_port.md)
- [vn_config.allowed_vip_port_sli.use_http_https_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_http_https_port.md)
- [vn_config.allowed_vip_port_sli.use_http_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_http_port.md)
- [vn_config.allowed_vip_port_sli.use_https_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_https_port.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
