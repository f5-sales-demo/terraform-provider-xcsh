---
page_title: "With description"
subcategory: ""
description: "With description for xcsh_user_identification."
xcsh_docs: {"aliases": ["with-description"], "body_bytes": 1106, "body_sha256": "sha256:0dbce3510d504ce11ebb173a01dae4aa49fee0bfb93a1d7a9c20f9df7df54a45", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:28e44c81af88136d36f13495510529b754168c2724c3ce7cec3eeb9297eaa127", "source_path": "examples/resources/xcsh_user_identification/with-description.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:with-description", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "documentation/resources/user_identification/examples/with-description/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3230320200122111-3032130232331310-2113130211013131-1033013122310231-3211313200033213-1100023221222131-3030102212211220-0013013203022123", "registry_path": "docs/guides/resources--user_identification--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-description"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/with-description/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "With description for xcsh_user_identification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["user_identificationCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With description

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-description.tf`; digest `sha256:28e44c81af88136d36f13495510529b754168c2724c3ce7cec3eeb9297eaa127`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
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
  name        = "example"
  namespace   = "system"
  description = "example-value"

  rules {
    client_ip = {}
  }
}
```
