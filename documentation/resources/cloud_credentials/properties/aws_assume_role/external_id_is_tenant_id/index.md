---
page_title: "aws_assume_role.external_id_is_tenant_id"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["aws assume role external id is tenant id"], "body_bytes": 1323, "body_sha256": "sha256:80034dcbd67c5e6174ac0bca041bc28d6be61d7bb685afcb9832fa870d3a78ac", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "path": "documentation/resources/cloud_credentials/properties/aws_assume_role/external_id_is_tenant_id/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2203103211303300-3032003222210131-0131212122013230-0313303303201103-1023331120233333-1313220201120323-1032330303131223-0003332300111312", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_assume_role", "external_id_is_tenant_id"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/aws_assume_role/external_id_is_tenant_id/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_assume_role.external_id_is_tenant_id

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/)
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

- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
