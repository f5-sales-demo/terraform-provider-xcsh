---
page_title: "disable_ip_reputation"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable ip reputation"], "body_bytes": 1370, "body_sha256": "sha256:457a0f4b2c97a104a7af9be67d33d0dfca73364d657e94324e82ab176ffa2dcf", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:disable_ip_reputation", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/disable_ip_reputation/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1031001202212121-2320213202323013-2310203020210132-3222031002011331-1231231310110332-2313123223113001-3022212102330213-3000321230000321", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_ip_reputation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/disable_ip_reputation/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_ip_reputation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- disable_ip_reputation

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/disable_ip_reputation/#section)
- [enable_ip_reputation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_ip_reputation/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ip_reputation = {}
```

This is an empty object or choice marker. It has no direct properties.
