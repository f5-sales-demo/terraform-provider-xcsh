---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_upgrade_os."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1167, "body_sha256": "sha256:1e3202354208528735c7814b20460e665801566c33a0fc691be6048e0533545d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_upgrade_os:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:39e348141a455b8f8743fa7481580e9f9219671ca870734d07a968a5bcbd2acc", "source_path": "examples/actions/xcsh_site_upgrade_os/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_upgrade_os:example:action", "parent_id": "xcsh-docs:actions:site_upgrade_os:examples", "path": "documentation/actions/site_upgrade_os/examples/action/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_os", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "registry_anchor": "canonical-2020233322102210-1200223201102011-1232100213021313-1223022003212330-3022013310331221-2123133322200210-3201330022133300-0130322311002100", "registry_path": "docs/guides/actions--site_upgrade_os--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_os/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_site_upgrade_os.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
