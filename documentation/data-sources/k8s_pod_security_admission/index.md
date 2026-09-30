---
page_title: "xcsh_k8s_pod_security_admission"
subcategory: ""
description: "xcsh_k8s_pod_security_admission for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": [], "body_bytes": 1378, "body_sha256": "sha256:8731db741b3e8c2cc41af08e618e6def0192b3bf0f99fef9c1b62e8efa439f0d", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_admission:reference", "xcsh-docs:data-sources:k8s_pod_security_admission:examples"], "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_admission:fundamentals", "parent_id": null, "path": "documentation/data-sources/k8s_pod_security_admission/index.md", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_k8s_pod_security_admission for xcsh_k8s_pod_security_admission.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_k8s_pod_security_admission

Breadcrumbs:

- xcsh_k8s_pod_security_admission

Manages k8s\_pod\_security\_admission will create the object in the storage backend in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityAdmission Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SPodSecurityAdmission by name
data "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}

output "k8s_pod_security_admission_id" {
  value = data.xcsh_k8s_pod_security_admission.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/examples/)
