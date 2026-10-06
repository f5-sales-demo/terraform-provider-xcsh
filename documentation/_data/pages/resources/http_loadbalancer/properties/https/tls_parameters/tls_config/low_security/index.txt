---
page_title: "https.tls_parameters.tls_config.low_security"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https tls parameters tls config low security"], "body_bytes": 1330, "body_sha256": "sha256:6eabafa2939a58f42fd3ab3abd97025bf481eac7ad6e974d6323bbde58307991", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config", "path": "documentation/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/low_security/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3212011310330030-2002101322323302-0030022003030222-1222111202312132-0130202122103133-3211111312131222-3033100131100311-3013030233002023", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "tls_parameters", "tls_config", "low_security"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/low_security/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.tls_config.low_security

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/)
- [https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/)
- [https.tls_parameters.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/)
- https.tls_parameters.tls_config.low_security

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.
