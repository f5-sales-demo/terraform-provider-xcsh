---
page_title: "aws_provider.aws_tgw_site"
subcategory: ""
description: "aws_provider.aws_tgw_site for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1705, "body_sha256": "sha256:459b0960c7119e762abf6dfb5226c40642de7e464de38251e5c463c3d9e3e673", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:cred", "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:site", "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider", "path": "docs/guides/resources--cloud_connect--properties--aws_provider--aws_tgw_site.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [aws_provider](resources--cloud_connect--properties--aws_provider.md)
- aws_provider.aws_tgw_site

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS TGW Site Type. Cloud Connect AWS TGW Site Type.

Upstream description:

Cloud Connect AWS TGW Site Type.

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
aws_tgw_site {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cred](resources--cloud_connect--properties--aws_provider--aws_tgw_site--cred.md): complete subsection reference.

- [site](resources--cloud_connect--properties--aws_provider--aws_tgw_site--site.md): complete subsection reference.

- [vpc_attachments](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site.cred](resources--cloud_connect--properties--aws_provider--aws_tgw_site--cred.md)
- [aws_provider.aws_tgw_site.site](resources--cloud_connect--properties--aws_provider--aws_tgw_site--site.md)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- [aws_provider](resources--cloud_connect--properties--aws_provider.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
