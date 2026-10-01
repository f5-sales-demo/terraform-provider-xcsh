---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": [], "body_bytes": 1225, "body_sha256": "sha256:7c6dc49c391f7502a4413ec2592a103d8634527affbc16631902470bb3ea8734", "canonical_id": "xcsh-docs:resources:k8s_pod_security_admission:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:55b7bf85163231f6d1cb9fd77f147ddf87ba9bdf58b1aa2d6af5bb2885858dc2", "source_path": "examples/resources/xcsh_k8s_pod_security_admission/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:k8s_pod_security_admission:example:resource", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:examples", "path": "docs/guides/resources--k8s_pod_security_admission--example--resource.md", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_k8s_pod_security_admission.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md)
- [Examples](resources--k8s_pod_security_admission--examples.md)
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

- [Examples](resources--k8s_pod_security_admission--examples.md)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md)
