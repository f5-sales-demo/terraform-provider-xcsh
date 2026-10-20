---
page_title: "pod_security_admission_specs.baseline"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["pod security admission specs baseline"], "body_bytes": 1086, "body_sha256": "sha256:604308a76179df3975f66d6b41bf35d23421c8b9dd52373f18f8d5b6248bdc10", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs", "path": "documentation/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/baseline/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2133120101221330-1013013031332212-0111021021330212-0122112012131332-2003123222210011-3302203210312121-0002131200013010-1121202031110332", "registry_path": "docs/guides/resources--k8s_pod_security_admission--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["pod_security_admission_specs", "baseline"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/baseline/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pod_security_admission_specs.baseline

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/)
- [pod_security_admission_specs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/)
- pod_security_admission_specs.baseline

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
baseline = {}
```

This is an empty object or choice marker. It has no direct properties.
