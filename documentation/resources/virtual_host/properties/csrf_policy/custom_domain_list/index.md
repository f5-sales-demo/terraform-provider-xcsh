---
page_title: "csrf_policy.custom_domain_list"
subcategory: ""
description: "List of domain names used for Host header matching."
xcsh_docs: {"aliases": ["csrf policy custom domain list"], "body_bytes": 2953, "body_sha256": "sha256:9a0a0671361edcb9e1431df04d22cdc54eece625698bf31a6ba17cd1ff531f7b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "parent_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy", "path": "documentation/resources/virtual_host/properties/csrf_policy/custom_domain_list/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0210302321013100-1300220003111020-1330033301211103-3021331002113120-2311332011003021-2100020010022101-3320121011300333-3030122322323001", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [{"anchor": "schema-csrf_policy--custom_domain_list--domains", "enforcement": "provider-schema", "group": "csrf_policy.custom_domain_list:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["csrf_policy", "custom_domain_list"], "schema_version": 1, "sections": [{"aliases": ["csrf policy custom domain list domains"], "anchor": "schema-csrf_policy--custom_domain_list--domains", "description": "A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form.", "document_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["csrf_policy", "custom_domain_list", "domains"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/csrf_policy/custom_domain_list/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of domain names used for Host header matching.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# csrf_policy.custom_domain_list

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/)
- csrf_policy.custom_domain_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
```

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
custom_domain_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-csrf_policy--custom_domain_list--domains"></a>

### domains property

Type: `["list", "string"]`. Optional.

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
