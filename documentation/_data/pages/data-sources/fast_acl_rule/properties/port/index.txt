---
page_title: "port"
subcategory: ""
description: "L4 port numbers to match."
xcsh_docs: {"aliases": ["port"], "body_bytes": 2965, "body_sha256": "sha256:b7a3e18fa2f43dc42aa77e1ec92bc6be72513b52b7a4dd9ad31e05be59906208", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl_rule:properties:port:all", "xcsh-docs:data-sources:fast_acl_rule:properties:port:dns"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl_rule:properties:port", "parent_id": "xcsh-docs:data-sources:fast_acl_rule:reference", "path": "documentation/data-sources/fast_acl_rule/properties/port/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213", "registry_path": "docs/guides/data-sources--fast_acl_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["port"], "schema_version": 1, "sections": [{"aliases": ["all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:fast_acl_rule:properties:port:all", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port", "all"], "syntax": "attribute", "type": "object"}, {"aliases": ["dns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:fast_acl_rule:properties:port:dns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port", "dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["user defined"], "anchor": "schema-port--user_defined", "description": "Exclusive with Matches the user defined port.", "document_id": "xcsh-docs:data-sources:fast_acl_rule:properties:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port", "user_defined"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/properties/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "L4 port numbers to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# port

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/)
- port

<a id="section"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

## Direct properties

- [all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/port/all/): complete subsection reference.

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/port/dns/): complete subsection reference.

<a id="schema-port--user_defined"></a>

### user_defined property

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [port.all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/port/all/)
- [port.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/port/dns/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/)
- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/)
