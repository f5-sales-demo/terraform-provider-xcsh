---
page_title: "Cookie"
subcategory: ""
description: "Cookie for xcsh_user_identification."
xcsh_docs: {"aliases": ["cookie"], "body_bytes": 1043, "body_sha256": "sha256:2d694155b75e0ab54af8d97c719d88ea66695cc20898176182ef3094218429f5", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:5603d5f0bc4ec0397625fee3d3eb29ea6bfe17b42e387a451f62ff082893af26", "source_path": "examples/resources/xcsh_user_identification/cookie.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:cookie", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "documentation/resources/user_identification/examples/cookie/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0213113321013213-3210122030022033-0311331313313200-1000233122011303-2013121221231230-0233032300231222-3222020122113321-2330231013231033", "registry_path": "docs/guides/resources--user_identification--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["cookie"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/cookie/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Cookie for xcsh_user_identification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["user_identificationCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Cookie

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- Cookie

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/cookie.tf`; digest `sha256:5603d5f0bc4ec0397625fee3d3eb29ea6bfe17b42e387a451f62ff082893af26`.

```terraform
# Cookie — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  rules {
    cookie_name = "session_id"
  }
}
```
