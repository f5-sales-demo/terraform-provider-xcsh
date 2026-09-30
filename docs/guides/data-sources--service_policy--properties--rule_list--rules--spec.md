---
page_title: "rule_list.rules.spec"
subcategory: "Security"
description: "rule_list.rules.spec for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 14052, "body_sha256": "sha256:80d5e23b3cf04ce0830681fdb761b1fb89268cd6ee93f16329c0f248b5a4c51d", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_asn", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_client", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_ip", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:arg_matchers", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:asn_list", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:asn_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:body_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:bot_action", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:client_selector", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:cookie_matchers", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:domain_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:headers", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:http_method", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_prefix_list", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:jwt_claims", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:label_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:mum_action", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:path", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:port_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:query_params", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:tls_fingerprint_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:user_identity_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules", "path": "docs/guides/data-sources--service_policy--properties--rule_list--rules--spec.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules.spec

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Property reference](data-sources--service_policy--reference.md)
- [rule_list](data-sources--service_policy--properties--rule_list.md)
- [rule_list.rules](data-sources--service_policy--properties--rule_list--rules.md)
- rule_list.rules.spec

<a id="section"></a>

Type: `"single"`. Computed.

Shape of service\_policy\_rule in the storage backend.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_name\",\"client_name_matcher\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-dst_asn_choice": "[]",
  "x-ves-oneof-field-dst_ip_choice": "[]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"ja4_tls_fingerprint\",\"tls_fingerprint_matcher\"]"
}
```

## Direct properties

<a id="schema-rule_list--rules--spec--action"></a>

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  }
}
```

- [any_asn](data-sources--service_policy--properties--rule_list--rules--spec--any_asn.md): complete subsection reference.

- [any_client](data-sources--service_policy--properties--rule_list--rules--spec--any_client.md): complete subsection reference.

- [any_ip](data-sources--service_policy--properties--rule_list--rules--spec--any_ip.md): complete subsection reference.

- [api_group_matcher](data-sources--service_policy--properties--rule_list--rules--spec--api_group_matcher.md): complete subsection reference.

- [arg_matchers](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers.md): complete subsection reference.

- [asn_list](data-sources--service_policy--properties--rule_list--rules--spec--asn_list.md): complete subsection reference.

- [asn_matcher](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher.md): complete subsection reference.

- [body_matcher](data-sources--service_policy--properties--rule_list--rules--spec--body_matcher.md): complete subsection reference.

- [bot_action](data-sources--service_policy--properties--rule_list--rules--spec--bot_action.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--client_name"></a>

### client_name property

Type: `"string"`. Computed.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Upstream description:

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [client_name_matcher](data-sources--service_policy--properties--rule_list--rules--spec--client_name_matcher.md): complete subsection reference.

- [client_selector](data-sources--service_policy--properties--rule_list--rules--spec--client_selector.md): complete subsection reference.

- [cookie_matchers](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers.md): complete subsection reference.

- [domain_matcher](data-sources--service_policy--properties--rule_list--rules--spec--domain_matcher.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [headers](data-sources--service_policy--properties--rule_list--rules--spec--headers.md): complete subsection reference.

- [http_method](data-sources--service_policy--properties--rule_list--rules--spec--http_method.md): complete subsection reference.

- [ip_matcher](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher.md): complete subsection reference.

- [ip_prefix_list](data-sources--service_policy--properties--rule_list--rules--spec--ip_prefix_list.md): complete subsection reference.

- [ip_threat_category_list](data-sources--service_policy--properties--rule_list--rules--spec--ip_threat_category_list.md): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--service_policy--properties--rule_list--rules--spec--ja4_tls_fingerprint.md): complete subsection reference.

- [jwt_claims](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims.md): complete subsection reference.

- [label_matcher](data-sources--service_policy--properties--rule_list--rules--spec--label_matcher.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--log_rule_evaluation"></a>

### log_rule_evaluation property

Type: `"bool"`. Computed.

Log the rule match details along with the request and continue to evaluate rules in the sequence.

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

- [mum_action](data-sources--service_policy--properties--rule_list--rules--spec--mum_action.md): complete subsection reference.

- [path](data-sources--service_policy--properties--rule_list--rules--spec--path.md): complete subsection reference.

- [port_matcher](data-sources--service_policy--properties--rule_list--rules--spec--port_matcher.md): complete subsection reference.

- [query_params](data-sources--service_policy--properties--rule_list--rules--spec--query_params.md): complete subsection reference.

- [request_constraints](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md): complete subsection reference.

- [segment_policy](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy.md): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--service_policy--properties--rule_list--rules--spec--tls_fingerprint_matcher.md): complete subsection reference.

- [user_identity_matcher](data-sources--service_policy--properties--rule_list--rules--spec--user_identity_matcher.md): complete subsection reference.

- [waf_action](data-sources--service_policy--properties--rule_list--rules--spec--waf_action.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.any_asn](data-sources--service_policy--properties--rule_list--rules--spec--any_asn.md)
- [rule_list.rules.spec.any_client](data-sources--service_policy--properties--rule_list--rules--spec--any_client.md)
- [rule_list.rules.spec.any_ip](data-sources--service_policy--properties--rule_list--rules--spec--any_ip.md)
- [rule_list.rules.spec.api_group_matcher](data-sources--service_policy--properties--rule_list--rules--spec--api_group_matcher.md)
- [rule_list.rules.spec.arg_matchers](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers.md)
- [rule_list.rules.spec.asn_list](data-sources--service_policy--properties--rule_list--rules--spec--asn_list.md)
- [rule_list.rules.spec.asn_matcher](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher.md)
- [rule_list.rules.spec.body_matcher](data-sources--service_policy--properties--rule_list--rules--spec--body_matcher.md)
- [rule_list.rules.spec.bot_action](data-sources--service_policy--properties--rule_list--rules--spec--bot_action.md)
- [rule_list.rules.spec.client_name_matcher](data-sources--service_policy--properties--rule_list--rules--spec--client_name_matcher.md)
- [rule_list.rules.spec.client_selector](data-sources--service_policy--properties--rule_list--rules--spec--client_selector.md)
- [rule_list.rules.spec.cookie_matchers](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers.md)
- [rule_list.rules.spec.domain_matcher](data-sources--service_policy--properties--rule_list--rules--spec--domain_matcher.md)
- [rule_list.rules.spec.headers](data-sources--service_policy--properties--rule_list--rules--spec--headers.md)
- [rule_list.rules.spec.http_method](data-sources--service_policy--properties--rule_list--rules--spec--http_method.md)
- [rule_list.rules.spec.ip_matcher](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher.md)
- [rule_list.rules.spec.ip_prefix_list](data-sources--service_policy--properties--rule_list--rules--spec--ip_prefix_list.md)
- [rule_list.rules.spec.ip_threat_category_list](data-sources--service_policy--properties--rule_list--rules--spec--ip_threat_category_list.md)
- [rule_list.rules.spec.ja4_tls_fingerprint](data-sources--service_policy--properties--rule_list--rules--spec--ja4_tls_fingerprint.md)
- [rule_list.rules.spec.jwt_claims](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims.md)
- [rule_list.rules.spec.label_matcher](data-sources--service_policy--properties--rule_list--rules--spec--label_matcher.md)
- [rule_list.rules.spec.mum_action](data-sources--service_policy--properties--rule_list--rules--spec--mum_action.md)
- [rule_list.rules.spec.path](data-sources--service_policy--properties--rule_list--rules--spec--path.md)
- [rule_list.rules.spec.port_matcher](data-sources--service_policy--properties--rule_list--rules--spec--port_matcher.md)
- [rule_list.rules.spec.query_params](data-sources--service_policy--properties--rule_list--rules--spec--query_params.md)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy.md)
- [rule_list.rules.spec.tls_fingerprint_matcher](data-sources--service_policy--properties--rule_list--rules--spec--tls_fingerprint_matcher.md)
- [rule_list.rules.spec.user_identity_matcher](data-sources--service_policy--properties--rule_list--rules--spec--user_identity_matcher.md)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--properties--rule_list--rules--spec--waf_action.md)
- [rule_list.rules](data-sources--service_policy--properties--rule_list--rules.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
