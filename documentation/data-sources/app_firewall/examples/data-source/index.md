---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1279, "body_sha256": "sha256:6c22cddfab39f95a749e951ed21a347aa5b0f3864a1a9a11f313bee8fe5b048c", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3f7d455a2134b0926217d55dc017e3580269822e7d265dbe83b586db534e3931", "source_path": "examples/data-sources/xcsh_app_firewall/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_firewall:example:data-source", "parent_id": "xcsh-docs:data-sources:app_firewall:examples", "path": "documentation/data-sources/app_firewall/examples/data-source/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_firewall/data-source.tf`; digest `sha256:3f7d455a2134b0926217d55dc017e3580269822e7d265dbe83b586db534e3931`.

```terraform
# AppFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppFirewall by name
data "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}

output "app_firewall_id" {
  value = data.xcsh_app_firewall.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/examples/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
