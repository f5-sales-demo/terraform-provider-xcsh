---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_upgrade_os."
xcsh_docs: {"aliases": [], "body_bytes": 1167, "body_sha256": "sha256:1e3202354208528735c7814b20460e665801566c33a0fc691be6048e0533545d", "child_ids": [], "collection_id": "xcsh-docs:actions:site_upgrade_os:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:39e348141a455b8f8743fa7481580e9f9219671ca870734d07a968a5bcbd2acc", "source_path": "examples/actions/xcsh_site_upgrade_os/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_upgrade_os:example:action", "parent_id": "xcsh-docs:actions:site_upgrade_os:examples", "path": "documentation/actions/site_upgrade_os/examples/action/index.md", "provider_name": "site_upgrade_os", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_os/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_site_upgrade_os.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_site_upgrade_os](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_os/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_os/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_site_upgrade_os/action.tf`; digest `sha256:39e348141a455b8f8743fa7481580e9f9219671ca870734d07a968a5bcbd2acc`.

```terraform
# SiteUpgradeOS Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_upgrade_os" "example" {
  config {
    site       = "example-value"
    os_version = "example-value"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_os/examples/)
- [xcsh_site_upgrade_os](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_os/)
