---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_proxy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1020, "body_sha256": "sha256:fd390682c90a58da9752992be93f16e3485f7f69f60ba188efa7045ffce6fcc8", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8864854e827d63992024dc1c04e3c3a45871e2f9bca0bf3baae302eae31eb760", "source_path": "examples/data-sources/xcsh_dns_proxy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_proxy:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_proxy:examples", "path": "documentation/data-sources/dns_proxy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-3232113333222221-2310022121033100-1232222032132013-3021132202211022-2103103301221103-2221012022020230-0332310321000001-3322132022103120", "registry_path": "docs/guides/data-sources--dns_proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/examples/data-source/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Data source for xcsh_dns_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/examples/)
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
