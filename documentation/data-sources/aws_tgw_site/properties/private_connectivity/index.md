---
page_title: "private_connectivity"
subcategory: ""
description: "private_connectivity for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2031, "body_sha256": "sha256:421b2740d95248dd0a6b8143287f47c8904be1566828981a0bf6f2d49b2c730d", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:cloud_link", "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:inside", "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:outside"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "documentation/data-sources/aws_tgw_site/properties/private_connectivity/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["private_connectivity"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/private_connectivity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_connectivity for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- private_connectivity

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for private connectivity.

Upstream description:

Private Connect Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

## Direct properties

- [cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/cloud_link/): complete subsection reference.

- [inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/inside/): complete subsection reference.

- [outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/outside/): complete subsection reference.

## Next pages

- [private_connectivity.cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/cloud_link/)
- [private_connectivity.inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/inside/)
- [private_connectivity.outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/outside/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
