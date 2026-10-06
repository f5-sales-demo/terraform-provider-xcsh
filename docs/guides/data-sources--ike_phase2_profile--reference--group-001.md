---
page_title: "xcsh_ike_phase2_profile reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile reference."
---

# xcsh_ike_phase2_profile reference

<a id="canonical-1221020020223301-3321201330233132-0012122311312233-0211331333321112-1210022031000211-2123121212122103-1000331210101312-1310102310302122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130)
- Property reference

<a id="canonical-2112013232322113-3113101112320011-3203100000222200-2130101020312030-2013013113201012-3331210122103313-0322330210300102-3033031011120131"></a>

### Direct properties for `xcsh_ike_phase2_profile`

<a id="canonical-3001301100011003-1110113310021211-1011300122131213-2333132231100022-2011003023033301-1100310103233213-2113102032301000-2031311230230220"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-3023310110122012-0233201001120313-0122231020310211-3101112200101021-1323122202213031-3000210212113120-0220302331023320-0112212131110003"></a>

<a id="canonical-0110331030012313-1232130102203330-1022122113322212-3313123013130322-2132011212202132-1302033223013322-2002203312022331-0113220023002131"></a>

#### `authentication_algos` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3012322301023301-3220013212320112-3012002213303200-3100120133031303-0212033331330301-1302212311100022-2003302232312113-0212130123302132"></a>

<a id="canonical-1223233011003011-2133211221003110-1212210033221033-1213333130332333-1231300112011032-2220021130030210-0102001201131001-3302323123131123"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the IKEPhase2Profile.

Additional upstream details:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [dh_group_set](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0111123121220033-0201132202202103-2033200311320112-2123300322002312-0213003332021311-0012110313331122-0032012330333223-2110022330003013): complete subsection reference.

- [disable_pfs](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0031003200012310-1231032022221232-1312212101231301-2100310311221213-1230223032131131-1011223020122212-1313102003001323-2201022112312103): complete subsection reference.

<a id="canonical-3031111301320201-1113301310210303-0012220120031033-0302102001213121-0310210210102011-1103123122303103-2011020022330101-2033121321320321"></a>

<a id="canonical-2300033230003101-2111302031223011-0213112112033301-2002011012003202-0202022201112133-2133330102110121-1131011013033000-3333001210320132"></a>

#### `encryption_algos` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0022111001103300-2231221333001223-3313221030101101-2311103001103133-1112321133312023-3230332301122101-1213331123001003-2310301202130120"></a>

<a id="canonical-1221130313230212-1033100021313013-3333131020110303-0331210010110100-2332333031220113-1223312101303201-2113103233300222-1330312213031022"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0221231222222301-0211130333012203-0003230320030031-1010003201022212-1021323232023230-3102301221130200-0002230220230222-1213032102021021): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike_phase2_profile--reference--group-001.md#canonical-3110023100310321-3211000132212003-1211001313310313-2323132132231223-2131032031310313-1022320312031330-0302321200131310-2133333103210022): complete subsection reference.

<a id="canonical-3223203330210101-2032201301223020-0002132330210303-3002233133311130-2302201220102101-0123013212113312-0301212312301201-3103203020133112"></a>

<a id="canonical-2303221122301331-2023203033220130-3102203320202322-1223113122131003-2020222132131303-0011233122021310-2220202002023331-2020221333130210"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-0131131203230002-1311311011101212-0012100132222032-0132200012122200-0032203203020112-3302223212013123-2201330202320330-3003012201332010"></a>

<a id="canonical-2001100313202002-1121020221302003-2130110023133132-1300230310330303-2210321313121210-0000103112020102-0323302003132002-2322332213232132"></a>

#### `name` property

Type: `"string"`. Required.

Name of the IKEPhase2Profile.

Additional upstream details:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0320213103201131-3220231322010223-2002331121113000-1331003002130010-1012001301030031-2001121301023010-2312311211131000-2111001221011123"></a>

<a id="canonical-2200130103203031-1201201021113303-0121331210331201-1313021202010322-0122123021222011-2301221101011110-1201123321310303-2133020011023330"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the IKEPhase2Profile exists.

Additional upstream details:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [use_default_keylifetime](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1131211011201200-1001013132021311-2132032102321131-2212203122011112-2013131101331332-1113121310320203-2230212211000202-2002113332222103): complete subsection reference.

<a id="canonical-0222212312132102-0311001213102232-1233032103013223-0310021223132310-3132322020311301-2020032212130020-1320310212230200-1000101202201201"></a>

### All schema paths for `xcsh_ike_phase2_profile`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike_phase2_profile--reference--group-001.md#canonical-3001301100011003-1110113310021211-1011300122131213-2333132231100022-2011003023033301-1100310103233213-2113102032301000-2031311230230220) |
| `authentication_algos` | [authentication_algos](data-sources--ike_phase2_profile--reference--group-001.md#canonical-3023310110122012-0233201001120313-0122231020310211-3101112200101021-1323122202213031-3000210212113120-0220302331023320-0112212131110003) |
| `description` | [description](data-sources--ike_phase2_profile--reference--group-001.md#canonical-3012322301023301-3220013212320112-3012002213303200-3100120133031303-0212033331330301-1302212311100022-2003302232312113-0212130123302132) |
| `dh_group_set` | [dh_group_set](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1301102320001111-3111011031301000-2212003100322302-1113021011131201-3000333113112311-0111111323313330-1020333222313133-2023021003230103) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](data-sources--ike_phase2_profile--reference--group-001.md#canonical-2232113013301010-2223212331110231-0101111120022230-1213321300220310-3000200100212003-1333332220331203-3321211333323302-2331032311320101) |
| `disable_pfs` | [disable_pfs](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0133332113113011-0302200020101300-1130113230202003-2230123331023321-1222102121312210-2203132231010122-1333103022113310-1311121003311121) |
| `encryption_algos` | [encryption_algos](data-sources--ike_phase2_profile--reference--group-001.md#canonical-3031111301320201-1113301310210303-0012220120031033-0302102001213121-0310210210102011-1103123122303103-2011020022330101-2033121321320321) |
| `id` | [ID](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0022111001103300-2231221333001223-3313221030101101-2311103001103133-1112321133312023-3230332301122101-1213331123001003-2310301202130120) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike_phase2_profile--reference--group-001.md#canonical-2212001002312020-2332120231113330-3033332102010220-0031233313103233-0322101132033020-2022203032331013-1122031320332121-3120123010202120) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1200212332333023-0201232113111232-1233221311131020-0311211301301132-3031030123012211-3021132221223223-2001012300121120-3313212010312111) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1231110130300232-2111310031333222-2001333222313311-2112320333301321-2321030231333000-1310011123211002-3220313300223021-0112001310113101) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0212011123313203-1202200011330122-0302110121011002-2002302221223201-2332323100013202-1120333321232332-2230301220112231-2020130302220303) |
| `labels` | [labels](data-sources--ike_phase2_profile--reference--group-001.md#canonical-3223203330210101-2032201301223020-0002132330210303-3002233133311130-2302201220102101-0123013212113312-0301212312301201-3103203020133112) |
| `name` | [name](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0131131203230002-1311311011101212-0012100132222032-0132200012122200-0032203203020112-3302223212013123-2201330202320330-3003012201332010) |
| `namespace` | [namespace](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0320213103201131-3220231322010223-2002331121113000-1331003002130010-1012001301030031-2001121301023010-2312311211131000-2111001221011123) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1321211330130012-0332113101313310-2322121303232330-3313230033132032-2000111112130321-3203311030312012-1000122231300111-0203002310123221) |

<a id="canonical-0111123121220033-0201132202202103-2033200311320112-2123300322002312-0213003332021311-0012110313331122-0032012330333223-2110022330003013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dh_group_set` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1221020020223301-3321201330233132-0012122311312233-0211331333321112-1210022031000211-2123121212122103-1000331210101312-1310102310302122)
- dh_group_set

<a id="canonical-1301102320001111-3111011031301000-2212003100322302-1113021011131201-3000333113112311-0111111323313330-1020333222313133-2023021003230103"></a>

Type: `"single"`. Computed.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

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

- [dh_group_set](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1301102320001111-3111011031301000-2212003100322302-1113021011131201-3000333113112311-0111111323313330-1020333222313133-2023021003230103)
- [disable_pfs](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0133332113113011-0302200020101300-1130113230202003-2230123331023321-1222102121312210-2203132231010122-1333103022113310-1311121003311121)

Select alternatives according to the provider validators above.

<a id="canonical-2102002220030302-2321223033322130-0003212120121311-1010132111120002-0003213302303130-1003102231321110-0000031203311212-1320303220032033"></a>

### Direct properties for `dh_group_set`

<a id="canonical-2232113013301010-2223212331110231-0101111120022230-1213321300220310-3000200100212003-1333332220331203-3321211333323302-2331032311320101"></a>

#### `dh_group_set.dh_groups` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0031003200012310-1231032022221232-1312212101231301-2100310311221213-1230223032131131-1011223020122212-1313102003001323-2201022112312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_pfs` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1221020020223301-3321201330233132-0012122311312233-0211331333321112-1210022031000211-2123121212122103-1000331210101312-1310102310302122)
- disable_pfs

<a id="canonical-0133332113113011-0302200020101300-1130113230202003-2230123331023321-1222102121312210-2203132231010122-1333103022113310-1311121003311121"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221231222222301-0211130333012203-0003230320030031-1010003201022212-1021323232023230-3102301221130200-0002230220230222-1213032102021021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ike_keylifetime_hours` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1221020020223301-3321201330233132-0012122311312233-0211331333321112-1210022031000211-2123121212122103-1000331210101312-1310102310302122)
- ike_keylifetime_hours

<a id="canonical-2212001002312020-2332120231113330-3033332102010220-0031233313103233-0322101132033020-2022203032331013-1122031320332121-3120123010202120"></a>

Type: `"single"`. Computed.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Additional upstream details:

Input Hours.

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

- [ike_keylifetime_hours](data-sources--ike_phase2_profile--reference--group-001.md#canonical-2212001002312020-2332120231113330-3033332102010220-0031233313103233-0322101132033020-2022203032331013-1122031320332121-3120123010202120)
- [ike_keylifetime_minutes](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1231110130300232-2111310031333222-2001333222313311-2112320333301321-2321030231333000-1310011123211002-3220313300223021-0112001310113101)
- [use_default_keylifetime](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1321211330130012-0332113101313310-2322121303232330-3313230033132032-2000111112130321-3203311030312012-1000122231300111-0203002310123221)

Select alternatives according to the provider validators above.

<a id="canonical-1230022300211013-3110031121131032-1100210101131122-3210202211310133-1000312301223312-2133311132120003-3302121120120333-3101110130202231"></a>

### Direct properties for `ike_keylifetime_hours`

<a id="canonical-1200212332333023-0201232113111232-1233221311131020-0311211301301132-3031030123012211-3021132221223223-2001012300121120-3313212010312111"></a>

#### `ike_keylifetime_hours.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3110023100310321-3211000132212003-1211001313310313-2323132132231223-2131032031310313-1022320312031330-0302321200131310-2133333103210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ike_keylifetime_minutes` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1221020020223301-3321201330233132-0012122311312233-0211331333321112-1210022031000211-2123121212122103-1000331210101312-1310102310302122)
- ike_keylifetime_minutes

<a id="canonical-1231110130300232-2111310031333222-2001333222313311-2112320333301321-2321030231333000-1310011123211002-3220313300223021-0112001310113101"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Additional upstream details:

Set IKE Key Lifetime in minutes.

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

<a id="canonical-1331030021133200-0303103331312213-1133230232113210-0133330201331213-0122202131333202-1332022221033203-0011213103332310-1113203030010102"></a>

### Direct properties for `ike_keylifetime_minutes`

<a id="canonical-0212011123313203-1202200011330122-0302110121011002-2002302221223201-2332323100013202-1120333321232332-2230301220112231-2020130302220303"></a>

#### `ike_keylifetime_minutes.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1131211011201200-1001013132021311-2132032102321131-2212203122011112-2013131101331332-1113121310320203-2230212211000202-2002113332222103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_keylifetime` properties

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1221020020223301-3321201330233132-0012122311312233-0211331333321112-1210022031000211-2123121212122103-1000331210101312-1310102310302122)
- use_default_keylifetime

<a id="canonical-1321211330130012-0332113101313310-2322121303232330-3313230033132032-2000111112130321-3203311030312012-1000122231300111-0203002310123221"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.
