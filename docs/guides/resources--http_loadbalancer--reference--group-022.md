---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-9abbc413775c12e4afa6381b0e510bb6a4d98c4f9013e222abf51022bd45db3c"></a>

## description_spec property — policy_based_challenge.rule_list.rules.metadata / 3ea242021dd9 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-fef9c12c5b0b8a40ee6261cecdc1dc188b9d01242256b146f82e9936a4135c2f"></a>

<a id="canonical-383adbf272a343860b2e4806f469d7647ecaeb857703866572a51049fb9d889d"></a>

## name property — policy_based_challenge.rule_list.rules.metadata / 3ea242021dd9 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-d1883f656cb0421c0046ba34c874ca523541baa113e725c7672145ab51907550"></a>

## Next pages — policy_based_challenge.rule_list.rules.metadata / 3ea242021dd9 / 6

- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-595e68beca5673b8c0f8902b0a04e67dbc80bd9c71a5ee352c7239245de55b24"></a>

## policy_based_challenge.rule_list.rules.spec — policy_based_challenge.rule_list.rules.spec / 74ebca8f0fde / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-8acd51a5689de6ed8ae66c085d78fb1579571001bd32b772c52f0bb9296121ae"></a>

Type: `"object"`. single nested block, Optional.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_captcha_challenge"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("enable_captcha_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-9431a21755ee898b3a380b7b806e6393c66a136d256cf4f6391700bd8eaba3ac"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec / 74ebca8f0fde / 3

- [any_asn](resources--http_loadbalancer--reference--group-022.md#canonical-24c256b49a73d756a2425760a35a63b9b8f155434abdea6f1df2d780a4afd01f): complete subsection reference.

- [any_client](resources--http_loadbalancer--reference--group-022.md#canonical-c11f54f81fa65b7a2d684f6d88b7f8641693229beeb0ba17491a971e0de1a27e): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-022.md#canonical-c2a962e94757c3e08f0034c477796fced39dd6c81cf82d84877ca5651f2158cf): complete subsection reference.

- [arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-022.md#canonical-03c67751763ddfe791cc15c94b9a960d7d9d5e17dd4b99263b005be0ed5bea4e): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-bcba8dc798025cabc69a2ab757bb11c9ed8b966f3e954d8b47065e7486e9e518): complete subsection reference.

- [body_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-7a342cf4551145c4c401896b16362ecfe70450180e3acb11e8ae356f0b308c4d): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-022.md#canonical-f104f8ad9f7fb9d4b9be24f17487940b72859842e0ba92cb1437306b26deaadf): complete subsection reference.

- [cookie_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc): complete subsection reference.

- [disable_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-22cbb30d3ef9e29c7140b0efb284417070e37d6efecefec789fa81670edd927a): complete subsection reference.

- [domain_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-adeb9c1ca50e09d42c628bb43c362f9a62f3fa638ab8c6f11e2e1c745af570cc): complete subsection reference.

- [enable_captcha_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-bd4e0621310490ba1ebc8a4ea3482e0ac15374be351970bc2208da254b51d619): complete subsection reference.

- [enable_javascript_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-a9597852ca4de95178a33faab3f16331488f823590f32aa44606eccbdcdba2cb): complete subsection reference.

<a id="canonical-cc6c3ffbf980facb392c3eb0b46a95ddeb8a6af71a6f853d7e4d107de6d741a0"></a>

<a id="canonical-6954c0834b0218e1e2d024edb2800fdfbe6644b8360b40f773b4d488996a1b70"></a>

## expiration_timestamp property — policy_based_challenge.rule_list.rules.spec / 74ebca8f0fde / 4

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

- [headers](resources--http_loadbalancer--reference--group-022.md#canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5): complete subsection reference.

- [http_method](resources--http_loadbalancer--reference--group-022.md#canonical-bc152563e091075896817bbbd82667aee3dbc790b9f2bcd14be601f1b313317b): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-ddbb6fb7bd6975151d125412349de656d5834509446663c7eed4067c68b5a594): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-022.md#canonical-52cc79e99166ebccce4e5a198deeb5e7a4045546419946c41882e724d826e7e6): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-022.md#canonical-947327aa7fee2ee2917a57bfd86fd202b9de6a5104002e194ff7d5e557d36fa1): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-022.md#canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-5d3c8888dcffc18cc205b64c7baf988e6975e341fb257285637733dcc4349d37): complete subsection reference.

<a id="canonical-c364a07ad0f26d7b1619a4f5ced441552c0570f4242c1e463d3a33a72ae977a6"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec / 74ebca8f0fde / 5

- [policy_based_challenge.rule_list.rules.spec.any_asn](resources--http_loadbalancer--reference--group-022.md#canonical-24c256b49a73d756a2425760a35a63b9b8f155434abdea6f1df2d780a4afd01f)
- [policy_based_challenge.rule_list.rules.spec.any_client](resources--http_loadbalancer--reference--group-022.md#canonical-c11f54f81fa65b7a2d684f6d88b7f8641693229beeb0ba17491a971e0de1a27e)
- [policy_based_challenge.rule_list.rules.spec.any_ip](resources--http_loadbalancer--reference--group-022.md#canonical-c2a962e94757c3e08f0034c477796fced39dd6c81cf82d84877ca5651f2158cf)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b)
- [policy_based_challenge.rule_list.rules.spec.asn_list](resources--http_loadbalancer--reference--group-022.md#canonical-03c67751763ddfe791cc15c94b9a960d7d9d5e17dd4b99263b005be0ed5bea4e)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-bcba8dc798025cabc69a2ab757bb11c9ed8b966f3e954d8b47065e7486e9e518)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-7a342cf4551145c4c401896b16362ecfe70450180e3acb11e8ae356f0b308c4d)
- [policy_based_challenge.rule_list.rules.spec.client_selector](resources--http_loadbalancer--reference--group-022.md#canonical-f104f8ad9f7fb9d4b9be24f17487940b72859842e0ba92cb1437306b26deaadf)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-22cbb30d3ef9e29c7140b0efb284417070e37d6efecefec789fa81670edd927a)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-adeb9c1ca50e09d42c628bb43c362f9a62f3fa638ab8c6f11e2e1c745af570cc)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-bd4e0621310490ba1ebc8a4ea3482e0ac15374be351970bc2208da254b51d619)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-a9597852ca4de95178a33faab3f16331488f823590f32aa44606eccbdcdba2cb)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-022.md#canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5)
- [policy_based_challenge.rule_list.rules.spec.http_method](resources--http_loadbalancer--reference--group-022.md#canonical-bc152563e091075896817bbbd82667aee3dbc790b9f2bcd14be601f1b313317b)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-ddbb6fb7bd6975151d125412349de656d5834509446663c7eed4067c68b5a594)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](resources--http_loadbalancer--reference--group-022.md#canonical-52cc79e99166ebccce4e5a198deeb5e7a4045546419946c41882e724d826e7e6)
- [policy_based_challenge.rule_list.rules.spec.path](resources--http_loadbalancer--reference--group-022.md#canonical-947327aa7fee2ee2917a57bfd86fd202b9de6a5104002e194ff7d5e557d36fa1)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-022.md#canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-5d3c8888dcffc18cc205b64c7baf988e6975e341fb257285637733dcc4349d37)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-24c256b49a73d756a2425760a35a63b9b8f155434abdea6f1df2d780a4afd01f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2eb342cd4858d36bc817937b02afd4916ad30c92441802794bf253bfadf00ab4"></a>

## policy_based_challenge.rule_list.rules.spec.any_asn — policy_based_challenge.rule_list.rules.spec.any_asn / be0c62ab2031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-87aab8c3440a40622945657ac5284867e3aab33fa6b9ddd2f194becca8f0c8a3"></a>

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
any_asn = {}
```

<a id="canonical-df0f2a41d162a4acbae868c61f3ec33226238dcc6486544ad44a0816bfd19cdc"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_asn / be0c62ab2031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f9a14258c9672540a077f0b39151e3196905b804e74dfe9a8cd2e67081de03ec"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_asn / be0c62ab2031 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c11f54f81fa65b7a2d684f6d88b7f8641693229beeb0ba17491a971e0de1a27e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b3d34cf0eb15aaea31598db892768a42dbd78fc91f8aa113dd1451c2e6227dd"></a>

## policy_based_challenge.rule_list.rules.spec.any_client — policy_based_challenge.rule_list.rules.spec.any_client / 3877b27117cb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-c0abf02f2b4841d8710fc1db443f3d21e64f17fd20faf39b20fc4e872162196e"></a>

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
any_client = {}
```

<a id="canonical-77e2579b2181adb385fc9fe5a5f51448b70604862609ef1f35406d86ff7e9e45"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_client / 3877b27117cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3562f9a4109f3a2a331d1092c463c47bb53b5f857c58e7775ea4cbf33ba01757"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_client / 3877b27117cb / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c2a962e94757c3e08f0034c477796fced39dd6c81cf82d84877ca5651f2158cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-336f13cef7d10a6f84d923bab23dd4ebc48d7fc5e00acc4b473b6e4db7bf073d"></a>

## policy_based_challenge.rule_list.rules.spec.any_ip — policy_based_challenge.rule_list.rules.spec.any_ip / 1037a9f0db84 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-1be11da4fb120f14d8f4d03951c72a7652774b33c2368c79395cd8def550588f"></a>

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
any_ip = {}
```

<a id="canonical-0702573fd4bde9bd79ddaad51c74280d0d9e4b4ef690e62a948184aacde8c50f"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_ip / 1037a9f0db84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-344385c2cc152def80dd9abcc2ff3a0a273714d130301f5b7080b764e1f6740a"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_ip / 1037a9f0db84 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c72e4b9c6e64bbd0d255d5083375a0fe6c314eec384a57b9fa55b2e114d6368c"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers — policy_based_challenge.rule_list.rules.spec.arg_matchers / 5aef1163bdc2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-41dc9233afb9efaea1ed3a46df564899c8a6b570cfeee462238f6f82990e7982"></a>

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

<a id="canonical-5e2f6adc66774373c0aa0d06fbe4b43a8bd8d7eeea6e4930c3f60a4353230323"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers / 5aef1163bdc2 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-b322de1efa6a161ae47559a0d8cc69aebabd61097570cf7306e27b5963a089e5): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-022.md#canonical-5da87362e5bab0760763c7cd54ba7fbe9280b6e3954f4c7fbb6b0edb08b4fe70): complete subsection reference.

<a id="canonical-5ca7e535bfbd365f6f1f5764666d58e9af29d4188e3492fdc77d9a1011650e0e"></a>

<a id="canonical-4fe7a0f26e17d3a569c3bfbd86b7df452fe068f472851af5d06f7c2ef271b33e"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.arg_matchers / 5aef1163bdc2 / 4

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

- [item](resources--http_loadbalancer--reference--group-022.md#canonical-5310600794882b088e72fa786f3a7fd7e22561250f5c340686a90935e5920a6a): complete subsection reference.

<a id="canonical-add22d492afaab858b0af00b659d4ca177f0181203b0e14e33346537a3088208"></a>

<a id="canonical-ebabaa0edfb771efefebf9ee8c83668069999ff9852ecc505c446aab00e79bbf"></a>

## name property — policy_based_challenge.rule_list.rules.spec.arg_matchers / 5aef1163bdc2 / 5

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

<a id="canonical-53527c0da9f0103a074c9e6d7faa2ee47f4cb59b66da107230ad25eaafb40627"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers / 5aef1163bdc2 / 6

- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-b322de1efa6a161ae47559a0d8cc69aebabd61097570cf7306e27b5963a089e5)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present](resources--http_loadbalancer--reference--group-022.md#canonical-5da87362e5bab0760763c7cd54ba7fbe9280b6e3954f4c7fbb6b0edb08b4fe70)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.item](resources--http_loadbalancer--reference--group-022.md#canonical-5310600794882b088e72fa786f3a7fd7e22561250f5c340686a90935e5920a6a)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b322de1efa6a161ae47559a0d8cc69aebabd61097570cf7306e27b5963a089e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f5c536cf3f94e19bee5fe3d4f20bdea68204062f6c4271e4c13f43a961cb890"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 80d66ed538ed / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-1067e6d7fd166a15c68774b9d700e39428cf9f4f57ada00e0885bc431f62f591"></a>

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

<a id="canonical-c1d575eb46651713d2db78f8364fb24ed757e353fbdf3046a624e38436de4b7f"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 80d66ed538ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e89e42ab2ea5f98f4fa63f28eea674cf2450d93a2abb6642ffdb23dedf08d1de"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 80d66ed538ed / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5da87362e5bab0760763c7cd54ba7fbe9280b6e3954f4c7fbb6b0edb08b4fe70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cba5097f709f9b6afc7fd47fa3989a3e9da7ee854c957d8610f8ef227319aedd"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / ed506f4b9faa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-0d598fb1dbf99b7b61d8cea215b8f123e2fc5ecf7c17a67dcf051359ff0b031e"></a>

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

<a id="canonical-a613fa5012ed24c8603e8637d52a1c3cc3d7f5f4fbd7735ecb579eaedf91e8c5"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / ed506f4b9faa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-367adaab94218fa0c1578e372372ea9beece6cde04b9b7d461c2d3aef615f6ce"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / ed506f4b9faa / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5310600794882b088e72fa786f3a7fd7e22561250f5c340686a90935e5920a6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b9d585bf56d9f21c47c0ac3c67b4e372fbfb65a193477ea770543cf9fbe571f"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.item — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / a8dc75d55ac0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-bcaa516e6186e1699d00c2709d4ec5951b9dfcf5d985d7986af2a5d64b7fcc87"></a>

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

<a id="canonical-ad639e9dafc76d0355546c7688051c8b08084676fd32cd465187b8c9731e3979"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / a8dc75d55ac0 / 3

<a id="canonical-cadad48dc305bfb399b1a6aae0f4c63b15ba0ecf56cf2c67110e5fe7171c3235"></a>

<a id="canonical-dcc671525abd1d521ce6158dc4c2207c6aefc8e7fa0bfcb1e0f483d3241f4cc5"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / a8dc75d55ac0 / 4

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

<a id="canonical-ba9d48f137ac4914ce9720981fe90cee2b7d2b87f8a7f28196c1d30002c575e0"></a>

<a id="canonical-6518f5e1627f3e6d1c2a6ce793dfbee54b0ded86a09de9fa2d393deb418be33f"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / a8dc75d55ac0 / 5

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

<a id="canonical-9e7f344c511c3197a25ea18fe131c36c3de38f3f31e7bf799f8bac2f9cca8af7"></a>

<a id="canonical-004a3fc306fb5c6f495c384d0c97a91f9f9dde402d48319fe2680b68eb839fd8"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / a8dc75d55ac0 / 6

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

<a id="canonical-a9fb2e3a79b2f9cfc6fe8aa69e173dd1837fb338c6e5d3a7817aeedd9c868973"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / a8dc75d55ac0 / 7

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-51dc50704d876ffeb1288a19618bdd7bca1ce8fc40057595b01bc133c116f20b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-03c67751763ddfe791cc15c94b9a960d7d9d5e17dd4b99263b005be0ed5bea4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9aba8295b7f202eb829b1d862f5306947148c1b6fee7ee02dcad94b5901e0a0c"></a>

## policy_based_challenge.rule_list.rules.spec.asn_list — policy_based_challenge.rule_list.rules.spec.asn_list / 7aec75674d73 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-f761a1b476cf159d3cb739f49b8c012570a0418262d5ade71a2e2d3b7b80dffd"></a>

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

<a id="canonical-a70d42aade4e59757fe8ae7acaaa81826a43ed4a6b4e405d48259a3a58cf2d8e"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_list / 7aec75674d73 / 3

<a id="canonical-85241b7b80768fe86732d049aa96ff85e08c5b2d4ece2844cfd791f0cf9a5a4d"></a>

<a id="canonical-41808fc4d7be3b818b4783c0a1fc5b9748fc43821af607a1012520c1731ca7b1"></a>

## as_numbers property — policy_based_challenge.rule_list.rules.spec.asn_list / 7aec75674d73 / 4

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

<a id="canonical-d3d60e2de1c0f7023553750764efde1c10a3e17933cdad480e93e14c3d43c63e"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_list / 7aec75674d73 / 5

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bcba8dc798025cabc69a2ab757bb11c9ed8b966f3e954d8b47065e7486e9e518"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd6c345341737b6332ce1084ca41cf559fd67b6f4fd09da2c7199a87fea8a735"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher — policy_based_challenge.rule_list.rules.spec.asn_matcher / 78a050901adb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-fe21d55edbd9576f2963b2d60e23206c5bb9d13e86e35b51725c1536eb39087f"></a>

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

<a id="canonical-2daff3cff34bf515cf624418a56884533f6a88578cabde8c310618305b09234a"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_matcher / 78a050901adb / 3

- [asn_sets](resources--http_loadbalancer--reference--group-022.md#canonical-3d13b0dba0d9310f64744700932805a1f9658f1c4a8100b7176a1ce7237ba149): complete subsection reference.

<a id="canonical-7f6f4cd525b4f032cc9cbbf80161befa2537e1a80787c4be58223b80f84e49bc"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_matcher / 78a050901adb / 4

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-022.md#canonical-3d13b0dba0d9310f64744700932805a1f9658f1c4a8100b7176a1ce7237ba149)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3d13b0dba0d9310f64744700932805a1f9658f1c4a8100b7176a1ce7237ba149"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e49bec8f460058e976f6dbb917977650f994fee2df1a351552c8162a3fc4bbd"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / ab69a3c5d128 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-bcba8dc798025cabc69a2ab757bb11c9ed8b966f3e954d8b47065e7486e9e518)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-072733e4743e143adabe4c5b7cf2a029c3aa6ecb404a9ea219a8a7d43c5b50e3"></a>

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

<a id="canonical-f25d6d28d965bf85f590a817814817f2af6f7ce966c11c406081b17a6acb4a27"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / ab69a3c5d128 / 3

<a id="canonical-38a7335761bb2a87a81c0b7f2d44769d6e9e1332c7abbcc109cac1b43ee7a662"></a>

<a id="canonical-f8686e1e22e0669d2ad40019c6fb0c3ac4a48091d9b8139e4b5e1cb779c489b7"></a>

## kind property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / ab69a3c5d128 / 4

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

<a id="canonical-2cd568b2b11b147b75762c65a09be5e0770c4e1707ae2e6756f91f0f0eab047d"></a>

<a id="canonical-309e06261bcad863efd8d0b525483558753073b4a202574b335639893077c771"></a>

## name property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / ab69a3c5d128 / 5

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

<a id="canonical-e64046d3a178b9c8659b3ce19b3729488731a26a612d8339c02b3bc0fbf61101"></a>

<a id="canonical-b61b3b06759f3f0c25553cef68e6779e1f2d969565144d556b28c60eb23a3dcb"></a>

## namespace property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / ab69a3c5d128 / 6

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

<a id="canonical-68fc9fe4de44f8b3d3e5c9fcef02088b1b5d34ff608b89bb9cbd5911bcf7f245"></a>

<a id="canonical-9d03b32905e30b7067f5730a84c12e160b816275875f2192062c04c22d99d2fc"></a>

## tenant property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / ab69a3c5d128 / 7

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

<a id="canonical-2632633778fa2b8a782f5c235d01e6e96aee65a201825cbb2fb787f030bb974a"></a>

<a id="canonical-105f7f01b3170e1b0c8ae6e97fc631615dc39ddddbcee5aad5d252aa6260b2af"></a>

## uid property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / ab69a3c5d128 / 8

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

<a id="canonical-edcf0cc8150fd8b5d7cc56098d5fb655566b64708d64428a9bb826a55b07749d"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / ab69a3c5d128 / 9

- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-bcba8dc798025cabc69a2ab757bb11c9ed8b966f3e954d8b47065e7486e9e518)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7a342cf4551145c4c401896b16362ecfe70450180e3acb11e8ae356f0b308c4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0378756c259396776c3dcda3fb271dffb87e0a2cb4f777fbc6951999d5f8a52"></a>

## policy_based_challenge.rule_list.rules.spec.body_matcher — policy_based_challenge.rule_list.rules.spec.body_matcher / ca2c58c93ba4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-588880980f0a935551b55dfd09c7337488148e82c3b370d5a3f0731a0d63602e"></a>

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

<a id="canonical-6bb7a3f54c045ba19979101dbfdd0ec8c80bfbd2726ea1bb0aecb47f82c3331d"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.body_matcher / ca2c58c93ba4 / 3

<a id="canonical-6a53d8601019c5183510e4e68d8ab7c1ed62bf212ef9e77d135d0ce4e6c6f8e4"></a>

<a id="canonical-365497e3a4c55d9b11d7f6a12311f95038c6af8db2c76ad4eac6ef113c9844f2"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.body_matcher / ca2c58c93ba4 / 4

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

<a id="canonical-9141be77ddf1f51d45f758f1a5578ab8dceacd5164b1e8f7def5619aa87b3b2e"></a>

<a id="canonical-68f60588265f8d360f6e055682c0560b26fc8b4fa8c61a9591d34fe0ee767d16"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.body_matcher / ca2c58c93ba4 / 5

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

<a id="canonical-6643e738464755fb703025839fa87b40ff1e7f83cfd4574c25121262ffee4b90"></a>

<a id="canonical-ddb856b9cbadba24dedc7517fd3cd9a9a484ca598667fcd4e93b6710028c19ca"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.body_matcher / ca2c58c93ba4 / 6

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

<a id="canonical-6bc5f601743eec804e1be3d0e8917da9195f81f7b69a3c405e7876937044fbe4"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.body_matcher / ca2c58c93ba4 / 7

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f104f8ad9f7fb9d4b9be24f17487940b72859842e0ba92cb1437306b26deaadf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-035634179034a39db3b1d1e8ab5eea02b90c8b24c3462d969a8fd88dbd46305e"></a>

## policy_based_challenge.rule_list.rules.spec.client_selector — policy_based_challenge.rule_list.rules.spec.client_selector / 9b587d326aab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-19364626a02f013873b3e9ce89ae0efcc9eb3e40ca286f43228b523b85cef007"></a>

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

<a id="canonical-99fcbe4460971f73651482a4778dc5e75c6c51d6b32d78c589ec4c9dbf62d445"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.client_selector / 9b587d326aab / 3

<a id="canonical-32d9f48404d21dbaa2b7d6392bfb168a6c9a37940a2d2dfbe17f406d6b089568"></a>

<a id="canonical-e3492c4524a1660a0413b5ede3ecf9ac2407ca7c66b45e7e71e5f187e81608f3"></a>

## expressions property — policy_based_challenge.rule_list.rules.spec.client_selector / 9b587d326aab / 4

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

<a id="canonical-8e67917a25edc4a013bc54f40601530601b36be1a1874319c62966a53bc77eb9"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.client_selector / 9b587d326aab / 5

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6b2cee99c5162e918f70f6e414e9c5fc44aaad8ee9753102ad8231115a3e827"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers — policy_based_challenge.rule_list.rules.spec.cookie_matchers / 5234bb49b448 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-b5cb1320d7f3a5b6f8290adb15e931aed3db6f83f4962b54a30203de83ddebf4"></a>

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

<a id="canonical-c290a83eaad881493e2255eae63e155fe746db14ea37faa2f6fee3af0a1efb1f"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers / 5234bb49b448 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-13dfc88e3feeac25ed1194ad4ac3dbbd750ba91ee484278da0077d2d5e602327): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-022.md#canonical-451594cc4be75d3fa4d200488dacae51255930bd17c06168274b6906fdea5f7a): complete subsection reference.

<a id="canonical-b835e1eba09831655a5a7c05ea8511127b203928d07a375280c3eb79cea35d69"></a>

<a id="canonical-ec8c8f55d7eef4ef3565c07ef81017eafd57d7fb3818873f50925a74b745acd8"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.cookie_matchers / 5234bb49b448 / 4

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

- [item](resources--http_loadbalancer--reference--group-022.md#canonical-b163300f5852e145353f30cdc747b8d4d3ac0048866a3aaec837e988de369e6d): complete subsection reference.

<a id="canonical-b8a3ba2e2ae680d54e5c01c872df6503c6f7c871ffa1347aee82e7d659f8303a"></a>

<a id="canonical-320e1afcfe1e6522130b8a325fc78511d1bb2da671c433a327554a0a7b56c2f7"></a>

## name property — policy_based_challenge.rule_list.rules.spec.cookie_matchers / 5234bb49b448 / 5

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

<a id="canonical-d1ec9058317d1ff6f9a92d88b72fd4d83ba2f664edf6e74706a9a4986c7d22a0"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers / 5234bb49b448 / 6

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-13dfc88e3feeac25ed1194ad4ac3dbbd750ba91ee484278da0077d2d5e602327)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-022.md#canonical-451594cc4be75d3fa4d200488dacae51255930bd17c06168274b6906fdea5f7a)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.item](resources--http_loadbalancer--reference--group-022.md#canonical-b163300f5852e145353f30cdc747b8d4d3ac0048866a3aaec837e988de369e6d)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-13dfc88e3feeac25ed1194ad4ac3dbbd750ba91ee484278da0077d2d5e602327"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-889bafed69772d3204dcf51cce76aed59c241f18121660522995ec407c9b88be"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / d8ce982a608e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-71d2935fee16954ea192464d6422223a637bf2590e543ac263a9ba16d52a3b90"></a>

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

<a id="canonical-670eee007b7fae8b673125694f6e2526d3ff41e3c0c15197af5701f23d0e8681"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / d8ce982a608e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6f36c03d1594ed22c521e5a47526f9b4da07bb77395b3fc0b9121787db70bf6"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / d8ce982a608e / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-451594cc4be75d3fa4d200488dacae51255930bd17c06168274b6906fdea5f7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b163d91a9407c019468c5239b65005e83db5ecfc9de768902d08f645d69fa82"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / b2be6eca44cc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-9a0969c168a53ee0f59ac74ac4f14a06b06aff80701f46e540e392ce6b51e574"></a>

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

<a id="canonical-c73907876840844b10b164e5462b482cd8e80b73128b2c645ec9c6e441fec903"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / b2be6eca44cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-940806092dc2462b8ac34ecd4de6b6ac3e305e40c6831398be2017ea320dd6f7"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / b2be6eca44cc / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b163300f5852e145353f30cdc747b8d4d3ac0048866a3aaec837e988de369e6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34b2bceddbba44e1bc6773c01f9ca8824b2b543e770cddbf80b558e623b6b755"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.item — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / c36db4e207c6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-2fd206414437490842f22d8c16460beb90719b400f8e09ce02b0e6ef02c77d84"></a>

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

<a id="canonical-4fc6f454eb9d41c2935f2c48269abb337f8b0adb9c95fc8cabdc849d64148bef"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / c36db4e207c6 / 3

<a id="canonical-25e5e3dc7c35be52525785a7ea2c3247bfe5ebd0dbe11a0d50e506e4e48f4109"></a>

<a id="canonical-4a61ea0f8bc8b2fdc5f6f3d6a74ae261f6ae5a088df297515c56184807260d0f"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / c36db4e207c6 / 4

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

<a id="canonical-de73ac3aa5ab8b3b2ea200bc9e3dc1e2f80f1ea49a3308ce5cbd64d91f19a548"></a>

<a id="canonical-38b1b8e4b3c5448be0e6bc40603ff3cd3a3d55c29531995cee4344829c62a099"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / c36db4e207c6 / 5

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

<a id="canonical-4f28515e0c0fe73dadde848b72844d6ed9a1378ea591ed5ebebffe8f390ffbb6"></a>

<a id="canonical-73e9a8d8e24700d9f893697fbccb3c295848d9ac866ba0b4e4e2f08d1b54ef66"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / c36db4e207c6 / 6

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

<a id="canonical-fb476d98126e5c34457a637540b25c53f8482f65da9742ba21e27e719690e368"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / c36db4e207c6 / 7

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-56acce62b474ef10b24feb1d0183c8e9a6cf251acf654b905cc1cc97eef353cc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-22cbb30d3ef9e29c7140b0efb284417070e37d6efecefec789fa81670edd927a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9de298f0d069c7d1bf773677f80f4235169dba405befbb37b5ae0627171f2e08"></a>

## policy_based_challenge.rule_list.rules.spec.disable_challenge — policy_based_challenge.rule_list.rules.spec.disable_challenge / 1e09c4b19bb0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-b3ca0c4b2c6e4bb6d45efd56a5b3275eea867bf4e51413a6905d40004969f3ad"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable challenge.

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
disable_challenge = {}
```

<a id="canonical-624596e7652237ca5d7b1e42669f7bd80b16131c9df4d41f65edab7deead0e3b"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.disable_challenge / 1e09c4b19bb0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0d785516eaff039c809ce01ec5b682735e8bdfe6e034f737392d35e25db458d7"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.disable_challenge / 1e09c4b19bb0 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-adeb9c1ca50e09d42c628bb43c362f9a62f3fa638ab8c6f11e2e1c745af570cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31825e2e585db60fc50837eb753d626f3613321742bcd78bfb0bfdf181d0e549"></a>

## policy_based_challenge.rule_list.rules.spec.domain_matcher — policy_based_challenge.rule_list.rules.spec.domain_matcher / cb229a1938ec / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-6eec2fcf163b0a475fe1df7413ca1a80c1e0b1a0961e62a90e4609b5a6e46a23"></a>

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

<a id="canonical-9c012779a3747cd2bc9c00e39914c4a2b6d25d59072736ed182d7da3858ef956"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.domain_matcher / cb229a1938ec / 3

<a id="canonical-0795f393ee1d70b2ec5caa4e5e97d8c1d72ef2f11362bfe74899fba8c4712138"></a>

<a id="canonical-3a203a3a4216c949500214c8953e4765d9ab2a7afbf1567086ad4365eda94d00"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.domain_matcher / cb229a1938ec / 4

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

<a id="canonical-a2530eb787e4282d75b778901e0c689d5a5c8f453002b1e2e7128933e799c203"></a>

<a id="canonical-77995a744eaeb7a83856fcf1db9edc36a53c5fbe708c2975ce948d470b0444f4"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.domain_matcher / cb229a1938ec / 5

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

<a id="canonical-183c0b1afed876b294e6fc45daa8e57bb21140fd47be221984504b1a7c7be622"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.domain_matcher / cb229a1938ec / 6

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bd4e0621310490ba1ebc8a4ea3482e0ac15374be351970bc2208da254b51d619"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb410e23ab23b723bb5c60f8275d2900007fedc25224846496189ed6d4401fd9"></a>

## policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / de4d7b7618c9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-9ef63d386a82088f05f1454d367a14a207e171dad2d2405f589b631310e871ca"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable captcha challenge.

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
enable_captcha_challenge = {}
```

<a id="canonical-685ef9147cc0e4f5f2f79b46cb52ba9deb1d93ec046c4d814dc8d43c1163b21a"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / de4d7b7618c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8fcea65ecea063fc7c6abd4f0e4c0576a3dbfea5b7a7fa51deb3ba8a76fe4e37"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / de4d7b7618c9 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a9597852ca4de95178a33faab3f16331488f823590f32aa44606eccbdcdba2cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-467c177d4cf484cce6dedd62c3951f5a1fd0bae975320f3d04051c1e127d924b"></a>

## policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 456de449aa79 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-5c894bd39d483de8cc99cfcc816e5f0046b04c0bad6c54a2ea63fcfb46961888"></a>

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
enable_javascript_challenge = {}
```

<a id="canonical-f52b8cd292381165d1d0143a504a1d239f9dbffa38fafb7dfde8b0d07dadfdb3"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 456de449aa79 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c33f060144e5f59b2ac2a07adf55fe9e23bbc3267b804dfd92bf40fd3cbdf8b"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 456de449aa79 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e7f76520ca076c1f1967d940674f5fa11ec064c77f8e42d319ea8022b71859e"></a>

## policy_based_challenge.rule_list.rules.spec.headers — policy_based_challenge.rule_list.rules.spec.headers / 5ff1f2bc366d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-e766229902e7c8db2f6cedd0bcef6346fcdc23ec3cc82202a1d559033252d991"></a>

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

<a id="canonical-402649baa98b1166fbc680d4cbd52274a43f3590248aa7afe0661b3c1130d1a4"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers / 5ff1f2bc366d / 3

- [check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-bb22824c69695604400b4571d37a6d2dea02297bab2bb3fae018657a84640ad7): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-022.md#canonical-65ce57746b671a6b3d0cbf70c786a3c9e7b9a85d61e80477ed8c9118b19f6787): complete subsection reference.

<a id="canonical-4a04ac643185939c6436843e2f42f8082620adad95c4dc60d403f2df1c690226"></a>

<a id="canonical-95969ea40b56d71ace8ce30046b6432801143fb1568afdba66c893d00cd6a73b"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.headers / 5ff1f2bc366d / 4

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

- [item](resources--http_loadbalancer--reference--group-022.md#canonical-c09b332cbb440d339fcb2ed8721846b5ecb87caba9d79d1d37bd1c5b1b219008): complete subsection reference.

<a id="canonical-2b56b763c67c1e4423ce4f603e2b2ee884bdc316e50361d27d712794a2d2a43b"></a>

<a id="canonical-901f05f2ae3ad23fc2a10c8d589ed9f934a9dcdbd010649a7104b1c074089d8c"></a>

## name property — policy_based_challenge.rule_list.rules.spec.headers / 5ff1f2bc366d / 5

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

<a id="canonical-312e63f3394d669c627a61ceb43641c1d8acb304002b7c2bdb7c055d266700b9"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers / 5ff1f2bc366d / 6

- [policy_based_challenge.rule_list.rules.spec.headers.check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-bb22824c69695604400b4571d37a6d2dea02297bab2bb3fae018657a84640ad7)
- [policy_based_challenge.rule_list.rules.spec.headers.check_present](resources--http_loadbalancer--reference--group-022.md#canonical-65ce57746b671a6b3d0cbf70c786a3c9e7b9a85d61e80477ed8c9118b19f6787)
- [policy_based_challenge.rule_list.rules.spec.headers.item](resources--http_loadbalancer--reference--group-022.md#canonical-c09b332cbb440d339fcb2ed8721846b5ecb87caba9d79d1d37bd1c5b1b219008)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bb22824c69695604400b4571d37a6d2dea02297bab2bb3fae018657a84640ad7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28a8f2fa2516d7bfe3f9df80442479c3d688d247e6041ab3047c8c3b5f9c1948"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_not_present — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 2ee881767344 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-022.md#canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-02c374d6f14946b0f189a1964fbe2518009dfa03e230d6f19b36c90de5e99c30"></a>

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

<a id="canonical-c73eb6fb1f4e15f9886990489df557b5803142e02a450cdb5e486ded99ea95b0"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 2ee881767344 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-725a91ebeaef55aae7ab930cc87154e22110ed05310d5f0d695b6c1665baf40d"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 2ee881767344 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-022.md#canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-65ce57746b671a6b3d0cbf70c786a3c9e7b9a85d61e80477ed8c9118b19f6787"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83ebba4555792cc06a4b4287535c367e1798e5517c2adad5fce57c42adeb4a4e"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_present — policy_based_challenge.rule_list.rules.spec.headers.check_present / 3df947c03dcb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-022.md#canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-79f3e5c59159440fd55dbf145993aa3ed833b51ac229a83efb5ea5aa3dfaa95a"></a>

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

<a id="canonical-ff0c876e5efc2b106e70688af1190f875f87d747e2f8067536ebfb56a24e76fa"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.check_present / 3df947c03dcb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f29841c008bfd9cb5c5c91f80e77f4aa1db153fc6981eacc30e37cebb5806a8c"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.check_present / 3df947c03dcb / 4

- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-022.md#canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c09b332cbb440d339fcb2ed8721846b5ecb87caba9d79d1d37bd1c5b1b219008"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b447b6412553d43a7619072d84439e3314686621acc55bfc3c643658ab985a8e"></a>

## policy_based_challenge.rule_list.rules.spec.headers.item — policy_based_challenge.rule_list.rules.spec.headers.item / e2a32c84589e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-022.md#canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-156fc241deb11ee5073a4cc179cc8713f6c38841271c9a1643c28e9a2d0f4087"></a>

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

<a id="canonical-3357959f257fab3c13fd9fe61c16d5e9063af9a1f8c28a0e3e0925ec4f27e8af"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.item / e2a32c84589e / 3

<a id="canonical-3fac2aff4c3e89b82c9a678417f3ae07e5598d54b067c2faa7fb5849e9514401"></a>

<a id="canonical-3041017afbb17877987f725654873f76c79d5fbbfcd8464db95c307f8ae7c67d"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.headers.item / e2a32c84589e / 4

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

<a id="canonical-595d64046c60cf775a19f5124f33ae56bed0fc3a6cecd5619bc577c2bfe68135"></a>

<a id="canonical-4cf18b6594d5c00342247e0e4cc72418e516317fbf275d16f12a0b6ddfec2112"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.headers.item / e2a32c84589e / 5

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

<a id="canonical-65e405cc6cfc2d3265bcb16dfa9cf5c82d24b8a404cde15a89ef57da9c6d61ad"></a>

<a id="canonical-c97c054acae2bacf0725969ebce3238115da6a29dcb01c07dc4ab4e1896a2d21"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.headers.item / e2a32c84589e / 6

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

<a id="canonical-b4b9a167176e80af56396de769287b917ab295888a3176dba3be974856f6fc72"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.item / e2a32c84589e / 7

- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-022.md#canonical-7784d14f2325195caf490922ad83e29c26d8c5430cf0967c0519b645857754d5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bc152563e091075896817bbbd82667aee3dbc790b9f2bcd14be601f1b313317b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2d1636bc4648126507d4fd24f692ba33d8f7e8cd2f539b33be370b0969e8792"></a>

## policy_based_challenge.rule_list.rules.spec.http_method — policy_based_challenge.rule_list.rules.spec.http_method / 74521f56baf4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-3883538a37ce909148415dc14b6e68e386561519b831cbc9a253a8d7f957bb8c"></a>

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

<a id="canonical-43b41d4f720dfa09721060d84f3c493a0676590156550c9735c77d40df2c1ec4"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.http_method / 74521f56baf4 / 3

<a id="canonical-bab4053e47383d70afb6e178167d0d54d95199f9d88c984ff100b06edbd45cd7"></a>

<a id="canonical-5c2b66acfc77ae34b6aaab074b34e8c3f2b4ecf3a2500bfd2142e793b1de9483"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.http_method / 74521f56baf4 / 4

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

<a id="canonical-2ff9d76e6038c1f5dc3aa4ba704e47c7a0440665083f3caee36d5efac1313229"></a>

<a id="canonical-22fb729f5664a75171cf3083262e77a0083df119098488c7d9d3522d873240b7"></a>

## methods property — policy_based_challenge.rule_list.rules.spec.http_method / 74521f56baf4 / 5

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

<a id="canonical-8553e94a12e26446c1ee238f0d15baf69a2ae25e4f0e33e1bc1f59de462a7512"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.http_method / 74521f56baf4 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ddbb6fb7bd6975151d125412349de656d5834509446663c7eed4067c68b5a594"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5b81e671288f72cf4ed234ec3be5a6b20ef2328f78b2b72c6b94b063dd5cc48"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher — policy_based_challenge.rule_list.rules.spec.ip_matcher / d01ab25b887e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-0b90da7138c44338837946abe04e483579e154d58824e7a98c0e571c7d280678"></a>

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

<a id="canonical-3edecebb20136fc8112f72b256d8e44de331b2eba09c00fb0c9ee17649387e95"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_matcher / d01ab25b887e / 3

<a id="canonical-de624675ce7a690a31ab1114faf5d03a32f0d2bf6e1fd7e2c396ff0eaa6ba53f"></a>

<a id="canonical-f44bc1c6a4b62abcdbdcc4f7863c8c506d5d8870b477ee5b04eb49ed69762e8a"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.ip_matcher / d01ab25b887e / 4

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

- [prefix_sets](resources--http_loadbalancer--reference--group-022.md#canonical-26f86ad00e83285b9b703b524b2032995ffa68091e2445e122b5d2e39055b468): complete subsection reference.

<a id="canonical-9edd3c2e574e209576044531d2e391e791d69243986c298e2e0f0052074033f3"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_matcher / d01ab25b887e / 5

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-022.md#canonical-26f86ad00e83285b9b703b524b2032995ffa68091e2445e122b5d2e39055b468)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-26f86ad00e83285b9b703b524b2032995ffa68091e2445e122b5d2e39055b468"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f5a78078cdc1f269d84147d03724b2e64f44cf0fd4d7aa25cc927dfbf843d27"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 1f92761c0be8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-ddbb6fb7bd6975151d125412349de656d5834509446663c7eed4067c68b5a594)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-3c208b2409b7809404c68fc3f59f728abee992a102500c3974eac7bf554e196c"></a>

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

<a id="canonical-c500feede308f6dfbffd44a2101586e3b3ca3d61f8c4454c0c2b795b250a434c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 1f92761c0be8 / 3

<a id="canonical-c7d2ac70734f0d651477784d9812d2efdfd0ce5d094ad9baa825f05427a7029d"></a>

<a id="canonical-e25b498f7001bfc8a3dc8418f733ea4592db8a6a1fe224708e1c42d11c3e444b"></a>

## kind property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 1f92761c0be8 / 4

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

<a id="canonical-1fcffcad05f7b6a97c8172f2835721688eb0a792bee128ba9c7e455c13d558c4"></a>

<a id="canonical-6e9a4ed5a61660b91a87d028575e60b233b27cf6692d7aace571b47d28da2c29"></a>

## name property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 1f92761c0be8 / 5

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

<a id="canonical-626f18c7f2d325cecdc6299b216fd3c4369698f9ca5dc41b33e77c97041057db"></a>

<a id="canonical-d9ce819a510dd131d1919dc1900c5aaa6734f511d77d45a90369dc820241f7c9"></a>

## namespace property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 1f92761c0be8 / 6

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

<a id="canonical-8384256c31b9b3bb289bd19d7cb3739e19ace35c6162d93081484aaa90d571b6"></a>

<a id="canonical-29324b7080f619cdbf7d45c1e7411268057fd555e48dbd07fab2baf2252f3d01"></a>

## tenant property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 1f92761c0be8 / 7

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

<a id="canonical-3ec27f0ddc69198c47ab80b63750c0ddcf4da11f05d6b2ec9746bea15adf87c3"></a>

<a id="canonical-5e30fe40dde96ca7f099151b269708a4a212856cd30264091b09ff3b0f387576"></a>

## uid property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 1f92761c0be8 / 8

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

<a id="canonical-db07d15fb6e5f87deefc374104ec6ab0d7e35f304017ad991af222597b87d696"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 1f92761c0be8 / 9

- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-ddbb6fb7bd6975151d125412349de656d5834509446663c7eed4067c68b5a594)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-52cc79e99166ebccce4e5a198deeb5e7a4045546419946c41882e724d826e7e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57978bb69fd9b6ebe1457b27cd7792020b06e41654f1c8ba34ee5ef764ba2635"></a>

## policy_based_challenge.rule_list.rules.spec.ip_prefix_list — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / ea875a7ad641 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-e43e005386d4150cc92ac589d997b11b51df03e351bb7e234a1d147c50a244b6"></a>

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

<a id="canonical-5688917b540522acd516ba26396b98cddc646c4815dfa484f2be09c53361c9b5"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / ea875a7ad641 / 3

<a id="canonical-0bc362f96a5ae033f623fed0c47119eb927dce001568294c89b2c997251b2f4f"></a>

<a id="canonical-63b430b01a2e3f793030bb6706c02fcecca14054727a84db65febf96a1b493dc"></a>

## invert_match property — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / ea875a7ad641 / 4

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

<a id="canonical-1f394728b466b2081bd04911c39ba5d03efbf52ac51a25761ad1ecbcf387a5ea"></a>

<a id="canonical-e9e0a0a38cf5ee7badc16d26a1d7e8bda826ba910c8d22c434a808d2a3f424c9"></a>

## ip_prefixes property — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / ea875a7ad641 / 5

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

<a id="canonical-449edba2fbaf936d253fd3ab1cdf62e88bd6d24c61fb1ce5e921c3b0b552912e"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / ea875a7ad641 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-947327aa7fee2ee2917a57bfd86fd202b9de6a5104002e194ff7d5e557d36fa1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fa06f589a6af499cbc9407584820fd0c41e4ce54f95143c967eca573b5d9de9"></a>

## policy_based_challenge.rule_list.rules.spec.path — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-478dbf24461e04a803b007290ea7fbf32e84d29957a2283b5c93f9fee18d6460"></a>

Type: `"object"`. single nested block, Optional.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-8092519495b1ac995ef95936cf9263276d67b8a2158ad674f6d2e2ba9616e282"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 3

<a id="canonical-ba9ebfdfac56b9d75133ebb1b656d92241024a9252ff3a23ae2c81bc7262c3bd"></a>

<a id="canonical-60b776eac72ec2bd812dfbd3fd345799a9732e272b9fbd8a36faa5f3a98d47e4"></a>

## encoded_path_matcher property — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 4

Type: `"bool"`. Optional.

Match against the encoded, escaped path.

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

<a id="canonical-8e299ea05ecb05ce7cca71ee66273d43e94d0c07b9cb6ea10b015479d1e7386f"></a>

<a id="canonical-654303f0c2797a55f30457bba9f47686aa0bec8f844676f1890c6e9948525b0c"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 5

Type: `["list", "string"]`. Optional.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-13173c550894d2c12c486bffa5d6ab8e9f4c630f70cbc57fca819216075652ad"></a>

<a id="canonical-f36cea47a530a821a0cc67ef0626704740d1a0351e2549a275802719a1c87d55"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 6

Type: `"bool"`. Optional.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-7337db3481be129622bec618d1cdf8264e87f2f9980ac37c5ca0b8de0b8f5d35"></a>

<a id="canonical-f0775dbde6f7abd236aa30c99966a8f6e61cfb12c0d008f49e596099d8a27312"></a>

## prefix_values property — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 7

Type: `["list", "string"]`. Optional.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5c4a5f03aab3fc24f29d51b843efad9c2acdee933db73e2c50633415b95f6219"></a>

<a id="canonical-8e883e66ee162452c180b5038a9a5deeac9f0f4745ff389d4de1fed21c5905c4"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 8

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

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

<a id="canonical-29b8645677155214ad4fe47d34dd71e8801333c5e9844e5df0793f5eea98cca8"></a>

<a id="canonical-ed7f485bdd8f45ac854263031a62c2b9b9e5ea891bf384d3ad43b33851114f98"></a>

## suffix_values property — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 9

Type: `["list", "string"]`. Optional.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-76e8d268386b359a38b134f0d2379f00d47a77ac34f6ea739ec41bec35c8db40"></a>

<a id="canonical-bd263d09fed26e233ea786a01f63da25e2dfc2486b48be3ab0cdee7a7c47a428"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 10

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

<a id="canonical-6ca27aa57acf47c9aed84abcbccf6e393cd83cc39265449b982fa1caa5fcb9ea"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.path / 204dfb40089e / 11

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-525f98c17d96a8ecebb8ee01d843ddc4b6156d0ec30ed20ae6f6304bc150b9b8"></a>

## policy_based_challenge.rule_list.rules.spec.query_params — policy_based_challenge.rule_list.rules.spec.query_params / 139bb90f8007 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-a36245e5960244eed2a0809688f16b8991ac3dac36ce1d8bcd6804d1c70364b1"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
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
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-84c6152ed761427332e6face1122c9b95a81518e5ac194bf53e7d012a4d270cd"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params / 139bb90f8007 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-d3c046d8e0a0511ecf8a0315b37c001d352fb9b8076b035a2a042ee8be57cad2): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-022.md#canonical-960cc630edef6b603a8ffeebfa4d3a9dc68a26ee364f41210e78774b01aa5c36): complete subsection reference.

<a id="canonical-c9b5874174e285468b4a7b5f5078b2db7e7923d0666ae8b05b54dbf8fe334484"></a>

<a id="canonical-e4b7aeaa430146c1b188fdba983dfb66ecb64b1a5bdc1b3e17a2894981d4b74b"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.query_params / 139bb90f8007 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-022.md#canonical-e9ce29f6e58aef96d16722c9c124f823f116205a699a441a709dcb42fcd532d0): complete subsection reference.

<a id="canonical-a7e700b596f65bdee45fc5f91114d6aeca1f64f3d334524599b1dcf6eb727da2"></a>

<a id="canonical-d213c99bd2440c17e37dee4e48f58c01394f5979866f368120d508f36c798cef"></a>

## key property — policy_based_challenge.rule_list.rules.spec.query_params / 139bb90f8007 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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

<a id="canonical-ec19dd64502646a1ba48633e5150d0485390953ed9c0d74f08add17488b46401"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params / 139bb90f8007 / 6

- [policy_based_challenge.rule_list.rules.spec.query_params.check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-d3c046d8e0a0511ecf8a0315b37c001d352fb9b8076b035a2a042ee8be57cad2)
- [policy_based_challenge.rule_list.rules.spec.query_params.check_present](resources--http_loadbalancer--reference--group-022.md#canonical-960cc630edef6b603a8ffeebfa4d3a9dc68a26ee364f41210e78774b01aa5c36)
- [policy_based_challenge.rule_list.rules.spec.query_params.item](resources--http_loadbalancer--reference--group-022.md#canonical-e9ce29f6e58aef96d16722c9c124f823f116205a699a441a709dcb42fcd532d0)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d3c046d8e0a0511ecf8a0315b37c001d352fb9b8076b035a2a042ee8be57cad2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3ec83b6870a61f18d0dd6992aba09b69113cd3be99ee06f92b083e0ef70767d"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_not_present — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / 694d883112d3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-022.md#canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-51cd07b63e1fc3b73adc63ea63d536fdf1c39d52747859e8e09025dc7f85843e"></a>

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

<a id="canonical-9a28fad89327c74abcd3e61fcf3d6d3c6fe73b69fa7ed615336d2107d6a4e84b"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / 694d883112d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-364a7cfd4f96b082d7ae6cc921d4df10f64b8d247aee4538f0d093c2ef865fd3"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / 694d883112d3 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-022.md#canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-960cc630edef6b603a8ffeebfa4d3a9dc68a26ee364f41210e78774b01aa5c36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b8ec08b5fb22029428061a3d475d979cc120830bba3646d38f74e6168e42877"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_present — policy_based_challenge.rule_list.rules.spec.query_params.check_present / a626113194e2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-022.md#canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-86bacd858009219ab67a4519965e3e6e20b61444bbf3fc79a1c6a7df0b98b721"></a>

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

<a id="canonical-2c2cb0c66587645897302fc54b6133dde538c5e0b752017aab3b92e25792b1b4"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.check_present / a626113194e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12dd300cc8db473292bbee58f80bd127d29e9055e40a07470781a3ae52fe6a8d"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.check_present / a626113194e2 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-022.md#canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e9ce29f6e58aef96d16722c9c124f823f116205a699a441a709dcb42fcd532d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d4cca1ebdc03d2e761d57ebe238f9153e99e9910d447db1130058c6b4d4130e"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.item — policy_based_challenge.rule_list.rules.spec.query_params.item / 27d5ac8618b5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-022.md#canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-8955d35a6f3d94532e2247c9eddb3e7d3e440674e3c443bcc017254b344b6dd0"></a>

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

<a id="canonical-5d846c171fba78820e59a7f5d9566ee970d13cc8a90e2fd844785e5750aefb93"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.item / 27d5ac8618b5 / 3

<a id="canonical-8095bf8cb6f6cbdc2dff47502fb6dd72f08c951e2e7a5ac109f50f2e2c498526"></a>

<a id="canonical-9117f6fe8ba8ce72feb393a8e2966a557fa9533132ec14b08226236174565845"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.query_params.item / 27d5ac8618b5 / 4

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

<a id="canonical-ad964af2e9a39507aa75da0202c1b379ac9623ff0634e28fc7a46471e8d84d9a"></a>

<a id="canonical-dc6ac4db36d12096912602064eea1cfdb9180fa2130f2d3ccb560c0e35351cf5"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.query_params.item / 27d5ac8618b5 / 5

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

<a id="canonical-5cf03a019c63e3d8324f8eeabb079c55e52ea3ca8c9bd884aa6761b91fbe503f"></a>

<a id="canonical-3a9ef7f0a6fe7992db32409f2f99ef80c0609de2fccca57c1ba45aa5bd72b5aa"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.query_params.item / 27d5ac8618b5 / 6

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

<a id="canonical-f2708abda8c5f77eb20736a558970133a27daf04c231f478e530bcb14e3e7a11"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.item / 27d5ac8618b5 / 7

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-022.md#canonical-f1e47bdcd8a0a4b284312131637153c0919aa212b83529811f15e65224028e33)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5d3c8888dcffc18cc205b64c7baf988e6975e341fb257285637733dcc4349d37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c737bee7f764698616cf904c502cd2f5240608d5de09e295843f1a90129127d8"></a>

## policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / 289b6013c587 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-0b657d11c041495839dc76d0ab2dbdbd09a39ba038a636ededb5e46556530426"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-84273ee1c9f749c11928bd50c009dcc9e75c8d65e4418bfe5d8b2eb371213a09"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / 289b6013c587 / 3

<a id="canonical-495798e3f179f30284e8fd5073fb20f9e7588ad928a627e4330bf3a27206dc63"></a>

<a id="canonical-16ee44717d64cfcbcfc66aad9338ee4a1491e0e0b171b82815c563d69fbe2734"></a>

## classes property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / 289b6013c587 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a0727eba626637c9b5e3e4bdb85a3f455eaeea36712b7dd03ad4b0667e52a4cc"></a>

<a id="canonical-cfb0880d62d7c44e6b7fb3b0200b7fc612b97791b6a2616211b3eb313065b6c3"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / 289b6013c587 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-61726323b6a381fa38d41b6aae215e1854deb931e45f1d8f99860a8b16a9861c"></a>

<a id="canonical-03c0ae9c2cf7b7dd1e0595533090ac6cf0773a6fdcb47be668233eb75ca40402"></a>

## excluded_values property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / 289b6013c587 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e0b3e9c44429d7d3732148a21a2abe1f8f5a6fcead520d8a2c8cba66edb5e837"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / 289b6013c587 / 7

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3ee94511dbac7c8b39832e21b65490680ec7d621b40e6b37977278368aaaed41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41beb4d2b90b0ff9440be2fc9c75844af8954c1d004c8e3d55f69e8bd6707f27"></a>

## policy_based_challenge.temporary_user_blocking — policy_based_challenge.temporary_user_blocking / 163cd8ff092b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-67f08214210ed65c3a39a83f6abb1a546d69ceef27c11397e5986dc777173018"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration for temporary user blocking resulting from user behavior analysis. When
Malicious User Mitigation is enabled from service policy rules, users' accessing the application
will be analyzed for malicious activity and the configured mitigation actions will be taken on..

Upstream description:

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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
temporary_user_blocking {
  # Configure direct properties listed below.
}
```

<a id="canonical-ab8136e5a13417cfb8c95d52249bbd5fdcad930a75e6f12d6cae5a5065f1bb5c"></a>

## Direct properties — policy_based_challenge.temporary_user_blocking / 163cd8ff092b / 3

<a id="canonical-d88332aa789b8f16a52365d346c6043e1d3a5bde90001119da768e162451930b"></a>

<a id="canonical-c9279dc89aa79d02bea67b61e110a4dbb31cbc655be05dd86df920a811bceb43"></a>

## custom_page property — policy_based_challenge.temporary_user_blocking / 163cd8ff092b / 4

Type: `"string"`. Optional.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Upstream description:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in Base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3d272e6e7530bd62913db8094e90be53d8a65f5d991261e5aa7b7b8e1a7198e3"></a>

## Next pages — policy_based_challenge.temporary_user_blocking / 163cd8ff092b / 5

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b77fbe9c5547545075857dc9a5f94365365f416b679714bc283157f5422e0a84"></a>

## protected_cookies — protected_cookies / 8f799ef4d73b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- protected_cookies

<a id="canonical-06a36d15b37876ba54c877f4d9eb07bde5ef3ad759a3c101fd730ca76564cab0"></a>

Type: `"object"`. list nested block, Optional.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("disable_tampering_protection",
    "enable_tampering_protection"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_cookies {
  # Configure direct properties listed below.
}
```

<a id="canonical-f8be34111bf57bce8709fabcfbd31ced459851623f914ea2e89a5b4ac7a0194e"></a>

## Direct properties — protected_cookies / 8f799ef4d73b / 3

- [add_httponly](resources--http_loadbalancer--reference--group-022.md#canonical-f0c3768b1d45f65a15cb956978b470e21f575b9e8036c20eb3df57b4382767a5): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-022.md#canonical-ad3f20511eb775278d5a4aeaff7d05096fc161208b0511c6825b685852829f87): complete subsection reference.

- [disable_tampering_protection](resources--http_loadbalancer--reference--group-022.md#canonical-e7846a2ba5f98152cb7f729338767d59b402facc894b155440d6557568cd4eeb): complete subsection reference.

- [enable_tampering_protection](resources--http_loadbalancer--reference--group-022.md#canonical-98e6ac2ee19d8b79c2e1973dcf60aef7810613bfe76e10be7625a6dd5342dfa5): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-022.md#canonical-239948c9a0a2d624acd57352e92bd8668af6ffea4bf11027a6b5daada0016411): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-022.md#canonical-17b93ba76853276fa88701b16bf11619d22ec6b98907607de6cdec0e590c2555): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-023.md#canonical-6a9ea7044c39de2c7799d88e2c40e099a07a6a498e14fbbd6a71d471cf56333a): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-023.md#canonical-83689452932005334a2e222aa6e5083a574dddab755ceeefaa4789cfedb59cee): complete subsection reference.

<a id="canonical-7b636c32f8405d380588ff16f165c60631c55b4112f28984df0156574849f503"></a>

<a id="canonical-4bd4a66290ea770c005c84d7db0696dac7963566e697aeb10b8e16964526d762"></a>

## max_age_value property — protected_cookies / 8f799ef4d73b / 4

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-797fad6940775be31edf97732aa8dc512f36b1213d22170063d81fe4ffaaec56"></a>

<a id="canonical-e58aeb5bb1b397f04d3333ca434389cd2c3794914a9956d94f10d9d679e884dc"></a>

## name property — protected_cookies / 8f799ef4d73b / 5

Type: `"string"`. Optional.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-023.md#canonical-8070146ed034d844e58f95f940e633f4ae0b62926b81b715b9eb49b9804a1e1c): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-023.md#canonical-448b6ac4d1281770184a479374f166aec84e61611e646e69b89406b18bc26375): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-023.md#canonical-e20afc57f06bad2eea45c617dc036b74874c14831a5709620c66058ec84287a0): complete subsection reference.

<a id="canonical-18d63f068feb243faa30e7411136a37415a78aa55cf2b9b4a9494a20b85e10ad"></a>

## Next pages — protected_cookies / 8f799ef4d73b / 6

- [protected_cookies.add_httponly](resources--http_loadbalancer--reference--group-022.md#canonical-f0c3768b1d45f65a15cb956978b470e21f575b9e8036c20eb3df57b4382767a5)
- [protected_cookies.add_secure](resources--http_loadbalancer--reference--group-022.md#canonical-ad3f20511eb775278d5a4aeaff7d05096fc161208b0511c6825b685852829f87)
- [protected_cookies.disable_tampering_protection](resources--http_loadbalancer--reference--group-022.md#canonical-e7846a2ba5f98152cb7f729338767d59b402facc894b155440d6557568cd4eeb)
- [protected_cookies.enable_tampering_protection](resources--http_loadbalancer--reference--group-022.md#canonical-98e6ac2ee19d8b79c2e1973dcf60aef7810613bfe76e10be7625a6dd5342dfa5)
- [protected_cookies.ignore_httponly](resources--http_loadbalancer--reference--group-022.md#canonical-239948c9a0a2d624acd57352e92bd8668af6ffea4bf11027a6b5daada0016411)
- [protected_cookies.ignore_max_age](resources--http_loadbalancer--reference--group-022.md#canonical-17b93ba76853276fa88701b16bf11619d22ec6b98907607de6cdec0e590c2555)
- [protected_cookies.ignore_samesite](resources--http_loadbalancer--reference--group-023.md#canonical-6a9ea7044c39de2c7799d88e2c40e099a07a6a498e14fbbd6a71d471cf56333a)
- [protected_cookies.ignore_secure](resources--http_loadbalancer--reference--group-023.md#canonical-83689452932005334a2e222aa6e5083a574dddab755ceeefaa4789cfedb59cee)
- [protected_cookies.samesite_lax](resources--http_loadbalancer--reference--group-023.md#canonical-8070146ed034d844e58f95f940e633f4ae0b62926b81b715b9eb49b9804a1e1c)
- [protected_cookies.samesite_none](resources--http_loadbalancer--reference--group-023.md#canonical-448b6ac4d1281770184a479374f166aec84e61611e646e69b89406b18bc26375)
- [protected_cookies.samesite_strict](resources--http_loadbalancer--reference--group-023.md#canonical-e20afc57f06bad2eea45c617dc036b74874c14831a5709620c66058ec84287a0)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f0c3768b1d45f65a15cb956978b470e21f575b9e8036c20eb3df57b4382767a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7c979b01467d89b0708dde05bbb434c1b4119f377b4254e88f51705bf78e4b1"></a>

## protected_cookies.add_httponly — protected_cookies.add_httponly / f87856606506 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- protected_cookies.add_httponly

<a id="canonical-e5be07ac788637e4ea66d4a0d1d204bd5d3e54e3a46dd41f2c88d2f4a71e0068"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-a07d01f054771053a276c94914d8da7db8984f2f063ce62091dc91e0c8c3a14f"></a>

## Direct properties — protected_cookies.add_httponly / f87856606506 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8689d4907a08d97012b909dbfa150f7ef48af2ea55914c307a3134c7e176c85"></a>

## Next pages — protected_cookies.add_httponly / f87856606506 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ad3f20511eb775278d5a4aeaff7d05096fc161208b0511c6825b685852829f87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ca7412ee23ddf846c91a005705832aaf02cb8af84a4636a5740d75405045e7a"></a>

## protected_cookies.add_secure — protected_cookies.add_secure / ac5aad227d73 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- protected_cookies.add_secure

<a id="canonical-080b29c73962fe85b359a931280c0b0926f5724c8825995dc8c90aa7d5e2e32f"></a>

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
add_secure = {}
```

<a id="canonical-f5d8b8803ebfcccdbe55b5e5219330a5cb692309af507329996e67361ee87451"></a>

## Direct properties — protected_cookies.add_secure / ac5aad227d73 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-35b9a889cdb95b9376e0c116e6215b97938af8029f965cfab299077089e54c3c"></a>

## Next pages — protected_cookies.add_secure / ac5aad227d73 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e7846a2ba5f98152cb7f729338767d59b402facc894b155440d6557568cd4eeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5382c470a12161d94008b04120a10598d7f6910667460180b76f053e870dbd36"></a>

## protected_cookies.disable_tampering_protection — protected_cookies.disable_tampering_protection / 24b04637d10f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- protected_cookies.disable_tampering_protection

<a id="canonical-75067a9a358aa2ab184bdbe8dcfecbdf541ec3a3ebd20f9c9cfe6595d5e725f9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tampering protection.

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
disable_tampering_protection = {}
```

<a id="canonical-dd3b341deb5aa91e85156e480bfedfe005578bad8d96578491971e54dd256564"></a>

## Direct properties — protected_cookies.disable_tampering_protection / 24b04637d10f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ef04d68b219814815184cca8c1953e17ce4e9dc81570e272d6e967cab90d79b"></a>

## Next pages — protected_cookies.disable_tampering_protection / 24b04637d10f / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-98e6ac2ee19d8b79c2e1973dcf60aef7810613bfe76e10be7625a6dd5342dfa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d29c13e46e58f44ee567e6700efc06599db8004a90666163da0e2dc7907c549a"></a>

## protected_cookies.enable_tampering_protection — protected_cookies.enable_tampering_protection / e6d7235c6e59 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- protected_cookies.enable_tampering_protection

<a id="canonical-028a71563512f1e8726c53f80679765ba4a4bad4b5d731ee84c1f6a5ce33b130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tampering protection.

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
enable_tampering_protection = {}
```

<a id="canonical-f59316f5f8ca47ac06924c3ce193b9647ba2d34cc656a75da10e1626c2f7433c"></a>

## Direct properties — protected_cookies.enable_tampering_protection / e6d7235c6e59 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34783f3d7a2c7da1e993edd8da5cf0b41dfef1bddbe9f647dd128c5066a7125c"></a>

## Next pages — protected_cookies.enable_tampering_protection / e6d7235c6e59 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-239948c9a0a2d624acd57352e92bd8668af6ffea4bf11027a6b5daada0016411"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1be5fa68a507b2b2eb5e89ae451702f63a9f09627a83fae155c796baf92e48a0"></a>

## protected_cookies.ignore_httponly — protected_cookies.ignore_httponly / d950f6b22cc5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- protected_cookies.ignore_httponly

<a id="canonical-899466b70d37242e011f871f1d3dcd8fed2677a16cc3ebbe2fa11d3bd350b1eb"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-62d450e689062de48f32a9152240588657ef1ebf2d910ecf062b5b473c112d4a"></a>

## Direct properties — protected_cookies.ignore_httponly / d950f6b22cc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5393cc390a2e3b76408d099507a58760f4f80e12cbbed501f7f1f644c91c4d70"></a>

## Next pages — protected_cookies.ignore_httponly / d950f6b22cc5 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-17b93ba76853276fa88701b16bf11619d22ec6b98907607de6cdec0e590c2555"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
