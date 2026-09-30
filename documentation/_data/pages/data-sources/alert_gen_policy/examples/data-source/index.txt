---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_alert_gen_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1230, "body_sha256": "sha256:83fd40a6b86abc2393c43a5dbefca0622ab046a72c9e37994c69fcf3eee5dc91", "child_ids": [], "collection_id": "xcsh-docs:data-sources:alert_gen_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:349cbe33e45f53c629e7ce220830862d4e4d7dd9b1fd6b5c5a42c3556d1a1bed", "source_path": "examples/data-sources/xcsh_alert_gen_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:alert_gen_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:alert_gen_policy:examples", "path": "documentation/data-sources/alert_gen_policy/examples/data-source/index.md", "provider_name": "alert_gen_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_gen_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_alert_gen_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_gen_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_alert_gen_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_gen_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_gen_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_gen_policy/data-source.tf`; digest `sha256:349cbe33e45f53c629e7ce220830862d4e4d7dd9b1fd6b5c5a42c3556d1a1bed`.

```terraform
# AlertGenPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertGenPolicy by name
data "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}

output "alert_gen_policy_id" {
  value = data.xcsh_alert_gen_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_gen_policy/examples/)
- [xcsh_alert_gen_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_gen_policy/)
