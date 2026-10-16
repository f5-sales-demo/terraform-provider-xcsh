---
page_title: "tls_parameters.common_params.tls_certificates.use_system_defaults"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls parameters common params tls certificates use system defaults"], "body_bytes": 1465, "body_sha256": "sha256:7b8144b7e964c79ec6869f242db2527e68727938c3acb19eb3044fc4815df706", "capabilities": ["load-balancing.tls", "networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:use_system_defaults", "parent_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params:tls_certificates", "path": "documentation/resources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-1100113010020100-0120101113323322-1322320201022021-3331332012322122-2002021222301231-1003000102013223-2323210121330021-0223021013300112", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "use_system_defaults"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.use_system_defaults

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/common_params/)
- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.
