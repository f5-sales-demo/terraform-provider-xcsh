---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_access_active_sessions_terminate."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1007, "body_sha256": "sha256:41807fdd151d8030b34357237ed89e44a791d31522652b45926fe0e067e0c3f6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:access_active_sessions_terminate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e3e7c43f13f1ace8b437c1c8885aa667a9add78ae921c28c7a137377165d7c29", "source_path": "examples/actions/xcsh_access_active_sessions_terminate/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:access_active_sessions_terminate:example:action", "parent_id": "xcsh-docs:actions:access_active_sessions_terminate:examples", "path": "documentation/actions/access_active_sessions_terminate/examples/action/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions_terminate", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "actions", "registry_anchor": "canonical-0323022220300023-0020313113010222-2300010131030002-0211131012302230-2213331102112032-2023220223020302-0230331230022310-0020130211311311", "registry_path": "docs/guides/actions--access_active_sessions_terminate--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_sessions_terminate/examples/action/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Action for xcsh_access_active_sessions_terminate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": [], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_access_active_sessions_terminate/action.tf`; digest `sha256:e3e7c43f13f1ace8b437c1c8885aa667a9add78ae921c28c7a137377165d7c29`.

```terraform
# AccessActiveSessionsTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_sessions_terminate" "example" {
  config {
    namespace = "example-value"
  }
}
```
