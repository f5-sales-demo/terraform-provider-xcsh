---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing"
subcategory: ""
description: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1941, "body_sha256": "sha256:7725e8b26731c2967e2f69bdf2ba23058b79f765c2ecf6546dee5fe58da78fd0", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing:route_tables"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "path": "docs/guides/resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "custom_routing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [aws_provider](resources--cloud_connect--properties--aws_provider.md)
- [aws_provider.aws_tgw_site](resources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS Route Table List. AWS Route Table List.

Upstream description:

AWS Route Table List.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("route_tables")}
```

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
custom_routing {
  # Configure direct properties listed below.
}
```

## Direct properties

- [route_tables](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables.md): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
