---
page_title: "aws_provider.aws_assume_role.external_id_is_optional"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["aws provider aws assume role external id is optional"], "body_bytes": 1567, "body_sha256": "sha256:5da90b7eb893991169261bd1558767c5f368629bb2ed90caa797640d2db9198a", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "parent_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "path": "documentation/resources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_optional/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0111331230012032-3300101331303302-0102302032121321-1132003103313323-0231031210303310-0003123003123011-2221210012032123-1000330030113332", "registry_path": "docs/guides/resources--cloud_user_account--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_assume_role", "external_id_is_optional"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_optional/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_assume_role.external_id_is_optional

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/)
- [aws_provider.aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_assume_role/)
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

- [aws_provider.aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_assume_role/)
- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
