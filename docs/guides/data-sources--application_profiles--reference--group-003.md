---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-1302311233020213-2030012101003320-1022203113032033-2233230221023300-2313312002213001-1012011213321122-3313010131311223-0102120002220031"></a>

## `virtual_server.http3.http_client_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2132210212333322-0102002020003113-0313200000111313-3111331333013320-2330110023112311-3001003223132132-2202330323331111-3203001132221202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.http_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.http_server_profile

<a id="canonical-3221312202023033-1102223012313300-3203010330011230-1210002202322011-2313030323231301-2222112330223212-1201303211311211-3123213113100210"></a>

Type: `"list"`. Computed.

Configuration parameter for http server profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3013313120112010-3323322103123331-3223000103212133-2103301132120031-2320103023231022-2110211303200312-2313131110122212-0201213002112032"></a>

### Direct properties for `virtual_server.http3.http_server_profile`

<a id="canonical-0101013121221132-1202201020020332-2202330300302233-3313203200201322-1002310330030222-1012311322121200-2132232133011113-1033300123331332"></a>

#### `virtual_server.http3.http_server_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0013313312210120-3130313333213333-1333213021111233-3213323103121202-1301102233021123-3303133120310333-2032312231023020-1323111311320313"></a>

<a id="canonical-1313321033321233-0130101202013203-1011200300123100-2032322202132010-0203002133123123-1220302332301012-1313322322303312-3323330322101133"></a>

#### `virtual_server.http3.http_server_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2213223111233123-1123303202020100-1120121101130001-1010113022330302-2310102333101101-2123223003313132-3012031111320013-3220300102203001"></a>

<a id="canonical-2103320033220021-1312013331103200-3320031133132202-0222132123202131-0113323033332212-2021310332020032-1033333001312232-0021120313202230"></a>

#### `virtual_server.http3.http_server_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2030211003122121-2230233120231211-0113032010002221-3021110031232003-0100121021033311-0230121020212022-3132123120320200-0321222313213122"></a>

<a id="canonical-2123331121022321-2202112221011033-1231221322323321-0033333010311200-3231232002003313-1210102322202001-2101120111020200-2123023200320312"></a>

#### `virtual_server.http3.http_server_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1133202011021202-3203233311231300-0221232013231200-0120132201231022-2302233003031303-0000233300032303-2221030002333213-0101322032333203"></a>

<a id="canonical-2003021311100120-3032123212303310-0321203210130323-0220220310210303-3003111313113021-0020312222200331-0120200230230232-2202333013203123"></a>

#### `virtual_server.http3.http_server_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0100013232022300-0111113200101122-0031003230203222-1131311321302302-2321000312232033-2202310303231100-0321032330201012-2211321031201111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.quic_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.quic_profile

<a id="canonical-2311003220220232-3230111332311321-3322012111002332-1020121133211300-1233032213310202-3311330100212020-3100011331330303-2020320213010313"></a>

Type: `"list"`. Computed.

Configuration parameter for quic profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2033332121210102-2313033203131230-2311322123033001-0332313011230021-0002221210201000-3312301101003132-1203130002220300-3311211100220231"></a>

### Direct properties for `virtual_server.http3.quic_profile`

<a id="canonical-1203331001223123-1121330223313120-1033122102032331-3330213310013021-1112032331012131-1111030133002003-0330213110200010-3231112320003111"></a>

#### `virtual_server.http3.quic_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0111021202211021-3110031212030122-0003313012020102-1322022003100203-1313300131210223-1103332102031111-3000123132230023-1031100120212321"></a>

<a id="canonical-0200011230121302-3211122321023031-3033313020111001-1032232123200313-0302130312123122-1223231333322321-2331303310111210-3031013103302301"></a>

#### `virtual_server.http3.quic_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2320111311120023-2023011231032330-2013033301023013-2020001013232033-2131302110331312-3220312130013031-0130011102002220-1203222002102111"></a>

<a id="canonical-3223330321021110-3103111220032032-0113021231310123-0000321010133021-3332102021010000-1323222202031302-2302000032210002-0031223302331110"></a>

#### `virtual_server.http3.quic_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2132213323131023-1321220100310030-3313323012000221-1102232221332033-1232011223323110-0310111030032100-0121223023131133-3210001113331001"></a>

<a id="canonical-0221033020330100-2232002103233131-1000200003123223-2322320023321100-3113213332333013-2010222000213201-1122300103223131-2110113102003232"></a>

#### `virtual_server.http3.quic_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2023200220330321-1303111011233222-1230003323303202-1102333003332033-2022212213331330-2302133320032220-2332200210130310-2003123300223210"></a>

<a id="canonical-3120132101003321-2110021031231223-2122011101111113-1321313013220300-2202120120123233-3000303110011230-1200223211213132-0203032220231221"></a>

#### `virtual_server.http3.quic_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0033200333002111-3332200110001031-0202212132212132-2013332121113223-0031132032002310-3123011022020321-2122311103211212-3130200323302122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.server_ssl_profile

<a id="canonical-2202321321031233-2131133023103213-1033011033022203-2110233230020231-0003201231231231-1123321331123200-3032133011212211-3033330020221003"></a>

Type: `"list"`. Computed.

Configuration parameter for server SSL profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0201033012032231-2130000330113323-0333231110230121-0133002212111300-3222000020032312-3310121201033003-0332133221322232-2110112011220300"></a>

### Direct properties for `virtual_server.http3.server_ssl_profile`

<a id="canonical-0012302313310131-1103200323031103-0120323210001020-1301132123213310-2210330331101200-1130030112120121-2031233323002311-1121011332311312"></a>

#### `virtual_server.http3.server_ssl_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1032032130033032-0001002123020200-2320330323322011-2101111011000112-0212031120113311-1121131012021232-0202202313011301-2202300323013101"></a>

<a id="canonical-0020022211310232-0203203301002000-1103323200020030-3121030231312003-1123022131331222-1121220012122113-0000320110301112-3021130032231011"></a>

#### `virtual_server.http3.server_ssl_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2220313001030323-0113120120312012-0312303200101103-1033001000112200-1022312000013333-0232101221333330-3110131333312022-0032212001100100"></a>

<a id="canonical-1311211231221332-0223302000002032-1002212331311012-2332102022222123-0313203020012201-1103022001113332-0220222303022100-3111002032211111"></a>

#### `virtual_server.http3.server_ssl_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0121333110001302-3011222323013000-2320123123203122-2232230320201010-1233200001213303-2210332201313033-1233201001131230-3031332103113033"></a>

<a id="canonical-1113011011032312-1311223212332303-1303103201220030-1132023330230002-2013223323222133-0010303131012332-2233223011231113-0011210321213311"></a>

#### `virtual_server.http3.server_ssl_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3321212230302200-3033132130011312-0223201231333020-2211031231201202-2212202010312300-3121032210113012-0232102002202013-1301020232030002"></a>

<a id="canonical-3301001213103021-2011210132233202-2312303213213303-2200233131010133-2101203031010300-2031233012123111-1132021102201330-3311211101232223"></a>

#### `virtual_server.http3.server_ssl_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0213332302322013-0132020312101330-1222131011201110-0332001323001021-2221123031210203-1102100013030330-3113113211011122-1120001133003210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.tcp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.tcp_server_profile

<a id="canonical-2111211220132202-1113012320222200-1013130022202123-3033202211130113-0002201233323031-3303210132013210-2321202231101210-0233111122201231"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2302003002103012-3320113212333333-2102030111322113-0302232132033201-0301001310323231-0032023321303020-1110321113122123-2201202103331321"></a>

### Direct properties for `virtual_server.http3.tcp_server_profile`

<a id="canonical-2333132203103120-3312032123020120-0320323131013211-2323021212220322-3220223202233133-1210201130022213-3010220310120233-3132313323203001"></a>

#### `virtual_server.http3.tcp_server_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2321221311323011-2003030111232032-3233130311301302-0210122112230333-3133320213120232-2030132302320031-1223012120022102-1111300111011123"></a>

<a id="canonical-0101011121011020-0111113300003002-2302233011303133-3113011131120011-2123210330212303-2023233201210111-0122131001331131-3133120033200310"></a>

#### `virtual_server.http3.tcp_server_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2112112320211220-1031113122132000-0213222102131010-2010111132203013-0313310023333332-1033011133002230-3003220320003310-0302312203222222"></a>

<a id="canonical-0131302130220033-0301112003213003-2303021232213013-3022313010301313-1220212233213303-2012200300223311-3220022302002320-1312030003133021"></a>

#### `virtual_server.http3.tcp_server_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2321221222112321-3200021332030111-3123111333232112-0303003031221302-1002332202323233-3321232232331230-3311221121312121-2301231300000023"></a>

<a id="canonical-1300022101132210-0120321231022101-2303322012011330-0302212330100100-3033132030323220-0310332301010320-0312121121023130-1330233013101013"></a>

#### `virtual_server.http3.tcp_server_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3331011033001003-1230331200212230-1022023101120323-1303102002021121-0131323333023332-0313223021002301-0123303202321220-0103011111120322"></a>

<a id="canonical-2231101121322000-0003231103320312-2213222301133210-2001023103030122-3022120323123010-3321221130212111-0132013221301122-0322010023313231"></a>

#### `virtual_server.http3.tcp_server_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1201123121201311-2131112003333311-3011331232130201-0103122032230311-2013030111222030-1310213111333000-3320122110022121-3121022321232010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.udp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.udp_client_profile

<a id="canonical-3110111131023211-0330232112200211-3112303113213121-0121113033113231-3120200331110202-2211221310223221-2122023033133323-3120120311021221"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2130131231112132-2033232010023120-3301123110203320-2133003001322332-3231031000201001-3122300321210032-1120121033122111-2323232112103030"></a>

### Direct properties for `virtual_server.http3.udp_client_profile`

<a id="canonical-2203200122311100-1333100112021001-3012100102130012-0020003230301123-0232310230310103-3330000102320121-0122232132322310-3003121302130021"></a>

#### `virtual_server.http3.udp_client_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2113002232231032-2002133030000231-1203233301211311-3333013303123110-2020111212300030-3303230223011012-2313110012013321-2201113110023200"></a>

<a id="canonical-1230330322310132-1311333011122321-1210312022023300-0132022311130031-0211331013110002-3120202231012121-0301311020233111-2000001200323002"></a>

#### `virtual_server.http3.udp_client_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2203103302203021-2303332233323320-2303101123233113-3000030102020003-3000100100131132-3022010120133112-3000033120011332-3231211002132301"></a>

<a id="canonical-3210212132303313-2221103022121321-1210231133212103-1213103213323230-3221211212031112-2212032320022203-1000310023101122-3330211223233312"></a>

#### `virtual_server.http3.udp_client_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0200301032120213-3100013023102132-0132212020103313-0222322033333302-1233011110233112-1030311003100221-1110230100220110-0131110212023231"></a>

<a id="canonical-2123032323213210-2310111033101231-0111132122112032-2312232303303320-3033010301233210-1333110213102202-3231121020033312-3010133201331030"></a>

#### `virtual_server.http3.udp_client_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3133021312220110-2030102332311103-2020010201220102-1333302111113001-1133310131021113-2120300102023310-2000011100002313-1211201232232031"></a>

<a id="canonical-0033311013032112-0102132301212030-0012001212220103-3123203123321222-1012122011021032-0232332210023001-3231133131111001-2211011311310233"></a>

#### `virtual_server.http3.udp_client_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3211232310201103-2321212210101101-1222112010313233-0230303023333102-0300133202221120-3311302022122102-1000103222013012-2031210202330200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.udp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.udp_server_profile

<a id="canonical-0003321320221213-2220121121131132-0000313303113012-3330123230300312-3310311321000131-2222322030230303-3311223213331001-2212313203101203"></a>

Type: `"list"`. Computed.

Configuration parameter for udp server profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0230301311223011-3130133001121112-3202102211300300-2300102220331220-2133310003012023-1303022222122021-1302110113121111-1020231111210331"></a>

### Direct properties for `virtual_server.http3.udp_server_profile`

<a id="canonical-0132022331223331-3013330311300132-0133002321200022-1323113301002103-0033112113333331-2300010221011213-1222202212213110-0320112301010203"></a>

#### `virtual_server.http3.udp_server_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2121223101103300-1232331220320023-1310212030013311-1321000112232033-3200020302331210-1132030023212313-1223033320233300-2221003102023221"></a>

<a id="canonical-0110311130221113-0103031320033000-0122011022122113-2111333300303133-1003332213312232-1030121202132211-1302033321313023-1000132000233122"></a>

#### `virtual_server.http3.udp_server_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3113011111203131-0000010322121200-2030300203201222-1221011232220100-3111121130312010-0122222132121130-1213013310112133-0021002211022202"></a>

<a id="canonical-0313020001013212-0122200003323031-2120001333123033-1111331320312311-0333010222032220-2321212213221303-0013201111030032-2130020321033200"></a>

#### `virtual_server.http3.udp_server_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2113201101102202-0111211201203112-2332130131123121-2032111133310112-0220201320123330-0013220230131001-2030312000222333-0210110323330131"></a>

<a id="canonical-1100303323112213-2003001201221132-1203012031002302-3233103222102330-1231310002230132-1103003331310323-3223113121131311-0310100230222203"></a>

#### `virtual_server.http3.udp_server_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1231221011321021-3213212321120213-1011312332333333-3221100201313303-1320021001111033-3202003320012100-3112310302203103-1231022310231323"></a>

<a id="canonical-2132231120001032-0131312233300323-1213233033313031-1102020121110210-2231230033101122-0333010220110311-2020120200313202-0221200332120021"></a>

#### `virtual_server.http3.udp_server_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.https

<a id="canonical-3110112100311110-2102131302131202-3130311013113200-3302111333312320-0231310322001110-1202132200130320-3130103300220103-2113202112113123"></a>

Type: `"single"`. Computed.

HTTP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231232222211323-1020132203233330-1203313120210100-1312323111003002-1021003322010102-1013023330033211-0001211011003311-3213120111031120"></a>

### Direct properties for `virtual_server.https`

- [client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-3001003302112110-3303032312322223-2133130310311231-0131323012021332-0210112201333032-3112021303011312-1000213203131133-0232013202102203): complete subsection reference.

- [http2_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-0232211301312331-3031322222310010-2213321213322101-0030303001230311-2333130203133130-0231333112101010-0230220121312123-0300210202101312): complete subsection reference.

- [http2_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-0133103111122212-0030132222212100-1322222321332221-1120301021330131-2303022021222200-2011233020213313-1331022021303013-2011010231130030): complete subsection reference.

- [http_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-1033020221112231-2312002311211110-1202132231000101-1221103303220301-0002022023110333-1300122031222333-3211212113103022-2332232210331000): complete subsection reference.

- [http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-1122232102120303-3221012223022302-3211320120010213-0130131002322323-0312123133102233-2312312301131020-2020103320100202-1133220133210101): complete subsection reference.

- [ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-3322023320211211-1331022213302121-0133203222332103-1030110033331121-2132310011310012-3201031220331313-1103000123001311-1303223101200210): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-1232131212003202-3233200300332302-3020033131100312-0013102030313023-3233122321231022-1311133030120233-3112020100110211-2033312301303113): complete subsection reference.

- [stream_profile](data-sources--application_profiles--reference--group-003.md#canonical-0210002202230203-3332310302110202-1331311203103002-0130212013211202-2323203303013132-3322010323130333-0303211301321220-0302221122220103): complete subsection reference.

- [tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-2202313221112330-1012021332020313-2032233030332021-3020110300212031-2303012031001112-1120200231113030-3001212033333233-3231212323223302): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2323210303130233-0021121130131301-1201021331112013-3231311301121012-2000003101021203-1210303003100330-0230010033103000-2210101312311203): complete subsection reference.

- [websocket_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-2103323231301330-0000131213010232-3232000232201002-1333132013310011-1100032312012331-0013323330310330-1313023021002211-2002230001300010): complete subsection reference.

- [websocket_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-3012332000033101-1232030021020103-1233321320121310-1012010101221322-0111221200010030-0303213133221032-3302100321231030-3030101313212331): complete subsection reference.

<a id="canonical-3001003302112110-3303032312322223-2133130310311231-0131323012021332-0210112201333032-3112021303011312-1000213203131133-0232013202102203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.client_ssl_profile

<a id="canonical-1132322120022333-0312030202333221-3321011032223110-3301111312201221-1130120122132222-1002223013031310-0302311312132232-2230003122102313"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2233010023111012-2300023010002302-3210230131231310-0310232320022010-2031103211023231-0133213122001313-2000033332013102-3232120023032232"></a>

### Direct properties for `virtual_server.https.client_ssl_profile`

<a id="canonical-1110121312130230-2001103131200203-1002102313133131-3303021103132200-0130032231211223-1300220112302302-1310022311212221-1003213311001200"></a>

#### `virtual_server.https.client_ssl_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3213033313130003-3221300102030120-0100313221202330-2311332320111010-1102210020102132-3000121221120320-2031021232010320-3210222111211002"></a>

<a id="canonical-1203232202303210-3102130320221013-2310101302221230-3212133031120012-1230021103001011-1002310301112133-1233220023331103-0000232322230122"></a>

#### `virtual_server.https.client_ssl_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2033311002330301-2231010213033013-0030022112033122-3332110123102221-0113120203033013-1322003033232020-0311011023132232-3103020330101200"></a>

<a id="canonical-2303201121313303-3020032100032230-3232212200303303-2331221201111300-3020031303010021-1120200132200322-0120303031113303-3123201211112300"></a>

#### `virtual_server.https.client_ssl_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0112111010121310-3102223002232002-1203013211000031-0230311212231010-3021331010111232-0011332030130232-0100211101100303-1331233201010010"></a>

<a id="canonical-2213320220312303-3210012032011012-1203230331131201-2001131101122223-1233222101220303-0301223021112302-2301232120211010-2022202320303303"></a>

#### `virtual_server.https.client_ssl_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3102103202302101-1123202102012103-3301002133331202-2120011031320023-1331100003303233-3322102211013001-3233002212032231-3020020120022222"></a>

<a id="canonical-2213300211023111-2110222312010303-2000230033232230-2002300323211132-0110020321130301-3131122110110322-1000123202322233-3012102111130012"></a>

#### `virtual_server.https.client_ssl_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0232211301312331-3031322222310010-2213321213322101-0030303001230311-2333130203133130-0231333112101010-0230220121312123-0300210202101312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.http2_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.http2_client_profile

<a id="canonical-2300311230302121-0332301030310130-1130120012100320-2113332311032112-0221030233102130-1111102222010113-0033022301021033-2213113030330110"></a>

Type: `"list"`. Computed.

HTTP/2 Profile Client. Client-side configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3021313332003320-1213232023123211-1213000032200011-0102203321021223-1122320333213211-0230023133122032-0003001322230100-0133122230330113"></a>

### Direct properties for `virtual_server.https.http2_client_profile`

<a id="canonical-2203030312232013-3121232202213222-0130200302212031-1332233000303222-3133213310313232-1320020223300131-3103231333212211-2131010323301103"></a>

#### `virtual_server.https.http2_client_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3020012020102033-3031123323232103-1101332333121031-3310203103033000-0323113313003102-2231033220100231-1120320202220231-0021100003102003"></a>

<a id="canonical-3110121111221303-2021222001013221-3201300110323123-0203022233031322-2202001101301213-3103322021013333-0322010323210301-2110231130312300"></a>

#### `virtual_server.https.http2_client_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2031100001000231-1032101003301222-1130121203102332-2132333023131010-1313101131303212-3333120331003113-1210022321033001-2310122203121332"></a>

<a id="canonical-1222233013210132-0121132313220013-0113202123330322-1132113030332230-0111113033231322-1302132302101333-0222101221033333-3322312220003212"></a>

#### `virtual_server.https.http2_client_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2102323113323032-3303022213303210-3110011201110201-3102333032000120-3200233312023010-0222202311102321-3321200121212102-2331311230020032"></a>

<a id="canonical-2233330333001123-0010220023323322-1202221012310330-0300323212103222-3133110013033333-0312023201021321-0011112232302130-1322231200131232"></a>

#### `virtual_server.https.http2_client_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0113223121300303-0230023313301202-0133133133313111-2113020333313130-3011122232313030-2032023212113113-3100203330213232-0200001101003223"></a>

<a id="canonical-0120233102112323-2333030210030003-1122232210330120-3223231322111222-1313103302322020-0323210211302123-3110110130132233-3203210113303130"></a>

#### `virtual_server.https.http2_client_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0133103111122212-0030132222212100-1322222321332221-1120301021330131-2303022021222200-2011233020213313-1331022021303013-2011010231130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.http2_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.http2_server_profile

<a id="canonical-3032313001230132-3323121103023232-0030300210011030-1130303023121331-2203133002010120-0303200302222011-0302021130001313-3122222101303101"></a>

Type: `"list"`. Computed.

Configuration parameter for http2 server profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2130302031110013-2131022202013123-1220113300013122-0230033312233100-0320102101301222-3102201123231303-2210212213300010-2322021202312303"></a>

### Direct properties for `virtual_server.https.http2_server_profile`

<a id="canonical-1202301132213331-1203021030233232-3130223333223111-3232102130330231-3032012001222110-2030210230301101-1100131103020110-3201103013230121"></a>

#### `virtual_server.https.http2_server_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3223330231301021-2100323011202103-0303120122312233-1130131121332113-3302210111222212-0310231232121113-0001022121001002-0313100230201000"></a>

<a id="canonical-2000100200110220-0232123021312111-1111030002201120-3230332323110231-0301113123103231-1031102112002120-3203003103130203-3220013301321200"></a>

#### `virtual_server.https.http2_server_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0033232312013133-0001112331212301-2033100130301223-3023303120230001-1331320012111232-0222311310203001-1103012201231303-3102002200100102"></a>

<a id="canonical-0003111100021321-2221030332121011-3022211212230113-2231202110332210-3313123020020301-3330102300012121-3131320100000021-0021021130202133"></a>

#### `virtual_server.https.http2_server_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3121201231013100-2201323213221001-2113310322022100-0232323033123122-1232133331201202-2133003212303330-2221320121331011-0100133130003133"></a>

<a id="canonical-1222013203300002-2221131212112303-2302302211111331-1002201213321110-2130232030222120-0011322212213113-0231203230300003-2232220221132232"></a>

#### `virtual_server.https.http2_server_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0002203310232131-0133132223210032-3321200332310202-3002203210020131-0203032211210212-0000210210003031-3011322201110133-0231231033200011"></a>

<a id="canonical-3313230013311232-0130100301330031-3110132001320113-1023031122013122-2310013031130002-0332023313321203-3230201101032132-1203113022100320"></a>

#### `virtual_server.https.http2_server_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1033020221112231-2312002311211110-1202132231000101-1221103303220301-0002022023110333-1300122031222333-3211212113103022-2332232210331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.http_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.http_client_profile

<a id="canonical-3131300113300020-0011113001200100-2121210322120013-3211213213122233-0003210302021133-0330020220231003-3331332200203012-1213322003103003"></a>

Type: `"list"`. Computed.

HTTP Profile (Client). Client-side configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1012102133232013-1333323321310000-0223230021200121-3021332210332221-2322331100320323-0102221133110100-0121311013323011-0120030212030212"></a>

### Direct properties for `virtual_server.https.http_client_profile`

<a id="canonical-0311231031223312-1003233110001021-1330321330102201-3310001122320102-0233220002200231-1021012000031103-0223003332101120-1133323203202031"></a>

#### `virtual_server.https.http_client_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3210120202130120-3122200123122223-2101031210003201-1231020113131103-1120112130330211-2313010312302330-0223121021321122-1120033321301012"></a>

<a id="canonical-3222213122031220-1231232322321202-2102000222322332-0102131211312111-0220212231003032-3231330130233310-3321312313333033-0212333203213022"></a>

#### `virtual_server.https.http_client_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0113320320112000-2022020213132233-3201323212030132-3212202201011120-1021323133120231-1130330123031211-3123230000032131-1133011021322211"></a>

<a id="canonical-2231010223313220-1321102200003110-3033313300111230-2230202221122010-0220023120210031-0132310031122323-3213131130021123-1232010133223301"></a>

#### `virtual_server.https.http_client_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3013122132032220-3300200112313121-2133301131222123-3100233121013313-2331231100323012-3101110000001020-3023333113213212-2101030100013002"></a>

<a id="canonical-3321211120132301-1133100233032000-3000220021221321-1030101123313200-2211202102202032-1300323112033132-2323011131030312-0302213331310313"></a>

#### `virtual_server.https.http_client_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0212230031010012-0201111222230021-1302231113311132-3133123131023130-0021121001231123-1132301212233032-2112022201202112-0000022132301230"></a>

<a id="canonical-2000233013121311-0011201302023113-0020013222002333-3301012312023332-3230232122130221-0132130013111111-1212220202233021-0203211331103203"></a>

#### `virtual_server.https.http_client_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1122232102120303-3221012223022302-3211320120010213-0130131002322323-0312123133102233-2312312301131020-2020103320100202-1133220133210101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.http_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.http_server_profile

<a id="canonical-3132302132312210-2020323020222112-0001200000121223-0212232330220030-0003202223331003-2221331033000110-3033303120211121-0231001312112020"></a>

Type: `"list"`. Computed.

Configuration parameter for http server profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3333200222021020-3000330033323100-1002213113320313-1033110302312100-2223031030331110-2211111022023031-2233231101313022-0320231123102000"></a>

### Direct properties for `virtual_server.https.http_server_profile`

<a id="canonical-2313123202213030-3331122332020222-0231131311022311-2030221320303210-3111033333222111-2301100010102000-1303122020202000-1201232312322311"></a>

#### `virtual_server.https.http_server_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1033133132012220-0313302002231332-0322332232330211-3013130103131212-3133331010033001-2113322113310302-3321011000102031-2003003101331213"></a>

<a id="canonical-3030202103020113-3323013313231301-2030321122221201-3210000301231010-2113321022312310-0003130010200011-0222131231131330-1003202322131010"></a>

#### `virtual_server.https.http_server_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0203121001132023-0203020112110320-3102230001231011-2233302203012112-1031230303112010-3202203103000223-0213013312323102-0101212220333122"></a>

<a id="canonical-3330002302102001-0032122020213002-2111102301011122-3123033021012121-3321302313303031-3003122233010231-0202010333022003-0313013111023232"></a>

#### `virtual_server.https.http_server_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2210223023312222-3232112103301023-3000033211200110-3223131332100020-3001303110110331-0003102012120113-0102333211110310-0112210233032223"></a>

<a id="canonical-0220221022233001-0000331233331310-1201020113312221-3213213013102310-3101023131230200-0220303222023002-2013030031110133-1332302102232212"></a>

#### `virtual_server.https.http_server_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1012232303223130-1311121212303020-0211321020020313-3000230031033011-0123023111100110-0102021101333001-0011130133022030-1212310110012030"></a>

<a id="canonical-2320132310120013-0023222011310032-1333021222010130-2033031221223020-2333303300010322-0031122033132201-0100211210222313-3023100103012000"></a>

#### `virtual_server.https.http_server_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3322023320211211-1331022213302121-0133203222332103-1030110033331121-2132310011310012-3201031220331313-1103000123001311-1303223101200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.ocsp_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.ocsp_profile

<a id="canonical-2013002101022013-2331200321301103-1321203302300002-0113011110003003-2233212212030320-2102000221123111-0330031003132102-1112012323101021"></a>

Type: `"list"`. Computed.

Configuration parameter for ocsp profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1332010133201331-1312030103311220-1001010010112113-1323313010200010-1021310310111310-2333222101113212-3320110122211212-0110133313223320"></a>

### Direct properties for `virtual_server.https.ocsp_profile`

<a id="canonical-0333132003010323-1120331001331123-3310322213323130-0001032331212300-1133120131302323-3033121313311120-0211223222112230-3230022231333202"></a>

#### `virtual_server.https.ocsp_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2230023313120100-1330310233232010-1320231330332233-0103000301232031-0130023113011302-2232132112201320-1012220320013001-1032232031122212"></a>

<a id="canonical-2210011303303033-1100220021012233-3032321001221002-1331211000323211-0002213131102200-2321332110321020-3130133321011313-2113031112231011"></a>

#### `virtual_server.https.ocsp_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0331122031032201-2023303021010123-2333311332200113-2011112031123032-3302313123333113-0012233132033133-1030012303030221-3220113120200330"></a>

<a id="canonical-1222210310212132-3130212232222013-1010232031222002-1020121211322212-1301210202322131-3203320222010001-3033330212210302-1321231221303022"></a>

#### `virtual_server.https.ocsp_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3322233310021320-1312220233121113-1020031232221023-3201103031032002-3210033330122110-1013321200010323-0000100322132202-0013012333221102"></a>

<a id="canonical-0223132011022022-3211023320222222-1122202330232110-2002002123132010-3031020021003222-0102131001213202-0121221030310110-1201310200102211"></a>

#### `virtual_server.https.ocsp_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110303230120211-2032002231223301-1320102101230200-2002000300101110-1110113211211010-0030300222021022-1222112133013200-1300000130031200"></a>

<a id="canonical-0130021313001010-2022023221230110-2111002312000003-3231113001223012-0022031332110313-3013323311210313-1311102113123003-3100021300100132"></a>

#### `virtual_server.https.ocsp_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1232131212003202-3233200300332302-3020033131100312-0013102030313023-3233122321231022-1311133030120233-3112020100110211-2033312301303113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.server_ssl_profile

<a id="canonical-3303311223330312-3033332311203113-2213312330332231-2300030212111132-2232201300002011-2011301203332012-3231333312121132-0112202100122202"></a>

Type: `"list"`. Computed.

Configuration parameter for server SSL profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320020203212102-0101312111203103-0022001131312201-0331322332330322-0132211201311331-3130101313310003-0221320023022023-2001033332220000"></a>

### Direct properties for `virtual_server.https.server_ssl_profile`

<a id="canonical-2003120003221002-1223130000200002-1311221231030330-3323013323220100-2001121132220131-0102313123312332-1001220333133001-3021133133120310"></a>

#### `virtual_server.https.server_ssl_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1323230312222023-1102033003233303-3033113201322003-2301312203332210-1220230321310322-2031333012333011-0321233332123113-3101220303303233"></a>

<a id="canonical-1330211202233311-3123011131120313-1120101131310110-0211003233322222-2230033021300021-2232130012330301-2332021033031132-2212203112301021"></a>

#### `virtual_server.https.server_ssl_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0000120220131302-3312322122131013-2230331032101122-2001222121232320-3223201220221111-3231121103122201-3332022131220100-2300223323101021"></a>

<a id="canonical-2311100112230213-0011132232302232-2330012233111102-2203221102223332-2200020120013320-3301300111223123-1001101202313311-1023103323331203"></a>

#### `virtual_server.https.server_ssl_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3203330132232233-1123323113131011-2211210121320212-0102303120111012-3200212011203330-3122333003123130-1223213321023120-3222101303211133"></a>

<a id="canonical-2102220132331322-1032310113033111-0232102310101130-0220000320233011-2313020212302123-0300302213302013-3112110013211100-2131020331203212"></a>

#### `virtual_server.https.server_ssl_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1131310220201110-0200002213220302-0211122122322001-3121111100222010-1021010002120130-2200013233333033-2011122113132300-2300020121320232"></a>

<a id="canonical-0230321322200220-0031311120331123-3313203221033122-0120103102100103-3212223020131121-2322122100213030-1120322313111302-1200131303002120"></a>

#### `virtual_server.https.server_ssl_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0210002202230203-3332310302110202-1331311203103002-0130212013211202-2323203303013132-3322010323130333-0303211301321220-0302221122220103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.stream_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.stream_profile

<a id="canonical-2101112123203110-0231221321022130-3103201012031200-2300212020012011-1230032322101301-3311322301203000-1102100102000233-0331221321313011"></a>

Type: `"list"`. Computed.

Configuration parameter for stream profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0030221133033023-3022123221220021-1323120313213023-0120000121313230-0130220112020301-3033230130302012-2231203013011132-0212333313131123"></a>

### Direct properties for `virtual_server.https.stream_profile`

<a id="canonical-2121222122231220-3003121210123321-3022213100303002-2112102222210331-0122110232203201-2020110010230013-3310202321301001-1100211213332302"></a>

#### `virtual_server.https.stream_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1221211022022220-2320212233303101-1212300121102300-1321012331102123-0003011221003303-1033023031111032-3123122100121202-2200202113010331"></a>

<a id="canonical-1332201321302103-1300110200313332-0032011301000102-2032333321302130-3313323110313223-3300301032030210-0020032221301201-0313222112310301"></a>

#### `virtual_server.https.stream_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1323301301310000-2200120032213330-1233200232021002-3102101333111003-0032113310102103-3322013311201003-1333210302001003-3210030121102231"></a>

<a id="canonical-1221333120210230-1232210130321330-3211323113330221-1322310002130000-0200010202210221-0233300210001110-0010320022321111-2203323201011213"></a>

#### `virtual_server.https.stream_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3112011222113010-1100200331000332-0300133333112322-3020010302200310-0022133213200113-2022023232322223-3023223231113201-0211020312011022"></a>

<a id="canonical-2011013332011221-0102030200102113-0311313033313032-3012333213330023-2112330230220100-0030122000202010-2221302022301031-3310011232323313"></a>

#### `virtual_server.https.stream_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0032123333333230-0013230113220010-1313011122213332-2220233113113320-2323333000130303-3202120010003011-2211030223311031-3102301001030012"></a>

<a id="canonical-3221310333002121-1121202101220113-1030122202131133-3223122103012011-0303303011213213-2012110013322211-2321300123123103-0330313303310101"></a>

#### `virtual_server.https.stream_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2202313221112330-1012021332020313-2032233030332021-3020110300212031-2303012031001112-1120200231113030-3001212033333233-3231212323223302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.tcp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.tcp_client_profile

<a id="canonical-3311100221020133-0023232010132211-2010011331121213-0031000022333011-2010010233030320-1311003313310322-2210321223120132-2003213033101022"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333233132321111-2332011013120021-1100010102330212-3322031131221132-3032210221123331-0103130001130311-3311023323332303-0202212020230021"></a>

### Direct properties for `virtual_server.https.tcp_client_profile`

<a id="canonical-3000013103023010-2330212310323033-3331201020222200-2300032002011121-1030111032321030-3123233320330320-1120003212313332-3223330333223231"></a>

#### `virtual_server.https.tcp_client_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1103011220011233-0321202113110021-1233000302213123-1330032120031301-1032103112102200-3213311102120310-2131001220211321-1312231201322121"></a>

<a id="canonical-1123201212233323-2322032110211301-2220302033110202-0230120233333011-3111320233131013-0220311212321210-0301030000120113-2000032213210130"></a>

#### `virtual_server.https.tcp_client_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3232010022121310-1211101223223122-3211031002300112-2221020122011223-2223232313022020-1000003023320233-3332211113001112-1133001100310232"></a>

<a id="canonical-0123201033311200-3121230032111102-3021022220311100-3221011031102210-1202020123001223-2332331303033021-1322212113113321-1000110122102132"></a>

#### `virtual_server.https.tcp_client_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1010202122201212-3223101002323311-3323013011303332-2201003313100330-0002331213203202-0210133211322131-2013211312212333-0010231013133300"></a>

<a id="canonical-2021103222202101-1112120123210223-2330233132000222-2012113131032300-3212322300030112-3222213303130122-3311001231333012-2231321313123100"></a>

#### `virtual_server.https.tcp_client_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0020322001231033-0020201001103000-1130013231123021-1213213011110002-2313300100001021-3202223011030302-1003323133110233-2100012113311032"></a>

<a id="canonical-1001123113122001-1031300321320201-1201211302332023-3131032003231301-0210013022330021-1322333133312032-0200230223000021-1322321123312330"></a>

#### `virtual_server.https.tcp_client_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2323210303130233-0021121130131301-1201021331112013-3231311301121012-2000003101021203-1210303003100330-0230010033103000-2210101312311203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.tcp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.tcp_server_profile

<a id="canonical-2333122203230102-3331000211011113-0031023332320321-1200000031100101-3333133213321022-0131033020022213-2301223321102130-1311030113100202"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2220202023003112-0030212110020111-0202012132021112-1312222210222201-3021330333121202-1333021012131310-3012012222332313-0102003121103130"></a>

### Direct properties for `virtual_server.https.tcp_server_profile`

<a id="canonical-0013012323002312-2102331123231221-3321222103101202-3223212101113302-0001220321120300-1313303203230220-3100212321213302-2032130012332223"></a>

#### `virtual_server.https.tcp_server_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0101010133201312-1311303112221210-0313101110330203-0121102123023200-0311013002111200-3200231221011232-3330311203333100-1000320201312101"></a>

<a id="canonical-0121033123210223-2130010013032031-3121010301030201-2101311301322300-0333220323033002-3230301112030132-3033332111312021-1313212003013302"></a>

#### `virtual_server.https.tcp_server_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3213110330021022-1300303210103311-1200130123220003-2223303111133100-0022030323301221-3303203201230303-1100221103300031-3132030201310202"></a>

<a id="canonical-0302113233122103-1333231020113313-3101211120231011-0113101103330132-0202133331103212-2113100001121222-3221300210102112-1020220210033110"></a>

#### `virtual_server.https.tcp_server_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3332321010023303-2101121321323103-3022220032202230-3223230130111301-2311310131032001-3123030223220033-3010200221002202-0203201330120022"></a>

<a id="canonical-1102001223003030-1222213301300012-3300223123313031-0232213013002003-1130321301230132-3231211302001301-3102310223113330-3200331131010022"></a>

#### `virtual_server.https.tcp_server_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1223313010313313-2333222130310231-3023011022002022-3203010033002221-0201112013000013-1102301122020330-3111031123131230-0200330302303321"></a>

<a id="canonical-1111022020112023-1033303323330032-3301102301030200-0120011230022001-3212112330201131-0123130330132110-2310333131012123-3122112000222130"></a>

#### `virtual_server.https.tcp_server_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2103323231301330-0000131213010232-3232000232201002-1333132013310011-1100032312012331-0013323330310330-1313023021002211-2002230001300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.websocket_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.websocket_client_profile

<a id="canonical-3110112022222213-3031103003231113-2210033100330223-0310132300103132-0223331232220230-0023202223300301-2213232333332010-2300321331131331"></a>

Type: `"list"`. Computed.

WebSocket Profile Client. Web-related configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1012233031121320-0232311032320201-1132331211023020-0023003120023122-1012132133133002-1313033120300301-3232200233031331-0013202201000323"></a>

### Direct properties for `virtual_server.https.websocket_client_profile`

<a id="canonical-2131032331210311-0021210032233031-2212110110323130-2003022003203300-0220211333330021-1213012133012212-2313323222131102-1200120200323202"></a>

#### `virtual_server.https.websocket_client_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3330323230322020-2012031130030030-2113120330302100-0213200231332113-1013003120101320-2233132220023213-2213003311230130-3133031110103210"></a>

<a id="canonical-1112033120223212-1113320133101111-2212220320121321-0023312231133220-3210000211111120-3321332221021222-3210133332000310-0322120011131233"></a>

#### `virtual_server.https.websocket_client_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0110203121212333-2032333011013122-2231111302201211-0023122303121011-2120131233032302-3100031303021212-2210132220213330-1231322001022102"></a>

<a id="canonical-2103031200020312-0312213312212332-3230030330121131-3313301032011210-1022210030000003-0112313032321022-0333020230201302-2231010211122112"></a>

#### `virtual_server.https.websocket_client_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1010103130330300-3233000010210011-2300111333231231-0002203032303110-0323123333311023-3210313102301312-1222113131321323-1003020200223120"></a>

<a id="canonical-3200101033020033-2000122232030301-1331111211202123-2312012013213311-0103203330021223-1000222102013203-3302000333003221-2301102131210031"></a>

#### `virtual_server.https.websocket_client_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1303232103101303-3000223033320023-1032313332322333-2002313232312132-3023100201120110-0112023020122122-1032321302023332-3332313012112301"></a>

<a id="canonical-1212110010332013-0131131210113222-3112220001121221-3332221123003300-2203030103311010-1020121323022130-0302121323112200-0213003112322203"></a>

#### `virtual_server.https.websocket_client_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3012332000033101-1232030021020103-1233321320121310-1012010101221322-0111221200010030-0303213133221032-3302100321231030-3030101313212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.websocket_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.websocket_server_profile

<a id="canonical-2012111301231322-1111322030030322-2111202200121021-1030233113133301-1101212311113022-0101312003203001-0210220000030302-3210332010000033"></a>

Type: `"list"`. Computed.

WebSocket Profile Server. Web-related configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2220200011322113-3330101210313322-2202223320022102-2301302310031031-3103122100211023-2210311122110100-3202102300111320-0210101133112100"></a>

### Direct properties for `virtual_server.https.websocket_server_profile`

<a id="canonical-0113033232120111-1002201022313210-1030311130010023-0330023210033321-0321320202020100-0113311232230321-0011102320232210-3231213010232232"></a>

#### `virtual_server.https.websocket_server_profile.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021000021011333-1233331232122330-0220113002301313-1002332003100013-0122220032220202-1032332031021002-1332013333133021-1232312100301023"></a>

<a id="canonical-2322103002211333-0110011132103030-3333001122033212-0200233222301202-3132213330233113-3312131333112113-0301310002122110-0321220012003331"></a>

#### `virtual_server.https.websocket_server_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1231220022320220-3302223200010122-0200303000022131-1112222231322113-1203111323210202-0101212310121011-1110132012210213-2303211230123023"></a>

<a id="canonical-2013323111320222-3010022332312100-3213211320312233-2202210131101221-1031312101310133-3031010230303313-1230120123003321-2221203213033213"></a>

#### `virtual_server.https.websocket_server_profile.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0012033122103103-3012233012033213-0001202231333010-0301331321100133-2002120120102011-2212122300310322-2123102231131002-1031133331211032"></a>

<a id="canonical-0333123013130223-3033303331321320-2321011100103220-3323121312111302-0013133112323303-0013011032211300-3020233202302120-1312302222003322"></a>

#### `virtual_server.https.websocket_server_profile.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0200303012122222-2213200013030120-2112010231302332-3023210330121221-1212123022222013-2200211110011231-0331213032113332-3031132301122001"></a>

<a id="canonical-1022330313111331-2031100302003202-0022030323003003-1101023323113321-1303032300011001-2203200302330331-0333321320212023-3132111110130021"></a>

#### `virtual_server.https.websocket_server_profile.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.immediate_action_on_service_down` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.immediate_action_on_service_down

<a id="canonical-0210331121211012-2203311120332113-1110330020220232-1213101202321223-2133202200212011-3301130210013120-3000113220233111-2331332203033110"></a>

Type: `"single"`. Computed.

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.
None: Specifies that the system takes no immediate action if the virtual server is reported Offline
or Unavailable. Reset: Specifies that the system resets the connections when the virtual server is
reported Offline or Unavailable. Drop: Specifies that the system drops the connections when the
virtual server is reported Offline or Unavailable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-immediate_action_on_service_down_choice": "[\"immediate_action_on_service_down_drop\",\"immediate_action_on_service_down_none\",\"immediate_action_on_service_down_reset\"]"
}
```

<a id="canonical-0111202103020311-3330310301311111-2320303001030231-0223030332233222-0322110122003112-3010002001331131-2010313223221212-2210133010323321"></a>

### Direct properties for `virtual_server.immediate_action_on_service_down`

- [immediate_action_on_service_down_drop](data-sources--application_profiles--reference--group-003.md#canonical-0330110231211133-1013300112302232-0032121321233120-0110023110200023-2011011311313223-3303032311213000-0323031123013102-0310120212202120): complete subsection reference.

- [immediate_action_on_service_down_none](data-sources--application_profiles--reference--group-003.md#canonical-0320222200030133-1100011111332202-3023021021122020-2220000330011332-3001111211301001-3013213030203230-2101133123011203-2323110021001012): complete subsection reference.

- [immediate_action_on_service_down_reset](data-sources--application_profiles--reference--group-003.md#canonical-0223021032303220-1102232003130310-1302011323122020-3231023111301112-3210023001202313-0320211021220022-3320013113301310-2323013330213133): complete subsection reference.

<a id="canonical-0330110231211133-1013300112302232-0032121321233120-0110023110200023-2011011311313223-3303032311213000-0323031123013102-0310120212202120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop

<a id="canonical-3233003231010201-1213111003322323-2221030030111111-3203212212203322-2322001311323312-0200002223310232-2310212013022231-1232232032330123"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-0320222200030133-1100011111332202-3023021021122020-2220000330011332-3001111211301001-3013213030203230-2101133123011203-2323110021001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none

<a id="canonical-3000331313012003-1313113120221112-1321320333220223-3223310230103233-3322133303302333-0133332030023102-1120203222130030-3333112121013130"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-0223021032303220-1102232003130310-1302011323122020-3231023111301112-3210023001202313-0320211021220022-3320013113301310-2323013330213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset

<a id="canonical-3031312321022010-3333033133022131-3100131033230303-0330010323330111-3111100013013201-0013013212012130-1313113331031321-3002313232001233"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-0000301021323203-1003301002002311-0310011213320313-0102003200331211-3103021001233200-3212113122100001-3320222320102111-3313223020133023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.last_hop_pool` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.last_hop_pool

<a id="canonical-1331332331330133-1331300103320202-1223101301311323-3321312330210100-2311310201321032-0223302031132302-3110001232223011-0133002310032201"></a>

Type: `"list"`. Computed.

Directs reply traffic to the last hop router using the specified pool.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1030023320201230-1002302213021210-0300321010113013-0232131103032301-2312031010203322-2320323120213213-1132320110211001-1221023100023133"></a>

### Direct properties for `virtual_server.last_hop_pool`

<a id="canonical-3320230021230312-1112302231000133-3132223210220133-1333213323222103-2023332210021111-0103333110320202-2121301203201323-1310211133223332"></a>

#### `virtual_server.last_hop_pool.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1100132102022032-0321130101023003-1222000130122110-2213212022000021-3122321123032201-0033210123000313-0322011202122302-0203100331021230"></a>

<a id="canonical-1203002103030030-2100030121131023-3003111132331322-0110131131030100-0013310230110320-1322321133013201-1312211322021332-3010113300332222"></a>

#### `virtual_server.last_hop_pool.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1022223331012211-3012300010010122-1133012302110113-2111323031231331-3333221002320111-1131010121320023-2223003313103221-3213220311111132"></a>

<a id="canonical-2212221101121321-2001300131122310-0012313002320213-3332321031032030-3222200320013022-1112132312231211-1100230203122233-3030303213300213"></a>

#### `virtual_server.last_hop_pool.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0322032110112232-1120333001010313-2222111031233223-3300212301332020-0232023133213333-0010200330202303-0022102030021102-3312021210321200"></a>

<a id="canonical-1123220103231031-3220032021003312-0330032022010200-3233031330321232-2022221120200112-1221210300132333-1121103013200031-3222003020121323"></a>

#### `virtual_server.last_hop_pool.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0222231132230310-3110230303320121-1133232113122013-1301213010002020-3310012202021002-1233230323311331-2202231211123133-2030012323232211"></a>

<a id="canonical-3230021113123010-0203020332220112-3303101333203100-1222211031200110-2013201031002022-1000232320221123-2200121000132131-3022112100211302"></a>

#### `virtual_server.last_hop_pool.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.nat64` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.nat64

<a id="canonical-3112222212333220-2111210012320301-1230013102331102-1113102012213010-1223310022222321-0001023002031323-0231201223233030-0321201011113133"></a>

Type: `"single"`. Computed.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-nat64_choice": "[\"nat64_disable\",\"nat64_enable\"]"
}
```

<a id="canonical-1312232211012033-1001322312123320-0101330122212001-0033002010213211-3202333221111122-1231023111001113-3000133331002211-3123221001232230"></a>

### Direct properties for `virtual_server.nat64`

- [nat64_disable](data-sources--application_profiles--reference--group-004.md#canonical-1022331310002101-0221122112312111-2032231100313013-1022031020210202-3312201210321013-1211121020302023-3201110001133330-3003133303132200): complete subsection reference.

- [nat64_enable](data-sources--application_profiles--reference--group-004.md#canonical-2013103112220012-3012000211330010-0010213022332111-0130021120033031-2221020031113033-2333133200123101-0203122230130021-2130102133201003): complete subsection reference.
