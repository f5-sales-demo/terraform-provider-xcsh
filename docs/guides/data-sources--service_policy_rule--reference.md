---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 48467, "body_sha256": "sha256:6ff75c76e21ad7d89539d6ec40951fc35861524a1a3bf19fc53e90f460e7f99c", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:reference", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:any_asn", "xcsh-docs:data-sources:service_policy_rule:properties:any_client", "xcsh-docs:data-sources:service_policy_rule:properties:any_ip", "xcsh-docs:data-sources:service_policy_rule:properties:api_group_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:arg_matchers", "xcsh-docs:data-sources:service_policy_rule:properties:asn_list", "xcsh-docs:data-sources:service_policy_rule:properties:asn_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:body_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:bot_action", "xcsh-docs:data-sources:service_policy_rule:properties:client_name_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:client_selector", "xcsh-docs:data-sources:service_policy_rule:properties:cookie_matchers", "xcsh-docs:data-sources:service_policy_rule:properties:domain_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:headers", "xcsh-docs:data-sources:service_policy_rule:properties:http_method", "xcsh-docs:data-sources:service_policy_rule:properties:ip_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:ip_prefix_list", "xcsh-docs:data-sources:service_policy_rule:properties:ip_threat_category_list", "xcsh-docs:data-sources:service_policy_rule:properties:ja4_tls_fingerprint", "xcsh-docs:data-sources:service_policy_rule:properties:jwt_claims", "xcsh-docs:data-sources:service_policy_rule:properties:label_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:mum_action", "xcsh-docs:data-sources:service_policy_rule:properties:path", "xcsh-docs:data-sources:service_policy_rule:properties:port_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:query_params", "xcsh-docs:data-sources:service_policy_rule:properties:request_constraints", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy", "xcsh-docs:data-sources:service_policy_rule:properties:tls_fingerprint_matcher", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:reference", "parent_id": "xcsh-docs:data-sources:service_policy_rule:fundamentals", "path": "docs/guides/data-sources--service_policy_rule--reference.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
- Property reference

## Direct properties

<a id="schema-action"></a>

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

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [any_asn](data-sources--service_policy_rule--properties--any_asn.md): complete subsection reference.

- [any_client](data-sources--service_policy_rule--properties--any_client.md): complete subsection reference.

- [any_ip](data-sources--service_policy_rule--properties--any_ip.md): complete subsection reference.

- [api_group_matcher](data-sources--service_policy_rule--properties--api_group_matcher.md): complete subsection reference.

- [arg_matchers](data-sources--service_policy_rule--properties--arg_matchers.md): complete subsection reference.

- [asn_list](data-sources--service_policy_rule--properties--asn_list.md): complete subsection reference.

- [asn_matcher](data-sources--service_policy_rule--properties--asn_matcher.md): complete subsection reference.

- [body_matcher](data-sources--service_policy_rule--properties--body_matcher.md): complete subsection reference.

- [bot_action](data-sources--service_policy_rule--properties--bot_action.md): complete subsection reference.

<a id="schema-client_name"></a>

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

- [client_name_matcher](data-sources--service_policy_rule--properties--client_name_matcher.md): complete subsection reference.

- [client_selector](data-sources--service_policy_rule--properties--client_selector.md): complete subsection reference.

- [cookie_matchers](data-sources--service_policy_rule--properties--cookie_matchers.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the ServicePolicyRule.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [domain_matcher](data-sources--service_policy_rule--properties--domain_matcher.md): complete subsection reference.

<a id="schema-expiration_timestamp"></a>

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

- [headers](data-sources--service_policy_rule--properties--headers.md): complete subsection reference.

- [http_method](data-sources--service_policy_rule--properties--http_method.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_matcher](data-sources--service_policy_rule--properties--ip_matcher.md): complete subsection reference.

- [ip_prefix_list](data-sources--service_policy_rule--properties--ip_prefix_list.md): complete subsection reference.

- [ip_threat_category_list](data-sources--service_policy_rule--properties--ip_threat_category_list.md): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--service_policy_rule--properties--ja4_tls_fingerprint.md): complete subsection reference.

- [jwt_claims](data-sources--service_policy_rule--properties--jwt_claims.md): complete subsection reference.

- [label_matcher](data-sources--service_policy_rule--properties--label_matcher.md): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-log_rule_evaluation"></a>

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

- [mum_action](data-sources--service_policy_rule--properties--mum_action.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the ServicePolicyRule.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the ServicePolicyRule exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

- [path](data-sources--service_policy_rule--properties--path.md): complete subsection reference.

- [port_matcher](data-sources--service_policy_rule--properties--port_matcher.md): complete subsection reference.

- [query_params](data-sources--service_policy_rule--properties--query_params.md): complete subsection reference.

- [request_constraints](data-sources--service_policy_rule--properties--request_constraints.md): complete subsection reference.

- [segment_policy](data-sources--service_policy_rule--properties--segment_policy.md): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--service_policy_rule--properties--tls_fingerprint_matcher.md): complete subsection reference.

- [waf_action](data-sources--service_policy_rule--properties--waf_action.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--service_policy_rule--reference.md#schema-action) |
| `annotations` | [annotations](data-sources--service_policy_rule--reference.md#schema-annotations) |
| `any_asn` | [any_asn](data-sources--service_policy_rule--properties--any_asn.md#section) |
| `any_client` | [any_client](data-sources--service_policy_rule--properties--any_client.md#section) |
| `any_ip` | [any_ip](data-sources--service_policy_rule--properties--any_ip.md#section) |
| `api_group_matcher` | [api_group_matcher](data-sources--service_policy_rule--properties--api_group_matcher.md#section) |
| `api_group_matcher.invert_matcher` | [api_group_matcher.invert_matcher](data-sources--service_policy_rule--properties--api_group_matcher.md#schema-api_group_matcher--invert_matcher) |
| `api_group_matcher.match` | [api_group_matcher.match](data-sources--service_policy_rule--properties--api_group_matcher.md#schema-api_group_matcher--match) |
| `arg_matchers` | [arg_matchers](data-sources--service_policy_rule--properties--arg_matchers.md#section) |
| `arg_matchers.check_not_present` | [arg_matchers.check_not_present](data-sources--service_policy_rule--properties--arg_matchers--check_not_present.md#section) |
| `arg_matchers.check_present` | [arg_matchers.check_present](data-sources--service_policy_rule--properties--arg_matchers--check_present.md#section) |
| `arg_matchers.invert_matcher` | [arg_matchers.invert_matcher](data-sources--service_policy_rule--properties--arg_matchers.md#schema-arg_matchers--invert_matcher) |
| `arg_matchers.item` | [arg_matchers.item](data-sources--service_policy_rule--properties--arg_matchers--item.md#section) |
| `arg_matchers.item.exact_values` | [arg_matchers.item.exact_values](data-sources--service_policy_rule--properties--arg_matchers--item.md#schema-arg_matchers--item--exact_values) |
| `arg_matchers.item.regex_values` | [arg_matchers.item.regex_values](data-sources--service_policy_rule--properties--arg_matchers--item.md#schema-arg_matchers--item--regex_values) |
| `arg_matchers.item.transformers` | [arg_matchers.item.transformers](data-sources--service_policy_rule--properties--arg_matchers--item.md#schema-arg_matchers--item--transformers) |
| `arg_matchers.name` | [arg_matchers.name](data-sources--service_policy_rule--properties--arg_matchers.md#schema-arg_matchers--name) |
| `asn_list` | [asn_list](data-sources--service_policy_rule--properties--asn_list.md#section) |
| `asn_list.as_numbers` | [asn_list.as_numbers](data-sources--service_policy_rule--properties--asn_list.md#schema-asn_list--as_numbers) |
| `asn_matcher` | [asn_matcher](data-sources--service_policy_rule--properties--asn_matcher.md#section) |
| `asn_matcher.asn_sets` | [asn_matcher.asn_sets](data-sources--service_policy_rule--properties--asn_matcher--asn_sets.md#section) |
| `asn_matcher.asn_sets.kind` | [asn_matcher.asn_sets.kind](data-sources--service_policy_rule--properties--asn_matcher--asn_sets.md#schema-asn_matcher--asn_sets--kind) |
| `asn_matcher.asn_sets.name` | [asn_matcher.asn_sets.name](data-sources--service_policy_rule--properties--asn_matcher--asn_sets.md#schema-asn_matcher--asn_sets--name) |
| `asn_matcher.asn_sets.namespace` | [asn_matcher.asn_sets.namespace](data-sources--service_policy_rule--properties--asn_matcher--asn_sets.md#schema-asn_matcher--asn_sets--namespace) |
| `asn_matcher.asn_sets.tenant` | [asn_matcher.asn_sets.tenant](data-sources--service_policy_rule--properties--asn_matcher--asn_sets.md#schema-asn_matcher--asn_sets--tenant) |
| `asn_matcher.asn_sets.uid` | [asn_matcher.asn_sets.uid](data-sources--service_policy_rule--properties--asn_matcher--asn_sets.md#schema-asn_matcher--asn_sets--uid) |
| `body_matcher` | [body_matcher](data-sources--service_policy_rule--properties--body_matcher.md#section) |
| `body_matcher.exact_values` | [body_matcher.exact_values](data-sources--service_policy_rule--properties--body_matcher.md#schema-body_matcher--exact_values) |
| `body_matcher.regex_values` | [body_matcher.regex_values](data-sources--service_policy_rule--properties--body_matcher.md#schema-body_matcher--regex_values) |
| `body_matcher.transformers` | [body_matcher.transformers](data-sources--service_policy_rule--properties--body_matcher.md#schema-body_matcher--transformers) |
| `bot_action` | [bot_action](data-sources--service_policy_rule--properties--bot_action.md#section) |
| `bot_action.bot_skip_processing` | [bot_action.bot_skip_processing](data-sources--service_policy_rule--properties--bot_action--bot_skip_processing.md#section) |
| `bot_action.none` | [bot_action.none](data-sources--service_policy_rule--properties--bot_action--none.md#section) |
| `client_name` | [client_name](data-sources--service_policy_rule--reference.md#schema-client_name) |
| `client_name_matcher` | [client_name_matcher](data-sources--service_policy_rule--properties--client_name_matcher.md#section) |
| `client_name_matcher.exact_values` | [client_name_matcher.exact_values](data-sources--service_policy_rule--properties--client_name_matcher.md#schema-client_name_matcher--exact_values) |
| `client_name_matcher.regex_values` | [client_name_matcher.regex_values](data-sources--service_policy_rule--properties--client_name_matcher.md#schema-client_name_matcher--regex_values) |
| `client_selector` | [client_selector](data-sources--service_policy_rule--properties--client_selector.md#section) |
| `client_selector.expressions` | [client_selector.expressions](data-sources--service_policy_rule--properties--client_selector.md#schema-client_selector--expressions) |
| `cookie_matchers` | [cookie_matchers](data-sources--service_policy_rule--properties--cookie_matchers.md#section) |
| `cookie_matchers.check_not_present` | [cookie_matchers.check_not_present](data-sources--service_policy_rule--properties--cookie_matchers--check_not_present.md#section) |
| `cookie_matchers.check_present` | [cookie_matchers.check_present](data-sources--service_policy_rule--properties--cookie_matchers--check_present.md#section) |
| `cookie_matchers.invert_matcher` | [cookie_matchers.invert_matcher](data-sources--service_policy_rule--properties--cookie_matchers.md#schema-cookie_matchers--invert_matcher) |
| `cookie_matchers.item` | [cookie_matchers.item](data-sources--service_policy_rule--properties--cookie_matchers--item.md#section) |
| `cookie_matchers.item.exact_values` | [cookie_matchers.item.exact_values](data-sources--service_policy_rule--properties--cookie_matchers--item.md#schema-cookie_matchers--item--exact_values) |
| `cookie_matchers.item.regex_values` | [cookie_matchers.item.regex_values](data-sources--service_policy_rule--properties--cookie_matchers--item.md#schema-cookie_matchers--item--regex_values) |
| `cookie_matchers.item.transformers` | [cookie_matchers.item.transformers](data-sources--service_policy_rule--properties--cookie_matchers--item.md#schema-cookie_matchers--item--transformers) |
| `cookie_matchers.name` | [cookie_matchers.name](data-sources--service_policy_rule--properties--cookie_matchers.md#schema-cookie_matchers--name) |
| `description` | [description](data-sources--service_policy_rule--reference.md#schema-description) |
| `domain_matcher` | [domain_matcher](data-sources--service_policy_rule--properties--domain_matcher.md#section) |
| `domain_matcher.exact_values` | [domain_matcher.exact_values](data-sources--service_policy_rule--properties--domain_matcher.md#schema-domain_matcher--exact_values) |
| `domain_matcher.regex_values` | [domain_matcher.regex_values](data-sources--service_policy_rule--properties--domain_matcher.md#schema-domain_matcher--regex_values) |
| `expiration_timestamp` | [expiration_timestamp](data-sources--service_policy_rule--reference.md#schema-expiration_timestamp) |
| `headers` | [headers](data-sources--service_policy_rule--properties--headers.md#section) |
| `headers.check_not_present` | [headers.check_not_present](data-sources--service_policy_rule--properties--headers--check_not_present.md#section) |
| `headers.check_present` | [headers.check_present](data-sources--service_policy_rule--properties--headers--check_present.md#section) |
| `headers.invert_matcher` | [headers.invert_matcher](data-sources--service_policy_rule--properties--headers.md#schema-headers--invert_matcher) |
| `headers.item` | [headers.item](data-sources--service_policy_rule--properties--headers--item.md#section) |
| `headers.item.exact_values` | [headers.item.exact_values](data-sources--service_policy_rule--properties--headers--item.md#schema-headers--item--exact_values) |
| `headers.item.regex_values` | [headers.item.regex_values](data-sources--service_policy_rule--properties--headers--item.md#schema-headers--item--regex_values) |
| `headers.item.transformers` | [headers.item.transformers](data-sources--service_policy_rule--properties--headers--item.md#schema-headers--item--transformers) |
| `headers.name` | [headers.name](data-sources--service_policy_rule--properties--headers.md#schema-headers--name) |
| `http_method` | [http_method](data-sources--service_policy_rule--properties--http_method.md#section) |
| `http_method.invert_matcher` | [http_method.invert_matcher](data-sources--service_policy_rule--properties--http_method.md#schema-http_method--invert_matcher) |
| `http_method.methods` | [http_method.methods](data-sources--service_policy_rule--properties--http_method.md#schema-http_method--methods) |
| `id` | [id](data-sources--service_policy_rule--reference.md#schema-id) |
| `ip_matcher` | [ip_matcher](data-sources--service_policy_rule--properties--ip_matcher.md#section) |
| `ip_matcher.invert_matcher` | [ip_matcher.invert_matcher](data-sources--service_policy_rule--properties--ip_matcher.md#schema-ip_matcher--invert_matcher) |
| `ip_matcher.prefix_sets` | [ip_matcher.prefix_sets](data-sources--service_policy_rule--properties--ip_matcher--prefix_sets.md#section) |
| `ip_matcher.prefix_sets.kind` | [ip_matcher.prefix_sets.kind](data-sources--service_policy_rule--properties--ip_matcher--prefix_sets.md#schema-ip_matcher--prefix_sets--kind) |
| `ip_matcher.prefix_sets.name` | [ip_matcher.prefix_sets.name](data-sources--service_policy_rule--properties--ip_matcher--prefix_sets.md#schema-ip_matcher--prefix_sets--name) |
| `ip_matcher.prefix_sets.namespace` | [ip_matcher.prefix_sets.namespace](data-sources--service_policy_rule--properties--ip_matcher--prefix_sets.md#schema-ip_matcher--prefix_sets--namespace) |
| `ip_matcher.prefix_sets.tenant` | [ip_matcher.prefix_sets.tenant](data-sources--service_policy_rule--properties--ip_matcher--prefix_sets.md#schema-ip_matcher--prefix_sets--tenant) |
| `ip_matcher.prefix_sets.uid` | [ip_matcher.prefix_sets.uid](data-sources--service_policy_rule--properties--ip_matcher--prefix_sets.md#schema-ip_matcher--prefix_sets--uid) |
| `ip_prefix_list` | [ip_prefix_list](data-sources--service_policy_rule--properties--ip_prefix_list.md#section) |
| `ip_prefix_list.invert_match` | [ip_prefix_list.invert_match](data-sources--service_policy_rule--properties--ip_prefix_list.md#schema-ip_prefix_list--invert_match) |
| `ip_prefix_list.ip_prefixes` | [ip_prefix_list.ip_prefixes](data-sources--service_policy_rule--properties--ip_prefix_list.md#schema-ip_prefix_list--ip_prefixes) |
| `ip_threat_category_list` | [ip_threat_category_list](data-sources--service_policy_rule--properties--ip_threat_category_list.md#section) |
| `ip_threat_category_list.ip_threat_categories` | [ip_threat_category_list.ip_threat_categories](data-sources--service_policy_rule--properties--ip_threat_category_list.md#schema-ip_threat_category_list--ip_threat_categories) |
| `ja4_tls_fingerprint` | [ja4_tls_fingerprint](data-sources--service_policy_rule--properties--ja4_tls_fingerprint.md#section) |
| `ja4_tls_fingerprint.exact_values` | [ja4_tls_fingerprint.exact_values](data-sources--service_policy_rule--properties--ja4_tls_fingerprint.md#schema-ja4_tls_fingerprint--exact_values) |
| `jwt_claims` | [jwt_claims](data-sources--service_policy_rule--properties--jwt_claims.md#section) |
| `jwt_claims.check_not_present` | [jwt_claims.check_not_present](data-sources--service_policy_rule--properties--jwt_claims--check_not_present.md#section) |
| `jwt_claims.check_present` | [jwt_claims.check_present](data-sources--service_policy_rule--properties--jwt_claims--check_present.md#section) |
| `jwt_claims.invert_matcher` | [jwt_claims.invert_matcher](data-sources--service_policy_rule--properties--jwt_claims.md#schema-jwt_claims--invert_matcher) |
| `jwt_claims.item` | [jwt_claims.item](data-sources--service_policy_rule--properties--jwt_claims--item.md#section) |
| `jwt_claims.item.exact_values` | [jwt_claims.item.exact_values](data-sources--service_policy_rule--properties--jwt_claims--item.md#schema-jwt_claims--item--exact_values) |
| `jwt_claims.item.regex_values` | [jwt_claims.item.regex_values](data-sources--service_policy_rule--properties--jwt_claims--item.md#schema-jwt_claims--item--regex_values) |
| `jwt_claims.item.transformers` | [jwt_claims.item.transformers](data-sources--service_policy_rule--properties--jwt_claims--item.md#schema-jwt_claims--item--transformers) |
| `jwt_claims.name` | [jwt_claims.name](data-sources--service_policy_rule--properties--jwt_claims.md#schema-jwt_claims--name) |
| `label_matcher` | [label_matcher](data-sources--service_policy_rule--properties--label_matcher.md#section) |
| `label_matcher.keys` | [label_matcher.keys](data-sources--service_policy_rule--properties--label_matcher.md#schema-label_matcher--keys) |
| `labels` | [labels](data-sources--service_policy_rule--reference.md#schema-labels) |
| `log_rule_evaluation` | [log_rule_evaluation](data-sources--service_policy_rule--reference.md#schema-log_rule_evaluation) |
| `mum_action` | [mum_action](data-sources--service_policy_rule--properties--mum_action.md#section) |
| `mum_action.default` | [mum_action.default](data-sources--service_policy_rule--properties--mum_action--default.md#section) |
| `mum_action.skip_processing` | [mum_action.skip_processing](data-sources--service_policy_rule--properties--mum_action--skip_processing.md#section) |
| `name` | [name](data-sources--service_policy_rule--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--service_policy_rule--reference.md#schema-namespace) |
| `path` | [path](data-sources--service_policy_rule--properties--path.md#section) |
| `path.encoded_path_matcher` | [path.encoded_path_matcher](data-sources--service_policy_rule--properties--path.md#schema-path--encoded_path_matcher) |
| `path.exact_values` | [path.exact_values](data-sources--service_policy_rule--properties--path.md#schema-path--exact_values) |
| `path.invert_matcher` | [path.invert_matcher](data-sources--service_policy_rule--properties--path.md#schema-path--invert_matcher) |
| `path.prefix_values` | [path.prefix_values](data-sources--service_policy_rule--properties--path.md#schema-path--prefix_values) |
| `path.regex_values` | [path.regex_values](data-sources--service_policy_rule--properties--path.md#schema-path--regex_values) |
| `path.suffix_values` | [path.suffix_values](data-sources--service_policy_rule--properties--path.md#schema-path--suffix_values) |
| `path.transformers` | [path.transformers](data-sources--service_policy_rule--properties--path.md#schema-path--transformers) |
| `port_matcher` | [port_matcher](data-sources--service_policy_rule--properties--port_matcher.md#section) |
| `port_matcher.invert_matcher` | [port_matcher.invert_matcher](data-sources--service_policy_rule--properties--port_matcher.md#schema-port_matcher--invert_matcher) |
| `port_matcher.ports` | [port_matcher.ports](data-sources--service_policy_rule--properties--port_matcher.md#schema-port_matcher--ports) |
| `query_params` | [query_params](data-sources--service_policy_rule--properties--query_params.md#section) |
| `query_params.check_not_present` | [query_params.check_not_present](data-sources--service_policy_rule--properties--query_params--check_not_present.md#section) |
| `query_params.check_present` | [query_params.check_present](data-sources--service_policy_rule--properties--query_params--check_present.md#section) |
| `query_params.invert_matcher` | [query_params.invert_matcher](data-sources--service_policy_rule--properties--query_params.md#schema-query_params--invert_matcher) |
| `query_params.item` | [query_params.item](data-sources--service_policy_rule--properties--query_params--item.md#section) |
| `query_params.item.exact_values` | [query_params.item.exact_values](data-sources--service_policy_rule--properties--query_params--item.md#schema-query_params--item--exact_values) |
| `query_params.item.regex_values` | [query_params.item.regex_values](data-sources--service_policy_rule--properties--query_params--item.md#schema-query_params--item--regex_values) |
| `query_params.item.transformers` | [query_params.item.transformers](data-sources--service_policy_rule--properties--query_params--item.md#schema-query_params--item--transformers) |
| `query_params.key` | [query_params.key](data-sources--service_policy_rule--properties--query_params.md#schema-query_params--key) |
| `request_constraints` | [request_constraints](data-sources--service_policy_rule--properties--request_constraints.md#section) |
| `request_constraints.max_cookie_count_exceeds` | [request_constraints.max_cookie_count_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_cookie_count_exceeds) |
| `request_constraints.max_cookie_count_none` | [request_constraints.max_cookie_count_none](data-sources--service_policy_rule--properties--request_constraints--max_cookie_count_none.md#section) |
| `request_constraints.max_cookie_key_size_exceeds` | [request_constraints.max_cookie_key_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_cookie_key_size_exceeds) |
| `request_constraints.max_cookie_key_size_none` | [request_constraints.max_cookie_key_size_none](data-sources--service_policy_rule--properties--request_constraints--max_cookie_key_size_none.md#section) |
| `request_constraints.max_cookie_value_size_exceeds` | [request_constraints.max_cookie_value_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_cookie_value_size_exceeds) |
| `request_constraints.max_cookie_value_size_none` | [request_constraints.max_cookie_value_size_none](data-sources--service_policy_rule--properties--request_constraints--max_cookie_value_size_none.md#section) |
| `request_constraints.max_header_count_exceeds` | [request_constraints.max_header_count_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_header_count_exceeds) |
| `request_constraints.max_header_count_none` | [request_constraints.max_header_count_none](data-sources--service_policy_rule--properties--request_constraints--max_header_count_none.md#section) |
| `request_constraints.max_header_key_size_exceeds` | [request_constraints.max_header_key_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_header_key_size_exceeds) |
| `request_constraints.max_header_key_size_none` | [request_constraints.max_header_key_size_none](data-sources--service_policy_rule--properties--request_constraints--max_header_key_size_none.md#section) |
| `request_constraints.max_header_value_size_exceeds` | [request_constraints.max_header_value_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_header_value_size_exceeds) |
| `request_constraints.max_header_value_size_none` | [request_constraints.max_header_value_size_none](data-sources--service_policy_rule--properties--request_constraints--max_header_value_size_none.md#section) |
| `request_constraints.max_parameter_count_exceeds` | [request_constraints.max_parameter_count_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_parameter_count_exceeds) |
| `request_constraints.max_parameter_count_none` | [request_constraints.max_parameter_count_none](data-sources--service_policy_rule--properties--request_constraints--max_parameter_count_none.md#section) |
| `request_constraints.max_parameter_name_size_exceeds` | [request_constraints.max_parameter_name_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_parameter_name_size_exceeds) |
| `request_constraints.max_parameter_name_size_none` | [request_constraints.max_parameter_name_size_none](data-sources--service_policy_rule--properties--request_constraints--max_parameter_name_size_none.md#section) |
| `request_constraints.max_parameter_value_size_exceeds` | [request_constraints.max_parameter_value_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_parameter_value_size_exceeds) |
| `request_constraints.max_parameter_value_size_none` | [request_constraints.max_parameter_value_size_none](data-sources--service_policy_rule--properties--request_constraints--max_parameter_value_size_none.md#section) |
| `request_constraints.max_query_size_exceeds` | [request_constraints.max_query_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_query_size_exceeds) |
| `request_constraints.max_query_size_none` | [request_constraints.max_query_size_none](data-sources--service_policy_rule--properties--request_constraints--max_query_size_none.md#section) |
| `request_constraints.max_request_line_size_exceeds` | [request_constraints.max_request_line_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_request_line_size_exceeds) |
| `request_constraints.max_request_line_size_none` | [request_constraints.max_request_line_size_none](data-sources--service_policy_rule--properties--request_constraints--max_request_line_size_none.md#section) |
| `request_constraints.max_request_size_exceeds` | [request_constraints.max_request_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_request_size_exceeds) |
| `request_constraints.max_request_size_none` | [request_constraints.max_request_size_none](data-sources--service_policy_rule--properties--request_constraints--max_request_size_none.md#section) |
| `request_constraints.max_url_size_exceeds` | [request_constraints.max_url_size_exceeds](data-sources--service_policy_rule--properties--request_constraints.md#schema-request_constraints--max_url_size_exceeds) |
| `request_constraints.max_url_size_none` | [request_constraints.max_url_size_none](data-sources--service_policy_rule--properties--request_constraints--max_url_size_none.md#section) |
| `segment_policy` | [segment_policy](data-sources--service_policy_rule--properties--segment_policy.md#section) |
| `segment_policy.dst_any` | [segment_policy.dst_any](data-sources--service_policy_rule--properties--segment_policy--dst_any.md#section) |
| `segment_policy.dst_segments` | [segment_policy.dst_segments](data-sources--service_policy_rule--properties--segment_policy--dst_segments.md#section) |
| `segment_policy.dst_segments.segments` | [segment_policy.dst_segments.segments](data-sources--service_policy_rule--properties--segment_policy--dst_segments--segments.md#section) |
| `segment_policy.dst_segments.segments.name` | [segment_policy.dst_segments.segments.name](data-sources--service_policy_rule--properties--segment_policy--dst_segments--segments.md#schema-segment_policy--dst_segments--segments--name) |
| `segment_policy.dst_segments.segments.namespace` | [segment_policy.dst_segments.segments.namespace](data-sources--service_policy_rule--properties--segment_policy--dst_segments--segments.md#schema-segment_policy--dst_segments--segments--namespace) |
| `segment_policy.dst_segments.segments.tenant` | [segment_policy.dst_segments.segments.tenant](data-sources--service_policy_rule--properties--segment_policy--dst_segments--segments.md#schema-segment_policy--dst_segments--segments--tenant) |
| `segment_policy.intra_segment` | [segment_policy.intra_segment](data-sources--service_policy_rule--properties--segment_policy--intra_segment.md#section) |
| `segment_policy.src_any` | [segment_policy.src_any](data-sources--service_policy_rule--properties--segment_policy--src_any.md#section) |
| `segment_policy.src_segments` | [segment_policy.src_segments](data-sources--service_policy_rule--properties--segment_policy--src_segments.md#section) |
| `segment_policy.src_segments.segments` | [segment_policy.src_segments.segments](data-sources--service_policy_rule--properties--segment_policy--src_segments--segments.md#section) |
| `segment_policy.src_segments.segments.name` | [segment_policy.src_segments.segments.name](data-sources--service_policy_rule--properties--segment_policy--src_segments--segments.md#schema-segment_policy--src_segments--segments--name) |
| `segment_policy.src_segments.segments.namespace` | [segment_policy.src_segments.segments.namespace](data-sources--service_policy_rule--properties--segment_policy--src_segments--segments.md#schema-segment_policy--src_segments--segments--namespace) |
| `segment_policy.src_segments.segments.tenant` | [segment_policy.src_segments.segments.tenant](data-sources--service_policy_rule--properties--segment_policy--src_segments--segments.md#schema-segment_policy--src_segments--segments--tenant) |
| `tls_fingerprint_matcher` | [tls_fingerprint_matcher](data-sources--service_policy_rule--properties--tls_fingerprint_matcher.md#section) |
| `tls_fingerprint_matcher.classes` | [tls_fingerprint_matcher.classes](data-sources--service_policy_rule--properties--tls_fingerprint_matcher.md#schema-tls_fingerprint_matcher--classes) |
| `tls_fingerprint_matcher.exact_values` | [tls_fingerprint_matcher.exact_values](data-sources--service_policy_rule--properties--tls_fingerprint_matcher.md#schema-tls_fingerprint_matcher--exact_values) |
| `tls_fingerprint_matcher.excluded_values` | [tls_fingerprint_matcher.excluded_values](data-sources--service_policy_rule--properties--tls_fingerprint_matcher.md#schema-tls_fingerprint_matcher--excluded_values) |
| `waf_action` | [waf_action](data-sources--service_policy_rule--properties--waf_action.md#section) |
| `waf_action.app_firewall_detection_control` | [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md#section) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md#section) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_attack_type_contexts--context) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_attack_type_contexts--context_name) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_attack_type_contexts--exclude_attack_type) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md#section) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts` | [waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_signature_contexts.md#section) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_signature_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_signature_contexts--context) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_signature_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_signature_contexts--context_name) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_signature_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_signature_contexts--signature_id) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts` | [waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_violation_contexts.md#section) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_violation_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_violation_contexts--context) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_violation_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_violation_contexts--context_name) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_violation_contexts.md#schema-waf_action--app_firewall_detection_control--exclude_violation_contexts--exclude_violation) |
| `waf_action.none` | [waf_action.none](data-sources--service_policy_rule--properties--waf_action--none.md#section) |
| `waf_action.waf_skip_processing` | [waf_action.waf_skip_processing](data-sources--service_policy_rule--properties--waf_action--waf_skip_processing.md#section) |

## Next pages

- [any_asn](data-sources--service_policy_rule--properties--any_asn.md)
- [any_client](data-sources--service_policy_rule--properties--any_client.md)
- [any_ip](data-sources--service_policy_rule--properties--any_ip.md)
- [api_group_matcher](data-sources--service_policy_rule--properties--api_group_matcher.md)
- [arg_matchers](data-sources--service_policy_rule--properties--arg_matchers.md)
- [asn_list](data-sources--service_policy_rule--properties--asn_list.md)
- [asn_matcher](data-sources--service_policy_rule--properties--asn_matcher.md)
- [body_matcher](data-sources--service_policy_rule--properties--body_matcher.md)
- [bot_action](data-sources--service_policy_rule--properties--bot_action.md)
- [client_name_matcher](data-sources--service_policy_rule--properties--client_name_matcher.md)
- [client_selector](data-sources--service_policy_rule--properties--client_selector.md)
- [cookie_matchers](data-sources--service_policy_rule--properties--cookie_matchers.md)
- [domain_matcher](data-sources--service_policy_rule--properties--domain_matcher.md)
- [headers](data-sources--service_policy_rule--properties--headers.md)
- [http_method](data-sources--service_policy_rule--properties--http_method.md)
- [ip_matcher](data-sources--service_policy_rule--properties--ip_matcher.md)
- [ip_prefix_list](data-sources--service_policy_rule--properties--ip_prefix_list.md)
- [ip_threat_category_list](data-sources--service_policy_rule--properties--ip_threat_category_list.md)
- [ja4_tls_fingerprint](data-sources--service_policy_rule--properties--ja4_tls_fingerprint.md)
- [jwt_claims](data-sources--service_policy_rule--properties--jwt_claims.md)
- [label_matcher](data-sources--service_policy_rule--properties--label_matcher.md)
- [mum_action](data-sources--service_policy_rule--properties--mum_action.md)
- [path](data-sources--service_policy_rule--properties--path.md)
- [port_matcher](data-sources--service_policy_rule--properties--port_matcher.md)
- [query_params](data-sources--service_policy_rule--properties--query_params.md)
- [request_constraints](data-sources--service_policy_rule--properties--request_constraints.md)
- [segment_policy](data-sources--service_policy_rule--properties--segment_policy.md)
- [tls_fingerprint_matcher](data-sources--service_policy_rule--properties--tls_fingerprint_matcher.md)
- [waf_action](data-sources--service_policy_rule--properties--waf_action.md)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
