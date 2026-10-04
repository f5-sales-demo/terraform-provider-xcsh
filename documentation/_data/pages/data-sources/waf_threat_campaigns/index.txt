---
page_title: "xcsh_waf_threat_campaigns"
subcategory: ""
description: "Resource retrieval operation."
xcsh_docs: {"aliases": ["waf threat campaigns"], "body_bytes": 1222, "body_sha256": "sha256:2793f77cad1fbf6186431c1b8c54b1ed0efb11d42c8b02973dc668327f904954", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_threat_campaigns:reference", "xcsh-docs:data-sources:waf_threat_campaigns:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threat_campaigns:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/waf_threat_campaigns/index.md", "product": "distributed-cloud", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1323321121302311-1032111022111323-1120132100221231-1221202233212221-1010312121000210-3221113220023012-3001030210113301-0002023032103021", "registry_path": "docs/data-sources/waf_threat_campaigns.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource retrieval operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/examples/)
