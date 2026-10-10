---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-3302323021312013-0123100021032230-2300321112130323-0101023000212032-3201130221213303-1003213100111311-2101031231210311-3332000302302310"></a>

## `primary.rr_set_group.rr_set.loc_record.values.longitude_degree` property

Type: `"number"`. Optional.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 180),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3311213031211000-2203110013110303-3011131310313213-2102122010222302-0311321031303203-3303003111302301-1221312122122210-3213102101323012"></a>

<a id="canonical-3111113020223033-0112303333322330-0130002001102020-3122000223010111-1201111232021302-2333101120111023-3211331130010310-2102201301202033"></a>

## `primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere` property

Type: `"string"`. Optional.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["E","W"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("E",
    "W"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3333311133103112-1223122203322303-1101213130001103-2003122121110130-0121201213333112-0302012011133303-3302000002003231-0010030132101310"></a>

<a id="canonical-0000013322213310-1113300033310201-2110133312022000-0122310011011202-0232121331123132-1121313323131200-2323133230230001-3231133231020011"></a>

## `primary.rr_set_group.rr_set.loc_record.values.longitude_minute` property

Type: `"number"`. Optional.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-0003300311211321-3302033301013000-1303223131102001-0021310012301222-3002221110003121-3102100110322301-3301103010223123-0202202000321131"></a>

<a id="canonical-2033210000213122-1301120210312330-1101230022320020-0213032003021103-2111000102200032-2031030110123233-3110011230210330-2002221032212322"></a>

## `primary.rr_set_group.rr_set.loc_record.values.longitude_second` property

Type: `"number"`. Optional.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-3023332133221031-2213200201011103-3213012023033333-0312213013012003-2200302200332113-0320122133201000-3000300330100222-0030211201320001"></a>

<a id="canonical-2231131000002132-2031221313120212-0303131213012230-2110322110311023-2313030303320132-0032123023131012-3202233133022202-1312313222232202"></a>

## `primary.rr_set_group.rr_set.loc_record.values.vertical_precision` property

Type: `"number"`. Optional.

Vertical Precision. Vertical Precision in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-2003001322011302-1301231333320300-1203302312221330-2033303022131132-3121320130110031-2332100310010333-2130011113103032-1022320310102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.mx_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.mx_record

<a id="canonical-2132030313003332-3230003011110221-0130002031232112-3133020121102322-2223230102301312-2023030330230330-3131300100100313-3221132223101331"></a>

Type: `"object"`. single nested block, Optional.

DNSMXResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
mx_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122020202222122-3200033213010313-3211200032310000-0230313112203201-0032120202032121-1201232010310033-0302003001012100-3311112132232133"></a>

### Direct properties for `primary.rr_set_group.rr_set.mx_record`

<a id="canonical-0102031200201131-1132303302003222-3210331133220110-3332322020300030-0033310330203230-2010332323122023-0032030122310211-1112003010201310"></a>

#### `primary.rr_set_group.rr_set.mx_record.name` property

Type: `"string"`. Optional.

MX Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-3203011211313001-0002133210230122-3102011130012031-2123133001220331-1333031133312300-1102233313120122-3020331320013120-3302110102201210): complete subsection reference.

<a id="canonical-3203011211313001-0002133210230122-3102011130012031-2123133001220331-1333031133312300-1102233313120122-3020331320013120-3302110102201210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.mx_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-2003001322011302-1301231333320300-1203302312221330-2033303022131132-3121320130110031-2332100310010333-2130011113103032-1022320310102202)
- primary.rr_set_group.rr_set.mx_record.values

<a id="canonical-2202333202101131-1132102103013000-1112330121332320-3212223020221231-2023130231102331-3203303102210210-0033220002221200-0313132030222011"></a>

Type: `"object"`. list nested block, Optional.

MX Record Value. Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112200001220210-2121121231303102-0021110000130213-2323101201222231-1023130330321001-1133331023220311-2013112302031000-0211303322101231"></a>

### Direct properties for `primary.rr_set_group.rr_set.mx_record.values`

<a id="canonical-1133102322110311-3101323123323131-3101301203111231-3223212323211213-1331110102101010-3233312022332330-2030332033303201-1102021310032131"></a>

#### `primary.rr_set_group.rr_set.mx_record.values.domain` property

Type: `"string"`. Optional.

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3111313022123021-1133213322112032-2203121002200232-1121301000133330-3000133312233103-1333132030312300-1102030223123323-0022112233333301"></a>

<a id="canonical-2001120210103100-0122300202300200-3003302122022122-2100032032112130-1233101123231200-1130302013321301-0323113131012023-3003320232333220"></a>

#### `primary.rr_set_group.rr_set.mx_record.values.priority` property

Type: `"number"`. Optional.

Priority. Mail exchanger priority code.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.naptr_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.naptr_record

<a id="canonical-0113102211212002-1101222021022110-3101000310021221-3213123131122132-2222303310010210-3123021310123133-0312033331022113-3003012030033011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for naptr record.

Additional upstream details:

DNS NAPTR Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
naptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121111313002311-2101111110301202-2232110213020131-2130000321003001-1201222310210113-3203121023112103-1323012223321200-3322220212332022"></a>

### Direct properties for `primary.rr_set_group.rr_set.naptr_record`

<a id="canonical-1200012232231001-1301100312311031-2133300113100220-1030101021010222-2032223233100320-0333201103130110-0123103332311012-3120002221030012"></a>

#### `primary.rr_set_group.rr_set.naptr_record.name` property

Type: `"string"`. Optional.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-0100220222133303-1230232023133312-3132333010213303-0000021310121331-0121023031303003-2003023202002101-3320001002112000-1133232110312031): complete subsection reference.

<a id="canonical-0100220222133303-1230232023133312-3132333010213303-0000021310121331-0121023031303003-2003023202002101-3320001002112000-1133232110312031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.naptr_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223)
- primary.rr_set_group.rr_set.naptr_record.values

<a id="canonical-0222200233003323-0322031231313100-1201022110113322-1130023112121030-3103233101313121-1313030302302333-2322013132210231-0032111113103033"></a>

Type: `"object"`. list nested block, Optional.

NAPTR Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("flags",
    "order",
    "preference")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022320311320031-1300120011010013-3232321003033010-2132331011112131-2303030320101110-0303303131033320-0013330301033113-3313010120123021"></a>

### Direct properties for `primary.rr_set_group.rr_set.naptr_record.values`

<a id="canonical-2000332132310023-1130303122313003-2101233300213131-2030002031201201-3111212121131033-1302032203222301-3013111030321123-0112300300330223"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.flags` property

Type: `"string"`. Optional.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-1131013031102231-1311302210303031-1031033032330023-0013330201133112-1300021322013100-2210010012012122-3132301203200123-0231312001211031"></a>

<a id="canonical-0201333103133120-1302131002332330-2200312223320331-2211223221322010-3131323201013213-3320313221120201-2313212112223002-2211212311120221"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.order` property

Type: `"number"`. Optional.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2011121233210300-0121201122303221-3300003023003222-0132232213132022-2310301231212110-1203320320310023-2113033110111312-1203321032011212"></a>

<a id="canonical-3131323232010220-0300033020330113-0211123133203211-0332121110221120-3321133202100231-0311133132221013-0023020311210333-0001332123210123"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.preference` property

Type: `"number"`. Optional.

Preference when records have the same order. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1022122030201233-0030103013313101-1220222121123303-1230222122133123-2112013030212201-3332121011100211-0330212323001202-3302211120230300"></a>

<a id="canonical-3011300303002113-3010311223032000-2022010013110133-2103332022113102-0102200220313212-0120021200003233-2223221331202111-0202302131320323"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.regexp` property

Type: `"string"`. Optional.

Regular expression to construct the next domain name to lookup.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1202323213312222-3123130121131000-3331022313122333-2223222230203203-1020000320101032-3102220101303203-1212203303132201-1211223210113211"></a>

<a id="canonical-1231330021122112-3011101000110202-2222210001103003-1023330232322123-0132012303233213-1302310011100032-2313313220111121-1211303032122031"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.replacement` property

Type: `"string"`. Optional.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1030011011201021-2321133132211031-2303023321320202-1201311121123320-2320202102312212-1333111131112123-2131131301202323-1123331311213000"></a>

<a id="canonical-3321232322231220-3123032230323321-0311020312302312-0132321333001102-0210332313023133-3121302110333332-3221020121032133-2233300120001123"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.service` property

Type: `"string"`. Optional.

Specifies the service(s) available down this rewrite path.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-3112102233211200-0313123203112300-1130013102230113-2011023232120220-1302012230210220-3202131201133110-0011021300201313-1102311332102132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ns_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.ns_record

<a id="canonical-1231003122123303-2133013001222111-3132233130013001-3012311102221220-2103210112001212-3323000203323113-3120302331000232-2021211213312321"></a>

Type: `"object"`. single nested block, Optional.

DNSNSResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
ns_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100301123130020-3302220311011323-2310302223223013-3300121033133313-2223013333021311-2222110122020131-0233211200222033-1011310010003031"></a>

### Direct properties for `primary.rr_set_group.rr_set.ns_record`

<a id="canonical-3320020101210010-1122130010031303-3301323332322213-2301330010001022-0121210231130222-1032122301130321-1320031201003011-2112231102302113"></a>

#### `primary.rr_set_group.rr_set.ns_record.name` property

Type: `"string"`. Optional.

NS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2013013321002000-1221302210011312-1303231000120231-3302010111320010-1331010332113222-2302333313221033-3033000111223302-1212201201230302"></a>

<a id="canonical-0212130032303313-2221001321333220-2320203313132110-3000033022120133-3212220131001221-2303122122321301-2122103303210212-1202003031023312"></a>

#### `primary.rr_set_group.rr_set.ns_record.values` property

Type: `["list", "string"]`. Optional.

Name Servers. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3232000310031110-1111032130320130-1330010220023223-2023200213300002-3123322311111130-2121011102311023-1203013003000322-3300000321210110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ptr_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.ptr_record

<a id="canonical-3013221023311022-0020211130201303-3131202122033003-1300210133223000-3222010023132110-1001102031230113-0020222130133000-1122303200203200"></a>

Type: `"object"`. single nested block, Optional.

DNSPTRResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
ptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200232120201130-3023122202210111-0021131213323022-2310202220311210-2121113111021231-2012123010301031-2312012002303332-3111100220303323"></a>

### Direct properties for `primary.rr_set_group.rr_set.ptr_record`

<a id="canonical-3100223311113131-0103210012013210-0331322110221133-0020032321022320-2022203200101222-2310003322230001-0213311230111011-1202012120212032"></a>

#### `primary.rr_set_group.rr_set.ptr_record.name` property

Type: `"string"`. Optional.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2023312301232211-1132111221331013-2102012300332121-0312000310311303-0030132032012211-3231013321332002-1203032112210230-0230333012212213"></a>

<a id="canonical-0003003020300201-1203300332302302-2211313302121102-2201013003222122-0303220032222212-1332012001133103-0012002320012212-0220220120300033"></a>

#### `primary.rr_set_group.rr_set.ptr_record.values` property

Type: `["list", "string"]`. Optional.

Domain Name. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3331101323210103-3033131313122302-0130321211322012-2233211131121323-1203222333302212-2332301313331233-3123010110333331-2023103210211121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.srv_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.srv_record

<a id="canonical-2001132300321330-0200133330021203-0313213322013100-0111121110201021-2012001100220230-1000022333221202-3112310302033103-1021221233100113"></a>

Type: `"object"`. single nested block, Optional.

DNSSRVResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "values")}
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
srv_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010011213222131-1303312101320110-2122021302102011-1010210011031233-1121132211230010-3211103102001123-2022032031312023-0331210333231300"></a>

### Direct properties for `primary.rr_set_group.rr_set.srv_record`

<a id="canonical-3022131012200113-3123121211100100-0221003010123321-0213303131121032-3122323030123030-0320131123002113-3101323220311231-2330310322133201"></a>

#### `primary.rr_set_group.rr_set.srv_record.name` property

Type: `"string"`. Optional.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-1313232132213222-3323312131222220-1302231230301310-2210311220110031-0003120131301303-3032221203121320-2213022201330021-3223010010123021): complete subsection reference.

<a id="canonical-1313232132213222-3323312131222220-1302231230301310-2210311220110031-0003120131301303-3032221203121320-2213022201330021-3223010010123021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.srv_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-3331101323210103-3033131313122302-0130321211322012-2233211131121323-1203222333302212-2332301313331233-3123010110333331-2023103210211121)
- primary.rr_set_group.rr_set.srv_record.values

<a id="canonical-0330111210022221-0223330313120313-0123322001022203-3102121300120123-2000122203102101-1130011312000333-2023210331301111-1111113130321111"></a>

Type: `"object"`. list nested block, Optional.

SRV Value. Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010332321210210-3301110201001323-1231303223131022-3133233023211320-3223130113323100-1121122020300131-3220332010301120-1203132231231011"></a>

### Direct properties for `primary.rr_set_group.rr_set.srv_record.values`

<a id="canonical-0320200121330033-2102220303010332-2331133001210202-0122123203111103-3221312301130130-1333222312220221-1110032222223301-3010223032320201"></a>

#### `primary.rr_set_group.rr_set.srv_record.values.port` property

Type: `"number"`. Optional.

Port. Port on which the service can be found.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0330033232222012-2031220323210103-0232222212023222-1202232323320001-0113231211002301-3321331300333002-1132313110213332-2303222323230010"></a>

<a id="canonical-3100200121000233-2210102331010232-2101023020230123-0212322302333331-3212000110010230-3322303331332212-0000211012132133-2320211123222001"></a>

#### `primary.rr_set_group.rr_set.srv_record.values.priority` property

Type: `"number"`. Optional.

Priority of the target. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0230233002211220-0231113232232120-0113320001010100-3301022031010221-2213213123333233-3202101320301203-0310330333313131-2122120102222023"></a>

<a id="canonical-0200212333023121-1212111203231223-2312323233230222-2200012023010321-0330320001121122-0220122121230100-0131021223200221-3223301103211130"></a>

#### `primary.rr_set_group.rr_set.srv_record.values.target` property

Type: `"string"`. Optional.

Hostname of the machine providing the service.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-1111203100210131-0200111203203233-3202120003030231-0203121301223000-0202113011300021-3211321033111023-3301200001303300-1312311121000120"></a>

<a id="canonical-1300101200310120-2221121033322233-1220100033003103-0132300010001023-0100032010100112-0310020031121313-0013321212301030-0002022020313103"></a>

#### `primary.rr_set_group.rr_set.srv_record.values.weight` property

Type: `"number"`. Optional.

Weight of the target. A higher number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.sshfp_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.sshfp_record

<a id="canonical-0121121322011211-1303301211230003-0201033312000112-0101223100102212-2102100201221222-1233012320013113-2021333323302131-0213331312123121"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sshfp record.

Additional upstream details:

DNS SSHFP Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
sshfp_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233221330020202-2111121302011233-1310113101030013-1113103313013230-2101110331201312-1321011023312103-3332212230031023-0033211033203333"></a>

### Direct properties for `primary.rr_set_group.rr_set.sshfp_record`

<a id="canonical-1213032112302330-2022300300122131-3030013113030212-0301021202311320-1331110100322013-3112223312103322-1201011233013001-1013022130133033"></a>

#### `primary.rr_set_group.rr_set.sshfp_record.name` property

Type: `"string"`. Optional.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322): complete subsection reference.

<a id="canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.sshfp_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311)
- primary.rr_set_group.rr_set.sshfp_record.values

<a id="canonical-3300100220311322-3102300013023303-1002222232103113-1010021032220123-3021101232031310-0113110031110203-1030020300032022-2311333323323231"></a>

Type: `"object"`. list nested block, Optional.

SSHFP Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("sha1_fingerprint",
    "sha256_fingerprint")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102131103301222-1331101311233230-1320321203131320-0221313310232013-0333322203010231-3332133130231003-2302223211022003-2021321020220202"></a>

### Direct properties for `primary.rr_set_group.rr_set.sshfp_record.values`

<a id="canonical-2002221222101033-0010212011212120-3123322310313233-3331213133311320-1023101333213032-3320100113110003-2022031113231031-1121332202330311"></a>

#### `primary.rr_set_group.rr_set.sshfp_record.values.algorithm` property

Type: `"string"`. Optional.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DSA","ECDSA","Ed25519","Ed448","RSA","UNSPECIFIEDALGORITHM"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](resources--dns_zone--reference--group-003.md#canonical-2323200133101212-1233020310330331-1020330120003130-0301212011321313-0033033220010232-3103312203323111-2203311212112222-2000323133201001): complete subsection reference.

- [sha256_fingerprint](resources--dns_zone--reference--group-003.md#canonical-3103002202322302-2232010213112211-1102001223032112-3101301022320121-0102223210112131-3030021311311121-1133202232131312-2302110021303323): complete subsection reference.

<a id="canonical-2323200133101212-1233020310330331-1020330120003130-0301212011321313-0033033220010232-3103312203323111-2203311212112222-2000323133201001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311)
- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322)
- primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint

<a id="canonical-1130201023023123-3003233001022022-1103002023221203-3120103301012110-2002112000301333-0001111201012033-0332131013022210-3321213103103321"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 fingerprint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
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
sha1_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311120202111130-0331131332133223-2330301013220131-0013231231221123-1333231120313112-0211001323313311-1013112101311021-1110322031003231"></a>

### Direct properties for `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint`

<a id="canonical-3012211102132302-0032303133312012-3112122000300033-3333100023133023-0333322332032103-0010100000331323-1101213233331000-0322311032202022"></a>

#### `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint` property

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-3103002202322302-2232010213112211-1102001223032112-3101301022320121-0102223210112131-3030021311311121-1133202232131312-2302110021303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311)
- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322)
- primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint

<a id="canonical-2100332123103230-2213333330233032-1113231100012320-0211332122323122-0010033302130021-1112322003202312-0013033222301123-1012300030101212"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 fingerprint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
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
sha256_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200203131103221-1003123122312030-0010320203133130-2212002300023001-0003313302112123-0010103101330233-2212102321220321-0313330330131013"></a>

### Direct properties for `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint`

<a id="canonical-1203210033210132-2221003233232103-1230233003103323-3030131030010303-1010112211003312-0013220311321321-0112003033021310-1100203131022201"></a>

#### `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint` property

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-0031230210322331-3321210223301010-0132333332013133-0020032020210333-3021212331022022-3312031103113322-3133100300013111-0030030101213110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.tlsa_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.tlsa_record

<a id="canonical-3310230100220110-2003230230011010-1023111321302213-2230131233321320-3001222102312131-3232032301310331-0320213001312302-2322332212011113"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tlsa record.

Additional upstream details:

DNS TLSA Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
tlsa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311322302310121-3001232133213123-0001303010123322-1131012223021003-1311310012231332-3233031312101003-3312213111220123-3303300113322030"></a>

### Direct properties for `primary.rr_set_group.rr_set.tlsa_record`

<a id="canonical-1311310012330022-1003131133330232-1230203212321112-0012021010333112-2320101211130013-0012211102012103-0022200321311220-0230020023123011"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.name` property

Type: `"string"`. Optional.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-3133001001210321-3322222312020213-2011103303030121-0113113320130302-0332233032210232-0233210203113001-0330002202013030-3223320001022020): complete subsection reference.

<a id="canonical-3133001001210321-3322222312020213-2011103303030121-0113113320130302-0332233032210232-0233210203113001-0330002202013030-3223320001022020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.tlsa_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0031230210322331-3321210223301010-0132333332013133-0020032020210333-3021212331022022-3312031103113322-3133100300013111-0030030101213110)
- primary.rr_set_group.rr_set.tlsa_record.values

<a id="canonical-2211233033003132-3231130220321230-1111303101003111-3321000232232120-1212221101223233-3320200121020020-1232201123102121-0333003133110300"></a>

Type: `"object"`. list nested block, Optional.

TLSA Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_association_data")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210302230302221-0122220321103230-3112130111023123-1111133131303332-3210213221202120-3222310133233011-3111033101221320-0212122221330333"></a>

### Direct properties for `primary.rr_set_group.rr_set.tlsa_record.values`

<a id="canonical-2011321131202002-2100102112300020-3330121221020233-2132110032030031-3202131212312122-0100100011213313-3121100203120303-2020201311021230"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data` property

Type: `"string"`. Optional.

The actual data to be matched given the settings of the other fields.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1011330200330030-2123312103302022-2113100323023323-1302211300022100-1130312210331310-1103233313330103-3322203320322103-2330211330322212"></a>

<a id="canonical-0232120010000322-0121100102100233-3231222201113310-3310233221121301-3233232201301303-3230330022013011-2133011320032201-0010020002200312"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage` property

Type: `"string"`. Optional.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CertificateAuthorityConstraint","DomainIssuedCertificate","ServiceCertificateConstraint","TrustAnchorAssertion"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120103313311223-3311212033320132-2313021132103012-3120322113032010-2210102330331300-1231213031030011-2233131013311220-2333023230333201"></a>

<a id="canonical-3200332022320111-0210302232331112-1002332223131030-2231231321120032-2201032233033202-1211220300001301-3221012033322022-2322103213030031"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.values.matching_type` property

Type: `"string"`. Optional.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["NoHash","SHA256","SHA512"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NoHash",
    "SHA256",
    "SHA512"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1231000011230302-1002130201101212-3023110333020223-3030001121300001-0200012003331321-2103210113120121-0311330322323123-3221133300203330"></a>

<a id="canonical-3100132233202121-3120001020212231-2200322100010313-1121222201002323-1232300231211321-2021230310101121-1203102202113213-3000132113121011"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.values.selector` property

Type: `"string"`. Optional.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["FullCertificate","UseSubjectPublicKey"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("FullCertificate",
    "UseSubjectPublicKey"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1121131111220220-0333230311100110-3220333331113330-3330020320131230-1033022131302013-0333232321103112-3113013123231300-0230232233102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.txt_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.txt_record

<a id="canonical-1100331110322030-0201111130213202-3110023110000301-0003213100033201-1112001001223102-1320003200303132-3000323323020202-0033331320301332"></a>

Type: `"object"`. single nested block, Optional.

DNSTXTResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
txt_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221223201330311-0103102232111101-3030301301313302-3010221031001133-1031313203311020-0211002331010313-1021003111303333-2121221230113013"></a>

### Direct properties for `primary.rr_set_group.rr_set.txt_record`

<a id="canonical-2030201221211031-0012002020130111-1310002001030002-0230223203110013-2302201111203001-2113210032213321-2332032321221200-2223333202010210"></a>

#### `primary.rr_set_group.rr_set.txt_record.name` property

Type: `"string"`. Optional.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-1213202100110122-1000332210102110-2233200221331210-3112130123122210-3131301222211012-2011200123210122-0203022103313330-2032101302022210"></a>

<a id="canonical-3111010303301021-0100333022233020-1132100130323201-1022012100003032-2030311303302223-1210020223313021-2132112223332102-3220300100231030"></a>

#### `primary.rr_set_group.rr_set.txt_record.values` property

Type: `["list", "string"]`. Optional.

Text. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1020202022030312-2130233001011230-0123121101212023-0220103223030203-2021100102202011-2332311013003301-3000013203013302-0121111103000130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.soa_parameters` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.soa_parameters

<a id="canonical-2121220321110032-3211203313030203-0103020311122201-0023320210202020-3121000202111212-1213011301133030-3111130103212201-0011312132222212"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for soa parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("refresh",
    "retry")}
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
soa_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211033221331230-2112210011303023-1023321313012233-3233331001103133-0222020033300023-2210201130323203-0223013311213301-0220202213021103"></a>

### Direct properties for `primary.soa_parameters`

<a id="canonical-3230132300000033-3110101120232103-0102232131022221-0213122210211302-1321012320222201-0122020231231331-3331220311231102-3102322313020311"></a>

#### `primary.soa_parameters.expire` property

Type: `"number"`. Optional.

Expire value indicates when secondary nameservers should stop answering request for this zone if
primary does not respond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-2321101100133013-2133313330023311-0230331332003001-3012220301332222-0102123320302303-1312030230112333-3112320201003201-1330311010030021"></a>

<a id="canonical-3221210121321221-3323221123013020-2131000200200021-0022320023001120-1132012221200200-1002031102300101-0221003110203310-3030102322210223"></a>

#### `primary.soa_parameters.negative_ttl` property

Type: `"number"`. Optional.

Negative TTL value indicates how long to cache non-existent resource record for this zone.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-2312322201230120-2102133302301003-1020333031120112-3232213210311122-3322000112101202-3011312002030010-3001102333031203-2200333102220312"></a>

<a id="canonical-2202110132121202-0333303330312000-2012001033122232-3000000121300111-2001032203303100-0121330030212000-3211033210103032-0023312001011001"></a>

#### `primary.soa_parameters.refresh` property

Type: `"number"`. Optional.

Refresh value indicates when secondary nameservers should query for the SOA record to detect zone
changes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(3600, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-0132223313310300-3002321121310323-1322333201102312-3002131301301323-2310010230213030-1221131321220023-0222013132031032-0201101302133211"></a>

<a id="canonical-1030301223021301-2201200022322003-3000100030132213-0121002332112301-1302220130110230-2300302232112230-2331311032301333-1320122223332023"></a>

#### `primary.soa_parameters.retry` property

Type: `"number"`. Optional.

Retry value indicates when secondary nameservers should retry to request the serial number if
primary does not respond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3031121330002222-1122010031033213-2002203113122001-2220320201322012-3303123323323323-0312201022120333-3312123322122323-2113131330202000"></a>

<a id="canonical-3301232330013020-2011231320002332-2311200201012223-2112130110212333-0003122211002210-2333311122212221-3022313212012332-2003313231021031"></a>

#### `primary.soa_parameters.ttl` property

Type: `"number"`. Optional.

TTL. SOA record time to live (in seconds)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `secondary` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- secondary

<a id="canonical-0223010131102331-0310121320000300-3131133100321303-1312200031031101-2101222213223300-3110301303132223-1023132100303123-3012223223200121"></a>

Type: `"object"`. single nested block, Optional.

SecondaryDNSCreateSpecType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_servers")}
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
secondary {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222013133003013-1021032013210130-1110302021230000-2112222133102000-3131023110210030-1101203003011232-1303222202332201-2020102021010000"></a>

### Direct properties for `secondary`

<a id="canonical-3022322020202002-2320120232202130-1332110110120132-3231131222233013-3023023031131313-2000321032320123-0222022323233201-0223032300003022"></a>

#### `secondary.primary_servers` property

Type: `["list", "string"]`. Optional.

Configuration parameter for primary servers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131221012001030-0213220100322000-0121123330330200-3122302010233231-3322322030130132-1232101300233203-2022220033103233-2112112211323031"></a>

<a id="canonical-3320121013310121-0010123111303231-0030000231021102-1213221331320010-2112311110310012-3333333213022031-2302001033303322-0230120230011003"></a>

#### `secondary.tsig_key_algorithm` property

Type: `"string"`. Optional.

\[Enum: HMAC\_MD5|UNDEFINED|HMAC\_SHA1|HMAC\_SHA224|HMAC\_SHA256|HMAC\_SHA384|HMAC\_SHA512\] TSIG
key-value must be compatible with the specified algorithm - UNDEFINED: UNDEFINED - HMAC\_MD5:
HMAC\_MD5 - HMAC\_SHA1: HMAC\_SHA1 - HMAC\_SHA224: HMAC\_SHA224 - HMAC\_SHA256: HMAC\_SHA256 -
HMAC\_SHA384: HMAC\_SHA384 - HMAC\_SHA512: HMAC\_SHA512. Possible values are \`HMAC\_MD5\`,
\`UNDEFINED\`, \`HMAC\_SHA1\`, \`HMAC\_SHA224\`, \`HMAC\_SHA256\`, \`HMAC\_SHA384\`,
\`HMAC\_SHA512\`. Defaults to \`UNDEFINED\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["HMAC_MD5","HMAC_SHA1","HMAC_SHA224","HMAC_SHA256","HMAC_SHA384","HMAC_SHA512","UNDEFINED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNDEFINED",
  "enum": [
    "HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1312113303301123-0131313203312122-2000203020002221-2012011313232100-0302313003131321-3031110232312212-3300111032232002-3020122021101320"></a>

<a id="canonical-2003000301233323-0123110111021213-1301133331333321-3002100130301123-2100213103310132-0022121103222110-1100131031332213-3321322302322122"></a>

#### `secondary.tsig_key_name` property

Type: `"string"`. Optional.

TSIG key name as used in TSIG protocol extension.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013): complete subsection reference.

<a id="canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `secondary.tsig_key_value` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021)
- secondary.tsig_key_value

<a id="canonical-1133331300221133-0123021301132223-1331302203102230-0220012222331103-1201221303001002-3103010003311200-3002102213203003-3012013312111321"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
tsig_key_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320033122100010-3011202300022113-1031220231312233-0201031130303002-0003230202311002-1012312313121132-2023130000221003-2121010000330310"></a>

### Direct properties for `secondary.tsig_key_value`

- [blindfold_secret_info](resources--dns_zone--reference--group-003.md#canonical-0021202031221212-3320231031003200-0031123320023132-3131020332100030-1003112133001300-1112332332210033-1323202102222310-2123013130232330): complete subsection reference.

- [clear_secret_info](resources--dns_zone--reference--group-003.md#canonical-3302323322210222-2331232232101310-2233022013321120-2000233221232030-3211302223220000-2103131213212123-2121333200012003-2133032302002300): complete subsection reference.

<a id="canonical-0021202031221212-3320231031003200-0031123320023132-3131020332100030-1003112133001300-1112332332210033-1323202102222310-2123013130232330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `secondary.tsig_key_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021)
- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013)
- secondary.tsig_key_value.blindfold_secret_info

<a id="canonical-3301012122133333-0110202122221123-0220023330200102-2000211100012300-0130332230302202-1322101030213033-2113110313221230-1332020312120123"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310112332030222-1120111103011033-2021021013321031-3310202102213112-2001223120300023-3202301300211122-1213020301330232-1330302130000011"></a>

### Direct properties for `secondary.tsig_key_value.blindfold_secret_info`

<a id="canonical-2033213322032312-0311201200230231-2001101223023201-1330222303203323-0301121110300103-1313212221133112-0100323011301232-3132213223101233"></a>

#### `secondary.tsig_key_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3023330002313330-0310233121031302-3013213121310201-3103011010012001-1331201301011322-0131330133132222-2031020230021022-1303333022323331"></a>

<a id="canonical-1233232233123313-1111322232221213-2211231132310331-1102301331323030-3103102233330010-1211211302133101-1311021022110012-0013301333022230"></a>

#### `secondary.tsig_key_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0231213330033102-3003110201202211-0202013213022303-0023000031003200-2101223332313210-0102312301311012-0013003223003233-0112100130310201"></a>

<a id="canonical-0133013332322201-3310112100231232-3100330202102301-1133220232012031-1000203323011313-2320221130210033-3201120023333220-0321320210213333"></a>

#### `secondary.tsig_key_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3302323322210222-2331232232101310-2233022013321120-2000233221232030-3211302223220000-2103131213212123-2121333200012003-2133032302002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `secondary.tsig_key_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021)
- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013)
- secondary.tsig_key_value.clear_secret_info

<a id="canonical-2023132111012003-0311221323013101-1122000230223032-2103123200220102-1023020011231121-3211122310021031-3232230201032312-0013130200310230"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321112001020332-1010001233221131-2332233233132002-1301222113313311-1130221222302330-0231010220302101-3131321121021301-1020300132102333"></a>

### Direct properties for `secondary.tsig_key_value.clear_secret_info`

<a id="canonical-0313132220201333-2033222033331013-0030330303311112-0322211020220301-0102331021000110-3213030222322230-3021000030321003-0103110133023013"></a>

#### `secondary.tsig_key_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2222023222322022-1101121232133022-3110222320132113-3212123330200313-3031101201201313-0210120111300132-2003201020113203-2113012332131101"></a>

<a id="canonical-1132110110112221-1011020333010120-0312213330311223-0001112030120131-0012033313023102-2120100103133223-0122003231123301-3102320132211313"></a>

#### `secondary.tsig_key_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1121023120323203-0300301033303311-3210213210020103-1313032000222330-2030312122211110-2210030103233231-0231221302231331-3010100002112333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- timeouts

<a id="canonical-0300320321333023-0310332010031331-3031222113133311-3100030302022003-3212120202021302-0032200302202012-1100200031002133-3222120010110101"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213302201023132-3320101001313231-1030133033302300-0303133313122032-3123312232032003-3220113030212102-1103211310223232-1332212013313002"></a>

### Direct properties for `timeouts`

<a id="canonical-0300131231131222-2103011332202020-0301112203313120-0211010303001323-2310331101022332-1211200311113120-0121002013210222-0231323222020212"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3321101022212312-1233211320310103-0013133032000303-0231010331321222-0213012223111311-0323213230221330-1133220321110110-2203021323101221"></a>

<a id="canonical-3020022210333132-0301020230022032-1101103330231130-1332003021112331-1301320000330003-1301122223010130-0332121130331111-1232001223323301"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0321030200103001-1021031232022232-1030103103022221-2103022311101312-3111010301003211-1221312222231132-3121023010300312-3023202012113020"></a>

<a id="canonical-3323013313320030-3130123230322232-3002132200311121-2100132010303210-0022212313001011-1211213321032123-3320022320312200-3121210222010323"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3323010230210303-0302232133303033-0311231221113031-1231331231312321-3203133220303021-3331110110321332-3232112303121310-3220332113032321"></a>

<a id="canonical-0332220011101031-3202002103220220-1202131101030223-2000102313020210-3332102222210230-0313020302120012-0133100300010010-0032201320102010"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
