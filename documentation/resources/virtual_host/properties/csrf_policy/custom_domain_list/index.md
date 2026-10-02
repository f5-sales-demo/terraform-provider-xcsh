---
page_title: "csrf_policy.custom_domain_list"
subcategory: ""
description: "List of domain names used for Host header matching."
xcsh_docs: {"aliases": ["csrf policy custom domain list"], "body_bytes": 3316, "body_sha256": "sha256:3bee6c5fc238124562750c9f3422e6df0ade323c115219e5a1a755c20208538a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "parent_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy", "path": "documentation/resources/virtual_host/properties/csrf_policy/custom_domain_list/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0210302321013100-1300220003111020-1330033301211103-3021331002113120-2311332011003021-2100020010022101-3320121011300333-3030122322323001", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [{"anchor": "schema-csrf_policy--custom_domain_list--domains", "enforcement": "provider-schema", "group": "csrf_policy.custom_domain_list:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["csrf_policy", "custom_domain_list"], "schema_version": 1, "sections": [{"aliases": ["domains"], "anchor": "schema-csrf_policy--custom_domain_list--domains", "description": "A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form.", "document_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["csrf_policy", "custom_domain_list", "domains"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/csrf_policy/custom_domain_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of domain names used for Host header matching.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
