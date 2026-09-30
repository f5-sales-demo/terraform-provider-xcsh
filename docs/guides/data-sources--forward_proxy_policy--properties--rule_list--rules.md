---
page_title: "rule_list.rules"
subcategory: "Security"
description: "rule_list.rules for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 7966, "body_sha256": "sha256:305fd975c51b2832dfcd9070a45da9318e21a630c8ab3230d81275bbb86126d9", "canonical_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:all_destinations", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:all_sources", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_asn_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_asn_set", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_ip_prefix_set", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_label_selector", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:dst_prefix_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:http_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:ip_prefix_set", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:label_selector", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:metadata", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:no_http_connect_port", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:port_matcher", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:prefix_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:url_category_list"], "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list", "path": "docs/guides/data-sources--forward_proxy_policy--properties--rule_list--rules.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
- [Property reference](data-sources--forward_proxy_policy--reference.md)
- [rule_list](data-sources--forward_proxy_policy--properties--rule_list.md)
- rule_list.rules

<a id="section"></a>

Type: `"list"`. Computed.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

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

Upstream description:

The rule action determines the disposition of the input request API. If a policy matches a rule with
an ALLOW action, the processing of the request proceeds forward. If it matches a rule with a DENY
action, the processing of the request is terminated and an appropriate message/code returned to the
originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current policy
set terminates and evaluation of the next policy set in the chain begins.

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

- [all_destinations](data-sources--forward_proxy_policy--properties--rule_list--rules--all_destinations.md): complete subsection reference.

- [all_sources](data-sources--forward_proxy_policy--properties--rule_list--rules--all_sources.md): complete subsection reference.

- [dst_asn_list](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_asn_list.md): complete subsection reference.

- [dst_asn_set](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_asn_set.md): complete subsection reference.

- [dst_ip_prefix_set](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_ip_prefix_set.md): complete subsection reference.

- [dst_label_selector](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_label_selector.md): complete subsection reference.

- [dst_prefix_list](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_prefix_list.md): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--properties--rule_list--rules--http_list.md): complete subsection reference.

- [ip_prefix_set](data-sources--forward_proxy_policy--properties--rule_list--rules--ip_prefix_set.md): complete subsection reference.

- [label_selector](data-sources--forward_proxy_policy--properties--rule_list--rules--label_selector.md): complete subsection reference.

- [metadata](data-sources--forward_proxy_policy--properties--rule_list--rules--metadata.md): complete subsection reference.

- [no_http_connect_port](data-sources--forward_proxy_policy--properties--rule_list--rules--no_http_connect_port.md): complete subsection reference.

- [port_matcher](data-sources--forward_proxy_policy--properties--rule_list--rules--port_matcher.md): complete subsection reference.

- [prefix_list](data-sources--forward_proxy_policy--properties--rule_list--rules--prefix_list.md): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--properties--rule_list--rules--tls_list.md): complete subsection reference.

- [url_category_list](data-sources--forward_proxy_policy--properties--rule_list--rules--url_category_list.md): complete subsection reference.

## Next pages

- [rule_list.rules.all_destinations](data-sources--forward_proxy_policy--properties--rule_list--rules--all_destinations.md)
- [rule_list.rules.all_sources](data-sources--forward_proxy_policy--properties--rule_list--rules--all_sources.md)
- [rule_list.rules.dst_asn_list](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_asn_list.md)
- [rule_list.rules.dst_asn_set](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_asn_set.md)
- [rule_list.rules.dst_ip_prefix_set](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_ip_prefix_set.md)
- [rule_list.rules.dst_label_selector](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_label_selector.md)
- [rule_list.rules.dst_prefix_list](data-sources--forward_proxy_policy--properties--rule_list--rules--dst_prefix_list.md)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--properties--rule_list--rules--http_list.md)
- [rule_list.rules.ip_prefix_set](data-sources--forward_proxy_policy--properties--rule_list--rules--ip_prefix_set.md)
- [rule_list.rules.label_selector](data-sources--forward_proxy_policy--properties--rule_list--rules--label_selector.md)
- [rule_list.rules.metadata](data-sources--forward_proxy_policy--properties--rule_list--rules--metadata.md)
- [rule_list.rules.no_http_connect_port](data-sources--forward_proxy_policy--properties--rule_list--rules--no_http_connect_port.md)
- [rule_list.rules.port_matcher](data-sources--forward_proxy_policy--properties--rule_list--rules--port_matcher.md)
- [rule_list.rules.prefix_list](data-sources--forward_proxy_policy--properties--rule_list--rules--prefix_list.md)
- [rule_list.rules.tls_list](data-sources--forward_proxy_policy--properties--rule_list--rules--tls_list.md)
- [rule_list.rules.url_category_list](data-sources--forward_proxy_policy--properties--rule_list--rules--url_category_list.md)
- [rule_list](data-sources--forward_proxy_policy--properties--rule_list.md)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
