---
page_title: "re_acl.fast_acl_rules"
subcategory: ""
description: "Fast ACL rules to match."
xcsh_docs: {"aliases": ["re acl fast acl rules"], "body_bytes": 3292, "body_sha256": "sha256:844d561bd7281fc6a4525774dc961da886186bc91a5e140313f064f80c3c6e9a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:action", "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:metadata", "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:port", "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:prefix"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl", "path": "documentation/data-sources/fast_acl/properties/re_acl/fast_acl_rules/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "fast_acl_rules"], "schema_version": 1, "sections": [{"aliases": ["re acl fast acl rules action"], "anchor": "section", "description": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "action"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl fast acl rules ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl fast acl rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl fast acl rules port"], "anchor": "section", "description": "L4 port numbers to match.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "port"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl fast acl rules prefix"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:prefix", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "prefix"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/fast_acl_rules/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Fast ACL rules to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/)
- re_acl.fast_acl_rules

<a id="section"></a>

Type: `"list"`. Computed.

Rules. Fast ACL rules to match. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

Fast ACL rules to match.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/action/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/ip_prefix_set/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/metadata/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/port/): complete subsection reference.

- [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/prefix/): complete subsection reference.

## Next pages

- [re_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/action/)
- [re_acl.fast_acl_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/ip_prefix_set/)
- [re_acl.fast_acl_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/metadata/)
- [re_acl.fast_acl_rules.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/port/)
- [re_acl.fast_acl_rules.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/prefix/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
