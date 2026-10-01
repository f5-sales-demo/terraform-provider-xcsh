---
page_title: "aws_provider"
subcategory: ""
description: "aws_provider for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1443, "body_sha256": "sha256:2b8fff681ba15926f1319d1db195331d0c477ac9f2246086788e0e033268fb43", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider", "parent_id": "xcsh-docs:resources:cloud_connect:reference", "path": "docs/guides/resources--cloud_connect--properties--aws_provider.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- aws_provider

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [aws_provider](resources--cloud_connect--properties--aws_provider.md#section)
- [azure_vnet_site](resources--cloud_connect--properties--azure_vnet_site.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_provider {
  # Configure direct properties listed below.
}
```

## Direct properties

- [aws_tgw_site](resources--cloud_connect--properties--aws_provider--aws_tgw_site.md): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site](resources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [Property reference](resources--cloud_connect--reference.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
