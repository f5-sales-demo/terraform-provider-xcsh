---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1269, "body_sha256": "sha256:31d4a8561f444177aa70818d283dc8136a03527c433416ba9f970fad2fca68a2", "child_ids": [], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ae57799b0f1696fd5e6a11587ed3b1a1bc30f50fe8953264b9b955e84a34c726", "source_path": "examples/data-sources/xcsh_service_policy_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:service_policy_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:service_policy_rule:examples", "path": "documentation/data-sources/service_policy_rule/examples/data-source/index.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
