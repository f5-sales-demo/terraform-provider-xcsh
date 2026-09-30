---
page_title: "aws_assume_role.external_id_is_optional"
subcategory: "Infrastructure"
description: "aws_assume_role.external_id_is_optional for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 998, "body_sha256": "sha256:e30d6aa593c71ce0cb038129298036584514406ee788a36d65d2bf18d0a61fff", "canonical_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "path": "docs/guides/resources--cloud_credentials--properties--aws_assume_role--external_id_is_optional.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_assume_role", "external_id_is_optional"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/aws_assume_role/external_id_is_optional/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_assume_role.external_id_is_optional for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
