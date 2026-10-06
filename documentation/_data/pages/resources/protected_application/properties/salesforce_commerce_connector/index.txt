---
page_title: "salesforce_commerce_connector"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["salesforce commerce connector"], "body_bytes": 948, "body_sha256": "sha256:66ba17acb6c5380249da0060b67b1b4799aebe852aff14cb9bef479685b34cd1", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:salesforce_commerce_connector", "parent_id": "xcsh-docs:resources:protected_application:reference", "path": "documentation/resources/protected_application/properties/salesforce_commerce_connector/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2332033220132333-1032222222002332-0311220332202110-2323303102320312-2200130201223220-1133120130131231-2221031302220301-1011010202321210", "registry_path": "docs/guides/resources--protected_application--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["salesforce_commerce_connector"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/salesforce_commerce_connector/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# salesforce_commerce_connector

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- salesforce_commerce_connector

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for salesforce commerce connector.

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
salesforce_commerce_connector = {}
```

This is an empty object or choice marker. It has no direct properties.
