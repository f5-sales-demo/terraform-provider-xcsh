---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_upgrade_os."
xcsh_docs: {"aliases": ["action"], "body_bytes": 940, "body_sha256": "sha256:74d74a3b3f85878818842e024483cfcea0ad404dcc32eb6125c30d05f949b9da", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_upgrade_os:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:39e348141a455b8f8743fa7481580e9f9219671ca870734d07a968a5bcbd2acc", "source_path": "examples/actions/xcsh_site_upgrade_os/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_upgrade_os:example:action", "parent_id": "xcsh-docs:actions:site_upgrade_os:examples", "path": "documentation/actions/site_upgrade_os/examples/action/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_os", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "actions", "registry_anchor": "canonical-2020233322102210-1200223201102011-1232100213021313-1223022003212330-3022013310331221-2123133322200210-3201330022133300-0130322311002100", "registry_path": "docs/guides/actions--site_upgrade_os--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_os/examples/action/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Action for xcsh_site_upgrade_os.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
