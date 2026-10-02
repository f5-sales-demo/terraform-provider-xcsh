---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_app_firewall."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1279, "body_sha256": "sha256:6c22cddfab39f95a749e951ed21a347aa5b0f3864a1a9a11f313bee8fe5b048c", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3f7d455a2134b0926217d55dc017e3580269822e7d265dbe83b586db534e3931", "source_path": "examples/data-sources/xcsh_app_firewall/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_firewall:example:data-source", "parent_id": "xcsh-docs:data-sources:app_firewall:examples", "path": "documentation/data-sources/app_firewall/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0122112313333132-2212103030220033-2313030100310020-2013011031101232-1333221300331121-3312123312233031-2111002030001000-0202231311001123", "registry_path": "docs/guides/data-sources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
