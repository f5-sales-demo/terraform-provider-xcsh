---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-0222213320032200-0031210310103201-3122323213320302-0131233032302201-3100121200310223-2323002322301133-0231300123003121-3201332000130113"></a>

## `primary.default_rr_set_group.loc_record.values.latitude_second` property

Type: `"number"`. Computed.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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

<a id="canonical-1332020110030220-0110323332033122-0211221103213133-2132023210320332-3200130321302230-3113133201030012-0201213332203223-1203223113202330"></a>

<a id="canonical-1331130020302210-3113111233113203-1103030300310221-3312030023012010-0103333112130032-0012112123002122-3302112200011330-2110201200113100"></a>

## `primary.default_rr_set_group.loc_record.values.location_diameter` property

Type: `"number"`. Computed.

Diameter of a sphere enclosing the described entity, in meters.

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

<a id="canonical-3120232123033331-0300133111133213-3111211323203213-3310210223031100-1131111302113223-0211120310003212-3032312321123012-0002301303031332"></a>

<a id="canonical-1301002030313102-0303103030120011-2301211113312302-3232220021232320-0330210002201321-0100301213332122-2320222331112302-0333300010032012"></a>

## `primary.default_rr_set_group.loc_record.values.longitude_degree` property

Type: `"number"`. Computed.

Longitude degree, an integer between 0 and 180, including 0 and 180.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2100132230332122-0112003022133101-1111223133113332-1113333203031030-0011231300022013-1210110123230003-3321212000222003-3220101322211323"></a>

<a id="canonical-3000232211302310-0230012020101301-0200331102221000-2321023000102023-1211032231201111-3321213133210132-2020130331101033-2123230020101232"></a>

## `primary.default_rr_set_group.loc_record.values.longitude_hemisphere` property

Type: `"string"`. Computed.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

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

<a id="canonical-2023123133322101-3220201332030302-2223300002022200-1211012232222333-3213230220333023-0101130013312211-3132312322023222-1133300303332100"></a>

<a id="canonical-2020110220123102-3003211330323201-0032033230201323-3103231012302112-2333303310213200-3211310120103001-1332113320223032-2333323220312013"></a>

## `primary.default_rr_set_group.loc_record.values.longitude_minute` property

Type: `"number"`. Computed.

Longitude minute, an integer between 0 and 59, including 0 and 59.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0130330210323301-2302302203032002-3013230121123223-2012002223231113-3003312111101111-1133012123001123-1001222022213200-3233112113322123"></a>

<a id="canonical-1311010023102033-0110233231322100-3003111010003111-2100021310302312-0031232320302020-0002302322111333-1130123310312122-3302303230122112"></a>

## `primary.default_rr_set_group.loc_record.values.longitude_second` property

Type: `"number"`. Computed.

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

<a id="canonical-2111033231212233-2011201101302003-2221003230122121-1013131311032202-0001230011101203-3113101020101303-1213112220312030-3331012100003310"></a>

<a id="canonical-1200231221033210-1122132231201110-0223321231203031-3222323113303001-2311233122212130-3311012103133203-3201230230222132-0010220022020323"></a>

## `primary.default_rr_set_group.loc_record.values.vertical_precision` property

Type: `"number"`. Computed.

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

<a id="canonical-1210121201311120-2131010222033122-0330222031322111-3200022000122302-0333021032123112-2220203201001212-2120030320310132-2322032202111212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.mx_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.mx_record

<a id="canonical-1032311100202302-3103323103112013-1333010002323213-2002133011303012-0230220002213110-1003213032321220-1130322103013231-3303203111301121"></a>

Type: `"single"`. Computed.

DNSMXResourceRecord.

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

<a id="canonical-3210020110023020-1002001331011011-0111123221022300-2021313211130221-3331120003230321-2312100233203021-0021133013213330-1102201022103013"></a>

### Direct properties for `primary.default_rr_set_group.mx_record`

<a id="canonical-0333112301302200-3132212312222130-0022323331310010-2303001231202322-0313332131233111-2202322212320001-0012222330120330-2020010020321221"></a>

#### `primary.default_rr_set_group.mx_record.name` property

Type: `"string"`. Computed.

MX Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-1011130322331212-0301101323322001-0320200330211331-2122132002032322-1003202023031203-2021031220230300-1212030120033231-0033332212110320): complete subsection reference.

<a id="canonical-1011130322331212-0301101323322001-0320200330211331-2122132002032322-1003202023031203-2021031220230300-1212030120033231-0033332212110320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.mx_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-1210121201311120-2131010222033122-0330222031322111-3200022000122302-0333021032123112-2220203201001212-2120030320310132-2322032202111212)
- primary.default_rr_set_group.mx_record.values

<a id="canonical-3233232323213020-1102033211231011-1321202112322120-0333320230101121-0320302103020112-1331300012031212-0022010320131022-3001202130130233"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3312032230300023-1301133310102232-3330311222023331-2131330122122203-1121200223120310-2102012210213112-2022213010211313-0202230122300212"></a>

### Direct properties for `primary.default_rr_set_group.mx_record.values`

<a id="canonical-3212233103010021-2112330330101021-3311113123112320-3210112120133231-1123310220120112-1301333133002022-2113311003233020-1203122111002323"></a>

#### `primary.default_rr_set_group.mx_record.values.domain` property

Type: `"string"`. Computed.

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3322212023020301-2102010023101223-1120101300130022-3000311103023230-3220132000213212-3210333213332211-3112113323201011-3332223233032202"></a>

<a id="canonical-3313322000200103-2310320322133022-2032013302100223-3020210022120123-1310233033320333-3332123231312031-1112211313001313-1223002300203232"></a>

#### `primary.default_rr_set_group.mx_record.values.priority` property

Type: `"number"`. Computed.

Priority. Mail exchanger priority code.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2000221213012233-1222212300001003-2201121122210020-3133301203202102-1332032202320301-1003102301122213-2203320000301213-0012113303201103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.naptr_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.naptr_record

<a id="canonical-1131003222002322-0331310113232231-0120321102333300-2000203102120231-0302103021322220-3232023223211033-2330023121211213-2302003331022033"></a>

Type: `"single"`. Computed.

Configuration parameter for naptr record.

Additional upstream details:

DNS NAPTR Record.

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

<a id="canonical-0133220221222203-1303310310332011-1212101231112031-2003032212321302-0312120000302100-0103330001231012-2123300130331330-3030221210220120"></a>

### Direct properties for `primary.default_rr_set_group.naptr_record`

<a id="canonical-1310020030321333-0101103311331302-2303012331310110-3003213333031333-1013020301321101-1322312211323103-3013310130323002-3013132232123002"></a>

#### `primary.default_rr_set_group.naptr_record.name` property

Type: `"string"`. Computed.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-0100201121020123-1031113322131133-1011321031132003-3321103133212313-3031002121201312-0312231201330322-1033220103223301-1023211031120113): complete subsection reference.

<a id="canonical-0100201121020123-1031113322131133-1011321031132003-3321103133212313-3031002121201312-0312231201330322-1033220103223301-1023211031120113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.naptr_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-2000221213012233-1222212300001003-2201121122210020-3133301203202102-1332032202320301-1003102301122213-2203320000301213-0012113303201103)
- primary.default_rr_set_group.naptr_record.values

<a id="canonical-2320200022012033-1033033032110021-0102202303021210-3030031212130021-0300133300102312-2203100103330133-2300012122032212-1100023312311020"></a>

Type: `"list"`. Computed.

NAPTR Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3022202101033110-0001013110121023-1120303322311032-0202332120303103-0132203000012211-3331300212301323-3310331133300120-2323123110230100"></a>

### Direct properties for `primary.default_rr_set_group.naptr_record.values`

<a id="canonical-3220320022000021-1233032000023232-0021330031311101-1003210020112310-0111322332300202-2330101220201211-0131110023033311-2302310101333232"></a>

#### `primary.default_rr_set_group.naptr_record.values.flags` property

Type: `"string"`. Computed.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0300031121323202-3023112232120133-0130222230303332-0302000303203010-0010113133012022-2012120103012321-2323100221100333-2202023212021101"></a>

<a id="canonical-0111332101221002-3021123011122320-2101133120300020-0032102013303332-3321312213302201-1212302301332100-2311012333011323-1030233320012031"></a>

#### `primary.default_rr_set_group.naptr_record.values.order` property

Type: `"number"`. Computed.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2011021203322230-2231013332301131-3132101230023310-3023000312210330-1333030302112031-2210123203222133-0111202001321110-2302302111210113"></a>

<a id="canonical-0021112211010312-2323330213210023-0332221031300130-0312200331023210-3111013120333220-1111002102103232-2330003231032233-1102212120120001"></a>

#### `primary.default_rr_set_group.naptr_record.values.preference` property

Type: `"number"`. Computed.

Preference when records have the same order. A lower number indicates a higher preference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3330121203020203-0221132100100003-2002333210221223-1311220321133311-3323323311333033-2122031012210132-1223102033211211-3032111300330202"></a>

<a id="canonical-3132223202032331-2002130020223213-2300221210103120-1220020010300100-2032311001330232-0000010200231010-2012331012111231-1223313031312211"></a>

#### `primary.default_rr_set_group.naptr_record.values.regexp` property

Type: `"string"`. Computed.

Regular expression to construct the next domain name to lookup.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3033113003322110-2210300231010000-0003313223021212-0212010220101222-1000113022332113-2210013200131230-1312001131132203-0331103032022122"></a>

<a id="canonical-3213130012332013-3013232113222102-2021231230120111-1311210220012011-0211332012001201-2211213322031211-3320302130223321-3222331030001100"></a>

#### `primary.default_rr_set_group.naptr_record.values.replacement` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1103001310312002-2033033120313313-3301110213102322-1200103113111121-0111201223100233-1210011003112100-3331010212033110-2111113103012012"></a>

<a id="canonical-1321122000102112-0011000121321122-2121131003320303-0222102032030303-3110131000201322-0122213013313231-2121310100132333-1322213311210122"></a>

#### `primary.default_rr_set_group.naptr_record.values.service` property

Type: `"string"`. Computed.

Specifies the service(s) available down this rewrite path.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2313231112212011-0201330312333230-1201033232031120-1123102230132211-2311130313033231-0220202113331330-3030032011212201-2022201023220103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.ns_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.ns_record

<a id="canonical-3120011311313020-0111201012211122-2213331000100130-1232101330013230-0030203001332200-2000311000020320-3311002011130312-1322011010013123"></a>

Type: `"single"`. Computed.

DNSNSResourceRecord.

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

<a id="canonical-1021003300030212-2213203212320133-0323010212313030-3312001122120303-0122212011332132-3123203223110001-3111033220210220-0300332301230133"></a>

### Direct properties for `primary.default_rr_set_group.ns_record`

<a id="canonical-2213320012022313-1021103303322033-0111031323202112-3312112010303203-3011210120120320-2230203102200111-2211022130202111-3010320210120211"></a>

#### `primary.default_rr_set_group.ns_record.name` property

Type: `"string"`. Computed.

NS Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0300120203320102-3303321332223332-0323100011233013-3230312112133013-3101333233332231-1033102102210322-3132203032100032-3220031211032023"></a>

<a id="canonical-1031030212322332-2223221201210333-2312121122113212-1121012300312011-0212013110003201-3032312310323221-2200012300303002-1123323013323230"></a>

#### `primary.default_rr_set_group.ns_record.values` property

Type: `["list", "string"]`. Computed.

Name Servers. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0032230300312201-2330203102230021-2122122330231113-1232330231302310-3001020022321003-0100000031302303-0320002132012223-3332123310323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.ptr_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.ptr_record

<a id="canonical-3303011132132001-3122103213121003-3311211220312210-1111031013000031-3212320011003331-3323012020013020-0120110300202132-1012322330233103"></a>

Type: `"single"`. Computed.

DNSPTRResourceRecord.

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

<a id="canonical-2301302133020033-3203320311100113-1222133021121003-1211200322001303-3331203323211132-0032120322123121-3223020011031230-1212031210023330"></a>

### Direct properties for `primary.default_rr_set_group.ptr_record`

<a id="canonical-1023002231030203-1111200200202010-1320211303302021-0033222220111300-1011203222322021-1122221223213310-0210233000330313-0013022333303210"></a>

#### `primary.default_rr_set_group.ptr_record.name` property

Type: `"string"`. Computed.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2313101211020213-0201123002020023-1113010113221221-2301313211102232-3200202311133321-3332021221031222-0321110113201202-0200112033033110"></a>

<a id="canonical-2213012211303220-2002200011123231-0203230000020231-1033122211020031-2102202323332330-1103322212213330-2030022021022031-0330101230232001"></a>

#### `primary.default_rr_set_group.ptr_record.values` property

Type: `["list", "string"]`. Computed.

Domain Name. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3022002303322101-2133310220222102-1221303323002002-2021320313022100-3022301033201311-3100301120122232-0121032103302112-2011211020313032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.srv_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.srv_record

<a id="canonical-2001102313121111-1031012001223021-1320032301003300-1221012201112033-0210231131333011-1223133113110100-3122001233211300-1233200332020001"></a>

Type: `"single"`. Computed.

DNSSRVResourceRecord.

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

<a id="canonical-2201103111222310-2010311122220201-2021132213110313-1120010132103012-2132020321312110-1130203310022111-3300233000000222-3303311220120023"></a>

### Direct properties for `primary.default_rr_set_group.srv_record`

<a id="canonical-0323130100321031-0322103030031330-0332231023332230-2002133130023113-3111213100333233-3133102030110313-3120321101233323-1223230200201300"></a>

#### `primary.default_rr_set_group.srv_record.name` property

Type: `"string"`. Computed.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-1320223020133013-1023223333033030-3122310013013203-0033013220230323-1203220300022200-0213310103002010-0022011222300321-0201010301013210): complete subsection reference.

<a id="canonical-1320223020133013-1023223333033030-3122310013013203-0033013220230323-1203220300022200-0213310103002010-0022011222300321-0201010301013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.srv_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-3022002303322101-2133310220222102-1221303323002002-2021320313022100-3022301033201311-3100301120122232-0121032103302112-2011211020313032)
- primary.default_rr_set_group.srv_record.values

<a id="canonical-3011201211313333-0132121313130231-1223101003211111-3213221203211211-2320300231031112-1033212012000030-3311101111020022-3113122301120032"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3002213123120032-3313100203112220-0023230300130020-2031132222221233-1123132233301232-1023312200310310-2233013110201020-1013210002033203"></a>

### Direct properties for `primary.default_rr_set_group.srv_record.values`

<a id="canonical-0001332202013201-2121012121222022-1020110233101232-1311321312220111-0023022022232021-0130000111100131-3221031321133220-0211320210103302"></a>

#### `primary.default_rr_set_group.srv_record.values.port` property

Type: `"number"`. Computed.

Port. Port on which the service can be found.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0203121331131200-1221202321111221-2331131303003330-3232111021110213-3010231102103002-0003330311113320-1330012010132020-2320212112122011"></a>

<a id="canonical-3222111213011132-3023010320000122-2323232212112103-0113102122013221-2121132311332323-3033012131123213-1102303102212113-1331310002300331"></a>

#### `primary.default_rr_set_group.srv_record.values.priority` property

Type: `"number"`. Computed.

Priority of the target. A lower number indicates a higher preference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0130212300322010-0021212210213323-3301110210322211-2131222211120020-2303212202032322-0322020110111220-2311332113310212-1303202022020203"></a>

<a id="canonical-1020222122221103-3130021203331100-2232021221200200-0201012333301312-1023212012233303-2333202013113122-3033033111331222-1212331130012303"></a>

#### `primary.default_rr_set_group.srv_record.values.target` property

Type: `"string"`. Computed.

Hostname of the machine providing the service.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1231123310210022-3110320103030330-1033213211212230-1102120031210331-2123020301223321-2202023200133222-2023030101310323-3302223332123313"></a>

<a id="canonical-1021012231230112-1001203333322321-3331223012223120-0332020223323000-0301333332023023-1001330003323032-1012313322121111-1121200332323101"></a>

#### `primary.default_rr_set_group.srv_record.values.weight` property

Type: `"number"`. Computed.

Weight of the target. A higher number indicates a higher preference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.sshfp_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.sshfp_record

<a id="canonical-2013012331332022-3033213302020033-1223002311330010-0102020002302130-2100313302123103-2003323030131310-2232212330323011-3212033020230212"></a>

Type: `"single"`. Computed.

Configuration parameter for sshfp record.

Additional upstream details:

DNS SSHFP Record.

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

<a id="canonical-2230132232203013-2000030110130320-3021013220020303-3030322231103033-1003222101320223-3120013331302030-0130322130231211-3310320213302213"></a>

### Direct properties for `primary.default_rr_set_group.sshfp_record`

<a id="canonical-1023122132223330-2001321301031200-1030312232001233-0201023130202100-2122322002301000-3102310331020333-3320013132232223-3030021323120222"></a>

#### `primary.default_rr_set_group.sshfp_record.name` property

Type: `"string"`. Computed.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103): complete subsection reference.

<a id="canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.sshfp_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021)
- primary.default_rr_set_group.sshfp_record.values

<a id="canonical-0002031333331122-1032020130110201-3102202233202031-1002322002200203-2030133203023222-3020310021213033-2031110213101210-2130013011222231"></a>

Type: `"list"`. Computed.

SSHFP Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3011022012313231-2130311020210332-0312002110330223-3110203022322323-2300323330021322-1223330221233300-0301320120302122-3112123022012011"></a>

### Direct properties for `primary.default_rr_set_group.sshfp_record.values`

<a id="canonical-0303030022003011-3020231322301030-0330110112122301-0330022000303231-2203332223112112-3201230320313201-1023202102110220-1001333003113030"></a>

#### `primary.default_rr_set_group.sshfp_record.values.algorithm` property

Type: `"string"`. Computed.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

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

- [sha1_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-3212233132102313-3020030030330330-1300213312321310-0320003102123111-2030132233231132-3223023333221311-2312331123120223-2012333000021131): complete subsection reference.

- [sha256_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-1033003133000023-0122010033130000-1001333131130010-2300232320011323-1020123031313230-0023020332032322-0022313002020000-2021012001232012): complete subsection reference.

<a id="canonical-3212233132102313-3020030030330330-1300213312321310-0320003102123111-2030132233231132-3223023333221311-2312331123120223-2012333000021131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021)
- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103)
- primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint

<a id="canonical-2301022221023201-1013221230312100-0313330033313231-3210020032122202-2321203113211310-2021331300331312-0010123010032000-2212302103211002"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 fingerprint.

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

<a id="canonical-2000220233310032-3231230321002033-1113113313322121-2123332320130333-2210233231220220-2320320222033321-1132132132120213-1211321103000130"></a>

### Direct properties for `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint`

<a id="canonical-2202012233001101-3131012212320120-2212030130212331-2331310220220331-0001230232101201-1221122111131023-0221201300130200-3003311132232222"></a>

#### `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint` property

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1033003133000023-0122010033130000-1001333131130010-2300232320011323-1020123031313230-0023020332032322-0022313002020000-2021012001232012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021)
- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103)
- primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint

<a id="canonical-1201210333203213-2223231311222211-0201320320001023-3110022121030130-2302312101110210-0301333010330033-0230022230112212-2210011303031321"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 fingerprint.

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

<a id="canonical-1222213300033130-0333232103332310-1031100110222221-0111133331202103-2032133203313020-3032131020020120-2321232123330001-2301220010233210"></a>

### Direct properties for `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint`

<a id="canonical-1113020303311333-0132332222011212-0232212100333233-3110110303222112-2023213310022210-2300321230330113-2222032010303302-0030102310221231"></a>

#### `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint` property

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2101121000001100-2301010023202222-0201001332133023-3100101303202310-3001221233210201-2132313122010110-3100111130011132-3022131111200230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.tlsa_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.tlsa_record

<a id="canonical-0032202110310221-3032213010113022-0333121311132030-3203010222233313-0303021332131032-3222131113020221-0011020030321203-0301021200311302"></a>

Type: `"single"`. Computed.

Configuration parameter for tlsa record.

Additional upstream details:

DNS TLSA Record.

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

<a id="canonical-0303333331303113-3201001233200302-2121030232203133-0311022231001321-3210000013033123-1022311332001210-2233030011330301-3331200023210130"></a>

### Direct properties for `primary.default_rr_set_group.tlsa_record`

<a id="canonical-1230213033331222-0330232100013210-3213221211033210-3211223220332020-1101010120033211-0233030231000310-0212212300203223-2210211111022110"></a>

#### `primary.default_rr_set_group.tlsa_record.name` property

Type: `"string"`. Computed.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-0120011031211132-3230333303020232-1031031230021310-0113032230023310-3311022011210212-1033010011101331-3112200132321022-2112113100132111): complete subsection reference.

<a id="canonical-0120011031211132-3230333303020232-1031031230021310-0113032230023310-3311022011210212-1033010011101331-3112200132321022-2112113100132111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.tlsa_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-2101121000001100-2301010023202222-0201001332133023-3100101303202310-3001221233210201-2132313122010110-3100111130011132-3022131111200230)
- primary.default_rr_set_group.tlsa_record.values

<a id="canonical-3002001033031230-0223021112020031-2031230013011312-2212321010023312-2130233001011032-3333311220331311-2220322001301203-3031032102121020"></a>

Type: `"list"`. Computed.

TLSA Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3320302201211120-1301023202221322-1202121110221232-3330220331331212-1112201032130123-1131203123333011-2111120133321001-2112300120130103"></a>

### Direct properties for `primary.default_rr_set_group.tlsa_record.values`

<a id="canonical-0101102321003212-3113212202302201-2110303021101311-3303012000332311-0311220231312102-0132122221022212-0211013111023301-1111100322111200"></a>

#### `primary.default_rr_set_group.tlsa_record.values.certificate_association_data` property

Type: `"string"`. Computed.

The actual data to be matched given the settings of the other fields.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0123210230200102-3213021213222203-2002331032101201-1123001303221233-3323100202113322-3130212113311100-3033330032301210-1231012020213333"></a>

<a id="canonical-0000232320323323-1200113311011323-0303332103001300-0130021302113312-1000223201020112-0331313131022010-0111133020313032-1231133120032331"></a>

#### `primary.default_rr_set_group.tlsa_record.values.certificate_usage` property

Type: `"string"`. Computed.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

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

<a id="canonical-1123030022332000-1322302232022012-3221001233221210-3233322123131031-2113030131111212-1200000322300232-2221233112232322-1131020233000320"></a>

<a id="canonical-2133211030333212-1003222000333211-3012212332201300-0200103312032323-2211110003233223-0002002202033133-2110022213032303-0011220322033122"></a>

#### `primary.default_rr_set_group.tlsa_record.values.matching_type` property

Type: `"string"`. Computed.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

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

<a id="canonical-1212311331323330-3121101311123002-3322032023130310-1333000010211313-0211023201030232-2111022020011002-3320313123203303-1302133303211202"></a>

<a id="canonical-3302121233011203-0103223101303320-0201112323002320-1110022101130022-3330213003010202-2111302011133321-1220230221131011-3321123122222001"></a>

#### `primary.default_rr_set_group.tlsa_record.values.selector` property

Type: `"string"`. Computed.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

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

<a id="canonical-2023003321332312-0023132123210230-1102330021223031-3112003211020100-3031102022010133-0321133333232212-0313320223232033-0232003031030313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.txt_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.txt_record

<a id="canonical-0113213300213103-1012321032121210-0320323211322311-2113133002212131-2200322100201330-2001332233103233-3303333311332233-0121112232302110"></a>

Type: `"single"`. Computed.

DNSTXTResourceRecord.

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

<a id="canonical-3112300033121301-0320131210013123-0213330031301222-3111301311122323-2232023211110233-3202321333221200-3221300132122203-1232022033300100"></a>

### Direct properties for `primary.default_rr_set_group.txt_record`

<a id="canonical-3132322012012100-2213013323323121-2113100213332123-3223330323222110-1300133212023222-0012333100203032-3211111032012332-2333213122021123"></a>

#### `primary.default_rr_set_group.txt_record.name` property

Type: `"string"`. Computed.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0232301130321032-1123011101323311-1103103322303022-1212000212321333-0213003222303132-3211223133203202-1111111310212333-3023300210011323"></a>

<a id="canonical-0332013112220110-1323320031233100-1323221302012131-2332101022322011-2311332120130013-1031230333333002-3323002232330101-3121210013231312"></a>

#### `primary.default_rr_set_group.txt_record.values` property

Type: `["list", "string"]`. Computed.

Text. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1013032121121211-0200310230111302-2302123231122131-1132011211233131-3233303200303110-1001330123313000-0122110323001111-2131223301100220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_soa_parameters` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.default_soa_parameters

<a id="canonical-1200102321322311-3223202022210202-3010323322310100-2013003011231200-1032331233311132-0021012320100102-3113332223310121-0101330133331122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default soa parameters.

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

<a id="canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.dnssec_mode` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.dnssec_mode

<a id="canonical-1003323000301330-1322020032333013-1121022131303132-2311101012002202-0130133003222100-0133113113112003-2121213120121032-2021301010121112"></a>

Type: `"single"`. Computed.

DNSSEC Mode.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-0212101102011133-2120120300102130-3122112021211131-0320020203220010-2300031321111022-3302300322210333-2123232222220302-3110320211233312"></a>

### Direct properties for `primary.dnssec_mode`

- [disable_spec](data-sources--dns_zone--reference--group-002.md#canonical-0203303200332322-0121103330333203-3311110011232221-2101231000322133-2102222111020223-3132302232302210-1213103200010130-0300321120212112): complete subsection reference.

- [enable](data-sources--dns_zone--reference--group-002.md#canonical-1221121130033221-2123222101202023-0201332012131001-0303022301203123-0122210132133332-2203321123031100-2100301320102120-1203023103200100): complete subsection reference.

<a id="canonical-0203303200332322-0121103330333203-3311110011232221-2101231000322133-2102222111020223-3132302232302210-1213103200010130-0300321120212112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.dnssec_mode.disable_spec` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102)
- primary.dnssec_mode.disable_spec

<a id="canonical-0133101022110312-1012230221200332-0112223131330022-3122233333232332-0303220302221003-1230223103002022-3112210203012202-3333110013323211"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221121130033221-2123222101202023-0201332012131001-0303022301203123-0122210132133332-2203321123031100-2100301320102120-1203023103200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.dnssec_mode.enable` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102)
- primary.dnssec_mode.enable

<a id="canonical-1021202202023231-1203202332111212-0003030323332333-2100202023101323-3223302330203000-3231022021132313-1010101012320311-0033321331223221"></a>

Type: `["object", {}]`. Computed.

Enable. DNSSEC enable.

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

<a id="canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.rr_set_group

<a id="canonical-2231202303211010-2021101001111030-2013023101223233-3321213130020213-3332320300011023-2121211132011013-2312031132210200-1131031002133031"></a>

Type: `"list"`. Computed.

Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed
by F5.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2321302230002300-0123313012322121-0231023323030011-2300222332033323-0101121133201303-2033102022133111-3232121332113200-1330230102100220"></a>

### Direct properties for `primary.rr_set_group`

- [metadata](data-sources--dns_zone--reference--group-002.md#canonical-2300002332200332-3300012320121033-1323312230331320-0222210111132311-1030011331333111-1212002231231222-1321113113133330-0222303310223013): complete subsection reference.

- [rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302): complete subsection reference.

<a id="canonical-2300002332200332-3300012320121033-1323312230331320-0222210111132311-1030011331333111-1212002231231222-1321113113133330-0222303310223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.metadata` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- primary.rr_set_group.metadata

<a id="canonical-1330003303230112-2023000311003232-0313021100232031-2130100211300330-1322102200310132-0111211310022101-3331032000001001-1003333223010012"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-1010113313101123-3113222020321022-2103203033123130-3110301311311222-3211213133003111-2121331132132210-1031223303312000-2111120000221023"></a>

### Direct properties for `primary.rr_set_group.metadata`

<a id="canonical-1322330123203331-2320210002222202-3221301311020012-2330303222321212-1031011322100311-2110232000022102-0032000113233121-0303020233122311"></a>

#### `primary.rr_set_group.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2202023131100022-1103302232221201-3011011233003132-1222131313102331-0013123113203032-1301132213112102-3002233120002323-1130303100213221"></a>

<a id="canonical-0122022112303300-0110123112132030-0321103100312000-0022333130020331-3002132032033213-3103123301231331-3000212332211233-1121023210131032"></a>

#### `primary.rr_set_group.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- primary.rr_set_group.rr_set

<a id="canonical-3313212012322332-1132311103230202-1103220203333130-3310221013011303-1033220103302101-1021000000330033-1210313021213300-0031330331100211"></a>

Type: `"list"`. Computed.

Resource Record Sets. Collection of DNS resource record sets.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

<a id="canonical-2002210200021330-3201320132220312-0332320022210122-2323221213203120-0000310010321223-3311100001121000-2110030310130221-2102031330003013"></a>

### Direct properties for `primary.rr_set_group.rr_set`

- [a_record](data-sources--dns_zone--reference--group-002.md#canonical-2133222331203223-0313012330312212-3123323310001222-2102203030233020-3301203301220002-1031033311120221-0000331011100133-0320223033201232): complete subsection reference.

- [aaaa_record](data-sources--dns_zone--reference--group-002.md#canonical-0300332203033212-2102110000033322-3321232013303310-3001001213000213-0230010313131203-0232112013213231-3212033301122213-3233101333323002): complete subsection reference.

- [afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-1300113000100303-0103301013032300-2333102222023331-1311010301000230-1312131033332333-3001303320321310-2023130101302221-1203303330012333): complete subsection reference.

- [alias_record](data-sources--dns_zone--reference--group-002.md#canonical-2232212032121023-3022121321313313-0013031312032122-3303023321031011-2033003002102300-0222023203002011-3310121213120113-3120313012020110): complete subsection reference.

- [caa_record](data-sources--dns_zone--reference--group-002.md#canonical-3021102303223021-0311330012113303-1211113233122120-0203302301302020-2223212013012002-2100133123121201-2130010021021103-0021000303323011): complete subsection reference.

- [cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121): complete subsection reference.

- [cert_record](data-sources--dns_zone--reference--group-002.md#canonical-2033130131100231-1310103201133222-0010203201322133-0113003033112230-3113331000213213-0323122330220111-1301022323030020-3123313211203310): complete subsection reference.

- [cname_record](data-sources--dns_zone--reference--group-002.md#canonical-3000110033203111-2312032311212322-2103322303022131-1123102210000211-0312322123020323-2222132212012111-1000102011220113-3123113121110030): complete subsection reference.

<a id="canonical-2000320222210232-1300311103033231-3231232123032330-3001223121001231-2331332212302100-1203320111201203-0121133303013112-1110222233321033"></a>

<a id="canonical-0223213321020232-1332201301000030-1032123033200312-2123202013032013-0232020231110002-0301330123223302-0302301022330131-0331202222202103"></a>

#### `primary.rr_set_group.rr_set.description_spec` property

Type: `"string"`. Computed.

Comment. Human-readable description text

- [ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331): complete subsection reference.

- [eui48_record](data-sources--dns_zone--reference--group-002.md#canonical-1300212020003230-1113311332101220-1132330003331301-3333111020100102-3003313132131220-0122230031303231-2203031223032120-2332332003333102): complete subsection reference.

- [eui64_record](data-sources--dns_zone--reference--group-002.md#canonical-0003303213311201-3003200303313223-1121023300112121-2002211020100233-1112212221032122-1331101211011320-2003332313321333-2110200232323101): complete subsection reference.

- [lb_record](data-sources--dns_zone--reference--group-002.md#canonical-1230020033200103-1330330300123130-3220203300231002-1211332020202310-3103022120112000-2013013012212211-2200123220201232-2330030320121121): complete subsection reference.

- [loc_record](data-sources--dns_zone--reference--group-002.md#canonical-3103002133110011-3300030122231322-1033330212002032-1120030230101113-1221213310123022-1300023322200233-0112012020111212-1320210023313010): complete subsection reference.

- [mx_record](data-sources--dns_zone--reference--group-002.md#canonical-2232020013010222-3011001301213203-0001030100231211-0233002301012011-0102021001102022-2003330212132112-3220310312321133-1330303033320022): complete subsection reference.

- [naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-2130200302222231-0201323023000030-2231013210302021-3320130221100012-2012002232303013-1033323023120320-3201231132301223-2120220121110010): complete subsection reference.

- [ns_record](data-sources--dns_zone--reference--group-002.md#canonical-1301020100313003-2113130112303333-0222200332023331-1011211333020200-3333320312030230-2202333330201232-2032210300303202-3111001112202203): complete subsection reference.

- [ptr_record](data-sources--dns_zone--reference--group-002.md#canonical-2211303020222311-3310323332010203-0232301233301110-3121123122223302-3103033101210322-2020123011002111-3032313210101120-2001120323033330): complete subsection reference.

- [srv_record](data-sources--dns_zone--reference--group-002.md#canonical-0223103023112010-0122311332230213-0020010203033121-2323203031213310-1321013110303211-1310102220022031-3130133320120313-1333011010122312): complete subsection reference.

- [sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332): complete subsection reference.

- [tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-1120111220133012-1033321320120331-3101013303033222-2303001213103323-1230312203202132-1302200000132221-0213201322100310-2310212123010121): complete subsection reference.

<a id="canonical-0113323032231011-0223333222302211-3322300003130030-0323113011131110-0313123301200332-3032113230120122-0323110023212100-0112221033132131"></a>

<a id="canonical-0013023330022121-2221010231221323-2332233220200113-0211132033310030-0303030003022212-2031121000213301-0311323323000112-0021323121321331"></a>

#### `primary.rr_set_group.rr_set.ttl` property

Type: `"number"`. Computed.

Time to live. Time-to-live duration in seconds

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [txt_record](data-sources--dns_zone--reference--group-003.md#canonical-2112000100003211-0010122023020330-2302300331333210-3123120333012012-1313313003312213-2120120330311231-1033321003121122-0030131112120013): complete subsection reference.

<a id="canonical-2133222331203223-0313012330312212-3123323310001222-2102203030233020-3301203301220002-1031033311120221-0000331011100133-0320223033201232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.a_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.a_record

<a id="canonical-3011212320111000-3202333013120221-0021212111332023-2032123232032111-0201112310120223-3310200123333112-0212223220022001-2001113312010031"></a>

Type: `"single"`. Computed.

DNSAResourceRecord. A Records

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

<a id="canonical-0223113110010333-3313310300000303-3330020122323211-1102031233033333-1101223202322132-0321313231322111-1232310302300333-0331003223010311"></a>

### Direct properties for `primary.rr_set_group.rr_set.a_record`

<a id="canonical-3213102231313033-0303331322121002-3300312033112130-0201111312112012-1003032023110211-2223300333323120-1211322302312013-2021322230330230"></a>

#### `primary.rr_set_group.rr_set.a_record.name` property

Type: `"string"`. Computed.

A Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0112323113010023-3023131013113031-3102131201322031-2233113002200212-3130333332313022-3002220000221331-1010312303011011-1023021030113300"></a>

<a id="canonical-0021213132010121-1322113133032112-1310321223332031-2113220303021331-1320030212312330-2311323332322002-2231031110222010-2011322010203013"></a>

#### `primary.rr_set_group.rr_set.a_record.values` property

Type: `["list", "string"]`. Computed.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0300332203033212-2102110000033322-3321232013303310-3001001213000213-0230010313131203-0232112013213231-3212033301122213-3233101333323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.aaaa_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.aaaa_record

<a id="canonical-2023331012311212-2202230303010330-2220121231320033-0011233220203122-1222030112301320-1131232001323310-2311111111133130-0132120210120321"></a>

Type: `"single"`. Computed.

Configuration parameter for aaaa record.

Additional upstream details:

RecordSet for AAAA Records.

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

<a id="canonical-1033200020223202-0301011012120212-2120120220301123-3232303323202301-1110200133113122-2121201000020030-2112202012202103-0133232212231122"></a>

### Direct properties for `primary.rr_set_group.rr_set.aaaa_record`

<a id="canonical-0132130000323000-0120210230321113-0332300303303101-2131331032001110-0103203311103130-3131232113100322-3103033011222301-3203031100330310"></a>

#### `primary.rr_set_group.rr_set.aaaa_record.name` property

Type: `"string"`. Computed.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1201210031130232-1121021132332032-0100200001133112-2013220012213033-0321033333113332-2002303321201102-3103030013230103-3032011201322021"></a>

<a id="canonical-2323233333101332-3010023200110002-2200020202111030-1322330321133230-1321332133130221-2310330332221103-0100323322132203-2313102013122321"></a>

#### `primary.rr_set_group.rr_set.aaaa_record.values` property

Type: `["list", "string"]`. Computed.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1300113000100303-0103301013032300-2333102222023331-1311010301000230-1312131033332333-3001303320321310-2023130101302221-1203303330012333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.afsdb_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.afsdb_record

<a id="canonical-3033302222102131-3033210012122221-1223220200233121-2000203301300123-2311103123330212-0213212012130122-0230201320011032-0231132102321330"></a>

Type: `"single"`. Computed.

Configuration parameter for afsdb record.

Additional upstream details:

DNS AFSDB Record.

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

<a id="canonical-3031002120213230-1002230131233010-1222013233102002-3313131333100021-1212331221321202-0322203311032323-0320333002122001-3303010221101123"></a>

### Direct properties for `primary.rr_set_group.rr_set.afsdb_record`

<a id="canonical-0032301233120202-0033231113100112-3111330023301212-2200010130112022-0331221002213211-0033300110200210-2322302222122332-1300113202211222"></a>

#### `primary.rr_set_group.rr_set.afsdb_record.name` property

Type: `"string"`. Computed.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2301311330323221-2332100303230210-0101102100201033-1203233130220313-2030201110213221-0131201211201033-1101120020202233-3111323103331201): complete subsection reference.

<a id="canonical-2301311330323221-2332100303230210-0101102100201033-1203233130220313-2030201110213221-0131201211201033-1101120020202233-3111323103331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.afsdb_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-1300113000100303-0103301013032300-2333102222023331-1311010301000230-1312131033332333-3001303320321310-2023130101302221-1203303330012333)
- primary.rr_set_group.rr_set.afsdb_record.values

<a id="canonical-1023001200301202-0110001232003000-2021300103333001-0123111002000032-1320201020302133-1233211310230010-2233331130101130-1032330312322202"></a>

Type: `"list"`. Computed.

AFSDB Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2233013220311330-1211301002012230-3231232130113233-3212230201101223-3121022211302010-1110010302022321-0010122120233020-2220011313331113"></a>

### Direct properties for `primary.rr_set_group.rr_set.afsdb_record.values`

<a id="canonical-1110231102320131-3211323223302120-0301301330103020-3102030311203001-1310333221022233-2201120001221211-1031312223130302-2000200323131022"></a>

#### `primary.rr_set_group.rr_set.afsdb_record.values.hostname` property

Type: `"string"`. Computed.

Server name of the AFS cell database server or the DCE name server.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3310302212033311-1100003023013330-3132210322002310-0121211123201130-2020120003011121-1210013320110013-2231203210201011-3323222313231230"></a>

<a id="canonical-2020002121001010-1120233120232121-3230033332323302-0223232011311003-3312323300033100-3202010002213132-3300002233300133-1130133032111330"></a>

#### `primary.rr_set_group.rr_set.afsdb_record.values.subtype` property

Type: `"string"`. Computed.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2232212032121023-3022121321313313-0013031312032122-3303023321031011-2033003002102300-0222023203002011-3310121213120113-3120313012020110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.alias_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.alias_record

<a id="canonical-3330101333113112-1201010022313322-3013003311010121-2021213321311301-1320223332310312-3313023020020102-0333222132313023-1313022000331133"></a>

Type: `"single"`. Computed.

Configuration parameter for alias record.

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

<a id="canonical-3013212011023011-0231303203331311-1220121000211310-3032120321011212-1300301212031132-3033133232010223-3030103212211020-1102223203331031"></a>

### Direct properties for `primary.rr_set_group.rr_set.alias_record`

<a id="canonical-0121102012223122-3013201022220312-1033133222301232-0231202321131130-3233312331331022-3330023120303030-0332131132301222-3301101221303330"></a>

#### `primary.rr_set_group.rr_set.alias_record.value` property

Type: `"string"`. Computed.

Domain. A valid domain name, for example: example.com.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-3021102303223021-0311330012113303-1211113233122120-0203302301302020-2223212013012002-2100133123121201-2130010021021103-0021000303323011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.caa_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.caa_record

<a id="canonical-3330112101203321-2003331021202232-1010102323102020-3030102201031133-2032203333030330-1012301022310203-3102020010230303-2121311003313311"></a>

Type: `"single"`. Computed.

DNSCAAResourceRecord.

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

<a id="canonical-2030320230121131-0333023012020200-2313023133013230-0031033001121133-3130331133310133-2031021020212233-1113232232103002-2010000112331032"></a>

### Direct properties for `primary.rr_set_group.rr_set.caa_record`

<a id="canonical-3203023010220211-2130302031120031-0012331023301333-3310303222120232-2113103010133130-2311001321201020-0132023102310033-1033300320133310"></a>

#### `primary.rr_set_group.rr_set.caa_record.name` property

Type: `"string"`. Computed.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2033002231100233-2130313021213102-0030032333030102-3301130122220233-0010121203021300-3101133012113020-1332312222131312-2301010021131131): complete subsection reference.

<a id="canonical-2033002231100233-2130313021213102-0030032333030102-3301130122220233-0010121203021300-3101133012113020-1332312222131312-2301010021131131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.caa_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-3021102303223021-0311330012113303-1211113233122120-0203302301302020-2223212013012002-2100133123121201-2130010021021103-0021000303323011)
- primary.rr_set_group.rr_set.caa_record.values

<a id="canonical-2233002323122031-0323113231103233-3332300130230020-2301012111103101-0202213230312031-0022032223213103-2220221031331222-2222213023101033"></a>

Type: `"list"`. Computed.

CAA Record Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

<a id="canonical-2323100221333222-3202201212110030-2203331001310100-0322000110231232-3021322232103100-0321332112300233-1320213002330301-2100103210332030"></a>

### Direct properties for `primary.rr_set_group.rr_set.caa_record.values`

<a id="canonical-2211321133212110-1031313232033132-3103331013220112-0303210120233121-0301021131013032-3201001113121002-3021012212122312-1232103220103322"></a>

#### `primary.rr_set_group.rr_set.caa_record.values.flags` property

Type: `"number"`. Computed.

This flag should be an integer between 0 and 255.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-1101210323223113-1331223213132323-1322200301010110-2000012101132313-2122322123301211-2023001023323010-2332013021031110-3221032200302220"></a>

<a id="canonical-0332010312311033-1000020332313323-0230201222132332-3202100000023200-2022333122233231-1221111033321312-1332112303220010-2230112313311101"></a>

#### `primary.rr_set_group.rr_set.caa_record.values.tag` property

Type: `"string"`. Computed.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-0230313221211321-3323301030103232-1121102112023002-3201100201022111-2333133220013211-3120020021013301-2311131201123333-3220101303303333"></a>

<a id="canonical-1311300032332103-2133313021000221-1102330032003032-3331322031010102-1201021013331231-0323202012210022-0011201013103031-3220133100213330"></a>

#### `primary.rr_set_group.rr_set.caa_record.values.value` property

Type: `"string"`. Computed.

Value. Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.cds_record

<a id="canonical-3112102222133222-3003220331112000-1320021313102202-0011300212321131-2113010310310211-3210200013001230-0010310233323110-3012303012123200"></a>

Type: `"single"`. Computed.

DNS CDS Record. DNS CDS Record.

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

<a id="canonical-2213211112232012-1013312031333000-1033331121311101-0300322120011120-2313033003002030-3131210010210221-1210200311332013-0333130101121213"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record`

<a id="canonical-3302113021000030-3032201310101020-1212330002102131-3202003013323323-3123322320212003-0232103121011002-2203201100131123-1031000300221300"></a>

#### `primary.rr_set_group.rr_set.cds_record.name` property

Type: `"string"`. Computed.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123): complete subsection reference.

<a id="canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- primary.rr_set_group.rr_set.cds_record.values

<a id="canonical-0323130310222111-2002023011012232-0210113302102211-1233200322033330-0113111301001120-2203020131211220-1322030002033312-3132100113001313"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2002232101001302-3220103222032113-3200103222130032-3201130103132311-2102031131020101-1330102330132301-2310120023030201-0212133303322033"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record.values`

<a id="canonical-1311020133210033-0320312100301033-0200030313011222-2333111231113311-1311030322303311-1101313022202023-2111222020331103-2030022111321310"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm` property

Type: `"string"`. Computed.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3113103231121323-3231030031230133-1300200122020033-3330323030313330-3011112122111203-2003302222012021-2322112322131103-1301213223313032"></a>

<a id="canonical-1002321031032113-1023300030030312-1310131212101311-1101023302232133-3201132011321010-0221113022001323-3033333002111121-3111031202331033"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.key_tag` property

Type: `"number"`. Computed.

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-1221310323233000-2020211221120032-2020200012022302-0213201022002321-3022012323310230-3023132002312331-0222220232101030-0133303001130101): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-0300231111232000-3003303231213201-3231031133010120-2113321011000103-1113022200201300-2032213301321031-3203302330300313-3230131031133200): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-2303011021010130-3223320101310233-2332323322220231-1000113113222221-3321223221113223-3301003120213030-2303202202020301-3111211300210013): complete subsection reference.

<a id="canonical-1221310323233000-2020211221120032-2020200012022302-0213201022002321-3022012323310230-3023132002312331-0222220232101030-0133303001130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record.values.sha1_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- primary.rr_set_group.rr_set.cds_record.values.sha1_digest

<a id="canonical-2321022202123020-0222121220103123-2222131300102200-0103002133131132-0210100012112032-1102002023303223-1012030202000002-1221002330103130"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-0221103111002021-3201302303013300-0131200133012222-1132232121000220-3332131012223333-3200011103011110-1223210002012133-2203020123321201"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record.values.sha1_digest`

<a id="canonical-2133322000331020-0030100322011210-0300332323232303-0333110303131021-1210111223031103-0302010333221131-0213212032011220-3201232230330122"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest` property

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40
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

<a id="canonical-0300231111232000-3003303231213201-3231031133010120-2113321011000103-1113022200201300-2032213301321031-3203302330300313-3230131031133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record.values.sha256_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- primary.rr_set_group.rr_set.cds_record.values.sha256_digest

<a id="canonical-2000202202001212-1210220312331330-0110131133032212-1300311323302110-3101211203203101-2321001303100120-3312232210121021-0332220100000133"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-1302012111223120-0313330031011320-3123302103221132-0310011111321312-0031333333102333-0010330101310111-2321032302200120-2021010032110201"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record.values.sha256_digest`

<a id="canonical-2202033002002021-1000111021220031-1331033210122233-3313031000000211-3000300233113102-2111230330123332-2230020201302231-0023002103223131"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest` property

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64
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

<a id="canonical-2303011021010130-3223320101310233-2332323322220231-1000113113222221-3321223221113223-3301003120213030-2303202202020301-3111211300210013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record.values.sha384_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- primary.rr_set_group.rr_set.cds_record.values.sha384_digest

<a id="canonical-2012033320333210-1320312201333310-0331302222232311-1003331302102232-1032101231112102-1231021021230101-3333201132021111-0232113010211320"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-3220231313102320-1131012213300323-0233313130002103-3303113113213122-2312121331033222-2212233222133230-1310320231302003-3211200201133332"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record.values.sha384_digest`

<a id="canonical-2211030303020013-0331200320121123-3212311320021110-1031231332012103-2131122222313110-2022313331101021-0200232030320220-1201101232112112"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest` property

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-2033130131100231-1310103201133222-0010203201322133-0113003033112230-3113331000213213-0323122330220111-1301022323030020-3123313211203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cert_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.cert_record

<a id="canonical-1333323223212332-2013302100133331-1221312031023331-0321203302021123-3333310322103003-1213310001113303-0211312030311301-3023103012323033"></a>

Type: `"single"`. Computed.

Configuration parameter for cert record.

Additional upstream details:

DNS CERT Record.

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

<a id="canonical-0002001213330011-3230320011023200-0330030303301311-0023220233213302-0312013101113023-0333313200112133-1101020132020000-3112330212012302"></a>

### Direct properties for `primary.rr_set_group.rr_set.cert_record`

<a id="canonical-0012022202310032-3221110201121102-1101331203011021-2020213332011022-1330102310002331-0000003100331121-2100120221011233-0303223210232310"></a>

#### `primary.rr_set_group.rr_set.cert_record.name` property

Type: `"string"`. Computed.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2032020132122123-2222213020312321-1332130201131112-3213000100131220-3033100331103133-3230013001320320-1302033020131202-2022220120133330): complete subsection reference.

<a id="canonical-2032020132122123-2222213020312321-1332130201131112-3213000100131220-3033100331103133-3230013001320320-1302033020131202-2022220120133330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cert_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-2033130131100231-1310103201133222-0010203201322133-0113003033112230-3113331000213213-0323122330220111-1301022323030020-3123313211203310)
- primary.rr_set_group.rr_set.cert_record.values

<a id="canonical-2130123321101122-3033233202331100-1111013012031301-1313123121032113-3023223120331121-2000332312300303-3000022210322301-2100102202121122"></a>

Type: `"list"`. Computed.

CERT Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1313332212101300-0312332033130301-1013201203310220-2232023323110000-1022300331012333-3000132333322010-0103300320020221-0231030312110313"></a>

### Direct properties for `primary.rr_set_group.rr_set.cert_record.values`

<a id="canonical-3333032312101212-2202131123321222-1020131320202200-3323220001231020-3333330012103133-2031012001301320-2003021211022132-3322000303031121"></a>

#### `primary.rr_set_group.rr_set.cert_record.values.algorithm` property

Type: `"string"`. Computed.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100313120323021-3110303122033131-1131312211110202-1133031331222312-3233300031121000-0113332033231221-1003232210002310-2001111102013021"></a>

<a id="canonical-2111230230221033-2133111300123213-0112222012211133-3030003122122300-0111232032232210-3302332021013111-2233313130200033-3001202032312012"></a>

#### `primary.rr_set_group.rr_set.cert_record.values.cert_key_tag` property

Type: `"number"`. Computed.

Key Tag. Tag for categorization and filtering

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1001133220032031-1213300121231022-1001202300323021-1232011122201033-2233112130131312-3300203202310010-2121310221121232-2033001023012003"></a>

<a id="canonical-1102321022103300-0022231312200200-1110233203122203-3230232032310001-1002132031332120-2333020213222222-2333303200123320-0233201031332103"></a>

#### `primary.rr_set_group.rr_set.cert_record.values.cert_type` property

Type: `"string"`. Computed.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022303100312121-2133320033312012-3220332213330112-0202030130002133-3112131122023121-2101032110003213-2222023222321232-3003020310021133"></a>

<a id="canonical-0310232312111103-1222201022212312-1131113211102330-0112133331211232-1002010132201300-0321213221021323-0110202301030111-1021220113111212"></a>

#### `primary.rr_set_group.rr_set.cert_record.values.certificate` property

Type: `"string"`. Computed.

Certificate. Certificate in base 64 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3000110033203111-2312032311212322-2103322303022131-1123102210000211-0312322123020323-2222132212012111-1000102011220113-3123113121110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cname_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.cname_record

<a id="canonical-1030012213303033-0002303312011021-3111130311232311-0020011033100021-0332202101323301-3133002320230131-2222332200333102-0102302112003002"></a>

Type: `"single"`. Computed.

DNSCNAMEResourceRecord.

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

<a id="canonical-2300010310330233-1112033021221223-0200300032223212-2301320321023223-1220130310033010-2013121013133321-1110123313323011-3011103101122230"></a>

### Direct properties for `primary.rr_set_group.rr_set.cname_record`

<a id="canonical-2003100302220033-3122311233212001-0113022310332003-3200210312121032-0000313000310133-0001200132032130-1301203111221020-0103323320203321"></a>

#### `primary.rr_set_group.rr_set.cname_record.name` property

Type: `"string"`. Computed.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3330130313321200-3212011323333300-3023021223312031-3210212311310232-2232213012131223-3301030123122211-2032112121113302-3021112323212230"></a>

<a id="canonical-0232032012201123-2313212102030001-1302002131331131-3313002002203112-0120230232033110-3112210212320103-2023323011213031-0133001111010111"></a>

#### `primary.rr_set_group.rr_set.cname_record.value` property

Type: `"string"`. Computed.

Domain. Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.ds_record

<a id="canonical-3323123231232232-3121122132113033-2010032032232020-0223223101013113-1320002203003203-0021121122213032-0110113000320310-3131320002200321"></a>

Type: `"single"`. Computed.

DNS DS Record. DNS DS Record.

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

<a id="canonical-3010122323310300-2113311301211120-0321201223303021-1331320120130210-0102203200113321-3313123113201213-1013312233220231-0311101222203121"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record`

<a id="canonical-2120010001021333-2330312330231221-3013231302312000-1012330233333101-2112330010102121-3230301012220311-2211313001301132-3201002122332000"></a>

#### `primary.rr_set_group.rr_set.ds_record.name` property

Type: `"string"`. Computed.

DS Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130): complete subsection reference.

<a id="canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- primary.rr_set_group.rr_set.ds_record.values

<a id="canonical-2002210212313213-0103320021303300-3103231001322333-2111012132002320-3231020112223333-1320100103000123-0230020232113323-1030130320013033"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3002122012010222-2122222033221221-2133221330300201-2233020003322320-1302323103313213-0330100300132013-1202310301102021-3033113112332031"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record.values`

<a id="canonical-2000112311312320-0220020111313121-0200111002020332-1212301331001213-1313232133301232-1332113323012120-1122200133112113-2122201122320213"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm` property

Type: `"string"`. Computed.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2230223230000300-1232232312112300-2032233211113100-0022102120111130-1120222113201100-2330003332012331-1323220230211111-0301333031331021"></a>

<a id="canonical-2033032023103301-0211133333311103-2200011030132331-0300110203111212-0010330023130321-0332000220113210-3000133001222202-3322011011231000"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.key_tag` property

Type: `"number"`. Computed.

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-3203221312201213-2031031221030023-3333333202131301-1331021012331001-2011030223330302-0223330221021222-1331321201021333-0102222111201331): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-2311203331210120-3312112321300113-1022020101003301-2131013301131322-0221033031112002-1122131211322321-0110130013320321-2121331211133023): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-3010013011001122-2212200033031212-2021022133102303-0230220332102103-1202131231131311-1000102210210112-2302010302202021-3000222221223232): complete subsection reference.

<a id="canonical-3203221312201213-2031031221030023-3333333202131301-1331021012331001-2011030223330302-0223330221021222-1331321201021333-0102222111201331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record.values.sha1_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- primary.rr_set_group.rr_set.ds_record.values.sha1_digest

<a id="canonical-1311001023013010-2120332201020300-0330303232223012-0003321111113220-3233132311300323-0331100222123203-0010100312122300-0202111211311011"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-0301131310331213-3311122233111100-1301021000323100-2000330131033131-2221310310332012-3313013300101131-3322303033111220-3202102310210023"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record.values.sha1_digest`

<a id="canonical-0211022311323301-2333203131012220-1031110301003033-1132331011133230-3030133201011313-2100123333112000-1122031102003013-1032211121233101"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest` property

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40
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

<a id="canonical-2311203331210120-3312112321300113-1022020101003301-2131013301131322-0221033031112002-1122131211322321-0110130013320321-2121331211133023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record.values.sha256_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- primary.rr_set_group.rr_set.ds_record.values.sha256_digest

<a id="canonical-3120333302320123-1021333022223330-3302111213102300-2111313030123132-0303031303033232-0210010310002012-0212000111133301-2111312021021231"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-3010013121113201-1110123203311002-3200232310030132-0203331133310101-0000133312202312-2022212133320121-3220202112230112-1131301102021231"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record.values.sha256_digest`

<a id="canonical-0000202333011311-2211122320113121-2311200101111101-3200233123201201-2222202310123300-3133121112121013-1213110333131131-3032030033221030"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest` property

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64
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

<a id="canonical-3010013011001122-2212200033031212-2021022133102303-0230220332102103-1202131231131311-1000102210210112-2302010302202021-3000222221223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record.values.sha384_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- primary.rr_set_group.rr_set.ds_record.values.sha384_digest

<a id="canonical-3003231321331023-2130302302202211-1301303101130013-1003130313231020-1111332313312312-1033021122012133-3101033031223101-2333020010213003"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-2103020311002013-0013111103223123-3023313320333313-0202202111002100-3310313211131112-0300001131321331-3111130012300132-3032102103003303"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record.values.sha384_digest`

<a id="canonical-1101221320221223-1231332032211112-3231113020303321-0300321103023000-1133300001133003-3132032231213321-0001011320020302-2303232101212231"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest` property

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-1300212020003230-1113311332101220-1132330003331301-3333111020100102-3003313132131220-0122230031303231-2203031223032120-2332332003333102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.eui48_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.eui48_record

<a id="canonical-2001000003012201-0103211033302211-3210331303310200-0132323302313312-2131011003203221-3101321110100030-1212121121002232-2010002212211130"></a>

Type: `"single"`. Computed.

Configuration parameter for eui48 record.

Additional upstream details:

DNS EUI48 Record.

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

<a id="canonical-0133202001213321-3230312233020033-0313003320203001-2120220100221101-0220133021131100-2111121301133122-2232203322323330-3202133311301323"></a>

### Direct properties for `primary.rr_set_group.rr_set.eui48_record`

<a id="canonical-3130212232120323-1122323212113010-0001012120200023-0201130300120321-2230111012303020-1110303000001301-0000030020001230-3011110120212102"></a>

#### `primary.rr_set_group.rr_set.eui48_record.name` property

Type: `"string"`. Computed.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1121220222210001-2202001311302222-3222323232333211-2020031120112200-2111201311232301-1330201121301331-3302310300220021-1123133321323333"></a>

<a id="canonical-0013213102312113-3230003232201320-1121231111323033-0301312002021103-1022233231333123-3310311320331002-1033310031112000-3220320301320112"></a>

#### `primary.rr_set_group.rr_set.eui48_record.value` property

Type: `"string"`. Computed.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-0003303213311201-3003200303313223-1121023300112121-2002211020100233-1112212221032122-1331101211011320-2003332313321333-2110200232323101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.eui64_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.eui64_record

<a id="canonical-1212201231123200-0110310002232221-1000113123120321-1200121211201322-1123100112113002-1230103230003231-3122120313201201-3133223022301123"></a>

Type: `"single"`. Computed.

Configuration parameter for eui64 record.

Additional upstream details:

DNS EUI64 Record.

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

<a id="canonical-1212322311103022-2000003102111333-0202110300302330-3122202321312313-0132120003322113-1001300030101230-0221003031302312-2010312321310321"></a>

### Direct properties for `primary.rr_set_group.rr_set.eui64_record`

<a id="canonical-3223031003000022-1301323031312210-1302322221021013-1133310210310132-1220023121201013-0220301331002032-0330103110030000-1211221101132120"></a>

#### `primary.rr_set_group.rr_set.eui64_record.name` property

Type: `"string"`. Computed.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2331200101303131-2201302132103211-1310001230130200-3321132230320223-3133322100001223-1011002231031201-3112211203302113-1311103100303112"></a>

<a id="canonical-2020233101313132-2332121100223221-2300321300003001-0022120130233301-2311330120212133-2011133201130101-2332031222303123-2003320110230100"></a>

#### `primary.rr_set_group.rr_set.eui64_record.value` property

Type: `"string"`. Computed.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-1230020033200103-1330330300123130-3220203300231002-1211332020202310-3103022120112000-2013013012212211-2200123220201232-2330030320121121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.lb_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.lb_record

<a id="canonical-1010002011111230-0030000301202311-2013002322300101-3033112001200130-3300230103231010-2032221112102132-3233310201023023-1110101302322332"></a>

Type: `"single"`. Computed.

DNS Load Balancer Record. DNS Load Balancer Record.

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

<a id="canonical-1021132233131202-2102113103011130-3323222100111222-1330133102130113-0202003222002112-1110301113202330-3031001022210301-3202313002221102"></a>

### Direct properties for `primary.rr_set_group.rr_set.lb_record`

<a id="canonical-0000321020211220-0303010000002220-0212332120131033-2232210021012112-3213233323302222-2022012000131232-0211310122022312-0213101122333323"></a>

#### `primary.rr_set_group.rr_set.lb_record.name` property

Type: `"string"`. Computed.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](data-sources--dns_zone--reference--group-002.md#canonical-0121010122012212-1330211002022212-1012203201000101-3332330100222032-3201223332222113-0021233010201332-1112313320320122-2313332311002133): complete subsection reference.

<a id="canonical-0121010122012212-1330211002022212-1012203201000101-3332330100222032-3201223332222113-0021233010201332-1112313320320122-2313332311002133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.lb_record.value` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-1230020033200103-1330330300123130-3220203300231002-1211332020202310-3103022120112000-2013013012212211-2200123220201232-2330030320121121)
- primary.rr_set_group.rr_set.lb_record.value

<a id="canonical-3121320223110223-1332123103002222-2313010003220230-2012231231333302-1010223011211013-2321221233000212-0331033121313101-0023233010031111"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3222111222333302-1001022313120120-1003002220320132-1111023003222223-1131232210332022-0213103321031310-1132122213122222-1132301200210200"></a>

### Direct properties for `primary.rr_set_group.rr_set.lb_record.value`

<a id="canonical-0000202220212313-3111321311123302-2131021122203101-3000230310120101-1112231002303110-1303023323321232-0013211232021121-2031223113331022"></a>

#### `primary.rr_set_group.rr_set.lb_record.value.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3313322000321000-3113213123012211-2032012223101033-2231033201320333-1310023100033020-0100030313121201-0210333130312303-2221022213230133"></a>

<a id="canonical-3131320133330101-2000310232213010-3313230330112000-0011002002132211-2113300231203312-0100220212022101-0102231032131233-3320130021022021"></a>

#### `primary.rr_set_group.rr_set.lb_record.value.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2330303013222230-0011223112102330-0022222023310020-2032231332020121-1232102130201023-1202203103200000-0203222211233301-3332202201032101"></a>

<a id="canonical-0230113230031230-3023112302213101-0123313133133301-0302033203202101-1033120230202232-1333201320111013-2123332110032230-2222233013233212"></a>

#### `primary.rr_set_group.rr_set.lb_record.value.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3103002133110011-3300030122231322-1033330212002032-1120030230101113-1221213310123022-1300023322200233-0112012020111212-1320210023313010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.loc_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.loc_record

<a id="canonical-2310233030001321-2302022120200120-2011323023012232-3330132222012231-3121102131100201-0332131311020001-2001331033212200-1023211110112202"></a>

Type: `"single"`. Computed.

DNS LOC Record. DNS LOC Record.

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

<a id="canonical-1120220103321202-1231313333023123-2122203122100011-0311313033210132-2210330202212130-1233022132330231-0033032220323311-0202101302311333"></a>

### Direct properties for `primary.rr_set_group.rr_set.loc_record`

<a id="canonical-3132112120333211-2233232033231211-1000331213201310-1101220230213201-0011032113331230-1030330122200111-1031020013333310-2333122001003212"></a>

#### `primary.rr_set_group.rr_set.loc_record.name` property

Type: `"string"`. Computed.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2330323032333122-2133122122112213-0010023311010233-1013321030002213-1010011203220211-1333111330011310-0110333330333322-2302320200003101): complete subsection reference.

<a id="canonical-2330323032333122-2133122122112213-0010023311010233-1013321030002213-1010011203220211-1333111330011310-0110333330333322-2302320200003101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.loc_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-3103002133110011-3300030122231322-1033330212002032-1120030230101113-1221213310123022-1300023322200233-0112012020111212-1320210023313010)
- primary.rr_set_group.rr_set.loc_record.values

<a id="canonical-1211112130033313-0001333210220013-1012303312003132-1302203221002201-1011022323200210-2103131232211132-1112322032113322-2101333131202112"></a>

Type: `"list"`. Computed.

LOC Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3201230022021200-3311221003223120-1002321322011113-1121131101330233-1132200003323032-2133120113001102-0023302223211013-1312220233010310"></a>

### Direct properties for `primary.rr_set_group.rr_set.loc_record.values`

<a id="canonical-2222132112000030-0310330123330132-0321123120201130-0111123130111021-2322113203022132-3221101203323322-3122223211202021-1220121231133330"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.altitude` property

Type: `"number"`. Computed.

Altitude. Altitude in meters.

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
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3022330112300003-2101102102123232-2033210031113203-3002322120120220-2213030031112110-2311010213002330-3012311300031232-3021121331102103"></a>

<a id="canonical-3111212122100220-1031111201201300-0131123022000332-2023023301220111-3010113111001312-1003110321110002-1020121013203121-0112103033211310"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.horizontal_precision` property

Type: `"number"`. Computed.

Horizontal Precision. Horizontal Precision in meters.

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

<a id="canonical-3110222123202312-3023302020220102-3303211313030132-2333010102201203-3221232320101031-0213101110130202-2121323110220213-1300013223030123"></a>

<a id="canonical-1211012113031210-3203212331221313-0200011321133301-0013011120231301-3113003130023032-2211323102313200-1123310030022021-3101010302110133"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.latitude_degree` property

Type: `"number"`. Computed.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0111330302313111-2332211212300211-3101302322000200-1233233202312011-0210232101131002-1303121021013232-3223312011021132-3201300323021120"></a>

<a id="canonical-1220003102021113-3030212231120232-2230310100322001-3220323201211200-2211321031230002-0113333020031002-0313103100313120-1202323033302010"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere` property

Type: `"string"`. Computed.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3200123113211231-3332120110000101-0110320132310200-0001222030032322-0310122331020230-2311203101032201-2301033033103122-2321133131203232"></a>

<a id="canonical-2303122232303031-3122221323031103-0133110123033101-0020000013221310-1020010131023123-1123000323000223-2322200000111333-1031220031212321"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.latitude_minute` property

Type: `"number"`. Computed.

Latitude minute, an integer between 0 and 59, including 0 and 59.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2210102333110102-0031330021011331-3020033131333202-1033111233002220-2110211233321231-0032032001331021-3103002220111222-0202321222020020"></a>

<a id="canonical-1002001120133010-3033321303311222-2012123202202123-3201200333303210-0203322130212031-1021231223003330-0121220023013110-3232020131122120"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.latitude_second` property

Type: `"number"`. Computed.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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

<a id="canonical-1212030312111130-1312210022310221-1010201220330032-0220302120333312-2222332330230100-0032301211312102-1011212222221011-0112212201020311"></a>

<a id="canonical-3223330333210201-2133023013211011-3010131001233332-0131121232132102-0320112322232232-2020311201011310-1222202121232100-0132222301200012"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.location_diameter` property

Type: `"number"`. Computed.

Diameter of a sphere enclosing the described entity, in meters.

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

<a id="canonical-2111001200013213-0110120330110322-0331221330203303-3102132333302121-3110020033013120-2310032301031003-2312013123323222-0010323001321223"></a>

<a id="canonical-1021231002330213-0023130330220223-0300130103122013-3313321333321311-3330113101130232-3310022103323123-3120310312202301-1332323031103222"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.longitude_degree` property

Type: `"number"`. Computed.

Longitude degree, an integer between 0 and 180, including 0 and 180.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3111201232032021-1220021122201120-2033010222130000-0022210021203103-3101002332121313-1311331012330212-3000003302112111-2123010131000321"></a>

<a id="canonical-2330122110211022-1212332002303110-1313330322200121-1101331000321130-3131330012312230-3013221101122233-2101112123223111-2113310133223330"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere` property

Type: `"string"`. Computed.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

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

<a id="canonical-1030211331000333-2301033202230211-1103101233121032-1032131212102230-1332223222132232-0312010111001312-3211332012212022-3020023001300003"></a>

<a id="canonical-0022323111011031-1330001321100231-0112233002103102-2321320013321212-1101100221031130-0211320103023333-3331331112312310-2301310301330233"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.longitude_minute` property

Type: `"number"`. Computed.

Longitude minute, an integer between 0 and 59, including 0 and 59.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2331223312002332-2222001221322123-0102102132203110-3321230121322300-1032102220010030-0211012221133102-1011212332200333-1031221332101102"></a>

<a id="canonical-3131023102111323-3220020220332231-0021132122232011-1032103320232221-0320113310220320-1001211113131101-3133023033223320-3112003020203313"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.longitude_second` property

Type: `"number"`. Computed.

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

<a id="canonical-2200331331113020-0321300301110030-1002113022231312-0123033233113033-2101130132120012-3012022111201103-2230222003012022-0332210312312211"></a>

<a id="canonical-0212032203122131-3110202033122133-3203303303110103-1130022023211303-1030031203033132-0111000023023202-3020000021303233-2231331220132232"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.vertical_precision` property

Type: `"number"`. Computed.

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

<a id="canonical-2232020013010222-3011001301213203-0001030100231211-0233002301012011-0102021001102022-2003330212132112-3220310312321133-1330303033320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.mx_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.mx_record

<a id="canonical-1032301230121302-2230303233101300-3200131020003011-3223231131120111-3211000013130303-3113013032123030-2313121333332032-0030212331102010"></a>

Type: `"single"`. Computed.

DNSMXResourceRecord.

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

<a id="canonical-0330323212233023-0020010233021011-1100131101003002-2220012233233202-3330101123332113-3201130320331130-3033003201330003-3232330130210310"></a>

### Direct properties for `primary.rr_set_group.rr_set.mx_record`

<a id="canonical-3311033132001020-0102022220030111-1232031110121310-2231022232031233-1133110222101113-1132313012310030-2313023133102322-1333311331220330"></a>

#### `primary.rr_set_group.rr_set.mx_record.name` property

Type: `"string"`. Computed.

MX Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-3312021201221301-3030330303323212-1020333131110122-1133200100220220-3301010123311000-1023002001023200-3221330220300001-3323322023301322): complete subsection reference.

<a id="canonical-3312021201221301-3030330303323212-1020333131110122-1133200100220220-3301010123311000-1023002001023200-3221330220300001-3323322023301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.mx_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-2232020013010222-3011001301213203-0001030100231211-0233002301012011-0102021001102022-2003330212132112-3220310312321133-1330303033320022)
- primary.rr_set_group.rr_set.mx_record.values

<a id="canonical-2010113303312310-1321103120212023-2203002233330103-3133021300012001-0313201120310203-2332220233012323-3310013003130201-2321220322121321"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1332100132001211-2130302333133300-2120132003332010-0123310010102013-2133303301313022-2102212010032013-2331222201020132-1021013120110013"></a>

### Direct properties for `primary.rr_set_group.rr_set.mx_record.values`

<a id="canonical-2212112331311011-1102003102311101-3303301010021132-1132200332102233-1303330021322112-2102312121323021-3020103332131030-3211011300301002"></a>

#### `primary.rr_set_group.rr_set.mx_record.values.domain` property

Type: `"string"`. Computed.

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1102132003010311-2221313113001020-3201012012000233-0212222330213313-0203330131320113-1211222100232300-3202231313221013-3313300300310020"></a>

<a id="canonical-0023312212102102-2312103132332311-0233211303230103-0221221221133311-2023110220311200-3102202100030332-1103132332222131-2300020100320022"></a>

#### `primary.rr_set_group.rr_set.mx_record.values.priority` property

Type: `"number"`. Computed.

Priority. Mail exchanger priority code.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2130200302222231-0201323023000030-2231013210302021-3320130221100012-2012002232303013-1033323023120320-3201231132301223-2120220121110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.naptr_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.naptr_record

<a id="canonical-0133232032210101-0001200232000323-1130202303131203-3331333233303332-0012131002202331-0320323131233332-0000212013233331-1011012221020233"></a>

Type: `"single"`. Computed.

Configuration parameter for naptr record.

Additional upstream details:

DNS NAPTR Record.

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

<a id="canonical-3220032000133032-2312300312030100-2020012011232112-0123230202321020-1302030333022113-2112333301010111-0300322023230302-0123201003012311"></a>

### Direct properties for `primary.rr_set_group.rr_set.naptr_record`

<a id="canonical-1212011312322322-0312203331301013-3322121111210101-2200211220000111-3032112033200333-0300322110002031-3202232010012220-1112222133303213"></a>

#### `primary.rr_set_group.rr_set.naptr_record.name` property

Type: `"string"`. Computed.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-1023231202220220-2030321031113022-3303221030012233-3132133203303320-1021013000001222-3211013210110322-2121322002201003-3010321130013113): complete subsection reference.

<a id="canonical-1023231202220220-2030321031113022-3303221030012233-3132133203303320-1021013000001222-3211013210110322-2121322002201003-3010321130013113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.naptr_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-2130200302222231-0201323023000030-2231013210302021-3320130221100012-2012002232303013-1033323023120320-3201231132301223-2120220121110010)
- primary.rr_set_group.rr_set.naptr_record.values

<a id="canonical-2200231101020302-0013202003222102-3000130322230201-1032213200122031-0022233212211123-0021232220010203-3020201223101211-0110020030232033"></a>

Type: `"list"`. Computed.

NAPTR Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0031333102012010-3321210310013333-3110203222322203-1100321221321302-1131020103003231-2300211110220303-2321232130220112-1203131113330331"></a>

### Direct properties for `primary.rr_set_group.rr_set.naptr_record.values`

<a id="canonical-1031012120012113-1313312322333103-2032103222123012-2121333312111012-3031130302031121-2202230212033331-0203301113002231-2312323013113013"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.flags` property

Type: `"string"`. Computed.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2311202231301323-3121333022301101-2223030001121022-0022120312113201-0303102103310032-2220312031222320-0312013323332111-2000122301101021"></a>

<a id="canonical-2121100211210300-3102022302031020-0301130030333011-2321012301332102-3131321100321032-2231103133131222-2322223121231103-2130313102130311"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.order` property

Type: `"number"`. Computed.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0032220230321232-0210110100300333-0111130031311110-3013012130120210-0211233310211013-1131033010331121-2223031133003030-2013111232122032"></a>

<a id="canonical-1101202023201303-3012320000110322-0231030111332020-1222322221313303-3203132221220330-1022330321320010-1311002321023020-0000113013023033"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.preference` property

Type: `"number"`. Computed.

Preference when records have the same order. A lower number indicates a higher preference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0101230230333232-1332200302210321-3130112222202020-0102100120101330-0121133023313321-3112011001110013-2323021301022023-0003332310132202"></a>

<a id="canonical-2320031020301320-2310002100233032-0202003110012023-3001323030330220-1012302322101302-2310120320212110-0210101333021201-1122212303000122"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.regexp` property

Type: `"string"`. Computed.

Regular expression to construct the next domain name to lookup.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2200200231122010-1201013102211212-2301011122201231-2223033312113200-2311302021012310-3030120312101310-0121001121111000-0222110200010111"></a>

<a id="canonical-1111202123332021-0112013030123220-1303131113300321-1320312012220301-1110133033322000-3212033200202330-3303101220001132-1310202222131000"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.replacement` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0222322101131033-2330013302030033-2322120201123121-2313311332010213-2210100232003131-3301000021132131-1001003200202201-0313032032003033"></a>

<a id="canonical-2322300330031211-2021031130032233-1213012231002303-2210330132313322-1023313210111332-2123032221010332-1222023300221302-0321121220101001"></a>

#### `primary.rr_set_group.rr_set.naptr_record.values.service` property

Type: `"string"`. Computed.

Specifies the service(s) available down this rewrite path.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1301020100313003-2113130112303333-0222200332023331-1011211333020200-3333320312030230-2202333330201232-2032210300303202-3111001112202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ns_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.ns_record

<a id="canonical-0330312201100120-2112312301203220-3230003000130221-0023232000030332-3220220102110023-1132123131100021-3103230321221112-3322223320131103"></a>

Type: `"single"`. Computed.

DNSNSResourceRecord.

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

<a id="canonical-1221333330213002-2312322233200133-1003313023122330-2303323032333130-3133112211110003-3103122121213222-2322021310223002-0122332223330130"></a>

### Direct properties for `primary.rr_set_group.rr_set.ns_record`

<a id="canonical-0200211233030201-3122121200101333-3220202313311001-1213312221301101-2333121220001202-3100112100122322-0301211011133220-3031313202011301"></a>

#### `primary.rr_set_group.rr_set.ns_record.name` property

Type: `"string"`. Computed.

NS Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2121333000030022-1111112311202130-1321222312321100-0231012012113003-0022310110111233-3320111233223231-1230011103023030-3130102330101301"></a>

<a id="canonical-2113130113223110-3323033030220131-1201202232200301-1200103322233102-1032031102023133-1231323002012021-3021203311231033-1313303002011023"></a>

#### `primary.rr_set_group.rr_set.ns_record.values` property

Type: `["list", "string"]`. Computed.

Name Servers. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2211303020222311-3310323332010203-0232301233301110-3121123122223302-3103033101210322-2020123011002111-3032313210101120-2001120323033330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ptr_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.ptr_record

<a id="canonical-3030302130012103-1031200320000313-3213320132112032-3011133211333100-2211102023313110-0000310203120132-2233122200311201-0012110010030112"></a>

Type: `"single"`. Computed.

DNSPTRResourceRecord.

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

<a id="canonical-3303322130222013-0123333021021123-2020321331201002-2302303310302123-0233332113101022-1103111132020313-1101221203320131-3331211121322000"></a>

### Direct properties for `primary.rr_set_group.rr_set.ptr_record`

<a id="canonical-2320112321012331-2211032133102013-2012112213232302-2200312233332133-2030122012311323-3000122321100123-2301210012312211-0232130132131110"></a>

#### `primary.rr_set_group.rr_set.ptr_record.name` property

Type: `"string"`. Computed.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3022000221212222-1223011132033230-1001032003202132-3013221033030122-0030010331211121-3002101102030231-2032222322202102-0122110132221302"></a>

<a id="canonical-3321211200233022-2111101000233023-1010100221302102-2031011300311122-3022203002300020-1201311330213102-1203201213211313-3030100301312202"></a>

#### `primary.rr_set_group.rr_set.ptr_record.values` property

Type: `["list", "string"]`. Computed.

Domain Name. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0223103023112010-0122311332230213-0020010203033121-2323203031213310-1321013110303211-1310102220022031-3130133320120313-1333011010122312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.srv_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.srv_record

<a id="canonical-1000221220030112-2010100232132001-3022002120021130-1103323000130332-1102212231213330-0122231210232320-2020130132322232-3121222113332300"></a>

Type: `"single"`. Computed.

DNSSRVResourceRecord.

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

<a id="canonical-1221020320210113-3111111212221203-3312213101001001-1203330120130132-1023301320301213-1223113203331001-3122231020223011-2121313320100111"></a>

### Direct properties for `primary.rr_set_group.rr_set.srv_record`

<a id="canonical-2302322211300310-3222300011022310-0210330303222112-1321232112133201-1223132210322111-1213310002320301-2223302100310113-0202212113031003"></a>

#### `primary.rr_set_group.rr_set.srv_record.name` property

Type: `"string"`. Computed.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-0211103321112302-0300323031310021-2302311310202221-3120013221210130-1001211300111011-2322132002232311-3113032000333100-1312130101032030): complete subsection reference.

<a id="canonical-0211103321112302-0300323031310021-2302311310202221-3120013221210130-1001211300111011-2322132002232311-3113032000333100-1312130101032030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.srv_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-0223103023112010-0122311332230213-0020010203033121-2323203031213310-1321013110303211-1310102220022031-3130133320120313-1333011010122312)
- primary.rr_set_group.rr_set.srv_record.values

<a id="canonical-2212233103113113-0311022000300322-0310231201320310-1100021111230311-2303130003123100-2223332202103301-0111000113330012-1221313232200032"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3110131200123110-0030000232120230-2300100013112130-0310110223121310-1301120231312003-1121012113311233-1112110112002232-3302122100233230"></a>

### Direct properties for `primary.rr_set_group.rr_set.srv_record.values`

<a id="canonical-3003201220200130-1021330232001031-2321011021133032-0002310022312232-2102213113331110-0132012022323033-1223031023100212-0321202210233230"></a>

#### `primary.rr_set_group.rr_set.srv_record.values.port` property

Type: `"number"`. Computed.

Port. Port on which the service can be found.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3210220321011103-0221123113220123-3101321201123011-3132321110103301-2232212102010012-3320012113021320-3201132330203101-1003103012001021"></a>

<a id="canonical-3323213110310313-0202013222021022-3302200321313210-2220220301313103-2021103120013021-2203231310310111-2210223102301323-0103202133102122"></a>

#### `primary.rr_set_group.rr_set.srv_record.values.priority` property

Type: `"number"`. Computed.

Priority of the target. A lower number indicates a higher preference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2222300302002330-0331121232303121-1323233320131120-2101102203212132-2213100223130100-1131220022202331-2100132311031101-0210020030000301"></a>

<a id="canonical-2012310322211332-3103303300132123-3111211232101310-3231320230120112-2203313020310323-0021213302023112-0111233233320020-0222031001301110"></a>

#### `primary.rr_set_group.rr_set.srv_record.values.target` property

Type: `"string"`. Computed.

Hostname of the machine providing the service.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0313220012311033-3131021312031102-2010003303132000-2313020302003312-3301221231001320-0300103203001322-1232303020123102-3030021033312210"></a>

<a id="canonical-1203101200323110-2202210220333221-0120102111011303-1202103212310022-2100223001111112-1323233032200110-3312122023223033-1112203223133123"></a>

#### `primary.rr_set_group.rr_set.srv_record.values.weight` property

Type: `"number"`. Computed.

Weight of the target. A higher number indicates a higher preference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.sshfp_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.sshfp_record

<a id="canonical-2121012320102322-0110331111102013-2132110133203320-3100031022001321-0032102321220212-3221111323222212-3100312000000230-0120322113132122"></a>

Type: `"single"`. Computed.

Configuration parameter for sshfp record.

Additional upstream details:

DNS SSHFP Record.

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

<a id="canonical-3022300333003101-3123022303110313-3120103131232130-3131022333100222-0021313311202132-0131321100110312-3231203111233231-1012102233030221"></a>

### Direct properties for `primary.rr_set_group.rr_set.sshfp_record`

<a id="canonical-1231031000113201-3321323320010331-3323322201120313-0222231133212021-1013111310313302-1232032001323032-1233021231002103-3113110111201021"></a>

#### `primary.rr_set_group.rr_set.sshfp_record.name` property

Type: `"string"`. Computed.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [values](data-sources--dns_zone--reference--group-002.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110): complete subsection reference.

<a id="canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.sshfp_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332)
- primary.rr_set_group.rr_set.sshfp_record.values

<a id="canonical-2331300330310133-0031203323201200-0020002201201200-0030033113010311-3330221213323321-1321210003322111-3011230300321203-0332322021001321"></a>

Type: `"list"`. Computed.

SSHFP Value. Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0012021110032011-2120321301320300-1000113333030101-2302232213320312-2103003332301213-3121223302320310-1321012011130022-2202311301301232"></a>

### Direct properties for `primary.rr_set_group.rr_set.sshfp_record.values`

<a id="canonical-1300023031211311-3313232011221303-0122212033232122-0212102021130121-1002320332020233-2310132330313320-1231033230020221-1213001232222300"></a>

#### `primary.rr_set_group.rr_set.sshfp_record.values.algorithm` property

Type: `"string"`. Computed.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

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

- [sha1_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-1212033231013101-1211210110120033-2201033211111131-0322310331021203-2200012100333231-0333231210321010-3113021122010322-1233011230030312): complete subsection reference.

- [sha256_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-0321003213330333-0032202123013222-1320212310111100-0222320102323300-1220220323000302-1020321111210022-3020001101001020-3133201013032301): complete subsection reference.
