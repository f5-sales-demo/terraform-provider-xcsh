---
page_title: "aws_provider"
subcategory: ""
description: "aws_provider for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1346, "body_sha256": "sha256:5f1f4923d9f7887e157ce6baf85e614d6c8a329d38029c337f703b87844f66ea", "canonical_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider", "child_ids": ["xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site"], "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider", "parent_id": "xcsh-docs:data-sources:cloud_connect:reference", "path": "docs/guides/data-sources--cloud_connect--properties--aws_provider.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/aws_provider/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
- [Property reference](data-sources--cloud_connect--reference.md)
- aws_provider

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: aws\_provider, azure\_vnet\_site\] Configuration parameter for aws provider.

Upstream description:

Cloud Connect with AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_type": "[\"aws_tgw_site\"]"
}
```

OneOf alternatives in this subsection:

- [aws_provider](data-sources--cloud_connect--properties--aws_provider.md#section)
- [azure_vnet_site](data-sources--cloud_connect--properties--azure_vnet_site.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [aws_tgw_site](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site.md): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [Property reference](data-sources--cloud_connect--reference.md)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
