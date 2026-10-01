---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_api_crawler."
xcsh_docs: {"aliases": [], "body_bytes": 1015, "body_sha256": "sha256:fa305ca7a2734e6a8d3f0a9ea1bb66c78c7159175ce3f8f7ecd0ff1154076c51", "canonical_id": "xcsh-docs:resources:api_crawler:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a73dd5a1c2c2dd7d614c3510b02a2ee9ecb351afd8ffe9c697a5142ace64fc30", "source_path": "examples/resources/xcsh_api_crawler/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_crawler:example:resource", "parent_id": "xcsh-docs:resources:api_crawler:examples", "path": "docs/guides/resources--api_crawler--example--resource.md", "provider_name": "api_crawler", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_api_crawler.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md)
- [Examples](resources--api_crawler--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_crawler/resource.tf`; digest `sha256:a73dd5a1c2c2dd7d614c3510b02a2ee9ecb351afd8ffe9c697a5142ace64fc30`.

```terraform
# APICrawler Resource Example
# Manages a API Crawler resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APICrawler configuration
resource "xcsh_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--api_crawler--examples.md)
- [xcsh_api_crawler](../resources/api_crawler.md)
