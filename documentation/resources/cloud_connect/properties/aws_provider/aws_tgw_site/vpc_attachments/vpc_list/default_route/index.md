---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route"
subcategory: ""
description: "Select Override Default Route Choice."
xcsh_docs: {"aliases": ["aws provider aws tgw site vpc attachments vpc list default route"], "body_bytes": 3206, "body_sha256": "sha256:e3890c913cf2012eb08ff585609e45949983c3adaad42b62c9ba58f19a537863", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:all_route_tables", "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:selective_route_tables"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "path": "documentation/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3233310033321120-2110120300102121-1222212213121011-3211203323031022-0323101010102011-3032033101120022-1321001101212331-1323210021310313", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route:ConflictingObjectAttributes:all_route_tables,selective_route_tables", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:all_route_tables", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route:ConflictingObjectAttributes:all_route_tables,selective_route_tables", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:selective_route_tables", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "default_route"], "schema_version": 1, "sections": [{"aliases": ["all route tables"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:all_route_tables", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "default_route", "all_route_tables"], "syntax": "attribute", "type": "object"}, {"aliases": ["selective route tables"], "anchor": "section", "description": "AWS Route Table.", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:selective_route_tables", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "default_route", "selective_route_tables"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select Override Default Route Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- [aws_provider.aws_tgw_site.vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_route_tables",
    "selective_route_tables")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_route_choice": "[\"all_route_tables\",\"selective_route_tables\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/all_route_tables/): complete subsection reference.

- [selective_route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/selective_route_tables/): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/all_route_tables/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/selective_route_tables/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
