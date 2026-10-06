---
page_title: "direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["direct connect enabled hosted vifs vif list same as site region"], "body_bytes": 1431, "body_sha256": "sha256:db70d6972361c3c498d84da114e300b9d0ea1547d06995daa2e6a5424b84db0f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list:same_as_site_region", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "path": "documentation/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/same_as_site_region/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1103203201203221-2232200121201212-3001332002021010-3011110001033223-3033312013300232-3010110101202312-2123331310303333-2203213223132331", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list", "same_as_site_region"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/same_as_site_region/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/)
- [direct_connect_enabled.hosted_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/)
- [direct_connect_enabled.hosted_vifs.vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
same_as_site_region = {}
```

This is an empty object or choice marker. It has no direct properties.
