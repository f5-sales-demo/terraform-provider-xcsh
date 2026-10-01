---
page_title: "voltstack_cluster.allowed_vip_port"
subcategory: "Infrastructure"
description: "voltstack_cluster.allowed_vip_port for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2888, "body_sha256": "sha256:1396effab39adb92f3b5adf96df20f8b9b6f7438052a568467d506a3510c6bc9", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port:custom_ports", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port:disable_allowed_vip_port", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port:use_http_https_port", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port:use_http_port", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port:use_https_port"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "path": "docs/guides/data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "allowed_vip_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.allowed_vip_port for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.allowed_vip_port

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [voltstack_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster.md)
- voltstack_cluster.allowed_vip_port

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

- [custom_ports](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--custom_ports.md): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--disable_allowed_vip_port.md): complete subsection reference.

- [use_http_https_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_http_https_port.md): complete subsection reference.

- [use_http_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_http_port.md): complete subsection reference.

- [use_https_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_https_port.md): complete subsection reference.

## Next pages

- [voltstack_cluster.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--custom_ports.md)
- [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--disable_allowed_vip_port.md)
- [voltstack_cluster.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_http_https_port.md)
- [voltstack_cluster.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_http_port.md)
- [voltstack_cluster.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_https_port.md)
- [voltstack_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
