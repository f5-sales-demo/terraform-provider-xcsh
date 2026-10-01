---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-17f106143a67f1c31413dd33f0b3c50ad87d7b3ac438ef4baef311c2fb33f917"></a>

## custom_security_group — custom_security_group / 4630f3b2ba8a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- custom_security_group

<a id="canonical-59152abbfa2de683070f08747663621c3f154c22740e88e3a0bb7516f1688fff"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_security\_group, f5xc\_security\_group\] Enter pre created security groups for
slo(Site Local Outside) and sli(Site Local Inside) interface. Supported only for sites deployed on
existing VPC.

Upstream description:

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

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

- [custom_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-59152abbfa2de683070f08747663621c3f154c22740e88e3a0bb7516f1688fff)
- [f5xc_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-d60c801bbd7a4ffe2c7b01eafc17214e0bf8bc4c935d5ae36e2fb9d0ab2a3c8b)

Select alternatives according to the provider validators above.

<a id="canonical-3461494c32d01b81345f7a0eea7e81395da0eb7e3877e920cd76f9cb2bcf8781"></a>

## Direct properties — custom_security_group / 4630f3b2ba8a / 3

<a id="canonical-8f7b6827edc8b41b7f4312ff6d8a93f73c92393387a5a3a2600fb6dfe2f0c935"></a>

<a id="canonical-04441418a11bf291de839bf41a228540da7589a0c571891ea0d6021f9fb13be7"></a>

## inside_security_group_id property — custom_security_group / 4630f3b2ba8a / 4

Type: `"string"`. Computed.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-37101993285dbd6d18accf978294df316fc9a4ef28c0a893a1b8bd4c41881b1f"></a>

<a id="canonical-0af93a530e52b8a7553dedef5bd9f0a41ddc9791f9b5f7f488f7351406cd3011"></a>

## outside_security_group_id property — custom_security_group / 4630f3b2ba8a / 5

Type: `"string"`. Computed.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-5c7932b053c4769f470bbd44e45a25e5f4e7d10e482df23a8fd1fe656845d4a4"></a>

## Next pages — custom_security_group / 4630f3b2ba8a / 6

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-a07d6b9e92df65b11b26bcede2ff5ea151145741fb6404881ba827c9c45df362"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d735b00b58105708f781728d4abb1be019007c2446a1cfe6cfe16b10e972d16a"></a>

## default_blocked_services — default_blocked_services / 1859c45059cf / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- default_blocked_services

<a id="canonical-16f59aaf40db507d24180fa9a4990a0a0b3bee4380e31b5e5a4d2a3449015f84"></a>

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

<a id="canonical-71819259a06d37a6089dfbb7968622ba7559bc23b2e0725b49bdd954f42ec1f5"></a>

## Direct properties — default_blocked_services / 1859c45059cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-52694a8524ba2ff6fec105d12bda9b95d242d81bb17620fe53808bb1b2b4f235"></a>

## Next pages — default_blocked_services / 1859c45059cf / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-99c0899d6bf0f8198853a25a8ad9829c746898d5f8959e6a601b46e61cfcbcf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-018131aa2a755d603bc0f8317bf5c5f77c50060f9939a6193eb3e8447049c8b2"></a>

## direct_connect_disabled — direct_connect_disabled / 94b6c999d1d2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- direct_connect_disabled

<a id="canonical-2c617222c25217dc0b439d73cd7069474b4675af810537f936c2dee162515414"></a>

Type: `["object", {}]`. Computed.

\[OneOf: direct\_connect\_disabled, direct\_connect\_enabled, private\_connectivity\] Enable this
option

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

- [direct_connect_disabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-2c617222c25217dc0b439d73cd7069474b4675af810537f936c2dee162515414)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-43a92841faa32b6ecee06bc534a421c302ba7c38832fd15815af04b5652ba3a1)
- [private_connectivity](data-sources--aws_vpc_site--reference--group-004.md#canonical-24ba5ea35ad5164775cda02456f1802f5c22784c2358d37c6fa4df6862c43fde)

Select alternatives according to the provider validators above.

<a id="canonical-c00acd4a985e4c0c53436a8210c086990c561dc56b109961ba00950abf4bd455"></a>

## Direct properties — direct_connect_disabled / 94b6c999d1d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c34faff0837799f6aaa8336e18a5af607014c44d6ddc172042a28897a2377168"></a>

## Next pages — direct_connect_disabled / 94b6c999d1d2 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e99baa0c2234585ce38637fde2fe1f6c099066c0f6165eb6c68d2976d88746d"></a>

## direct_connect_enabled — direct_connect_enabled / 7bb0dd24ffe7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- direct_connect_enabled

<a id="canonical-43a92841faa32b6ecee06bc534a421c302ba7c38832fd15815af04b5652ba3a1"></a>

Type: `"single"`. Computed.

Direct Connect Configuration. Direct Connect Configuration.

Upstream description:

Direct Connect Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-vif_choice": "[\"hosted_vifs\",\"standard_vifs\"]"
}
```

<a id="canonical-8466bb9a2c1e7e725439a442bd3344dd43a40583bab2f171e1e5f8f10d123865"></a>

## Direct properties — direct_connect_enabled / 7bb0dd24ffe7 / 3

- [auto_asn](data-sources--aws_vpc_site--reference--group-002.md#canonical-7a5536d28fed034d67630f3e262d707208dc5fc8b237255de8c6edfc3721c3cd): complete subsection reference.

<a id="canonical-87700e9942b8e0681a414f5166aee5acc77e8dfa8e3bd6617a75aefcd9f531e4"></a>

<a id="canonical-5f6fe093dc3babca9cd8fa1c4f4848d6b1be73eba75dd1f9e3ec2edfa1f8dcb9"></a>

## custom_asn property — direct_connect_enabled / 7bb0dd24ffe7 / 4

Type: `"number"`. Computed.

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Upstream description:

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248): complete subsection reference.

- [standard_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-81c2724e360c4d873b7bbee3c6a5aca224375b2030b6114dcc86b02b776e641b): complete subsection reference.

<a id="canonical-2054b1b29114ce013fe7c893c177c0317245678e6619dcc92bb0d39ab7ae20bf"></a>

## Next pages — direct_connect_enabled / 7bb0dd24ffe7 / 5

- [direct_connect_enabled.auto_asn](data-sources--aws_vpc_site--reference--group-002.md#canonical-7a5536d28fed034d67630f3e262d707208dc5fc8b237255de8c6edfc3721c3cd)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248)
- [direct_connect_enabled.standard_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-81c2724e360c4d873b7bbee3c6a5aca224375b2030b6114dcc86b02b776e641b)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-7a5536d28fed034d67630f3e262d707208dc5fc8b237255de8c6edfc3721c3cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b68c59909cba6590aed9a217a569d123367d082be69c5a523d7d6ec8a64d6a7"></a>

## direct_connect_enabled.auto_asn — direct_connect_enabled.auto_asn / 27838ee23de3 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- direct_connect_enabled.auto_asn

<a id="canonical-ff95c5d513e9ca3e0fb5766ea15926d414fe78ab80495923e5ccd79f3356a6fd"></a>

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

<a id="canonical-ee784735cc0f62153762a53244a59292840c4c08fba9b8bddf04b99dc0ca5b21"></a>

## Direct properties — direct_connect_enabled.auto_asn / 27838ee23de3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b3e6c2df62ab4bc6be295badb286fe80ecf872b843a8d311c98026d69a146b6"></a>

## Next pages — direct_connect_enabled.auto_asn / 27838ee23de3 / 4

- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0b1dfc3eb2a8924b8b81126e4b51fb9ece32ede622fb4b2a7e7a42ab984fc88"></a>

## direct_connect_enabled.hosted_vifs — direct_connect_enabled.hosted_vifs / b7cba04c3e0d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- direct_connect_enabled.hosted_vifs

<a id="canonical-3b1467446beb1aafda887fbcb3b9b7db8696e1c626eda9d6b13bc9545676b141"></a>

Type: `"single"`. Computed.

AWS Direct Connect Hosted VIF Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

<a id="canonical-f194b565d5b11646c4ab02632b74af2b4004b799cafb175717f13e7dfd2de16a"></a>

## Direct properties — direct_connect_enabled.hosted_vifs / b7cba04c3e0d / 3

- [site_registration_over_direct_connect](data-sources--aws_vpc_site--reference--group-002.md#canonical-ea71765fc4ae52bc16f0334417154ff945c5c1540c4050cc8fe54694255e2513): complete subsection reference.

- [site_registration_over_internet](data-sources--aws_vpc_site--reference--group-002.md#canonical-9391d282e248c71138f2891458f7723edb898fced8ec13b9e4047866cb6b3f3a): complete subsection reference.

- [vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-fbc7588ddb4d7129e835f80c39dec09c137c5c842ef489dfbf24ff9f93d84812): complete subsection reference.

<a id="canonical-91cc7ae11b53a386e42cdd439367efa772897652521c973a25033dcb3c9a1591"></a>

## Next pages — direct_connect_enabled.hosted_vifs / b7cba04c3e0d / 4

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_vpc_site--reference--group-002.md#canonical-ea71765fc4ae52bc16f0334417154ff945c5c1540c4050cc8fe54694255e2513)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_vpc_site--reference--group-002.md#canonical-9391d282e248c71138f2891458f7723edb898fced8ec13b9e4047866cb6b3f3a)
- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-fbc7588ddb4d7129e835f80c39dec09c137c5c842ef489dfbf24ff9f93d84812)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-ea71765fc4ae52bc16f0334417154ff945c5c1540c4050cc8fe54694255e2513"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aae220b68d26dbad1353c472c2c05bdab9582cf37258caabad8d15627958d30f"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / 93b42b85adfa / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-b88c0dd12095ea5de077c90ba05052b633b08316d6af82dea8776bb13601acc1"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

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

<a id="canonical-916b88a0bcd103b625d704d152cd2dcfc507f14dd95735b629623e481801a032"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / 93b42b85adfa / 3

<a id="canonical-2c95dac838294e91c99cc067b2fc25896b7fd94e58de7af3ca34ac757bcd069d"></a>

<a id="canonical-64e184ee83c89d061d190947cd17dc7512ad77378071be69550dd743ec589671"></a>

## cloudlink_network_name property — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / 93b42b85adfa / 4

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-55b9cb3203f00adc07e4924aa2aba5b4158f0fda01bc3c97a3e1d5c4b18e74d4"></a>

## Next pages — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / 93b42b85adfa / 5

- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-9391d282e248c71138f2891458f7723edb898fced8ec13b9e4047866cb6b3f3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27d4a0b4ce2fb8ee72f9d8f1fe574749236cc6eb32273564fa46c468c23b3daf"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_internet — direct_connect_enabled.hosted_vifs.site_registration_over_internet / 7075aa99c735 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-a67fd1b69bd6ad2b80d539705d1314f70e8af744ba51152e949df6d87dc0cc1f"></a>

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

<a id="canonical-b5020ad68aecbb5a3b5f1cbe2509ed871b705db815fb08d0cd57a65dc9c737a2"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.site_registration_over_internet / 7075aa99c735 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab2ecb3ada8cb33491957b4627e8c264ebba502b1c5dee31dad82357c5fa4b8c"></a>

## Next pages — direct_connect_enabled.hosted_vifs.site_registration_over_internet / 7075aa99c735 / 4

- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-fbc7588ddb4d7129e835f80c39dec09c137c5c842ef489dfbf24ff9f93d84812"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40231d50732d3d025697abebebcfc05edbd213625d96a68bb8ac95307978ace7"></a>

## direct_connect_enabled.hosted_vifs.vif_list — direct_connect_enabled.hosted_vifs.vif_list / 6d979f114a10 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-81791c20f393c64c06743821d0728623b0a6b48a5c96fa59bbbe1891d94a3e60"></a>

Type: `"list"`. Computed.

List of Hosted VIF Config. List of Hosted VIF Config.

Upstream description:

List of Hosted VIF Config.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 30,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 30,
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
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3f3adf8fd58754a5cc82606df4f34599e386e01c377a797d0e6aa94dabdeac78"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.vif_list / 6d979f114a10 / 3

<a id="canonical-d2e104ad110ded5dcf30906543c60724c8bf46fa0f2371a4aa354b52e090ce48"></a>

<a id="canonical-3141291e29a60b2ba87973bca722d34cf0e66e26007c4ffac143247f3c113033"></a>

## other_region property — direct_connect_enabled.hosted_vifs.vif_list / 6d979f114a10 / 4

Type: `"string"`. Computed.

\[Enum:
af-south-1|ap-east-1|ap-northeast-1|ap-northeast-2|ap-south-1|ap-southeast-1|ap-southeast-2|ap-southeast-3|ca-central-1|eu-central-1|eu-north-1|eu-south-1|eu-west-1|eu-west-2|eu-west-3|me-south-1|sa-east-1|us-east-1|us-east-2|us-west-1|us-west-2\]
Exclusive with \[same\_as\_site\_region\] Other Region. Possible values are \`af-south-1\`,
\`ap-east-1\`, \`ap-northeast-1\`, \`ap-northeast-2\`, \`ap-south-1\`, \`ap-southeast-1\`,
\`ap-southeast-2\`, \`ap-southeast-3\`, \`ca-central-1\`, \`eu-central-1\`, \`eu-north-1\`,
\`eu-south-1\`, \`eu-west-1\`, \`eu-west-2\`, \`eu-west-3\`, \`me-south-1\`, \`sa-east-1\`,
\`us-east-1\`, \`us-east-2\`, \`us-west-1\`, \`us-west-2\`.

Upstream description:

Exclusive with \[same\_as\_site\_region\] Other Region.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  }
}
```

- [same_as_site_region](data-sources--aws_vpc_site--reference--group-002.md#canonical-de8d2012e95aa5c78940091d214a80c0b82674bf17c7310e2f71cabd0f53e4a1): complete subsection reference.

<a id="canonical-1d3c3434295ac0408805dda5cf71279e7ecfa4f525331620fedef430b6f5c417"></a>

<a id="canonical-73a64b257129360b7a4ee28171d8f542c7768cf1be259d664a6a08709fb6e42b"></a>

## vif_id property — direct_connect_enabled.hosted_vifs.vif_list / 6d979f114a10 / 5

Type: `"string"`. Computed.

AWS Direct Connect VIF ID that needs to be connected to the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-2ae1cfe812d4ca2686b9f5f1fbaa02e623ee19fd7d65950741003d45eab60308"></a>

## Next pages — direct_connect_enabled.hosted_vifs.vif_list / 6d979f114a10 / 6

- [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](data-sources--aws_vpc_site--reference--group-002.md#canonical-de8d2012e95aa5c78940091d214a80c0b82674bf17c7310e2f71cabd0f53e4a1)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-de8d2012e95aa5c78940091d214a80c0b82674bf17c7310e2f71cabd0f53e4a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f38f16fe07bb8c74ec2ffa49b8ad55a40ff224cacfd502a657399c4fd124c75a"></a>

## direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / 753cee7241c1 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-72adbcd0b407d7e552c20494cfb6688be7ffc711043178cd61b6ebad1f65e248)
- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-fbc7588ddb4d7129e835f80c39dec09c137c5c842ef489dfbf24ff9f93d84812)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-f6f548d5eae93de74aa7230aec1453293535a0b11b9d22cad1b300fed13d6845"></a>

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

<a id="canonical-aa960067f30c3ef14e7c2e7f5f0e8b1167fb660b2aef5d693e929c480fd559a0"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / 753cee7241c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7524dfce299e50b9fa83db8d12b2b138210a4961293b1068ffa3f9e9ebecedfe"></a>

## Next pages — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / 753cee7241c1 / 4

- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-fbc7588ddb4d7129e835f80c39dec09c137c5c842ef489dfbf24ff9f93d84812)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-81c2724e360c4d873b7bbee3c6a5aca224375b2030b6114dcc86b02b776e641b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd9cbbccbc9aea27a1231d02851609eb11e701a67038f036ad83f46de2bfe6e8"></a>

## direct_connect_enabled.standard_vifs — direct_connect_enabled.standard_vifs / f27a3e39b932 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- direct_connect_enabled.standard_vifs

<a id="canonical-116d3f8e44034f7159af45880438dff56f4804ce0730e96df3cc8d27a00d3505"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for standard vifs.

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

<a id="canonical-22a09ecf9278b8494ac3488a261b0921250e5c335e23f58fef93b53640542c52"></a>

## Direct properties — direct_connect_enabled.standard_vifs / f27a3e39b932 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-926ed8148f89d0abdeaba2b772e6ae38980cca0040c995783f25b09a26e95c6c"></a>

## Next pages — direct_connect_enabled.standard_vifs / f27a3e39b932 / 4

- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-7516c1587d0df53ad876cae0477b2c38385dec4dd103145bda5084516dbca326"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d06f05d0b598562b4534300bc237001d739b37eeb1b231c6c5898a7bfcbfbab"></a>

## disable_encryption — disable_encryption / 2a8a89581671 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- disable_encryption

<a id="canonical-fffbfd32e36933134466c77289daef811a43e2eb3261a682831aa3e2b1651c73"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_encryption, enable\_encryption; Default: disable\_encryption\] Configuration
parameter for disable encryption.

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

- [disable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-fffbfd32e36933134466c77289daef811a43e2eb3261a682831aa3e2b1651c73)
- [enable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-f12821ee886d6be62e13aefc0b88f56f687993fdb3c8bed20d616976e32424de)

Select alternatives according to the provider validators above.

<a id="canonical-2282d6d6a787dcd76ddb9acd5ac6ba907cee784265d66dca16219c06a1ffb61a"></a>

## Direct properties — disable_encryption / 2a8a89581671 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d7aba0ffd45424f585544aec0e81b08283e116af6a9842add8562b8e6de11e86"></a>

## Next pages — disable_encryption / 2a8a89581671 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-1fa273c1c1a47fed8481715b68dfeeea56376ac2fd09470b7846e01af544bae7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aaee4f6ff6f4b4b7bbc53893c8fc6864d4246ef98fef90d937dd90b20c4f9a0f"></a>

## disable_internet_vip — disable_internet_vip / 882f3b77add2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- disable_internet_vip

<a id="canonical-0933b0e5d9c81b8e2cb92c5d815ae1885ec94bf6d2396c60dde2742e087b6959"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_internet\_vip, enable\_internet\_vip; Default: disable\_internet\_vip\] Enable
this option

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

- [disable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-0933b0e5d9c81b8e2cb92c5d815ae1885ec94bf6d2396c60dde2742e087b6959)
- [enable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-e9ec1ae00695b0d7b92a7a826fe89a1ad5b50947d7f25fd0a123ff27f6403d60)

Select alternatives according to the provider validators above.

<a id="canonical-1fa8cfa93c6f8dcd0fefc66589761f794cda00eba44fbfacc3cf91fa6de08d95"></a>

## Direct properties — disable_internet_vip / 882f3b77add2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb6ca187ff2f790122b60aa2d112e710f1acc33987451b9bd62f84fe6404398a"></a>

## Next pages — disable_internet_vip / 882f3b77add2 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-27eda6087dad60ebfab2bd9d183d63d4fd3d5ff5c8054c9ab718e31ae3b124e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e6150df04b896c00c5b317078169e8b3f9a31e398cd6a83c986ea7fe922cca1"></a>

## egress_gateway_default — egress_gateway_default / fd923b60c583 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- egress_gateway_default

<a id="canonical-04ae9d37f36abc850a51c14aa5065a12dc1299d51fe56cb03192ea4be5d6aa03"></a>

Type: `["object", {}]`. Computed.

\[OneOf: egress\_gateway\_default, egress\_nat\_gw, egress\_virtual\_private\_gateway; Default:
egress\_gateway\_default\] Configuration parameter for egress gateway default.

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

- [egress_gateway_default](data-sources--aws_vpc_site--reference--group-002.md#canonical-04ae9d37f36abc850a51c14aa5065a12dc1299d51fe56cb03192ea4be5d6aa03)
- [egress_nat_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-f9b981f1221b3bc14a00e436edcdf6531a5a0bbd875174e5cb016c74e7d2b225)
- [egress_virtual_private_gateway](data-sources--aws_vpc_site--reference--group-002.md#canonical-7c275335231a1534b98345464271e54a30ed72381f8e15caeafbd207f57dc8ba)

Select alternatives according to the provider validators above.

<a id="canonical-d352c544935bd8c77cb4f36200c58f81d875141d0c6bdc7a9952209e94351531"></a>

## Direct properties — egress_gateway_default / fd923b60c583 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-24bbd3a53f3c86549f33014aed72abff558d4c62e532c03f517ff1b4967707f0"></a>

## Next pages — egress_gateway_default / fd923b60c583 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-9e4ad9435e0b0e6716c7f67d167fce5206201d39fc76b01abb1877385ec5606b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0a233c33632ac225827c504c0c83022e94b43301abc1bbae2920504eedbfc4a"></a>

## egress_nat_gw — egress_nat_gw / 511e6a0c8301 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- egress_nat_gw

<a id="canonical-f9b981f1221b3bc14a00e436edcdf6531a5a0bbd875174e5cb016c74e7d2b225"></a>

Type: `"single"`. Computed.

With this option, egress site traffic will be routed through an Network Address Translation(NAT)
Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"nat_gw_id\"]"
}
```

<a id="canonical-d95075654b5f4cd305e0158e7b04ad0f33573b6e45dba4dc751635ba4d6ea216"></a>

## Direct properties — egress_nat_gw / 511e6a0c8301 / 3

<a id="canonical-b8ade35052c84311900517fa685e3ebefef41ba0f97387416eeee83e0bd65f40"></a>

<a id="canonical-44eefb56fb644191c703574310d701f6a4f2e61bd6b3ca2846924b2862113272"></a>

## nat_gw_id property — egress_nat_gw / 511e6a0c8301 / 4

Type: `"string"`. Computed.

Existing NAT Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 21,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-b5596c4e936c496ce0daf49a38f0bb1378931df6a1a740b8b8b717e3539f038c"></a>

## Next pages — egress_nat_gw / 511e6a0c8301 / 5

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-408490482dd6c239b209790d7fbc93eab5b8299b6bc34c608b552cf8779a4ac7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc3d93bd4ee3744574a1e0a78b8973103c999c63db6f0a40dbe294834d370e59"></a>

## egress_virtual_private_gateway — egress_virtual_private_gateway / 3728208ba412 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- egress_virtual_private_gateway

<a id="canonical-7c275335231a1534b98345464271e54a30ed72381f8e15caeafbd207f57dc8ba"></a>

Type: `"single"`. Computed.

With this option, egress site traffic will be routed through an Virtual Private Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"vgw_id\"]"
}
```

<a id="canonical-8119e61d75b9ea5d11bbbd3d6aabc0d5b6bdaeaa99e7e0c1023952140361a5ff"></a>

## Direct properties — egress_virtual_private_gateway / 3728208ba412 / 3

<a id="canonical-51625879836dd753cde6495349ce898e3ff0c2d349fea5cf14cde078a4b9bafa"></a>

<a id="canonical-32f6affc8e983bb1373bc28dd466dde63012da3f878325df19baee7be3c04d46"></a>

## vgw_id property — egress_virtual_private_gateway / 3728208ba412 / 4

Type: `"string"`. Computed.

Existing Virtual Private Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 21,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-5b6bf905e269d541728a3603ee789a13feae4c31afd81cdd3431df6e553e7a98"></a>

## Next pages — egress_virtual_private_gateway / 3728208ba412 / 5

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-ebce9f17e54eded7f5f4b5f98bed9bd8bc031751d987e072803886c4bfa51dfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-317bc3bab1a1fd9b95b50f167ae0d31d08aa086308277eb976930b352cd7497a"></a>

## enable_encryption — enable_encryption / fa2d2ee8f263 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- enable_encryption

<a id="canonical-f12821ee886d6be62e13aefc0b88f56f687993fdb3c8bed20d616976e32424de"></a>

Type: `"single"`. Computed.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

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

<a id="canonical-ce81bc09c7c6209b5754cdaa66ae78d263dfe3eac8779e86bc529a37498be11a"></a>

## Direct properties — enable_encryption / fa2d2ee8f263 / 3

<a id="canonical-6eb0bc33796fdbf1639a3313764a8d10659acee07a0f48f34702fd6d7db4910d"></a>

<a id="canonical-bd23f33e8691a59fa3d66a1e023b2b59be8468176de457ad16ba1025302a3542"></a>

## kms_key_id property — enable_encryption / fa2d2ee8f263 / 4

Type: `"string"`. Computed.

AWS KMS Key to be used to encrypt the disk attached to the VM.

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

<a id="canonical-3dce233a2c795e8c7096c3f16446e39b7ccd7177f27e6004c1337f725f854962"></a>

## Next pages — enable_encryption / fa2d2ee8f263 / 5

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-b2e9032440c417e11bfaa49c40aaf20079c79a6264edb2d63f11954983deeeca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5098664a4238170e52bae5efc3028397351142885b031727e31bbae94b5fd7e"></a>

## enable_internet_vip — enable_internet_vip / 382132ac5995 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- enable_internet_vip

<a id="canonical-e9ec1ae00695b0d7b92a7a826fe89a1ad5b50947d7f25fd0a123ff27f6403d60"></a>

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

<a id="canonical-5541d9938d3fbbcf209723ee71a97700e22dbfd313c6108abebfb2f210de672d"></a>

## Direct properties — enable_internet_vip / 382132ac5995 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1a1f4c03800be9a772fe4231c15cc399ee748ddd147331a5a820af0545c8f20"></a>

## Next pages — enable_internet_vip / 382132ac5995 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-cc60cce8b152be015a54c11c552724934297d0d6a25ba937a96666f5a1e6ba16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-421711bdae0f28625d3a3aa6d3733027753fb69fabede727a54eba9c15f5b2bb"></a>

## f5_orchestrated_routing — f5_orchestrated_routing / ca5560c31b6c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- f5_orchestrated_routing

<a id="canonical-727e6614fe7e42957a96f7e90ec05f73e047b974ef839990292469dc87ac9f85"></a>

Type: `["object", {}]`. Computed.

\[OneOf: f5\_orchestrated\_routing, manual\_routing\] Enable this option

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

- [f5_orchestrated_routing](data-sources--aws_vpc_site--reference--group-002.md#canonical-727e6614fe7e42957a96f7e90ec05f73e047b974ef839990292469dc87ac9f85)
- [manual_routing](data-sources--aws_vpc_site--reference--group-004.md#canonical-006673fa3223785abfa75471531e49830b0bab9c65f232e240dc7fd6200a09ee)

Select alternatives according to the provider validators above.

<a id="canonical-0725bb9a95f5dff35b27b5e36070cead20757186f6da0e48119a94f9e4f4ecf6"></a>

## Direct properties — f5_orchestrated_routing / ca5560c31b6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a22cad1aa34547611d6ff91532eac7dd1d0d8bc5395669bb38f5033fe6cab72"></a>

## Next pages — f5_orchestrated_routing / ca5560c31b6c / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-e52d2230b230a613685d90c39f2ec99ce98ccd1da9751e604aed0656f9695db2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e555aa7cba02d85e324a533632452dfbc1d04001467e88de60183420c4da85b"></a>

## f5xc_security_group — f5xc_security_group / dbe463e76aee / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- f5xc_security_group

<a id="canonical-d60c801bbd7a4ffe2c7b01eafc17214e0bf8bc4c935d5ae36e2fb9d0ab2a3c8b"></a>

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

<a id="canonical-f81860c4a9070619b8502b6d2f991925a9ad0689f336d210aa702ff0980bd83f"></a>

## Direct properties — f5xc_security_group / dbe463e76aee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dbc858c474763f20bb28f6775145794f2c98531e98911c55927c55722bb435fb"></a>

## Next pages — f5xc_security_group / dbe463e76aee / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ed2351ff653f147dbbe23283d6c619c2f8d406071ee2d3363cb517a648a1ac2"></a>

## ingress_egress_gw — ingress_egress_gw / d90f00ba0ab8 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- ingress_egress_gw

<a id="canonical-8558d68e84e4446e22acc7dc87c078b052dee2c7a631eb6fee6fc62d3325f80f"></a>

Type: `"single"`. Computed.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface AWS ingress/egress site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-8558d68e84e4446e22acc7dc87c078b052dee2c7a631eb6fee6fc62d3325f80f)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-2587d0e59e93e14e4342d3ebfcc4f857e71f684d5db3c70a77657e26d368a922)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-86b971c090b17b2b68826ac670d5250c7ea2fcc25680b525a5e4fba7db9673a0)

Select alternatives according to the provider validators above.

<a id="canonical-61c31d3500c0a4d767290989986ae1df3b41344c8238c29737e003f8d29b2107"></a>

## Direct properties — ingress_egress_gw / d90f00ba0ab8 / 3

- [active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-3e08df3441d57486a7cab3bed3eb6b091c21803e33c6432d10de88a11984842d): complete subsection reference.

- [active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-b995345905b1be88325d5afa0bb72c43fe9c4dc8b599da19955f8bc77b341a78): complete subsection reference.

- [active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-5c4c1f0818cb59da7114a06c5992f2250e1baaf7c90df8e395144e8357e97104): complete subsection reference.

- [allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e): complete subsection reference.

- [allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331): complete subsection reference.

<a id="canonical-e2a0e257d95d06fdb4e2629f5a8d7dffac5741f4b00f7b656595406b912df3c9"></a>

<a id="canonical-d52a857fd3d677c5adc1ddc5ce86ee8c81f75d552d16ea2b464c958fe30fc98c"></a>

## aws_certified_hw property — ingress_egress_gw / d90f00ba0ab8 / 4

Type: `"string"`. Computed.

\[Enum: aws-byol-multi-nic-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The
only possible value is \`aws-byol-multi-nic-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26): complete subsection reference.

- [dc_cluster_group_inside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-2a6852e810022ac31544b2447fb302bd35b1ae58b86a042440ae79d0fbd8dfbc): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-08b710f8cb19ae0b29bb024ace9c0c0ca77b40114ff5e92280fcc72bc1ed3078): complete subsection reference.

- [forward_proxy_allow_all](data-sources--aws_vpc_site--reference--group-002.md#canonical-42e3e0c67afb1f7e797175ec0767137bb70771de9c4fe9414268ea5d21bfff49): complete subsection reference.

- [global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279): complete subsection reference.

- [inside_static_routes](data-sources--aws_vpc_site--reference--group-002.md#canonical-1cd789f89f4736315a9dbca490acb46cd71979de8e8aad79b5856820c301d900): complete subsection reference.

- [no_dc_cluster_group](data-sources--aws_vpc_site--reference--group-003.md#canonical-97e60bedc56b9e67bae6c8a43b0d6840269aa474d34c73a4d5544f92a2c323da): complete subsection reference.

- [no_forward_proxy](data-sources--aws_vpc_site--reference--group-003.md#canonical-b61834664c9f8d671bb1aba2740b3530b5485154360698be1a2593686c64a113): complete subsection reference.

- [no_global_network](data-sources--aws_vpc_site--reference--group-003.md#canonical-c6ea3d4b5a26f8fc1de2a108458f449f5fd5cfc81d36e1ec9493041d9d1f8eaf): complete subsection reference.

- [no_inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-5afb431ab18fde7df627ddce14dc9a25c09abd10d36f15c3a1073241f45d3aac): complete subsection reference.

- [no_network_policy](data-sources--aws_vpc_site--reference--group-003.md#canonical-ba1bdd54c791c03c30f6f919c1e1f93cccb0b311f0a4aee22fad10dcda911709): complete subsection reference.

- [no_outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1a33c22d48421feb8b5032bbaa4704a5e77ec5ec78ba696a5c681ee414504178): complete subsection reference.

- [outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-5a02ebc39a14d11b09b5cfc0f75cf7e6b918fd72dd27aebccf83568053a3e7b6): complete subsection reference.

- [performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-ff5618d86b69ff7d9dd839203306e1999ab34d2d04061ae2f77e854cce12102c): complete subsection reference.

- [sm_connection_public_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-4ec8fc999958d21c9812de10672703581a0b9766b5624937d317c54f866fbf44): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-4d40f80d834113ad440ad36b53b3831f140fae60c3ae0cf54b271b716e545a29): complete subsection reference.

<a id="canonical-8d246186e9b59531f38d6e4f14aadf147a3b91f262e88df735d1e6e5ec510a5a"></a>

## Next pages — ingress_egress_gw / d90f00ba0ab8 / 5

- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-3e08df3441d57486a7cab3bed3eb6b091c21803e33c6432d10de88a11984842d)
- [ingress_egress_gw.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-b995345905b1be88325d5afa0bb72c43fe9c4dc8b599da19955f8bc77b341a78)
- [ingress_egress_gw.active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-5c4c1f0818cb59da7114a06c5992f2250e1baaf7c90df8e395144e8357e97104)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-2a6852e810022ac31544b2447fb302bd35b1ae58b86a042440ae79d0fbd8dfbc)
- [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-08b710f8cb19ae0b29bb024ace9c0c0ca77b40114ff5e92280fcc72bc1ed3078)
- [ingress_egress_gw.forward_proxy_allow_all](data-sources--aws_vpc_site--reference--group-002.md#canonical-42e3e0c67afb1f7e797175ec0767137bb70771de9c4fe9414268ea5d21bfff49)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-002.md#canonical-1cd789f89f4736315a9dbca490acb46cd71979de8e8aad79b5856820c301d900)
- [ingress_egress_gw.no_dc_cluster_group](data-sources--aws_vpc_site--reference--group-003.md#canonical-97e60bedc56b9e67bae6c8a43b0d6840269aa474d34c73a4d5544f92a2c323da)
- [ingress_egress_gw.no_forward_proxy](data-sources--aws_vpc_site--reference--group-003.md#canonical-b61834664c9f8d671bb1aba2740b3530b5485154360698be1a2593686c64a113)
- [ingress_egress_gw.no_global_network](data-sources--aws_vpc_site--reference--group-003.md#canonical-c6ea3d4b5a26f8fc1de2a108458f449f5fd5cfc81d36e1ec9493041d9d1f8eaf)
- [ingress_egress_gw.no_inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-5afb431ab18fde7df627ddce14dc9a25c09abd10d36f15c3a1073241f45d3aac)
- [ingress_egress_gw.no_network_policy](data-sources--aws_vpc_site--reference--group-003.md#canonical-ba1bdd54c791c03c30f6f919c1e1f93cccb0b311f0a4aee22fad10dcda911709)
- [ingress_egress_gw.no_outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1a33c22d48421feb8b5032bbaa4704a5e77ec5ec78ba696a5c681ee414504178)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-5a02ebc39a14d11b09b5cfc0f75cf7e6b918fd72dd27aebccf83568053a3e7b6)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-ff5618d86b69ff7d9dd839203306e1999ab34d2d04061ae2f77e854cce12102c)
- [ingress_egress_gw.sm_connection_public_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-4ec8fc999958d21c9812de10672703581a0b9766b5624937d317c54f866fbf44)
- [ingress_egress_gw.sm_connection_pvt_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-4d40f80d834113ad440ad36b53b3831f140fae60c3ae0cf54b271b716e545a29)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-3e08df3441d57486a7cab3bed3eb6b091c21803e33c6432d10de88a11984842d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-defa7c6456ee7b6d3a176b1014c9e6c17d00ae4ca6234d541b2f3f6ffa08b058"></a>

## ingress_egress_gw.active_enhanced_firewall_policies — ingress_egress_gw.active_enhanced_firewall_policies / 4dc00150d748 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.active_enhanced_firewall_policies

<a id="canonical-f4a1e550939e4a7e68aa4d4525621da7f58e6895becc49ca3a9c77b9d1659709"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

<a id="canonical-00f86ca2d77f9d4325cdf1818011f954683e31acb90d6102aa8b61315d178083"></a>

## Direct properties — ingress_egress_gw.active_enhanced_firewall_policies / 4dc00150d748 / 3

- [enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-f0b023ead38104155dba30ebee107cea6a4c3763239be307f0f746f8b9a93f98): complete subsection reference.

<a id="canonical-f9f34c28ffa12fc580fc35371e546c53bc0f23e4803ef3b753d68a2837d974c5"></a>

## Next pages — ingress_egress_gw.active_enhanced_firewall_policies / 4dc00150d748 / 4

- [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-f0b023ead38104155dba30ebee107cea6a4c3763239be307f0f746f8b9a93f98)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-f0b023ead38104155dba30ebee107cea6a4c3763239be307f0f746f8b9a93f98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad83c1767ac987774e80eb343ea0602caee17f441678cea7527ebf8785146904"></a>

## ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 95b26009de4f / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-3e08df3441d57486a7cab3bed3eb6b091c21803e33c6432d10de88a11984842d)
- ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-68b7057b9a6bb79c33ae9bb4c24ff914eba5c83ff2e2e840837ddb6e3988020e"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-7ec902c40add5327d08ea96127f64c555b55de04a610f08a346eced65c65c9db"></a>

## Direct properties — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 95b26009de4f / 3

<a id="canonical-6b4f2cb8f51347c8934aedc1eab6d17ad733b646888657f0f093103cc1af7dc8"></a>

<a id="canonical-4a0597df4fdc493efb00dceb0a762c0a284d0735792c240835930171994a4097"></a>

## name property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 95b26009de4f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-a46575a67db96ad25431e57272a00f44da963d63a037a66112ad01bfaf716c6d"></a>

<a id="canonical-9cbde938c659e6759127f30ed319a1e437257f2cf92fe5dcd26c0e6462c87e56"></a>

## namespace property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 95b26009de4f / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-87869e5a70428e895af6b0717689464bb7c86f1cbd5c383deba33d15dd0fa158"></a>

<a id="canonical-28161e8e877c8a5fab884e0c9d207f034a3acccfc19f96f633e18169ac6a21e7"></a>

## tenant property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 95b26009de4f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-8ae5d1f96d5edce5c642c3930a4eee63198b88fc81b34d4f93a84e6864ad5b54"></a>

## Next pages — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 95b26009de4f / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-3e08df3441d57486a7cab3bed3eb6b091c21803e33c6432d10de88a11984842d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-b995345905b1be88325d5afa0bb72c43fe9c4dc8b599da19955f8bc77b341a78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eeeb2a37fd2e47dc0ba79b138eafab88dc1a534523ca549057b25f5fc7759e4c"></a>

## ingress_egress_gw.active_forward_proxy_policies — ingress_egress_gw.active_forward_proxy_policies / 8021b4baaf9d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.active_forward_proxy_policies

<a id="canonical-7a327cb15ea1d59564ab2d30fcbcad843c0335543cdf05698e5bb6b023d34c0c"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

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

<a id="canonical-d2f7a34350d585c5fd630ef8a54a6d771c33cbceb8cb05544fd534a9fd542b45"></a>

## Direct properties — ingress_egress_gw.active_forward_proxy_policies / 8021b4baaf9d / 3

- [forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-307fccc279eab97d3d1f567cfc3a7a1f0f21182539b215dfb5d5702012123907): complete subsection reference.

<a id="canonical-df40ef25a8e069b39837a816b8536dbff1989354084b2bf34a97da2ff4a57ed7"></a>

## Next pages — ingress_egress_gw.active_forward_proxy_policies / 8021b4baaf9d / 4

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-307fccc279eab97d3d1f567cfc3a7a1f0f21182539b215dfb5d5702012123907)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-307fccc279eab97d3d1f567cfc3a7a1f0f21182539b215dfb5d5702012123907"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a276f1be3697648cd613ee21a42f060736eb89a8497ac24870e2ad731534e38d"></a>

## ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 456aa22f959f / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-b995345905b1be88325d5afa0bb72c43fe9c4dc8b599da19955f8bc77b341a78)
- ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-0a9b38ce7db69f72b267503f7caa9c87b0a1247961bab9a01ce1451aa9b6936e"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1f9ddca5a3e998db03cc2ecb48cad19f4032abce0af0a62da3791f84f4f235c8"></a>

## Direct properties — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 456aa22f959f / 3

<a id="canonical-4304c534b11e2749b35b7fde65e0bd2a41ad158ef3f61ebac4d2249f6cc4fd6f"></a>

<a id="canonical-b64b270a748118f3481b1a51e31d057a3e395b23d71e792d6b4d10316816090b"></a>

## name property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 456aa22f959f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-5a0469d922fde628b649c41cf3111a11693c169b3bec50056d5cb223124fa36d"></a>

<a id="canonical-52fd8d50a0c1926274eabc326b07dfcdae9ad3ceeeaa26135e3d70175c4c5042"></a>

## namespace property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 456aa22f959f / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-ad2afd32399442e87ed8997d76a820b8b325b21bd290493d2032ebe6e2eeb4bb"></a>

<a id="canonical-2676406753169d7fdf0c1b46565be494f73bf22a44f2d397f1b16cf251bc9c81"></a>

## tenant property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 456aa22f959f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-5236a0a26a1a296568b8a2b0cd593601467e1b48c6e710eff8461230d0ee270c"></a>

## Next pages — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 456aa22f959f / 7

- [ingress_egress_gw.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-b995345905b1be88325d5afa0bb72c43fe9c4dc8b599da19955f8bc77b341a78)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-5c4c1f0818cb59da7114a06c5992f2250e1baaf7c90df8e395144e8357e97104"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7426b5f625251e9e9b69607d26b8b26092f9a70ca6ba2b799f1ef4f0482d0fc5"></a>

## ingress_egress_gw.active_network_policies — ingress_egress_gw.active_network_policies / 9aebecc4dd4f / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.active_network_policies

<a id="canonical-4dfbaaf99a1a0c711279300c25f85e0f1cfb5d4bf72ce76703836d4a2cfa5c37"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

<a id="canonical-069415f680ef2f135e428977de06acc2c71e2b2d90ad0886d6f446147b455537"></a>

## Direct properties — ingress_egress_gw.active_network_policies / 9aebecc4dd4f / 3

- [network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-431e2132b30bf7ef1e8fb8f695606dc0c6da945a8bf1de250d448321c9fbe95f): complete subsection reference.

<a id="canonical-5ed1138d92842c27cb5f38ef8ae49acc3f58ed8dc5189a1cec2c7262bbea8eb3"></a>

## Next pages — ingress_egress_gw.active_network_policies / 9aebecc4dd4f / 4

- [ingress_egress_gw.active_network_policies.network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-431e2132b30bf7ef1e8fb8f695606dc0c6da945a8bf1de250d448321c9fbe95f)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-431e2132b30bf7ef1e8fb8f695606dc0c6da945a8bf1de250d448321c9fbe95f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06fbdfacb519b6fec6e6f37e053869c957b85614333b67966282cb6e43e9a31f"></a>

## ingress_egress_gw.active_network_policies.network_policies — ingress_egress_gw.active_network_policies.network_policies / 8b75b82d72be / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-5c4c1f0818cb59da7114a06c5992f2250e1baaf7c90df8e395144e8357e97104)
- ingress_egress_gw.active_network_policies.network_policies

<a id="canonical-d0e224b3ed476cbae5fb9e78b12bf64e9b3c006200c282e289bd97e85941caed"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2e6a795443fe622881cd21e17f6e1769a824c7057723e9d98cb0a0e8c8ab822f"></a>

## Direct properties — ingress_egress_gw.active_network_policies.network_policies / 8b75b82d72be / 3

<a id="canonical-4412e995e1f6725edaba613a077d38d71d63ae8e1f39ec7379cc485133a05b2d"></a>

<a id="canonical-48dec68f2b3bfbe730acd97557dfbbf150184d0777c8fc636992cc5ead6573ac"></a>

## name property — ingress_egress_gw.active_network_policies.network_policies / 8b75b82d72be / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-49db6808a6039be1e7a5297762ff1f4f37e5059a0ab845b3a89d3506f1e7244d"></a>

<a id="canonical-4f463b948321733b12fffbf32c8dc5159faa2cd53d55f1a5db39641d7ce8eb62"></a>

## namespace property — ingress_egress_gw.active_network_policies.network_policies / 8b75b82d72be / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-5cf83fa1435bb3be4bcac570c3d86bf52397d0ebb8b7830e06113ed6f3c16168"></a>

<a id="canonical-02763f858fe6312d3a76b5660dac4a21ab3163f78ab11bb580252edf45be4e18"></a>

## tenant property — ingress_egress_gw.active_network_policies.network_policies / 8b75b82d72be / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1a27c1f38067b7ab134b28101a4a9ae83b070bcbc4dca30409126c144078f3f5"></a>

## Next pages — ingress_egress_gw.active_network_policies.network_policies / 8b75b82d72be / 7

- [ingress_egress_gw.active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-5c4c1f0818cb59da7114a06c5992f2250e1baaf7c90df8e395144e8357e97104)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5086d0cd1e56845566e0b292c534636ba8d58a31eb2f37d5c34b0cd524e6ce8a"></a>

## ingress_egress_gw.allowed_vip_port — ingress_egress_gw.allowed_vip_port / 94fb13e202d2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.allowed_vip_port

<a id="canonical-e2dcf82a43c2e49d2780d22ed25ab13392b2e45615b73a8821c1554c2bd2113a"></a>

Type: `"single"`. Computed.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

<a id="canonical-45391c1edb17a0979287c668a3bea02a5c3e33b6385eda79d563f4c11fd34002"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port / 94fb13e202d2 / 3

- [custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-7b86656377b3241993e859f140ebe5433f4877e6f46d46760885744cc82a62ac): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-578c2e78a9e9678d23417dd81024afe84a6f7fc2025d6bf31c008504b02f318d): complete subsection reference.

- [use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-8d370ef231d793773a287f86f25894e6613bfabb7ff0ba7b1e03feb2cfb1f883): complete subsection reference.

- [use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-73359235d66fcc1d193663b2f7a45e900d4fb0d28f3f671d79c5734725058d70): complete subsection reference.

- [use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-cb1bb2247dd88cca8e5c3be523b115b36e1b4199cdf4b27bf9f958f153e5a583): complete subsection reference.

<a id="canonical-f7b20aa7df52de084dd4fdac7003f03f920a70ff795b5a8dceb290e2d3315407"></a>

## Next pages — ingress_egress_gw.allowed_vip_port / 94fb13e202d2 / 4

- [ingress_egress_gw.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-7b86656377b3241993e859f140ebe5433f4877e6f46d46760885744cc82a62ac)
- [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-578c2e78a9e9678d23417dd81024afe84a6f7fc2025d6bf31c008504b02f318d)
- [ingress_egress_gw.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-8d370ef231d793773a287f86f25894e6613bfabb7ff0ba7b1e03feb2cfb1f883)
- [ingress_egress_gw.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-73359235d66fcc1d193663b2f7a45e900d4fb0d28f3f671d79c5734725058d70)
- [ingress_egress_gw.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-cb1bb2247dd88cca8e5c3be523b115b36e1b4199cdf4b27bf9f958f153e5a583)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-7b86656377b3241993e859f140ebe5433f4877e6f46d46760885744cc82a62ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1fede6887adbbcc2e1cea5b61218b4dc73eea6eb96403b994853976c6616586"></a>

## ingress_egress_gw.allowed_vip_port.custom_ports — ingress_egress_gw.allowed_vip_port.custom_ports / 5035c7736935 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- ingress_egress_gw.allowed_vip_port.custom_ports

<a id="canonical-722d2f4f028de81b574b8726052cc8abc397d4c03ebab8184748b1b4f6d80e54"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

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

<a id="canonical-7b9ec49a7f0f7562ceb7f424d47cf18b82935f1acd23223d2568c4d77b4ec019"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.custom_ports / 5035c7736935 / 3

<a id="canonical-6b99cf6545bbb3694726c3f812992218832dfd688e9b939915629779b183cae0"></a>

<a id="canonical-f033317b4a19be2fcca9d0f2c6024d0eedb388cdb16db7d2507b628cab453437"></a>

## port_ranges property — ingress_egress_gw.allowed_vip_port.custom_ports / 5035c7736935 / 4

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-5e36a0f78d92f0f7de8e31828b8401a1cc92626229c3e5e8ab3d0b6547a0d753"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.custom_ports / 5035c7736935 / 5

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-578c2e78a9e9678d23417dd81024afe84a6f7fc2025d6bf31c008504b02f318d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e537ed4c4d099f3f58c4ef2881421651a0c9ee1ccbb780ae94a37c5fd701d7e"></a>

## ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port — ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port / 2df5c08f7d8b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-1df93eb95ccec2fe23bf8a9377649883a18109ca40fce8ca0ea8ee2dd0cb6b0e"></a>

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

<a id="canonical-fb62aaea2394759d2b2eacd0aa422802edc778418d57d9349457d04bb4d72db7"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port / 2df5c08f7d8b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-075a7175a571d6108c27d915e8fb6fb7808c980b9eaa1d1372f45a0957e965f8"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port / 2df5c08f7d8b / 4

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-8d370ef231d793773a287f86f25894e6613bfabb7ff0ba7b1e03feb2cfb1f883"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90806973195421bd771029e65b4b747cdbd0dc2efc93fdd80abc82a2c3fa54ce"></a>

## ingress_egress_gw.allowed_vip_port.use_http_https_port — ingress_egress_gw.allowed_vip_port.use_http_https_port / 963164f2a30c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- ingress_egress_gw.allowed_vip_port.use_http_https_port

<a id="canonical-7c121dbda2f667aff7aea2a3186ae1c29a48e35d13fa08524fb97126d108a7b0"></a>

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

<a id="canonical-5aa69b7c6f754c15f2c969134bab59e4052ce0928466b043f1a7f481fac0fe98"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.use_http_https_port / 963164f2a30c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b8dcba4a9dd83cfbf5e57ed7b7027a8ab29f1beb74d8610bcdfe7f65a5c56eb"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.use_http_https_port / 963164f2a30c / 4

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-73359235d66fcc1d193663b2f7a45e900d4fb0d28f3f671d79c5734725058d70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98049f3b19da3cc80a83c65480eb4a86c33cf37f64320d1cdb63f8aad0466d11"></a>

## ingress_egress_gw.allowed_vip_port.use_http_port — ingress_egress_gw.allowed_vip_port.use_http_port / 444e6ff1fc82 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- ingress_egress_gw.allowed_vip_port.use_http_port

<a id="canonical-fb12bd0262dd722998e25b5d97c1e116ff3d975c075c1599a10b5fcfdb978ab6"></a>

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

<a id="canonical-7632b717ec0034ab6e73c0b8b3168297dd467da69c7a0f42458edbf5e01cffd8"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.use_http_port / 444e6ff1fc82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-24ac2be3ecb2f9f5e3ebfc120e29c9cc03db34948b8afaed5c81455ea3085153"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.use_http_port / 444e6ff1fc82 / 4

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-cb1bb2247dd88cca8e5c3be523b115b36e1b4199cdf4b27bf9f958f153e5a583"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa41142e920dfb2bd1a0453583dfa7cf24b26b6d71d8932b96d7c38301a2f581"></a>

## ingress_egress_gw.allowed_vip_port.use_https_port — ingress_egress_gw.allowed_vip_port.use_https_port / e8cd58712f46 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- ingress_egress_gw.allowed_vip_port.use_https_port

<a id="canonical-a81b3919d58ca353f1bc3db302bd493fb62bca4feb8ac4bcacd0010efa1aeffd"></a>

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

<a id="canonical-0544f4c45779b7449833ec62ddf833af71f39581514c0a188e340668fe1f7b83"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.use_https_port / e8cd58712f46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ccbcecc06dcc85c2cdbe34270b2ca45429f6a050566a1c61052ccf9bea0a50d2"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.use_https_port / e8cd58712f46 / 4

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-27a8e85fcbea883eb494b892cf42624d4bcc83c4a4d041ba55b64fccabdd081e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-789fbce1be71a04592a2e56cb73823dec43b35203b7bf284b1edae84d09fa945"></a>

## ingress_egress_gw.allowed_vip_port_sli — ingress_egress_gw.allowed_vip_port_sli / f9290558a176 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.allowed_vip_port_sli

<a id="canonical-0fd5a127f65155ac00096004f1e41df56f316f015d2127e37a064cc58ee4a751"></a>

Type: `"single"`. Computed.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

<a id="canonical-062c13e48e1680f6d58df3b98b3cb5776c325a901321d30a292d4535df87d3a1"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli / f9290558a176 / 3

- [custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-5c59ed04689d6aed23f7e92fa934db9732995194d4599425e97600dafc848d57): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-08ec26a489d6dc9c51256b77d0ed7c260ed61f6aa8e12016ff86ff18a4970b94): complete subsection reference.

- [use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-89a1bc6907f64f6eef34025ba1275ed07e6f0333af5d552d2f9a3784849e4d4c): complete subsection reference.

- [use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-81437760668bc53a8ee2623a70f401f8dbea5d01f158bdc3a030255aefb2049b): complete subsection reference.

- [use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-908e166ba54d0c891c56bf22f23bd93550c8870b95844b3551d1723f9c1a2224): complete subsection reference.

<a id="canonical-45aee2ea7096d4a90258d1662654ccdc4c0e631595e5dd0e31bb056dd426b859"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli / f9290558a176 / 4

- [ingress_egress_gw.allowed_vip_port_sli.custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-5c59ed04689d6aed23f7e92fa934db9732995194d4599425e97600dafc848d57)
- [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-08ec26a489d6dc9c51256b77d0ed7c260ed61f6aa8e12016ff86ff18a4970b94)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-89a1bc6907f64f6eef34025ba1275ed07e6f0333af5d552d2f9a3784849e4d4c)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-81437760668bc53a8ee2623a70f401f8dbea5d01f158bdc3a030255aefb2049b)
- [ingress_egress_gw.allowed_vip_port_sli.use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-908e166ba54d0c891c56bf22f23bd93550c8870b95844b3551d1723f9c1a2224)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-5c59ed04689d6aed23f7e92fa934db9732995194d4599425e97600dafc848d57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e302d4c7b3088b52f7b5dd2d55041c09118f12ec67775483b6d50fb5ac4170c"></a>

## ingress_egress_gw.allowed_vip_port_sli.custom_ports — ingress_egress_gw.allowed_vip_port_sli.custom_ports / 1d66e4439641 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- ingress_egress_gw.allowed_vip_port_sli.custom_ports

<a id="canonical-9d6e538eba4e1c2e9682861721c83d8d59aea3dba0f374eab40ed9cb5aa72164"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

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

<a id="canonical-33ec8f3614f649af64c5ffdf58960ee86be4c8df10430ca0da68d811f4d0d9b5"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.custom_ports / 1d66e4439641 / 3

<a id="canonical-fa059e6d18e3cf35f4af3f554fcce780ccba3709ebef79fd4f89f3e771cdb2dd"></a>

<a id="canonical-8d32490423bbccec2b21350b24a23e3e22c1cae815e04d932551be96093d52fc"></a>

## port_ranges property — ingress_egress_gw.allowed_vip_port_sli.custom_ports / 1d66e4439641 / 4

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-e0076399be59fef398eacdfd72b7cfc234f65024b6362bed7da0a6fbd4a75ba1"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.custom_ports / 1d66e4439641 / 5

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-08ec26a489d6dc9c51256b77d0ed7c260ed61f6aa8e12016ff86ff18a4970b94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e5764bf6f7af786bf509ca15d5bf1584051694306a9cb05e3ec10565bc404f3"></a>

## ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port — ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port / 9964e2eb0e97 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port

<a id="canonical-657cf327e956a3828fa46f52d023471f38e4bdd8ccf77a409c90126b50efe3c9"></a>

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

<a id="canonical-0aab7e726e04949095ed743f2b330d1d29bb7d6e383daea38441c84d9739e3c6"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port / 9964e2eb0e97 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b303aadf0909c3374605e2382742128b0a15f30800c3e62ff6a591b15bca295d"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port / 9964e2eb0e97 / 4

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-89a1bc6907f64f6eef34025ba1275ed07e6f0333af5d552d2f9a3784849e4d4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ff0a482ec6e87639fd641391179d94969833f0b8af4a990363d4fc77a5324a4"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_http_https_port — ingress_egress_gw.allowed_vip_port_sli.use_http_https_port / 69ae1cf31162 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- ingress_egress_gw.allowed_vip_port_sli.use_http_https_port

<a id="canonical-457f833b30b3e5ca566e873daa37fb3053ef3948a269ba1e6cc047be2d2993a6"></a>

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

<a id="canonical-2ed9a1b5b60a0802be402d1ac92109eae2b1d4bc0e5beca4186ab9c89792c6ba"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.use_http_https_port / 69ae1cf31162 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48d640ae2f06df0ec3f71fe68fe0c2c802b3951196644df3d080572018e044b5"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.use_http_https_port / 69ae1cf31162 / 4

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-81437760668bc53a8ee2623a70f401f8dbea5d01f158bdc3a030255aefb2049b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df4487767b9a774f529f97a86892bf4715af531b5a1dc0828c03dd90a64002f9"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_http_port — ingress_egress_gw.allowed_vip_port_sli.use_http_port / a5b0477c3deb / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- ingress_egress_gw.allowed_vip_port_sli.use_http_port

<a id="canonical-e3002714a1b4a16ff09ee650ef9797a5e0e574e91a44fb61a4ed52da052eaf4b"></a>

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

<a id="canonical-57ba3952feb49285fc761b0d71a0bd96b391f1c4122878c8a908825ad2fa76f8"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.use_http_port / a5b0477c3deb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b882c78a5447690e0859adb9cdc34b8b0bf2271473add4afb7be73b42c510194"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.use_http_port / a5b0477c3deb / 4

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-908e166ba54d0c891c56bf22f23bd93550c8870b95844b3551d1723f9c1a2224"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a19162a92c04e20e2a0e8245c26211630618eff951c420991c95534e7b327c0"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_https_port — ingress_egress_gw.allowed_vip_port_sli.use_https_port / d959eaeae736 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- ingress_egress_gw.allowed_vip_port_sli.use_https_port

<a id="canonical-0ae522f689ecbdfa3eaa2598e3e37c99701fa357a585d2446132bcf43a3f0034"></a>

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

<a id="canonical-939cc7281be1f28e9da360ac34769615ff9f472b242744e57cc3a831540ff591"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.use_https_port / d959eaeae736 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4065e0e0cd6371f4e2445984abb4e9af2e77aa1e56f57bb9deacdc43fa7433a"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.use_https_port / d959eaeae736 / 4

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-e92f17fba0f34fe6f9dca381a5ce39b289988e2fe2143590f8cec588f3d72331)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89bfc8ff41d124871ce24a148d219fd127770bd7fe603a57c75114b40429b014"></a>

## ingress_egress_gw.az_nodes — ingress_egress_gw.az_nodes / be24e1a32faf / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.az_nodes

<a id="canonical-29e3b9fb11728347a51210058cae464686194a7e0d8d2c5549178c53ca52bb9c"></a>

Type: `"list"`. Computed.

Only Single AZ or Three AZ(s) nodes are supported currently.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

<a id="canonical-828a4a6078f631f14367575d5732085026a413684e254c001043f4733613bc72"></a>

## Direct properties — ingress_egress_gw.az_nodes / be24e1a32faf / 3

<a id="canonical-f2c111696289bdbf3919366325edccd101bfb9f2b4aba7d41fd540e64e27a70b"></a>

<a id="canonical-7d63878cb167441084e61055388903b7bb6d7f3daf749ee86e90e448f30aa14b"></a>

## aws_az_name property — ingress_egress_gw.az_nodes / be24e1a32faf / 4

Type: `"string"`. Computed.

AWS availability zone, must be consistent with the selected AWS region.

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

- [inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-336ab420c145832387f16f4c3b3017f6bb17b734f464e2fa356a5a72fa33c02e): complete subsection reference.

- [outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-4a786f6f81f89681bac3f860b207e1333016e4e12f3d85d73f3b447cd96c4a2c): complete subsection reference.

- [reserved_inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-765e6faaa877686879f4a5229926ecbd91b5c4568d0f640a813c3cafeffbe66f): complete subsection reference.

- [workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-45645b0bbe86ce64c51dfd50236619000d61754670a2550e45932d25441cf6c7): complete subsection reference.

<a id="canonical-a58b92a556af59c48ea92561cd4201818b02232dab8abf1e0658ef8697513f4f"></a>

## Next pages — ingress_egress_gw.az_nodes / be24e1a32faf / 5

- [ingress_egress_gw.az_nodes.inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-336ab420c145832387f16f4c3b3017f6bb17b734f464e2fa356a5a72fa33c02e)
- [ingress_egress_gw.az_nodes.outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-4a786f6f81f89681bac3f860b207e1333016e4e12f3d85d73f3b447cd96c4a2c)
- [ingress_egress_gw.az_nodes.reserved_inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-765e6faaa877686879f4a5229926ecbd91b5c4568d0f640a813c3cafeffbe66f)
- [ingress_egress_gw.az_nodes.workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-45645b0bbe86ce64c51dfd50236619000d61754670a2550e45932d25441cf6c7)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-336ab420c145832387f16f4c3b3017f6bb17b734f464e2fa356a5a72fa33c02e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2add5149fd443f82f8b7ccfc4209ed1238d241ceaf97db5d0c7796519dd7a03"></a>

## ingress_egress_gw.az_nodes.inside_subnet — ingress_egress_gw.az_nodes.inside_subnet / 3529523b4b43 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- ingress_egress_gw.az_nodes.inside_subnet

<a id="canonical-315137bda29437a98ee9aa6e7e81a6ee6629cfec48c7d9ce03d66f994cf80dbb"></a>

Type: `"single"`. Computed.

Configuration parameter for inside subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-41924885fa730ef3d9f8d43778e65b1206c67faa47ef8c27d70a56aae6d2d765"></a>

## Direct properties — ingress_egress_gw.az_nodes.inside_subnet / 3529523b4b43 / 3

<a id="canonical-e6f687636e7ea38071f275fca358391265eb2a8bf9ebba5b97e3e2b279b9fc99"></a>

<a id="canonical-d1509f17c08a92d873c94db971a4c17381c1cee50d3a7a16aab93840a8150a84"></a>

## existing_subnet_id property — ingress_egress_gw.az_nodes.inside_subnet / 3529523b4b43 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-11e33bc1f6d0ece152f36762ff763bc0ec1afd0839430e737e42871bed2471bf): complete subsection reference.

<a id="canonical-22cfbbbb27b4d264dd73a36f3bbd611e2ec5358a2f5b82a77c0e7bbf44352555"></a>

## Next pages — ingress_egress_gw.az_nodes.inside_subnet / 3529523b4b43 / 5

- [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-11e33bc1f6d0ece152f36762ff763bc0ec1afd0839430e737e42871bed2471bf)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-11e33bc1f6d0ece152f36762ff763bc0ec1afd0839430e737e42871bed2471bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e72855dac9f0e6a1c1c1199511a72ed3e62d3b7fe39d9147fe2bb48974b6c4a1"></a>

## ingress_egress_gw.az_nodes.inside_subnet.subnet_param — ingress_egress_gw.az_nodes.inside_subnet.subnet_param / d2a1ef33cda0 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- [ingress_egress_gw.az_nodes.inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-336ab420c145832387f16f4c3b3017f6bb17b734f464e2fa356a5a72fa33c02e)
- ingress_egress_gw.az_nodes.inside_subnet.subnet_param

<a id="canonical-cef5a8343b4dafff8e645341a11c546cc3330e3309c9b41ab97365e65c356a55"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-5ab1b2fdd49bad9775eed9d1946514ba1a5998d872fa0560ddade5f734915de1"></a>

## Direct properties — ingress_egress_gw.az_nodes.inside_subnet.subnet_param / d2a1ef33cda0 / 3

<a id="canonical-24db48ad6b0f43791d99ce284d071e7fc088562788774a781f08d49a14127536"></a>

<a id="canonical-1294c328ec56cfca3c5d6cc3579dc4a05e02da4e71931b664184029974333c85"></a>

## ipv4 property — ingress_egress_gw.az_nodes.inside_subnet.subnet_param / d2a1ef33cda0 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-9aca8f95c90447156ccd58db7859a24026d51971f17fc2420d3d240834814235"></a>

## Next pages — ingress_egress_gw.az_nodes.inside_subnet.subnet_param / d2a1ef33cda0 / 5

- [ingress_egress_gw.az_nodes.inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-336ab420c145832387f16f4c3b3017f6bb17b734f464e2fa356a5a72fa33c02e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-4a786f6f81f89681bac3f860b207e1333016e4e12f3d85d73f3b447cd96c4a2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0e777dc4f5fb99b15a040163cafe0fe9de7a77e7f69345a6c13a5a80742fa86"></a>

## ingress_egress_gw.az_nodes.outside_subnet — ingress_egress_gw.az_nodes.outside_subnet / 4740a345e5d4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- ingress_egress_gw.az_nodes.outside_subnet

<a id="canonical-b91a6b60febe15b641e8cab8db705ce10cc2f4b2402c20028974603e5ef4241e"></a>

Type: `"single"`. Computed.

Configuration parameter for outside subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-52e8a6febeee671334592e4f14f755d47df1b6dc547631bed059c1dac5cf80d5"></a>

## Direct properties — ingress_egress_gw.az_nodes.outside_subnet / 4740a345e5d4 / 3

<a id="canonical-244f4a27a6c57fe9a10d7a0548fa102790517e5eef2fed4af04e30420d5743b7"></a>

<a id="canonical-8a6a63d4d6c22154e655a84f1b6a887c246aa6d16cd35e44503d9b27b2c26002"></a>

## existing_subnet_id property — ingress_egress_gw.az_nodes.outside_subnet / 4740a345e5d4 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-de94564cbfb0ca0425dc3e4facc1a3abc4031f8271ef28954e5d5ac2d4247dc6): complete subsection reference.

<a id="canonical-076ed77e7a2dd6e8fade611623551848bee38fcc7989088cf33d9be56e90cee6"></a>

## Next pages — ingress_egress_gw.az_nodes.outside_subnet / 4740a345e5d4 / 5

- [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-de94564cbfb0ca0425dc3e4facc1a3abc4031f8271ef28954e5d5ac2d4247dc6)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-de94564cbfb0ca0425dc3e4facc1a3abc4031f8271ef28954e5d5ac2d4247dc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51844f63b24169f983d64ba1e916dfbe47c1624e6a69f85c56d5f57b318837a5"></a>

## ingress_egress_gw.az_nodes.outside_subnet.subnet_param — ingress_egress_gw.az_nodes.outside_subnet.subnet_param / d2233c36b8ef / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- [ingress_egress_gw.az_nodes.outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-4a786f6f81f89681bac3f860b207e1333016e4e12f3d85d73f3b447cd96c4a2c)
- ingress_egress_gw.az_nodes.outside_subnet.subnet_param

<a id="canonical-45b00ef1713fbc4f33be6b17dc45ff62e8959edf1408653d8b2ec2b95c363975"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-b69f72610a7c292d15452e1529f5ede69bd38534ee703a7269836d21d59de4cc"></a>

## Direct properties — ingress_egress_gw.az_nodes.outside_subnet.subnet_param / d2233c36b8ef / 3

<a id="canonical-6f325b172d23a7c4ac8474bdc4464a6cffb9c9837d78ad3855e6856a58885d07"></a>

<a id="canonical-49b52572c4c46fdac7e0186473af97155ecf1181454e323ecdbd37daa816783a"></a>

## ipv4 property — ingress_egress_gw.az_nodes.outside_subnet.subnet_param / d2233c36b8ef / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-ea18f595c5dac1f4f05bd42b1003118b24c0166994976d3f6d0c32a9ab1e0f4d"></a>

## Next pages — ingress_egress_gw.az_nodes.outside_subnet.subnet_param / d2233c36b8ef / 5

- [ingress_egress_gw.az_nodes.outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-4a786f6f81f89681bac3f860b207e1333016e4e12f3d85d73f3b447cd96c4a2c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-765e6faaa877686879f4a5229926ecbd91b5c4568d0f640a813c3cafeffbe66f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f80e32cf3187e6a4839e3953b29c50e06e1972708b1f06d88b0e454b05463e4e"></a>

## ingress_egress_gw.az_nodes.reserved_inside_subnet — ingress_egress_gw.az_nodes.reserved_inside_subnet / 696c9d7e5d03 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- ingress_egress_gw.az_nodes.reserved_inside_subnet

<a id="canonical-6b4170c9d6f64469d3ec880948163c65dca73a4aaf31dc60d98989a8faeb6cfa"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reserved inside subnet.

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

<a id="canonical-f87efe96762dfca5dcd805f53a343e10a01f3b336a10a7ed578810ebfc03cf8f"></a>

## Direct properties — ingress_egress_gw.az_nodes.reserved_inside_subnet / 696c9d7e5d03 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46fa82e3285ff9e6051def9c670fe71bec2d909c4b9acee561353501bc253790"></a>

## Next pages — ingress_egress_gw.az_nodes.reserved_inside_subnet / 696c9d7e5d03 / 4

- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-45645b0bbe86ce64c51dfd50236619000d61754670a2550e45932d25441cf6c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9bb30ba2decdb2b0d73862821e38d82d069674b578de6d06defd896fea37ce1"></a>

## ingress_egress_gw.az_nodes.workload_subnet — ingress_egress_gw.az_nodes.workload_subnet / fb520d90b18b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- ingress_egress_gw.az_nodes.workload_subnet

<a id="canonical-3864657aed666dd9ab4d27391d0d5d3255c6364f4c8ea2264625a2de016861fd"></a>

Type: `"single"`. Computed.

Configuration parameter for workload subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-d84332f33e68be1373aaaec2424c049b38e9d39f2f73d2a56d9d47d88fcce1bb"></a>

## Direct properties — ingress_egress_gw.az_nodes.workload_subnet / fb520d90b18b / 3

<a id="canonical-e261a82eb30f2345d70187b4a99a010a3ad6f11981a5fb5b6106e7d58c1e9174"></a>

<a id="canonical-be7bc51762898e4fbefa65b8a3589b26fc697cd47b8f4d4208d752177e78e62f"></a>

## existing_subnet_id property — ingress_egress_gw.az_nodes.workload_subnet / fb520d90b18b / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-7b9079d8507975e10f9c3c4e8b04e2b737b01aa28078bef4729a07c94f955579): complete subsection reference.

<a id="canonical-e04e57cf6ad1424cef4084906a58e486b57c8635a29126c6e474fac371f5c117"></a>

## Next pages — ingress_egress_gw.az_nodes.workload_subnet / fb520d90b18b / 5

- [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-7b9079d8507975e10f9c3c4e8b04e2b737b01aa28078bef4729a07c94f955579)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-7b9079d8507975e10f9c3c4e8b04e2b737b01aa28078bef4729a07c94f955579"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39d25bd7c1e64073b34dd6f303d0945bdf1e0cad82e113304c9ffc129cad88a3"></a>

## ingress_egress_gw.az_nodes.workload_subnet.subnet_param — ingress_egress_gw.az_nodes.workload_subnet.subnet_param / 7b8a6c492b60 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-90cc30ef730339904b03af07d44ba8a1a5305d520a35b9fc4bd67a75a54e9a26)
- [ingress_egress_gw.az_nodes.workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-45645b0bbe86ce64c51dfd50236619000d61754670a2550e45932d25441cf6c7)
- ingress_egress_gw.az_nodes.workload_subnet.subnet_param

<a id="canonical-bc883d794512cc03f2afee012b9966e7773109b0982c7ea23d5c22cb7951f44c"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-1cfcbcd57fd3b646424abe77987b5ce9cd02e87d7f241f7b74986035809ed64e"></a>

## Direct properties — ingress_egress_gw.az_nodes.workload_subnet.subnet_param / 7b8a6c492b60 / 3

<a id="canonical-79ecacdbb7b8f7c853cbcdcb7249ee86708b3b868d1c87c83cef9639fe5a0cc8"></a>

<a id="canonical-a78edafe16b97dfdbfa96594f837234061f0f3cca1976311a7cd0539ad4be7db"></a>

## ipv4 property — ingress_egress_gw.az_nodes.workload_subnet.subnet_param / 7b8a6c492b60 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-44fc10fd71955b61103b33366259befdd30783a1bfc9009d80f716d73c6e6889"></a>

## Next pages — ingress_egress_gw.az_nodes.workload_subnet.subnet_param / 7b8a6c492b60 / 5

- [ingress_egress_gw.az_nodes.workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-45645b0bbe86ce64c51dfd50236619000d61754670a2550e45932d25441cf6c7)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-2a6852e810022ac31544b2447fb302bd35b1ae58b86a042440ae79d0fbd8dfbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c8468d24d061bab90399534bfc037426e5056a21db62436900622bab0ebf3dd"></a>

## ingress_egress_gw.dc_cluster_group_inside_vn — ingress_egress_gw.dc_cluster_group_inside_vn / d7b949c434b6 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.dc_cluster_group_inside_vn

<a id="canonical-72f5dcfd839e3b2032383615ed0623c2754f259abe9e242836ecc7961b45de4f"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0791b46f8419b63c73745f5a78b360564dd056168fc25bbda6c40ed35d3abe2b"></a>

## Direct properties — ingress_egress_gw.dc_cluster_group_inside_vn / d7b949c434b6 / 3

<a id="canonical-61816814407d2fa22920ff5cd934c2e2e7c17be50ba3c5788d49be2282efd091"></a>

<a id="canonical-b107c27aa25f6de69f8da991da9aeaa98cbe29852ec958fc9af62fee389ad90e"></a>

## name property — ingress_egress_gw.dc_cluster_group_inside_vn / d7b949c434b6 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-91148a4776aa94d49c5fb4ee86aae40e54cc88516813848ecdec4f82b7330000"></a>

<a id="canonical-bb359314b8413ffa038c0fc498a4aec1d1654a593dfffc6aa579def7ccd2896f"></a>

## namespace property — ingress_egress_gw.dc_cluster_group_inside_vn / d7b949c434b6 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-f95cc387036a8a44d45e915be0bebea1aaf9b557c054cffbd1e221811c5350fc"></a>

<a id="canonical-656a8c0c857c45fbf7726df300bd69087f46872924d680d3f45ec5f5b7a77caa"></a>

## tenant property — ingress_egress_gw.dc_cluster_group_inside_vn / d7b949c434b6 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0a265e961c4d8ae8b13d2ad74d65bb64e6b4eadffc584772c53cc8abb65a51ab"></a>

## Next pages — ingress_egress_gw.dc_cluster_group_inside_vn / d7b949c434b6 / 7

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-08b710f8cb19ae0b29bb024ace9c0c0ca77b40114ff5e92280fcc72bc1ed3078"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22cb9b9302fcc4233c574e348f98c9c7a5789e531b5499a8b696504575eb3945"></a>

## ingress_egress_gw.dc_cluster_group_outside_vn — ingress_egress_gw.dc_cluster_group_outside_vn / b5c557c4e9d2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.dc_cluster_group_outside_vn

<a id="canonical-e2c05b684d8a16ff6994951ad9fd9460824854e559144b95e178c8cdaa52cb7b"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-78592a4626958206a43e6e7f89b3458d12830ed4544ac8a4ac576658cec7ab14"></a>

## Direct properties — ingress_egress_gw.dc_cluster_group_outside_vn / b5c557c4e9d2 / 3

<a id="canonical-2716edb4b0f594c6bfcd75e54921f68e91af7b2bb3d2dfdd0a75002d849d17fa"></a>

<a id="canonical-f2c30b520f491d4d114f0b95aaf72faa93e6824ea11df69b6bd4e9b16a4bd235"></a>

## name property — ingress_egress_gw.dc_cluster_group_outside_vn / b5c557c4e9d2 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-5f180e8ba8e7021176f4ae358b14337be3a403acf3a55ede767e20689a8e561a"></a>

<a id="canonical-da251496d2698ad965ed26250c0a59be23ae5e45a11867353469bdf84183d2ba"></a>

## namespace property — ingress_egress_gw.dc_cluster_group_outside_vn / b5c557c4e9d2 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-025eb969eb8dd591670569d1f6ad98e8f8d790c4d898177eb280ee94f530bb47"></a>

<a id="canonical-5912491a21a0db04485f06eba6bf4d80016698460611be31fa5a522680495c1f"></a>

## tenant property — ingress_egress_gw.dc_cluster_group_outside_vn / b5c557c4e9d2 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-b03276c0f8ce9b9f21094d852a69fe0e822480f7daf161902aa5c518396dbec1"></a>

## Next pages — ingress_egress_gw.dc_cluster_group_outside_vn / b5c557c4e9d2 / 7

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-42e3e0c67afb1f7e797175ec0767137bb70771de9c4fe9414268ea5d21bfff49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eac49e998664ca2c664474c4e43b664965c2f08b67b27d440f7d3db053dd84ce"></a>

## ingress_egress_gw.forward_proxy_allow_all — ingress_egress_gw.forward_proxy_allow_all / 49d5da979df0 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.forward_proxy_allow_all

<a id="canonical-4bf09b6012f00f3ba136d25aae4021f8b2db0ce973bb476260d15dee2dd25b54"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for forward proxy allow all.

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

<a id="canonical-87ac93585407bac32980c27ee459be2882d619461a768be44aa8435da0f97ae3"></a>

## Direct properties — ingress_egress_gw.forward_proxy_allow_all / 49d5da979df0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1435bfb17eb14a4d7482adfc87a29230610a57f5a93556b13985d63a452c0c7"></a>

## Next pages — ingress_egress_gw.forward_proxy_allow_all / 49d5da979df0 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6755ef88e1ef3a44b8b61a2879f18c434ea8010b23338785c027d509db24ac11"></a>

## ingress_egress_gw.global_network_list — ingress_egress_gw.global_network_list / 4f625a503990 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.global_network_list

<a id="canonical-5521a81083b62d98e7da46e24e813c488ec85e8c56c22ead1966d72608c3248a"></a>

Type: `"single"`. Computed.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

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

<a id="canonical-c434b1b457f3af3383be0f068078ced78a3e9b5ae99a108a3679d139b8074e37"></a>

## Direct properties — ingress_egress_gw.global_network_list / 4f625a503990 / 3

- [global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2): complete subsection reference.

<a id="canonical-6b0dfaae189aaa23a43aa45a8d0431d58d569b3114835f97481f0fbbbccf53eb"></a>

## Next pages — ingress_egress_gw.global_network_list / 4f625a503990 / 4

- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72950ff39f0525ebdf5be7642dc74375edcfcf4c0f3a44290c6d63be72d9cd06"></a>

## ingress_egress_gw.global_network_list.global_network_connections — ingress_egress_gw.global_network_list.global_network_connections / 6dc22e7672b3 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="canonical-05862f8af1d646fd3e5ffe76f006f3531ffae38f22c128a04754c3096dd901e4"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0484b894c09a2e26ed2aa892e3ffe1ce9cecf70b7c21b15ccbd937ec11bcf108"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections / 6dc22e7672b3 / 3

- [sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-2f7e0b4bc738780556e410378802c8160706bf73653ab06b1cc019f96f448863): complete subsection reference.

- [slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-5b88003b77cd031af5142ff93741877301fbcbdbad3043b1b5b01bd7e8cd9a91): complete subsection reference.

<a id="canonical-0e0ea370378d73ddea5642eae407db4dda911db9028b64e03444538e95265f77"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections / 6dc22e7672b3 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-2f7e0b4bc738780556e410378802c8160706bf73653ab06b1cc019f96f448863)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-5b88003b77cd031af5142ff93741877301fbcbdbad3043b1b5b01bd7e8cd9a91)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-2f7e0b4bc738780556e410378802c8160706bf73653ab06b1cc019f96f448863"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39aa36fd0e09eb5f529f1933a7cd280febe93eaa990dfdcc161ffb48eb224cfe"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 29991f76de7b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-4188507b8b4d44c76aad65e36eef3b4a412835924b1338aa61e5d9a97d144da7"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

<a id="canonical-576eb054ebb2153bf7a27ee0d35298cdae96b4ddb53a0f0d2835fbbf6249d2c4"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 29991f76de7b / 3

- [global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-7f7f2548ad154ef6057cfbf198e314a27d28c145bc43eb6d9d323fc66b4fc635): complete subsection reference.

<a id="canonical-c4fcec4eb511eed235a4b4f7dbc46a583a9ea29d06172101322c7ccae1633969"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 29991f76de7b / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-7f7f2548ad154ef6057cfbf198e314a27d28c145bc43eb6d9d323fc66b4fc635)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-7f7f2548ad154ef6057cfbf198e314a27d28c145bc43eb6d9d323fc66b4fc635"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fe365f2296882b06858bd915fc1dc938dfe19e88233e0a57c5c8eb5a76f07cb"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 8e215be284b7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2)
- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-2f7e0b4bc738780556e410378802c8160706bf73653ab06b1cc019f96f448863)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-29af3ba871a0a2b53d6a6084c6550214fe0812b9a3d979eacb9de7dd94b42273"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-52d63419f35a945878c2f55e8a4592de495177b50428f0f340b61664d1228aa6"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 8e215be284b7 / 3

<a id="canonical-8f5ddb808ab7d5239ecac9a808d9aef37eb560a9f8f6a957ef8966502fb41ae2"></a>

<a id="canonical-96e9c75c69ae2348b9a2fa65bfd31e16813a01dd69b5eea83e98e6dc80e6f25c"></a>

## name property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 8e215be284b7 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3a4ed092f9dd647fec552a41a8d339840f26626bf9e4a09d3f5fbc8816c78d81"></a>

<a id="canonical-f29bf15bddf90d739094ac767d47dcb238d5b4c8605da546254a43e46a7e92fc"></a>

## namespace property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 8e215be284b7 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-c969f95f27e91b4dc3218382eda1ec082c16cfc289de2910460c1947ae14b2a4"></a>

<a id="canonical-538187d8fd2977400eecd9d7822709ce936fe86a2900369059a846a8ba3221ca"></a>

## tenant property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 8e215be284b7 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-912a24939861219ed513b71850f8adf6764f638ce594b90860f2bb46d522c852"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 8e215be284b7 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-2f7e0b4bc738780556e410378802c8160706bf73653ab06b1cc019f96f448863)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-5b88003b77cd031af5142ff93741877301fbcbdbad3043b1b5b01bd7e8cd9a91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7da3bc78c73e3737013df53971dd6c93b6a74bdbda37c9dbe8cb12eec8d81bb9"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 290c298dfdc6 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-5e74501abe51c29e1f93ed0d4ff47a676becb8fb388e5b5294da9d8b5f4da4a1"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

<a id="canonical-07d55ef2ee8e1b9618738f016133546db6a23e045ab30b3869e0455ca9d357b5"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 290c298dfdc6 / 3

- [global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-efd1aefa988db06c8239b2e70d33c789c7635f6c9a5b9e0d43d8e103b9b84cdf): complete subsection reference.

<a id="canonical-69e386747bacd8e648f6444bb8a78dde7527d1d742eb45bb9327f532fbcf450a"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 290c298dfdc6 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-efd1aefa988db06c8239b2e70d33c789c7635f6c9a5b9e0d43d8e103b9b84cdf)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-efd1aefa988db06c8239b2e70d33c789c7635f6c9a5b9e0d43d8e103b9b84cdf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e730d2e50130ecfdc6516e607ac1a1fddea2928954d0affb60e0ad44724cd9d3"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 64d7a384db86 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-7e6bb59ec2a628d656adab14e28c86c74566aa20bdfc2dfbc015ec0688341279)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-3ca7f35fdeb2714af89d184e5dc292601283fd5addf942e7db7a4b4cd8b2b7e2)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-5b88003b77cd031af5142ff93741877301fbcbdbad3043b1b5b01bd7e8cd9a91)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-f7e17e9978f36d12ee61b1e91b3394693dd8d184303876d1e59d51de1c5be722"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-80c32381ffcd5bfeba480f2fc124a6050ff7753029b7dbf009860b2e7fb0e858"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 64d7a384db86 / 3

<a id="canonical-e5bfaf3ae0bbce04df3b1ad232e79b6bbcb5bd274741e6e2b9154a2b15ffa2a5"></a>

<a id="canonical-18dfcaffc87e7d3fff89c5e41a49e4cbe97b9db6b80d063f2e2bd05eb39a89b7"></a>

## name property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 64d7a384db86 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-4ac3962137596f256ef1287c4c2a19d6cd588fe1e587bcfc168d4c37c24e146f"></a>

<a id="canonical-f42a50a352bde9935882993cc8be5e38e821c70bd865724a60a109c8fae6ac9f"></a>

## namespace property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 64d7a384db86 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-10e1f1ed059fc725f0ccbdcd190b3fd478e7ce597e8510386b4cbc9219d8822e"></a>

<a id="canonical-69df1432eda5a1874c8096179376155fc8641a7cf36953787bdae6610a68f1b9"></a>

## tenant property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 64d7a384db86 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-6b4265750b1c9bcbad9a8011c755a802d38d24260d4c3ea11670d5e882bbdd4a"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 64d7a384db86 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-5b88003b77cd031af5142ff93741877301fbcbdbad3043b1b5b01bd7e8cd9a91)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-1cd789f89f4736315a9dbca490acb46cd71979de8e8aad79b5856820c301d900"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d432f6c4a0c371835ac1c99b57295eeadf1016f9e37795e661b9a95bec79f138"></a>

## ingress_egress_gw.inside_static_routes — ingress_egress_gw.inside_static_routes / cd18aa37dafe / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- ingress_egress_gw.inside_static_routes

<a id="canonical-3386420ad603ea21d4462c53d89bbadb38be73cb65e6ff506fb4b1b47b50aafc"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

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

<a id="canonical-ccdb4ebdd160854aea4a104c6fd9527181085e34267c7ad1a39e88a7704ed937"></a>

## Direct properties — ingress_egress_gw.inside_static_routes / cd18aa37dafe / 3

- [static_route_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-e571291ca7af70000fd40835f82139e158e9d0ff28c92aeed44aa2941aac831a): complete subsection reference.

<a id="canonical-b1abf892b3b053a34fd52c23f81ceb7d1ae39e944e216111103673f25a80480a"></a>

## Next pages — ingress_egress_gw.inside_static_routes / cd18aa37dafe / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-e571291ca7af70000fd40835f82139e158e9d0ff28c92aeed44aa2941aac831a)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-e571291ca7af70000fd40835f82139e158e9d0ff28c92aeed44aa2941aac831a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5d954692142116bb99d03986ce99cbad9e2b75c767460a303768372767914a6"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — ingress_egress_gw.inside_static_routes.static_route_list / c26ad08fbab8 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-002.md#canonical-1cd789f89f4736315a9dbca490acb46cd71979de8e8aad79b5856820c301d900)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-ebccf0a40a4275148d96809d79c748a69b1d6b27150583f6b12ce12e30ec4288"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-54a036a3b29d40da7cc89de596dd2fa037cdcf1c6fe583fb4b8429025bd8acae"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list / c26ad08fbab8 / 3

- [custom_static_route](data-sources--aws_vpc_site--reference--group-002.md#canonical-cd2f53aae84e47a1703cbce1e18308012ac3e5044f88102b0544cdee626b9eb2): complete subsection reference.

<a id="canonical-fdca5986a767f77eb2a988c5e69ee5a69cf99062c3eb547a651889071596e41b"></a>

<a id="canonical-a6785e8906b2c89b104774cd4544293e722ca71492fd4f2a9a8255927654e0c9"></a>

## simple_static_route property — ingress_egress_gw.inside_static_routes.static_route_list / c26ad08fbab8 / 4

Type: `"string"`. Computed.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-e171bbdb5bd8f809704e8477d4b77e368f0d730d60755a1f5a0c55fee72b678a"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list / c26ad08fbab8 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-002.md#canonical-cd2f53aae84e47a1703cbce1e18308012ac3e5044f88102b0544cdee626b9eb2)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-002.md#canonical-1cd789f89f4736315a9dbca490acb46cd71979de8e8aad79b5856820c301d900)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-cd2f53aae84e47a1703cbce1e18308012ac3e5044f88102b0544cdee626b9eb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-855303d70fc5e82fdc9e6455681196c044594a7588149c3f09392a59ead97074"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 04187dc6c757 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-002.md#canonical-1cd789f89f4736315a9dbca490acb46cd71979de8e8aad79b5856820c301d900)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-e571291ca7af70000fd40835f82139e158e9d0ff28c92aeed44aa2941aac831a)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-622c74b78190ca224666f357756c6e9b8e10b54fc9818b786013c8a0cc6865dc"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

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

<a id="canonical-8a34a980f326e380135b1340e837905dd86dc30fd248e51637db9fa94dfbcf02"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 04187dc6c757 / 3

<a id="canonical-e2fb42ce1ae32f79ebba007e358e6abbab237a9e47818bd6c009a94daffcf25d"></a>

<a id="canonical-950470e80504bbb762ca85c0bcbd5f48479c149a7364ec812771d601fc0748c7"></a>

## attrs property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 04187dc6c757 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--aws_vpc_site--reference--group-002.md#canonical-406a92a09330b00241b0d3979ece4af8ed6aff75bc560e41c63927062b9bd250): complete subsection reference.

- [nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-74a86d7a620199101b6bc3eb976f85c9d806afdd42efbb4dc69762ef45fb75c7): complete subsection reference.

- [subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-0a43afe403aa22a7b54b42c7bdedfa63b8c45d9ca6e5de7fad95efd2bd2852fb): complete subsection reference.

<a id="canonical-d869b52285dfbfa682722e987e2427ca98fb1628a7b6c03c69dbb1d14fb4b908"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 04187dc6c757 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-002.md#canonical-406a92a09330b00241b0d3979ece4af8ed6aff75bc560e41c63927062b9bd250)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-74a86d7a620199101b6bc3eb976f85c9d806afdd42efbb4dc69762ef45fb75c7)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-0a43afe403aa22a7b54b42c7bdedfa63b8c45d9ca6e5de7fad95efd2bd2852fb)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-e571291ca7af70000fd40835f82139e158e9d0ff28c92aeed44aa2941aac831a)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-406a92a09330b00241b0d3979ece4af8ed6aff75bc560e41c63927062b9bd250"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
