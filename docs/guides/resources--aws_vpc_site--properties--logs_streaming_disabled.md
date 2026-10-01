---
page_title: "logs_streaming_disabled"
subcategory: "Infrastructure"
description: "logs_streaming_disabled for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 904, "body_sha256": "sha256:782436800ce114603854bf93af6334f96632060b0835c517d39380e6aaa14c77", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:logs_streaming_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:logs_streaming_disabled", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "docs/guides/resources--aws_vpc_site--properties--logs_streaming_disabled.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["logs_streaming_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/logs_streaming_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "logs_streaming_disabled for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# logs_streaming_disabled

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- logs_streaming_disabled

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
logs_streaming_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
