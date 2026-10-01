---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1084, "body_sha256": "sha256:b116b68911cb267cee6532198340af939d51d7f71615dc55eee375d6ae222210", "canonical_id": "xcsh-docs:data-sources:fast_acl_rule:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d3a85aa2d2c4a98927959b827b07db2533ed58dbfde3c28ad82e3e2235d3a843", "source_path": "examples/data-sources/xcsh_fast_acl_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:fast_acl_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:fast_acl_rule:examples", "path": "docs/guides/data-sources--fast_acl_rule--example--data-source.md", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_fast_acl_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
- [Examples](data-sources--fast_acl_rule--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fast_acl_rule/data-source.tf`; digest `sha256:d3a85aa2d2c4a98927959b827b07db2533ed58dbfde3c28ad82e3e2235d3a843`.

```terraform
# FastACLRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACLRule by name
data "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}

output "fast_acl_rule_id" {
  value = data.xcsh_fast_acl_rule.example.id
}
```

## Next pages

- [Examples](data-sources--fast_acl_rule--examples.md)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
