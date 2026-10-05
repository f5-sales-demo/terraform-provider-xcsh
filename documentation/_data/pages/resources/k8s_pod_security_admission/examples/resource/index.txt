---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1431, "body_sha256": "sha256:1a413b14f1fe11c606ca92c08f3d060b2ee68d68f24eb5fb545d1f02904b4c82", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:55b7bf85163231f6d1cb9fd77f147ddf87ba9bdf58b1aa2d6af5bb2885858dc2", "source_path": "examples/resources/xcsh_k8s_pod_security_admission/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:k8s_pod_security_admission:example:resource", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:examples", "path": "documentation/resources/k8s_pod_security_admission/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1202203123123302-0101221202003123-2133202023323322-3231332220120210-1031221233313023-1122123230121120-3313032301323000-1323230331320202", "registry_path": "docs/guides/resources--k8s_pod_security_admission--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_k8s_pod_security_admission.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/examples/)
- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/)
