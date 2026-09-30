---
page_title: "xcsh_waf_threat_campaigns"
subcategory: ""
description: "xcsh_waf_threat_campaigns for xcsh_waf_threat_campaigns."
xcsh_docs: {"aliases": [], "body_bytes": 1038, "body_sha256": "sha256:414bdfd3028af8a3f2c49350e31844517b6e8590bf004b615f5ff975e4eaf3c2", "canonical_id": "xcsh-docs:data-sources:waf_threat_campaigns:fundamentals", "child_ids": ["xcsh-docs:data-sources:waf_threat_campaigns:reference", "xcsh-docs:data-sources:waf_threat_campaigns:examples"], "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threat_campaigns:fundamentals", "parent_id": null, "path": "docs/data-sources/waf_threat_campaigns.md", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_waf_threat_campaigns for xcsh_waf_threat_campaigns.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_waf_threat_campaigns

Breadcrumbs:

- xcsh_waf_threat_campaigns

Resource retrieval operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--waf_threat_campaigns--reference.md)
- [Examples](../guides/data-sources--waf_threat_campaigns--examples.md)
