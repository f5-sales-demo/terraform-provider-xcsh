---
page_title: "direct_connect_enabled.hosted_vifs.site_registration_over_internet"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["direct connect enabled hosted vifs site registration over internet"], "body_bytes": 1265, "body_sha256": "sha256:88044f5bf351bfcd20b0f2f3dc11334d4bc8814e223b3f886c92144ba3ba3138", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "path": "documentation/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2233303031303323-2200031230302211-1120123022102133-0030301332213231-3313302220320011-2101210010103233-1312233312211302-3332120200130330", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_internet"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.hosted_vifs.site_registration_over_internet

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/)
- [direct_connect_enabled.hosted_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

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
site_registration_over_internet = {}
```

This is an empty object or choice marker. It has no direct properties.
