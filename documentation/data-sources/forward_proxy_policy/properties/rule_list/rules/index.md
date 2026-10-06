---
page_title: "rule_list.rules"
subcategory: "Security"
description: "List of custom rules."
xcsh_docs: {"aliases": ["rule list rules"], "body_bytes": 6809, "body_sha256": "sha256:0c14ccd3e7d2428dc1808f4cb24b207920e530ec7c3fa1f9da66f4f744a9ce77", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:all_destinations", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:all_sources", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_asn_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_asn_set", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_ip_prefix_set", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_label_selector", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_prefix_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:http_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:ip_prefix_set", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:label_selector", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:metadata", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:no_http_connect_port", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:port_matcher", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:prefix_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:url_category_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list", "path": "documentation/data-sources/forward_proxy_policy/properties/rule_list/rules/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "sections": [{"aliases": ["rule list rules action"], "anchor": "schema-rule_list--rules--action", "description": "The rule action determines the disposition of the input request API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds forward. If it matches a rule with a DENY action, the processing of the request is terminated and an appropriate message/code returned to the originator. If it", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "action"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule list rules all destinations"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:all_destinations", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_destinations"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules all sources"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:all_sources", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_sources"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules dst asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_asn_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "dst_asn_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules dst asn set"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_asn_set", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "dst_asn_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules dst ip prefix set"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "dst_ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules dst label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_label_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "dst_label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules dst prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "dst_prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules http list"], "anchor": "section", "description": "URLListType.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:http_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "http_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules ip prefix set"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:label_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules no http connect port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:no_http_connect_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "no_http_connect_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules port matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A port matcher specifies a list of port ranges as match criteria. The match is considered successful if the input port falls within any of the port ranges. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:port_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "port_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules tls list"], "anchor": "section", "description": "DomainListType.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "tls_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules url category list"], "anchor": "section", "description": "List of URL categories.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:url_category_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "url_category_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of custom rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/)
- rule_list.rules

<a id="section"></a>

Type: `"list"`. Computed.

Custom Rule List. List of custom rules.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

<a id="schema-rule_list--rules--action"></a>

### action property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW|NEXT\_POLICY\] The rule action determines the disposition of the input request
API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds
forward. If it matches a rule with a DENY action, the processing of the request is terminated and an
appropriate message/code returned to.. Possible values are \`DENY\`, \`ALLOW\`, \`NEXT\_POLICY\`.
Defaults to \`DENY\`.

Additional upstream details:

The rule action determines the disposition of the input request API. If it matches a rule with a
DENY action, the processing of the request is terminated and an appropriate message/code returned to
the originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current
policy set terminates and evaluation of the next policy set in the chain begins.

&#8203;- DENY: DENY

Deny the request. &#8203;- ALLOW: ALLOW

Allow the request to proceed. &#8203;- NEXT\_POLICY\_SET: NEXT\_POLICY\_SET

Terminate evaluation of the current policy set and begin evaluating the next policy set in the
chain. Note that the evaluation of any remaining policies in the current policy set is skipped.
&#8203;- NEXT\_POLICY: NEXT\_POLICY

Terminate evaluation of the current policy and begin evaluating the next policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
LAST\_POLICY: LAST\_POLICY

Terminate evaluation of the current policy and begin evaluating the last policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
GOTO\_POLICY: GOTO\_POLICY

Terminate evaluation of the current policy and begin evaluating a specific policy in the policy set.
The policy is specified using the goto\_policy field in the rule and must be after the current
policy in the policy set.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW",
    "NEXT_POLICY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/all_destinations/): complete subsection reference.

- [all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/all_sources/): complete subsection reference.

- [dst_asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/dst_asn_list/): complete subsection reference.

- [dst_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/dst_asn_set/): complete subsection reference.

- [dst_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/dst_ip_prefix_set/): complete subsection reference.

- [dst_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/dst_label_selector/): complete subsection reference.

- [dst_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/dst_prefix_list/): complete subsection reference.

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/http_list/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/ip_prefix_set/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/metadata/): complete subsection reference.

- [no_http_connect_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/no_http_connect_port/): complete subsection reference.

- [port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/port_matcher/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/prefix_list/): complete subsection reference.

- [tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/): complete subsection reference.

- [url_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/url_category_list/): complete subsection reference.
