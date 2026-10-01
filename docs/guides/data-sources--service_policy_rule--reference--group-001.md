---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84e0a5daae54f6ec917c1d86a22c1f7e00eb4682e39e57df7b4a57f7b629db77"></a>

## Property reference — Property reference / dc726d2ff366 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- Property reference

<a id="canonical-7d6fa6ee012f8eabb25594514db266e9203b2bfa61bba69122dcf1e113b1915a"></a>

## Direct properties — Property reference / dc726d2ff366 / 3

<a id="canonical-f0a7f939451ae3ae5a28deb6c2f5b952e920e2631a84a42da29c981e585794ba"></a>

<a id="canonical-22844d5cf918ef46b4d106f045ecd07ed17d49a21a89270e85940fa299af79c6"></a>

## action property — Property reference / dc726d2ff366 / 4

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

<a id="canonical-9393bde47ea2742c802e1c1921e3db67925512c28fe65afa24422b8526082a31"></a>

<a id="canonical-054186f958d6c281d80cde932d4ceb914b359f5f5358bd2810ed63732ab5890a"></a>

## annotations property — Property reference / dc726d2ff366 / 5

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

- [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-c9fadc1947ea10969e9fda773fb6c34b205c8f4e3f8d9acfe60230a9aee5b48c): complete subsection reference.

- [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-a187c7105b0492c8f3bfade3fd49556f779f9196fe2f4b923fc6d8a097555b28): complete subsection reference.

- [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-84ebdc2357bec6ce8b16bbbc3baa8b4017493006464f42c495020e236eae1fb1): complete subsection reference.

- [api_group_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-bba1340cb0a6a7f136e7ead6bcb536cf8e805788d0033f246ac5769cb6039510): complete subsection reference.

- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6): complete subsection reference.

- [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-bb506919f0604945c9733e23725137ca93905f91f5750bb30a77adc2a31102d6): complete subsection reference.

- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-9b711c97a0ab1fdd142a22047f397fbfdbe65db43aa3a3e01006f75953bcca68): complete subsection reference.

- [body_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0b0bdac8179833c9adc7559d4d46d8c3073bbb200c6aae22807622a45623baa4): complete subsection reference.

- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-60950b56578b9a3d2e7f01afbf2f852364b3567038338171b36ef07220a513b8): complete subsection reference.

<a id="canonical-69dde0972f4803be8c008b5aaca1d11e91820582e9f1dbb4cfddfd54b3793d2c"></a>

<a id="canonical-9c8f800d65980ed9d4020b2cbcc2993050e7ff611b1d5a94727cf9e3116d5f95"></a>

## client_name property — Property reference / dc726d2ff366 / 6

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

- [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3fc5a45dd8f5f55acc9df34955eb8827a3995b6449583a84c5305f74f2e29351): complete subsection reference.

- [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-5bf025ca77ff33756a6188586d3d1dc585b82381fca91461ce07475b6b7719bf): complete subsection reference.

- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae): complete subsection reference.

<a id="canonical-41b402c8db3874ad19d5b8f9ae726dee9bc560207884a96ce7b1b6a9a99bf0b9"></a>

<a id="canonical-e784b1fc84978a9819a6e2fa0cb5e8cb739dc1d0817f38cd8ca90d1afe2570bc"></a>

## description property — Property reference / dc726d2ff366 / 7

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

- [domain_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-fa39c24dd24d6084f6b5f364ba1908e727261ce5b78f84e14d6f55bb96699695): complete subsection reference.

<a id="canonical-e3f37720e2613a117f6915c19fbfc6df3650fb7243ebe46c97772687c6609cab"></a>

<a id="canonical-b0e20b8bb086c72db1d15255fa6a4b0e82c801dec9f23a8b3d52b22857e82ccb"></a>

## expiration_timestamp property — Property reference / dc726d2ff366 / 8

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

- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8): complete subsection reference.

- [http_method](data-sources--service_policy_rule--reference--group-001.md#canonical-2ccfac68219f29dec73d5b10667fdddc8a3d5f24c54e6fdbd705961ddb8382bd): complete subsection reference.

<a id="canonical-d75958381fd77af37223a4d2d5c8a6d955a983062f91918c3b17ceab9812c587"></a>

<a id="canonical-dcc34806f0951466192956ed6969e76b1cce9772c6498686a369b65501d3cf40"></a>

## id property — Property reference / dc726d2ff366 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-4c64fdc40e1c3b6cadd843f10a9fcacaa6d87d0a051cb31ababde925f2d339dc): complete subsection reference.

- [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-f568569c6da6e1d479ee5ca4c660b0f3b3a6e5bdd4c77d6f54811e6a1ffd3098): complete subsection reference.

- [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-e2480d86388e5c9fb566e99804e4614053a7dba6f8db96780101bd841aaf01f7): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-fb9290901816a8a4f9e84e612d576c12566a1b4d776627310a573f99ef9fcfc3): complete subsection reference.

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36): complete subsection reference.

- [label_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-e4d40e388ca9969028510e765949d2ecf07e9d65a0c3bd535e05cee68412ce0b): complete subsection reference.

<a id="canonical-2e0e5113e680e19c0f4e8c544b25e9cf84ce9ecfce899548da5d27b3c03430d9"></a>

<a id="canonical-0bf61b03c2ed67b9ce6bb7d54d44559d5b2206a733bb54274137e65610042077"></a>

## labels property — Property reference / dc726d2ff366 / 10

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

<a id="canonical-51b39c1b7d26a64aa16eb836fbfe24076aa17e4752d0b95cedb31d56b1a52a3d"></a>

<a id="canonical-140ac30138b90f659bf51a0e12d963b8a01f772a3c517a1628994a1b79ce27b6"></a>

## log_rule_evaluation property — Property reference / dc726d2ff366 / 11

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

- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-137e55d694ff101b1eceed4b0f2170fade6d8b7a600a590d2a83bb1efb5d0eb3): complete subsection reference.

<a id="canonical-f61f250350b7a614efbdbce9c5c6ed8651b58798805698717fd523a323ad0848"></a>

<a id="canonical-6380fe6fc0ba5eba78d50c4ef91331971167abb27ebbf99d7e8a2af486881ce0"></a>

## name property — Property reference / dc726d2ff366 / 12

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

<a id="canonical-dcaf02c58a17dc42a260bbf07c8ab46a936b09c98c6130ee0464482a8165ea3d"></a>

<a id="canonical-00b40fa6c787968d747df2090e44cc96f8078d0049b7f710b4998cfa9233b703"></a>

## namespace property — Property reference / dc726d2ff366 / 13

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

- [path](data-sources--service_policy_rule--reference--group-002.md#canonical-62ecded1c4e84948a11dba301846deb27ceaebac027c0ed9896d35e68b2ae3c7): complete subsection reference.

- [port_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-91908ed3f38a27795fb956d472ae83414d3046dc441e006972b9304b7899f358): complete subsection reference.

- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3): complete subsection reference.

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4): complete subsection reference.

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-1fc6b6f8ae22871967881106ea255eba23defd48f77829156aa2fdf56202e7ca): complete subsection reference.

- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a): complete subsection reference.

<a id="canonical-b9f952b505c20e9d421c926847aa1a4a4f2b5922a82e3a1f1a31afa36660d6c1"></a>

## All schema paths — Property reference / dc726d2ff366 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--service_policy_rule--reference--group-001.md#canonical-f0a7f939451ae3ae5a28deb6c2f5b952e920e2631a84a42da29c981e585794ba) |
| `annotations` | [annotations](data-sources--service_policy_rule--reference--group-001.md#canonical-9393bde47ea2742c802e1c1921e3db67925512c28fe65afa24422b8526082a31) |
| `any_asn` | [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-53efe7095c1382b98b12fbcb16d318d032ad0782a66079042f357a2be7eb1749) |
| `any_client` | [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-aefd0b7d81871a8e4bd8bf2a976341dcbd6ccf876494361f2d155185914ea3e9) |
| `any_ip` | [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-37e2c167720c140231d0639a294bee12696a1f423fd3182aecb347bf93fb8a73) |
| `api_group_matcher` | [api_group_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-78a06280537d7f18fc43e762776e979f401ed05aad3218144634670e40a60431) |
| `api_group_matcher.invert_matcher` | [api_group_matcher.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-598426c9b7094e4909395205744aae9850f52bbc8e9bd64e692d607f9774b04a) |
| `api_group_matcher.match` | [api_group_matcher.match](data-sources--service_policy_rule--reference--group-001.md#canonical-4256bd8214ddc9409267603e515c1144b564800c084c3f9a4c4ed2f744764582) |
| `arg_matchers` | [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8b997cb3a29a65bb0850c07e6dca1d4245242ef8b30e1f6e0483b117927bfada) |
| `arg_matchers.check_not_present` | [arg_matchers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-8ad3893bf2dba3208ac168ced1b81864c9edd6dc570c193a8a48179500f89d77) |
| `arg_matchers.check_present` | [arg_matchers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-ffb79fe5bd5a6621289ca633fba77d0726323ec876f807ffaa92ccabe11ce76a) |
| `arg_matchers.invert_matcher` | [arg_matchers.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-7cc0d00a655443f6428fcf6f78bc4e2ba2facaaf54eeb96a8404317c5915eb24) |
| `arg_matchers.item` | [arg_matchers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-c0e6a9e264fc73074ac774a576f969a6cb382848bf6fc1ffa19b9577206b1889) |
| `arg_matchers.item.exact_values` | [arg_matchers.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-1300b9177b3ac3a5f45e7ce67d1e3b68005c27ff10d2cdfcd33e064b6fe654e9) |
| `arg_matchers.item.regex_values` | [arg_matchers.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-e29825c52bfba3ac0eb19d4f7fde81163273a9a7452b758da88da899231610ce) |
| `arg_matchers.item.transformers` | [arg_matchers.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-e561da3359f42b5d8cedf8f9a0a95a76a7fdfc407f6a1c00ba93d07b100be8ab) |
| `arg_matchers.name` | [arg_matchers.name](data-sources--service_policy_rule--reference--group-001.md#canonical-e30eb87ea1fd4929668e3b500cc0dedab68cdfbb337b2466a1998a321922277f) |
| `asn_list` | [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-7aecba9dbac4ee00ab5a73b7cf854c4fab432936e5f8f142643e2e82f3f96471) |
| `asn_list.as_numbers` | [asn_list.as_numbers](data-sources--service_policy_rule--reference--group-001.md#canonical-4a070f48a2f53d91b7413ce80578f57242c892e060741199b34cf80c44920d7d) |
| `asn_matcher` | [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-9913897497010a601864b014048080282dc44d25e56d1bf46013d44e83419468) |
| `asn_matcher.asn_sets` | [asn_matcher.asn_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-8f622b91ca00c53a4f2b0e13d5d8f153c23a6dddc28645b05ba18bf8d15bea2b) |
| `asn_matcher.asn_sets.kind` | [asn_matcher.asn_sets.kind](data-sources--service_policy_rule--reference--group-001.md#canonical-bfd2d3d58430cf88240f5ddc322946910bb90f9c11e348542b268640fa105928) |
| `asn_matcher.asn_sets.name` | [asn_matcher.asn_sets.name](data-sources--service_policy_rule--reference--group-001.md#canonical-729bcae3ee32f3dd518001beb94d1877235cbd255ce8b2bd82edff1566f5f0b5) |
| `asn_matcher.asn_sets.namespace` | [asn_matcher.asn_sets.namespace](data-sources--service_policy_rule--reference--group-001.md#canonical-dacc3dbbd65f7e493ae0acc40b1839afb145715e3364d7db23388595000263e7) |
| `asn_matcher.asn_sets.tenant` | [asn_matcher.asn_sets.tenant](data-sources--service_policy_rule--reference--group-001.md#canonical-b66656d0fea62a64037ad995975ff73e496362dbc80460dc8c1c95f915458ec2) |
| `asn_matcher.asn_sets.uid` | [asn_matcher.asn_sets.uid](data-sources--service_policy_rule--reference--group-001.md#canonical-0e2f6ac72d573b8164f5b4ac4f54cfbadf451f9fa0cdd3b8dcb1bfa8034cc084) |
| `body_matcher` | [body_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3a370b3af8ae5c9420f211f0fd5cc7199538fdd9708b406568a807cb52929b9c) |
| `body_matcher.exact_values` | [body_matcher.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-f32a724c14248bfafc47a88134f995576f169d86f6cefb4a07edffbfd01483f4) |
| `body_matcher.regex_values` | [body_matcher.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-2924009f40bc2b8a016f1d14c58dba7be783cd91b4afa416f09959ca7df39714) |
| `body_matcher.transformers` | [body_matcher.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-ea7aee7bf890cf40ff40e2650a08cc0f070642155f36b2d68bd0182af6f70165) |
| `bot_action` | [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-066eabe423a64de21295922960e696179b25993e080528d6fdf8ffe756bc899c) |
| `bot_action.bot_skip_processing` | [bot_action.bot_skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-5d0a167f4386e274396ff0eb40895aed41177cb16bf8f693210ce65ae3a1a27c) |
| `bot_action.none` | [bot_action.none](data-sources--service_policy_rule--reference--group-001.md#canonical-dd64af9c9c5c817b7c69765e2e6525b3df7404387b5cb14f2d017f13f33ce969) |
| `client_name` | [client_name](data-sources--service_policy_rule--reference--group-001.md#canonical-69dde0972f4803be8c008b5aaca1d11e91820582e9f1dbb4cfddfd54b3793d2c) |
| `client_name_matcher` | [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2e82196dc3b462d34175d5422b92c8dafa268e137093e6f8d76a663bf0c2e146) |
| `client_name_matcher.exact_values` | [client_name_matcher.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-20b5b70aec79de5574c3aa26645511436c1eaeca30c11a3dbc84dd50ad3bd7ab) |
| `client_name_matcher.regex_values` | [client_name_matcher.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-31f577399eda9adb266430d26bc01dab8d925a2b246923e85ed84e59f18009e4) |
| `client_selector` | [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-f8990b3fd7b3cae59cc8ab358a25ecadea2c1119a3b336ac0ecb373f9057e920) |
| `client_selector.expressions` | [client_selector.expressions](data-sources--service_policy_rule--reference--group-001.md#canonical-cef533245f8c252fe441da7af76566087fb3d538bd6ecfc9a533fa28101402c8) |
| `cookie_matchers` | [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-bbc24d0e2e13694ed93f99308a4a94a80e8fdacae5069e9d576cfdbf05256f1e) |
| `cookie_matchers.check_not_present` | [cookie_matchers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-877b8c6edf5751301d7f05733510626cb2488a05e3290302c298a4d0c7eb6601) |
| `cookie_matchers.check_present` | [cookie_matchers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-87d6362cbd50eaa5a5d70ec63bcca4c771b32a846354b20c22bfab15b9692799) |
| `cookie_matchers.invert_matcher` | [cookie_matchers.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-070b18ec3e1b40a2302eb64dab186cd57ec837c98db100da0b946c66e89a577a) |
| `cookie_matchers.item` | [cookie_matchers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-8832dd9b0b03963227b8bc6ea8244981a58e57d7fd2a8335f1f93e008e3bf646) |
| `cookie_matchers.item.exact_values` | [cookie_matchers.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-4eca8b358a12a7dfe755ebaeaf5da86de5e48a9d773f05b41ec60f3fa8694aca) |
| `cookie_matchers.item.regex_values` | [cookie_matchers.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-b9d816fae10396314d5e7338581b2724a37e1e8dee722414c79607bca6613d0e) |
| `cookie_matchers.item.transformers` | [cookie_matchers.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-5787dd0c64f61be099dd7b746776999bf6c0688dbc18aff010e568a847df4536) |
| `cookie_matchers.name` | [cookie_matchers.name](data-sources--service_policy_rule--reference--group-001.md#canonical-8b7bc3bc4b2a702c95986e53016e3e7b06e73c97b03419333770f630467d70e4) |
| `description` | [description](data-sources--service_policy_rule--reference--group-001.md#canonical-41b402c8db3874ad19d5b8f9ae726dee9bc560207884a96ce7b1b6a9a99bf0b9) |
| `domain_matcher` | [domain_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-36b300be0cc4c4c15dda34a7034779a72d633e0985ec653d727d43c5eec5a177) |
| `domain_matcher.exact_values` | [domain_matcher.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-6cabf7a9567a1156a71d24a176045de273c730e784be67ed950e2b494090e1e7) |
| `domain_matcher.regex_values` | [domain_matcher.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-a1a0ead876b7815bb163b7c3feb618ad9fc0c3a48dc547831f41141b1364d155) |
| `expiration_timestamp` | [expiration_timestamp](data-sources--service_policy_rule--reference--group-001.md#canonical-e3f37720e2613a117f6915c19fbfc6df3650fb7243ebe46c97772687c6609cab) |
| `headers` | [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-58c9722a52606797f2e46f83d88fda972cecc271d3d38923c01f1746f9662d10) |
| `headers.check_not_present` | [headers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-a5655798e0bb369c02a5f730c7da57d57bbe9661da7cba898a9da6ae3157600e) |
| `headers.check_present` | [headers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-03f662abf99723a1f087b2512052fdd2ee5d1cd2db5870aa87458ae845016e65) |
| `headers.invert_matcher` | [headers.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-06f9e5f331bdff84a77eea42067b2dfade15741e07b42d14aa2d126ab9154c03) |
| `headers.item` | [headers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-73cd52428eba6ff49c6c5351c520d3771ebea902832ed039975488e0cbd01462) |
| `headers.item.exact_values` | [headers.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-4f07bbb17807ca1f2944b3355d06b15e8088ff7f9f518c80974e1f0941b0904c) |
| `headers.item.regex_values` | [headers.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-0fb2700c93aeba6c4c78275fdb0bb7898ac65375b62dce94af1b85bcd773de05) |
| `headers.item.transformers` | [headers.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-3eacaa5f355ef7fce5997a8c9d0898e21515d2de69513037d1132f08509e0f04) |
| `headers.name` | [headers.name](data-sources--service_policy_rule--reference--group-001.md#canonical-d1a375df1099812d75f539c6c494f2627436f631053fff751b8e48e4894df918) |
| `http_method` | [http_method](data-sources--service_policy_rule--reference--group-001.md#canonical-f19c8dd827a26b5d469c865b20c30c0f09379c75b7882f8eabd8d6f5b134d6ba) |
| `http_method.invert_matcher` | [http_method.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-88d2303fc746cc8b12212f262be8e3985478d2d5e7fbfe84087815eb9a95c6a4) |
| `http_method.methods` | [http_method.methods](data-sources--service_policy_rule--reference--group-001.md#canonical-28842e6238d0a0bdf2c7f82da648fc8da8c4108d3e18d60bc27e82aeae53cf5d) |
| `id` | [id](data-sources--service_policy_rule--reference--group-001.md#canonical-d75958381fd77af37223a4d2d5c8a6d955a983062f91918c3b17ceab9812c587) |
| `ip_matcher` | [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-f2fe1d8d8a1cc10be12ab080a847dc9147eb8a1445ba7229336ebaa3da4c08f9) |
| `ip_matcher.invert_matcher` | [ip_matcher.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-4b77d6bd7758aae142ba6b6a834889cbea71caadb5a86704952f0e1200c99e0e) |
| `ip_matcher.prefix_sets` | [ip_matcher.prefix_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-4b3de46b55d3bb10bb709e635a1272ef715d628cf34f04e179cd76cfdadb8537) |
| `ip_matcher.prefix_sets.kind` | [ip_matcher.prefix_sets.kind](data-sources--service_policy_rule--reference--group-001.md#canonical-83cf436e6c0c0207961d7f5a0f8550108d98eaf39d5bd5fb548f0e5fa5dd6d27) |
| `ip_matcher.prefix_sets.name` | [ip_matcher.prefix_sets.name](data-sources--service_policy_rule--reference--group-001.md#canonical-2be3e9c55a91f54d75f5d2943b976d89f90ba176b4e002302d9750439e2c2d39) |
| `ip_matcher.prefix_sets.namespace` | [ip_matcher.prefix_sets.namespace](data-sources--service_policy_rule--reference--group-001.md#canonical-d20e03d868e9ab8bc4f245cf059c776f558c3aa989746acf99dad7eb33dcd397) |
| `ip_matcher.prefix_sets.tenant` | [ip_matcher.prefix_sets.tenant](data-sources--service_policy_rule--reference--group-001.md#canonical-82cebf422b341790bdb3f4fe06d68e4cb1ed085db23d164b1a4433a12422e249) |
| `ip_matcher.prefix_sets.uid` | [ip_matcher.prefix_sets.uid](data-sources--service_policy_rule--reference--group-001.md#canonical-db35f0c455c0015813f4595d1137af1282fa1af17edf5406d67b46bcc51ed6ad) |
| `ip_prefix_list` | [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-6696dbd25074da043944f7708ee1d094c2ee34bd0224758b8f8dee121b1bf2d0) |
| `ip_prefix_list.invert_match` | [ip_prefix_list.invert_match](data-sources--service_policy_rule--reference--group-001.md#canonical-a55e594b0d8b742ece741e845b2d01486306893e68f0cd4e62dfd8e77ecea14b) |
| `ip_prefix_list.ip_prefixes` | [ip_prefix_list.ip_prefixes](data-sources--service_policy_rule--reference--group-001.md#canonical-8e35c01c375275071c0ded269d2f186fcf44fc2e7bd0cc5c309e714a90b1736c) |
| `ip_threat_category_list` | [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-2edf112999a7fee20a7b6b204fd3bffb224f519efb9ca33941c1d2081a5803c3) |
| `ip_threat_category_list.ip_threat_categories` | [ip_threat_category_list.ip_threat_categories](data-sources--service_policy_rule--reference--group-001.md#canonical-0fae64cc3b0ac17f5c135d044c4b02e1ded8a61441ccb8e9206f5b121885217b) |
| `ja4_tls_fingerprint` | [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-d149fcd44be04087577cadf3b6123f95d9aee9f7342b68300a02e8277f5ef8e8) |
| `ja4_tls_fingerprint.exact_values` | [ja4_tls_fingerprint.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-36c99c5aaf1a105c34a70f7f198ced0a277128e6a677dd14d1893faedaec8ed1) |
| `jwt_claims` | [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-903651741a070ccaad3450637b0947ee47976ef53aca2e7fcda23d4dae1608ab) |
| `jwt_claims.check_not_present` | [jwt_claims.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-5865f83fa6374d4b02f40aad04f3ab4b089e4099ddf2d6756c308376c9c50814) |
| `jwt_claims.check_present` | [jwt_claims.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-7da7d5525238a9802e03b553405c631bf47ceddd9021b227f126fd144ecd1449) |
| `jwt_claims.invert_matcher` | [jwt_claims.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-6cf777cf61ccdf046266b907d6320ba044709822cfc38bd6786d2735f3b19885) |
| `jwt_claims.item` | [jwt_claims.item](data-sources--service_policy_rule--reference--group-001.md#canonical-65222b06e80f122fe3d8185fc13bbdb554354681ef2cd72ead7ef7f5d22f266c) |
| `jwt_claims.item.exact_values` | [jwt_claims.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-9fc7876d081525a785694fa2bf4c7208f1d70e86d33a19dd6f16b0ac9a21005b) |
| `jwt_claims.item.regex_values` | [jwt_claims.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-7e3db97c0aa0cc035933f6aa1ea7ef8892396e13f33741e4fb3048cadf4e23b6) |
| `jwt_claims.item.transformers` | [jwt_claims.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-51061f12207b14b41b17370ba64640dab9a226454137a72edbc72e6e33dc66cb) |
| `jwt_claims.name` | [jwt_claims.name](data-sources--service_policy_rule--reference--group-001.md#canonical-a592b14eb0f688210a25c22326875194ba6f273b953d9cff11f5783286000fce) |
| `label_matcher` | [label_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-20f34c8bf0cdf7c59b22cdd17daf333f7324a91cb4a87384b89411765cc65850) |
| `label_matcher.keys` | [label_matcher.keys](data-sources--service_policy_rule--reference--group-001.md#canonical-41700beadad031f073a63a1b7b43f8edb0d43f37e8c739800aa26b84ca01bf9e) |
| `labels` | [labels](data-sources--service_policy_rule--reference--group-001.md#canonical-2e0e5113e680e19c0f4e8c544b25e9cf84ce9ecfce899548da5d27b3c03430d9) |
| `log_rule_evaluation` | [log_rule_evaluation](data-sources--service_policy_rule--reference--group-001.md#canonical-51b39c1b7d26a64aa16eb836fbfe24076aa17e4752d0b95cedb31d56b1a52a3d) |
| `mum_action` | [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-6eb35d8e99f65232273f6c3f06bac3e5e03c321cb6798033257d0d9856cb317a) |
| `mum_action.default` | [mum_action.default](data-sources--service_policy_rule--reference--group-002.md#canonical-cfa1b83a8d5be05b8877e6dd9e43418e3c6a533a567ba11f1b3cc5c44148bca2) |
| `mum_action.skip_processing` | [mum_action.skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-10d9192d32a10830c275f5b628986ef6109007f710271b8a25aac62e0bfcb4f0) |
| `name` | [name](data-sources--service_policy_rule--reference--group-001.md#canonical-f61f250350b7a614efbdbce9c5c6ed8651b58798805698717fd523a323ad0848) |
| `namespace` | [namespace](data-sources--service_policy_rule--reference--group-001.md#canonical-dcaf02c58a17dc42a260bbf07c8ab46a936b09c98c6130ee0464482a8165ea3d) |
| `path` | [path](data-sources--service_policy_rule--reference--group-002.md#canonical-0d9e1e172dd195191983d5fe048baba005f96c513aae9ae293164121272c89ab) |
| `path.encoded_path_matcher` | [path.encoded_path_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-af553a240a535cc24966fa8d4e17937ac5f930e70455a89a23a1ecbd1e985a12) |
| `path.exact_values` | [path.exact_values](data-sources--service_policy_rule--reference--group-002.md#canonical-4959f2b757808631d6ed53b96654146b6bf607ba9a3780990dd670f6d2cf6f25) |
| `path.invert_matcher` | [path.invert_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-1a0aa613b0d6d05575147f6ca65dfc4ca09a4831e07e8b1e0ea771d582abf9b7) |
| `path.prefix_values` | [path.prefix_values](data-sources--service_policy_rule--reference--group-002.md#canonical-fba100c8bff57663f468ee1f0a15e6d61a6bdddeb1a00e9268a259adcdf31826) |
| `path.regex_values` | [path.regex_values](data-sources--service_policy_rule--reference--group-002.md#canonical-b5ddf173fed2571c218fc8a1700acba502054c43bb92f04ee10b16a21980034c) |
| `path.suffix_values` | [path.suffix_values](data-sources--service_policy_rule--reference--group-002.md#canonical-6ab019f53aa7cf0987168cfa0251ca752904544aacb40f907228da318a94ca95) |
| `path.transformers` | [path.transformers](data-sources--service_policy_rule--reference--group-002.md#canonical-1b4a2c22edd27c91ec14796e33888b37742378dd1c01ee36e1a38020ed80aa62) |
| `port_matcher` | [port_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-1d9c6a10ef27afc42e86601a626d30d91f085a40c8de2dc3734f3d1205357c56) |
| `port_matcher.invert_matcher` | [port_matcher.invert_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-fcfb0321a3f2f2c90239e0488687ce6348f6c5256a4cc69cb9da530fdd4d6d17) |
| `port_matcher.ports` | [port_matcher.ports](data-sources--service_policy_rule--reference--group-002.md#canonical-d1712ff694536f28c86c57ff2f5009fa6ad7c95b11ceed181a647a57c80b6594) |
| `query_params` | [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-142ab1ad452aa2896d9e4246882edda7180e3353549888197881edb7741c9d70) |
| `query_params.check_not_present` | [query_params.check_not_present](data-sources--service_policy_rule--reference--group-002.md#canonical-c54636e470c24448f9738d61f4b2c77ea7c9403b1a86931785efed68c0c16016) |
| `query_params.check_present` | [query_params.check_present](data-sources--service_policy_rule--reference--group-002.md#canonical-42d9f7fb2bdc3b01fd83f3427e644069866e02a5a5842ebdebbbf77873753603) |
| `query_params.invert_matcher` | [query_params.invert_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-533e091ae19e5a73ac1ffd866a36368a8a471adabc725e0089b2abd4b179bf59) |
| `query_params.item` | [query_params.item](data-sources--service_policy_rule--reference--group-002.md#canonical-3c45663aff7da99fbe6894ef7a402d5affddb55e5f2c7722299f1ccf92e40d67) |
| `query_params.item.exact_values` | [query_params.item.exact_values](data-sources--service_policy_rule--reference--group-002.md#canonical-2f605ba6bd5ee209fafc27ae5f24126ce4586487a6478c7f58a86deeae22f958) |
| `query_params.item.regex_values` | [query_params.item.regex_values](data-sources--service_policy_rule--reference--group-002.md#canonical-5decf1709b45fb79e9705dd3deba8cf8fd75981d02dcec298fe02fc67ddc1500) |
| `query_params.item.transformers` | [query_params.item.transformers](data-sources--service_policy_rule--reference--group-002.md#canonical-30e292a8c07688c7539ea7a498c4ae0b9b0ff3d6403966d10056ac0b410826b7) |
| `query_params.key` | [query_params.key](data-sources--service_policy_rule--reference--group-002.md#canonical-366485fae3ff67f7634f5e3fd94a9bcee1714d0000d6e37fc2eee959cbfdda38) |
| `request_constraints` | [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-19cdce4084ec9e481c5aee84aaa5b7c1ce661ee7646de0cfaf90c9518d616908) |
| `request_constraints.max_cookie_count_exceeds` | [request_constraints.max_cookie_count_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-9fb38ae8576c8b902f150d75a8acfb9fd577beedeeea4fabbe9f6221ee796d4a) |
| `request_constraints.max_cookie_count_none` | [request_constraints.max_cookie_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-4c5c8290d478f861c6e9274d8341bd424741e31e074ccc9c0a6d8af9a3227411) |
| `request_constraints.max_cookie_key_size_exceeds` | [request_constraints.max_cookie_key_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-73b5d0584160d52e531082eefd32bf2b7d78ae8b3ae9b5e20cc2e31fc3dc9607) |
| `request_constraints.max_cookie_key_size_none` | [request_constraints.max_cookie_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-4fffe56ac3c8ebff2511e3edf40f5c43e63d1c734e6b3996a5c6257e4d60f6b7) |
| `request_constraints.max_cookie_value_size_exceeds` | [request_constraints.max_cookie_value_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-13425e16f3cd32e029497732af6844291ec8e08e1eb914a925d7888dd2497057) |
| `request_constraints.max_cookie_value_size_none` | [request_constraints.max_cookie_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-24ed5c31551833027a28e37bc96ad20f29609f3bd9b544e061b4dcab89b91dbe) |
| `request_constraints.max_header_count_exceeds` | [request_constraints.max_header_count_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-45e3c7cb468594e1282141c0535dbe8933f4f5462aa54af019c1b1d4c6be54f1) |
| `request_constraints.max_header_count_none` | [request_constraints.max_header_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-4582e8d8582ea6dcc6f90fd72332e04da378adbbab77edc7b46a4c6ce5e4c2f0) |
| `request_constraints.max_header_key_size_exceeds` | [request_constraints.max_header_key_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-09964672a0c2420d458cd8c052d151daa4c1acdd79ebc662516f319b44122bb4) |
| `request_constraints.max_header_key_size_none` | [request_constraints.max_header_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-11ad84ec4196fc7b565f85d763cfa28185735048de0c291dbdd7e6265c555f42) |
| `request_constraints.max_header_value_size_exceeds` | [request_constraints.max_header_value_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-c810382529612d11205d678ae83a45a470dae42b36a2cccfd456bde22f03b29a) |
| `request_constraints.max_header_value_size_none` | [request_constraints.max_header_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-e8aaee10c6b78b111b952d4571899cbb07d4ac30685f2b3e3c91625871d00a50) |
| `request_constraints.max_parameter_count_exceeds` | [request_constraints.max_parameter_count_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-ac905d5de84f54b0d6883f9bb78215eac55464a2be12cb6b76c2d166b2514aae) |
| `request_constraints.max_parameter_count_none` | [request_constraints.max_parameter_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-70e74487b6feb26f29cb705a301c2079cbba978a20934b647b09ef798a21aaa6) |
| `request_constraints.max_parameter_name_size_exceeds` | [request_constraints.max_parameter_name_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-8566e4a7d5756e611a8092222e8e14e565c4ffd9610e739b613ed7a8ec9db56b) |
| `request_constraints.max_parameter_name_size_none` | [request_constraints.max_parameter_name_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-76e894999384d425ec37951ad2709ff510f193949ec6a3bea727e113915700c3) |
| `request_constraints.max_parameter_value_size_exceeds` | [request_constraints.max_parameter_value_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-62c234d54a3885120045d0087bb7515421719197c1e616d2fe1b36836d63f930) |
| `request_constraints.max_parameter_value_size_none` | [request_constraints.max_parameter_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-82fb4e4c80c9738af9561477fe2c1ca0223f9497c5f12c0d3ce99042333a1ac3) |
| `request_constraints.max_query_size_exceeds` | [request_constraints.max_query_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-955a5b60d99906a4a339a72e7d2ae550b3a24bf27862315ffeaa7dd498ff8859) |
| `request_constraints.max_query_size_none` | [request_constraints.max_query_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-7604655ac5aaded637ce980adb7e4a28b286559a891b7aaae9f88711ad5e4132) |
| `request_constraints.max_request_line_size_exceeds` | [request_constraints.max_request_line_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-525b735b57ae97eb091e22ad2e82db2e9902393e0c6b0f2737daddd667f86788) |
| `request_constraints.max_request_line_size_none` | [request_constraints.max_request_line_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3bfe1343298b3ea7650058a8934cff5c2c3e7eadca2ba352a247b86d74a7ad4a) |
| `request_constraints.max_request_size_exceeds` | [request_constraints.max_request_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-2b0ad8f0da212e754204168b5f91cec253d66befd7808b08af2723fde005f337) |
| `request_constraints.max_request_size_none` | [request_constraints.max_request_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-df35bf462fa8d275ab6e1e0b7fc1a249f3918457ebd16c82032bd6f63b5da096) |
| `request_constraints.max_url_size_exceeds` | [request_constraints.max_url_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-97a09d1ff6c9f0e998aae239212ede67be44d8e6bc653ceda3b0a3c9aa30b352) |
| `request_constraints.max_url_size_none` | [request_constraints.max_url_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-adfa33dc9c306e9a27f26e0a2e6f870d3f2999f79c81e788ca971f32ece4fe46) |
| `segment_policy` | [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-8d2067c794f0fb372809242bc002bd7197e79d94368e3cbd68cb115e886ed23b) |
| `segment_policy.dst_any` | [segment_policy.dst_any](data-sources--service_policy_rule--reference--group-002.md#canonical-f738d21847f22a4f9473da78950744d65a1a24132f601e9ee7039c16ebca6eca) |
| `segment_policy.dst_segments` | [segment_policy.dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-3266bf29bb4c0787465d1e2ebccada5d79ea39ac688650129674f8272cb4da6f) |
| `segment_policy.dst_segments.segments` | [segment_policy.dst_segments.segments](data-sources--service_policy_rule--reference--group-002.md#canonical-eeba5610d8b3350152c3a4b1ca2f3d118883a2bdb78309bfd30971d41e7ef7b6) |
| `segment_policy.dst_segments.segments.name` | [segment_policy.dst_segments.segments.name](data-sources--service_policy_rule--reference--group-002.md#canonical-c0fb29a12b2ab403e98ec6f2257bd5fbd9f4815cb27b06b996e829f7b26053df) |
| `segment_policy.dst_segments.segments.namespace` | [segment_policy.dst_segments.segments.namespace](data-sources--service_policy_rule--reference--group-002.md#canonical-20f27ac642e5778a96d471c9a7a277f1b7b44197897ecd5a62e0c82a736eb7f2) |
| `segment_policy.dst_segments.segments.tenant` | [segment_policy.dst_segments.segments.tenant](data-sources--service_policy_rule--reference--group-002.md#canonical-955c63f9b19e2637316255b522a969a6cdf8d51a5fd577c596386b30156daf98) |
| `segment_policy.intra_segment` | [segment_policy.intra_segment](data-sources--service_policy_rule--reference--group-002.md#canonical-d691d7584225f3718c19bc4c3af89da188b208da93dac6b295c5ae4e8dd1fd18) |
| `segment_policy.src_any` | [segment_policy.src_any](data-sources--service_policy_rule--reference--group-002.md#canonical-1a7a905bffddd3083ae6f5aca9c7891b46b4ad21d672b07d9157898878743fe6) |
| `segment_policy.src_segments` | [segment_policy.src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-7da2a0d339edabaf6811c80325bbb654d797964223e30d53ca16584b98d8e600) |
| `segment_policy.src_segments.segments` | [segment_policy.src_segments.segments](data-sources--service_policy_rule--reference--group-002.md#canonical-9a8bef16434e5720e50a05cefed43a3ad5fcedf0660925329159121d8941544b) |
| `segment_policy.src_segments.segments.name` | [segment_policy.src_segments.segments.name](data-sources--service_policy_rule--reference--group-002.md#canonical-338fea2490dd983c3f135abadcad4096f260055f4f120e8704f27db13f51e480) |
| `segment_policy.src_segments.segments.namespace` | [segment_policy.src_segments.segments.namespace](data-sources--service_policy_rule--reference--group-002.md#canonical-e86bd035467aa796d0fc7bdf5ca8fb74c90be0c80df598be07751ca518818a6d) |
| `segment_policy.src_segments.segments.tenant` | [segment_policy.src_segments.segments.tenant](data-sources--service_policy_rule--reference--group-002.md#canonical-908e58bf02ce5928f3c98b0b9aa523b4fae0c68fcdc087a74dd24d8c17ae3a32) |
| `tls_fingerprint_matcher` | [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-e44b775e4590083f8137f4b607de5d213b8808fc5179cdcd48fa9c663d5757c4) |
| `tls_fingerprint_matcher.classes` | [tls_fingerprint_matcher.classes](data-sources--service_policy_rule--reference--group-002.md#canonical-61d227bf85dccc3d5524ab6730e0f6e041ddc550221193aa3262e5e3257a3af8) |
| `tls_fingerprint_matcher.exact_values` | [tls_fingerprint_matcher.exact_values](data-sources--service_policy_rule--reference--group-002.md#canonical-15c34b3a37360974a324c488a89a53e199ec1e0a77c02195a2afb9dfbb2cee4b) |
| `tls_fingerprint_matcher.excluded_values` | [tls_fingerprint_matcher.excluded_values](data-sources--service_policy_rule--reference--group-002.md#canonical-b5d47cca86cf9a0256983ed2cb4b2c297534ace225976ee9bb7830211ebbe703) |
| `waf_action` | [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-68de118d4d4295e6bd3e781f066834486e0cf06b2f86b5b171d4d96966f991a2) |
| `waf_action.app_firewall_detection_control` | [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-bdd79412cece19057c6817382815fccd8e22e429a614569561109894560c72ba) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-94fc1e8fe6f1ed8bbc2188ad19a57df311e5522c4aec2368ca11dd0d7d5e3750) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](data-sources--service_policy_rule--reference--group-002.md#canonical-97d29978a70fe98644d4c2dadd31fff38b265d117f87ad871b0fc4fd1b051e30) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](data-sources--service_policy_rule--reference--group-002.md#canonical-f0085ec310937b8154f59e4dadadcbdbedd53f62125f8d61192d240967d6b9e9) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](data-sources--service_policy_rule--reference--group-002.md#canonical-653d9be481f1429466191edd6d0bbc117ec72f034c1dabd0810803ae66c93e26) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-7b18961e4d7afed2e51c84f996030ba68bd99ed32de86d93e5dc1ed728f10956) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](data-sources--service_policy_rule--reference--group-002.md#canonical-6e90e64d925b28fe35f4ab40638e5d6603a2b1a7aaad76aa008edf201f69b211) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts` | [waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-3ac0ce87875fe5c1e0549c08721a6d43620a4a184497953d826fe0e0c3191cac) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context](data-sources--service_policy_rule--reference--group-002.md#canonical-6dc9ad7a01aa8f6b52391f8f3cd84d004f8c11afb7d17a3129d8d1cb198c7da3) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](data-sources--service_policy_rule--reference--group-002.md#canonical-2324e7e3f052db760a47a873ad5849848e3f8fcb1ff95c96933c7c22e40fec47) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](data-sources--service_policy_rule--reference--group-002.md#canonical-2e3c8e1fc87986f39302b33fb6def2aa8b372634ed9f52332ec68ebc80184b38) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts` | [waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-c3150a67ff750387b960d408c13011d9c3ae274d7b8ec8356ff14c109c371bf9) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context](data-sources--service_policy_rule--reference--group-002.md#canonical-9b8f700d604dd0844c7837769f255de745bfd4fd1d0d7498388fd3a8e2764546) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](data-sources--service_policy_rule--reference--group-002.md#canonical-5bec8ad22a301eab8017f9feab3b198ab0d527ce5065f1ee7d21683641d1116f) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](data-sources--service_policy_rule--reference--group-002.md#canonical-b9dcdef566f530c384244ccba5d017b63393e0aec5ae131a2b4e30f804a265b8) |
| `waf_action.none` | [waf_action.none](data-sources--service_policy_rule--reference--group-002.md#canonical-7619e066b5db2aad474f45c3a82cf891a5c7f3b6f2a9d080b91beba3337b4ffa) |
| `waf_action.waf_skip_processing` | [waf_action.waf_skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-6a9db06d2e5f8785dc172f9af1937e7cfbe347a6da7c5b1788d9056d4a367785) |

<a id="canonical-f709f7dfc6eb5811676a74a67f9645505342f33427a8809b3f32761c82dc9bb2"></a>

## Next pages — Property reference / dc726d2ff366 / 15

- [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-c9fadc1947ea10969e9fda773fb6c34b205c8f4e3f8d9acfe60230a9aee5b48c)
- [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-a187c7105b0492c8f3bfade3fd49556f779f9196fe2f4b923fc6d8a097555b28)
- [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-84ebdc2357bec6ce8b16bbbc3baa8b4017493006464f42c495020e236eae1fb1)
- [api_group_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-bba1340cb0a6a7f136e7ead6bcb536cf8e805788d0033f246ac5769cb6039510)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6)
- [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-bb506919f0604945c9733e23725137ca93905f91f5750bb30a77adc2a31102d6)
- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-9b711c97a0ab1fdd142a22047f397fbfdbe65db43aa3a3e01006f75953bcca68)
- [body_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0b0bdac8179833c9adc7559d4d46d8c3073bbb200c6aae22807622a45623baa4)
- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-60950b56578b9a3d2e7f01afbf2f852364b3567038338171b36ef07220a513b8)
- [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3fc5a45dd8f5f55acc9df34955eb8827a3995b6449583a84c5305f74f2e29351)
- [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-5bf025ca77ff33756a6188586d3d1dc585b82381fca91461ce07475b6b7719bf)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae)
- [domain_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-fa39c24dd24d6084f6b5f364ba1908e727261ce5b78f84e14d6f55bb96699695)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8)
- [http_method](data-sources--service_policy_rule--reference--group-001.md#canonical-2ccfac68219f29dec73d5b10667fdddc8a3d5f24c54e6fdbd705961ddb8382bd)
- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-4c64fdc40e1c3b6cadd843f10a9fcacaa6d87d0a051cb31ababde925f2d339dc)
- [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-f568569c6da6e1d479ee5ca4c660b0f3b3a6e5bdd4c77d6f54811e6a1ffd3098)
- [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-e2480d86388e5c9fb566e99804e4614053a7dba6f8db96780101bd841aaf01f7)
- [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-fb9290901816a8a4f9e84e612d576c12566a1b4d776627310a573f99ef9fcfc3)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36)
- [label_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-e4d40e388ca9969028510e765949d2ecf07e9d65a0c3bd535e05cee68412ce0b)
- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-137e55d694ff101b1eceed4b0f2170fade6d8b7a600a590d2a83bb1efb5d0eb3)
- [path](data-sources--service_policy_rule--reference--group-002.md#canonical-62ecded1c4e84948a11dba301846deb27ceaebac027c0ed9896d35e68b2ae3c7)
- [port_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-91908ed3f38a27795fb956d472ae83414d3046dc441e006972b9304b7899f358)
- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-1fc6b6f8ae22871967881106ea255eba23defd48f77829156aa2fdf56202e7ca)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-c9fadc1947ea10969e9fda773fb6c34b205c8f4e3f8d9acfe60230a9aee5b48c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14aae7b8f1c698851fc1761cf746f236c929e5f4afd00383f70304421e47d0d0"></a>

## any_asn — any_asn / 54628aa89a84 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- any_asn

<a id="canonical-53efe7095c1382b98b12fbcb16d318d032ad0782a66079042f357a2be7eb1749"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_asn, asn\_list, asn\_matcher\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-53efe7095c1382b98b12fbcb16d318d032ad0782a66079042f357a2be7eb1749)
- [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-7aecba9dbac4ee00ab5a73b7cf854c4fab432936e5f8f142643e2e82f3f96471)
- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-9913897497010a601864b014048080282dc44d25e56d1bf46013d44e83419468)

Select alternatives according to the provider validators above.

<a id="canonical-308fbd167ef78788cc974a0fb58a4858cea71071f460aa35c694687cdf5d2268"></a>

## Direct properties — any_asn / 54628aa89a84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-27a20d67f626630fb00713c1a48ac4721eaed41770330ef5928499b10467682f"></a>

## Next pages — any_asn / 54628aa89a84 / 4

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-a187c7105b0492c8f3bfade3fd49556f779f9196fe2f4b923fc6d8a097555b28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40f738c9e6302c3932aa3eefff233c4ef9535fecfbef25778c5f0fa3704d21ed"></a>

## any_client — any_client / b703e0e1aeaa / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- any_client

<a id="canonical-aefd0b7d81871a8e4bd8bf2a976341dcbd6ccf876494361f2d155185914ea3e9"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_client, client\_name, client\_name\_matcher, client\_selector,
ip\_threat\_category\_list\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-aefd0b7d81871a8e4bd8bf2a976341dcbd6ccf876494361f2d155185914ea3e9)
- [client_name](data-sources--service_policy_rule--reference--group-001.md#canonical-69dde0972f4803be8c008b5aaca1d11e91820582e9f1dbb4cfddfd54b3793d2c)
- [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2e82196dc3b462d34175d5422b92c8dafa268e137093e6f8d76a663bf0c2e146)
- [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-f8990b3fd7b3cae59cc8ab358a25ecadea2c1119a3b336ac0ecb373f9057e920)
- [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-2edf112999a7fee20a7b6b204fd3bffb224f519efb9ca33941c1d2081a5803c3)

Select alternatives according to the provider validators above.

<a id="canonical-459a283958ee48f0d490523df9baeaf03f0814d867ca233ac8db478aa86d134d"></a>

## Direct properties — any_client / b703e0e1aeaa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb7acae2cca21b7bd7eebb39f9eae953cf8a11afb3306f16a822dcf2d5c5e015"></a>

## Next pages — any_client / b703e0e1aeaa / 4

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-84ebdc2357bec6ce8b16bbbc3baa8b4017493006464f42c495020e236eae1fb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c27649553a8323b32414a92ff7f439f83cbb25a6628dd4ff7fb0e34f900d04b"></a>

## any_ip — any_ip / 3c302c3e3163 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- any_ip

<a id="canonical-37e2c167720c140231d0639a294bee12696a1f423fd3182aecb347bf93fb8a73"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_ip, ip\_matcher, ip\_prefix\_list\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-37e2c167720c140231d0639a294bee12696a1f423fd3182aecb347bf93fb8a73)
- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-f2fe1d8d8a1cc10be12ab080a847dc9147eb8a1445ba7229336ebaa3da4c08f9)
- [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-6696dbd25074da043944f7708ee1d094c2ee34bd0224758b8f8dee121b1bf2d0)

Select alternatives according to the provider validators above.

<a id="canonical-57fed7607ffb3693ba85baf99bdc11a58113cfbf4f98534cae24b6ce0fd68fd6"></a>

## Direct properties — any_ip / 3c302c3e3163 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59c2581a58f6b52195e03ab21829ede2877f617a7960535af4497eafa32b40e1"></a>

## Next pages — any_ip / 3c302c3e3163 / 4

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-bba1340cb0a6a7f136e7ead6bcb536cf8e805788d0033f246ac5769cb6039510"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3592bd931f7f7d0f6ad08ae060ac2240d97bff63ad77d1048fbbf81ee7dd8c0d"></a>

## api_group_matcher — api_group_matcher / 18c980e04b52 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- api_group_matcher

<a id="canonical-78a06280537d7f18fc43e762776e979f401ed05aad3218144634670e40a60431"></a>

Type: `"single"`. Computed.

Matcher specifies a list of values for matching an input string. The match is considered successful
if the input value is present in the list. The result of the match is inverted if invert\_matcher is
true.

Upstream description:

A matcher specifies a list of values for matching an input string. The match is considered
successful if the input value is present in the list. The result of the match is inverted if
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

<a id="canonical-3582e8cee6e055db39443a39fc0d769d17ded10b447db89b09e890294038ec22"></a>

## Direct properties — api_group_matcher / 18c980e04b52 / 3

<a id="canonical-598426c9b7094e4909395205744aae9850f52bbc8e9bd64e692d607f9774b04a"></a>

<a id="canonical-2631ef5410806819c94aff76fbf4f25a815c326d111f5482e0d0600331920a56"></a>

## invert_matcher property — api_group_matcher / 18c980e04b52 / 4

Type: `"bool"`. Computed.

Invert String Matcher. Invert the match result.

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

<a id="canonical-4256bd8214ddc9409267603e515c1144b564800c084c3f9a4c4ed2f744764582"></a>

<a id="canonical-26b65604a8fda0056e1e91c0bbf294a491695803daf9017ed2b66e28ae59057a"></a>

## match property — api_group_matcher / 18c980e04b52 / 5

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-09a3d96693cad246c3898a96983a887158362fcedade926326c609f879734a3b"></a>

## Next pages — api_group_matcher / 18c980e04b52 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f94ea33e76acf820cb1b2aac1f42564ed0c2f9450833d991a2a10584bfda5900"></a>

## arg_matchers — arg_matchers / 7365bc5c8140 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- arg_matchers

<a id="canonical-8b997cb3a29a65bb0850c07e6dca1d4245242ef8b30e1f6e0483b117927bfada"></a>

Type: `"list"`. Computed.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-b1c80025d3823191656c365d5604e948b1138e3ffdecd53019f08bbd21e37c91"></a>

## Direct properties — arg_matchers / 7365bc5c8140 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-5bcf1e717e49a323139c96d65a29e9b0807dac0145c39879600613154fa26a45): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-baef8dea2ff694dd494aa52204f5db5a42d2fce7089bd44afabecf7cbb0405a1): complete subsection reference.

<a id="canonical-7cc0d00a655443f6428fcf6f78bc4e2ba2facaaf54eeb96a8404317c5915eb24"></a>

<a id="canonical-72f8a8feafd597fb33a3a879072dbf22d82218d04225daab9f2235207af2517f"></a>

## invert_matcher property — arg_matchers / 7365bc5c8140 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-be11fcec395ba9c9c9a849efa4511dadf919f0e593e8e200e48ee3ee94c4df49): complete subsection reference.

<a id="canonical-e30eb87ea1fd4929668e3b500cc0dedab68cdfbb337b2466a1998a321922277f"></a>

<a id="canonical-f250b7677f9cd3e6374ba6e23bfab8f62122912f41ac90cb7b22a2a79b3df1c0"></a>

## name property — arg_matchers / 7365bc5c8140 / 5

Type: `"string"`. Computed.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
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
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-5f46979e547cec050d95f08c6f141d011053ef9550524fcfad86e9f5d51dd64b"></a>

## Next pages — arg_matchers / 7365bc5c8140 / 6

- [arg_matchers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-5bcf1e717e49a323139c96d65a29e9b0807dac0145c39879600613154fa26a45)
- [arg_matchers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-baef8dea2ff694dd494aa52204f5db5a42d2fce7089bd44afabecf7cbb0405a1)
- [arg_matchers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-be11fcec395ba9c9c9a849efa4511dadf919f0e593e8e200e48ee3ee94c4df49)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-5bcf1e717e49a323139c96d65a29e9b0807dac0145c39879600613154fa26a45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3b3f231b69fa5a50faf09e3eb14addd73d83478ce76c00636daa7601d53ef4b"></a>

## arg_matchers.check_not_present — arg_matchers.check_not_present / 11add8f5ab02 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6)
- arg_matchers.check_not_present

<a id="canonical-8ad3893bf2dba3208ac168ced1b81864c9edd6dc570c193a8a48179500f89d77"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-9c03933e436eacca3fb8a719af22ad319a7c730700e72519e3f6280341cca476"></a>

## Direct properties — arg_matchers.check_not_present / 11add8f5ab02 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07485cce39ba267d99d33bbcf8cc9c4e135e54da03d6d9f894bd993326cd4bda"></a>

## Next pages — arg_matchers.check_not_present / 11add8f5ab02 / 4

- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-baef8dea2ff694dd494aa52204f5db5a42d2fce7089bd44afabecf7cbb0405a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bdcc1b37cce0903b0ceaa2e2086d98eb565aa8853c859798dd8ce937395c2a7"></a>

## arg_matchers.check_present — arg_matchers.check_present / 7e637ca1ea1d / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6)
- arg_matchers.check_present

<a id="canonical-ffb79fe5bd5a6621289ca633fba77d0726323ec876f807ffaa92ccabe11ce76a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3572c6be31da76847c3854f7f8552dec56c0c73cfc21e562e04011e25896e98a"></a>

## Direct properties — arg_matchers.check_present / 7e637ca1ea1d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eba98448d80194deeb46c6c80ee4dbc0798026cd1ffa3636d71a86e5a7d71c04"></a>

## Next pages — arg_matchers.check_present / 7e637ca1ea1d / 4

- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-be11fcec395ba9c9c9a849efa4511dadf919f0e593e8e200e48ee3ee94c4df49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfd623769d74e9cf7e4ac14de03905c73c484e28913fbe5c515790955d950193"></a>

## arg_matchers.item — arg_matchers.item / c98d67b0b823 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6)
- arg_matchers.item

<a id="canonical-c0e6a9e264fc73074ac774a576f969a6cb382848bf6fc1ffa19b9577206b1889"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-057cbe9bc9e523c41381cfdf5c30208d0100fa2c971f3e735f78e82f3af379ba"></a>

## Direct properties — arg_matchers.item / c98d67b0b823 / 3

<a id="canonical-1300b9177b3ac3a5f45e7ce67d1e3b68005c27ff10d2cdfcd33e064b6fe654e9"></a>

<a id="canonical-16c07e7ea02a007a7ea07449ac9f405dfb15c29d81731a97f023450d4b613e51"></a>

## exact_values property — arg_matchers.item / c98d67b0b823 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e29825c52bfba3ac0eb19d4f7fde81163273a9a7452b758da88da899231610ce"></a>

<a id="canonical-2c29be3acb118aaecc561f6908d9f914bcb5f0c1685517f640139bfd47282c8b"></a>

## regex_values property — arg_matchers.item / c98d67b0b823 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e561da3359f42b5d8cedf8f9a0a95a76a7fdfc407f6a1c00ba93d07b100be8ab"></a>

<a id="canonical-3c6d19c35bd8af68c8635aa7b1e0bb2d4c26424a5862e7a816087121597b92bd"></a>

## transformers property — arg_matchers.item / c98d67b0b823 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-240c8151be2d880f2e81e6eed140e97b658ed8a47d5bc30c6aedda72b143b8ee"></a>

## Next pages — arg_matchers.item / c98d67b0b823 / 7

- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-210c14a67cb5cfedddd8fbbe7353861ff0ea5968dc800e87ef4c8a2b066403d6)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-bb506919f0604945c9733e23725137ca93905f91f5750bb30a77adc2a31102d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82e70ad9b3e4e804fcbf89a197789303a92d2e0c9f5cf2174f23de822dea0c72"></a>

## asn_list — asn_list / 12660233814e / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- asn_list

<a id="canonical-7aecba9dbac4ee00ab5a73b7cf854c4fab432936e5f8f142643e2e82f3f96471"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-6130f76f92ea511b5f43e8213bcc0c52a1d295d57db28810a66ae73a42abffcc"></a>

## Direct properties — asn_list / 12660233814e / 3

<a id="canonical-4a070f48a2f53d91b7413ce80578f57242c892e060741199b34cf80c44920d7d"></a>

<a id="canonical-3e519520efbbbe21ccec54eb6b654e5be1e1db9ae348075c784cc8ee00703048"></a>

## as_numbers property — asn_list / 12660233814e / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-dcc9c854c533f8300101917dd921c3988047587dc38544e1d2f396d59803b78c"></a>

## Next pages — asn_list / 12660233814e / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-9b711c97a0ab1fdd142a22047f397fbfdbe65db43aa3a3e01006f75953bcca68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-324f64736c8f0ba2f6b85f89e5d127e39cfb5311dd11384462d080ea3b6c4a4b"></a>

## asn_matcher — asn_matcher / d2bb545d8e4b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- asn_matcher

<a id="canonical-9913897497010a601864b014048080282dc44d25e56d1bf46013d44e83419468"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-b60ca2de09fa8f567f6d797e2099d08ed1a038d1e776dbecd1a79b201fd3de3a"></a>

## Direct properties — asn_matcher / d2bb545d8e4b / 3

- [asn_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-cda8da68d3d1e605dcaabd2c22be11b03ea178538d2dbee65fa87a5df063158e): complete subsection reference.

<a id="canonical-e222e3042eb7dc702f38e9536ebd2c2426955970aede13d4343afa5e7186ae04"></a>

## Next pages — asn_matcher / d2bb545d8e4b / 4

- [asn_matcher.asn_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-cda8da68d3d1e605dcaabd2c22be11b03ea178538d2dbee65fa87a5df063158e)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-cda8da68d3d1e605dcaabd2c22be11b03ea178538d2dbee65fa87a5df063158e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b2be77cb13d62e2bc19932fd7db938181039d8f8c548ada73b5cdc13e0e0c51"></a>

## asn_matcher.asn_sets — asn_matcher.asn_sets / 93449b3727e4 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-9b711c97a0ab1fdd142a22047f397fbfdbe65db43aa3a3e01006f75953bcca68)
- asn_matcher.asn_sets

<a id="canonical-8f622b91ca00c53a4f2b0e13d5d8f153c23a6dddc28645b05ba18bf8d15bea2b"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-c25c1dbe1889345b1bec610b223863195c6f6264b8355bc704ed6e43a3b45b6b"></a>

## Direct properties — asn_matcher.asn_sets / 93449b3727e4 / 3

<a id="canonical-bfd2d3d58430cf88240f5ddc322946910bb90f9c11e348542b268640fa105928"></a>

<a id="canonical-64d1ca8a0565109bafe8158b76e983f85c5022aaa2825a1fb9ded6beaf55919d"></a>

## kind property — asn_matcher.asn_sets / 93449b3727e4 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

<a id="canonical-729bcae3ee32f3dd518001beb94d1877235cbd255ce8b2bd82edff1566f5f0b5"></a>

<a id="canonical-23ee629e3a1a9f7dcc3fbd0c41ffc74f9cd79e02d84e2369e40f885002c773d0"></a>

## name property — asn_matcher.asn_sets / 93449b3727e4 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

<a id="canonical-dacc3dbbd65f7e493ae0acc40b1839afb145715e3364d7db23388595000263e7"></a>

<a id="canonical-67982c717329cc7f6fbe00a9be6f5ff355ca91244394d604fa27cb70d39a9e7f"></a>

## namespace property — asn_matcher.asn_sets / 93449b3727e4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-b66656d0fea62a64037ad995975ff73e496362dbc80460dc8c1c95f915458ec2"></a>

<a id="canonical-643487d3c4f75bc9f47136d1517ba5debde7dc4f31feea11a088926fe622b711"></a>

## tenant property — asn_matcher.asn_sets / 93449b3727e4 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

<a id="canonical-0e2f6ac72d573b8164f5b4ac4f54cfbadf451f9fa0cdd3b8dcb1bfa8034cc084"></a>

<a id="canonical-794d62165a75d6106974bd9637643fdd48ed1d6d0fc154d80d489f5d8772faa9"></a>

## uid property — asn_matcher.asn_sets / 93449b3727e4 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

<a id="canonical-9e0a955875ab0a3f819bd154dead104f774c2ebf6ec0e6ba08d44e9f87011d00"></a>

## Next pages — asn_matcher.asn_sets / 93449b3727e4 / 9

- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-9b711c97a0ab1fdd142a22047f397fbfdbe65db43aa3a3e01006f75953bcca68)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-0b0bdac8179833c9adc7559d4d46d8c3073bbb200c6aae22807622a45623baa4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d8cd954f2f7268d1ce5ff4ff6ceae68cf44142d9f1c4c38ac80b38dc2227fb1"></a>

## body_matcher — body_matcher / 09ffaaf91e38 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- body_matcher

<a id="canonical-3a370b3af8ae5c9420f211f0fd5cc7199538fdd9708b406568a807cb52929b9c"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-7ae25a16dc7d2de5bc7f8fdca8443e65b11d4e20c6290a21fa17dffa551c5f30"></a>

## Direct properties — body_matcher / 09ffaaf91e38 / 3

<a id="canonical-f32a724c14248bfafc47a88134f995576f169d86f6cefb4a07edffbfd01483f4"></a>

<a id="canonical-ef40b3b2cb4b6f2587edd943c425bc52731db783afbb73914c3a57b9ed40e838"></a>

## exact_values property — body_matcher / 09ffaaf91e38 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2924009f40bc2b8a016f1d14c58dba7be783cd91b4afa416f09959ca7df39714"></a>

<a id="canonical-e213a3c0e4fe372e285b57e839714939862a11fa78f392987d170dbac2fd01c9"></a>

## regex_values property — body_matcher / 09ffaaf91e38 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ea7aee7bf890cf40ff40e2650a08cc0f070642155f36b2d68bd0182af6f70165"></a>

<a id="canonical-2363b09b5cd6d284e8196792b6335c78e66e3dfd40c54275f534574a62ab64cd"></a>

## transformers property — body_matcher / 09ffaaf91e38 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-42ee6323ef8a8a5a180525e99adad729df59a7f8ef4a8b82c539cf811d339ded"></a>

## Next pages — body_matcher / 09ffaaf91e38 / 7

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-60950b56578b9a3d2e7f01afbf2f852364b3567038338171b36ef07220a513b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74fd861e550f0063398b7fffedbb3982916f70f4fdce524d6cb86e81d166e401"></a>

## bot_action — bot_action / acecc9ed5517 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- bot_action

<a id="canonical-066eabe423a64de21295922960e696179b25993e080528d6fdf8ffe756bc899c"></a>

Type: `"single"`. Computed.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

<a id="canonical-86c8d52840bea804d847dba08fc9b948d496867a2aadc9c04c458b6bab7fd285"></a>

## Direct properties — bot_action / acecc9ed5517 / 3

- [bot_skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-dbf022e963a0047fd44e1f4654e3830481eb833d8749afe33f50e5290ffcc264): complete subsection reference.

- [none](data-sources--service_policy_rule--reference--group-001.md#canonical-a753ff5390f76d9de6332b6f899c0fdcd0b2c9f8582dca85bf90afa90496a00f): complete subsection reference.

<a id="canonical-2b72a999dd262dead8dc911cde3084f581a0c7e0b75ed5dc67bf60d53777e5cf"></a>

## Next pages — bot_action / acecc9ed5517 / 4

- [bot_action.bot_skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-dbf022e963a0047fd44e1f4654e3830481eb833d8749afe33f50e5290ffcc264)
- [bot_action.none](data-sources--service_policy_rule--reference--group-001.md#canonical-a753ff5390f76d9de6332b6f899c0fdcd0b2c9f8582dca85bf90afa90496a00f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-dbf022e963a0047fd44e1f4654e3830481eb833d8749afe33f50e5290ffcc264"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-801ae0aea62dee2cfbf2e5ed21a8478f7fd2ee6fc3bc025ddcb3b6a506b69a4f"></a>

## bot_action.bot_skip_processing — bot_action.bot_skip_processing / 84bde1a854d6 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-60950b56578b9a3d2e7f01afbf2f852364b3567038338171b36ef07220a513b8)
- bot_action.bot_skip_processing

<a id="canonical-5d0a167f4386e274396ff0eb40895aed41177cb16bf8f693210ce65ae3a1a27c"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-467897f07749b6b0f360a40d52aa2a6f555e0c1438606465b44c9bf55ae88ba9"></a>

## Direct properties — bot_action.bot_skip_processing / 84bde1a854d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58abe86ad4e871dcaa431e3105b3e5174df235fca35918e76db4affd13da8581"></a>

## Next pages — bot_action.bot_skip_processing / 84bde1a854d6 / 4

- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-60950b56578b9a3d2e7f01afbf2f852364b3567038338171b36ef07220a513b8)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-a753ff5390f76d9de6332b6f899c0fdcd0b2c9f8582dca85bf90afa90496a00f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aaa00a519768742c426b19c4d33a55e8c222cdd8f7a1e3427bfd060f39da0e7"></a>

## bot_action.none — bot_action.none / d9c133acd003 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-60950b56578b9a3d2e7f01afbf2f852364b3567038338171b36ef07220a513b8)
- bot_action.none

<a id="canonical-dd64af9c9c5c817b7c69765e2e6525b3df7404387b5cb14f2d017f13f33ce969"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-6ccd7f128d33a61ab0422859978d65f4ab46bd03f2b652e519f75419dbd6a713"></a>

## Direct properties — bot_action.none / d9c133acd003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-350454386025c935dc425eccbad6e5dc3d96137df7bebbdeff70bd0058beb0a1"></a>

## Next pages — bot_action.none / d9c133acd003 / 4

- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-60950b56578b9a3d2e7f01afbf2f852364b3567038338171b36ef07220a513b8)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-3fc5a45dd8f5f55acc9df34955eb8827a3995b6449583a84c5305f74f2e29351"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba63832adb9dbf1a854620d1d1ff41843af5d7c04874af29dfe100102145ff57"></a>

## client_name_matcher — client_name_matcher / b882c67b7feb / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- client_name_matcher

<a id="canonical-2e82196dc3b462d34175d5422b92c8dafa268e137093e6f8d76a663bf0c2e146"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-b1ce941ddbd5f669d0c93f8e8a5be1a2c63d8e7c5c49be807509e6cf05c40896"></a>

## Direct properties — client_name_matcher / b882c67b7feb / 3

<a id="canonical-20b5b70aec79de5574c3aa26645511436c1eaeca30c11a3dbc84dd50ad3bd7ab"></a>

<a id="canonical-c572ed56ffc490cb74f524981cfa6cccff8623c6e5de0b3a576ad3d5dc83982b"></a>

## exact_values property — client_name_matcher / b882c67b7feb / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-31f577399eda9adb266430d26bc01dab8d925a2b246923e85ed84e59f18009e4"></a>

<a id="canonical-8ba5446b939c865b3fde8bd05ab22588b960f4dbcdecb52568dbef44d9fad0ee"></a>

## regex_values property — client_name_matcher / b882c67b7feb / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-36e34edfc92a644da78dfa6b2f9af22b478ba731719b05b17d1badade3870aa0"></a>

## Next pages — client_name_matcher / b882c67b7feb / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-5bf025ca77ff33756a6188586d3d1dc585b82381fca91461ce07475b6b7719bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4e484fce90f2992b5f57d28e5afa74010601f324ee30858d651384081e882d3"></a>

## client_selector — client_selector / eeb0218782b3 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- client_selector

<a id="canonical-f8990b3fd7b3cae59cc8ab358a25ecadea2c1119a3b336ac0ecb373f9057e920"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-86354703fdede8721ca30573cedde2ec548067df3dc659ecf50c5bb895166b30"></a>

## Direct properties — client_selector / eeb0218782b3 / 3

<a id="canonical-cef533245f8c252fe441da7af76566087fb3d538bd6ecfc9a533fa28101402c8"></a>

<a id="canonical-3fc9341e057a001d93b7fe673b37d758a5a10d41468f55ef335d653bf1682c94"></a>

## expressions property — client_selector / eeb0218782b3 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-f47bafd26f0000e1760b178bc76c964d0eded5ed54b7f2f19d53acf710942e7b"></a>

## Next pages — client_selector / eeb0218782b3 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea3aabfb84c103809164d717abc277cbdf31bca0869e979d574e35a8b5bd9c59"></a>

## cookie_matchers — cookie_matchers / 8858594554f8 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- cookie_matchers

<a id="canonical-bbc24d0e2e13694ed93f99308a4a94a80e8fdacae5069e9d576cfdbf05256f1e"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-94a8aa1724c67e7d1076808c5e9882f1103155bea88beab955b81bc52d959b0e"></a>

## Direct properties — cookie_matchers / 8858594554f8 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2c62621d4f67cd8fd4872621f4bed7e76ab2619dbaec24ad4c420ff2528a339e): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-e5d858deb094650186a73ca51e568378e47561fb9684a84358a5bf546cad2784): complete subsection reference.

<a id="canonical-070b18ec3e1b40a2302eb64dab186cd57ec837c98db100da0b946c66e89a577a"></a>

<a id="canonical-b1a5a9c85171c90f45bbeb64c3d749e79c7cf07e9cef08f047d2fb669743d8d0"></a>

## invert_matcher property — cookie_matchers / 8858594554f8 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-8a2055ee79b30e4f0af93ad02cddef84e6877b6e992925dc2ce21c7682c2fb88): complete subsection reference.

<a id="canonical-8b7bc3bc4b2a702c95986e53016e3e7b06e73c97b03419333770f630467d70e4"></a>

<a id="canonical-4b84f0ac20c2fdbb96a399bd11b768430d010d1eca189ce253cbab1cc49d57f5"></a>

## name property — cookie_matchers / 8858594554f8 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
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
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-64343e3655b0a03ae0626053e43d6dd0bac8d15f80d9f369dea7daf8dc7a4c75"></a>

## Next pages — cookie_matchers / 8858594554f8 / 6

- [cookie_matchers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2c62621d4f67cd8fd4872621f4bed7e76ab2619dbaec24ad4c420ff2528a339e)
- [cookie_matchers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-e5d858deb094650186a73ca51e568378e47561fb9684a84358a5bf546cad2784)
- [cookie_matchers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-8a2055ee79b30e4f0af93ad02cddef84e6877b6e992925dc2ce21c7682c2fb88)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-2c62621d4f67cd8fd4872621f4bed7e76ab2619dbaec24ad4c420ff2528a339e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df03464a8eb4ba0e1c77fda555d908538f9059ce86f7a42a4cb1e6f463625dfa"></a>

## cookie_matchers.check_not_present — cookie_matchers.check_not_present / afa6744d185c / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae)
- cookie_matchers.check_not_present

<a id="canonical-877b8c6edf5751301d7f05733510626cb2488a05e3290302c298a4d0c7eb6601"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-379436536df5277b10b52482e5afa10d46c14e0f9d1bdf189ddd5bb77433727c"></a>

## Direct properties — cookie_matchers.check_not_present / afa6744d185c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce6e1925dc2f1a7ad88eb7ea86ec899dee8bbf3dbdcf3f03f492160682fc755e"></a>

## Next pages — cookie_matchers.check_not_present / afa6744d185c / 4

- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-e5d858deb094650186a73ca51e568378e47561fb9684a84358a5bf546cad2784"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b461037a591bb61db027d276a1daadef5973bebe5628e21f50e850500d6d29d2"></a>

## cookie_matchers.check_present — cookie_matchers.check_present / eb061869331f / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae)
- cookie_matchers.check_present

<a id="canonical-87d6362cbd50eaa5a5d70ec63bcca4c771b32a846354b20c22bfab15b9692799"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-d598ac5c591bb0025874e164494f8138e69ca6fad7be7daa530c59ae7635287a"></a>

## Direct properties — cookie_matchers.check_present / eb061869331f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4b8d8a0fd9c96bac10dca66700ce351c5886a1ef79b85756a89010be9e3b006"></a>

## Next pages — cookie_matchers.check_present / eb061869331f / 4

- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-8a2055ee79b30e4f0af93ad02cddef84e6877b6e992925dc2ce21c7682c2fb88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc28b766a44e55ce5d9ca7c7ba6666a6e139fdef70069586c84f6be4dcde5014"></a>

## cookie_matchers.item — cookie_matchers.item / 2827251d3802 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae)
- cookie_matchers.item

<a id="canonical-8832dd9b0b03963227b8bc6ea8244981a58e57d7fd2a8335f1f93e008e3bf646"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-2002b393891229aa69dab9141c00f9bee5b3a4a03957470ee56db5d648d139b1"></a>

## Direct properties — cookie_matchers.item / 2827251d3802 / 3

<a id="canonical-4eca8b358a12a7dfe755ebaeaf5da86de5e48a9d773f05b41ec60f3fa8694aca"></a>

<a id="canonical-698976e2a02aac9374589622e348bbe699f95fb0ac77e59456d4d6c848bc7b32"></a>

## exact_values property — cookie_matchers.item / 2827251d3802 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b9d816fae10396314d5e7338581b2724a37e1e8dee722414c79607bca6613d0e"></a>

<a id="canonical-4c78be1cc7e7fa03e06e18658aad675f0744f12a942ac944aba549e17f26420e"></a>

## regex_values property — cookie_matchers.item / 2827251d3802 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5787dd0c64f61be099dd7b746776999bf6c0688dbc18aff010e568a847df4536"></a>

<a id="canonical-51ce631211db41fc4b25a6e5d73825b94b912f7f1057f0f1084a9d8aac9030f7"></a>

## transformers property — cookie_matchers.item / 2827251d3802 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-628a0787f472e1006b95ef6ab6d708e890c29415ae3800f605f50a04a20487ab"></a>

## Next pages — cookie_matchers.item / 2827251d3802 / 7

- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-8d8d90141ba5f9e07d3bcdc81bedc61950b5a7955e9c8f51317a999ccf9d3aae)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-fa39c24dd24d6084f6b5f364ba1908e727261ce5b78f84e14d6f55bb96699695"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-377c32282561294826929641162f7c72a5e87a975a9d98828fd15d8022f50a9e"></a>

## domain_matcher — domain_matcher / e7eb81a6436b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- domain_matcher

<a id="canonical-36b300be0cc4c4c15dda34a7034779a72d633e0985ec653d727d43c5eec5a177"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-03d8d12e4c917da4956c85363651b3a10d472d40c1ca9282874a00e525502408"></a>

## Direct properties — domain_matcher / e7eb81a6436b / 3

<a id="canonical-6cabf7a9567a1156a71d24a176045de273c730e784be67ed950e2b494090e1e7"></a>

<a id="canonical-46d73b8a8cc4135a8c65821abf7196f3dcc1261e1e670e03c497d0f97b936c11"></a>

## exact_values property — domain_matcher / e7eb81a6436b / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a1a0ead876b7815bb163b7c3feb618ad9fc0c3a48dc547831f41141b1364d155"></a>

<a id="canonical-de37aea8207293397d2d7cdb2bd2a582940227c4448622646b63bc7096765245"></a>

## regex_values property — domain_matcher / e7eb81a6436b / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e483c48abb65183eb8ac713a9a7e1059d6161619f51bb86dce404693c3ea1cad"></a>

## Next pages — domain_matcher / e7eb81a6436b / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f66727132470d9357290b21854382d2c340300a39ede9c33dad9237edd9bde1"></a>

## headers — headers / 0cf6fb3656d9 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- headers

<a id="canonical-58c9722a52606797f2e46f83d88fda972cecc271d3d38923c01f1746f9662d10"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-6d4232011c9a9937cedfc9642ec656302631e62d2d18239c63a5cfc004fecf1f"></a>

## Direct properties — headers / 0cf6fb3656d9 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-b0845e278fb99fae5533c38108cfa6f88d62afac44edc63367e0bd1521a1ae26): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-1ed137b1bfd42683aa771b53cf1a3841fbb7665fd3f3867dd78034d4cc413ce3): complete subsection reference.

<a id="canonical-06f9e5f331bdff84a77eea42067b2dfade15741e07b42d14aa2d126ab9154c03"></a>

<a id="canonical-ce024743216742b21fc0bafb7b4ba89c8e5a8423f86a803b51ae878202c32a4d"></a>

## invert_matcher property — headers / 0cf6fb3656d9 / 4

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-7e3b4486ec8350747f1653855690f9f0edce9c762d1ff90a0e77fe33610b6682): complete subsection reference.

<a id="canonical-d1a375df1099812d75f539c6c494f2627436f631053fff751b8e48e4894df918"></a>

<a id="canonical-45893f1e66dcba9d33d265d7387c52c014d2c86934971c09bd81f203d36c96e4"></a>

## name property — headers / 0cf6fb3656d9 / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
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
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-08e1e1583af68b5d7a0b37c52a72ca083069047c7ba3da187443dbc3f8107343"></a>

## Next pages — headers / 0cf6fb3656d9 / 6

- [headers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-b0845e278fb99fae5533c38108cfa6f88d62afac44edc63367e0bd1521a1ae26)
- [headers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-1ed137b1bfd42683aa771b53cf1a3841fbb7665fd3f3867dd78034d4cc413ce3)
- [headers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-7e3b4486ec8350747f1653855690f9f0edce9c762d1ff90a0e77fe33610b6682)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-b0845e278fb99fae5533c38108cfa6f88d62afac44edc63367e0bd1521a1ae26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-850573fd9ad6cbfe16c95d4bab02382cdc94f8c722abb5734ab8bd3415c6dff3"></a>

## headers.check_not_present — headers.check_not_present / d3859ed4c3eb / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8)
- headers.check_not_present

<a id="canonical-a5655798e0bb369c02a5f730c7da57d57bbe9661da7cba898a9da6ae3157600e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0c92fc203dcb637a9d75672790b37a6d7d33e078ac3f8381eef504becd1dfa63"></a>

## Direct properties — headers.check_not_present / d3859ed4c3eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f27713960f6740793507be4cc8a2a0a82131a849e58740bee13e98a9b9d23bec"></a>

## Next pages — headers.check_not_present / d3859ed4c3eb / 4

- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-1ed137b1bfd42683aa771b53cf1a3841fbb7665fd3f3867dd78034d4cc413ce3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2df6b39afacbbafa65f49129e184b2b77c209805ae38926069beea6f5acf90c"></a>

## headers.check_present — headers.check_present / 7ee73b520d03 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8)
- headers.check_present

<a id="canonical-03f662abf99723a1f087b2512052fdd2ee5d1cd2db5870aa87458ae845016e65"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-98f70084495f82c0eba618d17d9b1b1fbf758756b4f16fef4883bab6075bb475"></a>

## Direct properties — headers.check_present / 7ee73b520d03 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f8f8677e1f48f7cc844cf6f5174df4549c0a2f45e3ac2aa5918e1b594a2699b2"></a>

## Next pages — headers.check_present / 7ee73b520d03 / 4

- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-7e3b4486ec8350747f1653855690f9f0edce9c762d1ff90a0e77fe33610b6682"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b813e6ec236a6d51b2a5379ad168d9e5a8893faff65f9d187b21d5804a77169"></a>

## headers.item — headers.item / 6764c29684f1 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8)
- headers.item

<a id="canonical-73cd52428eba6ff49c6c5351c520d3771ebea902832ed039975488e0cbd01462"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-ab58b2f69044384ded8388760cd4b528bec29d36d7293d86d5a18a265812c09f"></a>

## Direct properties — headers.item / 6764c29684f1 / 3

<a id="canonical-4f07bbb17807ca1f2944b3355d06b15e8088ff7f9f518c80974e1f0941b0904c"></a>

<a id="canonical-fbbacbae0655bbd5c3e39e8c74199386c0a2f452a8b523f02cd7b6d892cae67d"></a>

## exact_values property — headers.item / 6764c29684f1 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0fb2700c93aeba6c4c78275fdb0bb7898ac65375b62dce94af1b85bcd773de05"></a>

<a id="canonical-8cd9eae32cdb23e46581418aab354d72d20171f0a50881f6e97cec072cf77d18"></a>

## regex_values property — headers.item / 6764c29684f1 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3eacaa5f355ef7fce5997a8c9d0898e21515d2de69513037d1132f08509e0f04"></a>

<a id="canonical-ff91f0c2c0ac478254dc228ca430560496d00fccd86c1c3a5f14f8e1b7f50d13"></a>

## transformers property — headers.item / 6764c29684f1 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ab2f0d3a668575a3c078d91fde8eeb51d13341de8e7a1e9d156871f5360e52f6"></a>

## Next pages — headers.item / 6764c29684f1 / 7

- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-c92264cf3259df8496d02d921a164815a8eab57577896d7cb6f14bd53be1ffc8)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-2ccfac68219f29dec73d5b10667fdddc8a3d5f24c54e6fdbd705961ddb8382bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bd9fe63ab89c6f998b114c4ebd156d006a24559262bae037d3faae09a1eec68"></a>

## http_method — http_method / ed9c8efbda67 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- http_method

<a id="canonical-f19c8dd827a26b5d469c865b20c30c0f09379c75b7882f8eabd8d6f5b134d6ba"></a>

Type: `"single"`. Computed.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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

<a id="canonical-569111b4a5b25fae7bdb6717f26eb04068ed2b42b0997a1f158e6552d7344752"></a>

## Direct properties — http_method / ed9c8efbda67 / 3

<a id="canonical-88d2303fc746cc8b12212f262be8e3985478d2d5e7fbfe84087815eb9a95c6a4"></a>

<a id="canonical-1a4db1836467871a9c221e2d862d4b20364099d2e9d44caa3e236e6053662ada"></a>

## invert_matcher property — http_method / ed9c8efbda67 / 4

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-28842e6238d0a0bdf2c7f82da648fc8da8c4108d3e18d60bc27e82aeae53cf5d"></a>

<a id="canonical-b299c31c77818f794cb311b393a62544c129b286d47f65f75af28244378018ba"></a>

## methods property — http_method / ed9c8efbda67 / 5

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-68c5ef0edf9fecfc795fb567008bb1d3fd88f0d35eb847ac2598efec26f376a0"></a>

## Next pages — http_method / ed9c8efbda67 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-4c64fdc40e1c3b6cadd843f10a9fcacaa6d87d0a051cb31ababde925f2d339dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef36e7a62b5d01c6aedbfb74dd762bceb8bb11c38cb0a4965c0c1803fb4e40ec"></a>

## ip_matcher — ip_matcher / 33b097059f01 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- ip_matcher

<a id="canonical-f2fe1d8d8a1cc10be12ab080a847dc9147eb8a1445ba7229336ebaa3da4c08f9"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-a867078baf0cc8239000b1f923ee9fd6165396a3f2843641b62189d85e2039eb"></a>

## Direct properties — ip_matcher / 33b097059f01 / 3

<a id="canonical-4b77d6bd7758aae142ba6b6a834889cbea71caadb5a86704952f0e1200c99e0e"></a>

<a id="canonical-7a4795d29ef295b005c0557f1bb5f4cb19292bcec6c3ac8f0e8eab85ae0095cb"></a>

## invert_matcher property — ip_matcher / 33b097059f01 / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-260f4438b4738003f4f1025d3fc317322d5ec57a7997c3bff5b7f7080c3ac8fb): complete subsection reference.

<a id="canonical-0b0433cfd6fc6b6cb6db6055a3957b0a1354efc13ec7e8cd3ab72d80dcbae32d"></a>

## Next pages — ip_matcher / 33b097059f01 / 5

- [ip_matcher.prefix_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-260f4438b4738003f4f1025d3fc317322d5ec57a7997c3bff5b7f7080c3ac8fb)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-260f4438b4738003f4f1025d3fc317322d5ec57a7997c3bff5b7f7080c3ac8fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c8c33be2a1713018ab31933d3a70f606288f69e26e0338ab959e10aeac0b7ff"></a>

## ip_matcher.prefix_sets — ip_matcher.prefix_sets / a02f1938a529 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-4c64fdc40e1c3b6cadd843f10a9fcacaa6d87d0a051cb31ababde925f2d339dc)
- ip_matcher.prefix_sets

<a id="canonical-4b3de46b55d3bb10bb709e635a1272ef715d628cf34f04e179cd76cfdadb8537"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-f89e75e91e2d62d09cb10eb946c3e966d500459770ac6883676511a7d82aa688"></a>

## Direct properties — ip_matcher.prefix_sets / a02f1938a529 / 3

<a id="canonical-83cf436e6c0c0207961d7f5a0f8550108d98eaf39d5bd5fb548f0e5fa5dd6d27"></a>

<a id="canonical-6b117098f2b75243bbc04a1cd6fdd9712761f2a909d51441c87231e2e3b11592"></a>

## kind property — ip_matcher.prefix_sets / a02f1938a529 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

<a id="canonical-2be3e9c55a91f54d75f5d2943b976d89f90ba176b4e002302d9750439e2c2d39"></a>

<a id="canonical-c4f752ca16953a384ff211cab0e4f36adc7caaea43a38199a0e2428bb583cbc8"></a>

## name property — ip_matcher.prefix_sets / a02f1938a529 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

<a id="canonical-d20e03d868e9ab8bc4f245cf059c776f558c3aa989746acf99dad7eb33dcd397"></a>

<a id="canonical-5a768eb6751cc5e16fbd68a5c645e9e5a6da5d6ed65f929bbec718bb74fe2913"></a>

## namespace property — ip_matcher.prefix_sets / a02f1938a529 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-82cebf422b341790bdb3f4fe06d68e4cb1ed085db23d164b1a4433a12422e249"></a>

<a id="canonical-0a9597a91edb3210700d542c0a2a55e8fe0ee5b66482b3586384130cebb9a150"></a>

## tenant property — ip_matcher.prefix_sets / a02f1938a529 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

<a id="canonical-db35f0c455c0015813f4595d1137af1282fa1af17edf5406d67b46bcc51ed6ad"></a>

<a id="canonical-2a082987c30a6647df8dd9a9738e12ee175bd80e21070f35f1ce663ef9ec4e39"></a>

## uid property — ip_matcher.prefix_sets / a02f1938a529 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

<a id="canonical-125408f9ddd917f64e37028009fb3f6eaab38acd377a54ae369fef88b533e012"></a>

## Next pages — ip_matcher.prefix_sets / a02f1938a529 / 9

- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-4c64fdc40e1c3b6cadd843f10a9fcacaa6d87d0a051cb31ababde925f2d339dc)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-f568569c6da6e1d479ee5ca4c660b0f3b3a6e5bdd4c77d6f54811e6a1ffd3098"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b0d1ad865fbce55b225b35e714be5db78c7cb9dcfaeb1857a597bfa8897ee71"></a>

## ip_prefix_list — ip_prefix_list / b3b3874474c1 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- ip_prefix_list

<a id="canonical-6696dbd25074da043944f7708ee1d094c2ee34bd0224758b8f8dee121b1bf2d0"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-47657343a431cd4538302849b30b9bdbb138ea85b24b89c712c404dff67c1723"></a>

## Direct properties — ip_prefix_list / b3b3874474c1 / 3

<a id="canonical-a55e594b0d8b742ece741e845b2d01486306893e68f0cd4e62dfd8e77ecea14b"></a>

<a id="canonical-281ef21bd1fac3bfa801026342a0d1b282effaecee9f36dd715563c5431c98f4"></a>

## invert_match property — ip_prefix_list / b3b3874474c1 / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="canonical-8e35c01c375275071c0ded269d2f186fcf44fc2e7bd0cc5c309e714a90b1736c"></a>

<a id="canonical-a0b718fcc46fabfa61df52a00d4dae8385e5d31ebd12f18470ab83465e4b1d92"></a>

## ip_prefixes property — ip_prefix_list / b3b3874474c1 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e75fd52b28e862f2100102ddab669c3fb60e0a3fb5e4def00153f32f11f985a0"></a>

## Next pages — ip_prefix_list / b3b3874474c1 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-e2480d86388e5c9fb566e99804e4614053a7dba6f8db96780101bd841aaf01f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f98b54532b916154ab9dddba838d275eb48af4ad02a947534781a4ed427f3e3"></a>

## ip_threat_category_list — ip_threat_category_list / 81e6ad5ae6f5 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- ip_threat_category_list

<a id="canonical-2edf112999a7fee20a7b6b204fd3bffb224f519efb9ca33941c1d2081a5803c3"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-e005a6fb3d17cbd0ac698a656463fe9885b4d6b278d12859c7624b839ad38e92"></a>

## Direct properties — ip_threat_category_list / 81e6ad5ae6f5 / 3

<a id="canonical-0fae64cc3b0ac17f5c135d044c4b02e1ded8a61441ccb8e9206f5b121885217b"></a>

<a id="canonical-deffd87701ae4d4bb03f7a2ca42d14957ae6f90f3d1b2a97c722e82399cf4313"></a>

## ip_threat_categories property — ip_threat_category_list / 81e6ad5ae6f5 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-626a4e1e074494f60ca0f0ce6a8c1c9c44c70b52cd3ce16aab09782833098345"></a>

## Next pages — ip_threat_category_list / 81e6ad5ae6f5 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-fb9290901816a8a4f9e84e612d576c12566a1b4d776627310a573f99ef9fcfc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55096f0bb80a61845434a6999c2b6ea8ac1bdf6c849966468ad2c235e486f7dd"></a>

## ja4_tls_fingerprint — ja4_tls_fingerprint / c4b7fff3c5e2 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- ja4_tls_fingerprint

<a id="canonical-d149fcd44be04087577cadf3b6123f95d9aee9f7342b68300a02e8277f5ef8e8"></a>

Type: `"single"`. Computed.

\[OneOf: ja4\_tls\_fingerprint, tls\_fingerprint\_matcher\] Extended version of JA3 that includes
additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a
different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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

OneOf alternatives in this subsection:

- [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-d149fcd44be04087577cadf3b6123f95d9aee9f7342b68300a02e8277f5ef8e8)
- [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-e44b775e4590083f8137f4b607de5d213b8808fc5179cdcd48fa9c663d5757c4)

Select alternatives according to the provider validators above.

<a id="canonical-22687e8773cca9277c4d20be5743f18bbbbc88063a2bc0cfa10aaa446fa07c82"></a>

## Direct properties — ja4_tls_fingerprint / c4b7fff3c5e2 / 3

<a id="canonical-36c99c5aaf1a105c34a70f7f198ced0a277128e6a677dd14d1893faedaec8ed1"></a>

<a id="canonical-c6b841dc0ff857e9e718e13f1e5f342caaf1af2000bdb50a2983a828bc8a9b45"></a>

## exact_values property — ja4_tls_fingerprint / c4b7fff3c5e2 / 4

Type: `["list", "string"]`. Computed.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-eef5afb1e489e8ab9593ea9b174a7e7dc84ac1e0d8ff19534822b73f3f063821"></a>

## Next pages — ja4_tls_fingerprint / c4b7fff3c5e2 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a2a4fcbdf30af5aea427c12b9bc6adc9c9ac254f8f63c865154c0e0aef35970"></a>

## jwt_claims — jwt_claims / 06eace6db970 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- jwt_claims

<a id="canonical-903651741a070ccaad3450637b0947ee47976ef53aca2e7fcda23d4dae1608ab"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-72ce0797cfbbeee84686873bb86d4c3cf92dfffd57f1122145ae16f809239e35"></a>

## Direct properties — jwt_claims / 06eace6db970 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-8387f0a50bae0eaf04bf6f99c8debce31f82e8c09ec7ef88a336b36c6982f325): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-e4d5c3f01638ab204fdcb063fefe81a898fc7b32961d1ccf3539bdc231f98566): complete subsection reference.

<a id="canonical-6cf777cf61ccdf046266b907d6320ba044709822cfc38bd6786d2735f3b19885"></a>

<a id="canonical-2b57505bfd468491b5f8228926a508701dbd1dd1c89854e9bc3ef3c276331aca"></a>

## invert_matcher property — jwt_claims / 06eace6db970 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-7533f906b3fd0fd09959606201096f6528c438c29288e39cc32ecd6423f4f264): complete subsection reference.

<a id="canonical-a592b14eb0f688210a25c22326875194ba6f273b953d9cff11f5783286000fce"></a>

<a id="canonical-85966378f3bacc6e5b6793f1e1c4574a542725e2335334e44346932d89424b98"></a>

## name property — jwt_claims / 06eace6db970 / 5

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
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
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2d3a42f5338f86398121e77c717fec338264e8b9b06907ecccfc151bc1ea03ba"></a>

## Next pages — jwt_claims / 06eace6db970 / 6

- [jwt_claims.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-8387f0a50bae0eaf04bf6f99c8debce31f82e8c09ec7ef88a336b36c6982f325)
- [jwt_claims.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-e4d5c3f01638ab204fdcb063fefe81a898fc7b32961d1ccf3539bdc231f98566)
- [jwt_claims.item](data-sources--service_policy_rule--reference--group-001.md#canonical-7533f906b3fd0fd09959606201096f6528c438c29288e39cc32ecd6423f4f264)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-8387f0a50bae0eaf04bf6f99c8debce31f82e8c09ec7ef88a336b36c6982f325"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4ccdbb4407481fc648df20e70c0c402b55dd05dfd5c5ec05304d6fbbdb13bbb"></a>

## jwt_claims.check_not_present — jwt_claims.check_not_present / 0c6579f7faca / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36)
- jwt_claims.check_not_present

<a id="canonical-5865f83fa6374d4b02f40aad04f3ab4b089e4099ddf2d6756c308376c9c50814"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-ad714621cb2206ee14a5a25f612d98f49c735da862b4d710c08fccd33a108dad"></a>

## Direct properties — jwt_claims.check_not_present / 0c6579f7faca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc073a20435c50ff31f9326d700147e7c90a7d3e80971654ab0fcce5cf4ffcd5"></a>

## Next pages — jwt_claims.check_not_present / 0c6579f7faca / 4

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-e4d5c3f01638ab204fdcb063fefe81a898fc7b32961d1ccf3539bdc231f98566"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f161aab2c0916a36d5c541376e017a1728d4adf111d32b0bae1c26356e8d6161"></a>

## jwt_claims.check_present — jwt_claims.check_present / cedfe43164ea / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36)
- jwt_claims.check_present

<a id="canonical-7da7d5525238a9802e03b553405c631bf47ceddd9021b227f126fd144ecd1449"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0a4288718fea6c1749ae5bae27735b4c9bfcf14aae289cccbaebbd125bc3dded"></a>

## Direct properties — jwt_claims.check_present / cedfe43164ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02051cf282a8d81180a949391bb3a312ba197f8ff8d8c1c244b8962c01dcbe1b"></a>

## Next pages — jwt_claims.check_present / cedfe43164ea / 4

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-7533f906b3fd0fd09959606201096f6528c438c29288e39cc32ecd6423f4f264"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d18f9929149207021fc2e6315efffca70baab5d6dac0464fd849d1ada9ba71e6"></a>

## jwt_claims.item — jwt_claims.item / 3ac30a84eb75 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36)
- jwt_claims.item

<a id="canonical-65222b06e80f122fe3d8185fc13bbdb554354681ef2cd72ead7ef7f5d22f266c"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-310b1a86ec472d3c063391de96fedde9a48ddd78bd73a8487110746b582b633f"></a>

## Direct properties — jwt_claims.item / 3ac30a84eb75 / 3

<a id="canonical-9fc7876d081525a785694fa2bf4c7208f1d70e86d33a19dd6f16b0ac9a21005b"></a>

<a id="canonical-02225f6fabf3b46b3acbc64e2a333bf27ed8accd9398b0eafe2c42db65dc5063"></a>

## exact_values property — jwt_claims.item / 3ac30a84eb75 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7e3db97c0aa0cc035933f6aa1ea7ef8892396e13f33741e4fb3048cadf4e23b6"></a>

<a id="canonical-93444bc2eaf1ca05afe97721073535ebc4d9961f68ef80fe220b56733208a9f2"></a>

## regex_values property — jwt_claims.item / 3ac30a84eb75 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-51061f12207b14b41b17370ba64640dab9a226454137a72edbc72e6e33dc66cb"></a>

<a id="canonical-982dd4b23bfc1443a92294849acfc214ab22a0d23806a4d1c695d792012b4063"></a>

## transformers property — jwt_claims.item / 3ac30a84eb75 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-28e420b824cbdcb1eeb933a07a86fc003ad00ec10370acb30e84fd7226f62277"></a>

## Next pages — jwt_claims.item / 3ac30a84eb75 / 7

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-efcc8e2b21eacc6b551eb97ab526a6923754e5aee248d6c11b5cbf6853152e36)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-e4d40e388ca9969028510e765949d2ecf07e9d65a0c3bd535e05cee68412ce0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61411acc6a90a81cd44c7ef4edc97047006089e794efa34b9aecd2dfde3f959a"></a>

## label_matcher — label_matcher / de994b1c22c3 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- label_matcher

<a id="canonical-20f34c8bf0cdf7c59b22cdd17daf333f7324a91cb4a87384b89411765cc65850"></a>

Type: `"single"`. Computed.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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

<a id="canonical-2a8f41d08eca2171b01bf75a6f25527814f4e1e852067117ba4a5dbe5dcb3d54"></a>

## Direct properties — label_matcher / de994b1c22c3 / 3

<a id="canonical-41700beadad031f073a63a1b7b43f8edb0d43f37e8c739800aa26b84ca01bf9e"></a>

<a id="canonical-e4f98fd16bc62f510d4626d0c188e887e4259aad67410876382a503bddb97e6a"></a>

## keys property — label_matcher / de994b1c22c3 / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fcd18faba6e31921356aa9fa3bbb86896df5620fb66255e62c0ee65562370754"></a>

## Next pages — label_matcher / de994b1c22c3 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-137e55d694ff101b1eceed4b0f2170fade6d8b7a600a590d2a83bb1efb5d0eb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0620403b32f2befd7204cc68b899ce18af4fdc8ac412b8c820d68101fbeb58bd"></a>

## mum_action — mum_action / 41d644d5bfd8 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- mum_action

<a id="canonical-6eb35d8e99f65232273f6c3f06bac3e5e03c321cb6798033257d0d9856cb317a"></a>

Type: `"single"`. Computed.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

<a id="canonical-f1e5f17daed50057ea608b0c0a512637bb06ee44db3b5f925090ffcb304d00e2"></a>

## Direct properties — mum_action / 41d644d5bfd8 / 3

- [default](data-sources--service_policy_rule--reference--group-001.md#canonical-83223e6b1bbd942f5e5ec28eae169fe07663cbd0dd45ecb0440c803b2df81969): complete subsection reference.

- [skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-e6b52c574809f67bf9b2e3cc9c7f95a712b9d48e5d7f86fe7aa78e1c637e099e): complete subsection reference.

<a id="canonical-e919e8fd90987a59b52532091ecd9d9d378027b6f3030afddfc5dd1ad92d905f"></a>

## Next pages — mum_action / 41d644d5bfd8 / 4

- [mum_action.default](data-sources--service_policy_rule--reference--group-001.md#canonical-83223e6b1bbd942f5e5ec28eae169fe07663cbd0dd45ecb0440c803b2df81969)
- [mum_action.skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-e6b52c574809f67bf9b2e3cc9c7f95a712b9d48e5d7f86fe7aa78e1c637e099e)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-83223e6b1bbd942f5e5ec28eae169fe07663cbd0dd45ecb0440c803b2df81969"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
