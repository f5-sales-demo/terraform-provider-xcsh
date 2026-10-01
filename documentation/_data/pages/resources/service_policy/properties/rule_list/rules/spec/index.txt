---
page_title: "rule_list.rules.spec"
subcategory: "Security"
description: "rule_list.rules.spec for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 19103, "body_sha256": "sha256:403c3e9beb436e05192a30d54eb1bd920dc8c2e6230df9aa22ec2406aeabf6b4", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_asn", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_client", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_ip", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_list", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:body_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:bot_action", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_selector", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:cookie_matchers", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:domain_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:headers", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:http_method", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_prefix_list", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:jwt_claims", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:label_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:path", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:port_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:query_params", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:tls_fingerprint_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:user_identity_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action"], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/index.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["rule_list", "rules", "spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- rule_list.rules.spec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Shape of service\_policy\_rule in the storage backend.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("action",
    "waf_action"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_name"),
  validators.ConflictingObjectAttributes("any_client",
    "client_name_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("client_name",
    "client_name_matcher"),
  validators.ConflictingObjectAttributes("client_name",
    "client_selector"),
  validators.ConflictingObjectAttributes("client_name",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("client_name_matcher",
    "client_selector"),
  validators.ConflictingObjectAttributes("client_name_matcher",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("ja4_tls_fingerprint",
    "tls_fingerprint_matcher")}
```

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

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--spec--action"></a>

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  }
}
```

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_asn/): complete subsection reference.

- [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_client/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_ip/): complete subsection reference.

- [api_group_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/): complete subsection reference.

- [arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/): complete subsection reference.

- [body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/body_matcher/): complete subsection reference.

- [bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/bot_action/): complete subsection reference.

<a id="schema-rule_list--rules--spec--client_name"></a>

### client_name property

Type: `"string"`. Optional.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Upstream description:

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_name_matcher/): complete subsection reference.

- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_selector/): complete subsection reference.

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/): complete subsection reference.

- [domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/domain_matcher/): complete subsection reference.

<a id="schema-rule_list--rules--spec--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Optional.

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/): complete subsection reference.

- [http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/http_method/): complete subsection reference.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_prefix_list/): complete subsection reference.

- [ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_threat_category_list/): complete subsection reference.

- [ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ja4_tls_fingerprint/): complete subsection reference.

- [jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/label_matcher/): complete subsection reference.

<a id="schema-rule_list--rules--spec--log_rule_evaluation"></a>

### log_rule_evaluation property

Type: `"bool"`. Optional.

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

- [mum_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/): complete subsection reference.

- [port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/port_matcher/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/): complete subsection reference.

- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/): complete subsection reference.

- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/): complete subsection reference.

- [tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/): complete subsection reference.

- [user_identity_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/user_identity_matcher/): complete subsection reference.

- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_asn/)
- [rule_list.rules.spec.any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_client/)
- [rule_list.rules.spec.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_ip/)
- [rule_list.rules.spec.api_group_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/)
- [rule_list.rules.spec.arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/)
- [rule_list.rules.spec.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_list/)
- [rule_list.rules.spec.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/)
- [rule_list.rules.spec.body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/body_matcher/)
- [rule_list.rules.spec.bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/bot_action/)
- [rule_list.rules.spec.client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_name_matcher/)
- [rule_list.rules.spec.client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_selector/)
- [rule_list.rules.spec.cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/)
- [rule_list.rules.spec.domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/domain_matcher/)
- [rule_list.rules.spec.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/)
- [rule_list.rules.spec.http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/http_method/)
- [rule_list.rules.spec.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/)
- [rule_list.rules.spec.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_prefix_list/)
- [rule_list.rules.spec.ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_threat_category_list/)
- [rule_list.rules.spec.ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ja4_tls_fingerprint/)
- [rule_list.rules.spec.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/)
- [rule_list.rules.spec.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/label_matcher/)
- [rule_list.rules.spec.mum_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/)
- [rule_list.rules.spec.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/)
- [rule_list.rules.spec.port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/port_matcher/)
- [rule_list.rules.spec.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/)
- [rule_list.rules.spec.request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/)
- [rule_list.rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/)
- [rule_list.rules.spec.tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/)
- [rule_list.rules.spec.user_identity_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/user_identity_matcher/)
- [rule_list.rules.spec.waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
