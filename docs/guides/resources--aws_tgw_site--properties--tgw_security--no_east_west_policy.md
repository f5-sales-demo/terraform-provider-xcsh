---
page_title: "tgw_security.no_east_west_policy"
subcategory: ""
description: "tgw_security.no_east_west_policy for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1018, "body_sha256": "sha256:e8b197e0c29c93536f97f7447a00a9b1a66478f069079aeefc6f946bfa82d2a1", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "path": "docs/guides/resources--aws_tgw_site--properties--tgw_security--no_east_west_policy.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tgw_security", "no_east_west_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/tgw_security/no_east_west_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tgw_security.no_east_west_policy for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security.no_east_west_policy

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [tgw_security](resources--aws_tgw_site--properties--tgw_security.md)
- tgw_security.no_east_west_policy

<a id="section"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_east_west_policy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tgw_security](resources--aws_tgw_site--properties--tgw_security.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
