---
page_title: "egress_gateway_default"
subcategory: "Infrastructure"
description: "egress_gateway_default for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1414, "body_sha256": "sha256:849e299046d53e0b5e2591883840c2a26222198fa84c8678ee48cc6f1e924c3b", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:egress_gateway_default", "child_ids": [], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:egress_gateway_default", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:reference", "path": "docs/guides/data-sources--aws_vpc_site--properties--egress_gateway_default.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["egress_gateway_default"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/egress_gateway_default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_gateway_default for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_gateway_default

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- egress_gateway_default

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: egress\_gateway\_default, egress\_nat\_gw, egress\_virtual\_private\_gateway; Default:
egress\_gateway\_default\] Configuration parameter for egress gateway default.

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

OneOf alternatives in this subsection:

- [egress_gateway_default](data-sources--aws_vpc_site--properties--egress_gateway_default.md#section)
- [egress_nat_gw](data-sources--aws_vpc_site--properties--egress_nat_gw.md#section)
- [egress_virtual_private_gateway](data-sources--aws_vpc_site--properties--egress_virtual_private_gateway.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
