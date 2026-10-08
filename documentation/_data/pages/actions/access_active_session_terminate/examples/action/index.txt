---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_access_active_session_terminate."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1033, "body_sha256": "sha256:b736ae6b337c5688b3d5c2a674e6ee98d679cf2440327be10aed897b2e9aaaca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:access_active_session_terminate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1bbb807c7b4d5333250c2b20cf18f88631715519e818829cb7766f98ee751aa7", "source_path": "examples/actions/xcsh_access_active_session_terminate/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:access_active_session_terminate:example:action", "parent_id": "xcsh-docs:actions:access_active_session_terminate:examples", "path": "documentation/actions/access_active_session_terminate/examples/action/index.md", "product": "distributed-cloud", "provider_name": "access_active_session_terminate", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "actions", "registry_anchor": "canonical-2110223031311333-3310300201232301-3301310220122231-2200000222230101-2013031200123120-1223030202323330-1122220301101011-2132332032023110", "registry_path": "docs/guides/actions--access_active_session_terminate--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_session_terminate/examples/action/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Action for xcsh_access_active_session_terminate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_access_active_session_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_access_active_session_terminate/action.tf`; digest `sha256:1bbb807c7b4d5333250c2b20cf18f88631715519e818829cb7766f98ee751aa7`.

```terraform
# AccessActiveSessionTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_session_terminate" "example" {
  config {
    id        = "example-value"
    namespace = "example-value"
  }
}
```
