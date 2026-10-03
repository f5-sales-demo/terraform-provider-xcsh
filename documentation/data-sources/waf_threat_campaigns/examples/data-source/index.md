---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_threat_campaigns."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1264, "body_sha256": "sha256:ced1462f159ad6c859f4ec482fdc9b71fce509120d2417f784d6e2d087e3b440", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:87da0e1071c2bec565e2d7e725d9e896914625d820293d688bc4d4b1d617b7b5", "source_path": "examples/data-sources/xcsh_waf_threat_campaigns/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_threat_campaigns:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_threat_campaigns:examples", "path": "documentation/data-sources/waf_threat_campaigns/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0103000131000021-1121211022313202-3220322000022023-1100212322033323-3001103323211032-3030233103300230-0312310112231220-0133320133131230", "registry_path": "docs/guides/data-sources--waf_threat_campaigns--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_waf_threat_campaigns.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
