---
page_title: "rule_list.rules"
subcategory: "Security"
description: "rule_list.rules for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 13810, "body_sha256": "sha256:fc3008cd5b32376c8374073abdf600b63449cd18a525baf89ed53926889c9c08", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:all_destinations", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:all_sources", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:dst_asn_list", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:dst_asn_set", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:dst_ip_prefix_set", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:dst_label_selector", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:dst_prefix_list", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:ip_prefix_set", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:label_selector", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:metadata", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:no_http_connect_port", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:port_matcher", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:prefix_list", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:tls_list", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:url_category_list"], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules", "parent_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list", "path": "documentation/resources/forward_proxy_policy/properties/rule_list/rules/index.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/)
- rule_list.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_destinations",
    "dst_asn_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_asn_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "http_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "tls_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("all_sources",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_asn_set"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("no_http_connect_port",
    "port_matcher"),
  validators.ConflictingListObjectAttributes("tls_list",
    "url_category_list")}
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--action"></a>

### action property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW",
    "NEXT_POLICY"),
}
```

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

- [all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/all_destinations/): complete subsection reference.

- [all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/all_sources/): complete subsection reference.

- [dst_asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_list/): complete subsection reference.

- [dst_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_set/): complete subsection reference.

- [dst_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_ip_prefix_set/): complete subsection reference.

- [dst_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_label_selector/): complete subsection reference.

- [dst_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_prefix_list/): complete subsection reference.

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/ip_prefix_set/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/metadata/): complete subsection reference.

- [no_http_connect_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/no_http_connect_port/): complete subsection reference.

- [port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/port_matcher/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/prefix_list/): complete subsection reference.

- [tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/tls_list/): complete subsection reference.

- [url_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/url_category_list/): complete subsection reference.

## Next pages

- [rule_list.rules.all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/all_destinations/)
- [rule_list.rules.all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/all_sources/)
- [rule_list.rules.dst_asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_list/)
- [rule_list.rules.dst_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_set/)
- [rule_list.rules.dst_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_ip_prefix_set/)
- [rule_list.rules.dst_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_label_selector/)
- [rule_list.rules.dst_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_prefix_list/)
- [rule_list.rules.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/)
- [rule_list.rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/ip_prefix_set/)
- [rule_list.rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/label_selector/)
- [rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/metadata/)
- [rule_list.rules.no_http_connect_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/no_http_connect_port/)
- [rule_list.rules.port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/port_matcher/)
- [rule_list.rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/prefix_list/)
- [rule_list.rules.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/tls_list/)
- [rule_list.rules.url_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/url_category_list/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
