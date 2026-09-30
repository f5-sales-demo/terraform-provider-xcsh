---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_threat_campaigns."
xcsh_docs: {"aliases": [], "body_bytes": 959, "body_sha256": "sha256:93f63a2d97485b539c7a76bf3447186b71fbca222a7925b14d9af1c6eb3bcd1c", "canonical_id": "xcsh-docs:data-sources:waf_threat_campaigns:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:87da0e1071c2bec565e2d7e725d9e896914625d820293d688bc4d4b1d617b7b5", "source_path": "examples/data-sources/xcsh_waf_threat_campaigns/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_threat_campaigns:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_threat_campaigns:examples", "path": "docs/guides/data-sources--waf_threat_campaigns--example--data-source.md", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_waf_threat_campaigns.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md)
- [Examples](data-sources--waf_threat_campaigns--examples.md)
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

- [Examples](data-sources--waf_threat_campaigns--examples.md)
- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md)
