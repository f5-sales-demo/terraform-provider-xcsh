---
page_title: "aws_assume_role.external_id_is_optional"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["aws assume role external id is optional"], "body_bytes": 1020, "body_sha256": "sha256:014994360b13cbed0a9ef3f1daabeff62f74688b7b7c0bfadc98fe3c2cb93959", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "parent_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role", "path": "documentation/data-sources/cloud_credentials/properties/aws_assume_role/external_id_is_optional/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1111312120100210-0223030012102003-2002211112110022-1223030110102223-1113200123003201-1323023301021000-1202030011112320-2131312112312103", "registry_path": "docs/guides/data-sources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_assume_role", "external_id_is_optional"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/aws_assume_role/external_id_is_optional/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "This can be used for messages where no values are needed.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_assume_role.external_id_is_optional

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_assume_role/)
- aws_assume_role.external_id_is_optional

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external id is optional.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
