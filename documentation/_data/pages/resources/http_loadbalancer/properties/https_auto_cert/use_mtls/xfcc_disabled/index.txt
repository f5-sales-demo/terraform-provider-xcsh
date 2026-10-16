---
page_title: "https_auto_cert.use_mtls.xfcc_disabled"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https auto cert use mtls xfcc disabled"], "body_bytes": 1182, "body_sha256": "sha256:8530540cb85fd146b294040d33e88b263dae37f652c9efd08527309109d8c5b6", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:use_mtls:xfcc_disabled", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:use_mtls", "path": "documentation/resources/http_loadbalancer/properties/https_auto_cert/use_mtls/xfcc_disabled/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3203223313321313-1310203230020321-0010112031231321-0112210212013021-1121232000031112-1321201310001132-2202200001103323-0101210330332032", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_auto_cert", "use_mtls", "xfcc_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https_auto_cert/use_mtls/xfcc_disabled/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.use_mtls.xfcc_disabled

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https_auto_cert/)
- [https_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https_auto_cert/use_mtls/)
- https_auto_cert.use_mtls.xfcc_disabled

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
