---
page_title: "tgw_security.no_forward_proxy"
subcategory: ""
description: "tgw_security.no_forward_proxy for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1016, "body_sha256": "sha256:51bf24707891e739b592f6d7857e3775e5e2e98685789c39e03162dfd1049454", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "path": "docs/guides/resources--aws_tgw_site--properties--tgw_security--no_forward_proxy.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tgw_security", "no_forward_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/tgw_security/no_forward_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tgw_security.no_forward_proxy for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security.no_forward_proxy

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [tgw_security](resources--aws_tgw_site--properties--tgw_security.md)
- tgw_security.no_forward_proxy

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tgw_security](resources--aws_tgw_site--properties--tgw_security.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
