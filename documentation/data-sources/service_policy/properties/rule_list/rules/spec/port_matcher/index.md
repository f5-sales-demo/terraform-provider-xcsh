---
page_title: "rule_list.rules.spec.port_matcher"
subcategory: "Security"
description: "A port matcher specifies a list of port ranges as match criteria. The match is considered successful if the input port falls within any of the port ranges. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["rule list rules spec port matcher", "succeeded", "success", "successful"], "body_bytes": 3728, "body_sha256": "sha256:d8759df5dabafdf8646dea6f75ff7528799af8876d8e99535974a63672563be1", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:port_matcher", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "path": "documentation/data-sources/service_policy/properties/rule_list/rules/spec/port_matcher/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0002112321213331-2301331003031102-2101130012331010-0010130231301212-0111003003223033-3131022232020221-3200101100331330-2033223020101132", "registry_path": "docs/guides/data-sources--service_policy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "port_matcher"], "schema_version": 1, "sections": [{"aliases": ["rule list rules spec port matcher invert matcher"], "anchor": "schema-rule_list--rules--spec--port_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:port_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "port_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["rule list rules spec port matcher ports"], "anchor": "schema-rule_list--rules--spec--port_matcher--ports", "description": "A list of strings, each of which is a single port value or a tuple of start and end port values separated by \"-\". The start and end values are considered to be part of the range.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:port_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "port_matcher", "ports"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/port_matcher/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "A port matcher specifies a list of port ranges as match criteria. The match is considered successful if the input port falls within any of the port ranges. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.port_matcher

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- rule_list.rules.spec.port_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

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

<a id="schema-rule_list--rules--spec--port_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert Port Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="schema-rule_list--rules--spec--port_matcher--ports"></a>

### ports property

Type: `["list", "string"]`. Computed.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
