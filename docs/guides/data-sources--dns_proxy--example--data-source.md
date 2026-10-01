---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1033, "body_sha256": "sha256:3b3c9ffd579eb9e19972e72ed8bb69ceb07df547ad377fdd9c28ee9e75aac87c", "canonical_id": "xcsh-docs:data-sources:dns_proxy:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8864854e827d63992024dc1c04e3c3a45871e2f9bca0bf3baae302eae31eb760", "source_path": "examples/data-sources/xcsh_dns_proxy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_proxy:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_proxy:examples", "path": "docs/guides/data-sources--dns_proxy--example--data-source.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Examples](data-sources--dns_proxy--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_proxy/data-source.tf`; digest `sha256:8864854e827d63992024dc1c04e3c3a45871e2f9bca0bf3baae302eae31eb760`.

```terraform
# DNSProxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSProxy by name
data "xcsh_dns_proxy" "example" {
  name      = "example-dns-proxy"
  namespace = "system"
}

output "dns_proxy_id" {
  value = data.xcsh_dns_proxy.example.id
}
```

## Next pages

- [Examples](data-sources--dns_proxy--examples.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
