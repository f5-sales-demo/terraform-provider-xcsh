---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3130032003320203-0303103201031010-0123021130321021-1320113320022211-2103132333323320-3203123223020321-2323020212100123-1112112321213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-2131001030302230-3033011120331123-2302011200302313-0020322221023020-0013321313132120-3233101001313200-1022300022100300-2023303233112101"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-2020023331233203-2331031010030220-2030101132200111-1033210110121221-2211330221032022-2113312133332231-3223021322232102-3113310002212112"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.asn_list`

<a id="canonical-1001312112023123-2321102120013033-2013322322012320-1211222121103033-2122122120121103-1113131012130231-2021003000302002-2023231130230211"></a>

#### `ddos_mitigation_rules.ddos_client_source.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-1221200003223120-1123132330223022-0123031300233000-3032123301003023-2232023110333223-0020303131102133-0132010200130211-1321000101133023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-2330112201213102-2332311011222131-2220222303022201-3003332312233311-2020202022132032-0223330120232232-3201001203012133-1200001130130131"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ja4_tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132132300131323-0102112031323201-3112101130201031-2300020113220303-0300323323302002-0012200113321021-3222013213203123-3003033032022323"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher`

<a id="canonical-2030213330321333-1220012102102013-0202312332323302-1021322230100030-2100010211013122-0212122123330333-3331313121331102-0023320101311330"></a>

#### `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0232131223030222-2002023001132031-0113202320302321-2230333310201010-2203102312111021-1320300323110233-1111222133230010-0013202313112200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-1102221121013011-3220002311310113-2310023111232010-1023221121322301-1222121112210132-3010110220310101-2022303013232110-3111220010113000"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2121020020013033-2321200221221111-0211003313202302-1212300313100030-3003330210121011-2103211102011310-1323310202121312-0033221111321100"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher`

<a id="canonical-0010010101301122-3312122120133102-1333210011002112-2133310322133130-1300201103030132-2323013033012310-0113013023330010-1012023121011130"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2112021223101233-3033130211323103-2321332103123203-2102111211301002-1202202211210311-1123131131201101-1113121122123003-0231112331120033"></a>

<a id="canonical-2200333023012102-1112330003001022-1330122013113322-3033302232031232-3020131321312123-0013313120320033-3310123332333110-1032122201303113"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0212333233221133-0013123000030001-0100321001012303-2130030231223232-0303312122221122-2131300021303132-0200202302212012-0333333131300101"></a>

<a id="canonical-3113031220100012-3112213320121332-0300310300310112-1213323231003222-0012330232201212-2333333011111220-1023012333323101-2002013031002101"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1032202033010100-3123021202033002-2031020103231302-3003123132103022-3022131311221333-2033223333113210-1001110012203101-2113133220312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-1033120112120123-3030222311123121-3110333321012100-1311212100201231-1302312303232100-1300222123331023-0103100311321223-2232111020312231"></a>

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

<a id="canonical-1100202021300302-3032322313031332-1101123203202120-1110201103110011-2233310303301313-3000332122002002-1200301031110003-1233113311100020"></a>

### Direct properties for `ddos_mitigation_rules.ip_prefix_list`

<a id="canonical-0223213210203110-0203003323321110-1323210201300032-1330101122031233-0002033211302023-3112102232301300-1232130220131123-2112220112101322"></a>

#### `ddos_mitigation_rules.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="canonical-1103211303322330-3012331002310301-1013320011031010-2102123213122321-0301001031332201-3333212202111330-0121101112112322-2311123101230333"></a>

<a id="canonical-1323210203102320-2223203100211202-0323203330313002-3000200211030330-3323123121322233-2221332300033220-3211132102333333-0113120333012200"></a>

#### `ddos_mitigation_rules.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0021200313320233-2232231332010313-3112121020221101-0220223211022102-1131321221322321-0330021023003303-2230313223012220-2031330101232012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- ddos_mitigation_rules.metadata

<a id="canonical-3130323213113131-2301230133203230-3303200020112113-2320123223233023-0201312302132310-3323120331212110-3021322330303313-0012132123012301"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302231012230103-3122210330013310-0333000330023133-0100101312333103-0212212133223201-0002122110003003-0111123323201223-3220201130103321"></a>

### Direct properties for `ddos_mitigation_rules.metadata`

<a id="canonical-3011222013131130-1231031301131211-3211200033031200-2003112302112023-0021332101100121-0222010233212201-2002130121313210-2002211230122320"></a>

#### `ddos_mitigation_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2121101120131131-2230021113221212-2202022012112302-0011102023302320-1000113331332330-3001221003220301-3313212331020102-3320323211033321"></a>

<a id="canonical-0002223133331020-1111101030312102-3001302233203130-1112320321113333-3233211232102121-0332103132211300-0103223212211220-1331112220012130"></a>

#### `ddos_mitigation_rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2300021122222120-3103201200301030-2101110210210123-3021222103220303-1221111131320223-2202110300233100-2330133310101313-0202223013312313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_cache_action` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- default_cache_action

<a id="canonical-2033212333122032-2312110112111001-1233312213130130-1103312010323210-1023100121203202-0231332123122302-1303001230110013-2012231230210121"></a>

Type: `"object"`. single nested block, Optional.

Default Cache Behaviour. This defines a Default Cache Action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

Terraform syntax:

```terraform
default_cache_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030222233111300-2331213302201233-3321013002010200-3033210131133233-0110330202020130-0312030130132303-0010221020232332-1100111013012103"></a>

### Direct properties for `default_cache_action`

- [cache_disabled](resources--cdn_loadbalancer--reference--group-010.md#canonical-1112300232302300-1311300203022231-0122331013122023-3331202321321013-0223203021300303-2330123312133031-2330013031221033-2100200031312010): complete subsection reference.

<a id="canonical-0113101211112232-3111112303311302-2311330203121223-3133121201110322-1100320032022313-3002032121031100-3201330222303223-3211133132013331"></a>

<a id="canonical-3100001001201001-2332020232003031-1012332221123323-2032021332233001-2322021202002300-0112020110121001-0112020001301300-0221010000123303"></a>

#### `default_cache_action.cache_ttl_default` property

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-3201122112033133-2012200101300303-1232133123113222-0011121213202322-1320123013120223-3201003001233002-2333312131003320-1012130032010230"></a>

<a id="canonical-1300103310231322-3021321313000110-2203300220301301-2200030112122100-2322212010001102-1133113001130012-1312103021311013-1321311302330312"></a>

#### `default_cache_action.cache_ttl_override` property

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-1112300232302300-1311300203022231-0122331013122023-3331202321321013-0223203021300303-2330123312133031-2330013031221033-2100200031312010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_cache_action.cache_disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [default_cache_action](resources--cdn_loadbalancer--reference--group-010.md#canonical-2300021122222120-3103201200301030-2101110210210123-3021222103220303-1221111131320223-2202110300233100-2330133310101313-0202223013312313)
- default_cache_action.cache_disabled

<a id="canonical-3211233323103013-3220301332022022-2330331301333122-2012220120322331-3231022332000232-2123011210000013-2122232222332200-3021201022101133"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
cache_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310301221200230-0201112131201231-1131101200303021-2130030310111230-3203022302002332-2030103200001230-3233113321131031-0233232233132333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- default_sensitive_data_policy

<a id="canonical-0232201112201301-3232122211222210-1201200012003200-2231121011122220-1210020311103232-1111010111310203-0231021103002332-1301120220213033"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature.

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

OneOf alternatives in this subsection:

- [default_sensitive_data_policy](resources--cdn_loadbalancer--reference--group-010.md#canonical-0232201112201301-3232122211222210-1201200012003200-2231121011122220-1210020311103232-1111010111310203-0231021103002332-1301120220213033)
- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300020121333330-1303221303112300-0023322213301102-1002303223322122-2130031202131001-1332232121231102-1300111112133313-2022220010223230)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sensitive_data_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312231102230133-0002321102020331-1031013220300023-0101221003122222-2100330313000330-2331331200001221-0133311233301310-3133210321022010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_definition` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_api_definition

<a id="canonical-1013231301331321-3201321033031021-0303223333210301-2112110233330302-0302100120123301-0132331033332233-0312121102001302-1222030220330231"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_api_definition = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130210132113322-0123113132000212-3230220223101021-1200003113211031-0113012202110012-1311330331231000-1333100002101020-1131001031331003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_api_discovery

<a id="canonical-1322032121300031-1321200030022101-0230311311232112-1223332321131322-3320303002032002-2300111110032023-0301130102003102-3121221211001211"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-1322032121300031-1321200030022101-0230311311232112-1223332321131322-3320303002032002-2300111110032023-0301130102003102-3121221211001211)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-0001201201333322-1133021323220003-2333030001222131-0223013300101113-2133221013222213-2220202223030213-2303012213101101-0010230023302110)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_api_discovery = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322312112202232-3232301311330110-0300133101202313-0011202213000301-3113120123102301-0132223030100120-2202020201132330-2132133022103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_client_side_defense` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_client_side_defense

<a id="canonical-2323022203312322-3012211311100012-3031300132031113-3310223011233211-0120103103100233-1101032030002322-1221023301030013-2002330013332333"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_client_side_defense = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031001202212121-2320213202323013-2310203020210132-3222031002011331-1231231310110332-2313123223113001-3022212102330213-3000321230000321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ip_reputation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_ip_reputation

<a id="canonical-0232112311102100-0012331211221212-1001112022113233-1022321313211112-0223113010131232-2221332211301003-1223003001103130-3211023200221123"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-0232112311102100-0012331211221212-1001112022113233-1022321313211112-0223113010131232-2221332211301003-1223003001103130-3211023200221123)
- [enable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-1123223010010320-2010330211102303-3003132133301111-3013321110233132-3220003300022030-2113002030033003-3221112313032202-3320311130101320)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ip_reputation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323303332010223-3130303213321301-1233002121331100-0110203021132310-0033302313301323-2313122210131232-3213230202131103-1322110213131001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_malicious_user_detection

<a id="canonical-2333103103022200-0210221203102002-1120233210310133-0232012330110113-1331222330323203-0033030202200212-1220213202003111-1013012231033112"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-2333103103022200-0210221203102002-1120233210310133-0232012330110113-1331222330323203-0033030202200212-1220213202003111-1013012231033112)
- [enable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-1123232203103010-2221320123213112-1303032131321013-2032233321303102-1212120310102022-0332230300221210-2020311101222333-2212332330213300)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101211221210102-0323002031331010-2210220231120012-0023021303231032-1032303023320223-0003233332001012-3103103110223321-1232203302313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_rate_limit` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_rate_limit

<a id="canonical-3133211010101322-0323301002321331-2233011002232323-1123100210322211-1330323131213103-3011211112311233-3123030230331311-1230310111020221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable rate limit.

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
disable_rate_limit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300001220010210-2011013010302113-1031230110020223-2000000323133201-2331022100022333-0022210002101020-3232010103322200-0011200213203231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_threat_mesh` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_threat_mesh

<a id="canonical-3212121021213202-2312330010331332-1330021101321323-3001103211201202-0231201210232013-2112333212011003-2330002011233301-1323122311231320"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option

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

OneOf alternatives in this subsection:

- [disable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-3212121021213202-2312330010331332-1330021101321323-3001103211201202-0231201210232013-2112333212011003-2330002011233301-1323122311231320)
- [enable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-2123213233303030-1130113101312333-0212210331021330-3103233332320113-1210122212022213-2121031100211332-1031010102222213-2210032031223013)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_threat_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120023301221013-0010101221332312-2101301230211303-3223121323013122-2033123211312311-2201011122202012-0103112113322002-3003131323102011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_waf` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_waf

<a id="canonical-0330321120112012-0021333212231113-2232020002320311-0200223111331021-1210113213102002-2101003211223303-0223321332020310-1010211300000211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_api_discovery

<a id="canonical-0001201201333322-1133021323220003-2333030001222131-0223013300101113-2133221013222213-2220202223030213-2303012213101101-0010230023302110"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032100310032213-1102211120033313-0001323013010022-1302112103220302-1022012200220002-3233212132031231-3233210100332123-0030110213020220"></a>

### Direct properties for `enable_api_discovery`

- [api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302): complete subsection reference.

- [api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100): complete subsection reference.

- [custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212): complete subsection reference.

- [default_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-1013221211221310-1031120002332131-0232001133000302-3000130332131232-1303121022332312-3012002312133111-1112121330233321-2233231122001212): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331120230102020-0312102032323311-1020022132220113-1312101223301033-0301110002231123-3321223213133301-2033313330221330-1323102320001301): complete subsection reference.

- [discovered_api_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-2112030012103111-1212023301022210-3131303031130201-1232223310233020-1021210103230301-2110223201003332-1112033210222011-3100331330221113): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-0021233123301122-1013002210302002-1231011123110332-0101131021133112-1321303230003032-3032101231010222-0000133020332203-2201003101132010): complete subsection reference.

<a id="canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.api_crawler

<a id="canonical-3003102010320103-2001002303320310-3031102300112312-0303330213100202-3201321130230300-0131131021000230-1122313002322013-1302121212220232"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031132220122031-0022213112110331-1300210103302331-0333231332031201-1031012011002123-0221113103232310-2001210031322300-2330331200001333"></a>

### Direct properties for `enable_api_discovery.api_crawler`

- [api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021): complete subsection reference.

- [disable_api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-0223301322310022-2112210323323023-1010031200031021-3002003330300110-1311032202200232-0120121012010001-2331132132312022-1323130100110212): complete subsection reference.

<a id="canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-2122202003320020-0120333300302123-0330013012112013-3122023312102132-0303132002021121-1321312133200321-0013223212032013-0121323033102002"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

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
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331013313010023-2103302322110301-2003313023111100-0323132332132331-1302331233222002-0103111131310033-0022010100101113-3033023012303232"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config`

- [domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333): complete subsection reference.

<a id="canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-1033131131213212-1333112301230031-0120220313223123-2133321012032202-0011320302323302-3120230022131033-3022110333111010-0103110231112020"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202121031203013-3020003123123203-1311000313223033-2003200322203130-2023132232330103-1330303331223202-0330020120300001-1000030120033213"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains`

<a id="canonical-0230233313133232-1022303320313013-3330131233221032-3031021233230113-2312101102332113-1221321233132122-2223212030321213-3300031122011230"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.domain` property

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111): complete subsection reference.

<a id="canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-0302213302123230-2232310000301231-1300002000210303-1003320103213333-2211120131213223-2231201103113013-0030131101322300-3020333102303012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

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
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333121220112103-0220323013130101-2132212203011331-1032001130222013-2212013202201201-2100222120331313-2010322032110212-2111011200320113"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login`

- [password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220): complete subsection reference.

<a id="canonical-1323013001013013-0030010300130220-3330033320220123-2021223133110032-3022011023130012-2132311110100331-3011120230032211-3220130020310033"></a>

<a id="canonical-2302223212201221-2031223103123001-3030111022303031-0301031011211330-0101122020223332-3132110311033030-2310122000002233-3130002022210330"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.user` property

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-3332101023013122-2212130303321220-0003312321103212-2113013211200333-3330010232230023-2000233330103221-3332003232332023-2021330021010222"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031310133211012-3120302002233312-2131031112020122-3213212332232120-2111032321111202-0120230322300111-1300232200311121-0230122102312031"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-0203021132213020-1023230023031320-0001220210233022-3131201223200301-0332210022120002-2313010213221231-3201102221031211-1223020110012220): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-3011333222233131-0113333113011123-3012100120011302-1112101232033213-2032213333000311-1223123111101003-2131011023013300-1103130003100103): complete subsection reference.

<a id="canonical-0203021132213020-1023230023031320-0001220210233022-3131201223200301-0332210022120002-2313010213221231-3201102221031211-1223020110012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-2301201131312000-1123301133012000-3331211200230211-0102332200330130-3233103232202110-0321211311331013-2213112103223031-1012101110213333"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-3331121330300001-0003123232231321-0200002313033202-1211300101031010-3220202131323310-1012101132312001-0210201023113202-3012032102103303"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info`

<a id="canonical-2302132232012032-0131302000213103-3031312331231223-2000331033332023-2130212201213231-2222310131110302-2120322023222313-0012320233313013"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2212330103001132-2211032113010132-0202331221223110-3023001300132223-2013202312232331-0011302033030313-2323333013101312-0101223010033222"></a>

<a id="canonical-0120021220331211-0030311303201331-2121032301030301-2003001221113121-0203002311221222-0310123310132130-1021010201103313-2110220223033122"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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

<a id="canonical-1301231330230201-0021111020130311-3131212203310110-1313132211020121-2332330230111303-3200033313330111-3102001301200220-3010222202110003"></a>

<a id="canonical-0122120210330301-3202022230022122-2000103012112022-3101231300301033-0320323030232221-2100100001211120-3313213001213100-1012013231220300"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.store_provider` property

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

<a id="canonical-3011333222233131-0113333113011123-3012100120011302-1112101232033213-2032213333000311-1223123111101003-2131011023013300-1103130003100103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-2123113313220110-0310302301200313-3212113121210233-3200210232310020-0110010012200211-1300033200023320-0202132210212032-0010330003311023"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-3200221111202312-0013223323302122-0032030231030120-1230023110130233-0303311120220330-3211233013221030-1210003301033302-3231220123200101"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info`

<a id="canonical-3322003002213231-0230310333030102-1201313103300313-2312120312212110-3320222111221010-2231130331313002-3003011231302320-2333332011012303"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1331210200012223-3311013231010113-3302111203111003-1313113203320212-3203012011122210-2322112310002101-3212322121102220-3330313223320323"></a>

<a id="canonical-2320312111013312-3221230123132013-2103233102121111-0012122122130300-2211310022103323-1103013122200030-2033032113211323-2120323003203110"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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

<a id="canonical-0223301322310022-2112210323323023-1010031200031021-3002003330300110-1311032202200232-0120121012010001-2331132132312022-1323130100110212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.disable_api_crawler` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-3133131002302221-3223111311103332-0120003223003223-0020102313300121-1303332123233333-0312011031011130-0021320031010010-3112031301233302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_api_crawler = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-1012300311230310-0001121003022332-0223030222213323-1030300202303311-1300332101121211-3123203113211231-2221201300122130-1311020030100120"></a>

Type: `"object"`. single nested block, Optional.

Select codebase and Repositories.

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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302121320122110-0332000002303032-2001131200102122-3010221110120102-0302013002033123-0031111313103323-0330001323121113-2333111301322310"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan`

- [code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031): complete subsection reference.

<a id="canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-1133201031102003-3131001102211330-2110133203202212-2310302321232220-3030011203110303-2113121210013301-1120112103332321-3031023312012223"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for codebase integrations.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
code_base_integrations {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333013212300302-2132321122200333-1100313110333221-3232313110102101-2330010333110330-3323101021203203-1002330132231123-3221310031313111"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`

- [all_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-1133222312233233-3230003202230132-0331020231301303-1013330021223021-2003332220203123-1333213033231321-2101011233031021-1301003213132232): complete subsection reference.

- [code_base_integration](resources--cdn_loadbalancer--reference--group-010.md#canonical-2121123213033030-0101132301023222-0210321103113032-3130313132023121-1023321313310131-1122011031001221-2122200023311221-1003322232210223): complete subsection reference.

- [selected_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-3332003330020233-1321203101311112-1032320222112212-0303123303022320-0130233233212212-2331122210022100-2012123323130120-2300212333133331): complete subsection reference.

<a id="canonical-1133222312233233-3230003202230132-0331020231301303-1013330021223021-2003332220203123-1333213033231321-2101011233031021-1301003213132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-1000331013330222-3000113103113211-0001211010301300-1033013110223123-3221120202333012-2132002301113030-1001000130303322-0320300120233121"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_repos = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121123213033030-0101132301023222-0210321103113032-3130313132023121-1023321313310131-1122011031001221-2122200023311221-1003322232210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-3200022321013010-2212130332311312-3132123120032023-3231000103301002-2121202022103002-3030130133310002-3310023202310011-0330112210321032"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111021330021332-0020032120013313-2310231313101220-1321010210201303-0000321301311332-3102122011332303-0023301101312010-0101311323021131"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration`

<a id="canonical-2220112313013320-1132100032313302-1000023013203022-2111020010001130-1013130133110121-3101130033123200-0020031032133123-1213112011220020"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2031312132122123-2121030213300113-1202213111200320-3230112321322101-2201031000211321-1020003321312133-3233303111003332-3210033322011222"></a>

<a id="canonical-1111021122133120-1300100202300011-1123300301131212-0002110221311103-3030203111331210-0021013101232333-2123330223102132-2113311222211213"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300122230111202-2232322101032031-2012031232002231-3212101023303003-1231101131332320-2303301011231221-0311013221013232-2222110310033132"></a>

<a id="canonical-0032310122020112-2330023301133210-2121333231111111-0222231223113002-1100021003121202-0013130231301300-0233302332022113-0332212203310021"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3332003330020233-1321203101311112-1032320222112212-0303123303022320-0130233233212212-2331122210022100-2012123323130120-2300212333133331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-2023232133013100-2200211232213000-2201011220303123-1320133103031123-3102112121111312-2220023103310102-3100313313032210-1101013010020132"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

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
selected_repos {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230001300133222-1222320301010301-3213330003233221-1111010030212233-1203312011200033-3231112333333121-0331210210121000-3210122110100203"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos`

<a id="canonical-0303003011020122-3333212220000322-2133133022133233-1221133210312310-0022110323330303-2231201103001302-2333311331321221-2022031032102221"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo` property

Type: `["list", "string"]`. Optional.

Code repository which contain API endpoints.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-2000000012223230-0211010000233313-1010013330320313-3003002303303013-3003012223220213-3332302231300133-2301220230113223-0113213212221220"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

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
custom_api_auth_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231221221302103-0211202002021321-2323021110133323-2303023132001020-2022202331221132-3322011233021212-1211030012303320-2102020231112210"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery`

- [api_discovery_ref](resources--cdn_loadbalancer--reference--group-010.md#canonical-1032323002010020-1022331203312301-0223203013030002-1003233130313211-0303313013223302-0222322103303123-0212103001230001-1121320112232330): complete subsection reference.

<a id="canonical-1032323002010020-1022331203312301-0223203013030002-1003233130313211-0303313013223302-0222322103303123-0212103001230001-1121320112232330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-2100223211110032-0022232233231102-3002001313210133-2120130303331110-3000320010313301-2113212302322012-0110123000000132-0013123202113003"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212101222033110-2113330102332322-3101332221113302-3020132121331313-1330201030121130-2221123003122231-0303023031102033-0213121133101020"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref`

<a id="canonical-1202013332032113-1311221103102202-0313222201100110-0220120112122133-1210103221112011-3121121213013331-3220132123302120-1002201232013312"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0211313023320013-0302112230002111-0003010130213213-1233001202103230-0323302221323222-3312132022011001-2003332110301021-2100211103311310"></a>

<a id="canonical-1321203231133220-1010000332303122-3122033002303320-2122132212222221-0113311001011012-2022220312110100-0200132002213202-1121032202300320"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1230233100131211-1010220312033333-2312032321233232-1303012123221001-0333021220203130-3333013233232321-1022100113322231-0102233302231111"></a>

<a id="canonical-3103110223120001-3220122002332200-3000230231202213-0311022300200313-2011231100133231-3110203233000211-3313100131331002-1313031212121213"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1013221211221310-1031120002332131-0232001133000302-3000130332131232-1303121022332312-3012002312133111-1112121330233321-2233231122001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.default_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-3330312033101230-0332222200233332-0101032312130121-0031230222130131-0023200000201230-1103120303010303-1320313222002032-1020121123313123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_api_auth_discovery = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331120230102020-0312102032323311-1020022132220113-1312101223301033-0301110002231123-3321223213133301-2033313330221330-1323102320001301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.disable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-1013231022323130-0133022002100312-2200002332003020-0032222010313033-3210300023103030-2312320031121232-0001223313301012-0033220122221221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

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
disable_learn_from_redirect_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112030012103111-1212023301022210-3131303031130201-1232223310233020-1021210103230301-2110223201003332-1112033210222011-3100331330221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.discovered_api_settings

<a id="canonical-3211302123213233-2213103023232312-1322230201223031-3113130221113323-3213022012202033-3103002213022020-0110212023330223-1131133203203212"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

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
discovered_api_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221232302112131-0300300230320213-2213311201330311-2230032033031110-1112102323131210-1222333220010223-2032022332301002-2013112131112023"></a>

### Direct properties for `enable_api_discovery.discovered_api_settings`

<a id="canonical-0000013231022122-2030002133013321-3310220211312000-2200131132100102-3003021310122113-0002122013323233-1302113121220031-1110220311213120"></a>

#### `enable_api_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-0021233123301122-1013002210302002-1231011123110332-0101131021133112-1321303230003032-3032101231010222-0000133020332203-2201003101132010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.enable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-1330012322210220-2231101001022001-3233031123022210-0230202013013011-2230333011002231-3220333302031202-1002320112023320-1231321311011222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

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
enable_learn_from_redirect_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_challenge

<a id="canonical-1201312312020123-3230111211203203-0323103120212310-2002022312313232-2200320302313111-0002122232302200-2233200202311231-0212233210001213"></a>

Type: `"object"`. single nested block, Optional.

Configure auto mitigation i.e risk based challenges for malicious users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

Terraform syntax:

```terraform
enable_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312320100311220-2121111023223010-2222123023231000-3231230201333232-2000321013003312-3122121312131333-2311101310003322-3032200103202113"></a>

### Direct properties for `enable_challenge`

- [captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012112303301110-3222223022111121-0322323321333132-0233230321301001-3011120230103011-1321123300033331-0013312223020331-0102200200231012): complete subsection reference.

- [default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-2212120120103301-3233031210133001-0010332200203330-3033001222120213-3112132200312113-1322100201132022-0200002302133322-0333302223330212): complete subsection reference.

- [default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121110313000012-3310321102200200-3113100211012323-3003132311310212-1230221132012011-1002211112133301-0031330011112203-0302213312133110): complete subsection reference.

- [default_mitigation_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-3003123010123300-3301100122332312-3002011132220020-2333111001110100-1212100130022031-2213210313032210-0121003122010333-3322202033313003): complete subsection reference.

- [js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-0312101012232110-0001011211003323-1030010110111113-0310302333212012-0111321322111020-3021022301102212-0232230030320121-0333232231023122): complete subsection reference.

- [malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-010.md#canonical-0323300102211113-0133120131332130-3221200230202231-0232201333102322-1131003112030231-0213300330113222-2103122033011023-1113122230202203): complete subsection reference.

<a id="canonical-1012112303301110-3222223022111121-0322323321333132-0233230321301001-3011120230103011-1321123300033331-0013312223020331-0102200200231012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-0211123311231020-0003112120310203-2220013131110110-0021333120132010-1232202233030213-1131200301211130-3001200210101112-0323130031020303"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022331321200331-2101333111332103-3211120120210111-3113020131111101-1020312332331130-1213113123003213-3122021112223312-3121002300110012"></a>

### Direct properties for `enable_challenge.captcha_challenge_parameters`

<a id="canonical-0113232331333331-3301201121111010-3210223121223023-3111113102333310-0313002211120202-0033130020101121-2320101211323120-0111223312232301"></a>

#### `enable_challenge.captcha_challenge_parameters.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3311131020313213-2132233320310202-2310102223101331-2332231120032002-1013210123311021-1133200320131330-1213200133232303-0011012332001002"></a>

<a id="canonical-2112000300021300-3033223223201021-1110032013133101-2223023020331232-1131322211233020-2221102003031030-3323132100031320-0232001211311013"></a>

#### `enable_challenge.captcha_challenge_parameters.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2212120120103301-3233031210133001-0010332200203330-3033001222120213-3112132200312113-1322100201132022-0200002302133322-0333302223330212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-3333113032103032-1001030133122233-0000132133113311-3120220111203201-0123211100233312-3311212113122220-2201021103002213-2021333032202220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default captcha challenge parameters.

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
default_captcha_challenge_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121110313000012-3310321102200200-3113100211012323-3003132311310212-1230221132012011-1002211112133301-0031330011112203-0302213312133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-1210331103002233-1230202212223322-0232231320321303-3010212202200130-3332331010221113-1011112221033100-3010101032032100-0011013032203001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default js challenge parameters.

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
default_js_challenge_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003123010123300-3301100122332312-3002011132220020-2333111001110100-1212100130022031-2213210313032210-0121003122010333-3322202033313003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_mitigation_settings

<a id="canonical-0300221213130223-0123223121003201-0322310232332121-0202121230200012-2321332300132320-0223302031303011-2223230010221233-3011330222120330"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_mitigation_settings = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312101012232110-0001011211003323-1030010110111113-0310302333212012-0111321322111020-3021022301102212-0232230030320121-0333232231023122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.js_challenge_parameters

<a id="canonical-0323203300302101-1222220231002133-1303201020320013-0122301013313202-1231200110111332-2233103301103001-2123203121202012-3010322110322023"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202103312023101-2330031321110031-2331000030310023-0233233313320020-3102213330000302-2001200210100022-3001212222000230-0330031121220110"></a>

### Direct properties for `enable_challenge.js_challenge_parameters`

<a id="canonical-0033003031323131-2020101113101303-1303311130110220-2331313032022201-3310330101322221-1013302110103030-2220012103001131-3020213102210113"></a>

#### `enable_challenge.js_challenge_parameters.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-2331310112023121-3121123302133023-2311110113133302-2100311211030323-0313310322203231-3220311320200112-3300131121012320-1122202120330202"></a>

<a id="canonical-2301000310332320-3313203212013322-3012332102102101-2330020211033202-2322230222330313-2303213003023010-0321323333203030-0220010133311100"></a>

#### `enable_challenge.js_challenge_parameters.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0202132112111103-3231121320123332-0013222333311010-2020101013220303-0230203203120232-3103212103230013-1213132122223213-2233212311210300"></a>

<a id="canonical-1310230120221300-0011231200200121-2100012232320232-3111020112312102-0312132220221310-0101122022313332-0132221102302200-2301033331022203"></a>

#### `enable_challenge.js_challenge_parameters.js_script_delay` property

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-0323300102211113-0133120131332130-3221200230202231-0232201333102322-1131003112030231-0213300330113222-2103122033011023-1113122230202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.malicious_user_mitigation

<a id="canonical-3011302213322302-0112111011213203-2211103221013111-1102302231012201-2103101231221103-3211100203310131-1122133201203103-3020123030303130"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011030003301312-0323131112213330-1232013330203320-0223310012003231-2311010312330121-1213202330311231-1032001111220331-0212332130012303"></a>

### Direct properties for `enable_challenge.malicious_user_mitigation`

<a id="canonical-3230211000302300-2312302011233301-1233333123111232-0001122110132201-0231133230330302-0112222301013123-1102121201133032-3232220323103333"></a>

#### `enable_challenge.malicious_user_mitigation.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3122013203130123-3220023132222120-1101033220011321-0312333222120001-3201133220312033-1221213222320332-1221332331021321-0132312333121211"></a>

<a id="canonical-3202031212211010-1000310230001220-1222320033110102-1120112023332202-2310133002110323-1332120232200302-0213033100311013-2232331031012221"></a>

#### `enable_challenge.malicious_user_mitigation.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2322103131322131-0122312130133303-0300133222030010-1000211333211300-2121310322330221-1113023202220030-2202123212232232-1101202220020330"></a>

<a id="canonical-3000321313331123-3233131101110002-2022100221330130-2100333101131323-3023001223213330-0332320213100300-3122010313131301-1120213111323213"></a>

#### `enable_challenge.malicious_user_mitigation.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3132020110023330-0221123212220300-2303122132103021-2301032210323100-0112321332310001-2233030031121302-2100003011131031-0321200323303122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ip_reputation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_ip_reputation

<a id="canonical-1123223010010320-2010330211102303-3003132133301111-3013321110233132-3220003300022030-2113002030033003-3221112313032202-3320311130101320"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List. List of IP threat categories.

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
enable_ip_reputation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032321122230201-3300312212102222-2111201333313311-1200211223033002-0303130300310102-3012200321100321-3303211021120200-1023322223303132"></a>

### Direct properties for `enable_ip_reputation`

<a id="canonical-2222233221300223-2333220231013132-2233000201012203-2203002130020121-2110220333122330-3013100100110121-3121002111113221-3113320200211320"></a>

#### `enable_ip_reputation.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1303213031102030-3122302223322303-3002003203030013-0131022002010001-0002312200302222-3331123211021133-0230132330122210-3301210122110300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_malicious_user_detection

<a id="canonical-1123232203103010-2221320123213112-1303032131321013-2032233321303102-1212120310102022-0332230300221210-2020311101222333-2212332330213300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable malicious user detection.

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
enable_malicious_user_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301230133222123-3331320000010320-3300322323323202-1103213233103032-2103000031113100-2212030000211223-1213133100220200-0023212123220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_threat_mesh` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_threat_mesh

<a id="canonical-2123213233303030-1130113101312333-0212210331021330-3103233332320113-1210122212022213-2121031100211332-1031010102222213-2210032031223013"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_threat_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.
