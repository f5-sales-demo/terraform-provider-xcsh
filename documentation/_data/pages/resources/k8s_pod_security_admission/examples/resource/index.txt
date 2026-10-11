---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1167, "body_sha256": "sha256:faf97bec9a0364efe346e4329615e8562b0cbc340bf57a2a39446a82cbf10b19", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:55b7bf85163231f6d1cb9fd77f147ddf87ba9bdf58b1aa2d6af5bb2885858dc2", "source_path": "examples/resources/xcsh_k8s_pod_security_admission/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:k8s_pod_security_admission:example:resource", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:examples", "path": "documentation/resources/k8s_pod_security_admission/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1202203123123302-0101221202003123-2133202023323322-3231332220120210-1031221233313023-1122123230121120-3313032301323000-1323230331320202", "registry_path": "docs/guides/resources--k8s_pod_security_admission--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_k8s_pod_security_admission.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_pod_security_admission/resource.tf`; digest `sha256:55b7bf85163231f6d1cb9fd77f147ddf87ba9bdf58b1aa2d6af5bb2885858dc2`.

```terraform
# K8SPodSecurityAdmission Resource Example
# Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityAdmission configuration
resource "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}
```
