---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1316, "body_sha256": "sha256:b160eeeda103ec3e932f13a0fd6a9cec9b8f974e39b31e5888bb9ba77f93b4bb", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:873fdc1e55067d3a5c09e53e53d5c6c7bcd1d7d8f83e917fef809f43dd1b0325", "source_path": "examples/resources/xcsh_waf_exclusion_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:waf_exclusion_policy:example:resource", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:examples", "path": "documentation/resources/waf_exclusion_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2230322233122230-1101312320013330-3203011100003232-1113232203023323-1012202020311211-1322303313201112-3213313330311311-2322010330201213", "registry_path": "docs/guides/resources--waf_exclusion_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_waf_exclusion_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
