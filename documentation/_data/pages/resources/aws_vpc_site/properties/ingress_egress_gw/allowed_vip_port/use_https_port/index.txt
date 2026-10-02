---
page_title: "ingress_egress_gw.allowed_vip_port.use_https_port"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ingress egress gw allowed vip port use https port"], "body_bytes": 1504, "body_sha256": "sha256:0cb0104c13bfb3bc7aabfc37bce4b6ab292f14dd3cee7b03f1af98e48aff6f60", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port:use_https_port", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port", "path": "documentation/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_https_port/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3121101102110133-3103111202111010-2200011313000113-2302111201131323-3330111310112123-3110223103221320-2203231321331232-0103101223320201", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "allowed_vip_port", "use_https_port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_https_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.allowed_vip_port.use_https_port

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/)
- ingress_egress_gw.allowed_vip_port.use_https_port

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
use_https_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
