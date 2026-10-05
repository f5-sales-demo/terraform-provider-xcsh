---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1316, "body_sha256": "sha256:b160eeeda103ec3e932f13a0fd6a9cec9b8f974e39b31e5888bb9ba77f93b4bb", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:873fdc1e55067d3a5c09e53e53d5c6c7bcd1d7d8f83e917fef809f43dd1b0325", "source_path": "examples/resources/xcsh_waf_exclusion_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:waf_exclusion_policy:example:resource", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:examples", "path": "documentation/resources/waf_exclusion_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2230322233122230-1101312320013330-3203011100003232-1113232203023323-1012202020311211-1322303313201112-3213313330311311-2322010330201213", "registry_path": "docs/guides/resources--waf_exclusion_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_waf_exclusion_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/examples/)
- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/)
