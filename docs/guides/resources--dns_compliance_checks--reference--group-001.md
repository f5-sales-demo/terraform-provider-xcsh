---
page_title: "xcsh_dns_compliance_checks reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks reference."
---

# xcsh_dns_compliance_checks reference

<a id="canonical-3232320123103113-0100311102312313-0121302010013132-2312320000130313-1222132032330322-2302213002320222-1233300132211231-3112210312300202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103012302212203-2323121220101300-0000303332032001-2300121301100100-1332002103112022-1230230320221312-0122022103123021-3113100320331232"></a>

## Property reference — Property reference / 300312012031 / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210)
- Property reference

<a id="canonical-1313232132200010-2101203123102102-2200331101333110-1103010203202321-3121123310313113-2103101210323200-1010223110030101-2001201112210313"></a>

## Direct properties — Property reference / 300312012031 / 3

<a id="canonical-2103310213221120-0212111312320110-3230033011120113-3233223302332021-2002003020110223-1000202301020110-0011023101023031-3100023300302003"></a>

<a id="canonical-3130201220001201-2313002022021201-0113010121230100-2301301300220001-3032022223023113-1032320301312200-3033132132220132-0121220010203031"></a>

## annotations property — Property reference / 300312012031 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
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

<a id="canonical-0131010331220112-1013100200002020-0233110030302101-2001301311322001-3333031300111013-1031331302231300-2002203031021212-0332211331000220"></a>

<a id="canonical-1033200003333201-2223200222112132-3132200211002223-3023323311320221-0103303102001122-0103032012220322-3101201122101130-2123223313132101"></a>

## description property — Property reference / 300312012031 / 5

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

<a id="canonical-1321000001302211-3111212231231021-2110010113222000-3303201200313300-3233102330322021-0203013203030231-3103330322123211-0223021110030003"></a>

<a id="canonical-2111021222212130-1312033223003113-1330210210302002-3021332210220021-1130311002013002-0123133330133311-3010010320031302-3123113321112322"></a>

## disable property — Property reference / 300312012031 / 6

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

<a id="canonical-0132303132202032-1302312200111200-1220322121122131-2000111023122020-3032223120312133-1001030122131101-2333031302321131-2233012110112213"></a>

<a id="canonical-2121100033323220-1221332121210010-2222010321223200-2312213011031110-3330103011011122-3011001313301212-2220222111122231-1123301213132023"></a>

## disallowed_query_type_list property — Property reference / 300312012031 / 7

Type: `["list", "string"]`. Optional.

\[Enum: QUERY|IQUERY|STATUS|NOTIFY|UPDATE\] Disallowed Query Type Values. Disallowed Query Type
Values. Possible values are \`QUERY\`, \`IQUERY\`, \`STATUS\`, \`NOTIFY\`, \`UPDATE\`. Defaults to
\`QUERY\`.

Upstream description:

Disallowed Query Type Values.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1133030122121032-3003330202300200-2111111032010110-1003311002121132-2102231122320132-1130300300033212-0022320333003031-0233121232112023"></a>

<a id="canonical-1320312101300123-1221133233300031-3231031012000101-0233022013322101-1330101313130312-1311211310200222-1233203313330131-2323333210033201"></a>

## disallowed_resource_record_type_list property — Property reference / 300312012031 / 8

Type: `["list", "string"]`. Optional.

\[Enum:
T|A|NS|MD|MF|CNAME|SOA|MB|MG|MR|NULL|WKS|PTR|HINFO|MINFO|MX|TXT|RP|AFSDB|X25|ISDN|RT|NSAP|NSAP\_PTR|SIG|KEY|PX|GPOS|AAAA|LOC|NXT|EID|NIMLOC|SRV|ATMA|NAPTR|KX|CERT|A6|DNAME|SINK|OPT|APL|DS|SSHFP|IPSECKEY|RRSIG|NSEC|DNSKEY|DHCID|NSEC3|NSEC3PARAM|TLSA|SMIMEA|HIP|NINFO|RKEY|TALINK|CDS|CDNSKEY|OPENPGPKEY|CSYNC|SPF|UINFO|UID|GID|UNSPEC|NID|L32|L64|LP|EUI48|EUI64|TKEY|TSIG|IXFR|AXFR|MAILB|MAILA|URI|CAA|TA|DLV\]
Disallowed Resource Record Types. Disallowed Resource Record Type List. Possible values are \`T\`,
\`A\`, \`NS\`, \`MD\`, \`MF\`, \`CNAME\`, \`SOA\`, \`MB\`, \`MG\`, \`MR\`, \`NULL\`, \`WKS\`,
\`PTR\`, \`HINFO\`, \`MINFO\`, \`MX\`, \`TXT\`, \`RP\`, \`AFSDB\`, \`X25\`, \`ISDN\`, \`RT\`,
\`NSAP\`, \`NSAP\_PTR\`, \`SIG\`, \`KEY\`, \`PX\`, \`GPOS\`, \`AAAA\`, \`LOC\`, \`NXT\`, \`EID\`,
\`NIMLOC\`, \`SRV\`, \`ATMA\`, \`NAPTR\`, \`KX\`, \`CERT\`, \`A6\`, \`DNAME\`, \`SINK\`, \`OPT\`,
\`APL\`, \`DS\`, \`SSHFP\`, \`IPSECKEY\`, \`RRSIG\`, \`NSEC\`, \`DNSKEY\`, \`DHCID\`, \`NSEC3\`,
\`NSEC3PARAM\`, \`TLSA\`, \`SMIMEA\`, \`HIP\`, \`NINFO\`, \`RKEY\`, \`TALINK\`, \`CDS\`,
\`CDNSKEY\`, \`OPENPGPKEY\`, \`CSYNC\`, \`SPF\`, \`UINFO\`, \`UID\`, \`GID\`, \`UNSPEC\`, \`NID\`,
\`L32\`, \`L64\`, \`LP\`, \`EUI48\`, \`EUI64\`, \`TKEY\`, \`TSIG\`, \`IXFR\`, \`AXFR\`, \`MAILB\`,
\`MAILA\`, \`URI\`, \`CAA\`, \`TA\`, \`DLV\`. Defaults to \`T\`.

Upstream description:

Disallowed Resource Record Type List.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3122230132322300-2020220032002221-0021230010301333-1330113232013223-1023101223202232-3010302002120020-3232002230013312-2110321200113313"></a>

<a id="canonical-3323003232231023-1212011210012331-3031032312133010-3112211232003133-3320100111132022-1003111030032030-0302020303323210-3221123011133313"></a>

## domain_denylist property — Property reference / 300312012031 / 9

Type: `["list", "string"]`. Required.

List of domains to be denied by configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(30),
}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.etld_plus_one": "true",
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.etld_plus_one": "true",
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1300332031233211-1312010200012131-1331000300012102-1310231320323211-0230330213333310-2212321022212302-1310212221131121-3213203000322332"></a>

<a id="canonical-3221103133012321-2320302031333313-2331113220302012-2201011232020301-0132112003322213-0011231203330010-3012000033100203-2022222113210122"></a>

## ID property — Property reference / 300312012031 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3113113222120021-1331333322113120-2033011323310301-0131322231203101-1102233022301220-0121232230033030-1300012123122030-3131203210000202"></a>

<a id="canonical-1021203302331233-1220303311011313-2212222021331200-1223311111332103-1120020111330331-3312011023022221-0103210222202233-3102331321302310"></a>

## labels property — Property reference / 300312012031 / 11

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
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

<a id="canonical-1032301213310120-2020100231030222-0332010303200112-3332103121202033-2233011231022200-1200111032203323-1203300203002031-3110303202231322"></a>

<a id="canonical-2013313202221130-1202332002112102-2101120331011030-2120311111131233-2231300100331311-1332310131022133-2123122212210201-3100201220003221"></a>

## name property — Property reference / 300312012031 / 12

Type: `"string"`. Required.

Name of the DNS Compliance Checks. Must be unique within the namespace.

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

<a id="canonical-0312010333100130-2232000230030113-0013123133120123-1333112113112230-3022001313111302-0233001111130223-2130202310322002-0230311222131131"></a>

<a id="canonical-2201132212133300-0330331021203033-2303322203100000-2133130123333231-0013013001002300-1300132010013323-2123223330002222-1311000231021133"></a>

## namespace property — Property reference / 300312012031 / 13

Type: `"string"`. Required.

Namespace where the DNS Compliance Checks is created.

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

- [timeouts](resources--dns_compliance_checks--reference--group-001.md#canonical-0121213120233133-1312330233101232-1322031333323222-1212200121211103-0022222020113321-3121131210333132-3013302012320332-1200122023220122): complete subsection reference.

<a id="canonical-0213121200102321-0202020203230203-2033101321231013-3310121030312020-2312001012121000-3332000320012030-3213211230023300-3012132203022213"></a>

## All schema paths — Property reference / 300312012031 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_compliance_checks--reference--group-001.md#canonical-2103310213221120-0212111312320110-3230033011120113-3233223302332021-2002003020110223-1000202301020110-0011023101023031-3100023300302003) |
| `description` | [description](resources--dns_compliance_checks--reference--group-001.md#canonical-0131010331220112-1013100200002020-0233110030302101-2001301311322001-3333031300111013-1031331302231300-2002203031021212-0332211331000220) |
| `disable` | [disable](resources--dns_compliance_checks--reference--group-001.md#canonical-1321000001302211-3111212231231021-2110010113222000-3303201200313300-3233102330322021-0203013203030231-3103330322123211-0223021110030003) |
| `disallowed_query_type_list` | [disallowed_query_type_list](resources--dns_compliance_checks--reference--group-001.md#canonical-0132303132202032-1302312200111200-1220322121122131-2000111023122020-3032223120312133-1001030122131101-2333031302321131-2233012110112213) |
| `disallowed_resource_record_type_list` | [disallowed_resource_record_type_list](resources--dns_compliance_checks--reference--group-001.md#canonical-1133030122121032-3003330202300200-2111111032010110-1003311002121132-2102231122320132-1130300300033212-0022320333003031-0233121232112023) |
| `domain_denylist` | [domain_denylist](resources--dns_compliance_checks--reference--group-001.md#canonical-3122230132322300-2020220032002221-0021230010301333-1330113232013223-1023101223202232-3010302002120020-3232002230013312-2110321200113313) |
| `id` | [ID](resources--dns_compliance_checks--reference--group-001.md#canonical-1300332031233211-1312010200012131-1331000300012102-1310231320323211-0230330213333310-2212321022212302-1310212221131121-3213203000322332) |
| `labels` | [labels](resources--dns_compliance_checks--reference--group-001.md#canonical-3113113222120021-1331333322113120-2033011323310301-0131322231203101-1102233022301220-0121232230033030-1300012123122030-3131203210000202) |
| `name` | [name](resources--dns_compliance_checks--reference--group-001.md#canonical-1032301213310120-2020100231030222-0332010303200112-3332103121202033-2233011231022200-1200111032203323-1203300203002031-3110303202231322) |
| `namespace` | [namespace](resources--dns_compliance_checks--reference--group-001.md#canonical-0312010333100130-2232000230030113-0013123133120123-1333112113112230-3022001313111302-0233001111130223-2130202310322002-0230311222131131) |
| `timeouts` | [timeouts](resources--dns_compliance_checks--reference--group-001.md#canonical-2120220002012231-2303213222000013-2023321232203133-2301230311222023-2332123131123020-3232213033321132-0320103223133031-0113220303201321) |
| `timeouts.create` | [timeouts.create](resources--dns_compliance_checks--reference--group-001.md#canonical-1010012001011310-2021020030301022-0001003302000322-0101020233332131-1210000332200310-3113130202113321-3101301032120010-3322202330211221) |
| `timeouts.delete` | [timeouts.delete](resources--dns_compliance_checks--reference--group-001.md#canonical-1303101100111302-3303000323113321-1000312201211322-1010231312323232-0311110210302203-2202130311023002-2200233311110322-0311322112302022) |
| `timeouts.read` | [timeouts.read](resources--dns_compliance_checks--reference--group-001.md#canonical-3322032302032332-0022102123123322-0302323202223010-3023020110313332-2311132330101021-3021231100301232-0210031322312312-2010303330010303) |
| `timeouts.update` | [timeouts.update](resources--dns_compliance_checks--reference--group-001.md#canonical-3311103303302302-3003131121231112-2131012023203313-2101313320202211-3301233033123000-0133030123123002-0010321302002212-1100131322101211) |

<a id="canonical-1113131313210010-3203321001102130-0013302310123001-2323320032230330-3130201131112330-3011233321021103-3022132203001313-1020310321303200"></a>

## Next pages — Property reference / 300312012031 / 15

- [timeouts](resources--dns_compliance_checks--reference--group-001.md#canonical-0121213120233133-1312330233101232-1322031333323222-1212200121211103-0022222020113321-3121131210333132-3013302012320332-1200122023220122)
- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210)

<a id="canonical-0121213120233133-1312330233101232-1322031333323222-1212200121211103-0022222020113321-3121131210333132-3013302012320332-1200122023220122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201133222000232-3212231132233130-1033322020322201-2123202011323022-3120123003030232-2010002233202223-2012323011031200-1133031132202112"></a>

## timeouts — timeouts / 122001201223 / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210)
- [Property reference](resources--dns_compliance_checks--reference--group-001.md#canonical-3232320123103113-0100311102312313-0121302010013132-2312320000130313-1222132032330322-2302213002320222-1233300132211231-3112210312300202)
- timeouts

<a id="canonical-2120220002012231-2303213222000013-2023321232203133-2301230311222023-2332123131123020-3232213033321132-0320103223133031-0113220303201321"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102103330310110-2131301231310202-0130300022013031-2132303211030231-1103202230221230-0122120312103100-2103101311201032-3231013100312323"></a>

## Direct properties — timeouts / 122001201223 / 3

<a id="canonical-1010012001011310-2021020030301022-0001003302000322-0101020233332131-1210000332200310-3113130202113321-3101301032120010-3322202330211221"></a>

<a id="canonical-3321200021132110-1201122011022232-1111002231011310-3103103212020302-2100102122220200-1030132213332303-2333302010223132-1220112223212310"></a>

## create property — timeouts / 122001201223 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1303101100111302-3303000323113321-1000312201211322-1010231312323232-0311110210302203-2202130311023002-2200233311110322-0311322112302022"></a>

<a id="canonical-2131103010002021-2221333110331101-1022033333001011-3132322322133022-2211111310102201-3322312101313023-2332113330022023-2330003033101231"></a>

## delete property — timeouts / 122001201223 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3322032302032332-0022102123123322-0302323202223010-3023020110313332-2311132330101021-3021231100301232-0210031322312312-2010303330010303"></a>

<a id="canonical-0002231132100101-1203320212230300-3010032231001202-0323010003211120-3120021111331032-0302003001232132-3300102203332203-1312213101021132"></a>

## read property — timeouts / 122001201223 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3311103303302302-3003131121231112-2131012023203313-2101313320202211-3301233033123000-0133030123123002-0010321302002212-1100131322101211"></a>

<a id="canonical-3112200202101230-3103121313302020-0021220103132330-1221330230322130-2212331231131030-0112312003121302-1110121211303112-0113313223020230"></a>

## update property — timeouts / 122001201223 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3001230331202322-3133123330322210-1000313133111033-2031033101123110-3122013313201320-2221332002221023-1202033333013122-2211010030003001"></a>

## Next pages — timeouts / 122001201223 / 8

- [Property reference](resources--dns_compliance_checks--reference--group-001.md#canonical-3232320123103113-0100311102312313-0121302010013132-2312320000130313-1222132032330322-2302213002320222-1233300132211231-3112210312300202)
- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210)
