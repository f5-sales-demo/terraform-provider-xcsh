---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-99b1d145f0696b64a55adb9a4b7bdeeab874a158430dccf073b231446527479e"></a>

## latitude property — coordinates / 4f412f35ce63 / 4

Type: `"number"`. Optional.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-56f3f87e4e5bfef252c15a29db34f605f7512f797dcf8eb527b861b8b87a6545"></a>

<a id="canonical-c9fbb4950ed99b894a75f4745bbf8bf4eca1b9f0c6a0e026dac4ee158567d12b"></a>

## longitude property — coordinates / 4f412f35ce63 / 5

Type: `"number"`. Optional.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-3ea63dfec0332d24c80ddb8e88b62b4b40095271aa73760cb7b34e3530d08b25"></a>

## Next pages — coordinates / 4f412f35ce63 / 6

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-09c7d204935c84c5375a63cbc14e06bf5157500ed84092fa1f923683ce3bca7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-435ed83a15b11363db43fd6ecda69973918e3cdc2f5a503a011bf64202d69891"></a>

## custom_dns — custom_dns / d02234779ea2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- custom_dns

<a id="canonical-c79fec7f021e02efd47b16fb9f88c76f9c5d56bc3e59006c71a1acb879fdcdeb"></a>

Type: `"object"`. single nested block, Optional.

Custom DNS is the configured for specify CE site.

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
custom_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d3762dcaac296fa395859c15bd858258b07819b4bec15e4d25f457eea0d0dec"></a>

## Direct properties — custom_dns / d02234779ea2 / 3

<a id="canonical-7a845f568fa98625bd6a456a50fa611073f8774467481718c33370da5c016fef"></a>

<a id="canonical-d8e8075c08be7ffc6b38b346a1197e0740bd0710ee75af92476d6699758bd4b5"></a>

## inside_nameserver property — custom_dns / d02234779ea2 / 4

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in inside network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-7e034b024b412634c1d2bdd397d48792a1030416dfae8a2e39adc18f382d2021"></a>

<a id="canonical-c3daf2676ee1b2080eeb167ca81d6b837f3d3ffe7394edab21bc95aa8aa0669e"></a>

## outside_nameserver property — custom_dns / d02234779ea2 / 5

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in outside network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-682d9901f932e5c8913c3dbb5430939889851b22ef13f9b74c4cf42144cdd889"></a>

## Next pages — custom_dns / d02234779ea2 / 6

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-5b7a3935b1838b49ee2d084b292388d079eab2dc745cc84a61f53f9b2aadde8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abe7f65e5d9fc1247fa6c596cf1f3e9851618296a9c23cd99ce12198983f5909"></a>

## custom_security_group — custom_security_group / 0041575857b6 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- custom_security_group

<a id="canonical-9274af8148c1ca3ab69710d7190b671ae70a85d8f46291700143952cfbf04e45"></a>

Type: `"object"`. single nested block, Optional.

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

- [custom_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-9274af8148c1ca3ab69710d7190b671ae70a85d8f46291700143952cfbf04e45)
- [f5xc_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-d5baaebb13185ddb9f08cf03b5a2186aeaf97eabc9c0cfc2c226fe71a81f2fef)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_security_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-ba752b2cc3a2068b364cf77482ea9bbe201cf2feb82f0e6ca00bee7f5179eea5"></a>

## Direct properties — custom_security_group / 0041575857b6 / 3

<a id="canonical-c24437fc82e6c4efbcdeb62d86076bc82c88e0d79eed39b1599f04d3f439d077"></a>

<a id="canonical-097c936e488699b23080c69ed224a0064755dec4930056e16ed037fc2832a8ea"></a>

## inside_security_group_id property — custom_security_group / 0041575857b6 / 4

Type: `"string"`. Optional.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

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

<a id="canonical-650ca3159c41242c8a6c1f9540c04597083d762f7aa0c5d6492569185230177b"></a>

<a id="canonical-a2514c04945d7c6afd0c4a73149dee3153f8c19ce121c6bc4377e5e9c76fee82"></a>

## outside_security_group_id property — custom_security_group / 0041575857b6 / 5

Type: `"string"`. Optional.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

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

<a id="canonical-eece22c8b275ebfff13a0fbc05796cdd9fb12ae91cc2961ce31434a756b9b010"></a>

## Next pages — custom_security_group / 0041575857b6 / 6

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-4b76809e41769ed01b26b6c290952e2476c7e3775f91c338121b0d7d0ad3d43d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5584b94c1e8a1f452b43e37c9861f2ad52a31e3a06ed6f857e798254a44f42a3"></a>

## default_blocked_services — default_blocked_services / addb1ee3821e / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- default_blocked_services

<a id="canonical-107d8daa2a0b4ba72fa0da3c56e81d32d6c424b7b554a889c0ef35890a717bb2"></a>

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
default_blocked_services = {}
```

<a id="canonical-fcb9029169cf3a6e8eeedae9130d46eda92d93e30b4edfc61d1919e359ba96f2"></a>

## Direct properties — default_blocked_services / addb1ee3821e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc2668084d8cefe043e1617914cc644ec2abb7475191351c2ede43ebc27ec1b4"></a>

## Next pages — default_blocked_services / addb1ee3821e / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-63bd3819b7e4dec701590266727d665f4d46dae0ae26408d7540c82ce07a8838"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca15192fd214d5481a05b4167238c46fbe9b1fda0d1c69c4043e48f21c599f2f"></a>

## direct_connect_disabled — direct_connect_disabled / ff05d11ec906 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- direct_connect_disabled

<a id="canonical-f847caf11fdd93d20920f7737442df073adfab1e70aa7fe26bc02898d8970fc6"></a>

Type: `["object", {}]`. Optional.

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

- [direct_connect_disabled](resources--aws_vpc_site--reference--group-002.md#canonical-f847caf11fdd93d20920f7737442df073adfab1e70aa7fe26bc02898d8970fc6)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-e3f04534e6c45f7507bd9c708a1e38b4496d0bc44167320cb22145f214117f81)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-f209eb63cc8a411874e064df721409acd5caf5601b01bf4bc66a44a3062759ea)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
direct_connect_disabled = {}
```

<a id="canonical-1da3d6e676f5efdb57f236159add75539828b459023dde952c066a7c861a033a"></a>

## Direct properties — direct_connect_disabled / ff05d11ec906 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-612e9814910decd871c0052b985bdef59b1b799b28368eebbd51b48f29341ff7"></a>

## Next pages — direct_connect_disabled / ff05d11ec906 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f58a47b9c8b4846361cf3180f41451483e37525be5867a6a02588471ad483c00"></a>

## direct_connect_enabled — direct_connect_enabled / 5f5b8eddd01c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- direct_connect_enabled

<a id="canonical-e3f04534e6c45f7507bd9c708a1e38b4496d0bc44167320cb22145f214117f81"></a>

Type: `"object"`. single nested block, Optional.

Direct Connect Configuration. Direct Connect Configuration.

Upstream description:

Direct Connect Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_asn",
    "custom_asn"),
  validators.ConflictingObjectAttributes("hosted_vifs",
    "standard_vifs")}
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
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-vif_choice": "[\"hosted_vifs\",\"standard_vifs\"]"
}
```

Terraform syntax:

```terraform
direct_connect_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-76fc340337b9906853cb5e0fb44e3e6469df6538c2a01856569ae9e0acbb74d1"></a>

## Direct properties — direct_connect_enabled / 5f5b8eddd01c / 3

- [auto_asn](resources--aws_vpc_site--reference--group-002.md#canonical-f13a5fa5be55afef3a52cbf98675db3bee488505afc538d16f35896ae92500a0): complete subsection reference.

<a id="canonical-04dd771165c335d4e8b254472bb55facaca39b8e6c1f850efeef1ca64f655658"></a>

<a id="canonical-0a0c6ac3340ec55ca317791552a36022e1bfb523a5695a7f2738625c97df7721"></a>

## custom_asn property — direct_connect_enabled / 5f5b8eddd01c / 4

Type: `"number"`. Optional.

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Upstream description:

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

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

- [hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8): complete subsection reference.

- [standard_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-e82a63d6388e2a641d004f096783b314f4d0d1713b91ddecf92b7fbdf3b5700a): complete subsection reference.

<a id="canonical-9ffc15c7170c0d40b3937531120e820201c0e6f12098e06e0600489bf163a8b8"></a>

## Next pages — direct_connect_enabled / 5f5b8eddd01c / 5

- [direct_connect_enabled.auto_asn](resources--aws_vpc_site--reference--group-002.md#canonical-f13a5fa5be55afef3a52cbf98675db3bee488505afc538d16f35896ae92500a0)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8)
- [direct_connect_enabled.standard_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-e82a63d6388e2a641d004f096783b314f4d0d1713b91ddecf92b7fbdf3b5700a)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-f13a5fa5be55afef3a52cbf98675db3bee488505afc538d16f35896ae92500a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c85ff90cedff66629694e1d5296fa577443c0fee6921a6defffb18e6ac9587dc"></a>

## direct_connect_enabled.auto_asn — direct_connect_enabled.auto_asn / 61e487da6159 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- direct_connect_enabled.auto_asn

<a id="canonical-fd773522f917ce202475ab2522843d013941f9d0ae37272456339fda758ca735"></a>

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
auto_asn = {}
```

<a id="canonical-a99dc6eeeab0bba312dc1f080e6f11108ada222c003ed8cd6d7b0c564e0aff4b"></a>

## Direct properties — direct_connect_enabled.auto_asn / 61e487da6159 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eee0755ecf84d04be04156e88ab9a798bf460a178d8da8432fcbc947cddb8434"></a>

## Next pages — direct_connect_enabled.auto_asn / 61e487da6159 / 4

- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55e3a14c6db86cbe6c233b14b82635d4e9a517bf1283684b742ce680b8e5f421"></a>

## direct_connect_enabled.hosted_vifs — direct_connect_enabled.hosted_vifs / 741d57e7314c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- direct_connect_enabled.hosted_vifs

<a id="canonical-5136bb3655ebd77b20c235dc359d9f098aa9d6a17ab77c19f0b8ccc3830fcfcc"></a>

Type: `"object"`. single nested block, Optional.

AWS Direct Connect Hosted VIF Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_registration_over_direct_connect",
    "site_registration_over_internet")}
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
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

Terraform syntax:

```terraform
hosted_vifs {
  # Configure direct properties listed below.
}
```

<a id="canonical-15ef465ca1469d4b0df11c4b010e5e9ce7240f95fe8e19c13203883724fa68a9"></a>

## Direct properties — direct_connect_enabled.hosted_vifs / 741d57e7314c / 3

- [site_registration_over_direct_connect](resources--aws_vpc_site--reference--group-002.md#canonical-bbecf31acc9b36583a11a4ca602e108a51073e9abd859bdaff41956037960fa0): complete subsection reference.

- [site_registration_over_internet](resources--aws_vpc_site--reference--group-002.md#canonical-a840843b76e45c451b001f474ca4483caccab6c3cefd21ccf3b2cc78e1d0c267): complete subsection reference.

- [vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-b75c7f1e74e869a720dd1e3fdca55397e7233d6dca660080c75c65d9c9d64bfe): complete subsection reference.

<a id="canonical-b585c6842cc3b6e8461f095d23cd9bab160eed4ebf1375279755633cd4be4060"></a>

## Next pages — direct_connect_enabled.hosted_vifs / 741d57e7314c / 4

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_vpc_site--reference--group-002.md#canonical-bbecf31acc9b36583a11a4ca602e108a51073e9abd859bdaff41956037960fa0)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_vpc_site--reference--group-002.md#canonical-a840843b76e45c451b001f474ca4483caccab6c3cefd21ccf3b2cc78e1d0c267)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-b75c7f1e74e869a720dd1e3fdca55397e7233d6dca660080c75c65d9c9d64bfe)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-bbecf31acc9b36583a11a4ca602e108a51073e9abd859bdaff41956037960fa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3569dc5c01835ed32970edf65f1798c5728f86ae9de16c5f36b8d5c104536417"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / eb7dba3c0e6b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-2f5e7703ea0be6222221e1b3960c1b6956521c7ccc0bf342a943a1ccdd77a652"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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
site_registration_over_direct_connect {
  # Configure direct properties listed below.
}
```

<a id="canonical-4a676ff6f2134cba073392b08e8ae6fb27b5111083d5de73379b6b944fb5720b"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / eb7dba3c0e6b / 3

<a id="canonical-1669a14dfb29df6733f772876be4e065fab5ed4d6df9fd40403d8f60e4004a13"></a>

<a id="canonical-e02e06e7122b75b690b02877094c6fc4ce7a7cabe7421d72dc82350def7e7d2d"></a>

## cloudlink_network_name property — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / eb7dba3c0e6b / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-78c81e5066cfabafe3a2a3fd9e7637627a7082db297b6dc416f4e94f47ad32af"></a>

## Next pages — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / eb7dba3c0e6b / 5

- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a840843b76e45c451b001f474ca4483caccab6c3cefd21ccf3b2cc78e1d0c267"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12b14a0f5a01b7fcf1683740ec7e24b9059897ccd9a0459b99b3209d285ccab1"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_internet — direct_connect_enabled.hosted_vifs.site_registration_over_internet / ec53c0668647 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-99ad18e46cc0b1697dc0fbb2ee4cb36ed3c89a9c3bbdb3895d8118a930be2e5c"></a>

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
site_registration_over_internet = {}
```

<a id="canonical-d8ddcb07d17fe02023fb9eb2ab4b10fa621c3dbbd2ef495479dc7fd695b3af0c"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.site_registration_over_internet / ec53c0668647 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c96d36fd0ed44d3e49e3008ef2552ed4d87e55c42c7945888a4d3cbc27d5165"></a>

## Next pages — direct_connect_enabled.hosted_vifs.site_registration_over_internet / ec53c0668647 / 4

- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-b75c7f1e74e869a720dd1e3fdca55397e7233d6dca660080c75c65d9c9d64bfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4db7927674562112bca4b93ee50da397b48021c81d35babe1c69fe12e57f1b8"></a>

## direct_connect_enabled.hosted_vifs.vif_list — direct_connect_enabled.hosted_vifs.vif_list / e0b1dc56b4bb / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-bf8529e422458d385edb52fd5338b95eb687124a211f74141e4e8c565fefd9a8"></a>

Type: `"object"`. list nested block, Optional.

List of Hosted VIF Config. List of Hosted VIF Config.

Upstream description:

List of Hosted VIF Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("vif_id"),
  validators.ConflictingListObjectAttributes("other_region",
    "same_as_site_region")}
```

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

Terraform syntax:

```terraform
vif_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a0d4deba1109036f67c27adfd50552b292534c7a6fa02390782710a3dc4e4a7"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.vif_list / e0b1dc56b4bb / 3

<a id="canonical-5c0854eaff3d40a8586fe569264c1baf2bb4487a5f136beafb5a8196ccdb1f8f"></a>

<a id="canonical-9499cbf9e777b91a64ae3e3829ec9fafd0443182b5916919b5f904807d33aeb5"></a>

## other_region property — direct_connect_enabled.hosted_vifs.vif_list / e0b1dc56b4bb / 4

Type: `"string"`. Optional.

\[Enum:
af-south-1|ap-east-1|ap-northeast-1|ap-northeast-2|ap-south-1|ap-southeast-1|ap-southeast-2|ap-southeast-3|ca-central-1|eu-central-1|eu-north-1|eu-south-1|eu-west-1|eu-west-2|eu-west-3|me-south-1|sa-east-1|us-east-1|us-east-2|us-west-1|us-west-2\]
Exclusive with \[same\_as\_site\_region\] Other Region. Possible values are \`af-south-1\`,
\`ap-east-1\`, \`ap-northeast-1\`, \`ap-northeast-2\`, \`ap-south-1\`, \`ap-southeast-1\`,
\`ap-southeast-2\`, \`ap-southeast-3\`, \`ca-central-1\`, \`eu-central-1\`, \`eu-north-1\`,
\`eu-south-1\`, \`eu-west-1\`, \`eu-west-2\`, \`eu-west-3\`, \`me-south-1\`, \`sa-east-1\`,
\`us-east-1\`, \`us-east-2\`, \`us-west-1\`, \`us-west-2\`.

Upstream description:

Exclusive with \[same\_as\_site\_region\] Other Region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("af-south-1",
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
    "us-west-2"),
}
```

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

- [same_as_site_region](resources--aws_vpc_site--reference--group-002.md#canonical-b0861291b8a613e43e0ecbdf9725b08bf10ce6f79000763c3489a9adfcb93d06): complete subsection reference.

<a id="canonical-071c6604bf5a556c5455833357811d71373b575678d68534081624aaa2360121"></a>

<a id="canonical-acf8bc451077d8c36071dfda8dcdff46b71d66ac615950f79b418e0c622f1490"></a>

## vif_id property — direct_connect_enabled.hosted_vifs.vif_list / e0b1dc56b4bb / 5

Type: `"string"`. Optional.

AWS Direct Connect VIF ID that needs to be connected to the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

<a id="canonical-3a66c283f0c6e48f55e692500ecd99b6f592747fb96ffa9e74aa168c6f6bacb3"></a>

## Next pages — direct_connect_enabled.hosted_vifs.vif_list / e0b1dc56b4bb / 6

- [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_vpc_site--reference--group-002.md#canonical-b0861291b8a613e43e0ecbdf9725b08bf10ce6f79000763c3489a9adfcb93d06)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-b0861291b8a613e43e0ecbdf9725b08bf10ce6f79000763c3489a9adfcb93d06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ef2f72a0429e38b08e84be04703f3f4f07b93dd2dd36f9a46585326d3768b52"></a>

## direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / e69aff76a8a8 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5a6fbd3d9fc106eb9097b84eaabc5a32700e158d00a4f257e7a9110a19d7a0b8)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-b75c7f1e74e869a720dd1e3fdca55397e7233d6dca660080c75c65d9c9d64bfe)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-6e94970acc1c3eb03e5603645cc9a2a72f7a3d76add735040d556d469d34cd7d"></a>

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
same_as_site_region = {}
```

<a id="canonical-cbb7971755f5ef7f9642609922d34b6ac02dfa5f37b8066af51036805de8b2b8"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / e69aff76a8a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5e150c2c9b55d7d29a7f19eb5e7c5ce1b65140fef478acc9c93386b38c2acc4f"></a>

## Next pages — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / e69aff76a8a8 / 4

- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-b75c7f1e74e869a720dd1e3fdca55397e7233d6dca660080c75c65d9c9d64bfe)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e82a63d6388e2a641d004f096783b314f4d0d1713b91ddecf92b7fbdf3b5700a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16d2b487a95beb6366d41cd3eba5bdd90ef079e4052574c1e7504702c4871c91"></a>

## direct_connect_enabled.standard_vifs — direct_connect_enabled.standard_vifs / bf88397923be / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- direct_connect_enabled.standard_vifs

<a id="canonical-df7e3a71460cdacbbb9ead82212834343803f9b79c808011b8c50f20d6d8488a"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
standard_vifs = {}
```

<a id="canonical-6874e1a2fdd798afadaac443aaac7d935d6182a3e0e79e94d3f99358473763c7"></a>

## Direct properties — direct_connect_enabled.standard_vifs / bf88397923be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7492c8949807f05054e24d2fb11cdbe6334f87eba368c143d42074a508b0d62e"></a>

## Next pages — direct_connect_enabled.standard_vifs / bf88397923be / 4

- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-6223edc061be8cc81e48316841a7cf498750d8a0c00acc2bf767f5d931febc4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fe9a905726bdff2e14db8cf909545026c5f2ab2386ce1c5dd19a33ad7abbe5b"></a>

## disable_encryption — disable_encryption / 4196378cf7f2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- disable_encryption

<a id="canonical-3bb62a0311a7a7d3e6fefedd77c0b189c55466526ef345b933577675a7f16b9e"></a>

Type: `["object", {}]`. Optional.

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

- [disable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-3bb62a0311a7a7d3e6fefedd77c0b189c55466526ef345b933577675a7f16b9e)
- [enable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-8cd5cbc5f3879c41d2173dfceebcde1ec00a37777e55755ee4a6c8a0b473d868)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_encryption = {}
```

<a id="canonical-731d5275650aa1dad555be98346fa94cd5c7ec0d5ab243815eb4c6ebdd2ea76a"></a>

## Direct properties — disable_encryption / 4196378cf7f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-adf410a13543d8ab3639958147c41dc025e39b5f0c1523991a9bdcc56c9e115e"></a>

## Next pages — disable_encryption / 4196378cf7f2 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-4068adc6014448bc07d7ab63746cf6ce80a3fc063b63c7e93e64491beecbb31c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bc5e940d14fedd4eefc471da83d07270fa8055ccbd66e67bfd0dd5146877da8"></a>

## disable_internet_vip — disable_internet_vip / f16266ea8b62 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- disable_internet_vip

<a id="canonical-74bf589be459697a74a7431450d870406155041507e6aa4c4444f0b53716d00b"></a>

Type: `["object", {}]`. Optional.

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

- [disable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-74bf589be459697a74a7431450d870406155041507e6aa4c4444f0b53716d00b)
- [enable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-f9e38a102bd0e36baa826ae5153a21b93d95c091608aa478d8b8a0ecff2049e3)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_internet_vip = {}
```

<a id="canonical-e7b4eaf5f943479550cf90cbe08ab2060e2ab810a9dec87e47866969b4eae5b0"></a>

## Direct properties — disable_internet_vip / f16266ea8b62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6333230f02b081ae8ef7c82fa4bb5f3588a1aeb229c0f2329ddbb4d4ec06355e"></a>

## Next pages — disable_internet_vip / f16266ea8b62 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-662cb1c766ab5942d9d12ddd323a9ab6cc30eb0ff961e565556e21ceab133049"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5277507a24c3c01bf00ed3720d21e3fca313e353ef912bdc7a299cf4ceb84d5c"></a>

## egress_gateway_default — egress_gateway_default / 8bbb56236023 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- egress_gateway_default

<a id="canonical-2b78c4d1a5f62696fbb388133b7f1e001e499e40c66ef19deb9a329d736e90dc"></a>

Type: `["object", {}]`. Optional.

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

- [egress_gateway_default](resources--aws_vpc_site--reference--group-002.md#canonical-2b78c4d1a5f62696fbb388133b7f1e001e499e40c66ef19deb9a329d736e90dc)
- [egress_nat_gw](resources--aws_vpc_site--reference--group-002.md#canonical-b44c33f0c52cfc87b4bf392a8aa7159ccd82fcabe7f1c773379d84edfe584f08)
- [egress_virtual_private_gateway](resources--aws_vpc_site--reference--group-002.md#canonical-9a3715518fb6c71afdfc766cf2841c92a8348809446875906f74e91fe63f629e)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
egress_gateway_default = {}
```

<a id="canonical-6d8a3ffed7060d3c96f480a56cb2eb270a992abf5b20edd0932ea09d34bc6531"></a>

## Direct properties — egress_gateway_default / 8bbb56236023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ea0609727421422e77e2054abf1dd3cac8eeb20994afee44b80a09f4315a5016"></a>

## Next pages — egress_gateway_default / 8bbb56236023 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-9efdb256997148a2eef145984b59ba9a87faa0ba2632b4025d7d8dc6d5d78c74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70f95b7c9af29fc1d5e0a569393696129b8f477a0461725244b48fbb2a36222e"></a>

## egress_nat_gw — egress_nat_gw / 2d22e3bbf5e5 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- egress_nat_gw

<a id="canonical-b44c33f0c52cfc87b4bf392a8aa7159ccd82fcabe7f1c773379d84edfe584f08"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
egress_nat_gw {
  # Configure direct properties listed below.
}
```

<a id="canonical-6fd9deff7d9f3d9709c2b5538349aa26aebce182fdb827a8e413d3ad3c593d1e"></a>

## Direct properties — egress_nat_gw / 2d22e3bbf5e5 / 3

<a id="canonical-2bd12ca9b7eac7bddd41ec8f458300144f610775715470b9dc7d5e1c9773baa6"></a>

<a id="canonical-2c72db3661379265d2f98b2649d3021a947777c6be3c2460be178ce055354791"></a>

## nat_gw_id property — egress_nat_gw / 2d22e3bbf5e5 / 4

Type: `"string"`. Optional.

Existing NAT Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(21),
}
```

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

<a id="canonical-bce80280f01fccd2101489a1df585b7601ab76b812a36b91281e6a0bae7bf870"></a>

## Next pages — egress_nat_gw / 2d22e3bbf5e5 / 5

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-175073f7a4cdbf9b1f82a592583d94135e809165c4307b252469a6fb34e853f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7fd5127bf791fc2d1971f0a002f921b151d6c0ea288c3c7d9535154b4f8d2b8"></a>

## egress_virtual_private_gateway — egress_virtual_private_gateway / 147d969fcbf3 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- egress_virtual_private_gateway

<a id="canonical-9a3715518fb6c71afdfc766cf2841c92a8348809446875906f74e91fe63f629e"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
egress_virtual_private_gateway {
  # Configure direct properties listed below.
}
```

<a id="canonical-50cd691bfde69ef1010bdc178dc68e3a46ba24034061b884a30b2ed52803483c"></a>

## Direct properties — egress_virtual_private_gateway / 147d969fcbf3 / 3

<a id="canonical-7b3a5500401f0a42abd87af588b7750e9de12ca834aa25b5863a1dc17201b226"></a>

<a id="canonical-27243d9292e1bc7966e82a2bff563b5adfeaf9bb8ab3e6e6de2fe011bd5b50e7"></a>

## vgw_id property — egress_virtual_private_gateway / 147d969fcbf3 / 4

Type: `"string"`. Optional.

Existing Virtual Private Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(21),
}
```

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

<a id="canonical-4957b1b844690e8947bddfcfaf482522505674f9d4b06acc707fa1f6890ed41c"></a>

## Next pages — egress_virtual_private_gateway / 147d969fcbf3 / 5

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-da48ecd52f76e8a8414bfa8792082b40c74f6ff78395c63e44830dbee422729a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29bde25e6fb5a45126a95fba50cfb0da129252cafc8ce3cecf517a3d1197bf38"></a>

## enable_encryption — enable_encryption / 253b0f8a6768 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- enable_encryption

<a id="canonical-8cd5cbc5f3879c41d2173dfceebcde1ec00a37777e55755ee4a6c8a0b473d868"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("kms_key_id")}
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
enable_encryption {
  # Configure direct properties listed below.
}
```

<a id="canonical-ad354581564a9d4be1be9713f100b177e79bda0da7e7ae5bccdeddc2ee2f5e6d"></a>

## Direct properties — enable_encryption / 253b0f8a6768 / 3

<a id="canonical-eb724f49d998087076908938e93cd781f280488f23ae27dfaf15f2101780154d"></a>

<a id="canonical-23e02923d1a060389ffed136b8faebe7206df79aa47bf9a7a1a3104184979dcb"></a>

## kms_key_id property — enable_encryption / 253b0f8a6768 / 4

Type: `"string"`. Optional.

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

<a id="canonical-309c82af8bab22ee0cb6330b63bdbb4677f2b8d564e103c59be119a185ca46aa"></a>

## Next pages — enable_encryption / 253b0f8a6768 / 5

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-f26a15bffe5278b54999ccbe58e244df51238820933712ed60f8437927c3253a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31c756179a92da9bb99166875d7c857a0c1478562808a175af5356d09c0d8d33"></a>

## enable_internet_vip — enable_internet_vip / 9b41dda09d61 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- enable_internet_vip

<a id="canonical-f9e38a102bd0e36baa826ae5153a21b93d95c091608aa478d8b8a0ecff2049e3"></a>

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
enable_internet_vip = {}
```

<a id="canonical-43f5c9c44fe012596f10b670cfa2ad2565b38e5c7ba29414712961b39f5640fe"></a>

## Direct properties — enable_internet_vip / 9b41dda09d61 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7140dc72ea6c0eab671086e1c673759107d409a3f7f0ba73b6f498247e6bee02"></a>

## Next pages — enable_internet_vip / 9b41dda09d61 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-2f43fbe97c023e023323ddbe139b16cf5e26356d9691ae04b777fca69e39128c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f981c1133bd3942ee65829acb6e1c7ab37fc2f31521f2b4d0edcd1fbe89fbe03"></a>

## f5_orchestrated_routing — f5_orchestrated_routing / c94318ca6de7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- f5_orchestrated_routing

<a id="canonical-6012e1886fc593a833a34f70ca376543a6289b992b21d7587c409f3707b4954b"></a>

Type: `["object", {}]`. Optional.

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

- [f5_orchestrated_routing](resources--aws_vpc_site--reference--group-002.md#canonical-6012e1886fc593a833a34f70ca376543a6289b992b21d7587c409f3707b4954b)
- [manual_routing](resources--aws_vpc_site--reference--group-004.md#canonical-c864d3ad531ccb13cd7fc450b170b6ea6f1c9458d3bbd77e0a81ece6574449a6)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
f5_orchestrated_routing = {}
```

<a id="canonical-0ff4321ac0caa1d233514395cbb72a02eee36456b6ad9df3d52df12c8893887a"></a>

## Direct properties — f5_orchestrated_routing / c94318ca6de7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07f85adafd99c907854158436d6189edadc1410422f0a6fa799df9edf3ecaf70"></a>

## Next pages — f5_orchestrated_routing / c94318ca6de7 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1a305484da30c6bdf1b10ab03e7ffaa9215fd2dd126b99629e0535a72ce56821"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7da5a2239dc8974e549ab8b0d9792fb7107faf96a057c732e6665b6b13c521ca"></a>

## f5xc_security_group — f5xc_security_group / 02cdeae3d3c4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- f5xc_security_group

<a id="canonical-d5baaebb13185ddb9f08cf03b5a2186aeaf97eabc9c0cfc2c226fe71a81f2fef"></a>

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
f5xc_security_group = {}
```

<a id="canonical-b199594c6b9bc952e76b5e289db5bdd6f9a5fa7dc40bc68b24626eb6bde410c5"></a>

## Direct properties — f5xc_security_group / 02cdeae3d3c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-efb7f67b88fe4844d12ba3f36b98c3d2925d9f61e44fdcfa60b8b1c0e5191d7a"></a>

## Next pages — f5xc_security_group / 02cdeae3d3c4 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df84e2319539956391f5b4732694a93f968149fbafe6cdbd3051fe2daf3e6c1e"></a>

## ingress_egress_gw — ingress_egress_gw / 3e26249ca519 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- ingress_egress_gw

<a id="canonical-d713a098ee4dcbcdee803063f8fb7e3d8ef97c713bc4cac861138c989fc8595a"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface AWS ingress/egress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_certified_hw",
    "az_nodes"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-d713a098ee4dcbcdee803063f8fb7e3d8ef97c713bc4cac861138c989fc8595a)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-673064a09e6a2668846219b579341e6df28ae8d7a84d2c33fa7479cf571b751e)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-4141c9dfecf713cd3dfd532024e078d75bda39e24b2abdbeace7b643dc73f1d6)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ingress_egress_gw {
  # Configure direct properties listed below.
}
```

<a id="canonical-341e26391c8651f1709b92c993b16f4e609a3cdcacc5870134f2245cc976c580"></a>

## Direct properties — ingress_egress_gw / 3e26249ca519 / 3

- [active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-4f64e92c686bb3bbfbe60aa25480a9a05ecbd06f8acf22f69461c62932d971e2): complete subsection reference.

- [active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-8e10b9343f6235f693db55aef98d5ab7ad91ef8cb9cc636e764c1d1d7c6ae0e7): complete subsection reference.

- [active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-9e49e2206b18bc33a3cab0a8cf388bab916fbbfbb2acd042328c4e7497c19e70): complete subsection reference.

- [allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43): complete subsection reference.

- [allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49): complete subsection reference.

<a id="canonical-b6fec99e2cbb49efd440d62abdeaed8d079e2b08590524b6e54ee49b48dce9a4"></a>

<a id="canonical-dfda90ac3c1dba14396e7e8e336d6c3c55360039889bedf56f3fd87d51567d1d"></a>

## aws_certified_hw property — ingress_egress_gw / 3e26249ca519 / 4

Type: `"string"`. Optional.

\[Enum: aws-byol-multi-nic-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The
only possible value is \`aws-byol-multi-nic-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("aws-byol-multi-nic-voltmesh"),
}
```

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

- [az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922): complete subsection reference.

- [dc_cluster_group_inside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-6250e0449953556e0f52164a22f9a506c4cb1a29f295372fbdfa57b9d18dda25): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-0514f74fd9de4c088d7e716d36bba6e3a81e7172ecc0b55bf0868f6e654185cd): complete subsection reference.

- [forward_proxy_allow_all](resources--aws_vpc_site--reference--group-002.md#canonical-bd1bb7489a9cdcacd2d5f27ba6dcb6ac8e03cfd472684d709f73916056e114e1): complete subsection reference.

- [global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9): complete subsection reference.

- [inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e): complete subsection reference.

- [no_dc_cluster_group](resources--aws_vpc_site--reference--group-003.md#canonical-925db483a8d11cc7301a046cbd60e34288e2be0b4b2956b9c174f80e93f2cb80): complete subsection reference.

- [no_forward_proxy](resources--aws_vpc_site--reference--group-003.md#canonical-304079dda7dd5d59b838510f594d8679fe4d9b033527672f60df99dfe6b3e501): complete subsection reference.

- [no_global_network](resources--aws_vpc_site--reference--group-003.md#canonical-b373260e05977dec450fb878a00d0aac7028b0bdcc82fcffee8f6ede03a51058): complete subsection reference.

- [no_inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-7f08c3dd89dc58192ae4aabb1edee5c4690eab8217a210fb8632de2eb1cf4aa5): complete subsection reference.

- [no_network_policy](resources--aws_vpc_site--reference--group-003.md#canonical-3f8d154444742593f7621c73fc280a8494f1eda2ed613ddca74c3f4fa208a26e): complete subsection reference.

- [no_outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-e6a7038350d6bdb13a48ed189a70488534b464456f1cd2c9ec10ea58a9a3dd51): complete subsection reference.

- [outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82): complete subsection reference.

- [performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61): complete subsection reference.

- [sm_connection_public_ip](resources--aws_vpc_site--reference--group-003.md#canonical-8245ef1b9ec742f19a0f2e280e11739d80e7622dbdf423a8363c438ac16c8992): complete subsection reference.

- [sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-003.md#canonical-8a09e1f2df9a35b9e17192f57980319cd9e580ea183832a4d46d0012d6e36149): complete subsection reference.

<a id="canonical-ee0eaddef9586595985f4d16ff07fca9844b5b676a2af66ad6caeb1bf97b2d46"></a>

## Next pages — ingress_egress_gw / 3e26249ca519 / 5

- [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-4f64e92c686bb3bbfbe60aa25480a9a05ecbd06f8acf22f69461c62932d971e2)
- [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-8e10b9343f6235f693db55aef98d5ab7ad91ef8cb9cc636e764c1d1d7c6ae0e7)
- [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-9e49e2206b18bc33a3cab0a8cf388bab916fbbfbb2acd042328c4e7497c19e70)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- [ingress_egress_gw.dc_cluster_group_inside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-6250e0449953556e0f52164a22f9a506c4cb1a29f295372fbdfa57b9d18dda25)
- [ingress_egress_gw.dc_cluster_group_outside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-0514f74fd9de4c088d7e716d36bba6e3a81e7172ecc0b55bf0868f6e654185cd)
- [ingress_egress_gw.forward_proxy_allow_all](resources--aws_vpc_site--reference--group-002.md#canonical-bd1bb7489a9cdcacd2d5f27ba6dcb6ac8e03cfd472684d709f73916056e114e1)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.no_dc_cluster_group](resources--aws_vpc_site--reference--group-003.md#canonical-925db483a8d11cc7301a046cbd60e34288e2be0b4b2956b9c174f80e93f2cb80)
- [ingress_egress_gw.no_forward_proxy](resources--aws_vpc_site--reference--group-003.md#canonical-304079dda7dd5d59b838510f594d8679fe4d9b033527672f60df99dfe6b3e501)
- [ingress_egress_gw.no_global_network](resources--aws_vpc_site--reference--group-003.md#canonical-b373260e05977dec450fb878a00d0aac7028b0bdcc82fcffee8f6ede03a51058)
- [ingress_egress_gw.no_inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-7f08c3dd89dc58192ae4aabb1edee5c4690eab8217a210fb8632de2eb1cf4aa5)
- [ingress_egress_gw.no_network_policy](resources--aws_vpc_site--reference--group-003.md#canonical-3f8d154444742593f7621c73fc280a8494f1eda2ed613ddca74c3f4fa208a26e)
- [ingress_egress_gw.no_outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-e6a7038350d6bdb13a48ed189a70488534b464456f1cd2c9ec10ea58a9a3dd51)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- [ingress_egress_gw.sm_connection_public_ip](resources--aws_vpc_site--reference--group-003.md#canonical-8245ef1b9ec742f19a0f2e280e11739d80e7622dbdf423a8363c438ac16c8992)
- [ingress_egress_gw.sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-003.md#canonical-8a09e1f2df9a35b9e17192f57980319cd9e580ea183832a4d46d0012d6e36149)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-4f64e92c686bb3bbfbe60aa25480a9a05ecbd06f8acf22f69461c62932d971e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-282c262730235d158dbea095503b2018c1c0e844cb2ef3de79084d4dfac6c606"></a>

## ingress_egress_gw.active_enhanced_firewall_policies — ingress_egress_gw.active_enhanced_firewall_policies / ae5ab440837a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.active_enhanced_firewall_policies

<a id="canonical-3239a16f6add5e06e358345b4282167501cf54c45a9525099d11af18a0b279ca"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1fc3a2039df8ad6ee5c592d518b5d44c44bd0629a00af4ce8f672fc43b16f06f"></a>

## Direct properties — ingress_egress_gw.active_enhanced_firewall_policies / ae5ab440837a / 3

- [enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-bdeb5b48fa221c37cc00cb7f5530ca059f77bf1161b2af9eb88d1a87baeaee15): complete subsection reference.

<a id="canonical-5183b6415f8276c3a8e7685016b6242a396fac1c5a0dc6dc5e71745ad0c00f50"></a>

## Next pages — ingress_egress_gw.active_enhanced_firewall_policies / ae5ab440837a / 4

- [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-bdeb5b48fa221c37cc00cb7f5530ca059f77bf1161b2af9eb88d1a87baeaee15)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-bdeb5b48fa221c37cc00cb7f5530ca059f77bf1161b2af9eb88d1a87baeaee15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d6cb343b30f549a34459a1b821b8a089f8fd78562a98e05c9ae0c8d6c7f8401"></a>

## ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 5ce494eb98a9 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-4f64e92c686bb3bbfbe60aa25480a9a05ecbd06f8acf22f69461c62932d971e2)
- ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-60e1a76442302d65cbf7fb30a8ce21ecaf6084bd5e44e0b5c529861e25f227ce"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-82708f8b12b71977b88cf0f2594a80b16d99f7dd78af268df5afa557bbc4170a"></a>

## Direct properties — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 5ce494eb98a9 / 3

<a id="canonical-8d63eb0b9afd40248dc62bb716283c78b98f1bdd3928ff6f62f307c7f55ee985"></a>

<a id="canonical-ec8564faf5d7772fa3ecf97f88d5962ec662bc26bcf9bbc71f16212ad6d3c889"></a>

## name property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 5ce494eb98a9 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1b8810b53cfec9a5a1c1d0693ca4df0bfbd1ea8d1297037dc9bff336f9487a5b"></a>

<a id="canonical-461172d3a4d995c844118014ed47479e90138fcb87893e15514d45e14d776562"></a>

## namespace property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 5ce494eb98a9 / 5

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
}
```

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

<a id="canonical-31e88374ff5953c806f3d41baa81f40f13a497c9d3ca1b3c3e7e2afe7867c8c6"></a>

<a id="canonical-43baa6d413d41d91812d81962ca3c23b061b81d676d06182068fbb7803e6dc80"></a>

## tenant property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 5ce494eb98a9 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2e23c135d6f881f58fe5c9f770230f493af23efabbc40f6964294744aca5f09e"></a>

## Next pages — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / 5ce494eb98a9 / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-4f64e92c686bb3bbfbe60aa25480a9a05ecbd06f8acf22f69461c62932d971e2)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-8e10b9343f6235f693db55aef98d5ab7ad91ef8cb9cc636e764c1d1d7c6ae0e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db385fcced7f9a84017b16be4fbd048f898a9bbe453a1413f90a57a6fff1c4be"></a>

## ingress_egress_gw.active_forward_proxy_policies — ingress_egress_gw.active_forward_proxy_policies / 26703a6f97cd / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.active_forward_proxy_policies

<a id="canonical-1f26b244ff524d37ca8adb44433a02ba1a461a7765bf812ebdf4f4bb72dbc5f1"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-16692165c77084eb8097332e4b748483d44d3b61c44ccce044c7150809eb40ae"></a>

## Direct properties — ingress_egress_gw.active_forward_proxy_policies / 26703a6f97cd / 3

- [forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-28d5f6183458de3357dce46d7509f3bbb125b1393a091f7d459af65872b10308): complete subsection reference.

<a id="canonical-8e814f111c270f346213101f3c9ebc999429038d05a4be56a1b3288c89fe6e39"></a>

## Next pages — ingress_egress_gw.active_forward_proxy_policies / 26703a6f97cd / 4

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-28d5f6183458de3357dce46d7509f3bbb125b1393a091f7d459af65872b10308)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-28d5f6183458de3357dce46d7509f3bbb125b1393a091f7d459af65872b10308"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f150974835af3129c7e95a50f6d656639065ef39915ad2b7bce6ab8b037768d3"></a>

## ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 96b8012b40d6 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-8e10b9343f6235f693db55aef98d5ab7ad91ef8cb9cc636e764c1d1d7c6ae0e7)
- ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-dfeccd1538d946594d84c36aeecb62a0cb8c8efbfcdf2e5871c8072edbaaff1d"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc8154bc411dce649af448e63a1e9fbf53b3c58d3e5baa5f869e31a60a093eea"></a>

## Direct properties — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 96b8012b40d6 / 3

<a id="canonical-f4063e4aa0a9070117e6ec193c9cb8927c5f221311b191acc9d68ac348eb7f4f"></a>

<a id="canonical-6198c2ca361178be9c023c49e08e8d88a63265489f066a8f9cbd039c33a3cc5a"></a>

## name property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 96b8012b40d6 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-e335b990e410395aa0c8cf610061607ffe147d5293b86e8ffce2d997823159bf"></a>

<a id="canonical-48df5ad4a9e3d7dda7792227928f7fd9e3268ee4173408c88ceef8a292684013"></a>

## namespace property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 96b8012b40d6 / 5

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
}
```

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

<a id="canonical-fe15a7fb67884c940a499a007ebe5b775af25c25f856f61e9ea69feda3ec9abc"></a>

<a id="canonical-3673a542b236ebfc0cef5ae8f64da7e8081f861d1901cf2e561087b5aaecadc7"></a>

## tenant property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 96b8012b40d6 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-6e4f205e26fa37b2800c265d9734a4efee82d712a2bac05c7860f6f00937260a"></a>

## Next pages — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 96b8012b40d6 / 7

- [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-8e10b9343f6235f693db55aef98d5ab7ad91ef8cb9cc636e764c1d1d7c6ae0e7)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-9e49e2206b18bc33a3cab0a8cf388bab916fbbfbb2acd042328c4e7497c19e70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6db2cbe6f2d45ed391b532ae2254b5f38fda8dba7b6fffd39f71ab3cbe692ec0"></a>

## ingress_egress_gw.active_network_policies — ingress_egress_gw.active_network_policies / af796b7e9dfb / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.active_network_policies

<a id="canonical-bbc2cc1cd2580a81127e40bb960c68f1884ac396b2c7ad20ec7a94d18457f05b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-b4cef4dd116bafb17dde05acc4f5302e980f6e5675373f9c93f10901250bd6a0"></a>

## Direct properties — ingress_egress_gw.active_network_policies / af796b7e9dfb / 3

- [network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-89c623f005bd4cd0a6e8e2d61a0d6941007134e39276e91605320d848f0c1517): complete subsection reference.

<a id="canonical-b703f41ded581c3778b9b1cfe8c93c00fa52e10b9b54a1347c43d9fee23a3500"></a>

## Next pages — ingress_egress_gw.active_network_policies / af796b7e9dfb / 4

- [ingress_egress_gw.active_network_policies.network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-89c623f005bd4cd0a6e8e2d61a0d6941007134e39276e91605320d848f0c1517)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-89c623f005bd4cd0a6e8e2d61a0d6941007134e39276e91605320d848f0c1517"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a78610192c0226bda98df240e8d1b7d0375018f446bbe19cbdb3b9915b8515db"></a>

## ingress_egress_gw.active_network_policies.network_policies — ingress_egress_gw.active_network_policies.network_policies / 863d5546a21f / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-9e49e2206b18bc33a3cab0a8cf388bab916fbbfbb2acd042328c4e7497c19e70)
- ingress_egress_gw.active_network_policies.network_policies

<a id="canonical-cc2258bd1d73f083126572fa371d6d999fe8eee88738b30f8fbccdb955fb7f1f"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-9cbb775896ed0cd1752162c951a8cef163af136fdff10d41d7d4457963229788"></a>

## Direct properties — ingress_egress_gw.active_network_policies.network_policies / 863d5546a21f / 3

<a id="canonical-c2256aca2f4af87695789fe2c8e6bb246c728b2daed0b4e2cf1e6a4aade49229"></a>

<a id="canonical-53754834c27400755200122e61d3eb4127ea30d2cdfa2ac24970b985b00a5fa2"></a>

## name property — ingress_egress_gw.active_network_policies.network_policies / 863d5546a21f / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-df4612a6cc220269e7889153de554e4dcdee19c3e4c5fb870fa925d78754f6fa"></a>

<a id="canonical-e59dd34f147f544e0c80eae5bd20c6d9833eaafe36df8e621ecbda57a305d2a8"></a>

## namespace property — ingress_egress_gw.active_network_policies.network_policies / 863d5546a21f / 5

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
}
```

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

<a id="canonical-9e77eae0ed4fba5eb45c8dae8555bb33c21fc7703ad673c122a34d3bc115f4eb"></a>

<a id="canonical-77fd5a8de959c0b35ab579c6ad0129f5e312c3c6eb17775acd639f74e1edcf72"></a>

## tenant property — ingress_egress_gw.active_network_policies.network_policies / 863d5546a21f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-4f0c52a7f4e6a6e8a999d7a6a955bba59a190317390a2fa943ebf63379b6c11f"></a>

## Next pages — ingress_egress_gw.active_network_policies.network_policies / 863d5546a21f / 7

- [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-9e49e2206b18bc33a3cab0a8cf388bab916fbbfbb2acd042328c4e7497c19e70)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5929605575f2c17b333d8971d7e254913f1e9cbb8cd844aa44ef57f1bb2ccf8"></a>

## ingress_egress_gw.allowed_vip_port — ingress_egress_gw.allowed_vip_port / 038dcbae1739 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.allowed_vip_port

<a id="canonical-8dccbf736c9289beda749760cb60532d5edf37fc96240dc3b0ddbe8764615682"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c7e676030688d1199da2f288508291fec43e32d2b15971bad1b65ee2a7190da"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port / 038dcbae1739 / 3

- [custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-bd0fc8d01edac00c1eb1e7e55c7fb1fb12851174acd4c04238e5dcf795da55eb): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-e13603423676e9ec874cd77fbfb19a2af5147c9fb282130f7d4ace14f5e0d2be): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-ee80dd985ed384ffdd62c817d8d5225e44b69932d2d1de5c4e429be032e8c834): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-8f7a1ecf7db862b19ddaea4186562960a879787c1a0114b257bd722cbc371d3f): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-d945251fd3562544a0177017b256177bfc57459bd4ad3a78a3b79f6e1346be21): complete subsection reference.

<a id="canonical-4376d5e83f56cb9eec72c8c81f5314fe205f79314df78b2a4d6658cacc25a0fb"></a>

## Next pages — ingress_egress_gw.allowed_vip_port / 038dcbae1739 / 4

- [ingress_egress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-bd0fc8d01edac00c1eb1e7e55c7fb1fb12851174acd4c04238e5dcf795da55eb)
- [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-e13603423676e9ec874cd77fbfb19a2af5147c9fb282130f7d4ace14f5e0d2be)
- [ingress_egress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-ee80dd985ed384ffdd62c817d8d5225e44b69932d2d1de5c4e429be032e8c834)
- [ingress_egress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-8f7a1ecf7db862b19ddaea4186562960a879787c1a0114b257bd722cbc371d3f)
- [ingress_egress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-d945251fd3562544a0177017b256177bfc57459bd4ad3a78a3b79f6e1346be21)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-bd0fc8d01edac00c1eb1e7e55c7fb1fb12851174acd4c04238e5dcf795da55eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b1c7c5c472fc10e5bf778d2339328076670b81f818729f6d021375c79c55a80"></a>

## ingress_egress_gw.allowed_vip_port.custom_ports — ingress_egress_gw.allowed_vip_port.custom_ports / cc28a4d33d66 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- ingress_egress_gw.allowed_vip_port.custom_ports

<a id="canonical-f63f9832c0fe7c1c7e610830d444a06c58f03bcc914f0d7afd58219b74383b7b"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-8dd51cf577f1f64c618e48cdfea04023886def350fb0cc017fb8175f245f6c0a"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.custom_ports / cc28a4d33d66 / 3

<a id="canonical-0bd6657cb2877c62a67bce2326a30cc00ce555f8c0e323ecdd6b9ba5f77977d6"></a>

<a id="canonical-2a521e7bda30d1875bf73b91772156a4fd8bc2f65f99f942f00be99fc5f9f993"></a>

## port_ranges property — ingress_egress_gw.allowed_vip_port.custom_ports / cc28a4d33d66 / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-a07f4f68bd57c5ced98c31fa1cf8e9da8c1ab4a05b792acbe3514acab4a400dd"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.custom_ports / cc28a4d33d66 / 5

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e13603423676e9ec874cd77fbfb19a2af5147c9fb282130f7d4ace14f5e0d2be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ff957a92cbe4c56ad67cae0710941f9257e60cb6b39a904da40366fa6392293"></a>

## ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port — ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port / aee57ad762f9 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-b00985accbc0eb02b082516319fd46904b30687dbab9be0159797a9e49d2658b"></a>

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
disable_allowed_vip_port = {}
```

<a id="canonical-494477527949eef3b333495d048dda5069b71d9f264596d4c31e2306f3347085"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port / aee57ad762f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-235e0eef845da25e1514e64fc2743e1b08cf6ee531fd9d7d5d294ee38a9317ce"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port / aee57ad762f9 / 4

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ee80dd985ed384ffdd62c817d8d5225e44b69932d2d1de5c4e429be032e8c834"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f18486068756282fc6133e11b1a23dfb64d9798f8afc6e90867ace5ef246f585"></a>

## ingress_egress_gw.allowed_vip_port.use_http_https_port — ingress_egress_gw.allowed_vip_port.use_http_https_port / 2e2e1d18cf1d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- ingress_egress_gw.allowed_vip_port.use_http_https_port

<a id="canonical-a30e7606f4316465c822a1bd2d13edd80963c6b9ded136f398916ce52ae14c81"></a>

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
use_http_https_port = {}
```

<a id="canonical-9e98165071622282b1d193eccd0024bd5b0b86e1d7c32dcb7942891971f8d6f7"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.use_http_https_port / 2e2e1d18cf1d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78269a51ab429aaf5bd235f45bebfbd68b0db0323c8bf2fb4c5fb4abe8b838bb"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.use_http_https_port / 2e2e1d18cf1d / 4

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-8f7a1ecf7db862b19ddaea4186562960a879787c1a0114b257bd722cbc371d3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c7abbe61ceecd2802441a2f0fd6150432c4b657352bf3ab6498ac99c1a1a07e"></a>

## ingress_egress_gw.allowed_vip_port.use_http_port — ingress_egress_gw.allowed_vip_port.use_http_port / 047d300bd900 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- ingress_egress_gw.allowed_vip_port.use_http_port

<a id="canonical-4134ae38c2c18ae124832722eba6c8b7f006d3b68843ea020133a0b0c49c5297"></a>

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
use_http_port = {}
```

<a id="canonical-e66a22ef41e62c9d11d6e4c5bb3df0873dbbbb7d60454b96b58ec385b317848b"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.use_http_port / 047d300bd900 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7088641b1b41edb6992e418d3a857b48cd43416613e5c367ca5cfb23e300c0a4"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.use_http_port / 047d300bd900 / 4

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-d945251fd3562544a0177017b256177bfc57459bd4ad3a78a3b79f6e1346be21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c98bc90fb65da98a534e3cdaedcb7e3a47a2f2600068b2051114d19e9a3bf58"></a>

## ingress_egress_gw.allowed_vip_port.use_https_port — ingress_egress_gw.allowed_vip_port.use_https_port / c9652802cbd1 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- ingress_egress_gw.allowed_vip_port.use_https_port

<a id="canonical-d182c7daa3febbb0473b55d8e80d8abdf8de167ad20207133c5563f8322a2e7b"></a>

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
use_https_port = {}
```

<a id="canonical-41fbed64d6dc52e4de6ab2b5b6b4bdde4f23f6bc8ae36b4a10222e346e5a41d4"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port.use_https_port / c9652802cbd1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-27e5db1d6068f271476ac390ac27d5f886cf5487287fc1118a06fd6238eb2e82"></a>

## Next pages — ingress_egress_gw.allowed_vip_port.use_https_port / c9652802cbd1 / 4

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-207dec8b79b5b53d737c3e38b2af5c401b5c86b56d6b50f2e5d394b41bf10b43)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07b23283e29ef950b45c1a834abb9a7201c473eb859f4ff16e919a5a9311b458"></a>

## ingress_egress_gw.allowed_vip_port_sli — ingress_egress_gw.allowed_vip_port_sli / 7aab93c4d610 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.allowed_vip_port_sli

<a id="canonical-194c8b543f177b8df15d9bbb4d35808592b800c1e97a0b865ace4e4194082853"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-2d2dc09dca034aa4db1fd67ca3445dd5b78087b8a79683a45960b57a95b03d82"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli / 7aab93c4d610 / 3

- [custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-da7ea1920982defa587cbb0dcc50a95c71efd8dcfaa168de106c351e8b56655e): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-75c41a909dda12aa0c403bfafc6f3b66edf61a873c9f14a56b4dc30aae1d8e90): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-b112e80e73fcb9909f340ce245e3747bca538c322bd2426cbbee2523b3ca7f81): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-e0f891082e8cb6ca58f426c79abdbc3eb4ed1eabcfdbe3f8ec326d64d3524bb2): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-1677c0a0b4a6766d3f092be8470793a991d6d3e5b640d9094d9f267053f2c119): complete subsection reference.

<a id="canonical-5b1a4bd0def1aa49778a658cafef8a3c64a444f16ca39490a573d74f6a785c59"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli / 7aab93c4d610 / 4

- [ingress_egress_gw.allowed_vip_port_sli.custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-da7ea1920982defa587cbb0dcc50a95c71efd8dcfaa168de106c351e8b56655e)
- [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-75c41a909dda12aa0c403bfafc6f3b66edf61a873c9f14a56b4dc30aae1d8e90)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-b112e80e73fcb9909f340ce245e3747bca538c322bd2426cbbee2523b3ca7f81)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-e0f891082e8cb6ca58f426c79abdbc3eb4ed1eabcfdbe3f8ec326d64d3524bb2)
- [ingress_egress_gw.allowed_vip_port_sli.use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-1677c0a0b4a6766d3f092be8470793a991d6d3e5b640d9094d9f267053f2c119)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-da7ea1920982defa587cbb0dcc50a95c71efd8dcfaa168de106c351e8b56655e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-788788c1603f2e735783065d803a0422b8f5b6c8868c4e51412d8c7fe34a6fc4"></a>

## ingress_egress_gw.allowed_vip_port_sli.custom_ports — ingress_egress_gw.allowed_vip_port_sli.custom_ports / 822828173bb4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- ingress_egress_gw.allowed_vip_port_sli.custom_ports

<a id="canonical-c7480beda7f70c03cd5a7c0ef927f4ca81d01d295a7a97e7ae1e1b98104fd807"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-3507a33d82899edcf94597624c1597947cd233356fd4628bd117be0c32ce9c06"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.custom_ports / 822828173bb4 / 3

<a id="canonical-6995e62b6656d27675fe68990e83f2396eead2b5c36bb7d226434b87bc01a9c5"></a>

<a id="canonical-1b63209c4f372513c6db1c7fff2687e483b61e4136d0a7dc98f4f3ab5fc08452"></a>

## port_ranges property — ingress_egress_gw.allowed_vip_port_sli.custom_ports / 822828173bb4 / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-b464cef288027f1d44dee831b3fe79bd66b9874b1b9dca8da93435e2d6b54147"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.custom_ports / 822828173bb4 / 5

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-75c41a909dda12aa0c403bfafc6f3b66edf61a873c9f14a56b4dc30aae1d8e90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e880626ab2b526b994e6e943780f87343d777e12f241c6a3f10906f727e44727"></a>

## ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port — ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port / db9ba9724ea1 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port

<a id="canonical-edda45df2e5c28d7858d94765b215cae18a53f86454ba80a1b614d53e54d73a0"></a>

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
disable_allowed_vip_port = {}
```

<a id="canonical-b92f7c2a70c7bf269d911c96a1ac46d2f53b03167b321f1d16484f502100830a"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port / db9ba9724ea1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-22fa102c340256d838bda7655b393ebbb9381c1e586a3a75861dd925ef7f8be3"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port / db9ba9724ea1 / 4

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-b112e80e73fcb9909f340ce245e3747bca538c322bd2426cbbee2523b3ca7f81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27c71017287dcc80e47402e752df6c18ffb803fc9b83a717317fe71226c49e0c"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_http_https_port — ingress_egress_gw.allowed_vip_port_sli.use_http_https_port / 33868dad1138 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- ingress_egress_gw.allowed_vip_port_sli.use_http_https_port

<a id="canonical-0a086d169fd327ac7971585af0cfc65a86b59fd093d84f07c697f7364c9f09c4"></a>

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
use_http_https_port = {}
```

<a id="canonical-e47e777ef21e39d0430a55491cfaa1ccae3398f7c8240077a6143e505174ada7"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.use_http_https_port / 33868dad1138 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a0a1d1f698420cba8164ff02d9dcd727bced42008ff846f0a2540f96fa26fcc"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.use_http_https_port / 33868dad1138 / 4

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e0f891082e8cb6ca58f426c79abdbc3eb4ed1eabcfdbe3f8ec326d64d3524bb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19d940de953f5f31b68bf80b40af1c5133e5057da3150d6c793c5dbeb59e2d73"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_http_port — ingress_egress_gw.allowed_vip_port_sli.use_http_port / 62b2d7096ae4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- ingress_egress_gw.allowed_vip_port_sli.use_http_port

<a id="canonical-7b18c927a105c086cb6ac2421d68b2148ae3c9794b216ccd7f315f11e86fd958"></a>

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
use_http_port = {}
```

<a id="canonical-e14445fff15b689cd3e98f2ff2487aeded425dd20d6e4f9de1e80fed2d9a037f"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.use_http_port / 62b2d7096ae4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-08fc3c6ac2e2cb1beddf2b30a6faadddacfd2251cf0b86a4cc7b2fb594fbb881"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.use_http_port / 62b2d7096ae4 / 4

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1677c0a0b4a6766d3f092be8470793a991d6d3e5b640d9094d9f267053f2c119"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b2d2f7dc608beed08fead4c7cbaf4553f33c93baf0d8fc2ae040b661e7fa562"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_https_port — ingress_egress_gw.allowed_vip_port_sli.use_https_port / b5c04b165e67 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- ingress_egress_gw.allowed_vip_port_sli.use_https_port

<a id="canonical-b47b948d225b54fa4228f0dc0a85e45191b0ad4f144bb0c052a8703bd5f44f1d"></a>

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
use_https_port = {}
```

<a id="canonical-3fdc45ba541d81916b8e84ed453e7670f6789c01958caad3fc11bc9dd4bc0b99"></a>

## Direct properties — ingress_egress_gw.allowed_vip_port_sli.use_https_port / b5c04b165e67 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d6f6f53c673581b03cc3374cbde3b0442dfe08c715d152968ad470732cc4767"></a>

## Next pages — ingress_egress_gw.allowed_vip_port_sli.use_https_port / b5c04b165e67 / 4

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-39812507a95ea4a4c506f2d602f8bc25369969c866d5d9d6c512271dc2e90a49)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99309cb355c66666649cd112b1032f2c430269c06d48ca997cdcff52a0c818b1"></a>

## ingress_egress_gw.az_nodes — ingress_egress_gw.az_nodes / 37e2d50ef84c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.az_nodes

<a id="canonical-c37f11b7ce3f1c15608232489609493a40d07c22758b99363b6ecf253ad9b273"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name"),
  validators.ConflictingListObjectAttributes("inside_subnet",
    "reserved_inside_subnet")}
```

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

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-24ab85535543e6b3279fe129762442c76814938f5050b59cc7a1b5d632c667e2"></a>

## Direct properties — ingress_egress_gw.az_nodes / 37e2d50ef84c / 3

<a id="canonical-79d88448ed88ef60eb1ad367805eda3925e0874b2224e74ff01c289158869654"></a>

<a id="canonical-da4e4a5690cba738fad29b67cdb33690e5678da857b6808d8ffe566a8bebd03b"></a>

## aws_az_name property — ingress_egress_gw.az_nodes / 37e2d50ef84c / 4

Type: `"string"`. Optional.

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

- [inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-ffdcd15f75a313464d89060320f84612d73e5f4cf14310663817760509a31231): complete subsection reference.

- [outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-607605e2b9d887b0e65c9729a6254c579cc8cd7556e21c0f9c7fd5664efbc581): complete subsection reference.

- [reserved_inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-ba3ed3f555e2f1ec7e7cddbeaed8d217f5d92652596f6755b98e34260c603762): complete subsection reference.

- [workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-a2cac5253cd0fb40e17e1b2b1d6bedcfb58e54c4e73fd809e18d13335350f71c): complete subsection reference.

<a id="canonical-82ff2e6965ff73c3f852ed7e35c709221b3fbb70a3e73b4cf514c84239573a69"></a>

## Next pages — ingress_egress_gw.az_nodes / 37e2d50ef84c / 5

- [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-ffdcd15f75a313464d89060320f84612d73e5f4cf14310663817760509a31231)
- [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-607605e2b9d887b0e65c9729a6254c579cc8cd7556e21c0f9c7fd5664efbc581)
- [ingress_egress_gw.az_nodes.reserved_inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-ba3ed3f555e2f1ec7e7cddbeaed8d217f5d92652596f6755b98e34260c603762)
- [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-a2cac5253cd0fb40e17e1b2b1d6bedcfb58e54c4e73fd809e18d13335350f71c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ffdcd15f75a313464d89060320f84612d73e5f4cf14310663817760509a31231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-094018dd78f054a7343e0f45be369f07a6f6b3758c7d057c236be419b0e73998"></a>

## ingress_egress_gw.az_nodes.inside_subnet — ingress_egress_gw.az_nodes.inside_subnet / 851937585ae3 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- ingress_egress_gw.az_nodes.inside_subnet

<a id="canonical-2f77f6ea17af2aee161cd5ffb55219d59a0483d0514a86ad4cafea2d672463e3"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-7a13184bc505d7b639e32d52ef08fe2180e14e24a73ca97ef82bc3631c21f5a3"></a>

## Direct properties — ingress_egress_gw.az_nodes.inside_subnet / 851937585ae3 / 3

<a id="canonical-0590169f208510cb5f339968988bf9391c12c2e9b9a1287973ee82d9d39adeeb"></a>

<a id="canonical-eace08bdad4f7ae5968456d3a155ee792cb037d2be02d597b6f92081062a932b"></a>

## existing_subnet_id property — ingress_egress_gw.az_nodes.inside_subnet / 851937585ae3 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

- [subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-0aac961a5e2f058032db5948f78dd89a6f8620523b2234df7e1a6c4914b9c477): complete subsection reference.

<a id="canonical-f3e83b51e35a4a3e3427fff7d5245d194e2ec700768f89870dd61894fd7d804f"></a>

## Next pages — ingress_egress_gw.az_nodes.inside_subnet / 851937585ae3 / 5

- [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-0aac961a5e2f058032db5948f78dd89a6f8620523b2234df7e1a6c4914b9c477)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-0aac961a5e2f058032db5948f78dd89a6f8620523b2234df7e1a6c4914b9c477"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b73068c102889457ce47cc151faadeb8b88a87c3b2089f356fe0556ad9f5010a"></a>

## ingress_egress_gw.az_nodes.inside_subnet.subnet_param — ingress_egress_gw.az_nodes.inside_subnet.subnet_param / a963facf7ddd / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-ffdcd15f75a313464d89060320f84612d73e5f4cf14310663817760509a31231)
- ingress_egress_gw.az_nodes.inside_subnet.subnet_param

<a id="canonical-d02b9a3cdff42365ad8160161c64c782b26884aae9ea05138227022654fa15bf"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-f6e1813e3d0cb27828c3b8a0602b686430103f9297d6710b594f8f94cfe38812"></a>

## Direct properties — ingress_egress_gw.az_nodes.inside_subnet.subnet_param / a963facf7ddd / 3

<a id="canonical-55af322486f01b721cc2645d01935d33d40d0f6e149740d56952e01d58f9e9f9"></a>

<a id="canonical-32dc1faf1722eee656767650c26a2f9b8757cd75aa3567ee19a09659038f4347"></a>

## ipv4 property — ingress_egress_gw.az_nodes.inside_subnet.subnet_param / a963facf7ddd / 4

Type: `"string"`. Optional.

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

<a id="canonical-06b51de494578609b416dccaa31f2f4a9e34927c4863ac5226a05ede2f390c3e"></a>

## Next pages — ingress_egress_gw.az_nodes.inside_subnet.subnet_param / a963facf7ddd / 5

- [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-ffdcd15f75a313464d89060320f84612d73e5f4cf14310663817760509a31231)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-607605e2b9d887b0e65c9729a6254c579cc8cd7556e21c0f9c7fd5664efbc581"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-512dca220266ed0cf32973a5c5df528153f867615e63a32404a6d396ab8ed279"></a>

## ingress_egress_gw.az_nodes.outside_subnet — ingress_egress_gw.az_nodes.outside_subnet / 9bb7d4afa773 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- ingress_egress_gw.az_nodes.outside_subnet

<a id="canonical-7ad5b08697ac2716d64e6c9514f9910e84cb0788fd2cf86a0a0dec7f1895c8bd"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
outside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0fe8d3d2af7c1ec08e2c43fb4c583a61bda35190710dd42cf33cfe01ccf7f8e"></a>

## Direct properties — ingress_egress_gw.az_nodes.outside_subnet / 9bb7d4afa773 / 3

<a id="canonical-14636ffe9d0465f7109ab1ca431e84c4b357f2588db70b306f114b945adc41ea"></a>

<a id="canonical-39a9fb9a70716d05403dba6a1ca3af9e2ec1fc6388cecbb2f9c980ea7038e195"></a>

## existing_subnet_id property — ingress_egress_gw.az_nodes.outside_subnet / 9bb7d4afa773 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

- [subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-38e64ecb95e75cc7b7aaeb500e97e9db982ce8b225cc275c4cda6e76243fc927): complete subsection reference.

<a id="canonical-67e085c21a35bd350aa727118f11d2505e650837269a9da2ee910b6fde4514f7"></a>

## Next pages — ingress_egress_gw.az_nodes.outside_subnet / 9bb7d4afa773 / 5

- [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-38e64ecb95e75cc7b7aaeb500e97e9db982ce8b225cc275c4cda6e76243fc927)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-38e64ecb95e75cc7b7aaeb500e97e9db982ce8b225cc275c4cda6e76243fc927"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-926a51a6b145c4a81cc09d53b9883996ea9f8b9a4703c59213d666ddf28fef45"></a>

## ingress_egress_gw.az_nodes.outside_subnet.subnet_param — ingress_egress_gw.az_nodes.outside_subnet.subnet_param / 60fdccc3f3fd / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-607605e2b9d887b0e65c9729a6254c579cc8cd7556e21c0f9c7fd5664efbc581)
- ingress_egress_gw.az_nodes.outside_subnet.subnet_param

<a id="canonical-9cd7e2c159349ef3761983231b06d77c33306b40882f63a7535b149bcf6258eb"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-f84d96c106eb139d5e8b9024e34c06889240b3f39779321b4730e91ccaa47391"></a>

## Direct properties — ingress_egress_gw.az_nodes.outside_subnet.subnet_param / 60fdccc3f3fd / 3

<a id="canonical-da07c752c5c6c13aeb56ebd85c76898a3e8f49265c1c576545b464eaefd934fa"></a>

<a id="canonical-6e8fe94217df9ed9bd811546370e7c253133647d36386163e8d1326dfc78fffb"></a>

## ipv4 property — ingress_egress_gw.az_nodes.outside_subnet.subnet_param / 60fdccc3f3fd / 4

Type: `"string"`. Optional.

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

<a id="canonical-4af7cdde7df5f851c2f13ae975083dcfe95ff9a16a5a787a1895a33be993fe4c"></a>

## Next pages — ingress_egress_gw.az_nodes.outside_subnet.subnet_param / 60fdccc3f3fd / 5

- [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-607605e2b9d887b0e65c9729a6254c579cc8cd7556e21c0f9c7fd5664efbc581)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ba3ed3f555e2f1ec7e7cddbeaed8d217f5d92652596f6755b98e34260c603762"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1aa10ff91e929ce79baee9b0c7f4160b7276d164c14d59a3b76606d0e80ef083"></a>

## ingress_egress_gw.az_nodes.reserved_inside_subnet — ingress_egress_gw.az_nodes.reserved_inside_subnet / 296795875c70 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- ingress_egress_gw.az_nodes.reserved_inside_subnet

<a id="canonical-56b5bdbf6d9d0355fadbb3d89827393b7f1cb26b6fc1809d4202dd5b9593556f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
reserved_inside_subnet = {}
```

<a id="canonical-dca29a6fba79c5c37411b8691947e7055214e4e6dab0b976e7545d7657b49438"></a>

## Direct properties — ingress_egress_gw.az_nodes.reserved_inside_subnet / 296795875c70 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f2bad7ae74d33f5a83c25460bca431fcc1a86eceda979fe6708cc7fec88e85e2"></a>

## Next pages — ingress_egress_gw.az_nodes.reserved_inside_subnet / 296795875c70 / 4

- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a2cac5253cd0fb40e17e1b2b1d6bedcfb58e54c4e73fd809e18d13335350f71c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20f445b1385f7ddfab779fad40ea57cdeb3074290bfdc82407e41fa3ea00b786"></a>

## ingress_egress_gw.az_nodes.workload_subnet — ingress_egress_gw.az_nodes.workload_subnet / c3174d594590 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- ingress_egress_gw.az_nodes.workload_subnet

<a id="canonical-cf055d6e18a05490bd7e2ad6ec3c61a8cc47116d16e1bd1f9c2410a02a99a470"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for workload subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
workload_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-e79aa6c4169212c3c0b7227e5ee60682d62bf9ce1d1a3da3d589b643feb8d987"></a>

## Direct properties — ingress_egress_gw.az_nodes.workload_subnet / c3174d594590 / 3

<a id="canonical-da1ac5ef0cd6941d9fd8750797aafdb4bb0bbe382659a602b2b4c98c8ed4cf50"></a>

<a id="canonical-c6a1756ad0d7628f863e0d7cf5de6b3ae119c9e1e9ee8130e9690a9be2287189"></a>

## existing_subnet_id property — ingress_egress_gw.az_nodes.workload_subnet / c3174d594590 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

- [subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-79a3741d8e97cb27c754fba002ac9d4e08aa9db6f3a70102ee75f73c0f751d98): complete subsection reference.

<a id="canonical-7471a8232d48212079326d7accad532c701167e5fef8d50da80ff6073e6b369f"></a>

## Next pages — ingress_egress_gw.az_nodes.workload_subnet / c3174d594590 / 5

- [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-79a3741d8e97cb27c754fba002ac9d4e08aa9db6f3a70102ee75f73c0f751d98)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-79a3741d8e97cb27c754fba002ac9d4e08aa9db6f3a70102ee75f73c0f751d98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1c897274382f710dd80c84b576649aaf81cd768da2a16be4f671e0065cceba8"></a>

## ingress_egress_gw.az_nodes.workload_subnet.subnet_param — ingress_egress_gw.az_nodes.workload_subnet.subnet_param / e08248e9faa0 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-efb535f744f131f9c7d6a0dc9034f269572dfea3282f6ebaf20c09ce9bc09922)
- [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-a2cac5253cd0fb40e17e1b2b1d6bedcfb58e54c4e73fd809e18d13335350f71c)
- ingress_egress_gw.az_nodes.workload_subnet.subnet_param

<a id="canonical-c52d753f12d51d6d63da352338e4c9b1f1f89365d7da5185e375133490c2f42e"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-c94e7f9629b1a9247cfca5c390a6e30d7fdc0806f5d9e4a6a9f11fad75882d4a"></a>

## Direct properties — ingress_egress_gw.az_nodes.workload_subnet.subnet_param / e08248e9faa0 / 3

<a id="canonical-bfa9913d02b0f54eb57fb808dbf75379818efe1f4b82d7b8e54551294579e64d"></a>

<a id="canonical-22d40d31ac832e1ba2bb9576841921502988c4631eb15d1d5f4f2af2906915b1"></a>

## ipv4 property — ingress_egress_gw.az_nodes.workload_subnet.subnet_param / e08248e9faa0 / 4

Type: `"string"`. Optional.

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

<a id="canonical-a89bdec5a932cf40b4d82c3a54e1de4a39ca5e96437bc08b1798d2a836d6e698"></a>

## Next pages — ingress_egress_gw.az_nodes.workload_subnet.subnet_param / e08248e9faa0 / 5

- [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-a2cac5253cd0fb40e17e1b2b1d6bedcfb58e54c4e73fd809e18d13335350f71c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-6250e0449953556e0f52164a22f9a506c4cb1a29f295372fbdfa57b9d18dda25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d4f1b9fdd25a72b64a499ca9cf924124e599d96f24c88b91538219bc2c1d2e5"></a>

## ingress_egress_gw.dc_cluster_group_inside_vn — ingress_egress_gw.dc_cluster_group_inside_vn / d4566603425a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.dc_cluster_group_inside_vn

<a id="canonical-5b308535412116d3dd70851d34bde75d1dfe333ec37cc6f9f361375333dd6895"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
dc_cluster_group_inside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-bd68bb1dc48486ad6c9c8320ddad6754bb3a8a31b94780503c35b5d46b681291"></a>

## Direct properties — ingress_egress_gw.dc_cluster_group_inside_vn / d4566603425a / 3

<a id="canonical-669526408deb9a178563868cd459d9d73f5851c1fc41d22dfcc7a1fb3fed55db"></a>

<a id="canonical-5f17c36e7f84ffa3aa185e86d57123b21b5472d8f99edc47e2868ecb6d919bd9"></a>

## name property — ingress_egress_gw.dc_cluster_group_inside_vn / d4566603425a / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-866fc46bc686e0a377996f8972b34a13aa177e33ff0226404d86bacfed21e48e"></a>

<a id="canonical-67d36a7c51cec410e7633223711db1db57586fda8d7ec93d6f8f5b6c3a815922"></a>

## namespace property — ingress_egress_gw.dc_cluster_group_inside_vn / d4566603425a / 5

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
}
```

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

<a id="canonical-2add9ca3c5407426812e25c930fc65b313fa04f14155f20e749ba16c4d0f6739"></a>

<a id="canonical-9fc3e988b5d0e6a8c93091e78757d19b4871b15008510a29355d0859f723a16a"></a>

## tenant property — ingress_egress_gw.dc_cluster_group_inside_vn / d4566603425a / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-de19a1712e6b9ced00fdfe2240a3727db5b54d0de80bb8d4dcda4dd1ddd64082"></a>

## Next pages — ingress_egress_gw.dc_cluster_group_inside_vn / d4566603425a / 7

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-0514f74fd9de4c088d7e716d36bba6e3a81e7172ecc0b55bf0868f6e654185cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52ab913500ebf27768a1ce456ee8f4296a3d1d745f0e914260710aff31c34804"></a>

## ingress_egress_gw.dc_cluster_group_outside_vn — ingress_egress_gw.dc_cluster_group_outside_vn / 4c6d3659691a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.dc_cluster_group_outside_vn

<a id="canonical-d11e6e57e9d4e99d82c96dddc56641d7e99d4540edea60c212bd1890a6550aee"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
dc_cluster_group_outside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-d30043e88911b5b98054c4c2738c5fecd34eea65cc2a0827ca6ade710a7a0005"></a>

## Direct properties — ingress_egress_gw.dc_cluster_group_outside_vn / 4c6d3659691a / 3

<a id="canonical-2e1c9e22c4f18f4ecc3f0b40aec88e73c2cac79aa39538d83820e1f8d504210f"></a>

<a id="canonical-25bf5d4752308498cbdbe83cde5b963594c6d1c27c8d569ba82ac3cf734d9d61"></a>

## name property — ingress_egress_gw.dc_cluster_group_outside_vn / 4c6d3659691a / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-51696de63176d949276085e49a6e185bb6b69adf5d77d04cbf4fbc2bb6cc35de"></a>

<a id="canonical-3196bf65e1d6eecfd3e07a879803555b22c508701867ab6f965e2e27eee278b1"></a>

## namespace property — ingress_egress_gw.dc_cluster_group_outside_vn / 4c6d3659691a / 5

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
}
```

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

<a id="canonical-ea1f0309d0e5a2a655a5c7b804c11d02ef020ed5372082bcd69d023217f24b83"></a>

<a id="canonical-56a6256dfb15e3d0cf0a7ac2ce8eb74db3e40b396f122731d0491e19f08195d0"></a>

## tenant property — ingress_egress_gw.dc_cluster_group_outside_vn / 4c6d3659691a / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-bc17df64b946384ff115431a5ae96e8fff6a45f8712ce3c3a2eff5c2f59d8eea"></a>

## Next pages — ingress_egress_gw.dc_cluster_group_outside_vn / 4c6d3659691a / 7

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-bd1bb7489a9cdcacd2d5f27ba6dcb6ac8e03cfd472684d709f73916056e114e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4537cf08591d56a78a0892f04d3ad3429bb2fe2cce57a561f6f7708e6adeb3c"></a>

## ingress_egress_gw.forward_proxy_allow_all — ingress_egress_gw.forward_proxy_allow_all / 4a80604c6bcf / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.forward_proxy_allow_all

<a id="canonical-c0345d2378a430d8870a5e4f94ec1650cfeed217a1583a5b4789e768608b7f94"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
forward_proxy_allow_all = {}
```

<a id="canonical-f59cc241fc8c7b2dead01aff96e629b0a6882629a6cfe5b1f6ff2ec4e987fcf4"></a>

## Direct properties — ingress_egress_gw.forward_proxy_allow_all / 4a80604c6bcf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f895f399fe75972337df1666809558979a7e902b5ecc3c1adc559a523af3a8d"></a>

## Next pages — ingress_egress_gw.forward_proxy_allow_all / 4a80604c6bcf / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23be9c04d0ccbfada536504c2ad66864b6f6864f96a70b1618a128fa48d69d93"></a>

## ingress_egress_gw.global_network_list — ingress_egress_gw.global_network_list / dbd2e54b8e41 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.global_network_list

<a id="canonical-b5c0e95591f50bbff5f178b79a718d9240e896ac1522bd406082b1b02dcc7887"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-50166c3a359ff7f6211fd3b6c98ce4bd191984dd91a764a4b73f0f766dd67a3e"></a>

## Direct properties — ingress_egress_gw.global_network_list / dbd2e54b8e41 / 3

- [global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19): complete subsection reference.

<a id="canonical-be50b7e109248acf1fd005a7c95e6ea8acad50817ec06aca91f1ce25704ff0e6"></a>

## Next pages — ingress_egress_gw.global_network_list / dbd2e54b8e41 / 4

- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3527ad11c328d68c5e83cd0276559b76dc31b97265888621c48132d43cc189b"></a>

## ingress_egress_gw.global_network_list.global_network_connections — ingress_egress_gw.global_network_list.global_network_connections / 9bcda5fc02c4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="canonical-a605a318fdaa12e16e233a002ffc79d2931865d82f3c690412d696759737439c"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

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

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-30b7a6498896b9a8624e105deee132a760a01764acfb9cb2734cecfe03e536ed"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections / 9bcda5fc02c4 / 3

- [sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-a48d8ab032da2b5331b45736c47e60330c98ffa6df01a21cb2a968e7d582cbae): complete subsection reference.

- [slo_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-ff5a02b59fae6111ec96950bd953b02e9f5ead8e1059bb190c8ea62b3aed5fdd): complete subsection reference.

<a id="canonical-15d6178b1ac9a6dd2d24ed5f37d59aecbde803e07cc5a0b17ea52a004b7735f4"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections / 9bcda5fc02c4 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-a48d8ab032da2b5331b45736c47e60330c98ffa6df01a21cb2a968e7d582cbae)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-ff5a02b59fae6111ec96950bd953b02e9f5ead8e1059bb190c8ea62b3aed5fdd)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a48d8ab032da2b5331b45736c47e60330c98ffa6df01a21cb2a968e7d582cbae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b14f329a90a2436a3e296fa54de2ebfa9d29862a9bcd4974844af11b007f98ce"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / d7896d2c8b88 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-51fcafcc196e64c296afcb97b06b11b94c58a61f06e87a6e28286ffd6cd901ab"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-fc19e174e84fea4586bb75f6e4b4f8272fdcd3662c5c0ff78c5fe590182dc49c"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / d7896d2c8b88 / 3

- [global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-609986f4cb2ec6515b20e494837c0d935a49d12fab9a6458c42cfc5842055713): complete subsection reference.

<a id="canonical-9c255302e892e42bdd0a3690b450ae8b7d7e837fc8468cccc5b36877baef75c2"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / d7896d2c8b88 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-609986f4cb2ec6515b20e494837c0d935a49d12fab9a6458c42cfc5842055713)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-609986f4cb2ec6515b20e494837c0d935a49d12fab9a6458c42cfc5842055713"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6faae51627a4c06be2c467f82b44a1a40736a5f14c9418d680d4698ddb15c02b"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 30c561e6724a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19)
- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-a48d8ab032da2b5331b45736c47e60330c98ffa6df01a21cb2a968e7d582cbae)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-2887107343287c2746dc652fadf606931aa71f0838e68e086593b92fb2c00fbc"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-23fbdbcfa0bd4688c35d1687761faee0cb509d5352d27e44f591aa263503d12d"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 30c561e6724a / 3

<a id="canonical-1738ca21fc112f6710fe0b854b4386250fc094b3d94dd9d3234d57a484be0f81"></a>

<a id="canonical-608a9a952d64d29e0068036dc5835f61a704da08df8ea091e6b8debd98f0325c"></a>

## name property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 30c561e6724a / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-964c21925d45389bd14dd491e7eeb739b740ed8c8c7975bda0404914e28ec372"></a>

<a id="canonical-e5c5b1ec6340d01a0468a82f7cfd8b0e43f6c88d4bb5239d4fad3dc605a9560f"></a>

## namespace property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 30c561e6724a / 5

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
}
```

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

<a id="canonical-57a89557d2cd276f20c3d60c57339fcc7335172218c49e887de08e2c4c43c975"></a>

<a id="canonical-c4dd77b2e6599d669ff91d0266c4c8336674b114ecc9a9663ac4376736290ada"></a>

## tenant property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 30c561e6724a / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-c75b6a77cf3f1df4ef54c483a2e83ab4f483fc5948f2192ca707df0eeab84915"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 30c561e6724a / 7

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-a48d8ab032da2b5331b45736c47e60330c98ffa6df01a21cb2a968e7d582cbae)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ff5a02b59fae6111ec96950bd953b02e9f5ead8e1059bb190c8ea62b3aed5fdd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b185e78d038612e38295722b8011824618d47346dd640b808f01ff22e4c43e2"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 116814f318da / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-61c965dcf99fe256cd4f12c04be05d49acf6c8d8e5e9c7819d4760b0d704f519"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-f77b03ea62d9f6f7c1e6849e376619a80945fdcc3d41b76efda3cb7abaff522e"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 116814f318da / 3

- [global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-daf9140827d10d15e7dedc5adda775481f9ba476be8acfb076b450dd3d3152cd): complete subsection reference.

<a id="canonical-3029d2032baced8e37771e549c202e53693bf843c9d9365890c18c7bf4fff27e"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 116814f318da / 4

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-daf9140827d10d15e7dedc5adda775481f9ba476be8acfb076b450dd3d3152cd)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-daf9140827d10d15e7dedc5adda775481f9ba476be8acfb076b450dd3d3152cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc28614c5b8f903359cdfa72a126b6cb1a53bedc575acb12b00c640ac5293f29"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 1c33da7fb961 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-145819feaf852fe679b92852e7363c3c64d25c60df132c2577aa3c710f4e01e9)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-ffcaa3ee6f0fe83fc41c3438b17fd9ada2e6fde94dd5b3c54db8df53ca790c19)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-ff5a02b59fae6111ec96950bd953b02e9f5ead8e1059bb190c8ea62b3aed5fdd)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-87e531582e6dea504dcde9d921a6ccbc68184327b7a919ea0a117513a3c64684"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-52674396b9ca39fca733e9a7acee79ac9ae5c420ad9d568b7b31907b403dc4f0"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 1c33da7fb961 / 3

<a id="canonical-24623e2c08e70f2c5cc3d3d00195d871606ade0677c8d48dc08fd91ec946c4c5"></a>

<a id="canonical-7f2ba02b8a47773215d00ce583833634218e3b3d18652fb83e0722fc3200a1f2"></a>

## name property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 1c33da7fb961 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-ea898a2769d92abbc086aed10aa402c1e4c020ef888085adf6dea81fc4db9041"></a>

<a id="canonical-ddea09f76771dfc1cf3201db2eea516a208ed706fee9a3323f1ebdde3383dc9b"></a>

## namespace property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 1c33da7fb961 / 5

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
}
```

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

<a id="canonical-4550d01e61fe43de53fafc4e08b8fbb729ee6f5ed27fbb0e93d3b8ef13d12886"></a>

<a id="canonical-a7ceda3381bc048e4aa1ff3d5705a23d8a900633683dded89b3b99745453f787"></a>

## tenant property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 1c33da7fb961 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3f6a866cfc97caf1d1befbe04eea746733f39f50dd6a569f185cd5910d549a38"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 1c33da7fb961 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-ff5a02b59fae6111ec96950bd953b02e9f5ead8e1059bb190c8ea62b3aed5fdd)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
