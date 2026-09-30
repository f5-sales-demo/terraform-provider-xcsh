---
page_title: "direct_connect_enabled.hosted_vifs"
subcategory: ""
description: "direct_connect_enabled.hosted_vifs for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2470, "body_sha256": "sha256:c4c9b43747662e2f663454be50983bca4b6c3b9494fce954d5357b49b395c7d3", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled", "path": "documentation/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "direct_connect_enabled.hosted_vifs for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# direct_connect_enabled.hosted_vifs

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/)
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

- [site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/): complete subsection reference.

- [site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/): complete subsection reference.

- [vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/): complete subsection reference.

## Next pages

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/)
- [direct_connect_enabled.hosted_vifs.vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
