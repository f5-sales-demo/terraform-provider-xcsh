---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ce6b1fcbb08338e0fc7443fdb474c0127a2046b7a9c860b59e670b06566b840"></a>

## Property reference — Property reference / 08f531c4f6cb / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- Property reference

<a id="canonical-77e8609c014382e9e64242ef2cf36699da70d1dc14f87e7dcbdd22c7e1d531a4"></a>

## Direct properties — Property reference / 08f531c4f6cb / 3

<a id="canonical-b6925c5497384af9075d9ccb981c6617e20063c0f4dc63bc8c2d7af13df392af"></a>

<a id="canonical-50aef9541c5a039759c2d79f0e0336a70d41562b57f6bc4bd6426124180fc2cd"></a>

## action property — Property reference / 08f531c4f6cb / 4

Type: `"string"`. Required.

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

<a id="canonical-e42a7ae17f58f9215f4ec2de533b03d2a2232253f4e494895c70cd7e47a72a74"></a>

<a id="canonical-0dca2d8e4e3db0c39f45debf05e98f6adbe9a38f07a3a769a3e351faf2c06616"></a>

## annotations property — Property reference / 08f531c4f6cb / 5

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [any_asn](resources--service_policy_rule--reference--group-001.md#canonical-29bacf87798d1bc903b9ed87078c261fb07693a10b86f3252e680ec149ce9124): complete subsection reference.

- [any_client](resources--service_policy_rule--reference--group-001.md#canonical-1b190696041b8893594e5e23d2300f111e413977553cc089857c9fb5bc1ead34): complete subsection reference.

- [any_ip](resources--service_policy_rule--reference--group-001.md#canonical-bb3ff9935dbfdd859eff4a15408750eee13afb0315c3936d0080835ee6462570): complete subsection reference.

- [api_group_matcher](resources--service_policy_rule--reference--group-001.md#canonical-8181b58d11566e2d3bea22f44f5becd3bb2e2b9c68f63c0de829acd0c755fbdc): complete subsection reference.

- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94): complete subsection reference.

- [asn_list](resources--service_policy_rule--reference--group-001.md#canonical-09fa9b0f03a116e8889aab4532d7ee5dcb87005230bb50915d9e937361308ec4): complete subsection reference.

- [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-737775f18217586217b199bb222bdcd58421af7102d5cf785545f7621ff6f6a4): complete subsection reference.

- [body_matcher](resources--service_policy_rule--reference--group-001.md#canonical-41478796dc9f16cc362d425250c15e826413f8248849eadf86e9eab58399f5e1): complete subsection reference.

- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-936c6a7953e5f16177b88a818686a49d30aa935b9f8caa61e07bfcbe6b11ae52): complete subsection reference.

<a id="canonical-e42a2762a0028e79d9947feeb6fd03ae69a307e752948729a0752e765ddf7c38"></a>

<a id="canonical-4342eb9083333f9bdddae8dd31f1a77b7dbd7ce221f1515ab9748e19e9c99e30"></a>

## client_name property — Property reference / 08f531c4f6cb / 6

Type: `"string"`. Optional, Computed.

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

- [client_name_matcher](resources--service_policy_rule--reference--group-001.md#canonical-3e6fcc4950ad994ad911391dcf12a4305b0b9c75569d289775436104117fd307): complete subsection reference.

- [client_selector](resources--service_policy_rule--reference--group-001.md#canonical-70d913f0f8a715b2ff59d453d1f84bd6685e80d1a6caa80a9fef7d9e8dc7195c): complete subsection reference.

- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f): complete subsection reference.

<a id="canonical-fe006e21674de84102e233c2b9f76aacf051b65ca6eb96840f9389de6a3d21b0"></a>

<a id="canonical-ed770bc6894872eee8bd9fd25397b63481fa9bb5c0f8b23feb88f5760059ac85"></a>

## description property — Property reference / 08f531c4f6cb / 7

Type: `"string"`. Optional.

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

<a id="canonical-79230ed586a89a0132afb38b93453512d9351bfc9579711f6d051d80070d60ff"></a>

<a id="canonical-0cbc3fb7ee2f2e1518dbf67188cbd9a61e8448352a1c84bce01642315720f958"></a>

## disable property — Property reference / 08f531c4f6cb / 8

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [domain_matcher](resources--service_policy_rule--reference--group-001.md#canonical-7ea9c15f3bbe81e53c56ac87cbe279b4bf1e79773e794059e812c84b8e4e2be4): complete subsection reference.

<a id="canonical-1427758097d7ed42defbed3d880163985626f9613135fe342843cb922b22643e"></a>

<a id="canonical-ca230d4ad20d9fb71b5a32e23234523659bf7c0f8a3893aa5fb981e893cfa1c2"></a>

## expiration_timestamp property — Property reference / 08f531c4f6cb / 9

Type: `"string"`. Optional, Computed.

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

- [headers](resources--service_policy_rule--reference--group-001.md#canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480): complete subsection reference.

- [http_method](resources--service_policy_rule--reference--group-001.md#canonical-f5affe06710c1b7d6014ccf515752fa30f078d58d20a6cf7e719436ab7db52ae): complete subsection reference.

<a id="canonical-e1e850d300a3673a52b3b650a96926d20920f8a36d4b5f1de86c09c22fdce34e"></a>

<a id="canonical-d9201534e8a36a95b5674673ed524ce83d4ebcd622f6319c86da47700cb6abc2"></a>

## id property — Property reference / 08f531c4f6cb / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-5131505e9c5d1806673eebf13dd26429a151fd79793834461f68747727fecbf6): complete subsection reference.

- [ip_prefix_list](resources--service_policy_rule--reference--group-001.md#canonical-413521cec6a72c537f18e294f798d20fdf3fef19b519f823e0f7d2e09417ab79): complete subsection reference.

- [ip_threat_category_list](resources--service_policy_rule--reference--group-001.md#canonical-06cb5e2962be24ee85db336adacd12583343c5a13b249653673224a7757d992f): complete subsection reference.

- [ja4_tls_fingerprint](resources--service_policy_rule--reference--group-001.md#canonical-aed76a8d199a5769ebe931308623edc55ce47695ebacff7f084193b3f4a9709f): complete subsection reference.

- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397): complete subsection reference.

- [label_matcher](resources--service_policy_rule--reference--group-002.md#canonical-885654158fb78e99195654bfd7eb4ede70c6b02f2d395cef3f137f6410b798c0): complete subsection reference.

<a id="canonical-af02cb4ea4a3d466df78fa00d44d53c01eb8d113b460036d60ba7607422ea2ea"></a>

<a id="canonical-1176d4f2a05322eafa640e89edd01e52cc0f93d422141d4fb7d70dda3c93097c"></a>

## labels property — Property reference / 08f531c4f6cb / 11

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-ead4ce652f5b40bb98f609b2fc7a5fd3ba8643991f882991355f5c091c996f85"></a>

<a id="canonical-451ee1643138795c737b87c5b9f5ca2ba43b417b8e71b247bd910fa7d2f65167"></a>

## log_rule_evaluation property — Property reference / 08f531c4f6cb / 12

Type: `"bool"`. Optional, Computed.

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

- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-8173a7b12b6dbcabe15c5a24896b2850104c373bdd419858980760f784308f9d): complete subsection reference.

<a id="canonical-f5f25530ba9a9663beecfc96d008a36e34ff91c2d21a6b6607841760d9ba7e69"></a>

<a id="canonical-db4d3c1a4bddf7134df92352b6ee113731f4f9f2b7a30218bf6281a873bd469b"></a>

## name property — Property reference / 08f531c4f6cb / 13

Type: `"string"`. Required.

Name of the Service Policy Rule. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-6cf3ba6a9ddea8886019e28ab3607055c5e8159b9eefbc712c1b892a4a62dd03"></a>

<a id="canonical-154b072c18bd16cd87f746c005b4c4ede679e93b5979611249e7933d287dff1b"></a>

## namespace property — Property reference / 08f531c4f6cb / 14

Type: `"string"`. Required.

Namespace where the Service Policy Rule is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [path](resources--service_policy_rule--reference--group-002.md#canonical-c453af0b534656a15fed7afe77bb9ad4115339a74c815501323d7c06337a777c): complete subsection reference.

- [port_matcher](resources--service_policy_rule--reference--group-002.md#canonical-a3aa60eda2c10f207494ee2277014d7f3dde42d696ff2dcc542c02648fd6d21f): complete subsection reference.

- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111): complete subsection reference.

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886): complete subsection reference.

- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9): complete subsection reference.

- [timeouts](resources--service_policy_rule--reference--group-002.md#canonical-49d1dcedd566f98a405b94194bc73262f7147467a3f23b9ded4df6aaafb63e75): complete subsection reference.

- [tls_fingerprint_matcher](resources--service_policy_rule--reference--group-002.md#canonical-a00396ea59b4ee1bd6e1ea84d1322f649375d00c7f7360e57c8af940324e416b): complete subsection reference.

- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76): complete subsection reference.

<a id="canonical-67c68d80b42d0045fa833ed8a9ca283e6f890e1d267a4c91fc57e01ee59a643b"></a>

## All schema paths — Property reference / 08f531c4f6cb / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--service_policy_rule--reference--group-001.md#canonical-b6925c5497384af9075d9ccb981c6617e20063c0f4dc63bc8c2d7af13df392af) |
| `annotations` | [annotations](resources--service_policy_rule--reference--group-001.md#canonical-e42a7ae17f58f9215f4ec2de533b03d2a2232253f4e494895c70cd7e47a72a74) |
| `any_asn` | [any_asn](resources--service_policy_rule--reference--group-001.md#canonical-89b8b4c875a6deb8aa66d324fe177e5402017a5d46c2255ab9ecd73b930658aa) |
| `any_client` | [any_client](resources--service_policy_rule--reference--group-001.md#canonical-3e6e098d9f0c2d679e2aa75ef9171a69a1451766e9dad02e49a69b38fd6f6b04) |
| `any_ip` | [any_ip](resources--service_policy_rule--reference--group-001.md#canonical-be87293db2b603e2353fd6ff5be02c5a124868c1ef4814a4aff4df8afbb6043d) |
| `api_group_matcher` | [api_group_matcher](resources--service_policy_rule--reference--group-001.md#canonical-08c17e728005ad0f20d0e9c5242f5d794473e394872aaffeef9346b0f54470fa) |
| `api_group_matcher.invert_matcher` | [api_group_matcher.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-36e25bdf3721d7e428f63c80a17213208c4d051e5cee977dd8f5bca466eed763) |
| `api_group_matcher.match` | [api_group_matcher.match](resources--service_policy_rule--reference--group-001.md#canonical-1b1453015df43a83fc69d090045282b2fec2cf8f16be44297daf99f56beef818) |
| `arg_matchers` | [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-6a00b2f13c0a10ac6a6f95f84ce38ab9ca85b90efb152df3cda82456fc9b5791) |
| `arg_matchers.check_not_present` | [arg_matchers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-21d21899231beb535e28ee76d26fff180758552339fbbad1f5828a10285de744) |
| `arg_matchers.check_present` | [arg_matchers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-1161fb80d13e3e33ea406481777e3cb67e839cf2f6cb4ef188eb6def55db6a78) |
| `arg_matchers.invert_matcher` | [arg_matchers.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-b2a52e82c842a73e123c6e0f7523a8e360e6975003633322f0bde65564416103) |
| `arg_matchers.item` | [arg_matchers.item](resources--service_policy_rule--reference--group-001.md#canonical-89e5bba94ccd8768818101d5af93915fca2c29e14b4d3179d97513a6b1ff6d42) |
| `arg_matchers.item.exact_values` | [arg_matchers.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-f9dbd40d4d6008ab516baef9e15696a6ecdad1c5f86f87f3f19d158cbc92ceb6) |
| `arg_matchers.item.regex_values` | [arg_matchers.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-fe76e1f86817b2feb030b03ec047922d53c7717ac4e7ac5c702a160315cd9d2b) |
| `arg_matchers.item.transformers` | [arg_matchers.item.transformers](resources--service_policy_rule--reference--group-001.md#canonical-e9a351d5526941c0c72d18b075229c7f4596a639c0670060a8a69774a456b5bb) |
| `arg_matchers.name` | [arg_matchers.name](resources--service_policy_rule--reference--group-001.md#canonical-ad89d2e6433ee0cfee3fa906ec87ae8b4f04e26161467960268a446f43d71d62) |
| `asn_list` | [asn_list](resources--service_policy_rule--reference--group-001.md#canonical-921e189a1413d60c1e3e453034432ee30e83bfc13f351e7cce703ec58e6fa8aa) |
| `asn_list.as_numbers` | [asn_list.as_numbers](resources--service_policy_rule--reference--group-001.md#canonical-73037715a40f7109f627140c03ee6e9f375bc95455d4ea992723794c993c3cd8) |
| `asn_matcher` | [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-6949d079257b07f0ce687309fc1bc8f855c9638a840d3113e6083b728e5a7394) |
| `asn_matcher.asn_sets` | [asn_matcher.asn_sets](resources--service_policy_rule--reference--group-001.md#canonical-32334d5527daec2e4f98e5647f509967663f940846113d42f2f56a22dccea6ab) |
| `asn_matcher.asn_sets.kind` | [asn_matcher.asn_sets.kind](resources--service_policy_rule--reference--group-001.md#canonical-966e8a696cc65276355423426e39ba7b9562f23c1e3cde3e58e0f6e70101caea) |
| `asn_matcher.asn_sets.name` | [asn_matcher.asn_sets.name](resources--service_policy_rule--reference--group-001.md#canonical-a1de0353b6d11047b2b655c3fccd78c5f31391d66ebf3c1ae34f68404ee9ae9d) |
| `asn_matcher.asn_sets.namespace` | [asn_matcher.asn_sets.namespace](resources--service_policy_rule--reference--group-001.md#canonical-64758910f97b47dd6d5fcf972b0f399ff051a1acd202e28e7c1ef5d4aa63e4ea) |
| `asn_matcher.asn_sets.tenant` | [asn_matcher.asn_sets.tenant](resources--service_policy_rule--reference--group-001.md#canonical-b6aacb140770277934c9554f36a848d7bdaf17aad6377bc3207b03561e4b5158) |
| `asn_matcher.asn_sets.uid` | [asn_matcher.asn_sets.uid](resources--service_policy_rule--reference--group-001.md#canonical-9071b2e1412e8aba118abe8879171273ba22e9ce2f103d2e8dace692e03b578e) |
| `body_matcher` | [body_matcher](resources--service_policy_rule--reference--group-001.md#canonical-5a2bbc5e6f43b1e6da37ddd219c7d8059fa7523c86bf03a3dba56411cbe9cfe3) |
| `body_matcher.exact_values` | [body_matcher.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-b662d9229907c473f47256d4c955783b0694362f0e5bf2b59908996a80af25db) |
| `body_matcher.regex_values` | [body_matcher.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-ef0506a1cb07786b53af416c149fb8eba0bd5e15e0a75971c27f2ff382ccc76d) |
| `body_matcher.transformers` | [body_matcher.transformers](resources--service_policy_rule--reference--group-001.md#canonical-97c5604e76d0a46016d776b45b00ddf2ed6bd271eb2117668693b42209f83263) |
| `bot_action` | [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-25ffaaff0624f097d7f3168ddfc3ffd3f2d981759fa2d95ba562a47b6fd600d2) |
| `bot_action.bot_skip_processing` | [bot_action.bot_skip_processing](resources--service_policy_rule--reference--group-001.md#canonical-be5ea125580f10de5815da936976d8ce6a8c4a2841834a4fa9fc34e938674007) |
| `bot_action.none` | [bot_action.none](resources--service_policy_rule--reference--group-001.md#canonical-82c59bf4c3c738268e9116268ffc4554cdd37264e8008d9cff443d64d64f6452) |
| `client_name` | [client_name](resources--service_policy_rule--reference--group-001.md#canonical-e42a2762a0028e79d9947feeb6fd03ae69a307e752948729a0752e765ddf7c38) |
| `client_name_matcher` | [client_name_matcher](resources--service_policy_rule--reference--group-001.md#canonical-56ffd1528ef78daa1c36b951f5930c56bb9d8518e6f213f9573bcef765ab2229) |
| `client_name_matcher.exact_values` | [client_name_matcher.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-540b25072558891fbc1f8c93bf8641283c30ed646ac647eea88f55be7cd2ab3c) |
| `client_name_matcher.regex_values` | [client_name_matcher.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-c46b8c917351da408e1a314664be2bfd2bfb3ae0609a9e7592e6e4ee7f030922) |
| `client_selector` | [client_selector](resources--service_policy_rule--reference--group-001.md#canonical-508074458c6693764711cb853df217dbee86ee28414773d8ceed90ae3d61e707) |
| `client_selector.expressions` | [client_selector.expressions](resources--service_policy_rule--reference--group-001.md#canonical-212b04199feb77daf60397a0f83757dd43b995658c6864df1dc152af7a018a38) |
| `cookie_matchers` | [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-731f7a55fcfed6015f8d07d61d9b070b62205751b67158e1b8e9a5e604ca98d0) |
| `cookie_matchers.check_not_present` | [cookie_matchers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-88ef488d54ef170fd15ac26cabeb3157f0a1b21165c268fff49367c733a4a674) |
| `cookie_matchers.check_present` | [cookie_matchers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-d0031653dd7c370c4c966f2ce1248e6895b86bc13f53dc899d13fafe3e96d28e) |
| `cookie_matchers.invert_matcher` | [cookie_matchers.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-06cd36205575267943a9d7aefc8cd157c8f834a00d10b1ed37b19b91f77f0cbb) |
| `cookie_matchers.item` | [cookie_matchers.item](resources--service_policy_rule--reference--group-001.md#canonical-de9b6a8dbc50b0a6793eacc2652a0e6edb964d9f19e6e822deb999e3f8ebde73) |
| `cookie_matchers.item.exact_values` | [cookie_matchers.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-14e3985a3301a43b9e8a111911565d2ef443dae2b6f98087556a4976d395f2a7) |
| `cookie_matchers.item.regex_values` | [cookie_matchers.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-3a0de9359efe760ee3dc5cd9fb01bb8b4e0bc4d89ae78a9caf9327f542abd605) |
| `cookie_matchers.item.transformers` | [cookie_matchers.item.transformers](resources--service_policy_rule--reference--group-001.md#canonical-fb9cd88bf9970d417d1d0555989c116d114e28f478f9f9a83283b5c313667d2e) |
| `cookie_matchers.name` | [cookie_matchers.name](resources--service_policy_rule--reference--group-001.md#canonical-219fdd8ab7a887ab05cd5c97c3db5b8fe46be06458c9a45355f1966cb28d066b) |
| `description` | [description](resources--service_policy_rule--reference--group-001.md#canonical-fe006e21674de84102e233c2b9f76aacf051b65ca6eb96840f9389de6a3d21b0) |
| `disable` | [disable](resources--service_policy_rule--reference--group-001.md#canonical-79230ed586a89a0132afb38b93453512d9351bfc9579711f6d051d80070d60ff) |
| `domain_matcher` | [domain_matcher](resources--service_policy_rule--reference--group-001.md#canonical-cdd4331f0a134618a66f111094d0ec997c72ddd6e8a96574127e1ab83555de79) |
| `domain_matcher.exact_values` | [domain_matcher.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-c55e84ad4f921e7dd31966e6bbbc4997658e444a2ecca6b2c06ba58f7c864a1b) |
| `domain_matcher.regex_values` | [domain_matcher.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-ac9403e2b0064e7593de15d158f2020f6ed374f47f46a1886927aab8bafc62d5) |
| `expiration_timestamp` | [expiration_timestamp](resources--service_policy_rule--reference--group-001.md#canonical-1427758097d7ed42defbed3d880163985626f9613135fe342843cb922b22643e) |
| `headers` | [headers](resources--service_policy_rule--reference--group-001.md#canonical-d0b74e79562f149eb94c59c2295718ac38b374e646b1015c449f821c7d63a275) |
| `headers.check_not_present` | [headers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-1bac64f166a5ad124a42a06a751ad96c2a23b2e4ecf54ea467a2726422226a69) |
| `headers.check_present` | [headers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-fd476633b87fac8dd6d1c9dc5833674b3814130094235adab0c94f34a630d179) |
| `headers.invert_matcher` | [headers.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-21eaff1127d61f19a4c87c798cb2a711351c19928b4071fa09d4aacc3c93a8a3) |
| `headers.item` | [headers.item](resources--service_policy_rule--reference--group-001.md#canonical-c2da0ba0fcf3e97f8c7eb74282aea3b8c2c66237bc8edef2fa30daff26154858) |
| `headers.item.exact_values` | [headers.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-0015cac2002affc20fc573583b8a9578ca7a772a06bc96cf368513e2e1e7de2b) |
| `headers.item.regex_values` | [headers.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-5b282d7a1c059a941b63fa727987be988bda322338d482e14af8cabaf660f97c) |
| `headers.item.transformers` | [headers.item.transformers](resources--service_policy_rule--reference--group-001.md#canonical-a4ea4cd98fe17e432dfa7f65131e767d34f746b23513b806c5773128063e2533) |
| `headers.name` | [headers.name](resources--service_policy_rule--reference--group-001.md#canonical-fcaa4c20b655063118afeacf17bd3efe1775f053c8089c7fe53f91705c11997a) |
| `http_method` | [http_method](resources--service_policy_rule--reference--group-001.md#canonical-7679c52108773459414d3c53b5c742b4f40b95503476788672bbee67d7e72bbf) |
| `http_method.invert_matcher` | [http_method.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0209d45cca158b1d9304b9a40ab747404f823d7e436e4fcc93a96fd82417dfb2) |
| `http_method.methods` | [http_method.methods](resources--service_policy_rule--reference--group-001.md#canonical-53789f8e323ecdd856cba93945ce57eca94bd73d8699b37fd4c7335f180716fa) |
| `id` | [id](resources--service_policy_rule--reference--group-001.md#canonical-e1e850d300a3673a52b3b650a96926d20920f8a36d4b5f1de86c09c22fdce34e) |
| `ip_matcher` | [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-62e6ea35a259a4d1cc1564db083dea4875104358518ab57a3ad5372c1c58114a) |
| `ip_matcher.invert_matcher` | [ip_matcher.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-f3c4d3f111de7367f60ec4661911b5e420e4c6355fbfe255f8afd080e1f518dc) |
| `ip_matcher.prefix_sets` | [ip_matcher.prefix_sets](resources--service_policy_rule--reference--group-001.md#canonical-ef1f1b1b5d5e4d9f174582af2187bf8356a4a442b265ccaa80b10b8ac1237561) |
| `ip_matcher.prefix_sets.kind` | [ip_matcher.prefix_sets.kind](resources--service_policy_rule--reference--group-001.md#canonical-c918bf1b669db3e0254d4692d38b21ad334793ce90eb12e18b8a0d78b317ff49) |
| `ip_matcher.prefix_sets.name` | [ip_matcher.prefix_sets.name](resources--service_policy_rule--reference--group-001.md#canonical-cf1c54b1e85802abe03acaa4c91ed23cbe4f3fe548f24837fcd14e105e6f5533) |
| `ip_matcher.prefix_sets.namespace` | [ip_matcher.prefix_sets.namespace](resources--service_policy_rule--reference--group-001.md#canonical-3402f71578ebb5b0addabb79a91689ae443689ff2473ecdbce57652921e39252) |
| `ip_matcher.prefix_sets.tenant` | [ip_matcher.prefix_sets.tenant](resources--service_policy_rule--reference--group-001.md#canonical-666f02d1be6607d3898da1b78007a6f957b0d97e51b2e5adaaf8f53ceb3e764b) |
| `ip_matcher.prefix_sets.uid` | [ip_matcher.prefix_sets.uid](resources--service_policy_rule--reference--group-001.md#canonical-94997a446790b620904587c5ff16c5230d8d5f4be188dc9262ab996f2780dae2) |
| `ip_prefix_list` | [ip_prefix_list](resources--service_policy_rule--reference--group-001.md#canonical-5f120eb5a39bf5998ae353e86ae3d0b7b2fd5ac2d601d68512970f8838346d0c) |
| `ip_prefix_list.invert_match` | [ip_prefix_list.invert_match](resources--service_policy_rule--reference--group-001.md#canonical-5cfdc0a79af2f690be32cef9a00e4815210349545fcf0bf240daa3292ee191c8) |
| `ip_prefix_list.ip_prefixes` | [ip_prefix_list.ip_prefixes](resources--service_policy_rule--reference--group-001.md#canonical-04e3dce9ad1940b64607a31d5cd5f01b4c65203801dde8399a0fcd506e20f3c5) |
| `ip_threat_category_list` | [ip_threat_category_list](resources--service_policy_rule--reference--group-001.md#canonical-32815767088964b8a209c509d3c0680e7e5365edf89016f9ff8d9355507a39bd) |
| `ip_threat_category_list.ip_threat_categories` | [ip_threat_category_list.ip_threat_categories](resources--service_policy_rule--reference--group-001.md#canonical-abc0b874c7d15d717bbca3cf406369555a60d471002e445eb87290bab5efe8e2) |
| `ja4_tls_fingerprint` | [ja4_tls_fingerprint](resources--service_policy_rule--reference--group-001.md#canonical-b858314b40a61f4215a2dc82656112027d75f3e3a83c31c58d8fc7f435447dc3) |
| `ja4_tls_fingerprint.exact_values` | [ja4_tls_fingerprint.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-b8e74593a4d7311274016be5e1975f90c5c8ab0c62776d2b27ee41dfe177d138) |
| `jwt_claims` | [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-2e55d13a4b1be9943bec8c8805bf9f0066edec72cfa3d71700e58a205ab4dcf1) |
| `jwt_claims.check_not_present` | [jwt_claims.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-e50e87a89580f07625bac57cc015e31094c0b4cd2370d938e0a8285fe973649e) |
| `jwt_claims.check_present` | [jwt_claims.check_present](resources--service_policy_rule--reference--group-001.md#canonical-c4bf72353f8d455ed45e629f48294a6eee15fe380329aeaa1b9a08a74e1eafed) |
| `jwt_claims.invert_matcher` | [jwt_claims.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-5132dc611148aa88ac951cafc5436065753b90b12cce50e20e44e7c7f6a463f0) |
| `jwt_claims.item` | [jwt_claims.item](resources--service_policy_rule--reference--group-001.md#canonical-da8dad4d38ac4cf5b32c958060861029b9ef178502b372547649eadae9c71145) |
| `jwt_claims.item.exact_values` | [jwt_claims.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-5d8c89425e1dfa5cbcefc80ff8600c5e89c52dfdbfbae7f281e9f44543311f7d) |
| `jwt_claims.item.regex_values` | [jwt_claims.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-3f96d62da947fb07bea0331838177e1dfaa360c72148dedad510252f19458ae0) |
| `jwt_claims.item.transformers` | [jwt_claims.item.transformers](resources--service_policy_rule--reference--group-002.md#canonical-a14020c707d8861ddb93d22eef6fe3dbf9c9c23207352a9ddef77ed0a3575a0a) |
| `jwt_claims.name` | [jwt_claims.name](resources--service_policy_rule--reference--group-001.md#canonical-6d70639cb1204851a0aa7138bf713e136b3ec9ccb5be54d2f233ebe73d5e1911) |
| `label_matcher` | [label_matcher](resources--service_policy_rule--reference--group-002.md#canonical-760e05075b5d983c33bbc79bdce61f39bf38bce4d424ccc0d72e2925b2dbebf0) |
| `label_matcher.keys` | [label_matcher.keys](resources--service_policy_rule--reference--group-002.md#canonical-b316ea3d9da36d4f504fbc815ee8fc222baff5f65396203b051b00fb3660af06) |
| `labels` | [labels](resources--service_policy_rule--reference--group-001.md#canonical-af02cb4ea4a3d466df78fa00d44d53c01eb8d113b460036d60ba7607422ea2ea) |
| `log_rule_evaluation` | [log_rule_evaluation](resources--service_policy_rule--reference--group-001.md#canonical-ead4ce652f5b40bb98f609b2fc7a5fd3ba8643991f882991355f5c091c996f85) |
| `mum_action` | [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-a96746381ab87215d5dc2b130643029aaf328e776d49cadf3ba42382ce035150) |
| `mum_action.default` | [mum_action.default](resources--service_policy_rule--reference--group-002.md#canonical-925d59f1cc3b8c5a11b829e63a969d20519140c09d18fa424aa595c74d4248f2) |
| `mum_action.skip_processing` | [mum_action.skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-36d25b7c63255352d3bc21986471f3267a2305ddf5633971fc4b1dc9272159f5) |
| `name` | [name](resources--service_policy_rule--reference--group-001.md#canonical-f5f25530ba9a9663beecfc96d008a36e34ff91c2d21a6b6607841760d9ba7e69) |
| `namespace` | [namespace](resources--service_policy_rule--reference--group-001.md#canonical-6cf3ba6a9ddea8886019e28ab3607055c5e8159b9eefbc712c1b892a4a62dd03) |
| `path` | [path](resources--service_policy_rule--reference--group-002.md#canonical-6af93d47ba88d484b4fe57b19fcde6ddd13a83e9beed16060baa2c6f10a54bf0) |
| `path.encoded_path_matcher` | [path.encoded_path_matcher](resources--service_policy_rule--reference--group-002.md#canonical-63e41612ec6d0152f469426f443a8eb3798b555191ba3debcd5dba431bf57d97) |
| `path.exact_values` | [path.exact_values](resources--service_policy_rule--reference--group-002.md#canonical-f094b2c510f0220e29ba80438b7378d7545cb2522353ecda3178cb472e96508a) |
| `path.invert_matcher` | [path.invert_matcher](resources--service_policy_rule--reference--group-002.md#canonical-1a893eeefedc018f0d0ce860739d00295a90e224e9e8602f2c8c907083197f14) |
| `path.prefix_values` | [path.prefix_values](resources--service_policy_rule--reference--group-002.md#canonical-0f76def79fccdab7fca45f761e2a30acefc2dd4c6e0fc2bc439de16e8f1b159c) |
| `path.regex_values` | [path.regex_values](resources--service_policy_rule--reference--group-002.md#canonical-ed8dbf45442704dfd0134b7f960a0e52a7c45d879953722ef190e129a3ec4d79) |
| `path.suffix_values` | [path.suffix_values](resources--service_policy_rule--reference--group-002.md#canonical-06f8ab0d215fd1a6293c348b8c851f70ba8b526e528b34930afaa402db9d2001) |
| `path.transformers` | [path.transformers](resources--service_policy_rule--reference--group-002.md#canonical-bea7491b7392e0c8459bc85a382fa3f93ec48d536b1f8c822711eccc6695d52c) |
| `port_matcher` | [port_matcher](resources--service_policy_rule--reference--group-002.md#canonical-ba598e2ad70c6507d4b690fa29610e739e52ad5f8c95cd922db927fa8b2ed825) |
| `port_matcher.invert_matcher` | [port_matcher.invert_matcher](resources--service_policy_rule--reference--group-002.md#canonical-18a750d040b0da907aec0d8e4d73e96d6454710ce418d272d391201ecee54b09) |
| `port_matcher.ports` | [port_matcher.ports](resources--service_policy_rule--reference--group-002.md#canonical-abe394e30b6a6ec1c2ac8e09e788842c6d72a7821ecee80cdd2c11d18466bf73) |
| `query_params` | [query_params](resources--service_policy_rule--reference--group-002.md#canonical-5e673c0aea73dd35c6a73e2a30cfa90986748ad923cadef6d7f5933f0558e13a) |
| `query_params.check_not_present` | [query_params.check_not_present](resources--service_policy_rule--reference--group-002.md#canonical-e2f016bd5cf28071a1626f2980fcda1e3d1fff9f0c16060d2ecae66660a7f882) |
| `query_params.check_present` | [query_params.check_present](resources--service_policy_rule--reference--group-002.md#canonical-2336808eb9e1aef93c5949fc245bf3b5b00d5c515facb34e533806b6db7f45d9) |
| `query_params.invert_matcher` | [query_params.invert_matcher](resources--service_policy_rule--reference--group-002.md#canonical-8f44b135345aac5290bb604550df17e35d29b00d129ec54173a3c63026c05d62) |
| `query_params.item` | [query_params.item](resources--service_policy_rule--reference--group-002.md#canonical-af4ee6679d50bdfa377f8c985903450f95cdc5fc1e780cb19124e0e7de3718f6) |
| `query_params.item.exact_values` | [query_params.item.exact_values](resources--service_policy_rule--reference--group-002.md#canonical-fb56437d5f6d7654dde992b75ac941a5059c7b71b101eb27fed90c7397d63b96) |
| `query_params.item.regex_values` | [query_params.item.regex_values](resources--service_policy_rule--reference--group-002.md#canonical-43d0e053feb8a6a24d97ab0dcfb6b14ce1e4fc7dc21dc729c0b5ac849b777949) |
| `query_params.item.transformers` | [query_params.item.transformers](resources--service_policy_rule--reference--group-002.md#canonical-212e604288705df5ac7c5a948e7adc1173b355e8f8ea985e7e321ee9acc26a31) |
| `query_params.key` | [query_params.key](resources--service_policy_rule--reference--group-002.md#canonical-90c647f4bfaf6c61dd033d1d7fafc991999fd9998dab93ff22731df6ed017df6) |
| `request_constraints` | [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-f9c2b05098e790f0f50c4e03b2a23eb2ffa42d922bc0c62423baf521f93257f0) |
| `request_constraints.max_cookie_count_exceeds` | [request_constraints.max_cookie_count_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-b7834970258e25db1418339e952bcfe485b8adba324bb364af27776cbdc239f0) |
| `request_constraints.max_cookie_count_none` | [request_constraints.max_cookie_count_none](resources--service_policy_rule--reference--group-002.md#canonical-fc0fa0b3d6153df4b47f004e31597090adc7cd0a57d316d565c2f3212df3d255) |
| `request_constraints.max_cookie_key_size_exceeds` | [request_constraints.max_cookie_key_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-88991a529d4d9b0ee5458c074c39b76f81972f77131815f597a9f074db4203d0) |
| `request_constraints.max_cookie_key_size_none` | [request_constraints.max_cookie_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1a78462c48b944d6e919b2e171462ce231d03d17e4cf7e7137f7bb6f7b0219f8) |
| `request_constraints.max_cookie_value_size_exceeds` | [request_constraints.max_cookie_value_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-4934ef3754d74a2b5caed9a85b114bd708fa066058402ab6b9fc29f41df5e0a9) |
| `request_constraints.max_cookie_value_size_none` | [request_constraints.max_cookie_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1c8b538dc92a8d67a2542c27d67da59cdd43885e6c2c34e28ac5da995c3e2f92) |
| `request_constraints.max_header_count_exceeds` | [request_constraints.max_header_count_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-eb15bcde201e57f87f7670bbc519158d2d13cd55585c940366d20b6e123bba33) |
| `request_constraints.max_header_count_none` | [request_constraints.max_header_count_none](resources--service_policy_rule--reference--group-002.md#canonical-2d5f318525a2ac86604417b6c41297a4effb06aa3a256c04308d94cbc6924277) |
| `request_constraints.max_header_key_size_exceeds` | [request_constraints.max_header_key_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-10a116c4d617b294577631c1fa31b80c95c894779f8775b6336ade57576c0956) |
| `request_constraints.max_header_key_size_none` | [request_constraints.max_header_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-23d7f11367093bd4fe0b1a10134fe20b01b889ed1e441edb0989f288c2399ace) |
| `request_constraints.max_header_value_size_exceeds` | [request_constraints.max_header_value_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-8acda1972bf9a58cabdb3f5a58e8d09b3ac4606f7d92edd7428a32aa3056d147) |
| `request_constraints.max_header_value_size_none` | [request_constraints.max_header_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-fe10e0da82d0da1620ede5c32e935ebe52237c6d1086abdc575d9e31db9387ae) |
| `request_constraints.max_parameter_count_exceeds` | [request_constraints.max_parameter_count_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-6f3238576fa36484c7fe387a68d844ecea249a8d924fddae267fee4309160780) |
| `request_constraints.max_parameter_count_none` | [request_constraints.max_parameter_count_none](resources--service_policy_rule--reference--group-002.md#canonical-54fec6f536eb9d9a7c95d6e104e1e56c15bba57609113eff56a4689f735b8fe1) |
| `request_constraints.max_parameter_name_size_exceeds` | [request_constraints.max_parameter_name_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-08fbe2fa9b079c2f14fecab50e37a502524564d64fe31a4f7c3a26a8728da845) |
| `request_constraints.max_parameter_name_size_none` | [request_constraints.max_parameter_name_size_none](resources--service_policy_rule--reference--group-002.md#canonical-6beabc2f1c5c9f7cbb876a1e7a2b9d1422b93631466e6a23452e9e89da3e8c65) |
| `request_constraints.max_parameter_value_size_exceeds` | [request_constraints.max_parameter_value_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-0ce1c05cffb71f89b89d255fbed9d105b830841ebf24eae9e17acab772f1156c) |
| `request_constraints.max_parameter_value_size_none` | [request_constraints.max_parameter_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-7bdec7e82a48311c9c0358d1a98be173b3e4dc91f52b3a0918e63ff8699da708) |
| `request_constraints.max_query_size_exceeds` | [request_constraints.max_query_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-eb244083c53ecfd9debecc5e1f0742e9fd4395cc6f65591f82df9d22fcfbd07e) |
| `request_constraints.max_query_size_none` | [request_constraints.max_query_size_none](resources--service_policy_rule--reference--group-002.md#canonical-63ba896825420c330a7ca3f560376955ba1eddd42cb8317f50a70ad6fbaae114) |
| `request_constraints.max_request_line_size_exceeds` | [request_constraints.max_request_line_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-9528912dcc8fad9d1fa851cfac4689f2bb5a04379a63f7cdc8ffd4eadf1b113e) |
| `request_constraints.max_request_line_size_none` | [request_constraints.max_request_line_size_none](resources--service_policy_rule--reference--group-002.md#canonical-cf14153bda216fbc2c93900467bc0fd5241339dba3a1c36609462dab422778e9) |
| `request_constraints.max_request_size_exceeds` | [request_constraints.max_request_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-c50de14a1cd202da88934803d55dc1c440aa4cc7b2abd9857b6d8b99b3444b03) |
| `request_constraints.max_request_size_none` | [request_constraints.max_request_size_none](resources--service_policy_rule--reference--group-002.md#canonical-5426e810c9d4ab01fdb5640f7e530a9b2a2a124e8ec32cd2e0e15e1fcc287b37) |
| `request_constraints.max_url_size_exceeds` | [request_constraints.max_url_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-c2f9271e4b4a722ff1e77173dcd75c8965369e21f61dc0d8c122cd2a1544c64a) |
| `request_constraints.max_url_size_none` | [request_constraints.max_url_size_none](resources--service_policy_rule--reference--group-002.md#canonical-e691a6c709605fd6bf2d448a4f89717c7ec2733b5a7b1faebf080bc0bbfd4e96) |
| `segment_policy` | [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-00d30008ccdbf9b04fdb050086064d49ff6f6e9cd059d638fa7b75607ebd845d) |
| `segment_policy.dst_any` | [segment_policy.dst_any](resources--service_policy_rule--reference--group-002.md#canonical-bee6e6af73eea55968109fad8d9789bc81191d5c344a54a43aaf6a0ca9b35eb7) |
| `segment_policy.dst_segments` | [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-d71707b3b9c1a09888ee79454d2819ac9a17ef0ad587841fc2e2aaca212bc5fb) |
| `segment_policy.dst_segments.segments` | [segment_policy.dst_segments.segments](resources--service_policy_rule--reference--group-002.md#canonical-a21e47c705b45b3cd504b78b408415d06de0dacc51120cb5892555321cd91040) |
| `segment_policy.dst_segments.segments.name` | [segment_policy.dst_segments.segments.name](resources--service_policy_rule--reference--group-002.md#canonical-0d04f116ff42a697bf8959fe7e3135e1571f30d830957018ebc934a837c4fb5f) |
| `segment_policy.dst_segments.segments.namespace` | [segment_policy.dst_segments.segments.namespace](resources--service_policy_rule--reference--group-002.md#canonical-27e9c498fb263e23682a08d045ecbd673e9ab3141db5a5f714a367a9f1e6d371) |
| `segment_policy.dst_segments.segments.tenant` | [segment_policy.dst_segments.segments.tenant](resources--service_policy_rule--reference--group-002.md#canonical-a9837cb0b771f0cc01ad332311761a46053bb4d35e5896c8adfdfe8d96febfc0) |
| `segment_policy.intra_segment` | [segment_policy.intra_segment](resources--service_policy_rule--reference--group-002.md#canonical-2c7b775b1b38ffe563b94954df6dc1d33ffdeda3f7d37d521f249064e800f365) |
| `segment_policy.src_any` | [segment_policy.src_any](resources--service_policy_rule--reference--group-002.md#canonical-9fc2e32cf2dec5f7ac59fe587e9437dede8f5b0539b9e8fdb37ea319eb76bad0) |
| `segment_policy.src_segments` | [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-44d5b0b5c5aeaccd04042c3346814703f68c3e261ea03a82c4f247228bd3ebfd) |
| `segment_policy.src_segments.segments` | [segment_policy.src_segments.segments](resources--service_policy_rule--reference--group-002.md#canonical-cfa1d3f129a3e6212c669a54703c4f78ed62f0f86fd36bb9673a43fb3e1295e7) |
| `segment_policy.src_segments.segments.name` | [segment_policy.src_segments.segments.name](resources--service_policy_rule--reference--group-002.md#canonical-49c0c26a784da1e28a6c324d2827450d7620da81078f6c723e3781de0dffb3a7) |
| `segment_policy.src_segments.segments.namespace` | [segment_policy.src_segments.segments.namespace](resources--service_policy_rule--reference--group-002.md#canonical-b7fb798e441421c20357f0973c3b92d37963aeb00f222876bdb00b8d63138b7d) |
| `segment_policy.src_segments.segments.tenant` | [segment_policy.src_segments.segments.tenant](resources--service_policy_rule--reference--group-002.md#canonical-31b4fe1d337f8c60095e42b0336c5491ca3b1e7e6927997b7c4c4493559a8cb9) |
| `timeouts` | [timeouts](resources--service_policy_rule--reference--group-002.md#canonical-e8d4a4db8c37bbc139975eac7fab78a32fe33ee7dff1c263ef6de7c838963d40) |
| `timeouts.create` | [timeouts.create](resources--service_policy_rule--reference--group-002.md#canonical-2afdc9249a3aadf139d1bc74edcaa4434ca595129d225216d6c1d5e68c8dc570) |
| `timeouts.delete` | [timeouts.delete](resources--service_policy_rule--reference--group-002.md#canonical-f425c20ce69983bf50624410f75ebdda2a1e4304250e2c38651074936a2d02af) |
| `timeouts.read` | [timeouts.read](resources--service_policy_rule--reference--group-002.md#canonical-c776f046ac6043d6fdffa17ecd30232d6ab9ea2f0bd34068269628168bcc45d1) |
| `timeouts.update` | [timeouts.update](resources--service_policy_rule--reference--group-002.md#canonical-5829707d9ce173638483a737b2d61f9bb6c546ff5be9634e5a038152bc7604f9) |
| `tls_fingerprint_matcher` | [tls_fingerprint_matcher](resources--service_policy_rule--reference--group-002.md#canonical-b02fed0ec6f83276f5054e82dee19e02e679bcd5a8c478e201220c87e66bfe6b) |
| `tls_fingerprint_matcher.classes` | [tls_fingerprint_matcher.classes](resources--service_policy_rule--reference--group-002.md#canonical-a75c1e7eac11e7a1dd8191bad45ab6214a8e4d4b56db2b4b4e7a1e114ee53245) |
| `tls_fingerprint_matcher.exact_values` | [tls_fingerprint_matcher.exact_values](resources--service_policy_rule--reference--group-002.md#canonical-35b1c82d4f4362107a7bf3d8210cfd344b47a7293dadbc7e7e14dd227543a04d) |
| `tls_fingerprint_matcher.excluded_values` | [tls_fingerprint_matcher.excluded_values](resources--service_policy_rule--reference--group-002.md#canonical-d8179763066ddebb518f5499e7a6990432d3f04d524284a500265b724fc20fbb) |
| `waf_action` | [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-165aaa4d7dd235e1809f7ecc1c949e698490430debc8667276bc964700f63892) |
| `waf_action.app_firewall_detection_control` | [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-b222bf073801535e0ee7715168a722014bd0ad58e8e9dcb296e2f52e610e93cb) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy_rule--reference--group-002.md#canonical-0141ea8fef8d5cde72450c9f0162648107f851fafe45ace0a9f860147ec4ac5a) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](resources--service_policy_rule--reference--group-002.md#canonical-565b7dbaf3dfb9b0402e314d970beb67d15c02018fd6a5a070756d47fd256c41) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](resources--service_policy_rule--reference--group-002.md#canonical-4249d3c942ed9b3ae919d149b40e9ee09aa613c909b53b1c3cd5c62e94c1294f) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](resources--service_policy_rule--reference--group-002.md#canonical-47ffff15f1fbfdc13e11a052bf4b3e6a98872b5572cb6869ae28b80c415e261d) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy_rule--reference--group-002.md#canonical-ad676add7878cdbfa7120a9be8c2910d7e2654dffab880b47c421344cb73b732) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](resources--service_policy_rule--reference--group-002.md#canonical-79c52146348d21bf97af927ed738af07e5058446bf1cc7e07c364f0344fbd9d9) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts` | [waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy_rule--reference--group-002.md#canonical-82b99c25eee77b989b8d931a54110abce86f7d2d637a08b5a21c8f2f3ca065a4) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context](resources--service_policy_rule--reference--group-002.md#canonical-32f70242a2c1d1daf64b9fe67df3798f23366006d6b1860df342786e4039fccc) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](resources--service_policy_rule--reference--group-002.md#canonical-2590fee4736265af2844c871416ae80a3c70bd10bbf96197c9eef8fad10472ea) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](resources--service_policy_rule--reference--group-002.md#canonical-73bfcbc190821141dfa07aa31d428f86d7947e608deb4e4fe60d4175176d9acb) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts` | [waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy_rule--reference--group-002.md#canonical-829d489cb333d4c1b7dbb00179b4adce85e82d1e9465e95aac54dbfad01cf079) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context](resources--service_policy_rule--reference--group-002.md#canonical-a77b6e7c95478cdeb4570343f550b560877f331f00947c695779d95719845615) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](resources--service_policy_rule--reference--group-002.md#canonical-365da9ec4d8481be8508de9fbae7b046fd8253f805f49c9d2fc450ede799142d) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](resources--service_policy_rule--reference--group-002.md#canonical-3f522efa589cdb4cfefa9745324ae230863c08545c41ffd6642c4c93e1012b0b) |
| `waf_action.none` | [waf_action.none](resources--service_policy_rule--reference--group-002.md#canonical-23dc649e6ea1d40792de0b55a22c583d39035e7d969b0c7bb3de265023dc7e12) |
| `waf_action.waf_skip_processing` | [waf_action.waf_skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-f70a698546ade0558fad984e3f0d5f4d7b92ccc4d7e29e09b0804699138e21b5) |

<a id="canonical-4d52352139d12e6e947280f3132f5b0e988faa7456c062ad14435cb07fdebb50"></a>

## Next pages — Property reference / 08f531c4f6cb / 16

- [any_asn](resources--service_policy_rule--reference--group-001.md#canonical-29bacf87798d1bc903b9ed87078c261fb07693a10b86f3252e680ec149ce9124)
- [any_client](resources--service_policy_rule--reference--group-001.md#canonical-1b190696041b8893594e5e23d2300f111e413977553cc089857c9fb5bc1ead34)
- [any_ip](resources--service_policy_rule--reference--group-001.md#canonical-bb3ff9935dbfdd859eff4a15408750eee13afb0315c3936d0080835ee6462570)
- [api_group_matcher](resources--service_policy_rule--reference--group-001.md#canonical-8181b58d11566e2d3bea22f44f5becd3bb2e2b9c68f63c0de829acd0c755fbdc)
- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94)
- [asn_list](resources--service_policy_rule--reference--group-001.md#canonical-09fa9b0f03a116e8889aab4532d7ee5dcb87005230bb50915d9e937361308ec4)
- [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-737775f18217586217b199bb222bdcd58421af7102d5cf785545f7621ff6f6a4)
- [body_matcher](resources--service_policy_rule--reference--group-001.md#canonical-41478796dc9f16cc362d425250c15e826413f8248849eadf86e9eab58399f5e1)
- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-936c6a7953e5f16177b88a818686a49d30aa935b9f8caa61e07bfcbe6b11ae52)
- [client_name_matcher](resources--service_policy_rule--reference--group-001.md#canonical-3e6fcc4950ad994ad911391dcf12a4305b0b9c75569d289775436104117fd307)
- [client_selector](resources--service_policy_rule--reference--group-001.md#canonical-70d913f0f8a715b2ff59d453d1f84bd6685e80d1a6caa80a9fef7d9e8dc7195c)
- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f)
- [domain_matcher](resources--service_policy_rule--reference--group-001.md#canonical-7ea9c15f3bbe81e53c56ac87cbe279b4bf1e79773e794059e812c84b8e4e2be4)
- [headers](resources--service_policy_rule--reference--group-001.md#canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480)
- [http_method](resources--service_policy_rule--reference--group-001.md#canonical-f5affe06710c1b7d6014ccf515752fa30f078d58d20a6cf7e719436ab7db52ae)
- [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-5131505e9c5d1806673eebf13dd26429a151fd79793834461f68747727fecbf6)
- [ip_prefix_list](resources--service_policy_rule--reference--group-001.md#canonical-413521cec6a72c537f18e294f798d20fdf3fef19b519f823e0f7d2e09417ab79)
- [ip_threat_category_list](resources--service_policy_rule--reference--group-001.md#canonical-06cb5e2962be24ee85db336adacd12583343c5a13b249653673224a7757d992f)
- [ja4_tls_fingerprint](resources--service_policy_rule--reference--group-001.md#canonical-aed76a8d199a5769ebe931308623edc55ce47695ebacff7f084193b3f4a9709f)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397)
- [label_matcher](resources--service_policy_rule--reference--group-002.md#canonical-885654158fb78e99195654bfd7eb4ede70c6b02f2d395cef3f137f6410b798c0)
- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-8173a7b12b6dbcabe15c5a24896b2850104c373bdd419858980760f784308f9d)
- [path](resources--service_policy_rule--reference--group-002.md#canonical-c453af0b534656a15fed7afe77bb9ad4115339a74c815501323d7c06337a777c)
- [port_matcher](resources--service_policy_rule--reference--group-002.md#canonical-a3aa60eda2c10f207494ee2277014d7f3dde42d696ff2dcc542c02648fd6d21f)
- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- [timeouts](resources--service_policy_rule--reference--group-002.md#canonical-49d1dcedd566f98a405b94194bc73262f7147467a3f23b9ded4df6aaafb63e75)
- [tls_fingerprint_matcher](resources--service_policy_rule--reference--group-002.md#canonical-a00396ea59b4ee1bd6e1ea84d1322f649375d00c7f7360e57c8af940324e416b)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-29bacf87798d1bc903b9ed87078c261fb07693a10b86f3252e680ec149ce9124"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b65c797184100d87207b2bc8a1c96a72dff3965045fd1bd49d6df73d43f631fd"></a>

## any_asn — any_asn / e6b3ab2d2cdb / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- any_asn

<a id="canonical-89b8b4c875a6deb8aa66d324fe177e5402017a5d46c2255ab9ecd73b930658aa"></a>

Type: `["object", {}]`. Optional.

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

- [any_asn](resources--service_policy_rule--reference--group-001.md#canonical-89b8b4c875a6deb8aa66d324fe177e5402017a5d46c2255ab9ecd73b930658aa)
- [asn_list](resources--service_policy_rule--reference--group-001.md#canonical-921e189a1413d60c1e3e453034432ee30e83bfc13f351e7cce703ec58e6fa8aa)
- [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-6949d079257b07f0ce687309fc1bc8f855c9638a840d3113e6083b728e5a7394)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_asn = {}
```

<a id="canonical-2ccc4f766353ca0d9960c280d6fbf5b5f31f07f034f58a3f0bc59f80ed25db2a"></a>

## Direct properties — any_asn / e6b3ab2d2cdb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6605cb8491b8df85b70d789a858c58af56962173ee171f2f7608b2a0b52ad8b1"></a>

## Next pages — any_asn / e6b3ab2d2cdb / 4

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-1b190696041b8893594e5e23d2300f111e413977553cc089857c9fb5bc1ead34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-823c4cc2e1c4ef78dd96ec160307b5d4c6c712e1b77b2f3a448c8e057dec791e"></a>

## any_client — any_client / ef5be1b52db5 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- any_client

<a id="canonical-3e6e098d9f0c2d679e2aa75ef9171a69a1451766e9dad02e49a69b38fd6f6b04"></a>

Type: `["object", {}]`. Optional.

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

- [any_client](resources--service_policy_rule--reference--group-001.md#canonical-3e6e098d9f0c2d679e2aa75ef9171a69a1451766e9dad02e49a69b38fd6f6b04)
- [client_name](resources--service_policy_rule--reference--group-001.md#canonical-e42a2762a0028e79d9947feeb6fd03ae69a307e752948729a0752e765ddf7c38)
- [client_name_matcher](resources--service_policy_rule--reference--group-001.md#canonical-56ffd1528ef78daa1c36b951f5930c56bb9d8518e6f213f9573bcef765ab2229)
- [client_selector](resources--service_policy_rule--reference--group-001.md#canonical-508074458c6693764711cb853df217dbee86ee28414773d8ceed90ae3d61e707)
- [ip_threat_category_list](resources--service_policy_rule--reference--group-001.md#canonical-32815767088964b8a209c509d3c0680e7e5365edf89016f9ff8d9355507a39bd)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_client = {}
```

<a id="canonical-2f14c711829b3518f071e961f7f6de702f1e50ef63a1f39645b42b381d734cbf"></a>

## Direct properties — any_client / ef5be1b52db5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9b93358b386cec091eb02628bc8a485bfa472b8d3d758c0331c8bffeb38a9591"></a>

## Next pages — any_client / ef5be1b52db5 / 4

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-bb3ff9935dbfdd859eff4a15408750eee13afb0315c3936d0080835ee6462570"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e662211dab0331819ad5d8f6bbc1e3d2719ee1f136a250025ea6b3f860a0fb3a"></a>

## any_ip — any_ip / c4a50382e5f7 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- any_ip

<a id="canonical-be87293db2b603e2353fd6ff5be02c5a124868c1ef4814a4aff4df8afbb6043d"></a>

Type: `["object", {}]`. Optional.

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

- [any_ip](resources--service_policy_rule--reference--group-001.md#canonical-be87293db2b603e2353fd6ff5be02c5a124868c1ef4814a4aff4df8afbb6043d)
- [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-62e6ea35a259a4d1cc1564db083dea4875104358518ab57a3ad5372c1c58114a)
- [ip_prefix_list](resources--service_policy_rule--reference--group-001.md#canonical-5f120eb5a39bf5998ae353e86ae3d0b7b2fd5ac2d601d68512970f8838346d0c)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_ip = {}
```

<a id="canonical-b3d29411e0fd6c736de06194fe78a2cd61d31e87445179601c06dc8d17c66adb"></a>

## Direct properties — any_ip / c4a50382e5f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8cfd98743bc3ff2db7d2b93bb741a270005fa1979b36e3e7603f37d426728b50"></a>

## Next pages — any_ip / c4a50382e5f7 / 4

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-8181b58d11566e2d3bea22f44f5becd3bb2e2b9c68f63c0de829acd0c755fbdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a5c3e576e9dc43b385965584024ef68f9fe73f00d6d7a646bd491ca3a5001b1"></a>

## api_group_matcher — api_group_matcher / 3c957acd3220 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- api_group_matcher

<a id="canonical-08c17e728005ad0f20d0e9c5242f5d794473e394872aaffeef9346b0f54470fa"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies a list of values for matching an input string. The match is considered successful
if the input value is present in the list. The result of the match is inverted if invert\_matcher is
true.

Upstream description:

A matcher specifies a list of values for matching an input string. The match is considered
successful if the input value is present in the list. The result of the match is inverted if
invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("match")}
```

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

Terraform syntax:

```terraform
api_group_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ffdd2a708e39d3ff142cd748047185257e07f4e0f144a095ae2d03d0236f1d9"></a>

## Direct properties — api_group_matcher / 3c957acd3220 / 3

<a id="canonical-36e25bdf3721d7e428f63c80a17213208c4d051e5cee977dd8f5bca466eed763"></a>

<a id="canonical-2d2e7d979947bca2387de4dc261f122ccbc8af352cc00d1014904b7cad803aa6"></a>

## invert_matcher property — api_group_matcher / 3c957acd3220 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-1b1453015df43a83fc69d090045282b2fec2cf8f16be44297daf99f56beef818"></a>

<a id="canonical-5f76eebae0ef245a1ecb896a0bddc272ab2bd92f7250e782786c80650ac1ef13"></a>

## match property — api_group_matcher / 3c957acd3220 / 5

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-66da1be601fd23fce58a4bbbae74c331810c6ac5bfb7d528de8de461ce05d735"></a>

## Next pages — api_group_matcher / 3c957acd3220 / 6

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-191362a64aaff58bf09a494b8df03671c92e4ed6d685ce6deb06566fac4bb26d"></a>

## arg_matchers — arg_matchers / 41c10bd2e0c8 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- arg_matchers

<a id="canonical-6a00b2f13c0a10ac6a6f95f84ce38ab9ca85b90efb152df3cda82456fc9b5791"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
arg_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-884abba80f522d24510e9396ff656ea8f401f836737d4a6d149ff070a9448cb4"></a>

## Direct properties — arg_matchers / 41c10bd2e0c8 / 3

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-69562d7637d9952ecea514624ca840172bbe54831a5a4d4b0faeb62f81d6f6ce): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-4186d44513e075434596a435c2c7ae16132b2b41aec66fc9f063267fd53e0aa1): complete subsection reference.

<a id="canonical-b2a52e82c842a73e123c6e0f7523a8e360e6975003633322f0bde65564416103"></a>

<a id="canonical-8a51099f784f5f9c1146a21296b4a43fa086826be41896bab19b535c441a3101"></a>

## invert_matcher property — arg_matchers / 41c10bd2e0c8 / 4

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-6baef29b73bda89f75ba39c47fe78e07d0cf512295463e8733303caace45d06b): complete subsection reference.

<a id="canonical-ad89d2e6433ee0cfee3fa906ec87ae8b4f04e26161467960268a446f43d71d62"></a>

<a id="canonical-2b26d35b67cf61b96596c232f854e4d8de7019cf572ca7f7e7c82f9bf60d0eda"></a>

## name property — arg_matchers / 41c10bd2e0c8 / 5

Type: `"string"`. Optional.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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

<a id="canonical-a9eb5e916a693a94efec87eef28f248a72493c590efd318f877fcc12f0d35197"></a>

## Next pages — arg_matchers / 41c10bd2e0c8 / 6

- [arg_matchers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-69562d7637d9952ecea514624ca840172bbe54831a5a4d4b0faeb62f81d6f6ce)
- [arg_matchers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-4186d44513e075434596a435c2c7ae16132b2b41aec66fc9f063267fd53e0aa1)
- [arg_matchers.item](resources--service_policy_rule--reference--group-001.md#canonical-6baef29b73bda89f75ba39c47fe78e07d0cf512295463e8733303caace45d06b)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-69562d7637d9952ecea514624ca840172bbe54831a5a4d4b0faeb62f81d6f6ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31a27e6f3fea358d5514bc99b6feedc1fe821ac2c58c845b950cc0b9d870ddb4"></a>

## arg_matchers.check_not_present — arg_matchers.check_not_present / 3158d2b052bd / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94)
- arg_matchers.check_not_present

<a id="canonical-21d21899231beb535e28ee76d26fff180758552339fbbad1f5828a10285de744"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-1b04aa76a71eda594d04a02230cb10fad46f2ad21d1df2fda7b34b2f6e314f74"></a>

## Direct properties — arg_matchers.check_not_present / 3158d2b052bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d19b4821f75b12b2dcb70342cfcd2d3ab3c53a8b956a7edfca3ef760d1c9504"></a>

## Next pages — arg_matchers.check_not_present / 3158d2b052bd / 4

- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-4186d44513e075434596a435c2c7ae16132b2b41aec66fc9f063267fd53e0aa1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbcbe0767ddc152348d25509d46b9c4b81c213d9c3dfa83852a0da33b67f01da"></a>

## arg_matchers.check_present — arg_matchers.check_present / b6600e5a4846 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94)
- arg_matchers.check_present

<a id="canonical-1161fb80d13e3e33ea406481777e3cb67e839cf2f6cb4ef188eb6def55db6a78"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-21923b93a994f0f2f10a11e8eb39ee7a7fd972699a86cf4d763b0035b64c1f19"></a>

## Direct properties — arg_matchers.check_present / b6600e5a4846 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cbb66dac38484b307da8a8391afc96025616397f165132968a3558aa162bc682"></a>

## Next pages — arg_matchers.check_present / b6600e5a4846 / 4

- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-6baef29b73bda89f75ba39c47fe78e07d0cf512295463e8733303caace45d06b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8519d741f7ee8ae837d90e4e2cf67bc3728e52f51d403ef17e244a7c4501ca0b"></a>

## arg_matchers.item — arg_matchers.item / 13bffe692ffe / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94)
- arg_matchers.item

<a id="canonical-89e5bba94ccd8768818101d5af93915fca2c29e14b4d3179d97513a6b1ff6d42"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-6cfa8102004570e137672aa1fc6b250edbcb720cfd149490f974f73da6b953b8"></a>

## Direct properties — arg_matchers.item / 13bffe692ffe / 3

<a id="canonical-f9dbd40d4d6008ab516baef9e15696a6ecdad1c5f86f87f3f19d158cbc92ceb6"></a>

<a id="canonical-aecc275d9b3c4b35e6dfc38291958724cdd49cbe090809989ace36b929acbe3f"></a>

## exact_values property — arg_matchers.item / 13bffe692ffe / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-fe76e1f86817b2feb030b03ec047922d53c7717ac4e7ac5c702a160315cd9d2b"></a>

<a id="canonical-090ef99ff72929058f3e8fb2d6d64cdcd050cc586d1ea08bd591d50f2e507a10"></a>

## regex_values property — arg_matchers.item / 13bffe692ffe / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-e9a351d5526941c0c72d18b075229c7f4596a639c0670060a8a69774a456b5bb"></a>

<a id="canonical-5e9246261a1727340960866b4a553519cc80cb230014d57f4e548a1fd71ac78f"></a>

## transformers property — arg_matchers.item / 13bffe692ffe / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-0c4343ec69ce93e8ce9ed8b1f82887e73c249ee3c747419c6db868f76427f6d1"></a>

## Next pages — arg_matchers.item / 13bffe692ffe / 7

- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-4e2d32722bf0ed73a3e24f0543a32f9ce83412888e1ca64633bd34ab8dee4e94)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-09fa9b0f03a116e8889aab4532d7ee5dcb87005230bb50915d9e937361308ec4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-054db77827dc45ec6c80216246d632dd1e1573d53408f4b0c379b6849b9ccf75"></a>

## asn_list — asn_list / a85d12c8fefd / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- asn_list

<a id="canonical-921e189a1413d60c1e3e453034432ee30e83bfc13f351e7cce703ec58e6fa8aa"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
```

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

Terraform syntax:

```terraform
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2f05be9b21f3a34f7271b21902f752c956fbfd4368fbd8e3883388f6fe9dae0c"></a>

## Direct properties — asn_list / a85d12c8fefd / 3

<a id="canonical-73037715a40f7109f627140c03ee6e9f375bc95455d4ea992723794c993c3cd8"></a>

<a id="canonical-a3f46f438fe1357224d5869a8ce18ca257def0677b9cb1747859481c3a9ee03f"></a>

## as_numbers property — asn_list / a85d12c8fefd / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-cc50c1eb2864d9a02f84882d99f768c84108693e7341a980d286dc3e4bfa2b8f"></a>

## Next pages — asn_list / a85d12c8fefd / 5

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-737775f18217586217b199bb222bdcd58421af7102d5cf785545f7621ff6f6a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d048917954e192937b33d89fcd6deafdf3838acbdad30b244e02983ca1420e46"></a>

## asn_matcher — asn_matcher / 455fb31531f7 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- asn_matcher

<a id="canonical-6949d079257b07f0ce687309fc1bc8f855c9638a840d3113e6083b728e5a7394"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
```

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

Terraform syntax:

```terraform
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2500c63d47a92c5e2958c87ce44e72ed98dbed6d4ccf50996dd8d40c04e5ffd2"></a>

## Direct properties — asn_matcher / 455fb31531f7 / 3

- [asn_sets](resources--service_policy_rule--reference--group-001.md#canonical-f40160d07f4143d95f3b33a6489c711abbbd1bf174f231040f66494f69afec37): complete subsection reference.

<a id="canonical-a0d56b210f0afcac64db71e9789f9753d55f22fdc868f105ec4dcc97d24e3c98"></a>

## Next pages — asn_matcher / 455fb31531f7 / 4

- [asn_matcher.asn_sets](resources--service_policy_rule--reference--group-001.md#canonical-f40160d07f4143d95f3b33a6489c711abbbd1bf174f231040f66494f69afec37)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-f40160d07f4143d95f3b33a6489c711abbbd1bf174f231040f66494f69afec37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d524ad0fbed77fe4592a5aaf94f76e5c0c3eee929017536ab7a2af8545e11661"></a>

## asn_matcher.asn_sets — asn_matcher.asn_sets / c69d2dd96ab2 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-737775f18217586217b199bb222bdcd58421af7102d5cf785545f7621ff6f6a4)
- asn_matcher.asn_sets

<a id="canonical-32334d5527daec2e4f98e5647f509967663f940846113d42f2f56a22dccea6ab"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-99eb60194f47bdd4e086baa5f84d00cb800c1c624b960c4d46a69145bbd9abd0"></a>

## Direct properties — asn_matcher.asn_sets / c69d2dd96ab2 / 3

<a id="canonical-966e8a696cc65276355423426e39ba7b9562f23c1e3cde3e58e0f6e70101caea"></a>

<a id="canonical-167dfbf6d8105a5c090967a8f478ccb4759ccfdaa3fb39239ca102584d957e9f"></a>

## kind property — asn_matcher.asn_sets / c69d2dd96ab2 / 4

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

<a id="canonical-a1de0353b6d11047b2b655c3fccd78c5f31391d66ebf3c1ae34f68404ee9ae9d"></a>

<a id="canonical-0d5dddefc2097086d71c1d14ae55f3370db8a2cd6a4838c5fc6a1003a2e99334"></a>

## name property — asn_matcher.asn_sets / c69d2dd96ab2 / 5

Type: `"string"`. Optional.

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

<a id="canonical-64758910f97b47dd6d5fcf972b0f399ff051a1acd202e28e7c1ef5d4aa63e4ea"></a>

<a id="canonical-68572aa03bd11cc71f7633b459fa0bc16165227a0c1ddd35ac37a584a71a68f1"></a>

## namespace property — asn_matcher.asn_sets / c69d2dd96ab2 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-b6aacb140770277934c9554f36a848d7bdaf17aad6377bc3207b03561e4b5158"></a>

<a id="canonical-f9b97a2d5abefc35aa2e6bb6923e2a523576748d7a8c46596eca9cc11622a7dc"></a>

## tenant property — asn_matcher.asn_sets / c69d2dd96ab2 / 7

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

<a id="canonical-9071b2e1412e8aba118abe8879171273ba22e9ce2f103d2e8dace692e03b578e"></a>

<a id="canonical-b9d2fd41bfbb118d4bcee928d395e0f879ea44ac1087853323e97e09b1f55509"></a>

## uid property — asn_matcher.asn_sets / c69d2dd96ab2 / 8

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

<a id="canonical-6821273b92f009199a1a232887316e937cf3a9e530d2e0e25a00f466495b5dbf"></a>

## Next pages — asn_matcher.asn_sets / c69d2dd96ab2 / 9

- [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-737775f18217586217b199bb222bdcd58421af7102d5cf785545f7621ff6f6a4)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-41478796dc9f16cc362d425250c15e826413f8248849eadf86e9eab58399f5e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04c3a2063ece52432a65bdb315265a183f19694c758ee0116f0b25db41b985ae"></a>

## body_matcher — body_matcher / 4c3d8825af4a / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- body_matcher

<a id="canonical-5a2bbc5e6f43b1e6da37ddd219c7d8059fa7523c86bf03a3dba56411cbe9cfe3"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
body_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-748e8efb0f691db7f200110f7b41a342002f3fb39cea81f0c6a3c01443c985e5"></a>

## Direct properties — body_matcher / 4c3d8825af4a / 3

<a id="canonical-b662d9229907c473f47256d4c955783b0694362f0e5bf2b59908996a80af25db"></a>

<a id="canonical-014b60fb3b2d02fa1ccfce6c01f7662f8cb19c9e8172dbaa888f719c7ad0e0d5"></a>

## exact_values property — body_matcher / 4c3d8825af4a / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-ef0506a1cb07786b53af416c149fb8eba0bd5e15e0a75971c27f2ff382ccc76d"></a>

<a id="canonical-71a822c10855eee76e97a9c94898136d23218278ee81543d8e7e86639cc0aff5"></a>

## regex_values property — body_matcher / 4c3d8825af4a / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-97c5604e76d0a46016d776b45b00ddf2ed6bd271eb2117668693b42209f83263"></a>

<a id="canonical-0d6bf50a15a37f482a6eb3107bafe636b6b7eeb2d6555f60daefb56cb6e53ea6"></a>

## transformers property — body_matcher / 4c3d8825af4a / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-5d78d24ed20f8f40f4fd5c36b7a9bf772c0713e7d009e2e59ae32886cb785798"></a>

## Next pages — body_matcher / 4c3d8825af4a / 7

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-936c6a7953e5f16177b88a818686a49d30aa935b9f8caa61e07bfcbe6b11ae52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06b5760eaa85bf6d111457b7717dca2e9a825b47452a196fc6a37c16ef9b7d91"></a>

## bot_action — bot_action / ebb30573cf33 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- bot_action

<a id="canonical-25ffaaff0624f097d7f3168ddfc3ffd3f2d981759fa2d95ba562a47b6fd600d2"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_skip_processing",
    "none")}
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
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

Terraform syntax:

```terraform
bot_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-e86a0e181e2d31d1f322919ba6a69bf4cf76bfd993bfa9886e4520e2abc5b95b"></a>

## Direct properties — bot_action / ebb30573cf33 / 3

- [bot_skip_processing](resources--service_policy_rule--reference--group-001.md#canonical-2fc47a010c757e9509febb91d8398e7bc58cb7917debeed50042cda045cc76a2): complete subsection reference.

- [none](resources--service_policy_rule--reference--group-001.md#canonical-da135dd49e7627c3b5db7532af22094d27677c12713ecf1007554406ec064ad2): complete subsection reference.

<a id="canonical-81e10bcc3bff0214f1f0234103f97aca7bbfb63d192acf0591e4e1ed6ad41e60"></a>

## Next pages — bot_action / ebb30573cf33 / 4

- [bot_action.bot_skip_processing](resources--service_policy_rule--reference--group-001.md#canonical-2fc47a010c757e9509febb91d8398e7bc58cb7917debeed50042cda045cc76a2)
- [bot_action.none](resources--service_policy_rule--reference--group-001.md#canonical-da135dd49e7627c3b5db7532af22094d27677c12713ecf1007554406ec064ad2)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-2fc47a010c757e9509febb91d8398e7bc58cb7917debeed50042cda045cc76a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bb3bc710d3d532867bd7d2c2eefed058529237d1512c806fa923754ac1ff584"></a>

## bot_action.bot_skip_processing — bot_action.bot_skip_processing / dac2c3eeda56 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-936c6a7953e5f16177b88a818686a49d30aa935b9f8caa61e07bfcbe6b11ae52)
- bot_action.bot_skip_processing

<a id="canonical-be5ea125580f10de5815da936976d8ce6a8c4a2841834a4fa9fc34e938674007"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
bot_skip_processing = {}
```

<a id="canonical-c3c12b92ea1dba79e8db1e7f2b4b514d65d3550c20bd8a87037b3564886c5b56"></a>

## Direct properties — bot_action.bot_skip_processing / dac2c3eeda56 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8af58bd1b6cef382765e0b5171e801ad049cca89dcc92eb3cb6636cbeab1f943"></a>

## Next pages — bot_action.bot_skip_processing / dac2c3eeda56 / 4

- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-936c6a7953e5f16177b88a818686a49d30aa935b9f8caa61e07bfcbe6b11ae52)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-da135dd49e7627c3b5db7532af22094d27677c12713ecf1007554406ec064ad2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-171269196035db36ec373fd29185c456c6c807b1cf9539285ff030718f22bc3b"></a>

## bot_action.none — bot_action.none / d7471ce3df37 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-936c6a7953e5f16177b88a818686a49d30aa935b9f8caa61e07bfcbe6b11ae52)
- bot_action.none

<a id="canonical-82c59bf4c3c738268e9116268ffc4554cdd37264e8008d9cff443d64d64f6452"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
none = {}
```

<a id="canonical-c43c78493d03506892287feeb657fe2896f03ab71b8e69a9acbf3826180ff8dc"></a>

## Direct properties — bot_action.none / d7471ce3df37 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4999914e54034b3571dff1da82a05e92c319dc727f0fe055745588d8a174628a"></a>

## Next pages — bot_action.none / d7471ce3df37 / 4

- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-936c6a7953e5f16177b88a818686a49d30aa935b9f8caa61e07bfcbe6b11ae52)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-3e6fcc4950ad994ad911391dcf12a4305b0b9c75569d289775436104117fd307"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-150af0f12c488b37dc78f2910dd93b4c6e4051628ed439b3879ed8907d729bd1"></a>

## client_name_matcher — client_name_matcher / ef48234266ad / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- client_name_matcher

<a id="canonical-56ffd1528ef78daa1c36b951f5930c56bb9d8518e6f213f9573bcef765ab2229"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
client_name_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-82483905de50f9f264dc240bda132d98d0f3582c87f850e3d211c454d3318315"></a>

## Direct properties — client_name_matcher / ef48234266ad / 3

<a id="canonical-540b25072558891fbc1f8c93bf8641283c30ed646ac647eea88f55be7cd2ab3c"></a>

<a id="canonical-3bed75d7583ac799d809d643d427954e07128e7dc0d786c8846f15d77a4c5236"></a>

## exact_values property — client_name_matcher / ef48234266ad / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-c46b8c917351da408e1a314664be2bfd2bfb3ae0609a9e7592e6e4ee7f030922"></a>

<a id="canonical-a4e3cd8e1b10c70faf92d4504e42229c115ea03cbcdea4641bdd24492d3951c2"></a>

## regex_values property — client_name_matcher / ef48234266ad / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-10dfedb9b91145fbab0eb31803157af0d315f74eb7c936e88b6c9dad1332a91d"></a>

## Next pages — client_name_matcher / ef48234266ad / 6

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-70d913f0f8a715b2ff59d453d1f84bd6685e80d1a6caa80a9fef7d9e8dc7195c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0eb263b852295435b2b3773af0de0788487c6ae659efd5bfb22759c3d34bf621"></a>

## client_selector — client_selector / 2d92bcef697d / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- client_selector

<a id="canonical-508074458c6693764711cb853df217dbee86ee28414773d8ceed90ae3d61e707"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
```

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

Terraform syntax:

```terraform
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-34c8e93140bdb97af698b093116b36a76392b2726e635cb36702c363c8e25494"></a>

## Direct properties — client_selector / 2d92bcef697d / 3

<a id="canonical-212b04199feb77daf60397a0f83757dd43b995658c6864df1dc152af7a018a38"></a>

<a id="canonical-62e4a0cd36aa47bcc65513195c13581bf5425d3b00e362c7a4f8a588c23bc5ae"></a>

## expressions property — client_selector / 2d92bcef697d / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-306a4809d32290ad03a96da10062a230f6f754e1f20b751128894d48cc67657d"></a>

## Next pages — client_selector / 2d92bcef697d / 5

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9a1df9fc98326d62ded3b6a9810311c7ceda91b07568ad780aa370a83ac6f44"></a>

## cookie_matchers — cookie_matchers / 087ac9039d0c / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- cookie_matchers

<a id="canonical-731f7a55fcfed6015f8d07d61d9b070b62205751b67158e1b8e9a5e604ca98d0"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-9032e8ea5d82918fdf835bd39695ef7d8806706431578594c0b899ad18fe9c7b"></a>

## Direct properties — cookie_matchers / 087ac9039d0c / 3

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-be4203db59653ef1725b635c4c35ffb47ccc01c9add118f1a4c807ceaeab31e3): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-3140c9d155ef1f9a54b6ab1531098cb49a19a31f6836dd3ac148a3beda5186b3): complete subsection reference.

<a id="canonical-06cd36205575267943a9d7aefc8cd157c8f834a00d10b1ed37b19b91f77f0cbb"></a>

<a id="canonical-8ef174f4aad6b75dcf576777c8df40179bd8c2fa1755a5ba36e8dc475870a1d4"></a>

## invert_matcher property — cookie_matchers / 087ac9039d0c / 4

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-0c37c073de6a88673a30ff7beb9812cf95931c3404d1d27fdd00eeaec584b574): complete subsection reference.

<a id="canonical-219fdd8ab7a887ab05cd5c97c3db5b8fe46be06458c9a45355f1966cb28d066b"></a>

<a id="canonical-2b985fff9b46c4338192fd1c1652badb24d913020abebf25ce687a856900824d"></a>

## name property — cookie_matchers / 087ac9039d0c / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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

<a id="canonical-9d61d97b6b78c3af2b8fe15dc9b0d4a5bfe7fdfb10dcdf23821c14e6b84fa51e"></a>

## Next pages — cookie_matchers / 087ac9039d0c / 6

- [cookie_matchers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-be4203db59653ef1725b635c4c35ffb47ccc01c9add118f1a4c807ceaeab31e3)
- [cookie_matchers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-3140c9d155ef1f9a54b6ab1531098cb49a19a31f6836dd3ac148a3beda5186b3)
- [cookie_matchers.item](resources--service_policy_rule--reference--group-001.md#canonical-0c37c073de6a88673a30ff7beb9812cf95931c3404d1d27fdd00eeaec584b574)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-be4203db59653ef1725b635c4c35ffb47ccc01c9add118f1a4c807ceaeab31e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18afd574e672ddc4c57b570d25caac18edcdb7eeeb3c2543d84008b182fcb5c4"></a>

## cookie_matchers.check_not_present — cookie_matchers.check_not_present / 1d155e084840 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f)
- cookie_matchers.check_not_present

<a id="canonical-88ef488d54ef170fd15ac26cabeb3157f0a1b21165c268fff49367c733a4a674"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-bfa42c0cf897568674ebafb36c71f4b2f27cf3b82ce3d6ab0291dd7ad68dcb6f"></a>

## Direct properties — cookie_matchers.check_not_present / 1d155e084840 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07cfffdbfebde6d6ec94c433bd1ee222f3a93ea5fc2d8f10ba634459533e1e69"></a>

## Next pages — cookie_matchers.check_not_present / 1d155e084840 / 4

- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-3140c9d155ef1f9a54b6ab1531098cb49a19a31f6836dd3ac148a3beda5186b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66dbeb42825725ac3b59e987392ef09e9697b12231f5c32d9e6a2562e21e5500"></a>

## cookie_matchers.check_present — cookie_matchers.check_present / e36a42df1709 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f)
- cookie_matchers.check_present

<a id="canonical-d0031653dd7c370c4c966f2ce1248e6895b86bc13f53dc899d13fafe3e96d28e"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-1f55a1a4b1717f4f3eccab01b4d9b48cbaa14dcea09b9e93dc1fb294f29cc1f3"></a>

## Direct properties — cookie_matchers.check_present / e36a42df1709 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cda07a73e03416708a925e07a8370117072584c981228a95e259155cdca9e5cf"></a>

## Next pages — cookie_matchers.check_present / e36a42df1709 / 4

- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-0c37c073de6a88673a30ff7beb9812cf95931c3404d1d27fdd00eeaec584b574"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15dfe3d3f2aa646d2491012453c10e487c3394c81a003ad8956712ef2d386d33"></a>

## cookie_matchers.item — cookie_matchers.item / 1d7ded81925e / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f)
- cookie_matchers.item

<a id="canonical-de9b6a8dbc50b0a6793eacc2652a0e6edb964d9f19e6e822deb999e3f8ebde73"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-8a3f5aa2be80975ad2ac0aa9afb23a2b4698ff86fb57ae4f9026e1862955e30d"></a>

## Direct properties — cookie_matchers.item / 1d7ded81925e / 3

<a id="canonical-14e3985a3301a43b9e8a111911565d2ef443dae2b6f98087556a4976d395f2a7"></a>

<a id="canonical-b62f688b348daed739d928f77914ad017375e03b09c263f86112b29cd6fdde00"></a>

## exact_values property — cookie_matchers.item / 1d7ded81925e / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3a0de9359efe760ee3dc5cd9fb01bb8b4e0bc4d89ae78a9caf9327f542abd605"></a>

<a id="canonical-289de6f97462d39db260431e8039cf5f0232d71018c4047e02dd115caeaf6288"></a>

## regex_values property — cookie_matchers.item / 1d7ded81925e / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-fb9cd88bf9970d417d1d0555989c116d114e28f478f9f9a83283b5c313667d2e"></a>

<a id="canonical-79df6717b616b6fc2d341b88aff04ec9acccfe51bd34c2d0c15709bb49209e3b"></a>

## transformers property — cookie_matchers.item / 1d7ded81925e / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-f2f2dda8df3fbb9607d11dd528f9355c488559f2a0de3d6e7d98e7c8a97543d5"></a>

## Next pages — cookie_matchers.item / 1d7ded81925e / 7

- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-3baf3cc67b667ac0f3da1982967a4fa830a28b38dac4bf0b5fac1044a3d1e13f)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-7ea9c15f3bbe81e53c56ac87cbe279b4bf1e79773e794059e812c84b8e4e2be4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f067abab9bd83d13ad472457f4bfa182c90550ded4b93fc235b11c35e7660ee8"></a>

## domain_matcher — domain_matcher / 7d09751a2cb1 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- domain_matcher

<a id="canonical-cdd4331f0a134618a66f111094d0ec997c72ddd6e8a96574127e1ab83555de79"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-115c19dacfd57dcd27a45be2c4632b2671af2742eaa1b4f727dc1c8f6f031981"></a>

## Direct properties — domain_matcher / 7d09751a2cb1 / 3

<a id="canonical-c55e84ad4f921e7dd31966e6bbbc4997658e444a2ecca6b2c06ba58f7c864a1b"></a>

<a id="canonical-754dd48ab203ed0eb445df4a0879390bfcf11950437ccbafa2121c9a1f9f4364"></a>

## exact_values property — domain_matcher / 7d09751a2cb1 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-ac9403e2b0064e7593de15d158f2020f6ed374f47f46a1886927aab8bafc62d5"></a>

<a id="canonical-a43654d6889b611cc8b03dece4efc3a8cd9d59180c20a18247dc8257d48865a4"></a>

## regex_values property — domain_matcher / 7d09751a2cb1 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-ec409320e527f5225b74c9f446ad4da8df59c407a97b4f06f2e29be69bb4bdeb"></a>

## Next pages — domain_matcher / 7d09751a2cb1 / 6

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-832057437a1638c2ca18c7175971a464bd04be46a0d8e20cac68852133b81570"></a>

## headers — headers / 1d82ac169249 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- headers

<a id="canonical-d0b74e79562f149eb94c59c2295718ac38b374e646b1015c449f821c7d63a275"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c2f5f93b8ae2051b2e106d992e27892cbd55af242101de23dbc17b039351051"></a>

## Direct properties — headers / 1d82ac169249 / 3

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-f1f256393ca6c428f2b3506cf1a45111646c64a023876cc3a75b9318ce62f684): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-18b85653cfaeb1cc78d34ba7cf0c1c4d622f781ae5d660583cf22ecd6faff81b): complete subsection reference.

<a id="canonical-21eaff1127d61f19a4c87c798cb2a711351c19928b4071fa09d4aacc3c93a8a3"></a>

<a id="canonical-dfa9c0c84aab2ce461f599ba34820fee23f9131dc3f42e921fc722aef9074ff4"></a>

## invert_matcher property — headers / 1d82ac169249 / 4

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-c390cb55a1783884615a551a13266833a4a641842760d10cb80f23844e9b33c8): complete subsection reference.

<a id="canonical-fcaa4c20b655063118afeacf17bd3efe1775f053c8089c7fe53f91705c11997a"></a>

<a id="canonical-dd387be0473a491bb1f479a7e4f4070bdb32e581dae11df9969d6bbc9ce43fa9"></a>

## name property — headers / 1d82ac169249 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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

<a id="canonical-2f353d2c055601b57df6178eff85390756e776c0e6803eafcf21c758dfe14e66"></a>

## Next pages — headers / 1d82ac169249 / 6

- [headers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-f1f256393ca6c428f2b3506cf1a45111646c64a023876cc3a75b9318ce62f684)
- [headers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-18b85653cfaeb1cc78d34ba7cf0c1c4d622f781ae5d660583cf22ecd6faff81b)
- [headers.item](resources--service_policy_rule--reference--group-001.md#canonical-c390cb55a1783884615a551a13266833a4a641842760d10cb80f23844e9b33c8)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-f1f256393ca6c428f2b3506cf1a45111646c64a023876cc3a75b9318ce62f684"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49a47a403b85ff746ade374621d00bad37f7c8bbc60e261532a748b95f13bb40"></a>

## headers.check_not_present — headers.check_not_present / 9352f6043a77 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [headers](resources--service_policy_rule--reference--group-001.md#canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480)
- headers.check_not_present

<a id="canonical-1bac64f166a5ad124a42a06a751ad96c2a23b2e4ecf54ea467a2726422226a69"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-390d746c1b97edb5573430ac7145cca0e4ec0a2d74efdd5dea9ca26fe0d0a37a"></a>

## Direct properties — headers.check_not_present / 9352f6043a77 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7868a409f05f0882bc3ffa1d0f1e6cb6b01227d875aba3d29da08e61183558cf"></a>

## Next pages — headers.check_not_present / 9352f6043a77 / 4

- [headers](resources--service_policy_rule--reference--group-001.md#canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-18b85653cfaeb1cc78d34ba7cf0c1c4d622f781ae5d660583cf22ecd6faff81b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d056264d250fe6e314694e598c24b4999eded5f46a763a3e5e1a84dc7df76606"></a>

## headers.check_present — headers.check_present / 571a6cc911a0 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [headers](resources--service_policy_rule--reference--group-001.md#canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480)
- headers.check_present

<a id="canonical-fd476633b87fac8dd6d1c9dc5833674b3814130094235adab0c94f34a630d179"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-dbe36b325e64f22c91449969f6e601740f84ecc8b58cc195e75888c232431fee"></a>

## Direct properties — headers.check_present / 571a6cc911a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cca768458e17c19aba6c8ab060e3644b34eb06475095ce766c0315e7186a2b66"></a>

## Next pages — headers.check_present / 571a6cc911a0 / 4

- [headers](resources--service_policy_rule--reference--group-001.md#canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-c390cb55a1783884615a551a13266833a4a641842760d10cb80f23844e9b33c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fec3e47cb8ea8851c994ece1b4a4429ad77399ef4a6b3260b0b77fe3f4b2b24"></a>

## headers.item — headers.item / e88a2b76c3b0 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [headers](resources--service_policy_rule--reference--group-001.md#canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480)
- headers.item

<a id="canonical-c2da0ba0fcf3e97f8c7eb74282aea3b8c2c66237bc8edef2fa30daff26154858"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-f9ef5ab116a93ad361a85832a53d7149e0ee74a242d1f784dc1bdc344d1e33ad"></a>

## Direct properties — headers.item / e88a2b76c3b0 / 3

<a id="canonical-0015cac2002affc20fc573583b8a9578ca7a772a06bc96cf368513e2e1e7de2b"></a>

<a id="canonical-6759f110ab13415242f9bd2211d1be3bd062b55a6266ec171b4cdaee1de6a7da"></a>

## exact_values property — headers.item / e88a2b76c3b0 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-5b282d7a1c059a941b63fa727987be988bda322338d482e14af8cabaf660f97c"></a>

<a id="canonical-12b7c649ae62ba7587091590ec7e1f97596aabe6d59088cb52994cff86cd0e04"></a>

## regex_values property — headers.item / e88a2b76c3b0 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-a4ea4cd98fe17e432dfa7f65131e767d34f746b23513b806c5773128063e2533"></a>

<a id="canonical-4a84de82a6af1d31cc6e3b78990ef0cc96ec631eb13afd28faa450c033646ce8"></a>

## transformers property — headers.item / e88a2b76c3b0 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-b5d42f18271d6d6c1576dcd1292a0849b5472cf5f66dad25287f5f617443ddbd"></a>

## Next pages — headers.item / e88a2b76c3b0 / 7

- [headers](resources--service_policy_rule--reference--group-001.md#canonical-88e619e0c344e09203974efd12b9cef3fa0a5535d393cce25c4bfe2c4c9a7480)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-f5affe06710c1b7d6014ccf515752fa30f078d58d20a6cf7e719436ab7db52ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b606d5f171bb0fa7d6cbbbef9b864669a6ddf9b3d4ddd5e022ac4a008eed831a"></a>

## http_method — http_method / 53bab50626ea / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- http_method

<a id="canonical-7679c52108773459414d3c53b5c742b4f40b95503476788672bbee67d7e72bbf"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-19df419496d68eaedb6efa4fabf1e2e533ba3d7bd3592d348d77c3b3e9e7b991"></a>

## Direct properties — http_method / 53bab50626ea / 3

<a id="canonical-0209d45cca158b1d9304b9a40ab747404f823d7e436e4fcc93a96fd82417dfb2"></a>

<a id="canonical-2e3511d84a3ffef6f8bf456073c77ee9a1725a4238579210e01aacf0c9ef86e0"></a>

## invert_matcher property — http_method / 53bab50626ea / 4

Type: `"bool"`. Optional.

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

<a id="canonical-53789f8e323ecdd856cba93945ce57eca94bd73d8699b37fd4c7335f180716fa"></a>

<a id="canonical-48c9778ad636b4b057087d9cc3ffee5db8326179b46af9199c3af7c926275fba"></a>

## methods property — http_method / 53bab50626ea / 5

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-2a74e3f15479f7aaf154519a29a072702ec5f4461a4148391ecc9ef4cd465c45"></a>

## Next pages — http_method / 53bab50626ea / 6

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-5131505e9c5d1806673eebf13dd26429a151fd79793834461f68747727fecbf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6551a2a08c0e9040c75501b9d6d5cc2d420fbf28571f912754c577c677844e36"></a>

## ip_matcher — ip_matcher / 52f6a129ffc8 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- ip_matcher

<a id="canonical-62e6ea35a259a4d1cc1564db083dea4875104358518ab57a3ad5372c1c58114a"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
```

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

Terraform syntax:

```terraform
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-b89b9270372c41e9b43292dd3ca73c6811f982f59acedf7b497bc74832c79e21"></a>

## Direct properties — ip_matcher / 52f6a129ffc8 / 3

<a id="canonical-f3c4d3f111de7367f60ec4661911b5e420e4c6355fbfe255f8afd080e1f518dc"></a>

<a id="canonical-2b9dd025cfd89d4fde23850c0fb4478e3d9071c2b937d02c759ec1660cc43c68"></a>

## invert_matcher property — ip_matcher / 52f6a129ffc8 / 4

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--service_policy_rule--reference--group-001.md#canonical-2b5f3f6d9ec245554d1a13f08a3aac05ec6dde98018e3aa9eec82bd65e06b036): complete subsection reference.

<a id="canonical-0e0f58339f68c1d50a58fff3a5760a5cd5ce48a338993d1618df5a1351e697f8"></a>

## Next pages — ip_matcher / 52f6a129ffc8 / 5

- [ip_matcher.prefix_sets](resources--service_policy_rule--reference--group-001.md#canonical-2b5f3f6d9ec245554d1a13f08a3aac05ec6dde98018e3aa9eec82bd65e06b036)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-2b5f3f6d9ec245554d1a13f08a3aac05ec6dde98018e3aa9eec82bd65e06b036"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38014fdc265190a029ecb81de280979dd7718831eab3ed48dbf4820d11342e33"></a>

## ip_matcher.prefix_sets — ip_matcher.prefix_sets / c5123e2e9b0b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-5131505e9c5d1806673eebf13dd26429a151fd79793834461f68747727fecbf6)
- ip_matcher.prefix_sets

<a id="canonical-ef1f1b1b5d5e4d9f174582af2187bf8356a4a442b265ccaa80b10b8ac1237561"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-483bb3bde8ffffe82c5526824ae749e59a5d73bc3fecff8020321a75f942d377"></a>

## Direct properties — ip_matcher.prefix_sets / c5123e2e9b0b / 3

<a id="canonical-c918bf1b669db3e0254d4692d38b21ad334793ce90eb12e18b8a0d78b317ff49"></a>

<a id="canonical-1c3bfb9941535d269fa28a659133469949b2710a02a21109489323f5afd137a8"></a>

## kind property — ip_matcher.prefix_sets / c5123e2e9b0b / 4

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

<a id="canonical-cf1c54b1e85802abe03acaa4c91ed23cbe4f3fe548f24837fcd14e105e6f5533"></a>

<a id="canonical-5a2c458695e05f880be398c9db54bda006fa945fdcd3376e20704c76d8fa276f"></a>

## name property — ip_matcher.prefix_sets / c5123e2e9b0b / 5

Type: `"string"`. Optional.

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

<a id="canonical-3402f71578ebb5b0addabb79a91689ae443689ff2473ecdbce57652921e39252"></a>

<a id="canonical-08396039257cb1becaaff7e03ff40035c21e304da6eb1f289ae9eaae9034779b"></a>

## namespace property — ip_matcher.prefix_sets / c5123e2e9b0b / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-666f02d1be6607d3898da1b78007a6f957b0d97e51b2e5adaaf8f53ceb3e764b"></a>

<a id="canonical-12b7716f8cedc4d95164b2e29caa7cbd5e114136baf22d6d78ae6474af221835"></a>

## tenant property — ip_matcher.prefix_sets / c5123e2e9b0b / 7

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

<a id="canonical-94997a446790b620904587c5ff16c5230d8d5f4be188dc9262ab996f2780dae2"></a>

<a id="canonical-964a9ab789aa6ce19a85998298cc41aa1efce119a4b448fd09a77b67ddc8bda7"></a>

## uid property — ip_matcher.prefix_sets / c5123e2e9b0b / 8

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

<a id="canonical-a9cc0226f940a1653c682efc46e48a6082e7eac256bec64c1c83859ca9c7044d"></a>

## Next pages — ip_matcher.prefix_sets / c5123e2e9b0b / 9

- [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-5131505e9c5d1806673eebf13dd26429a151fd79793834461f68747727fecbf6)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-413521cec6a72c537f18e294f798d20fdf3fef19b519f823e0f7d2e09417ab79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d9de887e9f34ae0134717d85c02e132584511a038990fdc029450067518319c"></a>

## ip_prefix_list — ip_prefix_list / 2ba2029b6de0 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- ip_prefix_list

<a id="canonical-5f120eb5a39bf5998ae353e86ae3d0b7b2fd5ac2d601d68512970f8838346d0c"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-cd15bbae9d336d7857e6cd5262508a625c282b63fd676b1f56f85b971bf846d2"></a>

## Direct properties — ip_prefix_list / 2ba2029b6de0 / 3

<a id="canonical-5cfdc0a79af2f690be32cef9a00e4815210349545fcf0bf240daa3292ee191c8"></a>

<a id="canonical-0563d22e9d80f46d7f0edaf09161a6d0f4313002dbb9e517ad6aacfb1c4c5283"></a>

## invert_match property — ip_prefix_list / 2ba2029b6de0 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-04e3dce9ad1940b64607a31d5cd5f01b4c65203801dde8399a0fcd506e20f3c5"></a>

<a id="canonical-e607065a7690ec1d930277ba146f08f13e949ff2f25ba2b955d296f93492a3cb"></a>

## ip_prefixes property — ip_prefix_list / 2ba2029b6de0 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-122d0b99f494cdaa56af9d51e90a8ba7b3d8484c64c3399aaa5de43e05343e6b"></a>

## Next pages — ip_prefix_list / 2ba2029b6de0 / 6

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-06cb5e2962be24ee85db336adacd12583343c5a13b249653673224a7757d992f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1af1e42a5e863dfa9be82956ef60212ac4136c973ab05ec16592e6f0f3d735a9"></a>

## ip_threat_category_list — ip_threat_category_list / c14ea4026504 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- ip_threat_category_list

<a id="canonical-32815767088964b8a209c509d3c0680e7e5365edf89016f9ff8d9355507a39bd"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
```

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

Terraform syntax:

```terraform
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3b365a2d82a0065fbcb889f35ea49f905ca39c69f3f740c6855494e5f5ceb620"></a>

## Direct properties — ip_threat_category_list / c14ea4026504 / 3

<a id="canonical-abc0b874c7d15d717bbca3cf406369555a60d471002e445eb87290bab5efe8e2"></a>

<a id="canonical-b490827e61005961d46161a8d6cc22232c302b90af96e77cb1925e6ac24b96d5"></a>

## ip_threat_categories property — ip_threat_category_list / c14ea4026504 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-7a321736e5ef95093be1d0619ff3934d25c2af4b284ced88728d76a29adde14e"></a>

## Next pages — ip_threat_category_list / c14ea4026504 / 5

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-aed76a8d199a5769ebe931308623edc55ce47695ebacff7f084193b3f4a9709f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23dafc5a9b5645cc06024fef768aaa36e4f09f8b86f70f4ec1053caed3ac0d9a"></a>

## ja4_tls_fingerprint — ja4_tls_fingerprint / 73356c6b72ff / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- ja4_tls_fingerprint

<a id="canonical-b858314b40a61f4215a2dc82656112027d75f3e3a83c31c58d8fc7f435447dc3"></a>

Type: `"object"`. single nested block, Optional.

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

- [ja4_tls_fingerprint](resources--service_policy_rule--reference--group-001.md#canonical-b858314b40a61f4215a2dc82656112027d75f3e3a83c31c58d8fc7f435447dc3)
- [tls_fingerprint_matcher](resources--service_policy_rule--reference--group-002.md#canonical-b02fed0ec6f83276f5054e82dee19e02e679bcd5a8c478e201220c87e66bfe6b)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ja4_tls_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc9779bd2a823a6fb34cdbcca43de57a8b5c24b66adf5001f9aef68e8f379f11"></a>

## Direct properties — ja4_tls_fingerprint / 73356c6b72ff / 3

<a id="canonical-b8e74593a4d7311274016be5e1975f90c5c8ab0c62776d2b27ee41dfe177d138"></a>

<a id="canonical-822f768c74e3c7ff9863543d82b2cbf53269e7d73ed388c410731bdd1bca68ac"></a>

## exact_values property — ja4_tls_fingerprint / 73356c6b72ff / 4

Type: `["list", "string"]`. Optional.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-e497fefe272616545b90468f83e1dba4715ecb28c512f55628d22f9d57686949"></a>

## Next pages — ja4_tls_fingerprint / 73356c6b72ff / 5

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faa995dd99f8ba8197d3204b3570e2cee9b27ea1aef4506f0ca701c62f38c969"></a>

## jwt_claims — jwt_claims / 9ed2c21dd68e / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- jwt_claims

<a id="canonical-2e55d13a4b1be9943bec8c8805bf9f0066edec72cfa3d71700e58a205ab4dcf1"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-a917c6a7fec8e32d71e2da7a7202aae4b69915fe4ecb1aee01ab5c15e5ca3ae5"></a>

## Direct properties — jwt_claims / 9ed2c21dd68e / 3

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-b3e1371e9148a368ef91a120deeedeaa62fb64b70d5f2b5369b252e879b6edc2): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-8d3266bd6d6bbdd172c224ce14aabfda0127714ba0d20010fc777883f862d3a4): complete subsection reference.

<a id="canonical-5132dc611148aa88ac951cafc5436065753b90b12cce50e20e44e7c7f6a463f0"></a>

<a id="canonical-ba14dd7f411e1c9b24c86ced773b0160743b0139e266e0d9a44cce9ede691156"></a>

## invert_matcher property — jwt_claims / 9ed2c21dd68e / 4

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-b5b84f21d2ef04609efdecfc3e892da9287664fc240bde304b604f24958f4ce1): complete subsection reference.

<a id="canonical-6d70639cb1204851a0aa7138bf713e136b3ec9ccb5be54d2f233ebe73d5e1911"></a>

<a id="canonical-122ad8477cd9325ba54843fbbda9e77c55237305f9bc689808bb35cfcce9180e"></a>

## name property — jwt_claims / 9ed2c21dd68e / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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

<a id="canonical-2499ebf9b7eca32291f117ba0e628b1a2638d91ae8691ef425f4c4b957642f42"></a>

## Next pages — jwt_claims / 9ed2c21dd68e / 6

- [jwt_claims.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-b3e1371e9148a368ef91a120deeedeaa62fb64b70d5f2b5369b252e879b6edc2)
- [jwt_claims.check_present](resources--service_policy_rule--reference--group-001.md#canonical-8d3266bd6d6bbdd172c224ce14aabfda0127714ba0d20010fc777883f862d3a4)
- [jwt_claims.item](resources--service_policy_rule--reference--group-001.md#canonical-b5b84f21d2ef04609efdecfc3e892da9287664fc240bde304b604f24958f4ce1)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-b3e1371e9148a368ef91a120deeedeaa62fb64b70d5f2b5369b252e879b6edc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b267bc2b79343340de9a72a0f54dbaa9ebd99518c8dcf6069f1327b84379537f"></a>

## jwt_claims.check_not_present — jwt_claims.check_not_present / 4cb260e4134a / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397)
- jwt_claims.check_not_present

<a id="canonical-e50e87a89580f07625bac57cc015e31094c0b4cd2370d938e0a8285fe973649e"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-cbb27f200c7a7f5c713e29b4a63e792709afd19e969a6c7cb531f57f5a044c93"></a>

## Direct properties — jwt_claims.check_not_present / 4cb260e4134a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bdd7c871a0177a2ee92bb9546c6253c47f705dc4b4e1bf85603511e47f9e5e3"></a>

## Next pages — jwt_claims.check_not_present / 4cb260e4134a / 4

- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-8d3266bd6d6bbdd172c224ce14aabfda0127714ba0d20010fc777883f862d3a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-092c6cd160015458b9dba9a9303afd203aac62d16f648f4b17f14983db213472"></a>

## jwt_claims.check_present — jwt_claims.check_present / 39ec63294df3 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397)
- jwt_claims.check_present

<a id="canonical-c4bf72353f8d455ed45e629f48294a6eee15fe380329aeaa1b9a08a74e1eafed"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-3dff40ab053c96680be9101501e1fb474825e42483529bbaeabd0ca736b83891"></a>

## Direct properties — jwt_claims.check_present / 39ec63294df3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bcc47a32c7cb8ace76ad5b22daa8856ed42175e9286a2f65a93dc1e25736f9f3"></a>

## Next pages — jwt_claims.check_present / 39ec63294df3 / 4

- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-b5b84f21d2ef04609efdecfc3e892da9287664fc240bde304b604f24958f4ce1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09abc417f10f8573b5cf08b9b338c80a6bb717c55919e410ba4cd20ad3a28cd4"></a>

## jwt_claims.item — jwt_claims.item / 1ea83c790ef2 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397)
- jwt_claims.item

<a id="canonical-da8dad4d38ac4cf5b32c958060861029b9ef178502b372547649eadae9c71145"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-a186985d69ab80ec16431f39985559f70b57051cedcbbde2aad36263dc6e077a"></a>

## Direct properties — jwt_claims.item / 1ea83c790ef2 / 3

<a id="canonical-5d8c89425e1dfa5cbcefc80ff8600c5e89c52dfdbfbae7f281e9f44543311f7d"></a>

<a id="canonical-3576739ce6be931b76964ff5ecb3270ab813f4f14066cacba07ad434e501430a"></a>

## exact_values property — jwt_claims.item / 1ea83c790ef2 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3f96d62da947fb07bea0331838177e1dfaa360c72148dedad510252f19458ae0"></a>
