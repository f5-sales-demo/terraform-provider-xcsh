---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_service_policy_rule."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1368, "body_sha256": "sha256:e294d521e01754bb363c86cfcd9abf58ff43e41720b7487c4e74f2686e6bf882", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ae57799b0f1696fd5e6a11587ed3b1a1bc30f50fe8953264b9b955e84a34c726", "source_path": "examples/data-sources/xcsh_service_policy_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:service_policy_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:service_policy_rule:examples", "path": "documentation/data-sources/service_policy_rule/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3320121011320201-2020022320330013-1010121120101230-3200001120211013-2022023012201323-0333232213320213-2103310221122313-3011332222023310", "registry_path": "docs/guides/data-sources--service_policy_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_service_policy_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy_rule/data-source.tf`; digest `sha256:ae57799b0f1696fd5e6a11587ed3b1a1bc30f50fe8953264b9b955e84a34c726`.

```terraform
# ServicePolicyRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicyRule by name
data "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"
}

output "service_policy_rule_id" {
  value = data.xcsh_service_policy_rule.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/examples/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
