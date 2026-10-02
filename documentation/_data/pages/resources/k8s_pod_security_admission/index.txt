---
page_title: "xcsh_k8s_pod_security_admission"
subcategory: ""
description: "Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "k8s pod security admission", "origin servers", "upstream servers"], "body_bytes": 1720, "body_sha256": "sha256:89cf12b428a335069c5b28cd5227e4d604bb3ee63817717ae11ccbaf8d1e94b8", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_pod_security_admission:reference", "xcsh-docs:resources:k8s_pod_security_admission:examples", "xcsh-docs:resources:k8s_pod_security_admission:import", "xcsh-docs:resources:k8s_pod_security_admission:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_admission:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/k8s_pod_security_admission/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023", "registry_path": "docs/resources/k8s_pod_security_admission.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/lifecycle/timeouts/)
