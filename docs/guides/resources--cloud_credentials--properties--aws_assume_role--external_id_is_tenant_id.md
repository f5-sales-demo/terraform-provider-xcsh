---
page_title: "aws_assume_role.external_id_is_tenant_id"
subcategory: "Infrastructure"
description: "aws_assume_role.external_id_is_tenant_id for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1066, "body_sha256": "sha256:a71464411b85a640089ff62084e1e813148aa6a8e67f87c72eda5d2a4be13df9", "canonical_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "path": "docs/guides/resources--cloud_credentials--properties--aws_assume_role--external_id_is_tenant_id.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_assume_role", "external_id_is_tenant_id"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/aws_assume_role/external_id_is_tenant_id/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_assume_role.external_id_is_tenant_id for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_assume_role.external_id_is_tenant_id

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
- [Property reference](resources--cloud_credentials--reference.md)
- [aws_assume_role](resources--cloud_credentials--properties--aws_assume_role.md)
- aws_assume_role.external_id_is_tenant_id

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
external_id_is_tenant_id = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [aws_assume_role](resources--cloud_credentials--properties--aws_assume_role.md)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
