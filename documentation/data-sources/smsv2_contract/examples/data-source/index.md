---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_smsv2_contract."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1815, "body_sha256": "sha256:da1ceb1fa46c463b987920d407a10d2183401ac7f709e055b04b442a8a6ef45d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_contract:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e80a034ec67b555714059852f638089936add60d6d1f3b5f77d3c2d3ae050f0f", "source_path": "examples/data-sources/xcsh_smsv2_contract/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:smsv2_contract:example:data-source", "parent_id": "xcsh-docs:data-sources:smsv2_contract:examples", "path": "documentation/data-sources/smsv2_contract/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "smsv2_contract", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2110122221312022-2232102303200000-0031102010330212-3302100323023133-1232123332011220-2332032003030211-2001011312231033-1333022133021312", "registry_path": "docs/guides/data-sources--smsv2_contract--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_contract/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_smsv2_contract.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_smsv2_contract](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_smsv2_contract/data-source.tf`; digest `sha256:e80a034ec67b555714059852f638089936add60d6d1f3b5f77d3c2d3ae050f0f`.

```terraform
# Read the immutable clean-break SMSv2 contract compiled into the provider.
# Required capabilities are checked during planning before any F5 API request.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_smsv2_contract" "current" {
  required_capabilities = ["runtime_status"]
}

output "smsv2_contract" {
  value = {
    id                               = data.xcsh_smsv2_contract.current.contract_id
    version                          = data.xcsh_smsv2_contract.current.contract_version
    api_release                      = data.xcsh_smsv2_contract.current.api_release_tag
    telemetry_schema_id              = data.xcsh_smsv2_contract.current.telemetry_schema_id
    capabilities                     = data.xcsh_smsv2_contract.current.capabilities
    f5xc_authorities                 = data.xcsh_smsv2_contract.current.f5xc_authorities
    aws_authorities                  = data.xcsh_smsv2_contract.current.aws_authorities
    azure_route_server_ebgp_multihop = data.xcsh_smsv2_contract.current.azure_route_server_ebgp_multihop
  }
}
```
