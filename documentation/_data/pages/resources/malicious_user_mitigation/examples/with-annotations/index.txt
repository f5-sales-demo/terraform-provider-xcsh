---
page_title: "With annotations"
subcategory: ""
description: "With annotations for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": ["with-annotations"], "body_bytes": 1445, "body_sha256": "sha256:11f670f1513055e8b3208c12d916ae282451f7f7236e0e4f813b3db80e081050", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:10af11e1fcedf0c1a38131c01416f578defc093012b6a6669770dd539168e15b", "source_path": "examples/resources/xcsh_malicious_user_mitigation/with-annotations.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:malicious_user_mitigation:example:with-annotations", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:examples", "path": "documentation/resources/malicious_user_mitigation/examples/with-annotations/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0332011323203221-2212021231330001-3022103131310300-1030022012220312-0103300311201021-3223103001302132-0300002300321220-3012322030320300", "registry_path": "docs/guides/resources--malicious_user_mitigation--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-annotations"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/examples/with-annotations/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "With annotations for xcsh_malicious_user_mitigation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With annotations

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/examples/)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/with-annotations.tf`; digest `sha256:10af11e1fcedf0c1a38131c01416f578defc093012b6a6669770dd539168e15b`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "= 0.13.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_malicious_user_mitigation" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "example-value"
  namespace  = xcsh_namespace.test.name

  annotations = {
    example-key = "example-value"
  }
}
```
