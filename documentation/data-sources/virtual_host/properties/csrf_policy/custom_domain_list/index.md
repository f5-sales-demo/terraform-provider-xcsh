---
page_title: "csrf_policy.custom_domain_list"
subcategory: ""
description: "List of domain names used for Host header matching."
xcsh_docs: {"aliases": ["csrf policy custom domain list"], "body_bytes": 2494, "body_sha256": "sha256:96b7e86cca23aa08bf68708e733f88529543384d7e04de8578cc33486f09322b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:csrf_policy:custom_domain_list", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:csrf_policy", "path": "documentation/data-sources/virtual_host/properties/csrf_policy/custom_domain_list/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3030123002112223-0132013220102110-2111213112222322-0230132133330032-1201203113221330-0111313331112302-1121111033101120-0212232121222303", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["csrf_policy", "custom_domain_list"], "schema_version": 1, "sections": [{"aliases": ["csrf policy custom domain list domains"], "anchor": "schema-csrf_policy--custom_domain_list--domains", "description": "A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:csrf_policy:custom_domain_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["csrf_policy", "custom_domain_list", "domains"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/csrf_policy/custom_domain_list/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "List of domain names used for Host header matching.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# csrf_policy.custom_domain_list

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/csrf_policy/)
- csrf_policy.custom_domain_list

<a id="section"></a>

Type: `"single"`. Computed.

List of domain names used for Host header matching.

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

## Direct properties

<a id="schema-csrf_policy--custom_domain_list--domains"></a>

### domains property

Type: `["list", "string"]`. Computed.

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
