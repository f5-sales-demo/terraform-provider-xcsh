---
page_title: "xcsh_waf_exclusion_policy"
subcategory: ""
description: "Manages WAF exclusion policy in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["waf exclusion policy"], "body_bytes": 1378, "body_sha256": "sha256:ba068929b07debe196711eb23b3f06fe1a48c130f6bdec7e664916e40a4b60af", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:waf_exclusion_policy:reference", "xcsh-docs:data-sources:waf_exclusion_policy:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_exclusion_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/waf_exclusion_policy/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203", "registry_path": "docs/data-sources/waf_exclusion_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_exclusion_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages WAF exclusion policy in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_waf_exclusion_policy

Breadcrumbs:

- xcsh_waf_exclusion_policy

Manages WAF exclusion policy in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFExclusionPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WAFExclusionPolicy by name
data "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}

output "waf_exclusion_policy_id" {
  value = data.xcsh_waf_exclusion_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/examples/)
