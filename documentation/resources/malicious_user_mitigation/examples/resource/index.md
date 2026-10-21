---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1182, "body_sha256": "sha256:d6800ad875cd45f9e08c97f446b763085423838dd90b253602d4331236d9eea0", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cb208f65f7f8d9724e94da5a2c35bd650a88bc8059b0aeced4b7db40e67c44ff", "source_path": "examples/resources/xcsh_malicious_user_mitigation/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:malicious_user_mitigation:example:resource", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:examples", "path": "documentation/resources/malicious_user_mitigation/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3301001202201031-1213330213031310-1333230003311100-2013300321320202-0113311010311321-0233322120202133-3020120221033023-3023220011010101", "registry_path": "docs/guides/resources--malicious_user_mitigation--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/examples/resource/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Resource for xcsh_malicious_user_mitigation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/resource.tf`; digest `sha256:cb208f65f7f8d9724e94da5a2c35bd650a88bc8059b0aeced4b7db40e67c44ff`.

```terraform
# MaliciousUserMitigation Resource Example
# Manages malicious_user_mitigation creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MaliciousUserMitigation configuration
resource "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}
```
