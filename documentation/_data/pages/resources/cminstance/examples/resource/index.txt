---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cminstance."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1089, "body_sha256": "sha256:943776066cace514fc8ff78f2fea629c147b87870f06999388cec77357adc251", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cminstance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8160520d3332724273c3557478a3147a7b31a2fd2927166c5bb6f2b977df1b95", "source_path": "examples/resources/xcsh_cminstance/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cminstance:example:resource", "parent_id": "xcsh-docs:resources:cminstance:examples", "path": "documentation/resources/cminstance/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0030121000332330-0011313010131230-2210212231210232-2200230213002310-0011030312023330-1113011221032011-2000100201032103-0020101333302020", "registry_path": "docs/guides/resources--cminstance--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cminstance/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_cminstance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cminstanceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cminstance/resource.tf`; digest `sha256:8160520d3332724273c3557478a3147a7b31a2fd2927166c5bb6f2b977df1b95`.

```terraform
# Cminstance Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cminstance configuration
resource "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"

  port     = 1
  username = "example-value"
}
```
