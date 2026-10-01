---
page_title: "aws_assume_role.external_id_is_optional"
subcategory: "Infrastructure"
description: "aws_assume_role.external_id_is_optional for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1097, "body_sha256": "sha256:8c05f9b2e060ddc95b8c4e4b68ef30abcfb6aa5a3f4cf2b5ae024de7ce902228", "canonical_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "path": "docs/guides/resources--cloud_credentials--properties--aws_assume_role--external_id_is_optional.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_assume_role", "external_id_is_optional"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/aws_assume_role/external_id_is_optional/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_assume_role.external_id_is_optional for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_assume_role.external_id_is_optional

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
- [Property reference](resources--cloud_credentials--reference.md)
- [aws_assume_role](resources--cloud_credentials--properties--aws_assume_role.md)
- aws_assume_role.external_id_is_optional

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

- [aws_assume_role](resources--cloud_credentials--properties--aws_assume_role.md)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
