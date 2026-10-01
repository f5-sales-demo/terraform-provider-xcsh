---
page_title: "ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port"
subcategory: "Infrastructure"
description: "ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1228, "body_sha256": "sha256:a900964e2cce8a51cb5fbc9e41b9a5107bb10de695a33a36e2283378fe46fb25", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:disable_allowed_vip_port", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:disable_allowed_vip_port", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port", "path": "docs/guides/resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port--disable_allowed_vip_port.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "allowed_vip_port", "disable_allowed_vip_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/disable_allowed_vip_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port.md)
- ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_allowed_vip_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
