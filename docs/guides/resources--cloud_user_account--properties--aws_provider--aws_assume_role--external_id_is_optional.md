---
page_title: "aws_provider.aws_assume_role.external_id_is_optional"
subcategory: ""
description: "aws_provider.aws_assume_role.external_id_is_optional for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 1261, "body_sha256": "sha256:015443105c183d11b9246218cb879d2c8f4ff5983b547993459ec78b8b9e3b58", "canonical_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "parent_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "path": "docs/guides/resources--cloud_user_account--properties--aws_provider--aws_assume_role--external_id_is_optional.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_assume_role", "external_id_is_optional"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_optional/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_assume_role.external_id_is_optional for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_assume_role.external_id_is_optional

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
- [Property reference](resources--cloud_user_account--reference.md)
- [aws_provider](resources--cloud_user_account--properties--aws_provider.md)
- [aws_provider.aws_assume_role](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md)
- aws_provider.aws_assume_role.external_id_is_optional

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for external id is optional.

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
external_id_is_optional = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [aws_provider.aws_assume_role](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
