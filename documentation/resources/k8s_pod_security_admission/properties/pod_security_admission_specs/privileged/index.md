---
page_title: "pod_security_admission_specs.privileged"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["pod security admission specs privileged"], "body_bytes": 1092, "body_sha256": "sha256:d033bc5127e2262c46af26cf22e7d4a5c99f61cf2fc60ed142c737c4fcdfc6df", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs", "path": "documentation/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/privileged/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2221201033113120-3102221030010231-3013121211311033-3310133213032232-2333013322013120-1100103021023312-2222300310113132-1000322132313333", "registry_path": "docs/guides/resources--k8s_pod_security_admission--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["pod_security_admission_specs", "privileged"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/privileged/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pod_security_admission_specs.privileged

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/)
- [pod_security_admission_specs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/)
- pod_security_admission_specs.privileged

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
privileged = {}
```

This is an empty object or choice marker. It has no direct properties.
