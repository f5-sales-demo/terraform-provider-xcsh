---
page_title: "user_defined_api_discovery_policy.exclusive.ignore"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["user defined api discovery policy exclusive ignore"], "body_bytes": 1253, "body_sha256": "sha256:b88ad372b25fc99a920cbef6af2c74af645461fb82d4445f0ab32f97977f054c", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore", "parent_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "path": "documentation/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/ignore/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3031133300212101-0001123213300321-1333301200111021-3103100331030121-2321100311233312-1020100101032310-3110232112303200-2222232323111002", "registry_path": "docs/guides/resources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "exclusive", "ignore"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/ignore/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.exclusive.ignore

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/)
- [user_defined_api_discovery_policy.exclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/)
- user_defined_api_discovery_policy.exclusive.ignore

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
ignore = {}
```

This is an empty object or choice marker. It has no direct properties.
