---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-0102332200231322-2200202032302311-2332002033322011-3322112021130030-0301131002103321-3233232010011302-1222012102033322-1220122221203223"></a>

## Next pages — use_system_defaults / 323213311212 / 4

- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113103213311210-3302310220332300-1212131110331013-2022331010002033-1211101023020002-3210002001210102-1102220233201321-3323213100111202"></a>

## tls_parameters.common_params.validation_params — validation_params / 313321101023 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- tls_parameters.common_params.validation_params

<a id="canonical-0133133102312101-1332332333220020-0232232312232220-0021220323221021-0021132030110220-3222222000230321-3201120003120312-1321333333221100"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232100100101230-0132032000130312-3320113230102013-2002331312220222-2323312003320020-1233301303013311-0012120001123012-3210121323333133"></a>

## Direct properties — validation_params / 313321101023 / 3

<a id="canonical-2211003312223203-2030300222232211-3222310330011031-1320332023303302-0102131300203033-0232310003321230-0311211033300003-3030133333300110"></a>

<a id="canonical-0301122000100320-0233231121120110-0110313123203010-2120110120020223-3110000022121300-0000210033113223-2303313323330211-1323220330301330"></a>

## skip_hostname_verification property — validation_params / 313321101023 / 4

Type: `"bool"`. Optional.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](resources--cluster--reference--group-002.md#canonical-2123020121303021-2110101300212030-3021001233002120-3111103121230120-2023100101002103-1211222000220210-2232311111312300-2313222211011230): complete subsection reference.

<a id="canonical-1133321011103332-0130311203223102-1033303111022323-3021001010013233-3220202112100111-0323130031302131-1233211222330322-3212121011322212"></a>

<a id="canonical-0301012003323033-2210331200203201-0313323330222212-0111311232130213-1212302032132332-1212303103213333-3302323302300232-0232312230013212"></a>

## trusted_ca_url property — validation_params / 313321101023 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2110310211123103-2102100003023020-1003312030333020-0213311011332002-0113213222133001-2300230301123331-2332002030300323-1102303130322300"></a>

<a id="canonical-3303032022202211-0130331330011030-3202003311333113-0020101320111022-0301011321213201-2323120033003332-2011321320022322-1212232120120102"></a>

## verify_subject_alt_names property — validation_params / 313321101023 / 6

Type: `["list", "string"]`. Optional.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-1101122310310133-2103100222330311-2323232103102331-1113201102213123-0122311231203023-0203232101230031-2303023110000022-2013322200001030"></a>

## Next pages — validation_params / 313321101023 / 7

- [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-002.md#canonical-2123020121303021-2110101300212030-3021001233002120-3111103121230120-2023100101002103-1211222000220210-2232311111312300-2313222211011230)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2123020121303021-2110101300212030-3021001233002120-3111103121230120-2023100101002103-1211222000220210-2232311111312300-2313222211011230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111030002330112-0220032202213230-2221230332110202-2121011102111112-1212031011022120-1210330103131020-0200331222010213-3003201323303000"></a>

## tls_parameters.common_params.validation_params.trusted_ca — trusted_ca / 303100013122 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-002.md#canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-0113001301213033-0313332111203011-3221223322213132-0022130301130121-2320330012322003-3102000323101202-2002332201000203-2302030023233101"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333303310103310-1112202320231111-1323232023101331-3233220232200121-0113010213110033-3011232322112022-3323121113133022-2013230020312012"></a>

## Direct properties — trusted_ca / 303100013122 / 3

- [trusted_ca_list](resources--cluster--reference--group-002.md#canonical-2031332030331232-1103121133023121-1011133313302011-2210130002320220-3121303200312203-2233202100102010-1012212230111202-0333000000121000): complete subsection reference.

<a id="canonical-2032332101312003-1331012203220201-3321133012031332-0310012020323102-2022002121310101-1120320130033122-3203000223100321-1203130002230011"></a>

## Next pages — trusted_ca / 303100013122 / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-002.md#canonical-2031332030331232-1103121133023121-1011133313302011-2210130002320220-3121303200312203-2233202100102010-1012212230111202-0333000000121000)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-002.md#canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2031332030331232-1103121133023121-1011133313302011-2210130002320220-3121303200312203-2233202100102010-1012212230111202-0333000000121000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110332230020023-3122002023222023-1303001331113302-3130001320100030-0032200103100332-2110023113303022-0322110210011212-1112200200021323"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 302032020021 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-002.md#canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103)
- [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-002.md#canonical-2123020121303021-2110101300212030-3021001233002120-3111103121230120-2023100101002103-1211222000220210-2232311111312300-2313222211011230)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-1313122211012132-2323103001001011-0212010311201302-1221103302011231-3213123332332012-1223113322323302-0103033010333230-2311102132201233"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121103323201301-0311211010312120-0130311023323321-0213210203211313-1232003320112010-0302201032121132-2002020132030010-0311230330312212"></a>

## Direct properties — trusted_ca_list / 302032020021 / 3

<a id="canonical-3012103321133010-1133133303032000-1031310302000021-1311033132220110-1221023310030311-2002232310120001-3320023323313130-2101322023022232"></a>

<a id="canonical-3032221332033021-0031111312101212-0302233032032031-0131131013320003-1311330220221312-2203220103321131-2302213322132223-1022011213202123"></a>

## kind property — trusted_ca_list / 302032020021 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
  }
}
```

<a id="canonical-2101333123213122-0023033112202223-1013022330210321-0102133302312312-0201223303312123-3012032220321220-0232010010213310-2302203210300010"></a>

<a id="canonical-1231100121120022-3223203313310023-1000131301103312-2021001222321201-1111332131300100-3102212333302121-2120300103223001-1012123020200301"></a>

## name property — trusted_ca_list / 302032020021 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1000103232230201-0210021202300100-3030102033020013-0331201222231322-2311122312303331-3202022113301330-0011120131201310-2032023101211310"></a>

<a id="canonical-0330203331233030-2232313131123310-1103310130210321-1132030000112110-0300033123203102-0020022002312033-2333130123330002-1220111330031333"></a>

## namespace property — trusted_ca_list / 302032020021 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-0122102332323321-1130221101222111-3212312033312222-2002302321330120-3101201003302113-3033232002312110-1320322320030312-0122210312110213"></a>

<a id="canonical-3233110031132230-0023133033322320-1023211322300023-0101323110211331-2223102331001112-3301310000033100-3011031002330230-3303122311222230"></a>

## tenant property — trusted_ca_list / 302032020021 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3303133130220113-1232100232231223-3023230300311322-3122002310312030-1310232312133203-0031002020131301-3222202323233310-0333312110012122"></a>

<a id="canonical-2320221121100322-3202100100211011-1110203213130331-0312331020133101-2330120111122233-0213010030310321-2313222131203231-3211110333123330"></a>

## uid property — trusted_ca_list / 302032020021 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0222200212201003-0000201111030321-0112102222313131-0330332032212230-1133333000200210-1313301323210131-0321130011301320-0012112310202230"></a>

## Next pages — trusted_ca_list / 302032020021 / 9

- [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-002.md#canonical-2123020121303021-2110101300212030-3021001233002120-3111103121230120-2023100101002103-1211222000220210-2232311111312300-2313222211011230)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3111210320132103-3023301033212102-3010212200103223-1330312002311102-2203031033111312-1310030010122203-0223131200231111-2130111222200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202231323313000-3303122020230223-1030323023203002-3101001011300220-0002321200222221-0222313023223102-3001023030203133-2212223202313012"></a>

## tls_parameters.default_session_key_caching — default_session_key_caching / 101323211120 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.default_session_key_caching

<a id="canonical-0303022213110012-1300303123201131-3303103203230323-0022123313320131-0210210011130022-3223230022013201-1311321323330233-1010212312123302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default session key caching.

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
default_session_key_caching = {}
```

<a id="canonical-1022001222123310-3120213310101303-3030203000212133-1103222133233200-1133022313023232-1310331122022122-0120312022310110-2002222203230010"></a>

## Direct properties — default_session_key_caching / 101323211120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233130121223222-2312112201332021-3023313012231303-1222201123003120-1212120103323310-3102000000313020-3033112331313012-0301112031013123"></a>

## Next pages — default_session_key_caching / 101323211120 / 4

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2102323013232123-2122130023210330-2223011133213113-3010120223212310-3033220231110101-2003103300202020-3311031203312302-0210012033113310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033101102020320-3033112001113303-2212231121123330-1213303123020311-1332202013131101-2230102212212320-3231130000322113-3020003002031222"></a>

## tls_parameters.disable_session_key_caching — disable_session_key_caching / 013010002122 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.disable_session_key_caching

<a id="canonical-2303301023023213-3311333113121301-1002223201031203-1332130331003033-0322113131301333-3013101131220232-3233111300232102-0233122312021011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

<a id="canonical-2102210133332131-3132303002213211-1201211103211312-0102310333330133-2313201033001110-2031101301013113-1211011213211100-1031332302313102"></a>

## Direct properties — disable_session_key_caching / 013010002122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032033033130311-2002211200231122-2000302302212120-2302123330123201-2220302113211223-3032020233231321-3022021321121232-0213022102202031"></a>

## Next pages — disable_session_key_caching / 013010002122 / 4

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1231201030310032-2232133130332230-3102203310013132-0231112133233202-2112002222030130-0312121223003313-1302023203330231-0310011232322132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010220020102102-1322132201232303-2002223103222331-2001331322132300-2211211323032103-2002310121033223-2120100233320331-0310230011001330"></a>

## tls_parameters.disable_sni — disable_sni / 301233110123 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.disable_sni

<a id="canonical-3011031101302320-2132321032112032-3201030210032322-0003233113001212-0133230213031010-0031313020011233-0001000333100302-2003230021320213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

<a id="canonical-0111332333321011-2321313310310032-3311132203232201-1131202103212222-1130101122032120-2110132201212231-1213221123130200-2220111202013210"></a>

## Direct properties — disable_sni / 301233110123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211321201323330-2212320100000032-3122313013322033-1220100122112201-3321100310000103-0022121220333110-2320113320130111-2003231120332023"></a>

## Next pages — disable_sni / 301233110123 / 4

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3303030113320302-0313223030232122-1230002312320323-0112201033032021-3032303300003222-3030123202200331-3032022101232333-0330331332030030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022321303021303-0301221310122020-2011030023032130-1110101130103201-2001112123212033-0300301131021033-1113003001000113-0221222031203003"></a>

## tls_parameters.use_host_header_as_sni — use_host_header_as_sni / 233123212133 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.use_host_header_as_sni

<a id="canonical-2012212221130301-0303332301113230-3112332122220113-2220031222031121-0110221211330111-3111022033002232-2012200203321233-2010130310233323"></a>

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
use_host_header_as_sni = {}
```

<a id="canonical-2312110321002311-0231210333012211-3012302100232111-0002302310210303-0102132133103301-1001330331300311-1113033110232032-1210131330332302"></a>

## Direct properties — use_host_header_as_sni / 233123212133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323222332001032-1030031311221120-3100231010131013-3132003032232323-2222212121333133-3031103220332303-0122210322000012-3022031223023231"></a>

## Next pages — use_host_header_as_sni / 233123212133 / 4

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320331001230102-2231121201220333-1110302213302303-0233320111111023-3233022310230302-2100003101020110-0131021233211200-1103130213111101"></a>

## upstream_conn_pool_reuse_type — upstream_conn_pool_reuse_type / 312012113311 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- upstream_conn_pool_reuse_type

<a id="canonical-2123332212301000-2200010021323223-2330213113021121-0221021110000130-0001101030103313-0133101111321103-0322301331033113-0121000000212220"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
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
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313000320131113-3112330121011222-1330003313220221-2220222033133032-2302031313231233-1223122332303011-0313211111112220-1122331313130331"></a>

## Direct properties — upstream_conn_pool_reuse_type / 312012113311 / 3

- [disable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-3323023220002122-1301220010310231-1031013200230113-2100310231030123-2113121100122000-2033203013032233-0220222232110011-1222231110321022): complete subsection reference.

- [enable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-2303230031132101-0221312032202231-3131333202122000-0321210220013310-2012022020012310-2203211202111212-1223313210332112-0131300233211203): complete subsection reference.

<a id="canonical-3033220130113322-0003013223010210-1312132021313133-1101333332101323-1013321113113001-2330133220220223-1221310021102110-0203222330012122"></a>

## Next pages — upstream_conn_pool_reuse_type / 312012113311 / 4

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-3323023220002122-1301220010310231-1031013200230113-2100310231030123-2113121100122000-2033203013032233-0220222232110011-1222231110321022)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-2303230031132101-0221312032202231-3131333202122000-0321210220013310-2012022020012310-2203211202111212-1223313210332112-0131300233211203)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3323023220002122-1301220010310231-1031013200230113-2100310231030123-2113121100122000-2033203013032233-0220222232110011-1222231110321022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331321112123331-1332030333213232-1111222321030020-1203102300232311-1230223111230100-0023000210110103-0113200231201313-1230022032130231"></a>

## upstream_conn_pool_reuse_type.disable_conn_pool_reuse — disable_conn_pool_reuse / 232000002212 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-0313323233003322-0302021233211011-3111230310312230-3103300120210113-0212200233332101-0102223123320220-0331313020002320-3113233303110330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

<a id="canonical-3212201212031313-2232020313210220-3330031321003123-0233020131222110-0301203222103223-1120320312220000-0333211313112321-0233012100011133"></a>

## Direct properties — disable_conn_pool_reuse / 232000002212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103313001013100-0000012002002313-1233100123010302-2003330103301230-0111120033012110-3000110111230103-0112122131333000-1130333330331111"></a>

## Next pages — disable_conn_pool_reuse / 232000002212 / 4

- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2303230031132101-0221312032202231-3131333202122000-0321210220013310-2012022020012310-2203211202111212-1223313210332112-0131300233211203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132221022131111-0010303311101322-1313012123333201-1102223200221003-3321001020331032-3313302212131230-0311003322123313-3320221012002233"></a>

## upstream_conn_pool_reuse_type.enable_conn_pool_reuse — enable_conn_pool_reuse / 020033333221 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-2121321022333202-2233112003101001-3233223201112031-1201321001032030-2020012003332033-3313010300213130-2123000233333301-0313003330003133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

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
enable_conn_pool_reuse = {}
```

<a id="canonical-1122021033020223-3230332110130002-2323110000110212-3120032132322031-1230322323100113-2230013033323011-3301132131310011-0110222310322020"></a>

## Direct properties — enable_conn_pool_reuse / 020033333221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000112230031010-1232012020332303-1132013103001201-0213211110120023-2030323030232023-1220303001122312-2320133203103122-0022210221223310"></a>

## Next pages — enable_conn_pool_reuse / 020033333221 / 4

- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
