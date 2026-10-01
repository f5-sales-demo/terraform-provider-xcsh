---
page_title: "tgw_security.active_east_west_service_policies"
subcategory: ""
description: "tgw_security.active_east_west_service_policies for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1191, "body_sha256": "sha256:8eae04bdb5e70174e42812c7d68b602bf40467540baf3836a9f1e5e94d716d14", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies:service_policies"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security", "path": "docs/guides/data-sources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tgw_security", "active_east_west_service_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tgw_security.active_east_west_service_policies for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security.active_east_west_service_policies

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [tgw_security](data-sources--aws_tgw_site--properties--tgw_security.md)
- tgw_security.active_east_west_service_policies

<a id="section"></a>

Type: `"single"`. Computed.

Active service policies for the east-west proxy.

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

## Direct properties

- [service_policies](data-sources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md): complete subsection reference.

## Next pages

- [tgw_security.active_east_west_service_policies.service_policies](data-sources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md)
- [tgw_security](data-sources--aws_tgw_site--properties--tgw_security.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
