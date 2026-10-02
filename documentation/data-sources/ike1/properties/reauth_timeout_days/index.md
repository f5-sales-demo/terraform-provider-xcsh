---
page_title: "reauth_timeout_days"
subcategory: ""
description: "Set Duration in days."
xcsh_docs: {"aliases": ["duration", "operation timeout", "reauth timeout days"], "body_bytes": 1826, "body_sha256": "sha256:311efbbfedbb2e3ed3e2a3b2f4f0537a38e3908ec348de94625730f4e6fe11de", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike1:properties:reauth_timeout_days", "parent_id": "xcsh-docs:data-sources:ike1:reference", "path": "documentation/data-sources/ike1/properties/reauth_timeout_days/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2133111320003023-0122010212102113-2210003202133112-0212222121011111-1232101211020212-3120030011222322-1002232301311221-2221301000112022", "registry_path": "docs/guides/data-sources--ike1--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["reauth_timeout_days"], "schema_version": 1, "sections": [{"aliases": ["duration"], "anchor": "schema-reauth_timeout_days--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:data-sources:ike1:properties:reauth_timeout_days", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["reauth_timeout_days", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike1/properties/reauth_timeout_days/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Set Duration in days.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike1CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_timeout_days

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/)
- reauth_timeout_days

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

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

<a id="schema-reauth_timeout_days--duration"></a>

### duration property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/)
- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/)
