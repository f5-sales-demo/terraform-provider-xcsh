---
page_title: "direct_connect_enabled.hosted_vifs"
subcategory: "Infrastructure"
description: "direct_connect_enabled.hosted_vifs for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1919, "body_sha256": "sha256:bc3eb47181017cca4ae282821de9c7cd3cb5ebbb1edc39e743cbcd7d764249b5", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:vif_list"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled", "path": "docs/guides/data-sources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "direct_connect_enabled.hosted_vifs for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# direct_connect_enabled.hosted_vifs

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [direct_connect_enabled](data-sources--aws_vpc_site--properties--direct_connect_enabled.md)
- direct_connect_enabled.hosted_vifs

<a id="section"></a>

Type: `"single"`. Computed.

AWS Direct Connect Hosted VIF Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

## Direct properties

- [site_registration_over_direct_connect](data-sources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md): complete subsection reference.

- [site_registration_over_internet](data-sources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_internet.md): complete subsection reference.

- [vif_list](data-sources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md): complete subsection reference.

## Next pages

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_internet.md)
- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md)
- [direct_connect_enabled](data-sources--aws_vpc_site--properties--direct_connect_enabled.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
