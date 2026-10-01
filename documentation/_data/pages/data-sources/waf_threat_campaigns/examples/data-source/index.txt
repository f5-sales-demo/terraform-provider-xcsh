---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_threat_campaigns."
xcsh_docs: {"aliases": [], "body_bytes": 1264, "body_sha256": "sha256:ced1462f159ad6c859f4ec482fdc9b71fce509120d2417f784d6e2d087e3b440", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:87da0e1071c2bec565e2d7e725d9e896914625d820293d688bc4d4b1d617b7b5", "source_path": "examples/data-sources/xcsh_waf_threat_campaigns/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_threat_campaigns:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_threat_campaigns:examples", "path": "documentation/data-sources/waf_threat_campaigns/examples/data-source/index.md", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_waf_threat_campaigns.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_waf_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_threat_campaigns/data-source.tf`; digest `sha256:87da0e1071c2bec565e2d7e725d9e896914625d820293d688bc4d4b1d617b7b5`.

```terraform
# WAFThreatCampaigns DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threat_campaigns" "example" {
}

output "waf_threat_campaigns_result" {
  value = data.xcsh_waf_threat_campaigns.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/examples/)
- [xcsh_waf_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/)
