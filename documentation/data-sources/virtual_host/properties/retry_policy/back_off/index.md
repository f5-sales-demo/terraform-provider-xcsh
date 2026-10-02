---
page_title: "retry_policy.back_off"
subcategory: ""
description: "Specifies parameters that control retry back off."
xcsh_docs: {"aliases": ["retry policy back off"], "body_bytes": 2544, "body_sha256": "sha256:0bee8893ef69020c7107c7e34059f1905ba3e5b5abcf37e72a62fc3e0b9c3d28", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:retry_policy:back_off", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:retry_policy", "path": "documentation/data-sources/virtual_host/properties/retry_policy/back_off/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3213001230013032-0020221021322333-1332330301110002-0020303010112211-3321113230333322-3203333212321332-0203333021032210-3112130033232321", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["retry_policy", "back_off"], "schema_version": 1, "sections": [{"aliases": ["base interval"], "anchor": "schema-retry_policy--back_off--base_interval", "description": "Specifies the base interval between retries in milliseconds.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:retry_policy:back_off", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["retry_policy", "back_off", "base_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["max interval"], "anchor": "schema-retry_policy--back_off--max_interval", "description": "Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The default is 10 times the base_interval.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:retry_policy:back_off", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["retry_policy", "back_off", "max_interval"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/retry_policy/back_off/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specifies parameters that control retry back off.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# retry_policy.back_off

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/retry_policy/)
- retry_policy.back_off

<a id="section"></a>

Type: `"single"`. Computed.

Specifies parameters that control retry back off.

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

<a id="schema-retry_policy--back_off--base_interval"></a>

### base_interval property

Type: `"number"`. Computed.

Specifies the base interval between retries in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="schema-retry_policy--back_off--max_interval"></a>

### max_interval property

Type: `"number"`. Computed.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Upstream description:

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The default is 10 times the base\_interval.

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

## Next pages

- [retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/retry_policy/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
