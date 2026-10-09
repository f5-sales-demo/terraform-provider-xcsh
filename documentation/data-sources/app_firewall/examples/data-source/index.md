---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_app_firewall."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1051, "body_sha256": "sha256:51bfc00416d6a333da0cdf31f7f6a57adea7523ff9feada5dd1b0fa4de39f8f7", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3f7d455a2134b0926217d55dc017e3580269822e7d265dbe83b586db534e3931", "source_path": "examples/data-sources/xcsh_app_firewall/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_firewall:example:data-source", "parent_id": "xcsh-docs:data-sources:app_firewall:examples", "path": "documentation/data-sources/app_firewall/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0122112313333132-2212103030220033-2313030100310020-2013011031101232-1333221300331121-3312123312233031-2111002030001000-0202231311001123", "registry_path": "docs/guides/data-sources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["app_firewallCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
