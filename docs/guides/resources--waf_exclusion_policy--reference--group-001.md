---
page_title: "xcsh_waf_exclusion_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy reference."
---

# xcsh_waf_exclusion_policy reference

<a id="canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d1500ebe6ea1688244d6762c9ec7dabfe79bbde99891b0797f59463f34e14be"></a>

## Property reference — Property reference / c13f0a4cc1a4 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- Property reference

<a id="canonical-b2ddaf45bbb30c2a658dd4b57bea7f673e31ce0f1305f05ebba646217ac2e1dc"></a>

## Direct properties — Property reference / c13f0a4cc1a4 / 3

<a id="canonical-95551cfe7885e9711531a799be43fdff2718ad9edaceafd635c40bd14bbd4fb1"></a>

<a id="canonical-a4c81e5dd42fbbf43088aad8743cdf85690f3e9fa27ca120fbd38fd98effa0f5"></a>

## annotations property — Property reference / c13f0a4cc1a4 / 4

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

<a id="canonical-8bb6fa1d43235720eb725f3d8cba4097d71b3e4925301a23019e5d6cedbd40ae"></a>

<a id="canonical-2f5554c42e9a0577c57811a0bec347f16af49b10687a7c71b277fcb4dac58bff"></a>

## description property — Property reference / c13f0a4cc1a4 / 5

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

<a id="canonical-92400322daec0b3bf565d7896cff1dbb1d40bf4a0b2e5e8cea19313265116de3"></a>

<a id="canonical-84495b6aa889c5d30d45974486c38e45513d73b17cb5b5595d4678dd60c17e03"></a>

## disable property — Property reference / c13f0a4cc1a4 / 6

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

<a id="canonical-1ad95fc5306d7b34c025bdd088c05de72bf2ee685a1bec6a12d901f2fd18c446"></a>

<a id="canonical-f79f8149145ad583c113b5595e3d31673117498dcb127643c50c637d3839a03a"></a>

## id property — Property reference / c13f0a4cc1a4 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-eb604c14e2f60b9fb47e3c63e021658c285a3b7a962662bec6930e1cfb6e879e"></a>

<a id="canonical-bb243bc2f97cd4c91ad0bdbd370cadcaa33fc9ec6660704410aced13bb7e6bda"></a>

## labels property — Property reference / c13f0a4cc1a4 / 8

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

<a id="canonical-881ed1245a36a3272e61e10cad1ff6e9d90565e0522e886b14c29bc002d20425"></a>

<a id="canonical-9163554dd0179afb544f6d16e0802797994d52b3bb52064baa7c02cd498fc30c"></a>

## name property — Property reference / c13f0a4cc1a4 / 9

Type: `"string"`. Required.

Name of the WAF Exclusion Policy. Must be unique within the namespace.

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

<a id="canonical-9a191c5426aa6aa1222f9b2e4a99208d22bb57d1fc42284f115cb2b06d375bc8"></a>

<a id="canonical-d2e305b8a944ba3fbdfac1292bdabf11b1993cf34a5d7bd166a5931300e0a785"></a>

## namespace property — Property reference / c13f0a4cc1a4 / 10

Type: `"string"`. Required.

Namespace where the WAF Exclusion Policy is created.

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

- [timeouts](resources--waf_exclusion_policy--reference--group-001.md#canonical-e4a32f3dd7395885d1dfa832332805bcdd1089508d1dce7f4acdaca2d952881e): complete subsection reference.

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846): complete subsection reference.

<a id="canonical-a0c1ed716bece68421395156363afd8fe525ba8137ff715a355fe69edec51314"></a>

## All schema paths — Property reference / c13f0a4cc1a4 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--waf_exclusion_policy--reference--group-001.md#canonical-95551cfe7885e9711531a799be43fdff2718ad9edaceafd635c40bd14bbd4fb1) |
| `description` | [description](resources--waf_exclusion_policy--reference--group-001.md#canonical-8bb6fa1d43235720eb725f3d8cba4097d71b3e4925301a23019e5d6cedbd40ae) |
| `disable` | [disable](resources--waf_exclusion_policy--reference--group-001.md#canonical-92400322daec0b3bf565d7896cff1dbb1d40bf4a0b2e5e8cea19313265116de3) |
| `id` | [id](resources--waf_exclusion_policy--reference--group-001.md#canonical-1ad95fc5306d7b34c025bdd088c05de72bf2ee685a1bec6a12d901f2fd18c446) |
| `labels` | [labels](resources--waf_exclusion_policy--reference--group-001.md#canonical-eb604c14e2f60b9fb47e3c63e021658c285a3b7a962662bec6930e1cfb6e879e) |
| `name` | [name](resources--waf_exclusion_policy--reference--group-001.md#canonical-881ed1245a36a3272e61e10cad1ff6e9d90565e0522e886b14c29bc002d20425) |
| `namespace` | [namespace](resources--waf_exclusion_policy--reference--group-001.md#canonical-9a191c5426aa6aa1222f9b2e4a99208d22bb57d1fc42284f115cb2b06d375bc8) |
| `timeouts` | [timeouts](resources--waf_exclusion_policy--reference--group-001.md#canonical-454c38af9b33eeff31f0122821267b692b44f3b626a344026fe30ba8e91bdac4) |
| `timeouts.create` | [timeouts.create](resources--waf_exclusion_policy--reference--group-001.md#canonical-33746edf84d3bb64e9a771651f0c593b9f957f633fe1fb89b0aaed69c795b61a) |
| `timeouts.delete` | [timeouts.delete](resources--waf_exclusion_policy--reference--group-001.md#canonical-7b50fe8103e3a1583c0d576ebc696ff44e941620b6a5b72eb4153495623b17ce) |
| `timeouts.read` | [timeouts.read](resources--waf_exclusion_policy--reference--group-001.md#canonical-6021d152566b1a1030f4fffda13370a3f8ba4acce70241446d89cfdca7692e90) |
| `timeouts.update` | [timeouts.update](resources--waf_exclusion_policy--reference--group-001.md#canonical-756bfb6788065d8f9a4d1e79e470d16d58a6a69e5d5932697d0ff499bbbd784e) |
| `waf_exclusion_rules` | [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-fe5d396bb0f35c89fc9f1cf1dfac86b8ef3035c1db2d76902cf561d7aa30f2b9) |
| `waf_exclusion_rules.any_domain` | [waf_exclusion_rules.any_domain](resources--waf_exclusion_policy--reference--group-001.md#canonical-3c292af3b1a2f260a3a5e707f308ffe2dbebb04c876fd25e350b0471b288b427) |
| `waf_exclusion_rules.any_path` | [waf_exclusion_rules.any_path](resources--waf_exclusion_policy--reference--group-001.md#canonical-c42dd49f1b538b645fb016447d8d0633b5ef5e426f6980fece19b9fba9ab8100) |
| `waf_exclusion_rules.app_firewall_detection_control` | [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-44e317f3b39ea5168ea906b29b868e7c3c5d5564b6465ad68c2caaa0ca301cb3) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-6361f5f346fc96d8417c749d44244e789728973feed8d17e471bf88641afe417) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context](resources--waf_exclusion_policy--reference--group-001.md#canonical-1862c20c89b02255bc77cac72210923c574bba66383fff88ca478fca8ed33e46) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name](resources--waf_exclusion_policy--reference--group-001.md#canonical-f2b461b1fb989e7c1dda57370c9df8b30c2748919a6eaa63353632adde1ceec8) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](resources--waf_exclusion_policy--reference--group-001.md#canonical-f8ca2dcd876cc2d6210852c06bf6f851362760325e98b1f1391728467d4d8d1a) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-ef1ed4f9b3b3fb644dd1f5788f0de159ed2eafbd19011095818bb685729d130a) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](resources--waf_exclusion_policy--reference--group-001.md#canonical-210407145a427e372089122f3242c6f9344a371b693f361f246a1e0ed75c4f49) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-df5b236b4d132f9bfb5c62a334576557091e701520132de0bcb3622d3c492299) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context](resources--waf_exclusion_policy--reference--group-001.md#canonical-b8238de6db87b1a9aeeb32f3d98948ba4ab551950d95a8e36df77aada37fd37e) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name](resources--waf_exclusion_policy--reference--group-001.md#canonical-eebcf6a126181584bd03f024e0ece66a65033b6586a1a9496c2e1c23930cb257) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id](resources--waf_exclusion_policy--reference--group-001.md#canonical-e1b2a9dcb8560b4e555d7f126b5f391dd18687b616fc83639660f21567885c74) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-09fec633db117d7ec5f45b7113f1d2827bded6c962cc49c6f6c7a719e5c40fc8) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context](resources--waf_exclusion_policy--reference--group-001.md#canonical-609d0d07ef13a20fdf1cf8d83a3c769eb7c397f4b92b55e1969ea7649308d794) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name](resources--waf_exclusion_policy--reference--group-001.md#canonical-1fd8facae94a069f40bfbefa232032a9084334fa4c1d6f7c68f3a5978853d473) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](resources--waf_exclusion_policy--reference--group-001.md#canonical-5bc81af04c06331e2c75c95d27df47feeaaebcab990bf2ac5aad9cb6d9f9e531) |
| `waf_exclusion_rules.exact_value` | [waf_exclusion_rules.exact_value](resources--waf_exclusion_policy--reference--group-001.md#canonical-113e78e59aa065bbfc87391fa598b7ff324263df9e99fd0cddbdf5278461c920) |
| `waf_exclusion_rules.expiration_timestamp` | [waf_exclusion_rules.expiration_timestamp](resources--waf_exclusion_policy--reference--group-001.md#canonical-e43770c4a3c6218a18e4ac1f414fe97ac193c5c9ce4bc4f98fef716e1faf228a) |
| `waf_exclusion_rules.metadata` | [waf_exclusion_rules.metadata](resources--waf_exclusion_policy--reference--group-001.md#canonical-ead7f6e8bb7051ae4a1e100d6d687ea1e547e11e2652ef81e5564f3ddf021b16) |
| `waf_exclusion_rules.metadata.description_spec` | [waf_exclusion_rules.metadata.description_spec](resources--waf_exclusion_policy--reference--group-001.md#canonical-b4bb6263a5e12542109221cb7bfd3d559a3c06bd11c4e1e087d4cb9c1f67858d) |
| `waf_exclusion_rules.metadata.name` | [waf_exclusion_rules.metadata.name](resources--waf_exclusion_policy--reference--group-001.md#canonical-a011eef796ae9c89d48d59512dcaf1589758fad0172c896c82928dbdb66dbdd1) |
| `waf_exclusion_rules.methods` | [waf_exclusion_rules.methods](resources--waf_exclusion_policy--reference--group-001.md#canonical-f3f44cabfdc2d0b21a4979cda5392df14b8fbdce3d3bd36ea8ac9976a4aec3b7) |
| `waf_exclusion_rules.path_prefix` | [waf_exclusion_rules.path_prefix](resources--waf_exclusion_policy--reference--group-001.md#canonical-2969cc02ddde3268253c9479367e8a6cc97fc55f1e9dcb457a20829984821b2f) |
| `waf_exclusion_rules.path_regex` | [waf_exclusion_rules.path_regex](resources--waf_exclusion_policy--reference--group-001.md#canonical-eaf8917613181a8fac0a0fb0c01324818be16438dffba45df8b406afc60d0bda) |
| `waf_exclusion_rules.suffix_value` | [waf_exclusion_rules.suffix_value](resources--waf_exclusion_policy--reference--group-001.md#canonical-060931b47cba28715f76c3c80fb961a63d5c91492439be80388298ae6049f07e) |
| `waf_exclusion_rules.waf_skip_processing` | [waf_exclusion_rules.waf_skip_processing](resources--waf_exclusion_policy--reference--group-001.md#canonical-55ac36caede975a7391e3eabb8e8a0a5b1d36f09bf3b3e0fce1148e96b799e29) |

<a id="canonical-e6f3e6a1af693205528f7d8473f3679d050d3fb6012b741fbc60e39a391eb922"></a>

## Next pages — Property reference / c13f0a4cc1a4 / 12

- [timeouts](resources--waf_exclusion_policy--reference--group-001.md#canonical-e4a32f3dd7395885d1dfa832332805bcdd1089508d1dce7f4acdaca2d952881e)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-e4a32f3dd7395885d1dfa832332805bcdd1089508d1dce7f4acdaca2d952881e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c399b21f75de8ebc0ffb66584253137da30a7e41aefe615012568f5e2133993e"></a>

## timeouts — timeouts / 97a852d2c10f / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- timeouts

<a id="canonical-454c38af9b33eeff31f0122821267b692b44f3b626a344026fe30ba8e91bdac4"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-68013ba22e8fd7ab8450422a893a1917f2ab22d71ad23d0c9e859ad0cc23e9bc"></a>

## Direct properties — timeouts / 97a852d2c10f / 3

<a id="canonical-33746edf84d3bb64e9a771651f0c593b9f957f633fe1fb89b0aaed69c795b61a"></a>

<a id="canonical-e3d8b07f36e18909fe87473b904320fcde4f8e61707059d6ca9d343e5df3891b"></a>

## create property — timeouts / 97a852d2c10f / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7b50fe8103e3a1583c0d576ebc696ff44e941620b6a5b72eb4153495623b17ce"></a>

<a id="canonical-4541ca27db95fcc5df5f378fa69b3834afcdce118621764a4b443bbcd3160eaa"></a>

## delete property — timeouts / 97a852d2c10f / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-6021d152566b1a1030f4fffda13370a3f8ba4acce70241446d89cfdca7692e90"></a>

<a id="canonical-9cbce662ecb70432b2e1874057f5fa85f3fd77f0199c16a4f786b9820c89e29e"></a>

## read property — timeouts / 97a852d2c10f / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-756bfb6788065d8f9a4d1e79e470d16d58a6a69e5d5932697d0ff499bbbd784e"></a>

<a id="canonical-f256f43906b87b2a4ae63d859e0954b51a6cabebf549050bf8c06d09587b80f1"></a>

## update property — timeouts / 97a852d2c10f / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-b6d37064ac4ce26ccc72b7c8de53b5518bdc33c8bdfc36a80d1fa8e3d675d792"></a>

## Next pages — timeouts / 97a852d2c10f / 8

- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05ca05997bad8919f962a90bd213f763605f2e8a384ddeb7d30b540f7292f388"></a>

## waf_exclusion_rules — waf_exclusion_rules / 2c7f289eddcc / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- waf_exclusion_rules

<a id="canonical-fe5d396bb0f35c89fc9f1cf1dfac86b8ef3035c1db2d76902cf561d7aa30f2b9"></a>

Type: `"object"`. list nested block, Optional.

WAF Exclusion Rules. An ordered list of rules.

Upstream description:

An ordered list of rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex"),
  validators.ConflictingListObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_prefix",
    "path_regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
waf_exclusion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-44da0b106dddab1cfa0284ca9b3b52c5a99d7e925cef693c676f0124066991b9"></a>

## Direct properties — waf_exclusion_rules / 2c7f289eddcc / 3

- [any_domain](resources--waf_exclusion_policy--reference--group-001.md#canonical-fc12426f1c87b467495932cd74dac2bc74e531d2b26c88e5b9532e4df8adf6b2): complete subsection reference.

- [any_path](resources--waf_exclusion_policy--reference--group-001.md#canonical-212c86177b01bd727f4074afc6560e2932a8ef1fc81f2bb5b4f7be6a307bb5bc): complete subsection reference.

- [app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7): complete subsection reference.

<a id="canonical-113e78e59aa065bbfc87391fa598b7ff324263df9e99fd0cddbdf5278461c920"></a>

<a id="canonical-2ebb9bf759bce99a7caed805f82e1790f00198ea10a76276c7b40f0d9c7ce5e7"></a>

## exact_value property — waf_exclusion_rules / 2c7f289eddcc / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-e43770c4a3c6218a18e4ac1f414fe97ac193c5c9ce4bc4f98fef716e1faf228a"></a>

<a id="canonical-d2bb325132cbbcea445cdc49f94e1020a2c25d21f900765655bf02d58062091b"></a>

## expiration_timestamp property — waf_exclusion_rules / 2c7f289eddcc / 5

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

- [metadata](resources--waf_exclusion_policy--reference--group-001.md#canonical-b5edcfead7f13076c8c45a56d3e8b246732f16e59de6d84d8c4e7a973206d453): complete subsection reference.

<a id="canonical-f3f44cabfdc2d0b21a4979cda5392df14b8fbdce3d3bd36ea8ac9976a4aec3b7"></a>

<a id="canonical-4a78d3ddbb7ed9870639a366e848bd13048e36f4a5502ba612553695e20ec208"></a>

## methods property — waf_exclusion_rules / 2c7f289eddcc / 6

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

<a id="canonical-2969cc02ddde3268253c9479367e8a6cc97fc55f1e9dcb457a20829984821b2f"></a>

<a id="canonical-3cb0f566b970d631273589322061edf23bea8ed981d419cb15d03b5a92170a5d"></a>

## path_prefix property — waf_exclusion_rules / 2c7f289eddcc / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-eaf8917613181a8fac0a0fb0c01324818be16438dffba45df8b406afc60d0bda"></a>

<a id="canonical-efe3484dc9d7eea0ab49d349124b3776fef2e537ae7fa5ee4eef76635761b48f"></a>

## path_regex property — waf_exclusion_rules / 2c7f289eddcc / 8

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-060931b47cba28715f76c3c80fb961a63d5c91492439be80388298ae6049f07e"></a>

<a id="canonical-cef04ed56a95be27e84daedc000daaafc5fb0f547b7fd648c9468a394bf59eca"></a>

## suffix_value property — waf_exclusion_rules / 2c7f289eddcc / 9

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [waf_skip_processing](resources--waf_exclusion_policy--reference--group-001.md#canonical-655b33145072d2fee1008461c7b96b3264f34bc42fffe456a77bf197a7a33a9b): complete subsection reference.

<a id="canonical-ea025e389f560f508b6daccdaef489742905c828b9f85b913c49c5c816a11dde"></a>

## Next pages — waf_exclusion_rules / 2c7f289eddcc / 10

- [waf_exclusion_rules.any_domain](resources--waf_exclusion_policy--reference--group-001.md#canonical-fc12426f1c87b467495932cd74dac2bc74e531d2b26c88e5b9532e4df8adf6b2)
- [waf_exclusion_rules.any_path](resources--waf_exclusion_policy--reference--group-001.md#canonical-212c86177b01bd727f4074afc6560e2932a8ef1fc81f2bb5b4f7be6a307bb5bc)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- [waf_exclusion_rules.metadata](resources--waf_exclusion_policy--reference--group-001.md#canonical-b5edcfead7f13076c8c45a56d3e8b246732f16e59de6d84d8c4e7a973206d453)
- [waf_exclusion_rules.waf_skip_processing](resources--waf_exclusion_policy--reference--group-001.md#canonical-655b33145072d2fee1008461c7b96b3264f34bc42fffe456a77bf197a7a33a9b)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-fc12426f1c87b467495932cd74dac2bc74e531d2b26c88e5b9532e4df8adf6b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f738e25392f2b46231e8797c3d7946808ade92687a6678100a4233714df5f68a"></a>

## waf_exclusion_rules.any_domain — waf_exclusion_rules.any_domain / 115038692d74 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- waf_exclusion_rules.any_domain

<a id="canonical-3c292af3b1a2f260a3a5e707f308ffe2dbebb04c876fd25e350b0471b288b427"></a>

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
any_domain = {}
```

<a id="canonical-268bad5d45fb87ab0302cf2bb8b1b06c6989a4a8296f865f7c1121ffb642d761"></a>

## Direct properties — waf_exclusion_rules.any_domain / 115038692d74 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a73bc25fe55d7ed82b58959d400cf1133b77f566e6189096d197500043cb1cdd"></a>

## Next pages — waf_exclusion_rules.any_domain / 115038692d74 / 4

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-212c86177b01bd727f4074afc6560e2932a8ef1fc81f2bb5b4f7be6a307bb5bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cad6c49697dfc0a8bf601905b62b9cfa6927327c7c34192946d8f0fe8446f158"></a>

## waf_exclusion_rules.any_path — waf_exclusion_rules.any_path / ed629652df39 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- waf_exclusion_rules.any_path

<a id="canonical-c42dd49f1b538b645fb016447d8d0633b5ef5e426f6980fece19b9fba9ab8100"></a>

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
any_path = {}
```

<a id="canonical-83a5d6fcdb6607f409d4ecb6a361d0265ff74bb5fd8880323c12d132b2475460"></a>

## Direct properties — waf_exclusion_rules.any_path / ed629652df39 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f9cb9c756b2830505ecc44ce0144015edf9826ae8a16c3cb0971e89a8809451"></a>

## Next pages — waf_exclusion_rules.any_path / ed629652df39 / 4

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e92b0a6361dab4f85462bd17ceab8e8d7bf7df61e905b6e6e9feaf4605fa903e"></a>

## waf_exclusion_rules.app_firewall_detection_control — waf_exclusion_rules.app_firewall_detection_control / 39138a434fe8 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- waf_exclusion_rules.app_firewall_detection_control

<a id="canonical-44e317f3b39ea5168ea906b29b868e7c3c5d5564b6465ad68c2caaa0ca301cb3"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-c6da3948f38258cb28df068625bd3df6f8d18ab84b3f4a22a311d6f0a26aece1"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control / 39138a434fe8 / 3

- [exclude_attack_type_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-0e5ddad18eff2c365925e10bf87909a41e1ebe219cc4c21bced3fa5b7c586763): complete subsection reference.

- [exclude_bot_name_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-8bd88d072b0f86b1f22c232ae900b9f367be8015839650f12e8cfe2074a38b5e): complete subsection reference.

- [exclude_signature_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-6c3b3916ab591f2c011d2dfa2664c11a3e69ba9cb378bdbdf0e8540a90606978): complete subsection reference.

- [exclude_violation_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-5f161400357222a21c70ccc7d27a637e1e2a37b595a97f2c9060e00e8795d35c): complete subsection reference.

<a id="canonical-ab6ffdba345638e626346aabdccb13a40ff4898894da67b20e344436aa1a01f4"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control / 39138a434fe8 / 4

- [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-0e5ddad18eff2c365925e10bf87909a41e1ebe219cc4c21bced3fa5b7c586763)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-8bd88d072b0f86b1f22c232ae900b9f367be8015839650f12e8cfe2074a38b5e)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-6c3b3916ab591f2c011d2dfa2664c11a3e69ba9cb378bdbdf0e8540a90606978)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-5f161400357222a21c70ccc7d27a637e1e2a37b595a97f2c9060e00e8795d35c)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-0e5ddad18eff2c365925e10bf87909a41e1ebe219cc4c21bced3fa5b7c586763"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00c34ef8c3586ae494926e0a96190bfaed385833c13dd6208364e2f3c4e20ddf"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 39debfbb0e5c / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-6361f5f346fc96d8417c749d44244e789728973feed8d17e471bf88641afe417"></a>

Type: `"object"`. list nested block, Optional.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-6ca2cffa891594a1b36f827e6701be1a7d9917122d861b429d7c92411c8a2cc8"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 39debfbb0e5c / 3

<a id="canonical-1862c20c89b02255bc77cac72210923c574bba66383fff88ca478fca8ed33e46"></a>

<a id="canonical-55c3c26c9ae21d566629af246890bddbfa9fc4283740762fb957894a6fe231da"></a>

## context property — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 39debfbb0e5c / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f2b461b1fb989e7c1dda57370c9df8b30c2748919a6eaa63353632adde1ceec8"></a>

<a id="canonical-5e580abfa43ae50431a1bc80483f0374ef8a39acecbe6ca96fc7abcff396566c"></a>

## context_name property — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 39debfbb0e5c / 5

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-f8ca2dcd876cc2d6210852c06bf6f851362760325e98b1f1391728467d4d8d1a"></a>

<a id="canonical-3952b987188b5d4b570bd858703e88b6d27257521c08181c0033377a91d0e696"></a>

## exclude_attack_type property — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 39debfbb0e5c / 6

Type: `"string"`. Optional.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fc9643429fd87d97f85e1c5c9ac45ca202c30ad99c974555c99c0498d9756419"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 39debfbb0e5c / 7

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-8bd88d072b0f86b1f22c232ae900b9f367be8015839650f12e8cfe2074a38b5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c590481bd3cc9fd95ba3747a92f1a0f803cd653909b0c229496967e30a669f78"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts — waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts / ed604e1c6d31 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-ef1ed4f9b3b3fb644dd1f5788f0de159ed2eafbd19011095818bb685729d130a"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-effa746a750d800cd16a0f3b52f9758fc254396e84c98b11ddd99243dcc11752"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts / ed604e1c6d31 / 3

<a id="canonical-210407145a427e372089122f3242c6f9344a371b693f361f246a1e0ed75c4f49"></a>

<a id="canonical-9815e387b15e6b6a1fb2e6225d5dc09c893739ec6f3779b116ee4eaa29061a63"></a>

## bot_name property — waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts / ed604e1c6d31 / 4

Type: `"string"`. Optional.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-e4ee95158eb3e1dc10ef479eedacfaa65ca888ba4ec5e2286dae14e88fe027fd"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts / ed604e1c6d31 / 5

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-6c3b3916ab591f2c011d2dfa2664c11a3e69ba9cb378bdbdf0e8540a90606978"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfbd64107613321c45020eed36d7bcfc77883389ef8b12a1a481790c49990dbe"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / 67c259d3d0bd / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-df5b236b4d132f9bfb5c62a334576557091e701520132de0bcb3622d3c492299"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-38d2e3e95cd95416ef334fcda749819987578d3ea026cf8b4cb2fb80a0576303"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / 67c259d3d0bd / 3

<a id="canonical-b8238de6db87b1a9aeeb32f3d98948ba4ab551950d95a8e36df77aada37fd37e"></a>

<a id="canonical-b62af73d27e094f38cb5453e6988d38774876d36e644e84ba81805f72709cbbe"></a>

## context property — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / 67c259d3d0bd / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-eebcf6a126181584bd03f024e0ece66a65033b6586a1a9496c2e1c23930cb257"></a>

<a id="canonical-8f1c94e75f238d9a4d8ba0c1bd5dc758af5e6075a5ed5ffebb717b5b433082e0"></a>

## context_name property — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / 67c259d3d0bd / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-e1b2a9dcb8560b4e555d7f126b5f391dd18687b616fc83639660f21567885c74"></a>

<a id="canonical-8c2a3d0885de03ee44b6ebc061a3984376fa2d23071abf6e2a9b4b3cd22d59d7"></a>

## signature_id property — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / 67c259d3d0bd / 6

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-19d145446376d55758637ceedaf5b54f6f8dc8febcd6683e1c0f623b78c0d7df"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / 67c259d3d0bd / 7

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-5f161400357222a21c70ccc7d27a637e1e2a37b595a97f2c9060e00e8795d35c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d0a155d394fe0bd429acf1390f7cbf3adb576f76d6f9b74423b38e337345cdc"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 1e75fab0a55a / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-09fec633db117d7ec5f45b7113f1d2827bded6c962cc49c6f6c7a719e5c40fc8"></a>

Type: `"object"`. list nested block, Optional.

Violations to be excluded for the defined match criteria.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-fafdff0be6c117d5ca39f60a443fe74ca6f31927388050fae836f421331c0fdd"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 1e75fab0a55a / 3

<a id="canonical-609d0d07ef13a20fdf1cf8d83a3c769eb7c397f4b92b55e1969ea7649308d794"></a>

<a id="canonical-98af13ee1c867cd68fc617cab1c6aece0c68b0d384a6dcfbe282c20826a707db"></a>

## context property — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 1e75fab0a55a / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1fd8facae94a069f40bfbefa232032a9084334fa4c1d6f7c68f3a5978853d473"></a>

<a id="canonical-2c770b0e26d9536f09ca98a0055a69c8dacea8d5228176ba520175b1886e60d1"></a>

## context_name property — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 1e75fab0a55a / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-5bc81af04c06331e2c75c95d27df47feeaaebcab990bf2ac5aad9cb6d9f9e531"></a>

<a id="canonical-77d49d9bfba3bdb53929ca177dfb5e4567a947e3c4091cac89a5502cd1662b1c"></a>

## exclude_violation property — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 1e75fab0a55a / 6

Type: `"string"`. Optional.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fabbf2f7fa4c488e8a9c6f3adb239de7a5b8c0d8a62f886319be35072f092766"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 1e75fab0a55a / 7

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-cff60fc925822cffdf3ab6435ccc0c1daff22f6664b115f3256cfda9187ba2b7)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-b5edcfead7f13076c8c45a56d3e8b246732f16e59de6d84d8c4e7a973206d453"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20c605cd0762540ef1834a7bef5e5e391b7e8bdb553715a69d6428fdd751b68b"></a>

## waf_exclusion_rules.metadata — waf_exclusion_rules.metadata / 75d35a57f476 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- waf_exclusion_rules.metadata

<a id="canonical-ead7f6e8bb7051ae4a1e100d6d687ea1e547e11e2652ef81e5564f3ddf021b16"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-96b74b9067168033895f861576b83909d4167b72741327f06a4e16f9aca3fe43"></a>

## Direct properties — waf_exclusion_rules.metadata / 75d35a57f476 / 3

<a id="canonical-b4bb6263a5e12542109221cb7bfd3d559a3c06bd11c4e1e087d4cb9c1f67858d"></a>

<a id="canonical-a12a1fa0227a24095cee1d686626dce3e0db5e6a43426551f9cdfe2f3cab444e"></a>

## description_spec property — waf_exclusion_rules.metadata / 75d35a57f476 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-a011eef796ae9c89d48d59512dcaf1589758fad0172c896c82928dbdb66dbdd1"></a>

<a id="canonical-c1979a181042e255e2cd825642ca4e062b0c6f8776d0c4665fb676afe9346d70"></a>

## name property — waf_exclusion_rules.metadata / 75d35a57f476 / 5

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

<a id="canonical-e55b8325aaa2404efb69bc27310578c572774c92e4888d36598615d52720a134"></a>

## Next pages — waf_exclusion_rules.metadata / 75d35a57f476 / 6

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-655b33145072d2fee1008461c7b96b3264f34bc42fffe456a77bf197a7a33a9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40b9282f8ca8eff13ef79417b8434ed609b80ee96313b8f50059431ba85523af"></a>

## waf_exclusion_rules.waf_skip_processing — waf_exclusion_rules.waf_skip_processing / cb80400a2263 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- waf_exclusion_rules.waf_skip_processing

<a id="canonical-55ac36caede975a7391e3eabb8e8a0a5b1d36f09bf3b3e0fce1148e96b799e29"></a>

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
waf_skip_processing = {}
```

<a id="canonical-e8d62bae1ca6dee09233be1fe5744c2dc0b25afafbb8965d2173aff083a26be9"></a>

## Direct properties — waf_exclusion_rules.waf_skip_processing / cb80400a2263 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-61a924503e8472eb0ca1c9cee17623c37d31927bb45cdfba4616116f05680708"></a>

## Next pages — waf_exclusion_rules.waf_skip_processing / cb80400a2263 / 4

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3ae118acad1f904d215e4591c24a5fa3078679efa5b9d86f5e4dc6ed9044e846)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
