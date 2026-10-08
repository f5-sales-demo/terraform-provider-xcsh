---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1070, "body_sha256": "sha256:c3b79157ee1611742fc8a15421c5fe02e6f38b01db7ed1bfc5e1d1925b5b0cf6", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:873fdc1e55067d3a5c09e53e53d5c6c7bcd1d7d8f83e917fef809f43dd1b0325", "source_path": "examples/resources/xcsh_waf_exclusion_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:waf_exclusion_policy:example:resource", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:examples", "path": "documentation/resources/waf_exclusion_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2230322233122230-1101312320013330-3203011100003232-1113232203023323-1012202020311211-1322303313201112-3213313330311311-2322010330201213", "registry_path": "docs/guides/resources--waf_exclusion_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_waf_exclusion_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_waf_exclusion_policy/resource.tf`; digest `sha256:873fdc1e55067d3a5c09e53e53d5c6c7bcd1d7d8f83e917fef809f43dd1b0325`.

```terraform
# WAFExclusionPolicy Resource Example
# Manages WAF exclusion policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WAFExclusionPolicy configuration
resource "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}
```
