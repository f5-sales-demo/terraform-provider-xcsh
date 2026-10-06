---
page_title: "xcsh_ike_phase2_profile reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile reference."
---

# xcsh_ike_phase2_profile reference

<a id="canonical-2020002230212010-1313022100100310-3220332211132032-0022322013313122-2101323310333120-1102230003321002-1003312302120232-3300001312010212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- Property reference

<a id="canonical-3312212211312322-2211201222200121-0010113013100133-2323123330012321-1132232220303301-1302032113312301-2132130301201100-2011001023100200"></a>

### Direct properties for `xcsh_ike_phase2_profile`

<a id="canonical-2120032311123130-2333033013113210-2320010321013110-0312201300321200-2333302112210311-3121211221113031-2010103311103321-2032003003213322"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

<a id="canonical-0313200212131132-3000203033332330-2223033021121020-2201130131223130-0013323032010130-1231112320123200-0231030322103031-0222210111112330"></a>

<a id="canonical-2201113112010321-2112213003230002-3121203001132212-3020033121211003-3212230133202303-0111132321303231-0203032313010022-1013132033323113"></a>

#### `authentication_algos` property

Type: `["list", "string"]`. Required.

\[Enum: AUTH\_ALG\_DEFAULT|SHA256\_HMAC|SHA384\_HMAC|SHA512\_HMAC|AUTH\_ALG\_NONE\] Choose one or
more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption
algorithms. Possible values are \`AUTH\_ALG\_DEFAULT\`, \`SHA256\_HMAC\`, \`SHA384\_HMAC\`,
\`SHA512\_HMAC\`, \`AUTH\_ALG\_NONE\`. Defaults to \`AUTH\_ALG\_DEFAULT\`.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3010201121302101-0000131322220221-3111013002032103-0223023303221033-0321331131032202-0020232003030130-2030321200233232-2030010322202023"></a>

<a id="canonical-2211331133011200-2123222100221111-3113233300202102-2103132330213201-2322211120002000-2122101232002010-1033031311032001-3131302003311331"></a>

#### `description` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [dh_group_set](resources--ike_phase2_profile--reference--group-001.md#canonical-0133212113121220-2130201020223302-2301133222102133-0032320323013023-1212213333211310-3200230010202120-2303020033323020-2303101213112202): complete subsection reference.

<a id="canonical-2013022012203001-3210302121233201-2202233010333201-3110023213203030-1213001320023330-2230103001213033-2331133331000210-3333333313300010"></a>

<a id="canonical-0032310010103132-1210333002231022-2323010332200113-0320133202011221-2323321021231302-3301033201213301-3331131332020031-0122110101021213"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

- [disable_pfs](resources--ike_phase2_profile--reference--group-001.md#canonical-3031211031320220-0013203313003321-1310032113330010-1101300130320003-1211220101113133-1300032010002332-2012121201003130-0332131300131013): complete subsection reference.

<a id="canonical-3203213123321130-2232210221102201-2302021232100131-3000200301122020-1300313033001230-0112131101201203-2333221010023112-2303133001302101"></a>

<a id="canonical-0330021002311312-1233302012133220-0302213011210303-3302011222211333-1202120310303222-0021332322110103-0033000012012030-1202231130332213"></a>

#### `encryption_algos` property

Type: `["list", "string"]`. Required.

\[Enum:
ENC\_ALG\_DEFAULT|AES128\_CBC|AES192\_CBC|AES256\_CBC|TRIPLE\_DES\_CBC|AES128\_GCM|AES192\_GCM|AES256\_GCM\]
Choose one or more encryption algorithms. Possible values are \`ENC\_ALG\_DEFAULT\`,
\`AES128\_CBC\`, \`AES192\_CBC\`, \`AES256\_CBC\`, \`TRIPLE\_DES\_CBC\`, \`AES128\_GCM\`,
\`AES192\_GCM\`, \`AES256\_GCM\`. Defaults to \`ENC\_ALG\_DEFAULT\`.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2211132031020023-3313113210103202-3032300002122311-0111003131021030-1233100302002220-0212230011101111-2111111201122322-0130231321322022"></a>

<a id="canonical-3131330301233131-3103011031031030-0022110323022112-0222300031323211-2322011331202100-1231300102300333-3212213010320001-0023121313013101"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](resources--ike_phase2_profile--reference--group-001.md#canonical-3032303322112330-0322031301203323-2311300222303030-1002101131310020-2103031232210111-1131312232231103-2020331320111212-0002020120033013): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike_phase2_profile--reference--group-001.md#canonical-3013020132011201-3031230131033311-1222301300203032-0032302231133013-2323200223031321-3010001230222303-0120122032301323-1023103330212020): complete subsection reference.

<a id="canonical-0233330320212112-3332002322311101-2332311121021331-3331231113021231-1300320302113201-3021201313313113-0322030100130311-2300133311002113"></a>

<a id="canonical-3310120322032320-2223322102002100-1033031232000210-1033110113233311-2212211330130022-2301322013111100-2201103001030023-2023211031023313"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

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

<a id="canonical-3130122201032020-3023131000212302-3110011222033021-0130031213032333-2113121311111332-1030123032102132-0020133131032023-1100023031110223"></a>

<a id="canonical-0332100122032303-0013333230003023-3031313033122322-2321113311121211-3012131221111121-0203313212230022-3030122022223202-3020012113230310"></a>

#### `name` property

Type: `"string"`. Required.

Name of the IKE Phase2 Profile. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1000103301103322-2031121210200001-0000121213032132-3201233121023132-3201110323210023-0321302130220312-1332010123230211-0310223332111010"></a>

<a id="canonical-3101300130230013-2013300323130030-0102313001003131-2022330110213023-3101232001112110-2101310323102101-3331123121221022-1030002122001221"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the IKE Phase2 Profile is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [timeouts](resources--ike_phase2_profile--reference--group-001.md#canonical-3101332331313122-2110323302100320-3013322103011212-1110230011312133-1203231332233113-3333330120131210-2311230130323223-0003221123222333): complete subsection reference.

- [use_default_keylifetime](resources--ike_phase2_profile--reference--group-001.md#canonical-3311120103212311-2300230312003320-0230300311311003-2300313113323101-0330132023231103-1112301201122112-1323111222011123-1000331221220110): complete subsection reference.

<a id="canonical-2000302301031200-2302113013222212-1230022310213331-2321100222010321-0212212000020030-2122212121132332-2121230211203030-1331323200013020"></a>

### All schema paths for `xcsh_ike_phase2_profile`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike_phase2_profile--reference--group-001.md#canonical-2120032311123130-2333033013113210-2320010321013110-0312201300321200-2333302112210311-3121211221113031-2010103311103321-2032003003213322) |
| `authentication_algos` | [authentication_algos](resources--ike_phase2_profile--reference--group-001.md#canonical-0313200212131132-3000203033332330-2223033021121020-2201130131223130-0013323032010130-1231112320123200-0231030322103031-0222210111112330) |
| `description` | [description](resources--ike_phase2_profile--reference--group-001.md#canonical-3010201121302101-0000131322220221-3111013002032103-0223023303221033-0321331131032202-0020232003030130-2030321200233232-2030010322202023) |
| `dh_group_set` | [dh_group_set](resources--ike_phase2_profile--reference--group-001.md#canonical-2002312203112232-2202013332010131-1300321220022300-3203133110203122-2212332013010311-1003333133312003-3212332100301002-3230311221313233) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](resources--ike_phase2_profile--reference--group-001.md#canonical-0131110223101032-1131011021222102-3023011213122120-1123312212133111-0112222331131322-1222221123020312-0022232113122102-1320230202300302) |
| `disable` | [disable](resources--ike_phase2_profile--reference--group-001.md#canonical-2013022012203001-3210302121233201-2202233010333201-3110023213203030-1213001320023330-2230103001213033-2331133331000210-3333333313300010) |
| `disable_pfs` | [disable_pfs](resources--ike_phase2_profile--reference--group-001.md#canonical-1122302010012121-1131132333111221-3222222113102130-1021032120323211-1113230022032331-1111000101131032-0031103320202310-1002301202221202) |
| `encryption_algos` | [encryption_algos](resources--ike_phase2_profile--reference--group-001.md#canonical-3203213123321130-2232210221102201-2302021232100131-3000200301122020-1300313033001230-0112131101201203-2333221010023112-2303133001302101) |
| `id` | [ID](resources--ike_phase2_profile--reference--group-001.md#canonical-2211132031020023-3313113210103202-3032300002122311-0111003131021030-1233100302002220-0212230011101111-2111111201122322-0130231321322022) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike_phase2_profile--reference--group-001.md#canonical-1022332323000102-0332333112300200-1020301211301002-1210031023122333-2131212320033032-1312113002032331-0011220312132320-2122011323113232) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike_phase2_profile--reference--group-001.md#canonical-0011202121131331-0323211312012211-1100202120101330-2122200220021333-0220323223132310-3002021120020313-3310023132002331-2222033130111032) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike_phase2_profile--reference--group-001.md#canonical-3001103033322310-1120213321013202-3012002233211222-0302210330212322-1200222121013212-0320112333233210-3201301022022313-3321223011030300) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike_phase2_profile--reference--group-001.md#canonical-3300310011133200-0312232103101122-2011121010303131-3231310122320202-1101032220303313-2222200222130212-0102033330021011-2230220313320212) |
| `labels` | [labels](resources--ike_phase2_profile--reference--group-001.md#canonical-0233330320212112-3332002322311101-2332311121021331-3331231113021231-1300320302113201-3021201313313113-0322030100130311-2300133311002113) |
| `name` | [name](resources--ike_phase2_profile--reference--group-001.md#canonical-3130122201032020-3023131000212302-3110011222033021-0130031213032333-2113121311111332-1030123032102132-0020133131032023-1100023031110223) |
| `namespace` | [namespace](resources--ike_phase2_profile--reference--group-001.md#canonical-1000103301103322-2031121210200001-0000121213032132-3201233121023132-3201110323210023-0321302130220312-1332010123230211-0310223332111010) |
| `timeouts` | [timeouts](resources--ike_phase2_profile--reference--group-001.md#canonical-1231001211003321-3112312123100021-1300220032103321-1003331213230001-1300030200213033-1100320020313010-2111030311231111-0010322001133122) |
| `timeouts.create` | [timeouts.create](resources--ike_phase2_profile--reference--group-001.md#canonical-3101331000133030-2333130132032220-3120010012230001-2300333122213302-3222121202022323-2321122100102222-2012133302132211-0313222133030130) |
| `timeouts.delete` | [timeouts.delete](resources--ike_phase2_profile--reference--group-001.md#canonical-0023233122020123-0022033131222323-0023102013103020-1321011232221103-1133300132030132-0010231002202130-3100013310302022-2202003300033031) |
| `timeouts.read` | [timeouts.read](resources--ike_phase2_profile--reference--group-001.md#canonical-2100113023031233-0122021300102111-2010110303231130-0012103212100231-3331321101300123-2002331220333300-0003020013331210-2311121201332103) |
| `timeouts.update` | [timeouts.update](resources--ike_phase2_profile--reference--group-001.md#canonical-2232103122032013-1212301302303312-0232113122300012-1033221112003001-2203303221001013-2101213100031231-3322333021101032-3021223020213112) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike_phase2_profile--reference--group-001.md#canonical-1132201132113333-1330020002221323-3332321121201211-3033032322103020-3233331302012313-0031213231301130-2322230211130301-1312221111000133) |

<a id="canonical-0133212113121220-2130201020223302-2301133222102133-0032320323013023-1212213333211310-3200230010202120-2303020033323020-2303101213112202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dh_group_set` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-2020002230212010-1313022100100310-3220332211132032-0022322013313122-2101323310333120-1102230003321002-1003312302120232-3300001312010212)
- dh_group_set

<a id="canonical-2002312203112232-2202013332010131-1300321220022300-3203133110203122-2212332013010311-1003333133312003-3212332100301002-3230311221313233"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dh_groups")}
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

OneOf alternatives in this subsection:

- [dh_group_set](resources--ike_phase2_profile--reference--group-001.md#canonical-2002312203112232-2202013332010131-1300321220022300-3203133110203122-2212332013010311-1003333133312003-3212332100301002-3230311221313233)
- [disable_pfs](resources--ike_phase2_profile--reference--group-001.md#canonical-1122302010012121-1131132333111221-3222222113102130-1021032120323211-1113230022032331-1111000101131032-0031103320202310-1002301202221202)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dh_group_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001001320310121-2122113130302212-1300230300210023-0213131023202210-2131123110122232-2331112013202320-1200301231022110-1100201322233211"></a>

### Direct properties for `dh_group_set`

<a id="canonical-0131110223101032-1131011021222102-3023011213122120-1123312212133111-0112222331131322-1222221123020312-0022232113122102-1320230202300302"></a>

#### `dh_group_set.dh_groups` property

Type: `["list", "string"]`. Optional.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3031211031320220-0013203313003321-1310032113330010-1101300130320003-1211220101113133-1300032010002332-2012121201003130-0332131300131013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_pfs` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-2020002230212010-1313022100100310-3220332211132032-0022322013313122-2101323310333120-1102230003321002-1003312302120232-3300001312010212)
- disable_pfs

<a id="canonical-1122302010012121-1131132333111221-3222222113102130-1021032120323211-1113230022032331-1111000101131032-0031103320202310-1002301202221202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable pfs.

Additional upstream details:

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
disable_pfs = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032303322112330-0322031301203323-2311300222303030-1002101131310020-2103031232210111-1131312232231103-2020331320111212-0002020120033013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ike_keylifetime_hours` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-2020002230212010-1313022100100310-3220332211132032-0022322013313122-2101323310333120-1102230003321002-1003312302120232-3300001312010212)
- ike_keylifetime_hours

<a id="canonical-1022332323000102-0332333112300200-1020301211301002-1210031023122333-2131212320033032-1312113002032331-0011220312132320-2122011323113232"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Additional upstream details:

Input Hours.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](resources--ike_phase2_profile--reference--group-001.md#canonical-1022332323000102-0332333112300200-1020301211301002-1210031023122333-2131212320033032-1312113002032331-0011220312132320-2122011323113232)
- [ike_keylifetime_minutes](resources--ike_phase2_profile--reference--group-001.md#canonical-3001103033322310-1120213321013202-3012002233211222-0302210330212322-1200222121013212-0320112333233210-3201301022022313-3321223011030300)
- [use_default_keylifetime](resources--ike_phase2_profile--reference--group-001.md#canonical-1132201132113333-1330020002221323-3332321121201211-3033032322103020-3233331302012313-0031213231301130-2322230211130301-1312221111000133)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020301003322123-0131320210333120-0021203122121101-0033122100213212-1223210230112211-1200210033221000-1102312013120300-2011022211203021"></a>

### Direct properties for `ike_keylifetime_hours`

<a id="canonical-0011202121131331-0323211312012211-1100202120101330-2122200220021333-0220323223132310-3002021120020313-3310023132002331-2222033130111032"></a>

#### `ike_keylifetime_hours.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-3013020132011201-3031230131033311-1222301300203032-0032302231133013-2323200223031321-3010001230222303-0120122032301323-1023103330212020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ike_keylifetime_minutes` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-2020002230212010-1313022100100310-3220332211132032-0022322013313122-2101323310333120-1102230003321002-1003312302120232-3300001312010212)
- ike_keylifetime_minutes

<a id="canonical-3001103033322310-1120213321013202-3012002233211222-0302210330212322-1200222121013212-0320112333233210-3201301022022313-3321223011030300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ike keylifetime minutes.

Additional upstream details:

Set IKE Key Lifetime in minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
ike_keylifetime_minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301320111330212-0220322301233001-2322012003031003-2321310110331100-3231112002322233-1011303120222011-1232013321122302-3311131133000201"></a>

### Direct properties for `ike_keylifetime_minutes`

<a id="canonical-3300310011133200-0312232103101122-2011121010303131-3231310122320202-1101032220303313-2222200222130212-0102033330021011-2230220313320212"></a>

#### `ike_keylifetime_minutes.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(10, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-3101332331313122-2110323302100320-3013322103011212-1110230011312133-1203231332233113-3333330120131210-2311230130323223-0003221123222333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-2020002230212010-1313022100100310-3220332211132032-0022322013313122-2101323310333120-1102230003321002-1003312302120232-3300001312010212)
- timeouts

<a id="canonical-1231001211003321-3112312123100021-1300220032103321-1003331213230001-1300030200213033-1100320020313010-2111030311231111-0010322001133122"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222120312001103-2313122310210203-3122211133223030-1102301232320222-3112323310101310-3110213100100211-3230120212221132-1321111221032122"></a>

### Direct properties for `timeouts`

<a id="canonical-3101331000133030-2333130132032220-3120010012230001-2300333122213302-3222121202022323-2321122100102222-2012133302132211-0313222133030130"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0023233122020123-0022033131222323-0023102013103020-1321011232221103-1133300132030132-0010231002202130-3100013310302022-2202003300033031"></a>

<a id="canonical-0120231122303102-3332321101233200-2211001222000301-3011012212013132-2011231320310030-0100102113212323-0011123033333033-1213100013011321"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2100113023031233-0122021300102111-2010110303231130-0012103212100231-3331321101300123-2002331220333300-0003020013331210-2311121201332103"></a>

<a id="canonical-3210003031123010-1133102113230122-0201211001022222-1321000112231000-3132200022201130-2121233202100000-2131130031013120-0023203303102333"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2232103122032013-1212301302303312-0232113122300012-1033221112003001-2203303221001013-2101213100031231-3322333021101032-3021223020213112"></a>

<a id="canonical-2332322012223133-0223023013232021-1111302020120213-0131122032100133-3113322200021203-0200311232111220-0220013000112322-2000013333310110"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3311120103212311-2300230312003320-0230300311311003-2300313113323101-0330132023231103-1112301201122112-1323111222011123-1000331221220110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_keylifetime` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-2020002230212010-1313022100100310-3220332211132032-0022322013313122-2101323310333120-1102230003321002-1003312302120232-3300001312010212)
- use_default_keylifetime

<a id="canonical-1132201132113333-1330020002221323-3332321121201211-3033032322103020-3233331302012313-0031213231301130-2322230211130301-1312221111000133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use default keylifetime.

Additional upstream details:

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
use_default_keylifetime = {}
```

This is an empty object or choice marker. It has no direct properties.
