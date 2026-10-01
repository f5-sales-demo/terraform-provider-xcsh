---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1110, "body_sha256": "sha256:c418007f4ec19cc3c7b1d9dce051b046674fbbad39b1da52766ca423568f9026", "canonical_id": "xcsh-docs:resources:waf_exclusion_policy:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:873fdc1e55067d3a5c09e53e53d5c6c7bcd1d7d8f83e917fef809f43dd1b0325", "source_path": "examples/resources/xcsh_waf_exclusion_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:waf_exclusion_policy:example:resource", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:examples", "path": "docs/guides/resources--waf_exclusion_policy--example--resource.md", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_waf_exclusion_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md)
- [Examples](resources--waf_exclusion_policy--examples.md)
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

## Next pages

- [Examples](resources--waf_exclusion_policy--examples.md)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md)
